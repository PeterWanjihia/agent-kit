package chain

import (
	"context"

	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)
func MakeAuth(privateKeyHex string, chainID int64) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(
		privateKey,
		big.NewInt(chainID),
	)
	if err != nil {
		return nil, err
	}

	auth.Context = context.Background()
	return auth, nil
}

func BindCallerOpts(ctx context.Context,address *common.Address)(*bind.CallOpts,error){

	if ctx ==nil{
		ctx = context.Background()
	}
	opts := &bind.CallOpts{
		Context: ctx,
	}
	if address !=nil{
		opts.From = *address
	}
	return opts,nil

}