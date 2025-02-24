package kvstore

import (
	"errors"
	"fmt"
)

var (
	// ErrKeyNotFound is returned when an op. doesn't find the given key.
	ErrKeyNotFound = errors.New("key not found")

	EmptyPrefix = KeyPrefix{}
)

type Realm = []byte
type KeyPrefix = []byte
type Key = []byte
type Value = []byte

// IterDirection specifies the direction for iterations.
type IterDirection byte

const (
	IterDirectionForward IterDirection = iota
	IterDirectionBackward
)

// IteratorKeyValueConsumerFunc is a consumer function for an iterating function which iterates over keys and values.
// They key must not be prefixed with the realm.
// Returning false from this function indicates to abort the iteration.
type IteratorKeyValueConsumerFunc func(key Key, value Value) bool

// IteratorKeyConsumerFunc is a consumer function for an iterating function which iterates only over keys.
// They key must not be prefixed with the realm.
// Returning false from this function indicates to abort the iteration.
type IteratorKeyConsumerFunc func(key Key) bool

// BatchedMutations represents batched mutations to the storage.
type BatchedMutations interface {

	// Set sets the given key and value.
	Set(key Key, value Value) error

	// Delete deletes the entry for the given key.
	Delete(key Key) error

	// Cancel cancels the batched mutations.
	Cancel()

	// Commit commits/flushes the mutations.
	Commit() error
}

// KVStore persists, deletes and retrieves data.
type KVStore interface {
	// WithRealm is a factory method for using the same underlying storage with a different realm.
	WithRealm(realm Realm) KVStore

	// Realm returns the configured realm.
	Realm() Realm

	// Shutdown marks the store as shutdown.
	Shutdown()

	// Iterate iterates over all keys and values with the provided prefix. You can pass kvstore.EmptyPrefix to iterate over all keys and values.
	// Optionally the direction for the iteration can be passed (default: IterDirectionForward).
	Iterate(prefix KeyPrefix, kvConsumerFunc IteratorKeyValueConsumerFunc, direction ...IterDirection) error

	// IterateKeys iterates over all keys with the provided prefix. You can pass kvstore.EmptyPrefix to iterate over all keys.
	// Optionally the direction for the iteration can be passed (default: IterDirectionForward).
	IterateKeys(prefix KeyPrefix, consumerFunc IteratorKeyConsumerFunc, direction ...IterDirection) error

	// Clear clears the realm.
	Clear() error

	// Get gets the given key or nil if it doesn't exist or an error if an error occurred.
	Get(key Key) (value Value, err error)

	// Set sets the given key and value.
	Set(key Key, value Value) error

	// Has checks whether the given key exists.
	Has(key Key) (bool, error)

	// Delete deletes the entry for the given key.
	Delete(key Key) error

	// DeletePrefix deletes all the entries matching the given key prefix.
	DeletePrefix(prefix KeyPrefix) error

	// Batched returns a BatchedMutations interface to execute batched mutations.
	Batched() BatchedMutations

	// Flush persists all outstanding write operations to disc.
	Flush() error

	// Close closes the database file handles.
	Close() error
}

// GetIterDirection returns the direction to use for an iteration.
// If no direction is given, it defaults to IterDirectionForward.
func GetIterDirection(iterDirection ...IterDirection) IterDirection {
	direction := IterDirectionForward
	if len(iterDirection) > 0 {
		switch iterDirection[0] {
		case IterDirectionForward:
			break
		case IterDirectionBackward:
			direction = iterDirection[0]
		default:
			panic(fmt.Sprintf("unknown iteration direction: %d", iterDirection[0]))
		}
	}
	return direction
}
