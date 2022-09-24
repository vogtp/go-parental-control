//go:generate stringer -type=LockState

package sessions

type LockState uint

const (
	Locked   LockState = 0
	Unlocked LockState = 1
	Unknown  LockState = 4294967295
)
