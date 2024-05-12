package locku

var (
	defaultGlobalReentrantLock = NewReentrantLock(
		WithLockClearSec(3600 * 24), // 1天清理一次
	)
)

func LockKeysInMemory(entry LockEntry) {
	defaultGlobalReentrantLock.LockKeysInMemory(entry)
}

func UnlockKeysInMemory(entry LockEntry) {
	defaultGlobalReentrantLock.UnlockKeysInMemory(entry)
}

func DefaultGlobalMemoryLockOpt(opts ...LockOpt) {
	defaultGlobalReentrantLock.WithOpts(opts...)
}

//
//func LockKeysInMemory(entry LockEntry) {
//
//	xLockInMemoryWithAction(entry, func(lock *lockItem) bool {
//		lock.lock()
//		defer lock.unlock()
//		if lock.owner == entry.RequestID || lock.cnt == 0 {
//			lock.cnt++
//			lock.owner = entry.RequestID
//			return true
//		}
//		lock.wait()
//		return false
//	})
//}
//
//func UnLLockKeysInMemory(entry LockEntry) {
//
//	xLockInMemoryWithAction(entry, func(lock *lockItem) bool {
//		lock.lock()
//		defer lock.unlock()
//		if lock.owner != entry.RequestID {
//			panic(fmt.Sprintf("try[id:%s] unlock a not holding lock[%s]", entry.RequestID, lock.owner))
//		}
//		// id equal
//		if lock.cnt <= 0 {
//			panic(fmt.Sprintf("try[id:%s] unlock a not exist lock", entry.RequestID))
//		}
//		lock.cnt--
//		return true
//	})
//}
//
//func xLockInMemoryWithAction(entry LockEntry, checkLockAndDo func(*lockItem) bool) {
//	keys := entry.Keys
//	if !slices.IsSortedFunc(keys, func(a, b string) int { return strings.Compare(a, b) }) {
//		sort.Strings(keys)
//	}
//	if len(keys) == 0 {
//		return
//	}
//
//	toLocks := make([]*lockItem, 0, len(keys))
//	memLockMux.Lock()
//	for _, key := range keys {
//		rel, exist := memLock[key]
//		if !exist {
//			rel = &lockItem{
//				cond:  NewCond(),
//				cnt:   0,
//				owner: "",
//				key:   key,
//			}
//			memLock[key] = rel
//		}
//		toLocks = append(toLocks, rel)
//	}
//	memLockMux.Unlock()
//	for _, lock := range toLocks {
//		// lock
//		for !checkLockAndDo(lock) {
//		}
//	}
//}
