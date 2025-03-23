package node_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justinwongcn/go-ethlibs/eth"
	"github.com/justinwongcn/go-ethlibs/node"
)

func getRopstenClient(t *testing.T, ctx context.Context) node.Client {
	// These test require a ropsten websocket URL to test with, for example ws://localhost:8546 or wss://ropsten.infura.io/ws/v3/:YOUR_PROJECT_ID
	url := os.Getenv("ETHLIBS_TEST_ROPSTEN_WS_URL")
	if url == "" {
		t.Skip("ETHLIBS_TEST_ROPSTEN_WS_URL not set, skipping test.  Set to a valid websocket URL to execute this test.")
	}

	conn, err := node.NewClient(ctx, url)
	require.NoError(t, err, "creating websocket connection should not fail")
	return conn
}

func TestConnection_GetTransactionCount(t *testing.T) {
	ctx := context.Background()
	conn := getRopstenClient(t, ctx)

	// Checks the current pending nonce for account can be retrieved
	blockNum1 := eth.MustBlockNumberOrTag("latest")
	pendingNonce1, err := conn.GetTransactionCount(ctx, "0xed28874e52A12f0D42118653B0FBCee0ACFadC00", *blockNum1)
	require.NoError(t, err)
	require.NotEmpty(t, pendingNonce1, "pending nonce must not be nil")

	// Should catch failure since it is looking for a nonce of a future block
	blockNum2 := eth.MustBlockNumberOrTag("0x7654321")
	pendingNonce2, err := conn.GetTransactionCount(ctx, "0xed28874e52A12f0D42118653B0FBCee0ACFadC00", *blockNum2)
	require.Error(t, err)
	require.Empty(t, pendingNonce2, "pending nonce must not exist since it is a future block")
}

func TestConnection_GetBalance(t *testing.T) {
	ctx := context.Background()
	conn := getRopstenClient(t, ctx)

	// Test with valid address and latest block
	blockNum1 := eth.MustBlockNumberOrTag("latest")
	balance1, err := conn.GetBalance(ctx, "0xed28874e52A12f0D42118653B0FBCee0ACFadC00", *blockNum1)
	require.NoError(t, err)
	require.NotEqual(t, uint64(0), balance1, "balance must not be zero")

	// Test with invalid address format
	balance2, err := conn.GetBalance(ctx, "invalid", *blockNum1)
	require.Error(t, err, "requesting with invalid address should return an error")
	require.Equal(t, uint64(0), balance2, "balance should be 0 for invalid address")

	// Test with future block
	blockNum2 := eth.MustBlockNumberOrTag("0x7654321")
	balance3, err := conn.GetBalance(ctx, "0xed28874e52A12f0D42118653B0FBCee0ACFadC00", *blockNum2)
	require.Error(t, err)
	require.Equal(t, uint64(0), balance3, "balance must be zero for future block")
}

func TestConnection_EstimateGas(t *testing.T) {
	ctx := context.Background()
	conn := getRopstenClient(t, ctx)

	from := eth.MustAddress("0xed28874e52A12f0D42118653B0FBCee0ACFadC00")
	tx := eth.Transaction{
		Nonce:    eth.QuantityFromUInt64(146),
		GasPrice: eth.OptionalQuantityFromInt(3000000000),
		Gas:      eth.QuantityFromUInt64(22000),
		To:       eth.MustAddress("0x43700db832E9Ac990D36d6279A846608643c904E"),
		Value:    *eth.OptionalQuantityFromInt(100),
		From:     *from,
	}

	gas, err := conn.EstimateGas(ctx, tx)
	require.NoError(t, err)
	require.NotEqual(t, gas, 0, "estimate gas cannot be equal to zero.")
}

func TestConnection_MaxPriorityFeePerGas(t *testing.T) {
	ctx := context.Background()
	conn := getRopstenClient(t, ctx)

	fee, err := conn.MaxPriorityFeePerGas(ctx)
	require.NoError(t, err)
	require.NotEqual(t, fee, 0, "fee cannot be equal to 0")
}

func TestConnection_GasPrice(t *testing.T) {
	ctx := context.Background()
	conn := getRopstenClient(t, ctx)

	gasPrice, err := conn.GasPrice(ctx)
	require.NoError(t, err)
	require.NotEqual(t, gasPrice, 0, "gas price cannot be equal to 0")
}

func TestConnection_NetVersion(t *testing.T) {
	ctx := context.Background()
	conn := getRopstenClient(t, ctx)

	netVersion, err := conn.NetVersion(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, netVersion, "net version id must not be nil")
}

func TestConnection_ChainId(t *testing.T) {
	ctx := context.Background()
	conn := getRopstenClient(t, ctx)

	chainId, err := conn.ChainId(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, chainId, "chain id must not be nil")
}

func TestConnection_SendRawTransactionInValidEmpty(t *testing.T) {
	ctx := context.Background()
	conn := getRopstenClient(t, ctx)

	txHash, err := conn.SendRawTransaction(ctx, "0x0")
	require.Error(t, err)
	require.Empty(t, txHash, "txHash must be nil")
}

func TestConnection_SendRawTransactionInValidOldNonce(t *testing.T) {
	ctx := context.Background()
	conn := getRopstenClient(t, ctx)

	data := eth.MustData("0x02f8f70338849502f8f3849502f8f3826c3994b78ab5a21c74451906d6a113072e6aa2f2d905b980b88cf56256c730783078343836353663366336663030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303078307835373666373236633634323130303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030c001a0e2fd5de027d939a99df69954cd36a9f7cac6f3c4bf96eff48b7980be9394a1d7a06f0e4b4fa4642afa99f5caa74f004c93707c6503c7beb7e746352081d77ec054")
	txHash, err := conn.SendRawTransaction(ctx, data.String())
	require.Error(t, err)
	require.Equal(t, err.Error(), "{\"code\":-32000,\"message\":\"nonce too low\"}")
	require.Empty(t, txHash, "txHash must be nil")
}

func TestConnection_FutureBlockByNumber(t *testing.T) {
	ctx := context.Background()
	conn := getRopstenClient(t, ctx)

	blockNumber, err := conn.BlockNumber(ctx)
	require.NoError(t, err)

	next, err := conn.BlockByNumber(ctx, blockNumber+1000, false)
	require.Nil(t, next, "future block should be nil")
	require.Error(t, err, "requesting a future block should return an error")
	require.Equal(t, node.ErrBlockNotFound, err)

	// get a the genesis block by number which should _not_ fail
	genesis, err := conn.BlockByNumber(ctx, 0, false)
	require.NoError(t, err, "requesting genesis block by number should not fail")
	require.NotNil(t, genesis, "genesis block must not be nil")
}

func TestConnection_InvalidBlockByHash(t *testing.T) {
	ctx := context.Background()
	conn := getRopstenClient(t, ctx)

	b, err := conn.BlockByHash(ctx, "invalid", false)
	require.Error(t, err, "requesting an invalid hash should return an error")
	require.Nil(t, b, "block from invalid hash should be nil")

	b, err = conn.BlockByHash(ctx, "0x1234", false)
	require.Error(t, err, "requesting an invalid hash should return an error")
	require.Nil(t, b, "block from invalid hash should be nil")

	b, err = conn.BlockByHash(ctx, "0x0badf00dbadf00dbadf00dbadf00dbadf00dbadf00dbadf00dbadf00dbadf00d", false)
	require.Error(t, err, "requesting a non-existent block should should return an error")
	require.Nil(t, b, "block from non-existent hash should be nil")
	require.Equal(t, node.ErrBlockNotFound, err)

	// get the genesis block which should _not_ fail
	b, err = conn.BlockByHash(ctx, "0x41941023680923e0fe4d74a34bdac8141f2540e3ae90623718e47d66d1ca4a2d", true)
	require.NoError(t, err, "genesis block hash should not return an error")
	require.NotNil(t, b, "genesis block should be retrievable by hash")
}

func TestConnection_GetBlockTransactionCountByHash(t *testing.T) {
	ctx := context.Background()
	conn := getRopstenClient(t, ctx)

	// Test with valid block hash (genesis block)
	count1, err := conn.GetBlockTransactionCountByHash(ctx, "0x41941023680923e0fe4d74a34bdac8141f2540e3ae90623718e47d66d1ca4a2d")
	require.NoError(t, err)
	require.NotNil(t, count1, "transaction count must not be nil for valid block hash")

	// Test with invalid hash format
	count2, err := conn.GetBlockTransactionCountByHash(ctx, "invalid")
	require.Error(t, err, "requesting with invalid hash format should return an error")
	require.Equal(t, uint64(0), count2, "transaction count should be 0 for invalid hash format")

	// Test with non-existent block hash
	count3, err := conn.GetBlockTransactionCountByHash(ctx, "0x0badf00dbadf00dbadf00dbadf00dbadf00dbadf00dbadf00dbadf00dbadf00d")
	require.Error(t, err)
	require.Equal(t, uint64(0), count3, "transaction count should be 0 for non-existent block")
	require.Equal(t, node.ErrBlockNotFound, err)
}

func TestConnection_GetBlockTransactionCountByNumber(t *testing.T) {
	ctx := context.Background()
	conn := getRopstenClient(t, ctx)

	// Test with valid block number (genesis block)
	count1, err := conn.GetBlockTransactionCountByNumber(ctx, *eth.MustBlockNumberOrTag("0x0"))
	require.NoError(t, err)
	require.NotNil(t, count1, "transaction count must not be nil for valid block number")

	// Test with latest block tag
	count2, err := conn.GetBlockTransactionCountByNumber(ctx, *eth.MustBlockNumberOrTag("latest"))
	require.NoError(t, err)
	require.NotNil(t, count2, "transaction count must not be nil for latest block")

	// Test with future block number
	count3, err := conn.GetBlockTransactionCountByNumber(ctx, *eth.MustBlockNumberOrTag("0x7654321"))
	require.Error(t, err)
	require.Equal(t, uint64(0), count3, "transaction count should be 0 for future block")
	require.Equal(t, node.ErrBlockNotFound, err)
}

func TestConnection_GetCode(t *testing.T) {
	ctx := context.Background()
	conn := getRopstenClient(t, ctx)

	// Test with valid address and latest block
	blockNum1 := eth.MustBlockNumberOrTag("latest")
	code1, err := conn.GetCode(ctx, "0xed28874e52A12f0D42118653B0FBCee0ACFadC00", *blockNum1)
	require.NoError(t, err)
	require.NotEmpty(t, code1, "code must not be empty for valid address")

	// Test with invalid address format
	code2, err := conn.GetCode(ctx, "invalid", *blockNum1)
	require.Error(t, err, "requesting with invalid address should return an error")
	require.Empty(t, code2, "code should be empty for invalid address")

	// Test with future block
	blockNum2 := eth.MustBlockNumberOrTag("0x7654321")
	code3, err := conn.GetCode(ctx, "0xed28874e52A12f0D42118653B0FBCee0ACFadC00", *blockNum2)
	require.Error(t, err)
	require.Empty(t, code3, "code must be empty for future block")
}

func TestConnection_Sign(t *testing.T) {
	ctx := context.Background()
	conn := getRopstenClient(t, ctx)

	// Test with valid address and message
	signature1, err := conn.Sign(ctx, "0xed28874e52A12f0D42118653B0FBCee0ACFadC00", "0xdeadbeaf")
	require.NoError(t, err)
	require.NotEmpty(t, signature1, "signature must not be empty for valid parameters")

	// Test with invalid address format
	signature2, err := conn.Sign(ctx, "invalid", "0xdeadbeaf")
	require.Error(t, err, "requesting with invalid address should return an error")
	require.Empty(t, signature2, "signature should be empty for invalid address")

	// Test with empty message
	signature3, err := conn.Sign(ctx, "0xed28874e52A12f0D42118653B0FBCee0ACFadC00", "")
	require.Error(t, err)
	require.Empty(t, signature3, "signature must be empty for empty message")
}

func TestConnection_Call(t *testing.T) {
	ctx := context.Background()
	conn := getRopstenClient(t, ctx)

	// Test with valid transaction
	tx := eth.Transaction{
		From:     *eth.MustAddress("0xed28874e52A12f0D42118653B0FBCee0ACFadC00"),
		To:       eth.MustAddress("0x43700db832E9Ac990D36d6279A846608643c904E"),
		Gas:      eth.QuantityFromUInt64(30400),
		GasPrice: eth.OptionalQuantityFromInt(10000000000000),
		Value:    *eth.OptionalQuantityFromInt(2441406250),
		Input:    *eth.MustInput("0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"),
	}
	blockNum := eth.MustBlockNumberOrTag("latest")
	result1, err := conn.Call(ctx, tx, *blockNum)
	require.NoError(t, err)
	require.NotEmpty(t, result1, "call result must not be empty for valid transaction")

	// Test with invalid from address
	tx.From = *eth.MustAddress("0x0000000000000000000000000000000000000000")
	result2, err := conn.Call(ctx, tx, *blockNum)
	require.Error(t, err)
	require.Empty(t, result2, "call result should be empty for invalid from address")

	// Test with future block
	futureBlock := eth.MustBlockNumberOrTag("0x7654321")
	result3, err := conn.Call(ctx, tx, *futureBlock)
	require.Error(t, err)
	require.Empty(t, result3, "call result must be empty for future block")
}

func TestConnection_SendTransaction(t *testing.T) {
	ctx := context.Background()
	conn := getRopstenClient(t, ctx)

	// Test with valid transaction
	tx := eth.Transaction{
		From:     *eth.MustAddress("0xed28874e52A12f0D42118653B0FBCee0ACFadC00"),
		To:       eth.MustAddress("0x43700db832E9Ac990D36d6279A846608643c904E"),
		Gas:      eth.QuantityFromUInt64(30400),
		GasPrice: eth.OptionalQuantityFromInt(10000000000000),
		Value:    *eth.OptionalQuantityFromInt(2441406250),
		Input:    *eth.MustInput("0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"),
	}
	txHash1, err := conn.SendTransaction(ctx, tx)
	require.NoError(t, err)
	require.NotEmpty(t, txHash1, "transaction hash must not be empty for valid transaction")

	// Test with invalid from address
	tx.From = *eth.MustAddress("0x0000000000000000000000000000000000000000")
	txHash2, err := conn.SendTransaction(ctx, tx)
	require.Error(t, err)
	require.Empty(t, txHash2, "transaction hash should be empty for invalid from address")

	// Test with empty transaction
	txHash3, err := conn.SendTransaction(ctx, eth.Transaction{})
	require.Error(t, err)
	require.Empty(t, txHash3, "transaction hash must be empty for empty transaction")
}

func TestConnection_InvalidTransactionByHash(t *testing.T) {
	ctx := context.Background()
	conn := getRopstenClient(t, ctx)

	tx, err := conn.TransactionByHash(ctx, "invalid")
	require.Error(t, err, "requesting an invalid hash should return an error")
	require.Nil(t, tx, "tx from invalid hash should be nil")

	tx, err = conn.TransactionByHash(ctx, "0x1234")
	require.Error(t, err, "requesting an invalid hash should return an error")
	require.Nil(t, tx, "tx from invalid hash should be nil")

	tx, err = conn.TransactionByHash(ctx, "0x0badf00dbadf00dbadf00dbadf00dbadf00dbadf00dbadf00dbadf00dbadf00d")
	require.Error(t, err, "requesting an non-existent hash should return an error")
	require.Nil(t, tx, "tx from non-existent hash should be nil")
	require.Equal(t, node.ErrTransactionNotFound, err)

	// get an early tx which should _not_ fail
	tx, err = conn.TransactionByHash(ctx, "0x230f6e1739286f9cbf768e34a9ff3d69a2a72b92c8c3383fbdf163035c695332")
	require.NoError(t, err, "early tx should not return an error")
	require.NotNil(t, tx, "early tx should be retrievable by hash")
}

func TestConnection_GetTransactionByBlockHashAndIndex(t *testing.T) {
	ctx := context.Background()
	conn := getRopstenClient(t, ctx)

	// Get a block with transactions to test with
	latestBlock, err := conn.BlockByNumberOrTag(ctx, *eth.MustBlockNumberOrTag("latest"), true)
	require.NoError(t, err, "getting latest block should not fail")

	// If the latest block has transactions, use it for testing
	var blockHash string
	var txIndex uint64
	if len(latestBlock.Transactions) > 0 {
		blockHash = latestBlock.Hash.String()
		txIndex = 0 // Use the first transaction
	} else {
		// Use a known block with transactions (genesis block or another known block)
		// For this test, we'll use a known transaction from Ropsten
		blockHash = "0x41941023680923e0fe4d74a34bdac8141f2540e3ae90623718e47d66d1ca4a2d" // Genesis block or another known block
		txIndex = 0                                                                      // Adjust based on the block you choose
	}

	// Test with valid block hash and index
	tx1, err := conn.GetTransactionByBlockHashAndIndex(ctx, blockHash, txIndex)
	if err == nil {
		require.NotNil(t, tx1, "transaction must not be nil for valid block hash and index")
		require.NotEmpty(t, tx1.Hash, "transaction hash must not be empty")
	} else {
		// If the specific block doesn't have transactions at the index, this is expected
		require.Equal(t, node.ErrTransactionNotFound, err)
	}

	// Test with invalid hash format
	tx2, err := conn.GetTransactionByBlockHashAndIndex(ctx, "invalid", 0)
	require.Error(t, err, "requesting with invalid hash format should return an error")
	require.Nil(t, tx2, "transaction should be nil for invalid hash format")

	// Test with non-existent block hash
	tx3, err := conn.GetTransactionByBlockHashAndIndex(ctx, "0x0badf00dbadf00dbadf00dbadf00dbadf00dbadf00dbadf00dbadf00dbadf00d", 0)
	require.Error(t, err, "requesting with non-existent block hash should return an error")
	require.Nil(t, tx3, "transaction should be nil for non-existent block hash")

	// Test with valid block hash but out of range index
	tx4, err := conn.GetTransactionByBlockHashAndIndex(ctx, blockHash, 9999)
	require.Error(t, err, "requesting with out of range index should return an error")
	require.Nil(t, tx4, "transaction should be nil for out of range index")
	require.Equal(t, node.ErrTransactionNotFound, err)
}

func TestConnection_GetTransactionByBlockNumberAndIndex(t *testing.T) {
	ctx := context.Background()
	conn := getRopstenClient(t, ctx)

	// Get the latest block number
	blockNumber, err := conn.BlockNumber(ctx)
	require.NoError(t, err, "getting block number should not fail")

	// Get a block with transactions to test with
	latestBlock, err := conn.BlockByNumber(ctx, blockNumber, true)
	require.NoError(t, err, "getting latest block should not fail")

	// If the latest block has transactions, use it for testing
	var testBlockNumber uint64
	var txIndex uint64
	if len(latestBlock.Transactions) > 0 {
		testBlockNumber = blockNumber
		txIndex = 0 // Use the first transaction
	} else {
		// Use a known block with transactions (genesis block or another known block)
		testBlockNumber = 0 // Genesis block or another known block with transactions
		txIndex = 0         // Adjust based on the block you choose
	}

	// Test with valid block number and index
	blockNum1 := eth.MustBlockNumberOrTag(eth.QuantityFromUInt64(testBlockNumber).String())
	tx1, err := conn.GetTransactionByBlockNumberAndIndex(ctx, *blockNum1, txIndex)
	if err == nil {
		require.NotNil(t, tx1, "transaction must not be nil for valid block number and index")
		require.NotEmpty(t, tx1.Hash, "transaction hash must not be empty")
	} else {
		// If the specific block doesn't have transactions at the index, this is expected
		require.Equal(t, node.ErrTransactionNotFound, err)
	}

	// Test with latest block tag
	blockNum2 := eth.MustBlockNumberOrTag("latest")
	tx2, err := conn.GetTransactionByBlockNumberAndIndex(ctx, *blockNum2, txIndex)
	if err == nil {
		require.NotNil(t, tx2, "transaction must not be nil for latest block and valid index")
		require.NotEmpty(t, tx2.Hash, "transaction hash must not be empty")
	} else {
		// If the latest block doesn't have transactions at the index, this is expected
		require.Equal(t, node.ErrTransactionNotFound, err)
	}

	// Test with future block number
	blockNum3 := eth.MustBlockNumberOrTag("0x7654321")
	tx3, err := conn.GetTransactionByBlockNumberAndIndex(ctx, *blockNum3, 0)
	require.Error(t, err, "requesting with future block number should return an error")
	require.Nil(t, tx3, "transaction should be nil for future block number")

	// Test with valid block number but out of range index
	tx4, err := conn.GetTransactionByBlockNumberAndIndex(ctx, *blockNum1, 9999)
	require.Error(t, err, "requesting with out of range index should return an error")
	require.Nil(t, tx4, "transaction should be nil for out of range index")
	require.Equal(t, node.ErrTransactionNotFound, err)
}

func TestConnection_GetUncleByBlockHashAndIndex(t *testing.T) {
	ctx := context.Background()
	conn := getRopstenClient(t, ctx)

	// Get a block with uncles to test with
	latestBlock, err := conn.BlockByNumberOrTag(ctx, *eth.MustBlockNumberOrTag("latest"), true)
	require.NoError(t, err, "getting latest block should not fail")

	// If the latest block has uncles, use it for testing
	var blockHash string
	var uncleIndex uint64
	if len(latestBlock.Uncles) > 0 {
		blockHash = latestBlock.Hash.String()
		uncleIndex = 0 // Use the first uncle
	} else {
		// Use a known block with uncles
		blockHash = "0x41941023680923e0fe4d74a34bdac8141f2540e3ae90623718e47d66d1ca4a2d" // Genesis block or another known block
		uncleIndex = 0
	}

	// Test with valid block hash and index
	uncle1, err := conn.GetUncleByBlockHashAndIndex(ctx, blockHash, uncleIndex)
	if err == nil {
		require.NotNil(t, uncle1, "uncle must not be nil for valid block hash and index")
		require.NotEmpty(t, uncle1.Hash, "uncle hash must not be empty")
	} else {
		// If the specific block doesn't have uncles at the index, this is expected
		require.Equal(t, node.ErrBlockNotFound, err)
	}

	// Test with invalid hash format
	uncle2, err := conn.GetUncleByBlockHashAndIndex(ctx, "invalid", 0)
	require.Error(t, err, "requesting with invalid hash format should return an error")
	require.Nil(t, uncle2, "uncle should be nil for invalid hash format")

	// Test with non-existent block hash
	uncle3, err := conn.GetUncleByBlockHashAndIndex(ctx, "0x0badf00dbadf00dbadf00dbadf00dbadf00dbadf00dbadf00dbadf00dbadf00d", 0)
	require.Error(t, err, "requesting with non-existent block hash should return an error")
	require.Nil(t, uncle3, "uncle should be nil for non-existent block hash")

	// Test with valid block hash but out of range index
	uncle4, err := conn.GetUncleByBlockHashAndIndex(ctx, blockHash, 9999)
	require.Error(t, err, "requesting with out of range index should return an error")
	require.Nil(t, uncle4, "uncle should be nil for out of range index")
	require.Equal(t, node.ErrBlockNotFound, err)
}

func TestConnection_GetUncleByBlockNumberAndIndex(t *testing.T) {
	ctx := context.Background()
	conn := getRopstenClient(t, ctx)

	// Get the latest block number
	blockNumber, err := conn.BlockNumber(ctx)
	require.NoError(t, err, "getting block number should not fail")

	// Get a block with uncles to test with
	latestBlock, err := conn.BlockByNumber(ctx, blockNumber, true)
	require.NoError(t, err, "getting latest block should not fail")

	// If the latest block has uncles, use it for testing
	var testBlockNumber uint64
	var uncleIndex uint64
	if len(latestBlock.Uncles) > 0 {
		testBlockNumber = blockNumber
		uncleIndex = 0 // Use the first uncle
	} else {
		// Use a known block with uncles
		testBlockNumber = 0 // Genesis block or another known block with uncles
		uncleIndex = 0
	}

	// Test with valid block number and index
	blockNum1 := eth.MustBlockNumberOrTag(eth.QuantityFromUInt64(testBlockNumber).String())
	uncle1, err := conn.GetUncleByBlockNumberAndIndex(ctx, *blockNum1, uncleIndex)
	if err == nil {
		require.NotNil(t, uncle1, "uncle must not be nil for valid block number and index")
		require.NotEmpty(t, uncle1.Hash, "uncle hash must not be empty")
	} else {
		// If the specific block doesn't have uncles at the index, this is expected
		require.Equal(t, node.ErrBlockNotFound, err)
	}

	// Test with latest block tag
	blockNum2 := eth.MustBlockNumberOrTag("latest")
	uncle2, err := conn.GetUncleByBlockNumberAndIndex(ctx, *blockNum2, uncleIndex)
	if err == nil {
		require.NotNil(t, uncle2, "uncle must not be nil for latest block and valid index")
		require.NotEmpty(t, uncle2.Hash, "uncle hash must not be empty")
	} else {
		// If the latest block doesn't have uncles at the index, this is expected
		require.Equal(t, node.ErrBlockNotFound, err)
	}

	// Test with future block number
	blockNum3 := eth.MustBlockNumberOrTag("0x7654321")
	uncle3, err := conn.GetUncleByBlockNumberAndIndex(ctx, *blockNum3, 0)
	require.Error(t, err, "requesting with future block number should return an error")
	require.Nil(t, uncle3, "uncle should be nil for future block number")

	// Test with valid block number but out of range index
	uncle4, err := conn.GetUncleByBlockNumberAndIndex(ctx, *blockNum1, 9999)
	require.Error(t, err, "requesting with out of range index should return an error")
	require.Nil(t, uncle4, "uncle should be nil for out of range index")
	require.Equal(t, node.ErrBlockNotFound, err)
}
