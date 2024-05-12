package memlock

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strings"
	"sync"
	"time"
)

type ReentrantLock interface {
	LockKeysInMemory(entry LockEntry)
	UnlockKeysInMemory(entry LockEntry)
	WithOpts(opts ...LockOpt) // will hold lock
}

type LockOpt func(rl *reentrantLock)

// LockDebugFunc only called on locking succeed or unlocking
type LockDebugFunc func(action, key, id string, lockedTimes int)
type ClearDebugFunc func(keys []string)

func WithLockDebugFunc(lckDebugFunc LockDebugFunc) LockOpt {
	return func(rl *reentrantLock) {
		rl.lckDebugFunctions = append(rl.lckDebugFunctions, lckDebugFunc)
	}
}
func WithClearDebugFunc(clrDebugFunc ClearDebugFunc) LockOpt {
	return func(rl *reentrantLock) {
		rl.clrDebugFunctions = append(rl.clrDebugFunctions, clrDebugFunc)
	}
}

func WithLockClearSec(clearDurationSeconds int) LockOpt {

	seconds := time.Duration(clearDurationSeconds)
	return func(rl *reentrantLock) {
		// <= 0 means never do clear
		if seconds <= 0 {
			rl.clearPeriodDur = -1
			return
		}
		// 如果 > 10年，可能调用方错误设置`seconds`单位为`nanoseconds`
		//   但如果原计划设置 < 315.36ms(3600*24*365*10/1000_1000_1000)  时设置错单位，这里不会修正，会导致后续设置错误
		//   eg: 计划设置300ms清理，按ns传值为 300_000_000，此处按s处理，结果约为9.5年，会以9.5为周期进行清理，避免真需要9年这种清理周期
		if seconds > 3600*24*365*10 {
			// must set seconds as nanoseconds : fix it
			seconds = math.MaxInt64 / time.Second
		}
		rl.clearPeriodDur = seconds * time.Second
	}
}

func NewReentrantLock(opts ...LockOpt) ReentrantLock {
	rl := &reentrantLock{
		locksMap:       make(map[string]*lockItem, 100),
		locksMapMux:    sync.Mutex{},
		needClearTime:  time.Now(),
		clearPeriodDur: -1,
	}
	for _, opt := range opts {
		opt(rl)
	}
	return rl
}

type reentrantLock struct {
	locksMap       map[string]*lockItem
	locksMapMux    sync.Mutex
	needClearTime  time.Time
	clearPeriodDur time.Duration

	lckDebugFunctions []LockDebugFunc
	clrDebugFunctions []ClearDebugFunc
}

func (rl *reentrantLock) LockKeysInMemory(entry LockEntry) {

	rl.lockInMemoryWithAction(entry, func(lock *lockItem) bool {
		lock.lock()
		defer lock.unlock()
		if lock.owner == entry.RequestID || lock.cnt == 0 {
			lock.cnt++
			lock.owner = entry.RequestID
			rl.runLckDebugFunctions("locked", lock.key, lock.owner, lock.cnt)
			return true
		}
		lock.wait()
		return false
	})
}

func (rl *reentrantLock) UnlockKeysInMemory(entry LockEntry) {

	rl.lockInMemoryWithAction(entry, func(lock *lockItem) bool {
		lock.lock()
		defer lock.unlock()
		if lock.owner != entry.RequestID {
			panic(fmt.Sprintf("try[id:%s] unlock a not holding lock[%s]", entry.RequestID, lock.owner))
		}
		// id equal
		if lock.cnt <= 0 {
			panic(fmt.Sprintf("try[id:%s] unlock a not exist lock", entry.RequestID))
		}
		lock.cnt--
		rl.runLckDebugFunctions("unlocked", lock.key, lock.owner, lock.cnt)
		return true
	})
}

func (rl *reentrantLock) lockInMemoryWithAction(entry LockEntry, checkLockAndDo func(*lockItem) bool) {
	keys := entry.Keys
	if !slices.IsSortedFunc(keys, func(a, b string) int { return strings.Compare(a, b) }) {
		sort.Strings(keys)
	}

	if len(keys) == 0 {
		return
	}

	toLocks := make([]*lockItem, 0, len(keys))
	rl.locksMapMux.Lock()

	// do clear first
	rl.freeNotUsingLocks()
	// then get lock
	for _, key := range keys {
		rel, exist := rl.locksMap[key]
		if !exist {
			rel = &lockItem{
				cond:  NewCond(),
				cnt:   0,
				owner: "",
				key:   key,
			}
			rl.locksMap[key] = rel
		}
		toLocks = append(toLocks, rel)
	}
	rl.locksMapMux.Unlock()
	for _, lock := range toLocks {
		// lock
		for !checkLockAndDo(lock) {
		}
	}
}

func (rl *reentrantLock) runLckDebugFunctions(action, key, id string, lockedTimes int) {
	for _, debugFunc := range rl.lckDebugFunctions {
		debugFunc(action, key, id, lockedTimes)
	}
}
func (rl *reentrantLock) runClrDebugFunctions(keys []string) {
	for _, debugFunc := range rl.clrDebugFunctions {
		debugFunc(keys)
	}
}

// must lock with reentrantLock.locksMapMux
func (rl *reentrantLock) freeNotUsingLocks() {
	if rl.clearPeriodDur <= 0 {
		return
	}
	now := time.Now()
	// not need to clear
	if rl.needClearTime.After(now) {
		return
	}
	// need to clear
	defer func() { rl.needClearTime = now.Add(rl.clearPeriodDur) }()
	toDelKeys := make([]string, 0, len(rl.locksMap)/2)
	for k, v := range rl.locksMap {
		if v.cnt == 0 {
			toDelKeys = append(toDelKeys, k)
		}
	}
	for _, del := range toDelKeys {
		delete(rl.locksMap, del)
	}
	rl.runClrDebugFunctions(toDelKeys)
}

func (rl *reentrantLock) WithOpts(opts ...LockOpt) {
	rl.locksMapMux.Lock()
	defer rl.locksMapMux.Unlock()
	for _, opt := range opts {
		opt(rl)
	}
}

type lockItem struct {
	cond  Cond
	cnt   int
	owner string
	key   string
}

func (rel *lockItem) lock() {
	rel.cond.Lock()
}

func (rel *lockItem) wait() {
	rel.cond.Wait()
}

func (rel *lockItem) unlock() {
	rel.cond.Unlock()
}
