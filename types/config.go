package types

import (
	"time"
)

//-----------------------------------------------------------------------------
// BaseConfig

// BaseConfig defines the base configuration for a Tendermint node
type BaseConfig struct {
	// A custom human readable name for this node
	//Moniker string `json:"moniker" mapstructure:"moniker"`
}

//-----------------------------------------------------------------------------
// RPCConfig

// RPCConfig defines the configuration options for the Tendermint RPC server
type RPCConfig struct {
	// TCP or UNIX socket address for the RPC server to listen on
	ListenAddress string `json:"laddr" mapstructure:"laddr"`

	// A list of origins a cross-domain request can be executed from.
	// If the special '*' value is present in the list, all origins will be allowed.
	// An origin may contain a wildcard (*) to replace 0 or more characters (i.e.: http://*.domain.com).
	// Only one wildcard can be used per origin.
	CORSAllowedOrigins []string `json:"cors_allowed_origins" mapstructure:"cors_allowed_origins"`
}

//-----------------------------------------------------------------------------
// ConsensusConfig

// ConsensusConfig defines the configuration for the Tendermint consensus service,
// including timeouts and details about the WAL and the block structure.
type ConsensusConfig struct {
	TimeoutCommit time.Duration `json:"timeout_commit" mapstructure:"timeout_commit"`
	// EmptyBlocks mode and possible interval between empty blocks
	CreateEmptyBlocks         bool          `json:"create_empty_blocks" mapstructure:"create_empty_blocks"`
	CreateEmptyBlocksInterval time.Duration `json:"create_empty_blocks_interval" mapstructure:"create_empty_blocks_interval"`
}

type Config struct {
	// Top level options use an anonymous struct
	BaseConfig `mapstructure:",squash"`

	// Options for services
	RPC       *RPCConfig       `mapstructure:"rpc" json:"rpc"`
	Consensus *ConsensusConfig `mapstructure:"consensus" json:"consensus"`
}
