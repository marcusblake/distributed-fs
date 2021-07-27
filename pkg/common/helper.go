package common

func OperationToPermissionType(op OperationType) PermissionType {
	switch op {
	case OperationType_DELETE:
		return PermissionType_Delete
	case OperationType_APPEND:
		return PermissionType_Write
	default:
		return PermissionType_Read
	}
}
