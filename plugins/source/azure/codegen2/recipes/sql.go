// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"

func init() {
	tables := []Table{
		{
			Service:        "armsql",
			Name:           "instance_pools",
			Struct:         &armsql.InstancePool{},
			ResponseStruct: &armsql.InstancePoolsClientListResponse{},
			Client:         &armsql.InstancePoolsClient{},
			ListFunc:       (&armsql.InstancePoolsClient{}).NewListPager,
			NewFunc:        armsql.NewInstancePoolsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/instancePools",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_sql)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armsql",
			Name:           "managed_instances",
			Struct:         &armsql.ManagedInstance{},
			ResponseStruct: &armsql.ManagedInstancesClientListResponse{},
			Client:         &armsql.ManagedInstancesClient{},
			ListFunc:       (&armsql.ManagedInstancesClient{}).NewListPager,
			NewFunc:        armsql.NewManagedInstancesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/managedInstances",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_sql)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armsql",
			Name:           "virtual_clusters",
			Struct:         &armsql.VirtualCluster{},
			ResponseStruct: &armsql.VirtualClustersClientListResponse{},
			Client:         &armsql.VirtualClustersClient{},
			ListFunc:       (&armsql.VirtualClustersClient{}).NewListPager,
			NewFunc:        armsql.NewVirtualClustersClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/virtualClusters",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_sql)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}
