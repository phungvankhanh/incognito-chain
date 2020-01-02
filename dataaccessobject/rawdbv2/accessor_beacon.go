package rawdbv2

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/incognitochain/incognito-chain/common"
	"github.com/incognitochain/incognito-chain/incdb"
)

// StoreBeaconBlock store block hash => block value and block index => block hash
// record1: prefix-index-hash => empty
// record2: prefix-hash => block value
func StoreBeaconBlock(db incdb.Database, index uint64, hash common.Hash, v interface{}) error {
	keyHash := GetBeaconHashToBlockKey(hash)
	if ok, _ := db.Has(keyHash); ok {
		return NewRawdbError(StoreBeaconBlockError, fmt.Errorf("key %+v already exists", keyHash))
	}
	keyIndex := GetBeaconIndexToBlockHashKey(index, hash)
	if ok, _ := db.Has(keyIndex); ok {
		return NewRawdbError(StoreBeaconBlockError, fmt.Errorf("key %+v already exists", keyIndex))
	}
	val, err := json.Marshal(v)
	if err != nil {
		return NewRawdbError(StoreBeaconBlockError, err)
	}
	if err := db.Put(keyIndex, []byte{}); err != nil {
		return NewRawdbError(StoreBeaconBlockError, err)
	}
	if err := db.Put(keyHash, val); err != nil {
		return NewRawdbError(StoreBeaconBlockError, err)
	}
	return nil
}

// StoreBeaconBlockIndex store block hash => block index
// key: i-{hash}
// value: {index-shardID}
func StoreBeaconBlockIndex(db incdb.Database, hash common.Hash, index uint64) error {
	key := GetBeaconBlockHashToIndexKey(hash)
	buf := common.Uint64ToBytes(index)
	err := db.Put(key, buf)
	if err != nil {
		return NewRawdbError(StoreBeaconBlockIndexError, err)
	}
	return nil
}

func HasBeaconBlock(db incdb.Database, hash common.Hash) (bool, error) {
	keyHash := GetBeaconHashToBlockKey(hash)
	if ok, err := db.Has(keyHash); err != nil {
		return false, NewRawdbError(HasBeaconBlockError, fmt.Errorf("has key %+v failed", keyHash))
	} else if ok {
		return true, nil
	}
	return false, nil
}

func GetBeaconBlockByHash(db incdb.Database, hash common.Hash) ([]byte, error) {
	keyHash := GetBeaconHashToBlockKey(hash)
	if ok, err := db.Has(keyHash); err != nil {
		return []byte{}, NewRawdbError(GetBeaconBlockByHashError, fmt.Errorf("has key %+v failed", keyHash))
	} else if !ok {
		return []byte{}, NewRawdbError(GetBeaconBlockByHashError, fmt.Errorf("block %+v not exist", hash))
	}
	block, err := db.Get(keyHash)
	if err != nil {
		return nil, NewRawdbError(GetBeaconBlockByHashError, err)
	}
	ret := make([]byte, len(block))
	copy(ret, block)
	return ret, nil
}

func GetBeaconBlockByIndex(db incdb.Database, index uint64) (map[common.Hash][]byte, error) {
	m := make(map[common.Hash][]byte)
	indexPrefix := GetBeaconIndexToBlockHashPrefix(index)
	iterator := db.NewIteratorWithPrefix(indexPrefix)
	for iterator.Next() {
		key := iterator.Key()
		strs := strings.Split(string(key), string(splitter))
		tempHash := []byte(strs[len(strs)-1])
		hash := common.BytesToHash(tempHash)
		keyHash := GetBeaconHashToBlockKey(hash)
		if ok, err := db.Has(keyHash); err != nil {
			return nil, NewRawdbError(GetBeaconBlockByIndexError, fmt.Errorf("has key %+v failed", keyHash))
		} else if !ok {
			return nil, NewRawdbError(GetBeaconBlockByIndexError, fmt.Errorf("block %+v not exist", hash))
		}
		block, err := db.Get(keyHash)
		if err != nil {
			return nil, NewRawdbError(GetBeaconBlockByIndexError, err)
		}
		ret := make([]byte, len(block))
		copy(ret, block)
		m[hash] = ret
	}
	return m, nil
}

func GetIndexOfBeaconBlock(db incdb.Database, hash common.Hash) (uint64, error) {
	key := GetBeaconBlockHashToIndexKey(hash)
	buf, err := db.Get(key)
	if err != nil {
		return 0, NewRawdbError(GetIndexOfBeaconBlockError, err)
	}
	index, err := common.BytesToUint64(buf)
	if err != nil {
		return 0, NewRawdbError(GetIndexOfBeaconBlockError, err)
	}
	return index, nil
}

func DeleteBeaconBlock(db incdb.Database, index uint64, hash common.Hash) error {
	keyHash := GetBeaconHashToBlockKey(hash)
	keyIndex := GetBeaconIndexToBlockHashKey(index, hash)
	if err := db.Delete(keyHash); err != nil {
		return NewRawdbError(DeleteBeaconBlockError, err)
	}
	if err := db.Delete(keyIndex); err != nil {
		return NewRawdbError(DeleteBeaconBlockError, err)
	}
	return nil
}

func GetBeaconBlockHashByIndex(db incdb.Database, index uint64) ([]common.Hash, error) {
	beaconBlockHashes := []common.Hash{}
	indexPrefix := GetBeaconIndexToBlockHashPrefix(index)
	iterator := db.NewIteratorWithPrefix(indexPrefix)
	for iterator.Next() {
		key := iterator.Key()
		strs := strings.Split(string(key), string(splitter))
		tempHash := []byte(strs[len(strs)-1])
		hash := common.BytesToHash(tempHash)
		beaconBlockHashes = append(beaconBlockHashes, hash)
	}
	return beaconBlockHashes, nil
}

func StoreBeaconBestState(db incdb.Database, v interface{}) error {
	val, err := json.Marshal(v)
	if err != nil {
		return NewRawdbError(StoreBeaconBestStateError, err)
	}
	key := GetBeaconBestStateKey()
	if err := db.Put(key, val); err != nil {
		return NewRawdbError(StoreBeaconBestStateError, err)
	}
	return nil
}

func FetchBeaconBestState(db incdb.Database) ([]byte, error) {
	key := GetBeaconBestStateKey()
	block, err := db.Get(key)
	if err != nil {
		return nil, NewRawdbError(FetchBeaconBestStateError, err)
	}
	return block, nil
}