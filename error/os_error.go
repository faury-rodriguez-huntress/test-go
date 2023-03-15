package error

import (
	"io/fs"
	"os"
	"syscall"
)

type OsError struct {
	WrappedError error
}

func (e OsError) Error() string { return e.WrappedError.Error() }

func (e OsError) Unwrap() error {
	return e.WrappedError
}

func (e OsError) Is(target error) bool {
	errno, ok := e.getErrno()
	if !ok {
		return false
	}

	switch target {
	case fs.ErrPermission:
		return e.isErrPermission(errno)

	case fs.ErrNotExist:
		return e.isErrNotExist(errno)

	case fs.ErrInvalid:
		return e.isErrInvalid(errno)
	}

	return false
}

func (e OsError) getErrno() (errno syscall.Errno, ok bool) {
	switch x := e.WrappedError.(type) {
	case *os.PathError:
		errno, ok := x.Err.(syscall.Errno)
		return errno, ok
	}

	return errno, false
}
