package common

// Operation indicates the operation type
type FileOperation uint8

// Defines the type of operations that are permitted
var Operation = struct {
	Open     FileOperation
	Close    FileOperation
	Read     FileOperation
	Append   FileOperation
	Delete   FileOperation
	Snapshot FileOperation
}{
	Open:     1,
	Close:    2,
	Read:     3,
	Append:   4,
	Delete:   5,
	Snapshot: 6,
}
