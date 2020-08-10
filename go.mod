module github.com/cosmos/cosmos-sdk

require (
	github.com/99designs/keyring v1.1.4
	github.com/alecthomas/gocyclo v0.0.0-20150208221726-aa8f8b160214 // indirect
	github.com/alexkohler/nakedret v1.0.0 // indirect
	github.com/bgentry/speakeasy v0.1.0
	github.com/btcsuite/btcd v0.0.0-20190115013929-ed77733ec07d
	github.com/cosmos/go-bip39 v0.0.0-20180819234021-555e2067c45d
	github.com/cosmos/ledger-cosmos-go v0.11.1
	github.com/gibson042/canonicaljson-go v1.0.3
	github.com/gogo/protobuf v1.3.1
	github.com/golang/mock v1.4.3
	github.com/golang/protobuf v1.4.0-rc.4
	github.com/gordonklaus/ineffassign v0.0.0-20200809085317-e36bfde3bb78 // indirect
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/hashicorp/golang-lru v0.5.4
	github.com/mattn/go-isatty v0.0.12
	github.com/mdempsky/maligned v0.0.0-20180708014732-6e39bd26a8c8 // indirect
	github.com/mdempsky/unconvert v0.0.0-20200228143138-95ecdbfc0b5f // indirect
	github.com/mibk/dupl v1.0.0 // indirect
	github.com/opennota/check v0.0.0-20180911053232-0c771f5545ff // indirect
	github.com/otiai10/copy v1.1.1
	github.com/pelletier/go-toml v1.6.0
	github.com/pkg/errors v0.9.1
	github.com/rakyll/statik v0.1.7
	github.com/regen-network/cosmos-proto v0.1.1-0.20200213154359-02baa11ea7c2
	github.com/rs/cors v1.7.0
	github.com/spf13/afero v1.2.1 // indirect
	github.com/spf13/cobra v0.0.7
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.5.1
	github.com/tendermint/btcd v0.1.1
	github.com/tendermint/crypto v0.0.0-20191022145703-50d29ede1e15
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/iavl v0.13.2
	github.com/tendermint/tendermint v0.33.2
	github.com/tendermint/tm-db v0.5.1
	github.com/tsenart/deadcode v0.0.0-20160724212837-210d2dc333e9 // indirect
	github.com/walle/lll v1.0.1 // indirect
	google.golang.org/protobuf v1.20.1 // indirect
	gopkg.in/yaml.v2 v2.2.8
	mvdan.cc/unparam v0.0.0-20200501210554-b37ab49443f7 // indirect
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.1

go 1.14
