// Code generated by codegen; DO NOT EDIT.

package fsx

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func FileSystems() *schema.Table {
	return &schema.Table{
		Name:      "aws_fsx_file_systems",
		Resolver:  fetchFsxFileSystems,
		Multiplex: client.ServiceAccountRegionMultiplexer("fsx"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:          "administrative_actions",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("AdministrativeActions"),
				IgnoreInTests: true,
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "dns_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DNSName"),
			},
			{
				Name:     "failure_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FailureDetails"),
			},
			{
				Name:     "file_system_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FileSystemId"),
			},
			{
				Name:     "file_system_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FileSystemType"),
			},
			{
				Name:     "file_system_type_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FileSystemTypeVersion"),
			},
			{
				Name:     "kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KmsKeyId"),
			},
			{
				Name:     "lifecycle",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Lifecycle"),
			},
			{
				Name:     "lustre_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LustreConfiguration"),
			},
			{
				Name:     "network_interface_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("NetworkInterfaceIds"),
			},
			{
				Name:     "ontap_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("OntapConfiguration"),
			},
			{
				Name:     "open_zfs_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("OpenZFSConfiguration"),
			},
			{
				Name:     "owner_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OwnerId"),
			},
			{
				Name:     "storage_capacity",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("StorageCapacity"),
			},
			{
				Name:     "storage_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StorageType"),
			},
			{
				Name:     "subnet_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SubnetIds"),
			},
			{
				Name:     "vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcId"),
			},
			{
				Name:     "windows_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("WindowsConfiguration"),
			},
		},
	}
}
