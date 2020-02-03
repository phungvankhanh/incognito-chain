package block

import (
	"strconv"

	"github.com/incognitochain/incognito-chain/common"
	"github.com/incognitochain/incognito-chain/database"
	"github.com/incognitochain/incognito-chain/metadata"
	"github.com/pkg/errors"
)

func storeBurningConfirm(block *ShardBlock, bd *[]database.BatchData, blockchain BlockChain, logger common.Logger) error {
	for _, inst := range block.Body.Instructions {
		if inst[0] != strconv.Itoa(metadata.BurningConfirmMeta) {
			continue
		}
		logger.Infof("storeBurningConfirm for block %d, inst %v", block.Header.Height, inst)

		txID, err := common.Hash{}.NewHashFromStr(inst[5])
		if err != nil {
			return errors.Wrap(err, "txid invalid")
		}
		if err := blockchain.GetDatabase().StoreBurningConfirm(*txID, block.Header.Height, bd); err != nil {
			return errors.Wrapf(err, "store failed, txID: %x", txID)
		}
	}
	return nil
}
