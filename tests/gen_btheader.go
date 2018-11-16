// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package tests

import (
	"encoding/json"
	"math/big"

	"github.com/timenewbank/go-mit/common"
	"github.com/timenewbank/go-mit/common/hexutil"
	"github.com/timenewbank/go-mit/common/math"
	"github.com/timenewbank/go-mit/core/types"
)

var _ = (*btHeaderMarshaling)(nil)

func (b btHeader) MarshalJSON() ([]byte, error) {
	type btHeader struct {
		Bloom            types.Bloom
		Coinbase         common.Address
		MixHash          common.Hash
		Nonce            types.BlockNonce
		Number           *math.HexOrDecimal256
		Hash             common.Hash
		ParentHash       common.Hash
		ReceiptTrie      common.Hash
		StateRoot        common.Hash
		TransactionsTrie common.Hash
		UncleHash        common.Hash
		ExtraData        hexutil.Bytes
		Difficulty       *math.HexOrDecimal256
		GasLimit         math.HexOrDecimal64
		GasUsed          math.HexOrDecimal64
		Timestamp        *math.HexOrDecimal256
	}
	var enc btHeader
	enc.Bloom = b.Bloom
	enc.Coinbase = b.Coinbase
	enc.MixHash = b.MixHash
	enc.Nonce = b.Nonce
	enc.Number = (*math.HexOrDecimal256)(b.Number)
	enc.Hash = b.Hash
	enc.ParentHash = b.ParentHash
	enc.ReceiptTrie = b.ReceiptTrie
	enc.StateRoot = b.StateRoot
	enc.TransactionsTrie = b.TransactionsTrie
	enc.UncleHash = b.UncleHash
	enc.ExtraData = b.ExtraData
	enc.Difficulty = (*math.HexOrDecimal256)(b.Difficulty)
	enc.GasLimit = math.HexOrDecimal64(b.GasLimit)
	enc.GasUsed = math.HexOrDecimal64(b.GasUsed)
	enc.Timestamp = (*math.HexOrDecimal256)(b.Timestamp)
	return json.Marshal(&enc)
}

func (b *btHeader) UnmarshalJSON(input []byte) error {
	type btHeader struct {
		Bloom            *types.Bloom
		Coinbase         *common.Address
		MixHash          *common.Hash
		Nonce            *types.BlockNonce
		Number           *math.HexOrDecimal256
		Hash             *common.Hash
		ParentHash       *common.Hash
		ReceiptTrie      *common.Hash
		StateRoot        *common.Hash
		TransactionsTrie *common.Hash
		UncleHash        *common.Hash
		ExtraData        *hexutil.Bytes
		Difficulty       *math.HexOrDecimal256
		GasLimit         *math.HexOrDecimal64
		GasUsed          *math.HexOrDecimal64
		Timestamp        *math.HexOrDecimal256
	}
	var dec btHeader
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Bloom != nil {
		b.Bloom = *dec.Bloom
	}
	if dec.Coinbase != nil {
		b.Coinbase = *dec.Coinbase
	}
	if dec.MixHash != nil {
		b.MixHash = *dec.MixHash
	}
	if dec.Nonce != nil {
		b.Nonce = *dec.Nonce
	}
	if dec.Number != nil {
		b.Number = (*big.Int)(dec.Number)
	}
	if dec.Hash != nil {
		b.Hash = *dec.Hash
	}
	if dec.ParentHash != nil {
		b.ParentHash = *dec.ParentHash
	}
	if dec.ReceiptTrie != nil {
		b.ReceiptTrie = *dec.ReceiptTrie
	}
	if dec.StateRoot != nil {
		b.StateRoot = *dec.StateRoot
	}
	if dec.TransactionsTrie != nil {
		b.TransactionsTrie = *dec.TransactionsTrie
	}
	if dec.UncleHash != nil {
		b.UncleHash = *dec.UncleHash
	}
	if dec.ExtraData != nil {
		b.ExtraData = *dec.ExtraData
	}
	if dec.Difficulty != nil {
		b.Difficulty = (*big.Int)(dec.Difficulty)
	}
	if dec.GasLimit != nil {
		b.GasLimit = uint64(*dec.GasLimit)
	}
	if dec.GasUsed != nil {
		b.GasUsed = uint64(*dec.GasUsed)
	}
	if dec.Timestamp != nil {
		b.Timestamp = (*big.Int)(dec.Timestamp)
	}
	return nil
}
