// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"

	pb "google.golang.org/genproto/googleapis/cloud/compute/v1"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Instances() *schema.Table {
	return &schema.Table{
		Name:      "gcp_compute_instances",
		Resolver:  fetchInstances,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "advanced_machine_features",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AdvancedMachineFeatures"),
			},
			{
				Name:     "can_ip_forward",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("CanIpForward"),
			},
			{
				Name:     "confidential_instance_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ConfidentialInstanceConfig"),
			},
			{
				Name:     "cpu_platform",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CpuPlatform"),
			},
			{
				Name:     "creation_timestamp",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreationTimestamp"),
			},
			{
				Name:     "deletion_protection",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DeletionProtection"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "disks",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Disks"),
			},
			{
				Name:     "display_device",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DisplayDevice"),
			},
			{
				Name:     "fingerprint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Fingerprint"),
			},
			{
				Name:     "guest_accelerators",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("GuestAccelerators"),
			},
			{
				Name:     "hostname",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Hostname"),
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "key_revocation_action_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeyRevocationActionType"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "label_fingerprint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LabelFingerprint"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "last_start_timestamp",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastStartTimestamp"),
			},
			{
				Name:     "last_stop_timestamp",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastStopTimestamp"),
			},
			{
				Name:     "last_suspended_timestamp",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastSuspendedTimestamp"),
			},
			{
				Name:     "machine_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MachineType"),
			},
			{
				Name:     "metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Metadata"),
			},
			{
				Name:     "min_cpu_platform",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MinCpuPlatform"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "network_interfaces",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NetworkInterfaces"),
			},
			{
				Name:     "network_performance_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NetworkPerformanceConfig"),
			},
			{
				Name:     "params",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Params"),
			},
			{
				Name:     "private_ipv_6_google_access",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PrivateIpv6GoogleAccess"),
			},
			{
				Name:     "reservation_affinity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ReservationAffinity"),
			},
			{
				Name:     "resource_policies",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ResourcePolicies"),
			},
			{
				Name:     "satisfies_pzs",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SatisfiesPzs"),
			},
			{
				Name:     "scheduling",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Scheduling"),
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "service_accounts",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ServiceAccounts"),
			},
			{
				Name:     "shielded_instance_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ShieldedInstanceConfig"),
			},
			{
				Name:     "shielded_instance_integrity_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ShieldedInstanceIntegrityPolicy"),
			},
			{
				Name:     "source_machine_image",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceMachineImage"),
			},
			{
				Name:     "source_machine_image_encryption_key",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SourceMachineImageEncryptionKey"),
			},
			{
				Name:     "start_restricted",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("StartRestricted"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "status_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StatusMessage"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Zone"),
			},
		},
	}
}

func fetchInstances(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.AggregatedListInstancesRequest{
		Project: c.ProjectId,
	}
	it := c.Services.ComputeInstancesClient.AggregatedList(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return errors.WithStack(err)
		}

		res <- resp.Value.Instances

	}
	return nil
}
