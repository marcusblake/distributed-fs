package master

// Namespace is a struct which will contains the file namespace as a tree
type Namespace struct {
	root *Node
}

// Node is a node in the namespace tree
type Node struct {
	Name     string
	Children []*Node
}

// NewNamespace creates a new namespace struct
func NewNamespace() *Namespace {
	return &Namespace{
		&Node{
			Name:     "/",
			Children: []*Node{},
		},
	}
}

// Open opens a file in the namespace
func Open(filepath string, namespace *Namespace) error {
	return nil
}

// Close closes a file in the namespace
func Close(filepath string, namespace *Namespace) error {
	return nil
}

// Read file indicates that a file is going to be read
func Read(filepath string, namespace *Namespace) error {
	return nil
}

// Append indicates that data will be appended to the file
func Append(filepath string, namespace *Namespace) error {
	return nil
}

// Delete indicates that data will be deleted from the file
func Delete(filepath string, namespace *Namespace) error {
	return nil
}

// Snapshot indicates that a snapshot will be taken
func Snapshot(filepath string, namespace *Namespace) error {
	return nil
}
