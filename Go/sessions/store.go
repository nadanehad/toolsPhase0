package sessions

import "sync"

var SessionStore = struct {
	sync.RWMutex
	Sessions map[string]uint
}{Sessions: make(map[string]uint)}
