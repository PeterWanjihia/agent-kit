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