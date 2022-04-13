module github.com/celestiaorg/demochain

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.44.3
	github.com/cosmos/ibc-go v1.2.0
	github.com/gogo/protobuf v1.3.3
	github.com/golang/glog v1.0.0 // indirect
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.4.0
	github.com/stretchr/testify v1.7.1
	github.com/tendermint/spm v0.1.6
	github.com/tendermint/tendermint v0.34.14
	github.com/tendermint/tm-db v0.6.6
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	google.golang.org/genproto v0.0.0-20220317150908-0efb43f6373e
	google.golang.org/grpc v1.45.0
)

replace (
	github.com/99designs/keyring => github.com/cosmos/keyring v1.1.7-0.20210622111912-ef00f8ac3d76
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
)

replace github.com/cosmos/cosmos-sdk => github.com/celestiaorg/cosmos-sdk v0.44.3-0.20220413122210-0ecef447970b
