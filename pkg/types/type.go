package types


type ContractsRegistry struct{
	Identity string
	Reputation string

}
type Erc8004ContractsRegistry struct{
	ChainName string
	ContractAddress ContractsRegistry

}


type NetworkConfig struct{
	ChainName string
	RPC_URL string
	
}

type RegisterResult struct {
	TxHash   string `json:"txHash"`
	URI      string `json:"uri"`
	Chain    string `json:"chain"`
	Network  string `json:"network"`
	ChainID  int64  `json:"chainId"`
	Registry string `json:"registry"`
}

type RegisterOptions struct {
	AgentURI   string
	ChainName  string
	Network    string
	PrivateKey string
	JsonOutput bool
}