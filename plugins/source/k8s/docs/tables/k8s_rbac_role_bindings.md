
# Table: k8s_rbac_role_bindings
RoleBinding references a role, but does not contain it
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|context|text|Name of the context from k8s configuration.|
|kind|text|Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds +optional|
|api_version|text|APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources +optional|
|name|text|Name must be unique within a namespace|
|generate_name|text|GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed|
|namespace|text|Namespace defines the space within which each name must be unique|
|self_link|text|SelfLink is a URL representing this object. Populated by the system. Read-only.  DEPRECATED Kubernetes will stop propagating this field in 1.20 release and the field is planned to be removed in 1.21 release. +optional|
|uid|text|UID is the unique in time and space value for this object|
|resource_version|text|An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed|
|generation|bigint|A sequence number representing a specific generation of the desired state. Populated by the system|
|deletion_grace_period_seconds|bigint|Number of seconds allowed for this object to gracefully terminate before it will be removed from the system|
|labels|jsonb|Map of string keys and values that can be used to organize and categorize (scope and select) objects|
|annotations|jsonb|Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata|
|owner_references|jsonb|List of objects depended by this object|
|finalizers|text[]|Must be empty before the object is deleted from the registry|
|zzz_cluster_name|text|Deprecated: ClusterName is a legacy field that was always cleared by the system and never used; it will be removed in the future. The name in the database is changed to help clients detect accidental use.|
|managed_fields|jsonb|ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow|
|role_ref_api_group|text|APIGroup is the group for the resource being referenced|
|role_ref_kind|text|Kind is the type of resource being referenced|
|role_ref_name|text|Name is the name of resource being referenced|
