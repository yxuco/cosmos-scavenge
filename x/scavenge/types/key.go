package types

const (
	// ModuleName is the name of the module
	ModuleName = "scavenge"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querierer msgs
	QuerierRoute = ModuleName

	// ScavengePrefix is key prefix for scavenge store
	ScavengePrefix = "sk-"

	// CommitPrefix is key prefix for commit store
	CommitPrefix = "ck-"
)
