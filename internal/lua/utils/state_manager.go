package utils

import (
	"sync"
	"time"

	lua "github.com/yuin/gopher-lua"
)

// StateManager provides a thread-safe mechanism to manage persistent and temporary states.
type StateManager struct {
	mu    sync.Mutex
	store map[string]*stateItem
}

// stateItem represents an individual state with optional expiry.
type stateItem struct {
	Value     lua.LValue
	ExpiresAt *time.Time // nil means no expiry
}

// NewStateManager initializes a new StateManager.
func NewStateManager() *StateManager {
	sm := &StateManager{
		store: make(map[string]*stateItem),
	}

	go func(sm *StateManager) {
		for {
			time.Sleep(1 * time.Minute)
			sm.mu.Lock()
			for key, item := range sm.store {
				if item.ExpiresAt != nil && time.Now().After(*item.ExpiresAt) {
					delete(sm.store, key)
				}
			}
			sm.mu.Unlock()
		}
	}(sm)

	return sm
}

// Set stores a value with an optional expiry time.
func (sm *StateManager) Set(key string, value lua.LValue, expirySeconds int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	var expiresAt *time.Time
	if expirySeconds > 0 {
		exp := time.Now().Add(time.Duration(expirySeconds) * time.Second)
		expiresAt = &exp
	}

	sm.store[key] = &stateItem{
		Value:     value,
		ExpiresAt: expiresAt,
	}
}

// Get retrieves a value by key. Returns nil if expired or not found.
func (sm *StateManager) Get(key string) lua.LValue {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	item, exists := sm.store[key]
	if !exists {
		return lua.LNil
	}

	if item.ExpiresAt != nil && time.Now().After(*item.ExpiresAt) {
		delete(sm.store, key) // Remove expired item
		return lua.LNil
	}

	return item.Value
}

// Clear removes a specific key and its value.
func (sm *StateManager) Clear(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	delete(sm.store, key)
}
