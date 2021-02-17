module github.com/konstellation/cosmodrome

go 1.15

require (
	github.com/bramvdbogaerde/go-scp v0.0.0-20200119201711-987556b8bdd7
	github.com/cosmos/cosmos-sdk v0.41.0
	github.com/ethereum/go-ethereum v1.9.11
	github.com/google/uuid v1.1.2
	github.com/konstellation/kn-sdk v0.2.0
	github.com/konstellation/konstellation v0.2.0
	github.com/mitchellh/mapstructure v1.1.2
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/tendermint/iavl v0.12.4
	github.com/tendermint/tendermint v0.34.3
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
