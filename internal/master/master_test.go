// +build unit

package master

import (
	"errors"
	"os"
	"testing"

	"github.com/distributed-fs/internal/rpctype"
	"github.com/distributed-fs/pkg/common"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	testFile = "test.txt"
)

var master *Master
var fakeAppId uuid.UUID

var uuids = []string{
	"f47ac10b-58cc-0372-8567-0e02b2c3d479",
	"ga4ac10b-58cc-2372-8567-0e02b2c3d479",
	"dd7ac10b-58cc-2372-8567-0e02b2c3d479",
}

func TestMain(m *testing.M) {
	initialize()
	code := m.Run()
	os.Exit(code)
}

func initialize() {
	master = NewMaster()

	namespace := NewNamespace()
	fakeAppId, _ = uuid.Parse(uuids[0])
	files := map[string]*File{
		testFile: {
			name:  testFile,
			owner: uuids[1],
			group: map[string]bool{fakeAppId.String(): true},
			permissions: map[common.PermissionGroup]common.PermissionType{
				common.GroupPermissions.Application:      common.Permission.Read | common.Permission.Write | common.Permission.Delete,
				common.GroupPermissions.ApplicationGroup: common.Permission.Read | common.Permission.Write,
			},
		},
	}

	namespace.files = files
	master.namespace = namespace
}

func TestOperationRequest_SucceedsOnRead(t *testing.T) {
	request := rpctype.OperationRequest{
		ApplicationId: fakeAppId,
		Operation:     common.Operation.Open,
		Filename:      testFile,
	}

	var response rpctype.OperationResponse
	err := master.OperationRequest(&request, &response)

	assert.Nil(t, err)
	assert.True(t, response.Ok)
	assert.NotNil(t, response.Token)
	assert.NotEmpty(t, response.Token)
}

func TestOperationRequest_ReturnsErrorWhenInsufficientPermission(t *testing.T) {
	var expectedError = errors.New("application does not have permission to perform requested operation on this file")
	request := rpctype.OperationRequest{
		ApplicationId: fakeAppId,
		Operation:     common.Operation.Delete,
		Filename:      testFile,
	}

	var response rpctype.OperationResponse
	actualError := master.OperationRequest(&request, &response)

	assert.Error(t, actualError)
	assert.Equal(t, expectedError, actualError)
	assert.False(t, response.Ok)
}

func TestOperationRequest_ReturnsErrorOnAppNotInGroup(t *testing.T) {
	var expectedError = errors.New("application does not have permission to perform requested operation on this file")
	newFakeAppId, _ := uuid.Parse(uuids[2])
	request := rpctype.OperationRequest{
		ApplicationId: newFakeAppId,
		Operation:     common.Operation.Read,
		Filename:      testFile,
	}

	var response rpctype.OperationResponse
	actualError := master.OperationRequest(&request, &response)

	assert.Error(t, actualError)
	assert.Equal(t, expectedError, actualError)
	assert.False(t, response.Ok)
}
