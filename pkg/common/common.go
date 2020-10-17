package common

// Operation indicates the operation type
type Operation uint8

// Permission indicates the file permission
type Permission uint8

// Defines the type of operations that are permitted
const (
	Open Operation = iota
	Close
	Read
	Append
	Delete
	Snapshot
)

// Defines permissions that applications have for a file
const (
	Reader Permission = 1
	Writer Permission = 1 << 1
)
