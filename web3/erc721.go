// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package web3

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// LandContractAddresses is an auto generated low-level Go binding around an user-defined struct.
type LandContractAddresses struct {
	AlphaContract common.Address
	BetaContract  common.Address
	TokenContract common.Address
}

// LandContributorAmount is an auto generated low-level Go binding around an user-defined struct.
type LandContributorAmount struct {
	Contributor common.Address
	Amount      *big.Int
}

// LandLandAmount is an auto generated low-level Go binding around an user-defined struct.
type LandLandAmount struct {
	Alpha      *big.Int
	Beta       *big.Int
	PublicSale *big.Int
	Future     *big.Int
}

// LandMetadata is an auto generated low-level Go binding around an user-defined struct.
type LandMetadata struct {
	MetadataHash      [32]byte
	ShuffledArrayHash [32]byte
	StartIndex        *big.Int
	EndIndex          *big.Int
}

// Erc721MetaData contains all meta data concerning the Erc721 contract.
var Erc721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"alphaContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"betaContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"}],\"internalType\":\"structLand.ContractAddresses\",\"name\":\"addresses\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"alpha\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"beta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicSale\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"future\",\"type\":\"uint256\"}],\"internalType\":\"structLand.LandAmount\",\"name\":\"amount\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structLand.ContributorAmount[]\",\"name\":\"_contributors\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_vrfCoordinator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_linkTokenAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_vrfKeyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_vrfFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"claimableActive\",\"type\":\"bool\"}],\"name\":\"ClaimableStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"ContributorsClaimStart\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"ContributorsClaimStop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_saleDuration\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_saleStartTime\",\"type\":\"uint256\"}],\"name\":\"LandPublicSaleStart\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_currentPrice\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_timeElapsed\",\"type\":\"uint256\"}],\"name\":\"LandPublicSaleStop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"numLands\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"mintPrice\",\"type\":\"uint256\"}],\"name\":\"PublicSaleMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_alphaOffset\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_betaOffset\",\"type\":\"uint256\"}],\"name\":\"StartingIndexSetAlphaBeta\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_startingIndex\",\"type\":\"uint256\"}],\"name\":\"StartingIndexSetPublicSale\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_ALPHA_NFT_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BETA_NFT_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FUTURE_LANDS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_LANDS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_LANDS_WITH_FUTURE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_MINT_PER_BLOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_PUBLIC_SALE_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESERVED_CONTRIBUTORS_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminClaimStarted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"alphaClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alphaClaimedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alphaContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alphaOffset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"betaClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"betaClaimedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"betaContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"betaNftIdCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"betaOffset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"claimUnclaimedAndUnsoldLands\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxAmount\",\"type\":\"uint256\"}],\"name\":\"claimUnclaimedAndUnsoldLandsWithAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimableActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"contributors\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contributorsClaimActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"contributorsClaimLand\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentNumLandsMintedPublicSale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flipClaimableState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"futureLandsNftIdCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"futureMinter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isRandomRequestForPublicSaleAndContributors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kycMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"shuffledArrayHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"internalType\":\"structLand.Metadata\",\"name\":\"_landMetadata\",\"type\":\"tuple\"}],\"name\":\"loadLandMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxMintPerAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxMintPerTx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"metadataHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"shuffledArrayHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"mintFutureLands\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxAmount\",\"type\":\"uint256\"}],\"name\":\"mintFutureLandsWithAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintIndexPublicSaleAndContributors\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numLands\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"mintLands\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintedPerAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"alphaTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"betaTokenIds\",\"type\":\"uint256[]\"}],\"name\":\"nftOwnerClaimLand\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerClaimRandomnessRequested\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicSaleActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicSaleAndContributorsOffset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicSaleAndContributorsRandomnessRequested\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicSaleEndingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicSalePriceLoweringDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicSaleStartPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicSaleStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"shuffledArrayHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"internalType\":\"structLand.Metadata\",\"name\":\"_landMetadata\",\"type\":\"tuple\"}],\"name\":\"putLandMetadataAtIndex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"randomness\",\"type\":\"uint256\"}],\"name\":\"rawFulfillRandomness\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestRandomnessForOwnerClaim\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestRandomnessForPublicSaleAndContributors\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_futureMinter\",\"type\":\"address\"}],\"name\":\"setFutureMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isKycCheckRequired\",\"type\":\"bool\"}],\"name\":\"setKycCheckRequired\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_kycMerkleRoot\",\"type\":\"bytes32\"}],\"name\":\"setKycMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxMintPerAddress\",\"type\":\"uint256\"}],\"name\":\"setMaxMintPerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxMintPerTx\",\"type\":\"uint256\"}],\"name\":\"setMaxMintPerTx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startContributorsClaimPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_publicSalePriceLoweringDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_publicSaleStartPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_publicSaleEndingPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxMintPerTx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxMintPerAddress\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isKycCheckRequired\",\"type\":\"bool\"}],\"name\":\"startPublicSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopContributorsClaimPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopPublicSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Erc721ABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc721MetaData.ABI instead.
var Erc721ABI = Erc721MetaData.ABI

// Erc721 is an auto generated Go binding around an Ethereum contract.
type Erc721 struct {
	Erc721Caller     // Read-only binding to the contract
	Erc721Transactor // Write-only binding to the contract
	Erc721Filterer   // Log filterer for contract events
}

// Erc721Caller is an auto generated read-only Go binding around an Ethereum contract.
type Erc721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc721Session struct {
	Contract     *Erc721           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc721CallerSession struct {
	Contract *Erc721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Erc721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc721TransactorSession struct {
	Contract     *Erc721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc721Raw is an auto generated low-level Go binding around an Ethereum contract.
type Erc721Raw struct {
	Contract *Erc721 // Generic contract binding to access the raw methods on
}

// Erc721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc721CallerRaw struct {
	Contract *Erc721Caller // Generic read-only contract binding to access the raw methods on
}

// Erc721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc721TransactorRaw struct {
	Contract *Erc721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewErc721 creates a new instance of Erc721, bound to a specific deployed contract.
func NewErc721(address common.Address, backend bind.ContractBackend) (*Erc721, error) {
	contract, err := bindErc721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc721{Erc721Caller: Erc721Caller{contract: contract}, Erc721Transactor: Erc721Transactor{contract: contract}, Erc721Filterer: Erc721Filterer{contract: contract}}, nil
}

// NewErc721Caller creates a new read-only instance of Erc721, bound to a specific deployed contract.
func NewErc721Caller(address common.Address, caller bind.ContractCaller) (*Erc721Caller, error) {
	contract, err := bindErc721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc721Caller{contract: contract}, nil
}

// NewErc721Transactor creates a new write-only instance of Erc721, bound to a specific deployed contract.
func NewErc721Transactor(address common.Address, transactor bind.ContractTransactor) (*Erc721Transactor, error) {
	contract, err := bindErc721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc721Transactor{contract: contract}, nil
}

// NewErc721Filterer creates a new log filterer instance of Erc721, bound to a specific deployed contract.
func NewErc721Filterer(address common.Address, filterer bind.ContractFilterer) (*Erc721Filterer, error) {
	contract, err := bindErc721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc721Filterer{contract: contract}, nil
}

// bindErc721 binds a generic wrapper to an already deployed contract.
func bindErc721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc721MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc721 *Erc721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc721.Contract.Erc721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc721 *Erc721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721.Contract.Erc721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc721 *Erc721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc721.Contract.Erc721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc721 *Erc721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc721 *Erc721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc721 *Erc721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc721.Contract.contract.Transact(opts, method, params...)
}

// MAXALPHANFTAMOUNT is a free data retrieval call binding the contract method 0xae851f51.
//
// Solidity: function MAX_ALPHA_NFT_AMOUNT() view returns(uint256)
func (_Erc721 *Erc721Caller) MAXALPHANFTAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "MAX_ALPHA_NFT_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXALPHANFTAMOUNT is a free data retrieval call binding the contract method 0xae851f51.
//
// Solidity: function MAX_ALPHA_NFT_AMOUNT() view returns(uint256)
func (_Erc721 *Erc721Session) MAXALPHANFTAMOUNT() (*big.Int, error) {
	return _Erc721.Contract.MAXALPHANFTAMOUNT(&_Erc721.CallOpts)
}

// MAXALPHANFTAMOUNT is a free data retrieval call binding the contract method 0xae851f51.
//
// Solidity: function MAX_ALPHA_NFT_AMOUNT() view returns(uint256)
func (_Erc721 *Erc721CallerSession) MAXALPHANFTAMOUNT() (*big.Int, error) {
	return _Erc721.Contract.MAXALPHANFTAMOUNT(&_Erc721.CallOpts)
}

// MAXBETANFTAMOUNT is a free data retrieval call binding the contract method 0xec607ce2.
//
// Solidity: function MAX_BETA_NFT_AMOUNT() view returns(uint256)
func (_Erc721 *Erc721Caller) MAXBETANFTAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "MAX_BETA_NFT_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBETANFTAMOUNT is a free data retrieval call binding the contract method 0xec607ce2.
//
// Solidity: function MAX_BETA_NFT_AMOUNT() view returns(uint256)
func (_Erc721 *Erc721Session) MAXBETANFTAMOUNT() (*big.Int, error) {
	return _Erc721.Contract.MAXBETANFTAMOUNT(&_Erc721.CallOpts)
}

// MAXBETANFTAMOUNT is a free data retrieval call binding the contract method 0xec607ce2.
//
// Solidity: function MAX_BETA_NFT_AMOUNT() view returns(uint256)
func (_Erc721 *Erc721CallerSession) MAXBETANFTAMOUNT() (*big.Int, error) {
	return _Erc721.Contract.MAXBETANFTAMOUNT(&_Erc721.CallOpts)
}

// MAXFUTURELANDS is a free data retrieval call binding the contract method 0x0a3ed148.
//
// Solidity: function MAX_FUTURE_LANDS() view returns(uint256)
func (_Erc721 *Erc721Caller) MAXFUTURELANDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "MAX_FUTURE_LANDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFUTURELANDS is a free data retrieval call binding the contract method 0x0a3ed148.
//
// Solidity: function MAX_FUTURE_LANDS() view returns(uint256)
func (_Erc721 *Erc721Session) MAXFUTURELANDS() (*big.Int, error) {
	return _Erc721.Contract.MAXFUTURELANDS(&_Erc721.CallOpts)
}

// MAXFUTURELANDS is a free data retrieval call binding the contract method 0x0a3ed148.
//
// Solidity: function MAX_FUTURE_LANDS() view returns(uint256)
func (_Erc721 *Erc721CallerSession) MAXFUTURELANDS() (*big.Int, error) {
	return _Erc721.Contract.MAXFUTURELANDS(&_Erc721.CallOpts)
}

// MAXLANDS is a free data retrieval call binding the contract method 0xec10abc6.
//
// Solidity: function MAX_LANDS() view returns(uint256)
func (_Erc721 *Erc721Caller) MAXLANDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "MAX_LANDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXLANDS is a free data retrieval call binding the contract method 0xec10abc6.
//
// Solidity: function MAX_LANDS() view returns(uint256)
func (_Erc721 *Erc721Session) MAXLANDS() (*big.Int, error) {
	return _Erc721.Contract.MAXLANDS(&_Erc721.CallOpts)
}

// MAXLANDS is a free data retrieval call binding the contract method 0xec10abc6.
//
// Solidity: function MAX_LANDS() view returns(uint256)
func (_Erc721 *Erc721CallerSession) MAXLANDS() (*big.Int, error) {
	return _Erc721.Contract.MAXLANDS(&_Erc721.CallOpts)
}

// MAXLANDSWITHFUTURE is a free data retrieval call binding the contract method 0x0029d729.
//
// Solidity: function MAX_LANDS_WITH_FUTURE() view returns(uint256)
func (_Erc721 *Erc721Caller) MAXLANDSWITHFUTURE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "MAX_LANDS_WITH_FUTURE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXLANDSWITHFUTURE is a free data retrieval call binding the contract method 0x0029d729.
//
// Solidity: function MAX_LANDS_WITH_FUTURE() view returns(uint256)
func (_Erc721 *Erc721Session) MAXLANDSWITHFUTURE() (*big.Int, error) {
	return _Erc721.Contract.MAXLANDSWITHFUTURE(&_Erc721.CallOpts)
}

// MAXLANDSWITHFUTURE is a free data retrieval call binding the contract method 0x0029d729.
//
// Solidity: function MAX_LANDS_WITH_FUTURE() view returns(uint256)
func (_Erc721 *Erc721CallerSession) MAXLANDSWITHFUTURE() (*big.Int, error) {
	return _Erc721.Contract.MAXLANDSWITHFUTURE(&_Erc721.CallOpts)
}

// MAXMINTPERBLOCK is a free data retrieval call binding the contract method 0xae510a58.
//
// Solidity: function MAX_MINT_PER_BLOCK() view returns(uint256)
func (_Erc721 *Erc721Caller) MAXMINTPERBLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "MAX_MINT_PER_BLOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXMINTPERBLOCK is a free data retrieval call binding the contract method 0xae510a58.
//
// Solidity: function MAX_MINT_PER_BLOCK() view returns(uint256)
func (_Erc721 *Erc721Session) MAXMINTPERBLOCK() (*big.Int, error) {
	return _Erc721.Contract.MAXMINTPERBLOCK(&_Erc721.CallOpts)
}

// MAXMINTPERBLOCK is a free data retrieval call binding the contract method 0xae510a58.
//
// Solidity: function MAX_MINT_PER_BLOCK() view returns(uint256)
func (_Erc721 *Erc721CallerSession) MAXMINTPERBLOCK() (*big.Int, error) {
	return _Erc721.Contract.MAXMINTPERBLOCK(&_Erc721.CallOpts)
}

// MAXPUBLICSALEAMOUNT is a free data retrieval call binding the contract method 0xb48a0539.
//
// Solidity: function MAX_PUBLIC_SALE_AMOUNT() view returns(uint256)
func (_Erc721 *Erc721Caller) MAXPUBLICSALEAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "MAX_PUBLIC_SALE_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPUBLICSALEAMOUNT is a free data retrieval call binding the contract method 0xb48a0539.
//
// Solidity: function MAX_PUBLIC_SALE_AMOUNT() view returns(uint256)
func (_Erc721 *Erc721Session) MAXPUBLICSALEAMOUNT() (*big.Int, error) {
	return _Erc721.Contract.MAXPUBLICSALEAMOUNT(&_Erc721.CallOpts)
}

// MAXPUBLICSALEAMOUNT is a free data retrieval call binding the contract method 0xb48a0539.
//
// Solidity: function MAX_PUBLIC_SALE_AMOUNT() view returns(uint256)
func (_Erc721 *Erc721CallerSession) MAXPUBLICSALEAMOUNT() (*big.Int, error) {
	return _Erc721.Contract.MAXPUBLICSALEAMOUNT(&_Erc721.CallOpts)
}

// RESERVEDCONTRIBUTORSAMOUNT is a free data retrieval call binding the contract method 0xdcae8d87.
//
// Solidity: function RESERVED_CONTRIBUTORS_AMOUNT() view returns(uint256)
func (_Erc721 *Erc721Caller) RESERVEDCONTRIBUTORSAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "RESERVED_CONTRIBUTORS_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RESERVEDCONTRIBUTORSAMOUNT is a free data retrieval call binding the contract method 0xdcae8d87.
//
// Solidity: function RESERVED_CONTRIBUTORS_AMOUNT() view returns(uint256)
func (_Erc721 *Erc721Session) RESERVEDCONTRIBUTORSAMOUNT() (*big.Int, error) {
	return _Erc721.Contract.RESERVEDCONTRIBUTORSAMOUNT(&_Erc721.CallOpts)
}

// RESERVEDCONTRIBUTORSAMOUNT is a free data retrieval call binding the contract method 0xdcae8d87.
//
// Solidity: function RESERVED_CONTRIBUTORS_AMOUNT() view returns(uint256)
func (_Erc721 *Erc721CallerSession) RESERVEDCONTRIBUTORSAMOUNT() (*big.Int, error) {
	return _Erc721.Contract.RESERVEDCONTRIBUTORSAMOUNT(&_Erc721.CallOpts)
}

// AdminClaimStarted is a free data retrieval call binding the contract method 0xb45385bd.
//
// Solidity: function adminClaimStarted() view returns(bool)
func (_Erc721 *Erc721Caller) AdminClaimStarted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "adminClaimStarted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AdminClaimStarted is a free data retrieval call binding the contract method 0xb45385bd.
//
// Solidity: function adminClaimStarted() view returns(bool)
func (_Erc721 *Erc721Session) AdminClaimStarted() (bool, error) {
	return _Erc721.Contract.AdminClaimStarted(&_Erc721.CallOpts)
}

// AdminClaimStarted is a free data retrieval call binding the contract method 0xb45385bd.
//
// Solidity: function adminClaimStarted() view returns(bool)
func (_Erc721 *Erc721CallerSession) AdminClaimStarted() (bool, error) {
	return _Erc721.Contract.AdminClaimStarted(&_Erc721.CallOpts)
}

// AlphaClaimed is a free data retrieval call binding the contract method 0x4f2a7abb.
//
// Solidity: function alphaClaimed(uint256 ) view returns(bool)
func (_Erc721 *Erc721Caller) AlphaClaimed(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "alphaClaimed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AlphaClaimed is a free data retrieval call binding the contract method 0x4f2a7abb.
//
// Solidity: function alphaClaimed(uint256 ) view returns(bool)
func (_Erc721 *Erc721Session) AlphaClaimed(arg0 *big.Int) (bool, error) {
	return _Erc721.Contract.AlphaClaimed(&_Erc721.CallOpts, arg0)
}

// AlphaClaimed is a free data retrieval call binding the contract method 0x4f2a7abb.
//
// Solidity: function alphaClaimed(uint256 ) view returns(bool)
func (_Erc721 *Erc721CallerSession) AlphaClaimed(arg0 *big.Int) (bool, error) {
	return _Erc721.Contract.AlphaClaimed(&_Erc721.CallOpts, arg0)
}

// AlphaClaimedAmount is a free data retrieval call binding the contract method 0x68d41e7d.
//
// Solidity: function alphaClaimedAmount() view returns(uint256)
func (_Erc721 *Erc721Caller) AlphaClaimedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "alphaClaimedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AlphaClaimedAmount is a free data retrieval call binding the contract method 0x68d41e7d.
//
// Solidity: function alphaClaimedAmount() view returns(uint256)
func (_Erc721 *Erc721Session) AlphaClaimedAmount() (*big.Int, error) {
	return _Erc721.Contract.AlphaClaimedAmount(&_Erc721.CallOpts)
}

// AlphaClaimedAmount is a free data retrieval call binding the contract method 0x68d41e7d.
//
// Solidity: function alphaClaimedAmount() view returns(uint256)
func (_Erc721 *Erc721CallerSession) AlphaClaimedAmount() (*big.Int, error) {
	return _Erc721.Contract.AlphaClaimedAmount(&_Erc721.CallOpts)
}

// AlphaContract is a free data retrieval call binding the contract method 0xe867afc0.
//
// Solidity: function alphaContract() view returns(address)
func (_Erc721 *Erc721Caller) AlphaContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "alphaContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AlphaContract is a free data retrieval call binding the contract method 0xe867afc0.
//
// Solidity: function alphaContract() view returns(address)
func (_Erc721 *Erc721Session) AlphaContract() (common.Address, error) {
	return _Erc721.Contract.AlphaContract(&_Erc721.CallOpts)
}

// AlphaContract is a free data retrieval call binding the contract method 0xe867afc0.
//
// Solidity: function alphaContract() view returns(address)
func (_Erc721 *Erc721CallerSession) AlphaContract() (common.Address, error) {
	return _Erc721.Contract.AlphaContract(&_Erc721.CallOpts)
}

// AlphaOffset is a free data retrieval call binding the contract method 0x52a97fc3.
//
// Solidity: function alphaOffset() view returns(uint256)
func (_Erc721 *Erc721Caller) AlphaOffset(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "alphaOffset")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AlphaOffset is a free data retrieval call binding the contract method 0x52a97fc3.
//
// Solidity: function alphaOffset() view returns(uint256)
func (_Erc721 *Erc721Session) AlphaOffset() (*big.Int, error) {
	return _Erc721.Contract.AlphaOffset(&_Erc721.CallOpts)
}

// AlphaOffset is a free data retrieval call binding the contract method 0x52a97fc3.
//
// Solidity: function alphaOffset() view returns(uint256)
func (_Erc721 *Erc721CallerSession) AlphaOffset() (*big.Int, error) {
	return _Erc721.Contract.AlphaOffset(&_Erc721.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Erc721 *Erc721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Erc721 *Erc721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Erc721.Contract.BalanceOf(&_Erc721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Erc721 *Erc721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Erc721.Contract.BalanceOf(&_Erc721.CallOpts, owner)
}

// BetaClaimed is a free data retrieval call binding the contract method 0x2f1f38ae.
//
// Solidity: function betaClaimed(uint256 ) view returns(bool)
func (_Erc721 *Erc721Caller) BetaClaimed(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "betaClaimed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BetaClaimed is a free data retrieval call binding the contract method 0x2f1f38ae.
//
// Solidity: function betaClaimed(uint256 ) view returns(bool)
func (_Erc721 *Erc721Session) BetaClaimed(arg0 *big.Int) (bool, error) {
	return _Erc721.Contract.BetaClaimed(&_Erc721.CallOpts, arg0)
}

// BetaClaimed is a free data retrieval call binding the contract method 0x2f1f38ae.
//
// Solidity: function betaClaimed(uint256 ) view returns(bool)
func (_Erc721 *Erc721CallerSession) BetaClaimed(arg0 *big.Int) (bool, error) {
	return _Erc721.Contract.BetaClaimed(&_Erc721.CallOpts, arg0)
}

// BetaClaimedAmount is a free data retrieval call binding the contract method 0xca694ac8.
//
// Solidity: function betaClaimedAmount() view returns(uint256)
func (_Erc721 *Erc721Caller) BetaClaimedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "betaClaimedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BetaClaimedAmount is a free data retrieval call binding the contract method 0xca694ac8.
//
// Solidity: function betaClaimedAmount() view returns(uint256)
func (_Erc721 *Erc721Session) BetaClaimedAmount() (*big.Int, error) {
	return _Erc721.Contract.BetaClaimedAmount(&_Erc721.CallOpts)
}

// BetaClaimedAmount is a free data retrieval call binding the contract method 0xca694ac8.
//
// Solidity: function betaClaimedAmount() view returns(uint256)
func (_Erc721 *Erc721CallerSession) BetaClaimedAmount() (*big.Int, error) {
	return _Erc721.Contract.BetaClaimedAmount(&_Erc721.CallOpts)
}

// BetaContract is a free data retrieval call binding the contract method 0x2dd98a97.
//
// Solidity: function betaContract() view returns(address)
func (_Erc721 *Erc721Caller) BetaContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "betaContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BetaContract is a free data retrieval call binding the contract method 0x2dd98a97.
//
// Solidity: function betaContract() view returns(address)
func (_Erc721 *Erc721Session) BetaContract() (common.Address, error) {
	return _Erc721.Contract.BetaContract(&_Erc721.CallOpts)
}

// BetaContract is a free data retrieval call binding the contract method 0x2dd98a97.
//
// Solidity: function betaContract() view returns(address)
func (_Erc721 *Erc721CallerSession) BetaContract() (common.Address, error) {
	return _Erc721.Contract.BetaContract(&_Erc721.CallOpts)
}

// BetaNftIdCurrent is a free data retrieval call binding the contract method 0xd0bab933.
//
// Solidity: function betaNftIdCurrent() view returns(uint256)
func (_Erc721 *Erc721Caller) BetaNftIdCurrent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "betaNftIdCurrent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BetaNftIdCurrent is a free data retrieval call binding the contract method 0xd0bab933.
//
// Solidity: function betaNftIdCurrent() view returns(uint256)
func (_Erc721 *Erc721Session) BetaNftIdCurrent() (*big.Int, error) {
	return _Erc721.Contract.BetaNftIdCurrent(&_Erc721.CallOpts)
}

// BetaNftIdCurrent is a free data retrieval call binding the contract method 0xd0bab933.
//
// Solidity: function betaNftIdCurrent() view returns(uint256)
func (_Erc721 *Erc721CallerSession) BetaNftIdCurrent() (*big.Int, error) {
	return _Erc721.Contract.BetaNftIdCurrent(&_Erc721.CallOpts)
}

// BetaOffset is a free data retrieval call binding the contract method 0x497e0f0d.
//
// Solidity: function betaOffset() view returns(uint256)
func (_Erc721 *Erc721Caller) BetaOffset(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "betaOffset")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BetaOffset is a free data retrieval call binding the contract method 0x497e0f0d.
//
// Solidity: function betaOffset() view returns(uint256)
func (_Erc721 *Erc721Session) BetaOffset() (*big.Int, error) {
	return _Erc721.Contract.BetaOffset(&_Erc721.CallOpts)
}

// BetaOffset is a free data retrieval call binding the contract method 0x497e0f0d.
//
// Solidity: function betaOffset() view returns(uint256)
func (_Erc721 *Erc721CallerSession) BetaOffset() (*big.Int, error) {
	return _Erc721.Contract.BetaOffset(&_Erc721.CallOpts)
}

// ClaimableActive is a free data retrieval call binding the contract method 0xd04db50e.
//
// Solidity: function claimableActive() view returns(bool)
func (_Erc721 *Erc721Caller) ClaimableActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "claimableActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ClaimableActive is a free data retrieval call binding the contract method 0xd04db50e.
//
// Solidity: function claimableActive() view returns(bool)
func (_Erc721 *Erc721Session) ClaimableActive() (bool, error) {
	return _Erc721.Contract.ClaimableActive(&_Erc721.CallOpts)
}

// ClaimableActive is a free data retrieval call binding the contract method 0xd04db50e.
//
// Solidity: function claimableActive() view returns(bool)
func (_Erc721 *Erc721CallerSession) ClaimableActive() (bool, error) {
	return _Erc721.Contract.ClaimableActive(&_Erc721.CallOpts)
}

// Contributors is a free data retrieval call binding the contract method 0x1f6d4942.
//
// Solidity: function contributors(address ) view returns(uint256)
func (_Erc721 *Erc721Caller) Contributors(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "contributors", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Contributors is a free data retrieval call binding the contract method 0x1f6d4942.
//
// Solidity: function contributors(address ) view returns(uint256)
func (_Erc721 *Erc721Session) Contributors(arg0 common.Address) (*big.Int, error) {
	return _Erc721.Contract.Contributors(&_Erc721.CallOpts, arg0)
}

// Contributors is a free data retrieval call binding the contract method 0x1f6d4942.
//
// Solidity: function contributors(address ) view returns(uint256)
func (_Erc721 *Erc721CallerSession) Contributors(arg0 common.Address) (*big.Int, error) {
	return _Erc721.Contract.Contributors(&_Erc721.CallOpts, arg0)
}

// ContributorsClaimActive is a free data retrieval call binding the contract method 0x5668aca0.
//
// Solidity: function contributorsClaimActive() view returns(bool)
func (_Erc721 *Erc721Caller) ContributorsClaimActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "contributorsClaimActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContributorsClaimActive is a free data retrieval call binding the contract method 0x5668aca0.
//
// Solidity: function contributorsClaimActive() view returns(bool)
func (_Erc721 *Erc721Session) ContributorsClaimActive() (bool, error) {
	return _Erc721.Contract.ContributorsClaimActive(&_Erc721.CallOpts)
}

// ContributorsClaimActive is a free data retrieval call binding the contract method 0x5668aca0.
//
// Solidity: function contributorsClaimActive() view returns(bool)
func (_Erc721 *Erc721CallerSession) ContributorsClaimActive() (bool, error) {
	return _Erc721.Contract.ContributorsClaimActive(&_Erc721.CallOpts)
}

// CurrentNumLandsMintedPublicSale is a free data retrieval call binding the contract method 0x0fa57d8a.
//
// Solidity: function currentNumLandsMintedPublicSale() view returns(uint256)
func (_Erc721 *Erc721Caller) CurrentNumLandsMintedPublicSale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "currentNumLandsMintedPublicSale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentNumLandsMintedPublicSale is a free data retrieval call binding the contract method 0x0fa57d8a.
//
// Solidity: function currentNumLandsMintedPublicSale() view returns(uint256)
func (_Erc721 *Erc721Session) CurrentNumLandsMintedPublicSale() (*big.Int, error) {
	return _Erc721.Contract.CurrentNumLandsMintedPublicSale(&_Erc721.CallOpts)
}

// CurrentNumLandsMintedPublicSale is a free data retrieval call binding the contract method 0x0fa57d8a.
//
// Solidity: function currentNumLandsMintedPublicSale() view returns(uint256)
func (_Erc721 *Erc721CallerSession) CurrentNumLandsMintedPublicSale() (*big.Int, error) {
	return _Erc721.Contract.CurrentNumLandsMintedPublicSale(&_Erc721.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Erc721 *Erc721Caller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Erc721 *Erc721Session) Fee() (*big.Int, error) {
	return _Erc721.Contract.Fee(&_Erc721.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Erc721 *Erc721CallerSession) Fee() (*big.Int, error) {
	return _Erc721.Contract.Fee(&_Erc721.CallOpts)
}

// FutureLandsNftIdCurrent is a free data retrieval call binding the contract method 0x05084e6b.
//
// Solidity: function futureLandsNftIdCurrent() view returns(uint256)
func (_Erc721 *Erc721Caller) FutureLandsNftIdCurrent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "futureLandsNftIdCurrent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureLandsNftIdCurrent is a free data retrieval call binding the contract method 0x05084e6b.
//
// Solidity: function futureLandsNftIdCurrent() view returns(uint256)
func (_Erc721 *Erc721Session) FutureLandsNftIdCurrent() (*big.Int, error) {
	return _Erc721.Contract.FutureLandsNftIdCurrent(&_Erc721.CallOpts)
}

// FutureLandsNftIdCurrent is a free data retrieval call binding the contract method 0x05084e6b.
//
// Solidity: function futureLandsNftIdCurrent() view returns(uint256)
func (_Erc721 *Erc721CallerSession) FutureLandsNftIdCurrent() (*big.Int, error) {
	return _Erc721.Contract.FutureLandsNftIdCurrent(&_Erc721.CallOpts)
}

// FutureMinter is a free data retrieval call binding the contract method 0x7951074a.
//
// Solidity: function futureMinter() view returns(address)
func (_Erc721 *Erc721Caller) FutureMinter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "futureMinter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureMinter is a free data retrieval call binding the contract method 0x7951074a.
//
// Solidity: function futureMinter() view returns(address)
func (_Erc721 *Erc721Session) FutureMinter() (common.Address, error) {
	return _Erc721.Contract.FutureMinter(&_Erc721.CallOpts)
}

// FutureMinter is a free data retrieval call binding the contract method 0x7951074a.
//
// Solidity: function futureMinter() view returns(address)
func (_Erc721 *Erc721CallerSession) FutureMinter() (common.Address, error) {
	return _Erc721.Contract.FutureMinter(&_Erc721.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Erc721 *Erc721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Erc721 *Erc721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Erc721.Contract.GetApproved(&_Erc721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Erc721 *Erc721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Erc721.Contract.GetApproved(&_Erc721.CallOpts, tokenId)
}

// GetMintPrice is a free data retrieval call binding the contract method 0xa7f93ebd.
//
// Solidity: function getMintPrice() view returns(uint256)
func (_Erc721 *Erc721Caller) GetMintPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "getMintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMintPrice is a free data retrieval call binding the contract method 0xa7f93ebd.
//
// Solidity: function getMintPrice() view returns(uint256)
func (_Erc721 *Erc721Session) GetMintPrice() (*big.Int, error) {
	return _Erc721.Contract.GetMintPrice(&_Erc721.CallOpts)
}

// GetMintPrice is a free data retrieval call binding the contract method 0xa7f93ebd.
//
// Solidity: function getMintPrice() view returns(uint256)
func (_Erc721 *Erc721CallerSession) GetMintPrice() (*big.Int, error) {
	return _Erc721.Contract.GetMintPrice(&_Erc721.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Erc721 *Erc721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Erc721 *Erc721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Erc721.Contract.IsApprovedForAll(&_Erc721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Erc721 *Erc721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Erc721.Contract.IsApprovedForAll(&_Erc721.CallOpts, owner, operator)
}

// IsRandomRequestForPublicSaleAndContributors is a free data retrieval call binding the contract method 0xab0752d5.
//
// Solidity: function isRandomRequestForPublicSaleAndContributors(bytes32 ) view returns(bool)
func (_Erc721 *Erc721Caller) IsRandomRequestForPublicSaleAndContributors(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "isRandomRequestForPublicSaleAndContributors", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRandomRequestForPublicSaleAndContributors is a free data retrieval call binding the contract method 0xab0752d5.
//
// Solidity: function isRandomRequestForPublicSaleAndContributors(bytes32 ) view returns(bool)
func (_Erc721 *Erc721Session) IsRandomRequestForPublicSaleAndContributors(arg0 [32]byte) (bool, error) {
	return _Erc721.Contract.IsRandomRequestForPublicSaleAndContributors(&_Erc721.CallOpts, arg0)
}

// IsRandomRequestForPublicSaleAndContributors is a free data retrieval call binding the contract method 0xab0752d5.
//
// Solidity: function isRandomRequestForPublicSaleAndContributors(bytes32 ) view returns(bool)
func (_Erc721 *Erc721CallerSession) IsRandomRequestForPublicSaleAndContributors(arg0 [32]byte) (bool, error) {
	return _Erc721.Contract.IsRandomRequestForPublicSaleAndContributors(&_Erc721.CallOpts, arg0)
}

// KeyHash is a free data retrieval call binding the contract method 0x61728f39.
//
// Solidity: function keyHash() view returns(bytes32)
func (_Erc721 *Erc721Caller) KeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "keyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// KeyHash is a free data retrieval call binding the contract method 0x61728f39.
//
// Solidity: function keyHash() view returns(bytes32)
func (_Erc721 *Erc721Session) KeyHash() ([32]byte, error) {
	return _Erc721.Contract.KeyHash(&_Erc721.CallOpts)
}

// KeyHash is a free data retrieval call binding the contract method 0x61728f39.
//
// Solidity: function keyHash() view returns(bytes32)
func (_Erc721 *Erc721CallerSession) KeyHash() ([32]byte, error) {
	return _Erc721.Contract.KeyHash(&_Erc721.CallOpts)
}

// KycMerkleRoot is a free data retrieval call binding the contract method 0xc324a2c2.
//
// Solidity: function kycMerkleRoot() view returns(bytes32)
func (_Erc721 *Erc721Caller) KycMerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "kycMerkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// KycMerkleRoot is a free data retrieval call binding the contract method 0xc324a2c2.
//
// Solidity: function kycMerkleRoot() view returns(bytes32)
func (_Erc721 *Erc721Session) KycMerkleRoot() ([32]byte, error) {
	return _Erc721.Contract.KycMerkleRoot(&_Erc721.CallOpts)
}

// KycMerkleRoot is a free data retrieval call binding the contract method 0xc324a2c2.
//
// Solidity: function kycMerkleRoot() view returns(bytes32)
func (_Erc721 *Erc721CallerSession) KycMerkleRoot() ([32]byte, error) {
	return _Erc721.Contract.KycMerkleRoot(&_Erc721.CallOpts)
}

// MaxMintPerAddress is a free data retrieval call binding the contract method 0x572849c4.
//
// Solidity: function maxMintPerAddress() view returns(uint256)
func (_Erc721 *Erc721Caller) MaxMintPerAddress(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "maxMintPerAddress")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMintPerAddress is a free data retrieval call binding the contract method 0x572849c4.
//
// Solidity: function maxMintPerAddress() view returns(uint256)
func (_Erc721 *Erc721Session) MaxMintPerAddress() (*big.Int, error) {
	return _Erc721.Contract.MaxMintPerAddress(&_Erc721.CallOpts)
}

// MaxMintPerAddress is a free data retrieval call binding the contract method 0x572849c4.
//
// Solidity: function maxMintPerAddress() view returns(uint256)
func (_Erc721 *Erc721CallerSession) MaxMintPerAddress() (*big.Int, error) {
	return _Erc721.Contract.MaxMintPerAddress(&_Erc721.CallOpts)
}

// MaxMintPerTx is a free data retrieval call binding the contract method 0xde7fcb1d.
//
// Solidity: function maxMintPerTx() view returns(uint256)
func (_Erc721 *Erc721Caller) MaxMintPerTx(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "maxMintPerTx")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMintPerTx is a free data retrieval call binding the contract method 0xde7fcb1d.
//
// Solidity: function maxMintPerTx() view returns(uint256)
func (_Erc721 *Erc721Session) MaxMintPerTx() (*big.Int, error) {
	return _Erc721.Contract.MaxMintPerTx(&_Erc721.CallOpts)
}

// MaxMintPerTx is a free data retrieval call binding the contract method 0xde7fcb1d.
//
// Solidity: function maxMintPerTx() view returns(uint256)
func (_Erc721 *Erc721CallerSession) MaxMintPerTx() (*big.Int, error) {
	return _Erc721.Contract.MaxMintPerTx(&_Erc721.CallOpts)
}

// MetadataHashes is a free data retrieval call binding the contract method 0xf9fd78c8.
//
// Solidity: function metadataHashes(uint256 ) view returns(bytes32 metadataHash, bytes32 shuffledArrayHash, uint256 startIndex, uint256 endIndex)
func (_Erc721 *Erc721Caller) MetadataHashes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	MetadataHash      [32]byte
	ShuffledArrayHash [32]byte
	StartIndex        *big.Int
	EndIndex          *big.Int
}, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "metadataHashes", arg0)

	outstruct := new(struct {
		MetadataHash      [32]byte
		ShuffledArrayHash [32]byte
		StartIndex        *big.Int
		EndIndex          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MetadataHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.ShuffledArrayHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.StartIndex = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndIndex = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MetadataHashes is a free data retrieval call binding the contract method 0xf9fd78c8.
//
// Solidity: function metadataHashes(uint256 ) view returns(bytes32 metadataHash, bytes32 shuffledArrayHash, uint256 startIndex, uint256 endIndex)
func (_Erc721 *Erc721Session) MetadataHashes(arg0 *big.Int) (struct {
	MetadataHash      [32]byte
	ShuffledArrayHash [32]byte
	StartIndex        *big.Int
	EndIndex          *big.Int
}, error) {
	return _Erc721.Contract.MetadataHashes(&_Erc721.CallOpts, arg0)
}

// MetadataHashes is a free data retrieval call binding the contract method 0xf9fd78c8.
//
// Solidity: function metadataHashes(uint256 ) view returns(bytes32 metadataHash, bytes32 shuffledArrayHash, uint256 startIndex, uint256 endIndex)
func (_Erc721 *Erc721CallerSession) MetadataHashes(arg0 *big.Int) (struct {
	MetadataHash      [32]byte
	ShuffledArrayHash [32]byte
	StartIndex        *big.Int
	EndIndex          *big.Int
}, error) {
	return _Erc721.Contract.MetadataHashes(&_Erc721.CallOpts, arg0)
}

// MintIndexPublicSaleAndContributors is a free data retrieval call binding the contract method 0x4cbe9043.
//
// Solidity: function mintIndexPublicSaleAndContributors() view returns(uint256)
func (_Erc721 *Erc721Caller) MintIndexPublicSaleAndContributors(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "mintIndexPublicSaleAndContributors")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintIndexPublicSaleAndContributors is a free data retrieval call binding the contract method 0x4cbe9043.
//
// Solidity: function mintIndexPublicSaleAndContributors() view returns(uint256)
func (_Erc721 *Erc721Session) MintIndexPublicSaleAndContributors() (*big.Int, error) {
	return _Erc721.Contract.MintIndexPublicSaleAndContributors(&_Erc721.CallOpts)
}

// MintIndexPublicSaleAndContributors is a free data retrieval call binding the contract method 0x4cbe9043.
//
// Solidity: function mintIndexPublicSaleAndContributors() view returns(uint256)
func (_Erc721 *Erc721CallerSession) MintIndexPublicSaleAndContributors() (*big.Int, error) {
	return _Erc721.Contract.MintIndexPublicSaleAndContributors(&_Erc721.CallOpts)
}

// MintedPerAddress is a free data retrieval call binding the contract method 0xd445b978.
//
// Solidity: function mintedPerAddress(address ) view returns(uint256)
func (_Erc721 *Erc721Caller) MintedPerAddress(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "mintedPerAddress", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintedPerAddress is a free data retrieval call binding the contract method 0xd445b978.
//
// Solidity: function mintedPerAddress(address ) view returns(uint256)
func (_Erc721 *Erc721Session) MintedPerAddress(arg0 common.Address) (*big.Int, error) {
	return _Erc721.Contract.MintedPerAddress(&_Erc721.CallOpts, arg0)
}

// MintedPerAddress is a free data retrieval call binding the contract method 0xd445b978.
//
// Solidity: function mintedPerAddress(address ) view returns(uint256)
func (_Erc721 *Erc721CallerSession) MintedPerAddress(arg0 common.Address) (*big.Int, error) {
	return _Erc721.Contract.MintedPerAddress(&_Erc721.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc721 *Erc721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc721 *Erc721Session) Name() (string, error) {
	return _Erc721.Contract.Name(&_Erc721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc721 *Erc721CallerSession) Name() (string, error) {
	return _Erc721.Contract.Name(&_Erc721.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Erc721 *Erc721Caller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Erc721 *Erc721Session) Operator() (common.Address, error) {
	return _Erc721.Contract.Operator(&_Erc721.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Erc721 *Erc721CallerSession) Operator() (common.Address, error) {
	return _Erc721.Contract.Operator(&_Erc721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc721 *Erc721Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc721 *Erc721Session) Owner() (common.Address, error) {
	return _Erc721.Contract.Owner(&_Erc721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc721 *Erc721CallerSession) Owner() (common.Address, error) {
	return _Erc721.Contract.Owner(&_Erc721.CallOpts)
}

// OwnerClaimRandomnessRequested is a free data retrieval call binding the contract method 0x396d91b5.
//
// Solidity: function ownerClaimRandomnessRequested() view returns(bool)
func (_Erc721 *Erc721Caller) OwnerClaimRandomnessRequested(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "ownerClaimRandomnessRequested")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OwnerClaimRandomnessRequested is a free data retrieval call binding the contract method 0x396d91b5.
//
// Solidity: function ownerClaimRandomnessRequested() view returns(bool)
func (_Erc721 *Erc721Session) OwnerClaimRandomnessRequested() (bool, error) {
	return _Erc721.Contract.OwnerClaimRandomnessRequested(&_Erc721.CallOpts)
}

// OwnerClaimRandomnessRequested is a free data retrieval call binding the contract method 0x396d91b5.
//
// Solidity: function ownerClaimRandomnessRequested() view returns(bool)
func (_Erc721 *Erc721CallerSession) OwnerClaimRandomnessRequested() (bool, error) {
	return _Erc721.Contract.OwnerClaimRandomnessRequested(&_Erc721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Erc721 *Erc721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Erc721 *Erc721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Erc721.Contract.OwnerOf(&_Erc721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Erc721 *Erc721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Erc721.Contract.OwnerOf(&_Erc721.CallOpts, tokenId)
}

// PublicSaleActive is a free data retrieval call binding the contract method 0xbc8893b4.
//
// Solidity: function publicSaleActive() view returns(bool)
func (_Erc721 *Erc721Caller) PublicSaleActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "publicSaleActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PublicSaleActive is a free data retrieval call binding the contract method 0xbc8893b4.
//
// Solidity: function publicSaleActive() view returns(bool)
func (_Erc721 *Erc721Session) PublicSaleActive() (bool, error) {
	return _Erc721.Contract.PublicSaleActive(&_Erc721.CallOpts)
}

// PublicSaleActive is a free data retrieval call binding the contract method 0xbc8893b4.
//
// Solidity: function publicSaleActive() view returns(bool)
func (_Erc721 *Erc721CallerSession) PublicSaleActive() (bool, error) {
	return _Erc721.Contract.PublicSaleActive(&_Erc721.CallOpts)
}

// PublicSaleAndContributorsOffset is a free data retrieval call binding the contract method 0xbcd3a192.
//
// Solidity: function publicSaleAndContributorsOffset() view returns(uint256)
func (_Erc721 *Erc721Caller) PublicSaleAndContributorsOffset(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "publicSaleAndContributorsOffset")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicSaleAndContributorsOffset is a free data retrieval call binding the contract method 0xbcd3a192.
//
// Solidity: function publicSaleAndContributorsOffset() view returns(uint256)
func (_Erc721 *Erc721Session) PublicSaleAndContributorsOffset() (*big.Int, error) {
	return _Erc721.Contract.PublicSaleAndContributorsOffset(&_Erc721.CallOpts)
}

// PublicSaleAndContributorsOffset is a free data retrieval call binding the contract method 0xbcd3a192.
//
// Solidity: function publicSaleAndContributorsOffset() view returns(uint256)
func (_Erc721 *Erc721CallerSession) PublicSaleAndContributorsOffset() (*big.Int, error) {
	return _Erc721.Contract.PublicSaleAndContributorsOffset(&_Erc721.CallOpts)
}

// PublicSaleAndContributorsRandomnessRequested is a free data retrieval call binding the contract method 0x61eede53.
//
// Solidity: function publicSaleAndContributorsRandomnessRequested() view returns(bool)
func (_Erc721 *Erc721Caller) PublicSaleAndContributorsRandomnessRequested(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "publicSaleAndContributorsRandomnessRequested")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PublicSaleAndContributorsRandomnessRequested is a free data retrieval call binding the contract method 0x61eede53.
//
// Solidity: function publicSaleAndContributorsRandomnessRequested() view returns(bool)
func (_Erc721 *Erc721Session) PublicSaleAndContributorsRandomnessRequested() (bool, error) {
	return _Erc721.Contract.PublicSaleAndContributorsRandomnessRequested(&_Erc721.CallOpts)
}

// PublicSaleAndContributorsRandomnessRequested is a free data retrieval call binding the contract method 0x61eede53.
//
// Solidity: function publicSaleAndContributorsRandomnessRequested() view returns(bool)
func (_Erc721 *Erc721CallerSession) PublicSaleAndContributorsRandomnessRequested() (bool, error) {
	return _Erc721.Contract.PublicSaleAndContributorsRandomnessRequested(&_Erc721.CallOpts)
}

// PublicSaleEndingPrice is a free data retrieval call binding the contract method 0xb62147e5.
//
// Solidity: function publicSaleEndingPrice() view returns(uint256)
func (_Erc721 *Erc721Caller) PublicSaleEndingPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "publicSaleEndingPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicSaleEndingPrice is a free data retrieval call binding the contract method 0xb62147e5.
//
// Solidity: function publicSaleEndingPrice() view returns(uint256)
func (_Erc721 *Erc721Session) PublicSaleEndingPrice() (*big.Int, error) {
	return _Erc721.Contract.PublicSaleEndingPrice(&_Erc721.CallOpts)
}

// PublicSaleEndingPrice is a free data retrieval call binding the contract method 0xb62147e5.
//
// Solidity: function publicSaleEndingPrice() view returns(uint256)
func (_Erc721 *Erc721CallerSession) PublicSaleEndingPrice() (*big.Int, error) {
	return _Erc721.Contract.PublicSaleEndingPrice(&_Erc721.CallOpts)
}

// PublicSalePriceLoweringDuration is a free data retrieval call binding the contract method 0x745ac965.
//
// Solidity: function publicSalePriceLoweringDuration() view returns(uint256)
func (_Erc721 *Erc721Caller) PublicSalePriceLoweringDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "publicSalePriceLoweringDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicSalePriceLoweringDuration is a free data retrieval call binding the contract method 0x745ac965.
//
// Solidity: function publicSalePriceLoweringDuration() view returns(uint256)
func (_Erc721 *Erc721Session) PublicSalePriceLoweringDuration() (*big.Int, error) {
	return _Erc721.Contract.PublicSalePriceLoweringDuration(&_Erc721.CallOpts)
}

// PublicSalePriceLoweringDuration is a free data retrieval call binding the contract method 0x745ac965.
//
// Solidity: function publicSalePriceLoweringDuration() view returns(uint256)
func (_Erc721 *Erc721CallerSession) PublicSalePriceLoweringDuration() (*big.Int, error) {
	return _Erc721.Contract.PublicSalePriceLoweringDuration(&_Erc721.CallOpts)
}

// PublicSaleStartPrice is a free data retrieval call binding the contract method 0x6faaf624.
//
// Solidity: function publicSaleStartPrice() view returns(uint256)
func (_Erc721 *Erc721Caller) PublicSaleStartPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "publicSaleStartPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicSaleStartPrice is a free data retrieval call binding the contract method 0x6faaf624.
//
// Solidity: function publicSaleStartPrice() view returns(uint256)
func (_Erc721 *Erc721Session) PublicSaleStartPrice() (*big.Int, error) {
	return _Erc721.Contract.PublicSaleStartPrice(&_Erc721.CallOpts)
}

// PublicSaleStartPrice is a free data retrieval call binding the contract method 0x6faaf624.
//
// Solidity: function publicSaleStartPrice() view returns(uint256)
func (_Erc721 *Erc721CallerSession) PublicSaleStartPrice() (*big.Int, error) {
	return _Erc721.Contract.PublicSaleStartPrice(&_Erc721.CallOpts)
}

// PublicSaleStartTime is a free data retrieval call binding the contract method 0x6bb7b1d9.
//
// Solidity: function publicSaleStartTime() view returns(uint256)
func (_Erc721 *Erc721Caller) PublicSaleStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "publicSaleStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicSaleStartTime is a free data retrieval call binding the contract method 0x6bb7b1d9.
//
// Solidity: function publicSaleStartTime() view returns(uint256)
func (_Erc721 *Erc721Session) PublicSaleStartTime() (*big.Int, error) {
	return _Erc721.Contract.PublicSaleStartTime(&_Erc721.CallOpts)
}

// PublicSaleStartTime is a free data retrieval call binding the contract method 0x6bb7b1d9.
//
// Solidity: function publicSaleStartTime() view returns(uint256)
func (_Erc721 *Erc721CallerSession) PublicSaleStartTime() (*big.Int, error) {
	return _Erc721.Contract.PublicSaleStartTime(&_Erc721.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc721 *Erc721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc721 *Erc721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc721.Contract.SupportsInterface(&_Erc721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc721 *Erc721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc721.Contract.SupportsInterface(&_Erc721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc721 *Erc721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc721 *Erc721Session) Symbol() (string, error) {
	return _Erc721.Contract.Symbol(&_Erc721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc721 *Erc721CallerSession) Symbol() (string, error) {
	return _Erc721.Contract.Symbol(&_Erc721.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Erc721 *Erc721Caller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Erc721 *Erc721Session) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Erc721.Contract.TokenByIndex(&_Erc721.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Erc721 *Erc721CallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Erc721.Contract.TokenByIndex(&_Erc721.CallOpts, index)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() view returns(address)
func (_Erc721 *Erc721Caller) TokenContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "tokenContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() view returns(address)
func (_Erc721 *Erc721Session) TokenContract() (common.Address, error) {
	return _Erc721.Contract.TokenContract(&_Erc721.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() view returns(address)
func (_Erc721 *Erc721CallerSession) TokenContract() (common.Address, error) {
	return _Erc721.Contract.TokenContract(&_Erc721.CallOpts)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Erc721 *Erc721Caller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Erc721 *Erc721Session) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Erc721.Contract.TokenOfOwnerByIndex(&_Erc721.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Erc721 *Erc721CallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Erc721.Contract.TokenOfOwnerByIndex(&_Erc721.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Erc721 *Erc721Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Erc721 *Erc721Session) TokenURI(tokenId *big.Int) (string, error) {
	return _Erc721.Contract.TokenURI(&_Erc721.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Erc721 *Erc721CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Erc721.Contract.TokenURI(&_Erc721.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc721 *Erc721Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc721 *Erc721Session) TotalSupply() (*big.Int, error) {
	return _Erc721.Contract.TotalSupply(&_Erc721.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc721 *Erc721CallerSession) TotalSupply() (*big.Int, error) {
	return _Erc721.Contract.TotalSupply(&_Erc721.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Erc721 *Erc721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Erc721 *Erc721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721.Contract.Approve(&_Erc721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Erc721 *Erc721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721.Contract.Approve(&_Erc721.TransactOpts, to, tokenId)
}

// ClaimUnclaimedAndUnsoldLands is a paid mutator transaction binding the contract method 0x401a2ab9.
//
// Solidity: function claimUnclaimedAndUnsoldLands(address recipient) returns()
func (_Erc721 *Erc721Transactor) ClaimUnclaimedAndUnsoldLands(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "claimUnclaimedAndUnsoldLands", recipient)
}

// ClaimUnclaimedAndUnsoldLands is a paid mutator transaction binding the contract method 0x401a2ab9.
//
// Solidity: function claimUnclaimedAndUnsoldLands(address recipient) returns()
func (_Erc721 *Erc721Session) ClaimUnclaimedAndUnsoldLands(recipient common.Address) (*types.Transaction, error) {
	return _Erc721.Contract.ClaimUnclaimedAndUnsoldLands(&_Erc721.TransactOpts, recipient)
}

// ClaimUnclaimedAndUnsoldLands is a paid mutator transaction binding the contract method 0x401a2ab9.
//
// Solidity: function claimUnclaimedAndUnsoldLands(address recipient) returns()
func (_Erc721 *Erc721TransactorSession) ClaimUnclaimedAndUnsoldLands(recipient common.Address) (*types.Transaction, error) {
	return _Erc721.Contract.ClaimUnclaimedAndUnsoldLands(&_Erc721.TransactOpts, recipient)
}

// ClaimUnclaimedAndUnsoldLandsWithAmount is a paid mutator transaction binding the contract method 0x372854e4.
//
// Solidity: function claimUnclaimedAndUnsoldLandsWithAmount(address recipient, uint256 maxAmount) returns()
func (_Erc721 *Erc721Transactor) ClaimUnclaimedAndUnsoldLandsWithAmount(opts *bind.TransactOpts, recipient common.Address, maxAmount *big.Int) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "claimUnclaimedAndUnsoldLandsWithAmount", recipient, maxAmount)
}

// ClaimUnclaimedAndUnsoldLandsWithAmount is a paid mutator transaction binding the contract method 0x372854e4.
//
// Solidity: function claimUnclaimedAndUnsoldLandsWithAmount(address recipient, uint256 maxAmount) returns()
func (_Erc721 *Erc721Session) ClaimUnclaimedAndUnsoldLandsWithAmount(recipient common.Address, maxAmount *big.Int) (*types.Transaction, error) {
	return _Erc721.Contract.ClaimUnclaimedAndUnsoldLandsWithAmount(&_Erc721.TransactOpts, recipient, maxAmount)
}

// ClaimUnclaimedAndUnsoldLandsWithAmount is a paid mutator transaction binding the contract method 0x372854e4.
//
// Solidity: function claimUnclaimedAndUnsoldLandsWithAmount(address recipient, uint256 maxAmount) returns()
func (_Erc721 *Erc721TransactorSession) ClaimUnclaimedAndUnsoldLandsWithAmount(recipient common.Address, maxAmount *big.Int) (*types.Transaction, error) {
	return _Erc721.Contract.ClaimUnclaimedAndUnsoldLandsWithAmount(&_Erc721.TransactOpts, recipient, maxAmount)
}

// ContributorsClaimLand is a paid mutator transaction binding the contract method 0x064e144f.
//
// Solidity: function contributorsClaimLand(uint256 amount, address recipient) returns()
func (_Erc721 *Erc721Transactor) ContributorsClaimLand(opts *bind.TransactOpts, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "contributorsClaimLand", amount, recipient)
}

// ContributorsClaimLand is a paid mutator transaction binding the contract method 0x064e144f.
//
// Solidity: function contributorsClaimLand(uint256 amount, address recipient) returns()
func (_Erc721 *Erc721Session) ContributorsClaimLand(amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Erc721.Contract.ContributorsClaimLand(&_Erc721.TransactOpts, amount, recipient)
}

// ContributorsClaimLand is a paid mutator transaction binding the contract method 0x064e144f.
//
// Solidity: function contributorsClaimLand(uint256 amount, address recipient) returns()
func (_Erc721 *Erc721TransactorSession) ContributorsClaimLand(amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Erc721.Contract.ContributorsClaimLand(&_Erc721.TransactOpts, amount, recipient)
}

// FlipClaimableState is a paid mutator transaction binding the contract method 0x5006f20a.
//
// Solidity: function flipClaimableState() returns()
func (_Erc721 *Erc721Transactor) FlipClaimableState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "flipClaimableState")
}

// FlipClaimableState is a paid mutator transaction binding the contract method 0x5006f20a.
//
// Solidity: function flipClaimableState() returns()
func (_Erc721 *Erc721Session) FlipClaimableState() (*types.Transaction, error) {
	return _Erc721.Contract.FlipClaimableState(&_Erc721.TransactOpts)
}

// FlipClaimableState is a paid mutator transaction binding the contract method 0x5006f20a.
//
// Solidity: function flipClaimableState() returns()
func (_Erc721 *Erc721TransactorSession) FlipClaimableState() (*types.Transaction, error) {
	return _Erc721.Contract.FlipClaimableState(&_Erc721.TransactOpts)
}

// LoadLandMetadata is a paid mutator transaction binding the contract method 0xe58537f4.
//
// Solidity: function loadLandMetadata((bytes32,bytes32,uint256,uint256) _landMetadata) returns()
func (_Erc721 *Erc721Transactor) LoadLandMetadata(opts *bind.TransactOpts, _landMetadata LandMetadata) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "loadLandMetadata", _landMetadata)
}

// LoadLandMetadata is a paid mutator transaction binding the contract method 0xe58537f4.
//
// Solidity: function loadLandMetadata((bytes32,bytes32,uint256,uint256) _landMetadata) returns()
func (_Erc721 *Erc721Session) LoadLandMetadata(_landMetadata LandMetadata) (*types.Transaction, error) {
	return _Erc721.Contract.LoadLandMetadata(&_Erc721.TransactOpts, _landMetadata)
}

// LoadLandMetadata is a paid mutator transaction binding the contract method 0xe58537f4.
//
// Solidity: function loadLandMetadata((bytes32,bytes32,uint256,uint256) _landMetadata) returns()
func (_Erc721 *Erc721TransactorSession) LoadLandMetadata(_landMetadata LandMetadata) (*types.Transaction, error) {
	return _Erc721.Contract.LoadLandMetadata(&_Erc721.TransactOpts, _landMetadata)
}

// MintFutureLands is a paid mutator transaction binding the contract method 0xebc113de.
//
// Solidity: function mintFutureLands(address recipient) returns()
func (_Erc721 *Erc721Transactor) MintFutureLands(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "mintFutureLands", recipient)
}

// MintFutureLands is a paid mutator transaction binding the contract method 0xebc113de.
//
// Solidity: function mintFutureLands(address recipient) returns()
func (_Erc721 *Erc721Session) MintFutureLands(recipient common.Address) (*types.Transaction, error) {
	return _Erc721.Contract.MintFutureLands(&_Erc721.TransactOpts, recipient)
}

// MintFutureLands is a paid mutator transaction binding the contract method 0xebc113de.
//
// Solidity: function mintFutureLands(address recipient) returns()
func (_Erc721 *Erc721TransactorSession) MintFutureLands(recipient common.Address) (*types.Transaction, error) {
	return _Erc721.Contract.MintFutureLands(&_Erc721.TransactOpts, recipient)
}

// MintFutureLandsWithAmount is a paid mutator transaction binding the contract method 0x71700b56.
//
// Solidity: function mintFutureLandsWithAmount(address recipient, uint256 maxAmount) returns()
func (_Erc721 *Erc721Transactor) MintFutureLandsWithAmount(opts *bind.TransactOpts, recipient common.Address, maxAmount *big.Int) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "mintFutureLandsWithAmount", recipient, maxAmount)
}

// MintFutureLandsWithAmount is a paid mutator transaction binding the contract method 0x71700b56.
//
// Solidity: function mintFutureLandsWithAmount(address recipient, uint256 maxAmount) returns()
func (_Erc721 *Erc721Session) MintFutureLandsWithAmount(recipient common.Address, maxAmount *big.Int) (*types.Transaction, error) {
	return _Erc721.Contract.MintFutureLandsWithAmount(&_Erc721.TransactOpts, recipient, maxAmount)
}

// MintFutureLandsWithAmount is a paid mutator transaction binding the contract method 0x71700b56.
//
// Solidity: function mintFutureLandsWithAmount(address recipient, uint256 maxAmount) returns()
func (_Erc721 *Erc721TransactorSession) MintFutureLandsWithAmount(recipient common.Address, maxAmount *big.Int) (*types.Transaction, error) {
	return _Erc721.Contract.MintFutureLandsWithAmount(&_Erc721.TransactOpts, recipient, maxAmount)
}

// MintLands is a paid mutator transaction binding the contract method 0x3fa8e1b5.
//
// Solidity: function mintLands(uint256 numLands, bytes32[] merkleProof) returns()
func (_Erc721 *Erc721Transactor) MintLands(opts *bind.TransactOpts, numLands *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "mintLands", numLands, merkleProof)
}

// MintLands is a paid mutator transaction binding the contract method 0x3fa8e1b5.
//
// Solidity: function mintLands(uint256 numLands, bytes32[] merkleProof) returns()
func (_Erc721 *Erc721Session) MintLands(numLands *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _Erc721.Contract.MintLands(&_Erc721.TransactOpts, numLands, merkleProof)
}

// MintLands is a paid mutator transaction binding the contract method 0x3fa8e1b5.
//
// Solidity: function mintLands(uint256 numLands, bytes32[] merkleProof) returns()
func (_Erc721 *Erc721TransactorSession) MintLands(numLands *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _Erc721.Contract.MintLands(&_Erc721.TransactOpts, numLands, merkleProof)
}

// NftOwnerClaimLand is a paid mutator transaction binding the contract method 0x7ca0a252.
//
// Solidity: function nftOwnerClaimLand(uint256[] alphaTokenIds, uint256[] betaTokenIds) returns()
func (_Erc721 *Erc721Transactor) NftOwnerClaimLand(opts *bind.TransactOpts, alphaTokenIds []*big.Int, betaTokenIds []*big.Int) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "nftOwnerClaimLand", alphaTokenIds, betaTokenIds)
}

// NftOwnerClaimLand is a paid mutator transaction binding the contract method 0x7ca0a252.
//
// Solidity: function nftOwnerClaimLand(uint256[] alphaTokenIds, uint256[] betaTokenIds) returns()
func (_Erc721 *Erc721Session) NftOwnerClaimLand(alphaTokenIds []*big.Int, betaTokenIds []*big.Int) (*types.Transaction, error) {
	return _Erc721.Contract.NftOwnerClaimLand(&_Erc721.TransactOpts, alphaTokenIds, betaTokenIds)
}

// NftOwnerClaimLand is a paid mutator transaction binding the contract method 0x7ca0a252.
//
// Solidity: function nftOwnerClaimLand(uint256[] alphaTokenIds, uint256[] betaTokenIds) returns()
func (_Erc721 *Erc721TransactorSession) NftOwnerClaimLand(alphaTokenIds []*big.Int, betaTokenIds []*big.Int) (*types.Transaction, error) {
	return _Erc721.Contract.NftOwnerClaimLand(&_Erc721.TransactOpts, alphaTokenIds, betaTokenIds)
}

// PutLandMetadataAtIndex is a paid mutator transaction binding the contract method 0x9c59b66d.
//
// Solidity: function putLandMetadataAtIndex(uint256 index, (bytes32,bytes32,uint256,uint256) _landMetadata) returns()
func (_Erc721 *Erc721Transactor) PutLandMetadataAtIndex(opts *bind.TransactOpts, index *big.Int, _landMetadata LandMetadata) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "putLandMetadataAtIndex", index, _landMetadata)
}

// PutLandMetadataAtIndex is a paid mutator transaction binding the contract method 0x9c59b66d.
//
// Solidity: function putLandMetadataAtIndex(uint256 index, (bytes32,bytes32,uint256,uint256) _landMetadata) returns()
func (_Erc721 *Erc721Session) PutLandMetadataAtIndex(index *big.Int, _landMetadata LandMetadata) (*types.Transaction, error) {
	return _Erc721.Contract.PutLandMetadataAtIndex(&_Erc721.TransactOpts, index, _landMetadata)
}

// PutLandMetadataAtIndex is a paid mutator transaction binding the contract method 0x9c59b66d.
//
// Solidity: function putLandMetadataAtIndex(uint256 index, (bytes32,bytes32,uint256,uint256) _landMetadata) returns()
func (_Erc721 *Erc721TransactorSession) PutLandMetadataAtIndex(index *big.Int, _landMetadata LandMetadata) (*types.Transaction, error) {
	return _Erc721.Contract.PutLandMetadataAtIndex(&_Erc721.TransactOpts, index, _landMetadata)
}

// RawFulfillRandomness is a paid mutator transaction binding the contract method 0x94985ddd.
//
// Solidity: function rawFulfillRandomness(bytes32 requestId, uint256 randomness) returns()
func (_Erc721 *Erc721Transactor) RawFulfillRandomness(opts *bind.TransactOpts, requestId [32]byte, randomness *big.Int) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "rawFulfillRandomness", requestId, randomness)
}

// RawFulfillRandomness is a paid mutator transaction binding the contract method 0x94985ddd.
//
// Solidity: function rawFulfillRandomness(bytes32 requestId, uint256 randomness) returns()
func (_Erc721 *Erc721Session) RawFulfillRandomness(requestId [32]byte, randomness *big.Int) (*types.Transaction, error) {
	return _Erc721.Contract.RawFulfillRandomness(&_Erc721.TransactOpts, requestId, randomness)
}

// RawFulfillRandomness is a paid mutator transaction binding the contract method 0x94985ddd.
//
// Solidity: function rawFulfillRandomness(bytes32 requestId, uint256 randomness) returns()
func (_Erc721 *Erc721TransactorSession) RawFulfillRandomness(requestId [32]byte, randomness *big.Int) (*types.Transaction, error) {
	return _Erc721.Contract.RawFulfillRandomness(&_Erc721.TransactOpts, requestId, randomness)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Erc721 *Erc721Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Erc721 *Erc721Session) RenounceOwnership() (*types.Transaction, error) {
	return _Erc721.Contract.RenounceOwnership(&_Erc721.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Erc721 *Erc721TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Erc721.Contract.RenounceOwnership(&_Erc721.TransactOpts)
}

// RequestRandomnessForOwnerClaim is a paid mutator transaction binding the contract method 0x6f977fbe.
//
// Solidity: function requestRandomnessForOwnerClaim() returns(bytes32 requestId)
func (_Erc721 *Erc721Transactor) RequestRandomnessForOwnerClaim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "requestRandomnessForOwnerClaim")
}

// RequestRandomnessForOwnerClaim is a paid mutator transaction binding the contract method 0x6f977fbe.
//
// Solidity: function requestRandomnessForOwnerClaim() returns(bytes32 requestId)
func (_Erc721 *Erc721Session) RequestRandomnessForOwnerClaim() (*types.Transaction, error) {
	return _Erc721.Contract.RequestRandomnessForOwnerClaim(&_Erc721.TransactOpts)
}

// RequestRandomnessForOwnerClaim is a paid mutator transaction binding the contract method 0x6f977fbe.
//
// Solidity: function requestRandomnessForOwnerClaim() returns(bytes32 requestId)
func (_Erc721 *Erc721TransactorSession) RequestRandomnessForOwnerClaim() (*types.Transaction, error) {
	return _Erc721.Contract.RequestRandomnessForOwnerClaim(&_Erc721.TransactOpts)
}

// RequestRandomnessForPublicSaleAndContributors is a paid mutator transaction binding the contract method 0x653220bc.
//
// Solidity: function requestRandomnessForPublicSaleAndContributors() returns(bytes32 requestId)
func (_Erc721 *Erc721Transactor) RequestRandomnessForPublicSaleAndContributors(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "requestRandomnessForPublicSaleAndContributors")
}

// RequestRandomnessForPublicSaleAndContributors is a paid mutator transaction binding the contract method 0x653220bc.
//
// Solidity: function requestRandomnessForPublicSaleAndContributors() returns(bytes32 requestId)
func (_Erc721 *Erc721Session) RequestRandomnessForPublicSaleAndContributors() (*types.Transaction, error) {
	return _Erc721.Contract.RequestRandomnessForPublicSaleAndContributors(&_Erc721.TransactOpts)
}

// RequestRandomnessForPublicSaleAndContributors is a paid mutator transaction binding the contract method 0x653220bc.
//
// Solidity: function requestRandomnessForPublicSaleAndContributors() returns(bytes32 requestId)
func (_Erc721 *Erc721TransactorSession) RequestRandomnessForPublicSaleAndContributors() (*types.Transaction, error) {
	return _Erc721.Contract.RequestRandomnessForPublicSaleAndContributors(&_Erc721.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721 *Erc721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721 *Erc721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721.Contract.SafeTransferFrom(&_Erc721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721 *Erc721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721.Contract.SafeTransferFrom(&_Erc721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Erc721 *Erc721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Erc721 *Erc721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Erc721.Contract.SafeTransferFrom0(&_Erc721.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Erc721 *Erc721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Erc721.Contract.SafeTransferFrom0(&_Erc721.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc721 *Erc721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc721 *Erc721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc721.Contract.SetApprovalForAll(&_Erc721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc721 *Erc721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc721.Contract.SetApprovalForAll(&_Erc721.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string uri) returns()
func (_Erc721 *Erc721Transactor) SetBaseURI(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "setBaseURI", uri)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string uri) returns()
func (_Erc721 *Erc721Session) SetBaseURI(uri string) (*types.Transaction, error) {
	return _Erc721.Contract.SetBaseURI(&_Erc721.TransactOpts, uri)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string uri) returns()
func (_Erc721 *Erc721TransactorSession) SetBaseURI(uri string) (*types.Transaction, error) {
	return _Erc721.Contract.SetBaseURI(&_Erc721.TransactOpts, uri)
}

// SetFutureMinter is a paid mutator transaction binding the contract method 0xd926f8fa.
//
// Solidity: function setFutureMinter(address _futureMinter) returns()
func (_Erc721 *Erc721Transactor) SetFutureMinter(opts *bind.TransactOpts, _futureMinter common.Address) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "setFutureMinter", _futureMinter)
}

// SetFutureMinter is a paid mutator transaction binding the contract method 0xd926f8fa.
//
// Solidity: function setFutureMinter(address _futureMinter) returns()
func (_Erc721 *Erc721Session) SetFutureMinter(_futureMinter common.Address) (*types.Transaction, error) {
	return _Erc721.Contract.SetFutureMinter(&_Erc721.TransactOpts, _futureMinter)
}

// SetFutureMinter is a paid mutator transaction binding the contract method 0xd926f8fa.
//
// Solidity: function setFutureMinter(address _futureMinter) returns()
func (_Erc721 *Erc721TransactorSession) SetFutureMinter(_futureMinter common.Address) (*types.Transaction, error) {
	return _Erc721.Contract.SetFutureMinter(&_Erc721.TransactOpts, _futureMinter)
}

// SetKycCheckRequired is a paid mutator transaction binding the contract method 0xff1d4080.
//
// Solidity: function setKycCheckRequired(bool _isKycCheckRequired) returns()
func (_Erc721 *Erc721Transactor) SetKycCheckRequired(opts *bind.TransactOpts, _isKycCheckRequired bool) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "setKycCheckRequired", _isKycCheckRequired)
}

// SetKycCheckRequired is a paid mutator transaction binding the contract method 0xff1d4080.
//
// Solidity: function setKycCheckRequired(bool _isKycCheckRequired) returns()
func (_Erc721 *Erc721Session) SetKycCheckRequired(_isKycCheckRequired bool) (*types.Transaction, error) {
	return _Erc721.Contract.SetKycCheckRequired(&_Erc721.TransactOpts, _isKycCheckRequired)
}

// SetKycCheckRequired is a paid mutator transaction binding the contract method 0xff1d4080.
//
// Solidity: function setKycCheckRequired(bool _isKycCheckRequired) returns()
func (_Erc721 *Erc721TransactorSession) SetKycCheckRequired(_isKycCheckRequired bool) (*types.Transaction, error) {
	return _Erc721.Contract.SetKycCheckRequired(&_Erc721.TransactOpts, _isKycCheckRequired)
}

// SetKycMerkleRoot is a paid mutator transaction binding the contract method 0x786867b5.
//
// Solidity: function setKycMerkleRoot(bytes32 _kycMerkleRoot) returns()
func (_Erc721 *Erc721Transactor) SetKycMerkleRoot(opts *bind.TransactOpts, _kycMerkleRoot [32]byte) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "setKycMerkleRoot", _kycMerkleRoot)
}

// SetKycMerkleRoot is a paid mutator transaction binding the contract method 0x786867b5.
//
// Solidity: function setKycMerkleRoot(bytes32 _kycMerkleRoot) returns()
func (_Erc721 *Erc721Session) SetKycMerkleRoot(_kycMerkleRoot [32]byte) (*types.Transaction, error) {
	return _Erc721.Contract.SetKycMerkleRoot(&_Erc721.TransactOpts, _kycMerkleRoot)
}

// SetKycMerkleRoot is a paid mutator transaction binding the contract method 0x786867b5.
//
// Solidity: function setKycMerkleRoot(bytes32 _kycMerkleRoot) returns()
func (_Erc721 *Erc721TransactorSession) SetKycMerkleRoot(_kycMerkleRoot [32]byte) (*types.Transaction, error) {
	return _Erc721.Contract.SetKycMerkleRoot(&_Erc721.TransactOpts, _kycMerkleRoot)
}

// SetMaxMintPerAddress is a paid mutator transaction binding the contract method 0x1e14d44b.
//
// Solidity: function setMaxMintPerAddress(uint256 _maxMintPerAddress) returns()
func (_Erc721 *Erc721Transactor) SetMaxMintPerAddress(opts *bind.TransactOpts, _maxMintPerAddress *big.Int) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "setMaxMintPerAddress", _maxMintPerAddress)
}

// SetMaxMintPerAddress is a paid mutator transaction binding the contract method 0x1e14d44b.
//
// Solidity: function setMaxMintPerAddress(uint256 _maxMintPerAddress) returns()
func (_Erc721 *Erc721Session) SetMaxMintPerAddress(_maxMintPerAddress *big.Int) (*types.Transaction, error) {
	return _Erc721.Contract.SetMaxMintPerAddress(&_Erc721.TransactOpts, _maxMintPerAddress)
}

// SetMaxMintPerAddress is a paid mutator transaction binding the contract method 0x1e14d44b.
//
// Solidity: function setMaxMintPerAddress(uint256 _maxMintPerAddress) returns()
func (_Erc721 *Erc721TransactorSession) SetMaxMintPerAddress(_maxMintPerAddress *big.Int) (*types.Transaction, error) {
	return _Erc721.Contract.SetMaxMintPerAddress(&_Erc721.TransactOpts, _maxMintPerAddress)
}

// SetMaxMintPerTx is a paid mutator transaction binding the contract method 0x616cdb1e.
//
// Solidity: function setMaxMintPerTx(uint256 _maxMintPerTx) returns()
func (_Erc721 *Erc721Transactor) SetMaxMintPerTx(opts *bind.TransactOpts, _maxMintPerTx *big.Int) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "setMaxMintPerTx", _maxMintPerTx)
}

// SetMaxMintPerTx is a paid mutator transaction binding the contract method 0x616cdb1e.
//
// Solidity: function setMaxMintPerTx(uint256 _maxMintPerTx) returns()
func (_Erc721 *Erc721Session) SetMaxMintPerTx(_maxMintPerTx *big.Int) (*types.Transaction, error) {
	return _Erc721.Contract.SetMaxMintPerTx(&_Erc721.TransactOpts, _maxMintPerTx)
}

// SetMaxMintPerTx is a paid mutator transaction binding the contract method 0x616cdb1e.
//
// Solidity: function setMaxMintPerTx(uint256 _maxMintPerTx) returns()
func (_Erc721 *Erc721TransactorSession) SetMaxMintPerTx(_maxMintPerTx *big.Int) (*types.Transaction, error) {
	return _Erc721.Contract.SetMaxMintPerTx(&_Erc721.TransactOpts, _maxMintPerTx)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_Erc721 *Erc721Transactor) SetOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "setOperator", _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_Erc721 *Erc721Session) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _Erc721.Contract.SetOperator(&_Erc721.TransactOpts, _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_Erc721 *Erc721TransactorSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _Erc721.Contract.SetOperator(&_Erc721.TransactOpts, _operator)
}

// StartContributorsClaimPeriod is a paid mutator transaction binding the contract method 0x7d48ca41.
//
// Solidity: function startContributorsClaimPeriod() returns()
func (_Erc721 *Erc721Transactor) StartContributorsClaimPeriod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "startContributorsClaimPeriod")
}

// StartContributorsClaimPeriod is a paid mutator transaction binding the contract method 0x7d48ca41.
//
// Solidity: function startContributorsClaimPeriod() returns()
func (_Erc721 *Erc721Session) StartContributorsClaimPeriod() (*types.Transaction, error) {
	return _Erc721.Contract.StartContributorsClaimPeriod(&_Erc721.TransactOpts)
}

// StartContributorsClaimPeriod is a paid mutator transaction binding the contract method 0x7d48ca41.
//
// Solidity: function startContributorsClaimPeriod() returns()
func (_Erc721 *Erc721TransactorSession) StartContributorsClaimPeriod() (*types.Transaction, error) {
	return _Erc721.Contract.StartContributorsClaimPeriod(&_Erc721.TransactOpts)
}

// StartPublicSale is a paid mutator transaction binding the contract method 0x848d075e.
//
// Solidity: function startPublicSale(uint256 _publicSalePriceLoweringDuration, uint256 _publicSaleStartPrice, uint256 _publicSaleEndingPrice, uint256 _maxMintPerTx, uint256 _maxMintPerAddress, bool _isKycCheckRequired) returns()
func (_Erc721 *Erc721Transactor) StartPublicSale(opts *bind.TransactOpts, _publicSalePriceLoweringDuration *big.Int, _publicSaleStartPrice *big.Int, _publicSaleEndingPrice *big.Int, _maxMintPerTx *big.Int, _maxMintPerAddress *big.Int, _isKycCheckRequired bool) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "startPublicSale", _publicSalePriceLoweringDuration, _publicSaleStartPrice, _publicSaleEndingPrice, _maxMintPerTx, _maxMintPerAddress, _isKycCheckRequired)
}

// StartPublicSale is a paid mutator transaction binding the contract method 0x848d075e.
//
// Solidity: function startPublicSale(uint256 _publicSalePriceLoweringDuration, uint256 _publicSaleStartPrice, uint256 _publicSaleEndingPrice, uint256 _maxMintPerTx, uint256 _maxMintPerAddress, bool _isKycCheckRequired) returns()
func (_Erc721 *Erc721Session) StartPublicSale(_publicSalePriceLoweringDuration *big.Int, _publicSaleStartPrice *big.Int, _publicSaleEndingPrice *big.Int, _maxMintPerTx *big.Int, _maxMintPerAddress *big.Int, _isKycCheckRequired bool) (*types.Transaction, error) {
	return _Erc721.Contract.StartPublicSale(&_Erc721.TransactOpts, _publicSalePriceLoweringDuration, _publicSaleStartPrice, _publicSaleEndingPrice, _maxMintPerTx, _maxMintPerAddress, _isKycCheckRequired)
}

// StartPublicSale is a paid mutator transaction binding the contract method 0x848d075e.
//
// Solidity: function startPublicSale(uint256 _publicSalePriceLoweringDuration, uint256 _publicSaleStartPrice, uint256 _publicSaleEndingPrice, uint256 _maxMintPerTx, uint256 _maxMintPerAddress, bool _isKycCheckRequired) returns()
func (_Erc721 *Erc721TransactorSession) StartPublicSale(_publicSalePriceLoweringDuration *big.Int, _publicSaleStartPrice *big.Int, _publicSaleEndingPrice *big.Int, _maxMintPerTx *big.Int, _maxMintPerAddress *big.Int, _isKycCheckRequired bool) (*types.Transaction, error) {
	return _Erc721.Contract.StartPublicSale(&_Erc721.TransactOpts, _publicSalePriceLoweringDuration, _publicSaleStartPrice, _publicSaleEndingPrice, _maxMintPerTx, _maxMintPerAddress, _isKycCheckRequired)
}

// StopContributorsClaimPeriod is a paid mutator transaction binding the contract method 0x5cb3a9c0.
//
// Solidity: function stopContributorsClaimPeriod() returns()
func (_Erc721 *Erc721Transactor) StopContributorsClaimPeriod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "stopContributorsClaimPeriod")
}

// StopContributorsClaimPeriod is a paid mutator transaction binding the contract method 0x5cb3a9c0.
//
// Solidity: function stopContributorsClaimPeriod() returns()
func (_Erc721 *Erc721Session) StopContributorsClaimPeriod() (*types.Transaction, error) {
	return _Erc721.Contract.StopContributorsClaimPeriod(&_Erc721.TransactOpts)
}

// StopContributorsClaimPeriod is a paid mutator transaction binding the contract method 0x5cb3a9c0.
//
// Solidity: function stopContributorsClaimPeriod() returns()
func (_Erc721 *Erc721TransactorSession) StopContributorsClaimPeriod() (*types.Transaction, error) {
	return _Erc721.Contract.StopContributorsClaimPeriod(&_Erc721.TransactOpts)
}

// StopPublicSale is a paid mutator transaction binding the contract method 0xda1b91c3.
//
// Solidity: function stopPublicSale() returns()
func (_Erc721 *Erc721Transactor) StopPublicSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "stopPublicSale")
}

// StopPublicSale is a paid mutator transaction binding the contract method 0xda1b91c3.
//
// Solidity: function stopPublicSale() returns()
func (_Erc721 *Erc721Session) StopPublicSale() (*types.Transaction, error) {
	return _Erc721.Contract.StopPublicSale(&_Erc721.TransactOpts)
}

// StopPublicSale is a paid mutator transaction binding the contract method 0xda1b91c3.
//
// Solidity: function stopPublicSale() returns()
func (_Erc721 *Erc721TransactorSession) StopPublicSale() (*types.Transaction, error) {
	return _Erc721.Contract.StopPublicSale(&_Erc721.TransactOpts)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721 *Erc721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721 *Erc721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721.Contract.TransferFrom(&_Erc721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721 *Erc721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721.Contract.TransferFrom(&_Erc721.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc721 *Erc721Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc721 *Erc721Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Erc721.Contract.TransferOwnership(&_Erc721.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc721 *Erc721TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Erc721.Contract.TransferOwnership(&_Erc721.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Erc721 *Erc721Transactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Erc721 *Erc721Session) Withdraw() (*types.Transaction, error) {
	return _Erc721.Contract.Withdraw(&_Erc721.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Erc721 *Erc721TransactorSession) Withdraw() (*types.Transaction, error) {
	return _Erc721.Contract.Withdraw(&_Erc721.TransactOpts)
}

// Erc721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc721 contract.
type Erc721ApprovalIterator struct {
	Event *Erc721Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc721Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721Approval represents a Approval event raised by the Erc721 contract.
type Erc721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Erc721 *Erc721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*Erc721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Erc721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Erc721ApprovalIterator{contract: _Erc721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Erc721 *Erc721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Erc721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721Approval)
				if err := _Erc721.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Erc721 *Erc721Filterer) ParseApproval(log types.Log) (*Erc721Approval, error) {
	event := new(Erc721Approval)
	if err := _Erc721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Erc721 contract.
type Erc721ApprovalForAllIterator struct {
	Event *Erc721ApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721ApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc721ApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721ApprovalForAll represents a ApprovalForAll event raised by the Erc721 contract.
type Erc721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Erc721 *Erc721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*Erc721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Erc721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Erc721ApprovalForAllIterator{contract: _Erc721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Erc721 *Erc721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Erc721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Erc721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721ApprovalForAll)
				if err := _Erc721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Erc721 *Erc721Filterer) ParseApprovalForAll(log types.Log) (*Erc721ApprovalForAll, error) {
	event := new(Erc721ApprovalForAll)
	if err := _Erc721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc721ClaimableStateChangedIterator is returned from FilterClaimableStateChanged and is used to iterate over the raw logs and unpacked data for ClaimableStateChanged events raised by the Erc721 contract.
type Erc721ClaimableStateChangedIterator struct {
	Event *Erc721ClaimableStateChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc721ClaimableStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721ClaimableStateChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc721ClaimableStateChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc721ClaimableStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721ClaimableStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721ClaimableStateChanged represents a ClaimableStateChanged event raised by the Erc721 contract.
type Erc721ClaimableStateChanged struct {
	ClaimableActive bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterClaimableStateChanged is a free log retrieval operation binding the contract event 0x00231f1eb7ad7923209c5cc8028852e71745bff7e8c7d9f8752f2d3a69f29974.
//
// Solidity: event ClaimableStateChanged(bool indexed claimableActive)
func (_Erc721 *Erc721Filterer) FilterClaimableStateChanged(opts *bind.FilterOpts, claimableActive []bool) (*Erc721ClaimableStateChangedIterator, error) {

	var claimableActiveRule []interface{}
	for _, claimableActiveItem := range claimableActive {
		claimableActiveRule = append(claimableActiveRule, claimableActiveItem)
	}

	logs, sub, err := _Erc721.contract.FilterLogs(opts, "ClaimableStateChanged", claimableActiveRule)
	if err != nil {
		return nil, err
	}
	return &Erc721ClaimableStateChangedIterator{contract: _Erc721.contract, event: "ClaimableStateChanged", logs: logs, sub: sub}, nil
}

// WatchClaimableStateChanged is a free log subscription operation binding the contract event 0x00231f1eb7ad7923209c5cc8028852e71745bff7e8c7d9f8752f2d3a69f29974.
//
// Solidity: event ClaimableStateChanged(bool indexed claimableActive)
func (_Erc721 *Erc721Filterer) WatchClaimableStateChanged(opts *bind.WatchOpts, sink chan<- *Erc721ClaimableStateChanged, claimableActive []bool) (event.Subscription, error) {

	var claimableActiveRule []interface{}
	for _, claimableActiveItem := range claimableActive {
		claimableActiveRule = append(claimableActiveRule, claimableActiveItem)
	}

	logs, sub, err := _Erc721.contract.WatchLogs(opts, "ClaimableStateChanged", claimableActiveRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721ClaimableStateChanged)
				if err := _Erc721.contract.UnpackLog(event, "ClaimableStateChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimableStateChanged is a log parse operation binding the contract event 0x00231f1eb7ad7923209c5cc8028852e71745bff7e8c7d9f8752f2d3a69f29974.
//
// Solidity: event ClaimableStateChanged(bool indexed claimableActive)
func (_Erc721 *Erc721Filterer) ParseClaimableStateChanged(log types.Log) (*Erc721ClaimableStateChanged, error) {
	event := new(Erc721ClaimableStateChanged)
	if err := _Erc721.contract.UnpackLog(event, "ClaimableStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc721ContributorsClaimStartIterator is returned from FilterContributorsClaimStart and is used to iterate over the raw logs and unpacked data for ContributorsClaimStart events raised by the Erc721 contract.
type Erc721ContributorsClaimStartIterator struct {
	Event *Erc721ContributorsClaimStart // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc721ContributorsClaimStartIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721ContributorsClaimStart)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc721ContributorsClaimStart)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc721ContributorsClaimStartIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721ContributorsClaimStartIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721ContributorsClaimStart represents a ContributorsClaimStart event raised by the Erc721 contract.
type Erc721ContributorsClaimStart struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterContributorsClaimStart is a free log retrieval operation binding the contract event 0xb821e7c7541dfb5a35afc6d252e3cfcd56e0e25852e9c38fb7504da18ae4209e.
//
// Solidity: event ContributorsClaimStart(uint256 _timestamp)
func (_Erc721 *Erc721Filterer) FilterContributorsClaimStart(opts *bind.FilterOpts) (*Erc721ContributorsClaimStartIterator, error) {

	logs, sub, err := _Erc721.contract.FilterLogs(opts, "ContributorsClaimStart")
	if err != nil {
		return nil, err
	}
	return &Erc721ContributorsClaimStartIterator{contract: _Erc721.contract, event: "ContributorsClaimStart", logs: logs, sub: sub}, nil
}

// WatchContributorsClaimStart is a free log subscription operation binding the contract event 0xb821e7c7541dfb5a35afc6d252e3cfcd56e0e25852e9c38fb7504da18ae4209e.
//
// Solidity: event ContributorsClaimStart(uint256 _timestamp)
func (_Erc721 *Erc721Filterer) WatchContributorsClaimStart(opts *bind.WatchOpts, sink chan<- *Erc721ContributorsClaimStart) (event.Subscription, error) {

	logs, sub, err := _Erc721.contract.WatchLogs(opts, "ContributorsClaimStart")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721ContributorsClaimStart)
				if err := _Erc721.contract.UnpackLog(event, "ContributorsClaimStart", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContributorsClaimStart is a log parse operation binding the contract event 0xb821e7c7541dfb5a35afc6d252e3cfcd56e0e25852e9c38fb7504da18ae4209e.
//
// Solidity: event ContributorsClaimStart(uint256 _timestamp)
func (_Erc721 *Erc721Filterer) ParseContributorsClaimStart(log types.Log) (*Erc721ContributorsClaimStart, error) {
	event := new(Erc721ContributorsClaimStart)
	if err := _Erc721.contract.UnpackLog(event, "ContributorsClaimStart", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc721ContributorsClaimStopIterator is returned from FilterContributorsClaimStop and is used to iterate over the raw logs and unpacked data for ContributorsClaimStop events raised by the Erc721 contract.
type Erc721ContributorsClaimStopIterator struct {
	Event *Erc721ContributorsClaimStop // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc721ContributorsClaimStopIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721ContributorsClaimStop)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc721ContributorsClaimStop)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc721ContributorsClaimStopIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721ContributorsClaimStopIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721ContributorsClaimStop represents a ContributorsClaimStop event raised by the Erc721 contract.
type Erc721ContributorsClaimStop struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterContributorsClaimStop is a free log retrieval operation binding the contract event 0x4018d3084dfacabf0eba098d4b7b8b4140b4eae436210480e5affa48fdfacd76.
//
// Solidity: event ContributorsClaimStop(uint256 _timestamp)
func (_Erc721 *Erc721Filterer) FilterContributorsClaimStop(opts *bind.FilterOpts) (*Erc721ContributorsClaimStopIterator, error) {

	logs, sub, err := _Erc721.contract.FilterLogs(opts, "ContributorsClaimStop")
	if err != nil {
		return nil, err
	}
	return &Erc721ContributorsClaimStopIterator{contract: _Erc721.contract, event: "ContributorsClaimStop", logs: logs, sub: sub}, nil
}

// WatchContributorsClaimStop is a free log subscription operation binding the contract event 0x4018d3084dfacabf0eba098d4b7b8b4140b4eae436210480e5affa48fdfacd76.
//
// Solidity: event ContributorsClaimStop(uint256 _timestamp)
func (_Erc721 *Erc721Filterer) WatchContributorsClaimStop(opts *bind.WatchOpts, sink chan<- *Erc721ContributorsClaimStop) (event.Subscription, error) {

	logs, sub, err := _Erc721.contract.WatchLogs(opts, "ContributorsClaimStop")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721ContributorsClaimStop)
				if err := _Erc721.contract.UnpackLog(event, "ContributorsClaimStop", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContributorsClaimStop is a log parse operation binding the contract event 0x4018d3084dfacabf0eba098d4b7b8b4140b4eae436210480e5affa48fdfacd76.
//
// Solidity: event ContributorsClaimStop(uint256 _timestamp)
func (_Erc721 *Erc721Filterer) ParseContributorsClaimStop(log types.Log) (*Erc721ContributorsClaimStop, error) {
	event := new(Erc721ContributorsClaimStop)
	if err := _Erc721.contract.UnpackLog(event, "ContributorsClaimStop", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc721LandPublicSaleStartIterator is returned from FilterLandPublicSaleStart and is used to iterate over the raw logs and unpacked data for LandPublicSaleStart events raised by the Erc721 contract.
type Erc721LandPublicSaleStartIterator struct {
	Event *Erc721LandPublicSaleStart // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc721LandPublicSaleStartIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721LandPublicSaleStart)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc721LandPublicSaleStart)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc721LandPublicSaleStartIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721LandPublicSaleStartIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721LandPublicSaleStart represents a LandPublicSaleStart event raised by the Erc721 contract.
type Erc721LandPublicSaleStart struct {
	SaleDuration  *big.Int
	SaleStartTime *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLandPublicSaleStart is a free log retrieval operation binding the contract event 0x03bbdfe69cc0e9bf6a00b606f78ef6f3391ea272251e9ab2d56ce08f96be745f.
//
// Solidity: event LandPublicSaleStart(uint256 indexed _saleDuration, uint256 indexed _saleStartTime)
func (_Erc721 *Erc721Filterer) FilterLandPublicSaleStart(opts *bind.FilterOpts, _saleDuration []*big.Int, _saleStartTime []*big.Int) (*Erc721LandPublicSaleStartIterator, error) {

	var _saleDurationRule []interface{}
	for _, _saleDurationItem := range _saleDuration {
		_saleDurationRule = append(_saleDurationRule, _saleDurationItem)
	}
	var _saleStartTimeRule []interface{}
	for _, _saleStartTimeItem := range _saleStartTime {
		_saleStartTimeRule = append(_saleStartTimeRule, _saleStartTimeItem)
	}

	logs, sub, err := _Erc721.contract.FilterLogs(opts, "LandPublicSaleStart", _saleDurationRule, _saleStartTimeRule)
	if err != nil {
		return nil, err
	}
	return &Erc721LandPublicSaleStartIterator{contract: _Erc721.contract, event: "LandPublicSaleStart", logs: logs, sub: sub}, nil
}

// WatchLandPublicSaleStart is a free log subscription operation binding the contract event 0x03bbdfe69cc0e9bf6a00b606f78ef6f3391ea272251e9ab2d56ce08f96be745f.
//
// Solidity: event LandPublicSaleStart(uint256 indexed _saleDuration, uint256 indexed _saleStartTime)
func (_Erc721 *Erc721Filterer) WatchLandPublicSaleStart(opts *bind.WatchOpts, sink chan<- *Erc721LandPublicSaleStart, _saleDuration []*big.Int, _saleStartTime []*big.Int) (event.Subscription, error) {

	var _saleDurationRule []interface{}
	for _, _saleDurationItem := range _saleDuration {
		_saleDurationRule = append(_saleDurationRule, _saleDurationItem)
	}
	var _saleStartTimeRule []interface{}
	for _, _saleStartTimeItem := range _saleStartTime {
		_saleStartTimeRule = append(_saleStartTimeRule, _saleStartTimeItem)
	}

	logs, sub, err := _Erc721.contract.WatchLogs(opts, "LandPublicSaleStart", _saleDurationRule, _saleStartTimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721LandPublicSaleStart)
				if err := _Erc721.contract.UnpackLog(event, "LandPublicSaleStart", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLandPublicSaleStart is a log parse operation binding the contract event 0x03bbdfe69cc0e9bf6a00b606f78ef6f3391ea272251e9ab2d56ce08f96be745f.
//
// Solidity: event LandPublicSaleStart(uint256 indexed _saleDuration, uint256 indexed _saleStartTime)
func (_Erc721 *Erc721Filterer) ParseLandPublicSaleStart(log types.Log) (*Erc721LandPublicSaleStart, error) {
	event := new(Erc721LandPublicSaleStart)
	if err := _Erc721.contract.UnpackLog(event, "LandPublicSaleStart", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc721LandPublicSaleStopIterator is returned from FilterLandPublicSaleStop and is used to iterate over the raw logs and unpacked data for LandPublicSaleStop events raised by the Erc721 contract.
type Erc721LandPublicSaleStopIterator struct {
	Event *Erc721LandPublicSaleStop // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc721LandPublicSaleStopIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721LandPublicSaleStop)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc721LandPublicSaleStop)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc721LandPublicSaleStopIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721LandPublicSaleStopIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721LandPublicSaleStop represents a LandPublicSaleStop event raised by the Erc721 contract.
type Erc721LandPublicSaleStop struct {
	CurrentPrice *big.Int
	TimeElapsed  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLandPublicSaleStop is a free log retrieval operation binding the contract event 0x3da9555b37cd6c211f437cd26ac71eb0716e111fa9458c73183e99711e4e34eb.
//
// Solidity: event LandPublicSaleStop(uint256 indexed _currentPrice, uint256 indexed _timeElapsed)
func (_Erc721 *Erc721Filterer) FilterLandPublicSaleStop(opts *bind.FilterOpts, _currentPrice []*big.Int, _timeElapsed []*big.Int) (*Erc721LandPublicSaleStopIterator, error) {

	var _currentPriceRule []interface{}
	for _, _currentPriceItem := range _currentPrice {
		_currentPriceRule = append(_currentPriceRule, _currentPriceItem)
	}
	var _timeElapsedRule []interface{}
	for _, _timeElapsedItem := range _timeElapsed {
		_timeElapsedRule = append(_timeElapsedRule, _timeElapsedItem)
	}

	logs, sub, err := _Erc721.contract.FilterLogs(opts, "LandPublicSaleStop", _currentPriceRule, _timeElapsedRule)
	if err != nil {
		return nil, err
	}
	return &Erc721LandPublicSaleStopIterator{contract: _Erc721.contract, event: "LandPublicSaleStop", logs: logs, sub: sub}, nil
}

// WatchLandPublicSaleStop is a free log subscription operation binding the contract event 0x3da9555b37cd6c211f437cd26ac71eb0716e111fa9458c73183e99711e4e34eb.
//
// Solidity: event LandPublicSaleStop(uint256 indexed _currentPrice, uint256 indexed _timeElapsed)
func (_Erc721 *Erc721Filterer) WatchLandPublicSaleStop(opts *bind.WatchOpts, sink chan<- *Erc721LandPublicSaleStop, _currentPrice []*big.Int, _timeElapsed []*big.Int) (event.Subscription, error) {

	var _currentPriceRule []interface{}
	for _, _currentPriceItem := range _currentPrice {
		_currentPriceRule = append(_currentPriceRule, _currentPriceItem)
	}
	var _timeElapsedRule []interface{}
	for _, _timeElapsedItem := range _timeElapsed {
		_timeElapsedRule = append(_timeElapsedRule, _timeElapsedItem)
	}

	logs, sub, err := _Erc721.contract.WatchLogs(opts, "LandPublicSaleStop", _currentPriceRule, _timeElapsedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721LandPublicSaleStop)
				if err := _Erc721.contract.UnpackLog(event, "LandPublicSaleStop", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLandPublicSaleStop is a log parse operation binding the contract event 0x3da9555b37cd6c211f437cd26ac71eb0716e111fa9458c73183e99711e4e34eb.
//
// Solidity: event LandPublicSaleStop(uint256 indexed _currentPrice, uint256 indexed _timeElapsed)
func (_Erc721 *Erc721Filterer) ParseLandPublicSaleStop(log types.Log) (*Erc721LandPublicSaleStop, error) {
	event := new(Erc721LandPublicSaleStop)
	if err := _Erc721.contract.UnpackLog(event, "LandPublicSaleStop", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc721OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Erc721 contract.
type Erc721OwnershipTransferredIterator struct {
	Event *Erc721OwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc721OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc721OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc721OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721OwnershipTransferred represents a OwnershipTransferred event raised by the Erc721 contract.
type Erc721OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc721 *Erc721Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Erc721OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Erc721.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Erc721OwnershipTransferredIterator{contract: _Erc721.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc721 *Erc721Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Erc721OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Erc721.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721OwnershipTransferred)
				if err := _Erc721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc721 *Erc721Filterer) ParseOwnershipTransferred(log types.Log) (*Erc721OwnershipTransferred, error) {
	event := new(Erc721OwnershipTransferred)
	if err := _Erc721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc721PublicSaleMintIterator is returned from FilterPublicSaleMint and is used to iterate over the raw logs and unpacked data for PublicSaleMint events raised by the Erc721 contract.
type Erc721PublicSaleMintIterator struct {
	Event *Erc721PublicSaleMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc721PublicSaleMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721PublicSaleMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc721PublicSaleMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc721PublicSaleMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721PublicSaleMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721PublicSaleMint represents a PublicSaleMint event raised by the Erc721 contract.
type Erc721PublicSaleMint struct {
	Sender    common.Address
	NumLands  *big.Int
	MintPrice *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPublicSaleMint is a free log retrieval operation binding the contract event 0x2c7d174a64b49c17bcea3a44c1ba1547c9a3f4997b68952c5dd3fcc1f17f7d6d.
//
// Solidity: event PublicSaleMint(address indexed sender, uint256 indexed numLands, uint256 indexed mintPrice)
func (_Erc721 *Erc721Filterer) FilterPublicSaleMint(opts *bind.FilterOpts, sender []common.Address, numLands []*big.Int, mintPrice []*big.Int) (*Erc721PublicSaleMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var numLandsRule []interface{}
	for _, numLandsItem := range numLands {
		numLandsRule = append(numLandsRule, numLandsItem)
	}
	var mintPriceRule []interface{}
	for _, mintPriceItem := range mintPrice {
		mintPriceRule = append(mintPriceRule, mintPriceItem)
	}

	logs, sub, err := _Erc721.contract.FilterLogs(opts, "PublicSaleMint", senderRule, numLandsRule, mintPriceRule)
	if err != nil {
		return nil, err
	}
	return &Erc721PublicSaleMintIterator{contract: _Erc721.contract, event: "PublicSaleMint", logs: logs, sub: sub}, nil
}

// WatchPublicSaleMint is a free log subscription operation binding the contract event 0x2c7d174a64b49c17bcea3a44c1ba1547c9a3f4997b68952c5dd3fcc1f17f7d6d.
//
// Solidity: event PublicSaleMint(address indexed sender, uint256 indexed numLands, uint256 indexed mintPrice)
func (_Erc721 *Erc721Filterer) WatchPublicSaleMint(opts *bind.WatchOpts, sink chan<- *Erc721PublicSaleMint, sender []common.Address, numLands []*big.Int, mintPrice []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var numLandsRule []interface{}
	for _, numLandsItem := range numLands {
		numLandsRule = append(numLandsRule, numLandsItem)
	}
	var mintPriceRule []interface{}
	for _, mintPriceItem := range mintPrice {
		mintPriceRule = append(mintPriceRule, mintPriceItem)
	}

	logs, sub, err := _Erc721.contract.WatchLogs(opts, "PublicSaleMint", senderRule, numLandsRule, mintPriceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721PublicSaleMint)
				if err := _Erc721.contract.UnpackLog(event, "PublicSaleMint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePublicSaleMint is a log parse operation binding the contract event 0x2c7d174a64b49c17bcea3a44c1ba1547c9a3f4997b68952c5dd3fcc1f17f7d6d.
//
// Solidity: event PublicSaleMint(address indexed sender, uint256 indexed numLands, uint256 indexed mintPrice)
func (_Erc721 *Erc721Filterer) ParsePublicSaleMint(log types.Log) (*Erc721PublicSaleMint, error) {
	event := new(Erc721PublicSaleMint)
	if err := _Erc721.contract.UnpackLog(event, "PublicSaleMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc721StartingIndexSetAlphaBetaIterator is returned from FilterStartingIndexSetAlphaBeta and is used to iterate over the raw logs and unpacked data for StartingIndexSetAlphaBeta events raised by the Erc721 contract.
type Erc721StartingIndexSetAlphaBetaIterator struct {
	Event *Erc721StartingIndexSetAlphaBeta // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc721StartingIndexSetAlphaBetaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721StartingIndexSetAlphaBeta)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc721StartingIndexSetAlphaBeta)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc721StartingIndexSetAlphaBetaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721StartingIndexSetAlphaBetaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721StartingIndexSetAlphaBeta represents a StartingIndexSetAlphaBeta event raised by the Erc721 contract.
type Erc721StartingIndexSetAlphaBeta struct {
	AlphaOffset *big.Int
	BetaOffset  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStartingIndexSetAlphaBeta is a free log retrieval operation binding the contract event 0x7ed9998d8bac64249deff15104738a8f3446fef3c2d1155d10baa100eeb4965a.
//
// Solidity: event StartingIndexSetAlphaBeta(uint256 indexed _alphaOffset, uint256 indexed _betaOffset)
func (_Erc721 *Erc721Filterer) FilterStartingIndexSetAlphaBeta(opts *bind.FilterOpts, _alphaOffset []*big.Int, _betaOffset []*big.Int) (*Erc721StartingIndexSetAlphaBetaIterator, error) {

	var _alphaOffsetRule []interface{}
	for _, _alphaOffsetItem := range _alphaOffset {
		_alphaOffsetRule = append(_alphaOffsetRule, _alphaOffsetItem)
	}
	var _betaOffsetRule []interface{}
	for _, _betaOffsetItem := range _betaOffset {
		_betaOffsetRule = append(_betaOffsetRule, _betaOffsetItem)
	}

	logs, sub, err := _Erc721.contract.FilterLogs(opts, "StartingIndexSetAlphaBeta", _alphaOffsetRule, _betaOffsetRule)
	if err != nil {
		return nil, err
	}
	return &Erc721StartingIndexSetAlphaBetaIterator{contract: _Erc721.contract, event: "StartingIndexSetAlphaBeta", logs: logs, sub: sub}, nil
}

// WatchStartingIndexSetAlphaBeta is a free log subscription operation binding the contract event 0x7ed9998d8bac64249deff15104738a8f3446fef3c2d1155d10baa100eeb4965a.
//
// Solidity: event StartingIndexSetAlphaBeta(uint256 indexed _alphaOffset, uint256 indexed _betaOffset)
func (_Erc721 *Erc721Filterer) WatchStartingIndexSetAlphaBeta(opts *bind.WatchOpts, sink chan<- *Erc721StartingIndexSetAlphaBeta, _alphaOffset []*big.Int, _betaOffset []*big.Int) (event.Subscription, error) {

	var _alphaOffsetRule []interface{}
	for _, _alphaOffsetItem := range _alphaOffset {
		_alphaOffsetRule = append(_alphaOffsetRule, _alphaOffsetItem)
	}
	var _betaOffsetRule []interface{}
	for _, _betaOffsetItem := range _betaOffset {
		_betaOffsetRule = append(_betaOffsetRule, _betaOffsetItem)
	}

	logs, sub, err := _Erc721.contract.WatchLogs(opts, "StartingIndexSetAlphaBeta", _alphaOffsetRule, _betaOffsetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721StartingIndexSetAlphaBeta)
				if err := _Erc721.contract.UnpackLog(event, "StartingIndexSetAlphaBeta", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStartingIndexSetAlphaBeta is a log parse operation binding the contract event 0x7ed9998d8bac64249deff15104738a8f3446fef3c2d1155d10baa100eeb4965a.
//
// Solidity: event StartingIndexSetAlphaBeta(uint256 indexed _alphaOffset, uint256 indexed _betaOffset)
func (_Erc721 *Erc721Filterer) ParseStartingIndexSetAlphaBeta(log types.Log) (*Erc721StartingIndexSetAlphaBeta, error) {
	event := new(Erc721StartingIndexSetAlphaBeta)
	if err := _Erc721.contract.UnpackLog(event, "StartingIndexSetAlphaBeta", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc721StartingIndexSetPublicSaleIterator is returned from FilterStartingIndexSetPublicSale and is used to iterate over the raw logs and unpacked data for StartingIndexSetPublicSale events raised by the Erc721 contract.
type Erc721StartingIndexSetPublicSaleIterator struct {
	Event *Erc721StartingIndexSetPublicSale // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc721StartingIndexSetPublicSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721StartingIndexSetPublicSale)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc721StartingIndexSetPublicSale)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc721StartingIndexSetPublicSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721StartingIndexSetPublicSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721StartingIndexSetPublicSale represents a StartingIndexSetPublicSale event raised by the Erc721 contract.
type Erc721StartingIndexSetPublicSale struct {
	StartingIndex *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterStartingIndexSetPublicSale is a free log retrieval operation binding the contract event 0x662707e4febdcde4fd5eca7d6311dc840e55b942dc58734aac52fff6d866da99.
//
// Solidity: event StartingIndexSetPublicSale(uint256 indexed _startingIndex)
func (_Erc721 *Erc721Filterer) FilterStartingIndexSetPublicSale(opts *bind.FilterOpts, _startingIndex []*big.Int) (*Erc721StartingIndexSetPublicSaleIterator, error) {

	var _startingIndexRule []interface{}
	for _, _startingIndexItem := range _startingIndex {
		_startingIndexRule = append(_startingIndexRule, _startingIndexItem)
	}

	logs, sub, err := _Erc721.contract.FilterLogs(opts, "StartingIndexSetPublicSale", _startingIndexRule)
	if err != nil {
		return nil, err
	}
	return &Erc721StartingIndexSetPublicSaleIterator{contract: _Erc721.contract, event: "StartingIndexSetPublicSale", logs: logs, sub: sub}, nil
}

// WatchStartingIndexSetPublicSale is a free log subscription operation binding the contract event 0x662707e4febdcde4fd5eca7d6311dc840e55b942dc58734aac52fff6d866da99.
//
// Solidity: event StartingIndexSetPublicSale(uint256 indexed _startingIndex)
func (_Erc721 *Erc721Filterer) WatchStartingIndexSetPublicSale(opts *bind.WatchOpts, sink chan<- *Erc721StartingIndexSetPublicSale, _startingIndex []*big.Int) (event.Subscription, error) {

	var _startingIndexRule []interface{}
	for _, _startingIndexItem := range _startingIndex {
		_startingIndexRule = append(_startingIndexRule, _startingIndexItem)
	}

	logs, sub, err := _Erc721.contract.WatchLogs(opts, "StartingIndexSetPublicSale", _startingIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721StartingIndexSetPublicSale)
				if err := _Erc721.contract.UnpackLog(event, "StartingIndexSetPublicSale", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStartingIndexSetPublicSale is a log parse operation binding the contract event 0x662707e4febdcde4fd5eca7d6311dc840e55b942dc58734aac52fff6d866da99.
//
// Solidity: event StartingIndexSetPublicSale(uint256 indexed _startingIndex)
func (_Erc721 *Erc721Filterer) ParseStartingIndexSetPublicSale(log types.Log) (*Erc721StartingIndexSetPublicSale, error) {
	event := new(Erc721StartingIndexSetPublicSale)
	if err := _Erc721.contract.UnpackLog(event, "StartingIndexSetPublicSale", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc721 contract.
type Erc721TransferIterator struct {
	Event *Erc721Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc721Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721Transfer represents a Transfer event raised by the Erc721 contract.
type Erc721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Erc721 *Erc721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*Erc721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Erc721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Erc721TransferIterator{contract: _Erc721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Erc721 *Erc721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Erc721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721Transfer)
				if err := _Erc721.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Erc721 *Erc721Filterer) ParseTransfer(log types.Log) (*Erc721Transfer, error) {
	event := new(Erc721Transfer)
	if err := _Erc721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
