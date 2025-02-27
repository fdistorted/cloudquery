package elasticache

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildElasticacheEngineVersions(t *testing.T, ctrl *gomock.Controller) client.Services {
	mockElasticache := mocks.NewMockElastiCache(ctrl)
	output := elasticache.DescribeCacheEngineVersionsOutput{}
	err := faker.FakeData(&output)
	output.Marker = nil
	if err != nil {
		t.Fatal(err)
	}

	mockElasticache.EXPECT().DescribeCacheEngineVersions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&output, nil)

	return client.Services{
		ElastiCache: mockElasticache,
	}
}

func TestElasticacheEngineVersions(t *testing.T) {
	client.AwsMockTestHelper(t, EngineVersions(), buildElasticacheEngineVersions, client.TestOptions{})
}
