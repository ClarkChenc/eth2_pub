package eth

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

const ethJSONRPCEndpointAddress = "https://mainnet.infura.io/v3/e094bd4872984c20a288050c1d291598"

func EthGasFee() {
	config := params.MainnetChainConfig
	ethClient, err := ethclient.DialContext(context.Background(), ethJSONRPCEndpointAddress)
	if err != nil {
		log.Panic(err)
	}

	bn, _ := ethClient.BlockNumber(context.Background())

	bignumBn := big.NewInt(0).SetUint64(bn)
	blk, _ := ethClient.BlockByNumber(context.Background(), bignumBn)
	baseFee := misc.CalcBaseFee(config, blk.Header())
	fmt.Printf("Base fee for block %d is %s\n", bn+1, baseFee.String())
}
