package metadata

import (
	"errors"

	"github.com/ninjadotorg/constant/common"
	privacy "github.com/ninjadotorg/constant/privacy-protocol"
)

type BuySellRequest struct {
	PaymentAddress privacy.PaymentAddress
	AssetType      common.Hash // token id (note: for bond, this one is just bond token id prefix)
	Amount         uint64
	BuyPrice       uint64 // in Constant unit

	SaleID []byte // only when requesting to DCB

	MetadataBase
}

func NewBuySellRequest(bsReqData map[string]interface{}) *BuySellRequest {
	return &BuySellRequest{
		PaymentAddress: bsReqData["paymentAddress"].(privacy.PaymentAddress),
		AssetType:      bsReqData["assetType"].(common.Hash),
		Amount:         uint64(bsReqData["amount"].(float64)),
		BuyPrice:       uint64(bsReqData["buyPrice"].(float64)),
		SaleID:         bsReqData["saleId"].([]byte),
	}
}

func (bsReq *BuySellRequest) ValidateTxWithBlockChain(txr Transaction, bcr BlockchainRetriever, chainID byte) (bool, error) {

	// check double spending on fee + buy/sell amount tx
	// err := self.ValidateDoubleSpend(&buySellReqTx.Tx, chainID)
	err := txr.ValidateConstDoubleSpendWithBlockchain(bcr, chainID)
	if err != nil {
		return false, err
	}

	// TODO: support and validate for either bonds or govs buy requests

	govParams := bcr.GetGOVParams()
	sellingBondsParams := govParams.SellingBonds
	if sellingBondsParams == nil {
		return false, errors.New("SellingBonds params are not existed.")
	}

	// check if buy price againsts SellingBonds params' BondPrice is correct or not
	if bsReq.BuyPrice < sellingBondsParams.BondPrice {
		return false, errors.New("Requested buy price is under SellingBonds params' buy price.")
	}
	return false, nil
}

func (bsReq *BuySellRequest) ValidateSanityData(txr Transaction) (bool, bool, error) {
	ok, err := txr.ValidateSanityData()
	if err != nil || !ok {
		return false, ok, err
	}
	if len(bsReq.PaymentAddress.Pk) == 0 {
		return false, false, errors.New("Wrong request info's payment address")
	}
	if bsReq.BuyPrice == 0 {
		return false, false, errors.New("Wrong request info's buy price")
	}
	if bsReq.Amount == 0 {
		return false, false, errors.New("Wrong request info's amount")
	}
	if len(bsReq.AssetType) != common.HashSize {
		return false, false, errors.New("Wrong request info's asset type")
	}
	return false, true, nil
}

func (bsReq *BuySellRequest) ValidateMetadataByItself() bool {
	// The validation just need to check at tx level, so returning true here
	return true
}

// func (bsReq *BuySellRequest) GetType() int {
// 	return BuySellRequestMeta
// }

func (bsReq *BuySellRequest) Hash() *common.Hash {
	record := string(bsReq.PaymentAddress.ToBytes())
	record += bsReq.AssetType.String()
	record += string(bsReq.Amount)
	record += string(bsReq.BuyPrice)
	record += string(bsReq.SaleID)

	// final hash
	hash := common.DoubleHashH([]byte(record))
	return &hash
}
