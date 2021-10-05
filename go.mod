module github.com/konstellation/cosmodrome

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.42.6
	github.com/goccy/go-yaml v1.9.2
	github.com/imdario/mergo v0.3.11
	github.com/konstellation/konstellation v0.4.3
	github.com/mitchellh/mapstructure v1.3.3
	github.com/pelletier/go-toml v1.8.1
	github.com/rs/zerolog v1.21.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/spm v0.0.0-20210625155357-5a2c8d79013b
	github.com/tendermint/tendermint v0.34.11
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
