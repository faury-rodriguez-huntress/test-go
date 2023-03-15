package error

import (
	"syscall"

	"golang.org/x/sys/windows"
)

func (e OsError) isErrPermission(errno syscall.Errno) bool {
	return errno == windows.ERROR_SHARING_VIOLATION ||
		errno == windows.ERROR_LOCK_VIOLATION
}

func (e OsError) isErrNotExist(errno syscall.Errno) bool {
	switch errno {
	case windows.ERROR_INVALID_PARAMETER,
		windows.ERROR_INVALID_NAME:
		return true

	default:
		return false
	}
}

func (e OsError) isErrInvalid(errno syscall.Errno) bool {
	switch errno {
	case windows.ERROR_INVALID_PARAMETER,
		windows.ERROR_INVALID_NAME:
		return true

	default:
		return false
	}
}
