//go:build linux

package main

import (
	"io"
	"os"
	"syscall"
)

// readFile opens the file at path, acquires a shared (read) advisory lock via
// flock(2), reads its entire content and then releases the lock. Holding a
// shared lock ensures the hyper-v kvp daemon is not writing to the pool file
// while it is being read. flock(2) is only available on Linux.
func readFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()

	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_SH); err != nil {
		return nil, err
	}
	defer func() {
		_ = syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	}()

	return io.ReadAll(f)
}
