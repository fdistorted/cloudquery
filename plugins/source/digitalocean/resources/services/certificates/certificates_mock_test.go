package certificates

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createCertificates(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCertificatesService(ctrl)

	var data []godo.Certificate
	if err := faker.FakeData(&data); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().List(gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)

	//add children mocks

	return client.Services{
		Certificates: m,
	}
}

func TestCertificates(t *testing.T) {
	client.MockTestHelper(t, Certificates(), createCertificates, client.TestOptions{})
}
