package chain

import "github.com/ronexlemon/agent-kit/pkg/types"

var Mainnet_RPCS = map[string]types.NetworkConfig{
	"Ethereum": {
		ChainName: "ethereum",
		RPC_URL:  "https://eth.llamarpc.com",
	},
	"Base": {
		ChainName: "base",
		RPC_URL:  "https://mainnet.base.org",
	},
	"Arbitrum": {
		ChainName: "arbitrum",
		RPC_URL:  "https://arb1.arbitrum.io/rpc",
	},
	"Avalanche": {
		ChainName: "avalanche",
		RPC_URL:  "https://api.avax.network/ext/bc/C/rpc",
	},
	"Celo": {
		ChainName: "celo",
		RPC_URL:  "https://forno.celo.org",
	},
	"Monad": {
		ChainName: "monad",
		RPC_URL:  "https://rpc.monad.xyz",
	},
}

var Testnet_RPCS = map[string]types.NetworkConfig{
	"Ethereum": {
		ChainName: "ethereum",
		RPC_URL:  "https://ethereum-sepolia-rpc.publicnode.com",
	},
	"Base": {
		ChainName: "base",
		RPC_URL:  "https://sepolia.base.org",
	},
	"Arbitrum": {
		ChainName: "arbitrum",
		RPC_URL:  "https://sepolia-rollup.arbitrum.io/rpc",
	},
	"Avalanche": {
		ChainName: "avalanche",
		RPC_URL:  "https://api.avax-test.network/ext/bc/C/rpc",
	},
	"Celo": {
		ChainName: "celo",
		RPC_URL:  "https://celo-sepolia.drpc.org",
	},
	"Monad": {
		ChainName: "monad",
		RPC_URL:  "https://testnet-rpc.monad.xyz",
	},
}