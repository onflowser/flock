//go:build js
// +build js

package flock

func (f *Flock) TryLock() (bool, error) {
	return true, nil
}

func (f *Flock) TryRLock() (bool, error) {
	return true, nil
}

func (f *Flock) Unlock() error {
	return nil
}
