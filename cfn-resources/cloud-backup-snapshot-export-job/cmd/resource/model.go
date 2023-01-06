// Code generated by 'cfn generate', changes will be undone by the next invocation. DO NOT EDIT.
// Updates to this type are made my editing the schema file and executing the 'generate' command.
package resource

// Model is autogenerated from the json schema
type Model struct {
	ApiKeys        *ApiKeyDefinition                         `json:",omitempty"`
	ClusterName    *string                                   `json:",omitempty"`
	Components     []ApiAtlasDiskBackupBaseRestoreMemberView `json:",omitempty"`
	CreatedAt      *string                                   `json:",omitempty"`
	CustomDataSet  []CustomData                              `json:",omitempty"`
	ExportBucketId *string                                   `json:",omitempty"`
	ExportId       *string                                   `json:",omitempty"`
	ExportStatus   *ExportStatus                             `json:",omitempty"`
	FinishedAt     *string                                   `json:",omitempty"`
	TestMode       *string                                   `json:",omitempty"`
	GroupId        *string                                   `json:",omitempty"`
	Links          []Link                                    `json:",omitempty"`
	Prefix         *string                                   `json:",omitempty"`
	SnapshotId     *string                                   `json:",omitempty"`
	State          *string                                   `json:",omitempty"`
}

// ApiKeyDefinition is autogenerated from the json schema
type ApiKeyDefinition struct {
	PrivateKey *string `json:",omitempty"`
	PublicKey  *string `json:",omitempty"`
}

// ApiAtlasDiskBackupBaseRestoreMemberView is autogenerated from the json schema
type ApiAtlasDiskBackupBaseRestoreMemberView struct {
	ExportID       *string `json:",omitempty"`
	ReplicaSetName *string `json:",omitempty"`
}

// CustomData is autogenerated from the json schema
type CustomData struct {
	Key   *string `json:",omitempty"`
	Value *string `json:",omitempty"`
}

// ExportStatus is autogenerated from the json schema
type ExportStatus struct {
	ExportedCollections *int `json:",omitempty"`
	TotalCollections    *int `json:",omitempty"`
}

// Link is autogenerated from the json schema
type Link struct {
	Href *string `json:",omitempty"`
	Rel  *string `json:",omitempty"`
}
