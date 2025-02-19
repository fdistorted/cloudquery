package fsx

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildStorageVmsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockFsxClient(ctrl)

	var vm types.StorageVirtualMachine
	require.NoError(t, faker.FakeData(&vm))
	m.EXPECT().DescribeStorageVirtualMachines(
		gomock.Any(),
		&fsx.DescribeStorageVirtualMachinesInput{MaxResults: aws.Int32(1000)},
	).Return(
		&fsx.DescribeStorageVirtualMachinesOutput{StorageVirtualMachines: []types.StorageVirtualMachine{vm}},
		nil,
	)
	return client.Services{
		FSX: m,
	}
}

func TestStorageVms(t *testing.T) {
	client.AwsMockTestHelper(t, StorageVms(), buildStorageVmsMock, client.TestOptions{})
}
