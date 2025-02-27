//go:build integration
// +build integration

package networking

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
)

func TestIntegrationNetworkPolicies(t *testing.T) {
	client.K8sTestHelper(t, NetworkPolicies(), "./snapshots")
}
