package types


type ContractsRegistry struct{
	Identity string
	Reputation string

}
type Erc8004ContractsRegistry struct{
	ChainName string
	ContractAddress ContractsRegistry
	Rpc_Url NetworkConfig

}


type NetworkConfig struct{
	RPC_URL string
	
}