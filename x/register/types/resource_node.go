package types

import (
	"bytes"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stratosnet/stratos-chain/x/register/exported"
	"github.com/tendermint/tendermint/crypto"
	"sort"
	"strings"
)

type ResourceNode struct {
	NetworkAddress string         `json:"network_address" yaml:"network_address"` // network address of the resource node
	PubKey         crypto.PubKey  `json:"pubkey" yaml:"pubkey"`                   // the public key of the resource node; bech encoded in JSON
	Suspend        bool           `json:"suspend" yaml:"suspend"`                 // has the resource node been suspended from bonded status?
	Status         sdk.BondStatus `json:"status" yaml:"status"`                   // resource node bond status (bonded/unbonding/unbonded)
	Tokens         sdk.Int        `json:"tokens" yaml:"tokens"`                   // delegated tokens
	OwnerAddress   sdk.AccAddress `json:"owner_address" yaml:"owner_address"`     // owner address of the resource node
	Description    Description    `json:"description" yaml:"description"`         // description terms for the resource node
}

// ResourceNodes is a collection of resource node
type ResourceNodes []ResourceNode

func (v ResourceNodes) String() (out string) {
	for _, node := range v {
		out += node.String() + "\n"
	}
	return strings.TrimSpace(out)
}

// ToSDKResourceNodes -  convenience function convert []ResourceNodes to []sdk.ResourceNodes
func (v ResourceNodes) ToSDKResourceNodes() (resourceNodes []exported.ResourceNodeI) {
	for _, node := range v {
		resourceNodes = append(resourceNodes, node)
	}
	return resourceNodes
}

// Sort ResourceNodes sorts ResourceNode array in ascending owner address order
func (v ResourceNodes) Sort() {
	sort.Sort(v)
}

// Implements sort interface
func (v ResourceNodes) Len() int {
	return len(v)
}

// Implements sort interface
func (v ResourceNodes) Less(i, j int) bool {
	return bytes.Compare(v[i].OwnerAddress, v[j].OwnerAddress) == -1
}

// Implements sort interface
func (v ResourceNodes) Swap(i, j int) {
	it := v[i]
	v[i] = v[j]
	v[j] = it
}

// NewResourceNode - initialize a new resource node
func NewResourceNode(networkAddr string, pubKey crypto.PubKey, ownerAddr sdk.AccAddress, description Description) ResourceNode {
	return ResourceNode{
		NetworkAddress: networkAddr,
		PubKey:         pubKey,
		Suspend:        false,
		Status:         sdk.Unbonded,
		Tokens:         sdk.ZeroInt(),
		OwnerAddress:   ownerAddr,
		Description:    description,
	}
}

// MustMarshalResourceNode returns the resourceNode bytes. Panics if fails
func MustMarshalResourceNode(cdc *codec.Codec, resourceNode ResourceNode) []byte {
	return cdc.MustMarshalBinaryLengthPrefixed(resourceNode)
}

// MustUnmarshalResourceNode unmarshal a resourceNode from a store value. Panics if fails
func MustUnmarshalResourceNode(cdc *codec.Codec, value []byte) ResourceNode {
	resourceNode, err := UnmarshalResourceNode(cdc, value)
	if err != nil {
		panic(err)
	}
	return resourceNode
}

// UnmarshalResourceNode unmarshal a resourceNode from a store value
func UnmarshalResourceNode(cdc *codec.Codec, value []byte) (resourceNode ResourceNode, err error) {
	err = cdc.UnmarshalBinaryLengthPrefixed(value, &resourceNode)
	return resourceNode, err
}

// String returns a human readable string representation of a resource node.
func (v ResourceNode) String() string {
	pubKey, err := sdk.Bech32ifyPubKey(sdk.Bech32PubKeyTypeConsPub, v.PubKey)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf(`ResourceNode:{
		Network Address:	%s
  		Pubkey:				%s
  		Suspend:			%v
  		Status:				%s
  		Tokens:				%s
		Owner Address: 		%s
  		Description:		%s
	}`, v.NetworkAddress, pubKey, v.Suspend, v.Status, v.Tokens, v.OwnerAddress, v.Description)
}

// get the power of the node
// a reduction of 10^6 from node tokens is applied
func (v ResourceNode) GetPower() int64 {
	if v.Status.Equal(sdk.Bonded) {
		return v.PotentialPower()
	}
	return 0
}

// potential power of the node
func (v ResourceNode) PotentialPower() int64 {
	return TokensToPower(v.Tokens)
}

// AddToken adds tokens to a resource node
func (v ResourceNode) AddToken(amount sdk.Int) ResourceNode {
	v.Tokens = v.Tokens.Add(amount)
	if v.Status.Equal(sdk.Unbonded) {
		v.Status = sdk.Bonded
	}
	return v
}

// RemoveToken removes tokens from a resource node
func (v ResourceNode) RemoveToken(tokens sdk.Int) ResourceNode {
	if tokens.IsNegative() {
		panic(fmt.Sprintf("should not happen: trying to remove negative tokens %v", tokens))
	}
	if v.Tokens.LT(tokens) {
		panic(fmt.Sprintf("should not happen: only have %v tokens, trying to remove %v", v.Tokens, tokens))
	}
	v.Tokens = v.Tokens.Sub(tokens)
	return v
}

func (v ResourceNode) IsSuspended() bool            { return v.Suspend }
func (v ResourceNode) GetMoniker() string           { return v.Description.Moniker }
func (v ResourceNode) GetStatus() sdk.BondStatus    { return v.Status }
func (v ResourceNode) GetNetworkAddr() string       { return v.NetworkAddress }
func (v ResourceNode) GetPubKey() crypto.PubKey     { return v.PubKey }
func (v ResourceNode) GetAddr() sdk.AccAddress      { return sdk.AccAddress(v.PubKey.Address()) }
func (v ResourceNode) GetTokens() sdk.Int           { return v.Tokens }
func (v ResourceNode) GetOwnerAddr() sdk.AccAddress { return v.OwnerAddress }
