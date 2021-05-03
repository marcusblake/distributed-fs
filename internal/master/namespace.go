package master

import "github.com/distributed-fs/pkg/common"

// Namespace is a struct which will contains the file namespace as a tree
type Namespace struct {
	root  *Node
	files map[string]*File // Keep this for now until namespace tree is implemented
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
		map[string]*File{
			"test": {
				name:  "test",
				owner: "1",
				permissions: map[common.PermissionGroup]common.PermissionType{
					common.GroupPermissions.Application: common.Permission.Read | common.Permission.Write,
				},
			},
		},
	}
}

func (nspace *Namespace) GetFileInformation(filename string) (*File, bool) {
	file, ok := nspace.files[filename]
	return file, ok
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
