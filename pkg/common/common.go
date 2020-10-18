package common

// Operation indicates the operation type
type Operation uint8

// PermissionGroup indicates the file permission group
type PermissionGroup uint8

// PermissionType indicates the actions allowed on a file
type PermissionType uint8

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
	User        PermissionGroup = 1
	Group       PermissionGroup = 1 << 1
	All         PermissionGroup = 1 << 2
	Readable    PermissionType  = 1 << 3
	Writeable   PermissionType  = 1 << 4
	Executeable PermissionType  = 1 << 5
)
