// +build unit

package master

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/distributed-fs/pkg/common"
	"github.com/stretchr/testify/assert"
)

const (
	testFile = "test.txt"
)

var master *Master
var fakeAppId string

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
	fakeAppId = uuids[0]
	files := map[string]*File{
		testFile: {
			name:  testFile,
			owner: uuids[1],
			group: map[string]bool{fakeAppId: true},
			permissions: map[common.PermissionGroup]common.PermissionType{
				common.PermissionGroup_Application:      common.PermissionType_Read | common.PermissionType_Write | common.PermissionType_Delete,
				common.PermissionGroup_ApplicationGroup: common.PermissionType_Read | common.PermissionType_Write,
			},
		},
	}

	namespace.files = files
	master.namespace = namespace
}

func TestOperationRequest_SucceedsOnRead(t *testing.T) {
	request := ClientRequest{
		ApplicationId: fakeAppId,
		Operation:     common.OperationType_OPEN,
		Filename:      testFile,
	}

	response, err := master.ClientOperationRequest(context.Background(), &request)

	assert.Nil(t, err)
	assert.True(t, response.Success)
	assert.NotNil(t, response.Token)
	assert.NotEmpty(t, response.Token)
}

func TestOperationRequest_ReturnsErrorWhenInsufficientPermission(t *testing.T) {
	var expectedError = errors.New("application does not have permission to perform requested operation on this file")
	request := ClientRequest{
		ApplicationId: fakeAppId,
		Operation:     common.OperationType_DELETE,
		Filename:      testFile,
	}

	response, actualError := master.ClientOperationRequest(context.Background(), &request)

	assert.Error(t, actualError)
	assert.Equal(t, expectedError, actualError)
	assert.False(t, response.Success)
}

func TestOperationRequest_ReturnsErrorOnAppNotInGroup(t *testing.T) {
	var expectedError = errors.New("application does not have permission to perform requested operation on this file")
	newFakeAppId := uuids[2]
	request := ClientRequest{
		ApplicationId: newFakeAppId,
		Operation:     common.OperationType_READ,
		Filename:      testFile,
	}

	response, actualError := master.ClientOperationRequest(context.Background(), &request)

	assert.Error(t, actualError)
	assert.Equal(t, expectedError, actualError)
	assert.False(t, response.Success)
}
