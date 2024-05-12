package locku

import (
	"context"
	"fmt"
	"github.com/stdxgo/common-utils/crash"
	"github.com/stdxgo/common-utils/strutil"
	"strings"
	"sync"
	"testing"
	"time"
)

func genLocks(lockNum int, keysNum int) []LockEntry {
	result := make([]LockEntry, 0, lockNum)
	for i := 0; i < lockNum; i++ {
		result = append(result, genLock(keysNum))
	}
	return result
}
func genLock(keysLen int) LockEntry {

	result := make([]string, 0, keysLen)
	for i := 0; i < keysLen; i++ {
		result = append(result, "key_"+strutil.Gen_Test_Num_Only_For_Test(1))
	}
	id := strutil.GenNum(idLen)
	return LockEntry{
		RequestID: "id_" + id,
		Type:      "",
		Keys:      result,
	}
}

var (
	_testKeysFormat       = func(keys []string) string { return strings.Join(keys, "_") }
	_testTimeFormatLayout = "2006-01-02 15:04:05.000"
	_testLockDebugFunc    = func(key, owner string, lockedTimes int) {
		fmt.Printf("%s\t%s\tcnt_%d\n", key, owner, lockedTimes)
	}
	_testLockClrDebugFunc = func(keys []string) {
		fmt.Printf("clear keys[%s]:%s\n", time.Now().Format(_testTimeFormatLayout), strings.Join(keys, ","))
	}
)

const (
	idLen = 2
)

func TestReentrant(t *testing.T) {

	defer crash.HandlerStk(context.TODO(), func(stk []byte) {
		fmt.Println(string(stk))
	})

	const lockN = 500
	const keysN = 1

	DefaultGlobalMemoryLockOpt(
		//WithLockDebugFunc(_testLockDebugFunc),
		WithLockClearSec(1),
		WithClearDebugFunc(_testLockClrDebugFunc),
	)

	les := genLocks(lockN, keysN)
	//bs, _ := json.Marshal(les)
	//fmt.Println(string(bs))
	now := time.Now()
	defer func() {
		fmt.Println(now.Format(_testTimeFormatLayout))
		fmt.Println(time.Now().Format(_testTimeFormatLayout))
	}()

	var wg sync.WaitGroup
	wg.Add(lockN)
	for i := 0; i < lockN; i++ {
		go func(x int) {
			le := les[x]
			defer wg.Done()
			LockKeysInMemory(le)
			//xInfo := fmt.Sprintf("%s\t%s", _testKeysFormat(le.Keys), le.RequestID)
			defer func() {
				//fmt.Println("unlock:", xInfo)
				UnlockKeysInMemory(le)
			}()

			//fmt.Println("  lock:", xInfo)
			time.Sleep(time.Second / 10)
		}(i)
	}
	wg.Wait()
}

func TestOneReentrant(t *testing.T) {

	le := LockEntry{
		RequestID: "123",
		Type:      "",
		Keys:      []string{"a", "b"},
	}

	LockKeysInMemory(le)

	le.Keys = []string{"b", "a"}
	LockKeysInMemory(le)

	le.RequestID = "2"
	//LockKeysInMemory(le)
}
