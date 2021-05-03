package master

import "github.com/distributed-fs/pkg/common"

type File struct {
	name        string
	permissions map[common.PermissionGroup]common.PermissionType
	owner       string
	group       map[string]bool
	location    string
}
