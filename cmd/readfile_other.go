//go:build !linux

package main

import "os"

// readFile reads the entire content of the file at path. Advisory read locking
// via flock(2) is only supported on Linux, so on other platforms the file is
// read without holding a lock.
func readFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
