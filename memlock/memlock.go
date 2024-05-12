package memlock

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
