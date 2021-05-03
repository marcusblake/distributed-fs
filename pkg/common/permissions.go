package common

// PermissionType indicates the actions allowed on a file
type PermissionType uint8

// PermissionGroup indicates the file permission group
type PermissionGroup uint8

// Defines the permissions for a particular group
var GroupPermissions = struct {
	Application      PermissionGroup
	ApplicationGroup PermissionGroup
	All              PermissionGroup
}{
	Application:      1,
	ApplicationGroup: 1 << 1,
	All:              1 << 2,
}

// Defines permissions that applications have for a file
var Permission = struct {
	Read   PermissionType
	Write  PermissionType
	Delete PermissionType
}{
	Read:   1,
	Write:  1 << 1,
	Delete: 1 << 2,
}

func OperationToPermissionType(op FileOperation) PermissionType {
	switch op {
	case Operation.Delete:
		return Permission.Delete
	case Operation.Append:
		return Permission.Write
	default:
		return Permission.Read
	}
}
