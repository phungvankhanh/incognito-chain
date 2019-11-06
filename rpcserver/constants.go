package rpcserver

// rpc cmd method
const (
	// test rpc server
	testHttpServer = "testrpcserver"
	startProfiling = "startprofiling"
	stopProfiling  = "stopprofiling"

	getNetworkInfo       = "getnetworkinfo"
	getConnectionCount   = "getconnectioncount"
	getAllConnectedPeers = "getallconnectedpeers"
	getAllPeers          = "getallpeers"
	getNodeRole          = "getnoderole"
	getInOutMessages     = "getinoutmessages"
	getInOutMessageCount = "getinoutmessagecount"

	estimateFee              = "estimatefee"
	estimateFeeWithEstimator = "estimatefeewithestimator"

	getActiveShards    = "getactiveshards"
	getMaxShardsNumber = "getmaxshardsnumber"

	getMiningInfo                 = "getmininginfo"
	getRawMempool                 = "getrawmempool"
	getNumberOfTxsInMempool       = "getnumberoftxsinmempool"
	getMempoolEntry               = "getmempoolentry"
	removeTxInMempool             = "removetxinmempool"
	getBeaconPoolState            = "getbeaconpoolstate"
	getShardPoolState             = "getshardpoolstate"
	getShardPoolLatestValidHeight = "getshardpoollatestvalidheight"
	//getShardToBeaconPoolState     = "getshardtobeaconpoolstate"
	//getCrossShardPoolState        = "getcrossshardpoolstate"
	getNextCrossShard           = "getnextcrossshard"
	getShardToBeaconPoolStateV2 = "getshardtobeaconpoolstatev2"
	getCrossShardPoolStateV2    = "getcrossshardpoolstatev2"
	getShardPoolStateV2         = "getshardpoolstatev2"
	getBeaconPoolStateV2        = "getbeaconpoolstatev2"
	//getFeeEstimator             = "getfeeestimator"

	getBestBlock        = "getbestblock"
	getBestBlockHash    = "getbestblockhash"
	getBlocks           = "getblocks"
	retrieveBlock       = "retrieveblock"
	retrieveBeaconBlock = "retrievebeaconblock"
	getBlockChainInfo   = "getblockchaininfo"
	getBlockCount       = "getblockcount"
	getBlockHash        = "getblockhash"

	listOutputCoins                            = "listoutputcoins"
	createRawTransaction                       = "createtransaction"
	sendRawTransaction                         = "sendtransaction"
	createAndSendTransaction                   = "createandsendtransaction"
	createAndSendCustomTokenTransaction        = "createandsendcustomtokentransaction"
	sendRawCustomTokenTransaction              = "sendrawcustomtokentransaction"
	createRawCustomTokenTransaction            = "createrawcustomtokentransaction"
	createRawPrivacyCustomTokenTransaction     = "createrawprivacycustomtokentransaction"
	sendRawPrivacyCustomTokenTransaction       = "sendrawprivacycustomtokentransaction"
	createAndSendPrivacyCustomTokenTransaction = "createandsendprivacycustomtokentransaction"
	getMempoolInfo                             = "getmempoolinfo"
	getPendingTxsInBlockgen                    = "getpendingtxsinblockgen"
	getCandidateList                           = "getcandidatelist"
	getCommitteeList                           = "getcommitteelist"
	canPubkeyStake                             = "canpubkeystake"
	getTotalTransaction                        = "gettotaltransaction"
	listUnspentCustomToken                     = "listunspentcustomtoken"
	getBalanceCustomToken                      = "getbalancecustomtoken"
	getTransactionByHash                       = "gettransactionbyhash"
	gettransactionhashbyreceiver               = "gettransactionhashbyreceiver"
	gettransactionbyreceiver                   = "gettransactionbyreceiver"
	listCustomToken                            = "listcustomtoken"
	listPrivacyCustomToken                     = "listprivacycustomtoken"
	getBalancePrivacyCustomToken               = "getbalanceprivacycustomtoken"
	customTokenTxs                             = "customtoken"
	listCustomTokenHolders                     = "customtokenholder"
	privacyCustomTokenTxs                      = "privacycustomtoken"
	checkHashValue                             = "checkhashvalue"
	getListCustomTokenBalance                  = "getlistcustomtokenbalance"
	getListPrivacyCustomTokenBalance           = "getlistprivacycustomtokenbalance"
	getBlockHeader                             = "getheader"
	getCrossShardBlock                         = "getcrossshardblock"
	randomCommitments                          = "randomcommitments"
	hasSerialNumbers                           = "hasserialnumbers"
	hasSnDerivators                            = "hassnderivators"
	listSnDerivators                           = "listsnderivators"
	listSerialNumbers                          = "listserialnumbers"
	listCommitments                            = "listcommitments"
	listCommitmentIndices                      = "listcommitmentindices"
	createAndSendStakingTransaction            = "createandsendstakingtransaction"
	createAndSendStopAutoStakingTransaction    = "createandsendstopautostakingtransaction"

	//===========For Testing and Benchmark==============
	getAndSendTxsFromFile   = "getandsendtxsfromfile"
	getAndSendTxsFromFileV2 = "getandsendtxsfromfilev2"
	unlockMempool           = "unlockmempool"
	//==================================================

	getShardBestState        = "getshardbeststate"
	getShardBestStateDetail  = "getshardbeststatedetail"
	getBeaconBestState       = "getbeaconbeststate"
	getBeaconBestStateDetail = "getbeaconbeststatedetail"

	// Wallet rpc cmd
	listAccounts               = "listaccounts"
	getAccount                 = "getaccount"
	getAddressesByAccount      = "getaddressesbyaccount"
	getAccountAddress          = "getaccountaddress"
	dumpPrivkey                = "dumpprivkey"
	importAccount              = "importaccount"
	removeAccount              = "removeaccount"
	listUnspentOutputCoins     = "listunspentoutputcoins"
	getBalance                 = "getbalance"
	getBalanceByPrivatekey     = "getbalancebyprivatekey"
	getBalanceByPaymentAddress = "getbalancebypaymentaddress"
	getReceivedByAccount       = "getreceivedbyaccount"
	setTxFee                   = "settxfee"

	// walletsta
	getPublicKeyFromPaymentAddress = "getpublickeyfrompaymentaddress"
	defragmentAccount              = "defragmentaccount"

	getStackingAmount = "getstackingamount"

	// utils
	hashToIdenticon = "hashtoidenticon"
	generateTokenID = "generatetokenid"

	createIssuingRequest             = "createissuingrequest"
	sendIssuingRequest               = "sendissuingrequest"
	createAndSendIssuingRequest      = "createandsendissuingrequest"
	createAndSendContractingRequest  = "createandsendcontractingrequest"
	createAndSendBurningRequest      = "createandsendburningrequest"
	createAndSendTxWithIssuingETHReq = "createandsendtxwithissuingethreq"
	checkETHHashIssued               = "checkethhashissued"
	getAllBridgeTokens               = "getallbridgetokens"
	getETHHeaderByHash               = "getethheaderbyhash"
	getBridgeReqWithStatus           = "getbridgereqwithstatus"

	// Incognito -> Ethereum bridge
	getBeaconSwapProof       = "getbeaconswapproof"
	getLatestBeaconSwapProof = "getlatestbeaconswapproof"
	getBridgeSwapProof       = "getbridgeswapproof"
	getLatestBridgeSwapProof = "getlatestbridgeswapproof"
	getBurnProof             = "getburnproof"

	// reward
	CreateRawWithDrawTransaction = "withdrawreward"
	getRewardAmount              = "getrewardamount"
	listRewardAmount             = "listrewardamount"

	revertbeaconchain = "revertbeaconchain"
	revertshardchain  = "revertshardchain"

	enableMining                = "enablemining"
	getChainMiningStatus        = "getchainminingstatus"
	getPublickeyMining          = "getpublickeymining"
	getPublicKeyRole            = "getpublickeyrole"
	getIncognitoPublicKeyRole   = "getincognitopublickeyrole"
	getMinerRewardFromMiningKey = "getminerrewardfromminingkey"

	// slash
	getProducersBlackList       = "getproducersblacklist"
	getProducersBlackListDetail = "getproducersblacklistdetail"

	// pde
	getPDEState                           = "getpdestate"
	createAndSendTxWithWithdrawalReq      = "createandsendtxwithwithdrawalreq"
	createAndSendTxWithPTokenTradeReq     = "createandsendtxwithptokentradereq"
	createAndSendTxWithPRVTradeReq        = "createandsendtxwithprvtradereq"
	createAndSendTxWithPTokenContribution = "createandsendtxwithptokencontribution"
	createAndSendTxWithPRVContribution    = "createandsendtxwithprvcontribution"
)

const (
	testSubcrice                                = "testsubcribe"
	subcribeNewShardBlock                       = "subcribenewshardblock"
	subcribeNewBeaconBlock                      = "subcribenewbeaconblock"
	subcribePendingTransaction                  = "subcribependingtransaction"
	subcribeShardCandidateByPublickey           = "subcribeshardcandidatebypublickey"
	subcribeShardPendingValidatorByPublickey    = "subcribeshardpendingvalidatorbypublickey"
	subcribeShardCommitteeByPublickey           = "subcribeshardcommitteebypublickey"
	subcribeBeaconCandidateByPublickey          = "subcribebeaconcandidatebypublickey"
	subcribeBeaconPendingValidatorByPublickey   = "subcribebeaconpendingvalidatorbypublickey"
	subcribeBeaconCommitteeByPublickey          = "subcribebeaconcommitteebypublickey"
	subcribeCrossOutputCoinByPrivateKey         = "subcribecrossoutputcoinbyprivatekey"
	subcribeCrossCustomTokenByPrivateKey        = "subcribecrosscustomtokenbyprivatekey"
	subcribeCrossCustomTokenPrivacyByPrivateKey = "subcribecrosscustomtokenprivacybyprivatekey"
	subcribeMempoolInfo                         = "subcribemempoolinfo"
	subcribeShardBestState                      = "subcribeshardbeststate"
	subcribeBeaconBestState                     = "subcribebeaconbeststate"
	subcribeBeaconPoolBeststate                 = "subcribebeaconpoolbeststate"
	subcribeShardPoolBeststate                  = "subcribeshardpoolbeststate"
)
