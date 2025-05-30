package eth_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justinwongcn/go-ethlibs/eth"
)

func TestTransactionReceipts(t *testing.T) {
	raw := `{
    "blockHash": "0xa37f46c4692db33012c105a27b9e4c582e822ed60a54667875fb92def52fd75a",
    "blockNumber": "0x72991c",
    "contractAddress": null,
    "cumulativeGasUsed": "0x7650c2",
    "from": "0x9e44b7d42125b7bb4e809406ed5e1079ff500969",
    "gasUsed": "0x5630",
    "logs": [
      {
        "address": "0x21ab6c9fac80c59d401b37cb43f81ea9dde7fe34",
        "blockHash": "0xa37f46c4692db33012c105a27b9e4c582e822ed60a54667875fb92def52fd75a",
        "blockNumber": "0x72991c",
        "data": "0x000000000000000000000000000000000000000000000000000000070560c8c0",
        "logIndex": "0xb7",
        "removed": false,
        "topics": [
          "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
          "0x0000000000000000000000009e44b7d42125b7bb4e809406ed5e1079ff500969",
          "0x000000000000000000000000fe5854255eb1eb921525fa856a3947ed2412a1d7"
        ],
        "transactionHash": "0x9d2fb08850a9b38173044ae6a61974fde4eacca504e399ffd9d5c8af567113cc",
        "transactionIndex": "0x8a"
      }
    ],
    "logsBloom": "0x00000001000000000000000000000000000000000000000000008000000000000000000000010000000000000000000000400000000000000000000000000008000000000000000000000008000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000200000000002000000000000000000000000000000000000000000000000000000000000000008000000000000000000000010000000000000000000000000000000",
    "status": "0x1",
    "to": "0x21ab6c9fac80c59d401b37cb43f81ea9dde7fe34",
    "transactionHash": "0x9d2fb08850a9b38173044ae6a61974fde4eacca504e399ffd9d5c8af567113cc",
    "transactionIndex": "0x8a"
  }`

	receipt := eth.TransactionReceipt{}
	err := json.Unmarshal([]byte(raw), &receipt)
	require.NoError(t, err, "unmarshal must succeed")

	require.Equal(t, uint64(0x72991c), receipt.BlockNumber.UInt64())
	require.Equal(t, "0xa37f46c4692db33012c105a27b9e4c582e822ed60a54667875fb92def52fd75a", receipt.BlockHash.String())
	require.Nil(t, receipt.ContractAddress)
	require.Equal(t, uint64(0x7650c2), receipt.CumulativeGasUsed.UInt64())
	require.Equal(t, uint64(0x5630), receipt.GasUsed.UInt64())
	require.Equal(t, *eth.MustAddress("0x9e44b7d42125b7bb4e809406ed5e1079ff500969"), receipt.From)
	require.Equal(t, eth.MustAddress("0x21ab6c9fac80c59d401b37cb43f81ea9dde7fe34"), receipt.To)
	require.Equal(t, uint64(1), receipt.Status.UInt64())
	require.Equal(t, *eth.MustHash("0x9d2fb08850a9b38173044ae6a61974fde4eacca504e399ffd9d5c8af567113cc"), receipt.TransactionHash)
	require.Equal(t, uint64(0x8a), receipt.TransactionIndex.UInt64())

	require.Equal(t, *eth.MustTopic("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"), receipt.Logs[0].Topics[0])
	require.Equal(t, *eth.MustData256("0x00000001000000000000000000000000000000000000000000008000000000000000000000010000000000000000000000400000000000000000000000000008000000000000000000000008000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000200000000002000000000000000000000000000000000000000000000000000000000000000008000000000000000000000010000000000000000000000000000000"), receipt.LogsBloom)
	require.Nil(t, receipt.Type)
	require.Equal(t, eth.TransactionTypeLegacy, receipt.TransactionType())

	// double check that we can back back to JSON as well
	b, err := json.Marshal(&receipt)
	require.NoError(t, err, "marshal must succeed")
	require.JSONEq(t, raw, string(b))

	// Let double check that contract creation receipts work too
	creation := `{
    "blockHash": "0xaacadbbc77f8962c0f2749ca12145ddfd09857c4ef4d6caa507d0afda7c200f5",
    "blockNumber": "0x578049",
    "contractAddress": "0x21ab6c9fac80c59d401b37cb43f81ea9dde7fe34",
    "cumulativeGasUsed": "0x4fac0d",
    "from": "0x4ea0d7225e384582d6ea31e34260bf7ac0c1127f",
    "gasUsed": "0xe85d5",
    "logs": [],
    "logsBloom": "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
    "status": "0x1",
    "to": null,
    "transactionHash": "0x5e07f2daaa7fad0d59a19ba2bce54d1b58be25f5c8dd684e1f27e59c3d5f92c6",
    "transactionIndex": "0x6a"
  }`

	{
		receipt := eth.TransactionReceipt{}
		err := json.Unmarshal([]byte(creation), &receipt)
		require.NoError(t, err, "unmarshal must succeed")

		require.Equal(t, eth.MustAddress("0x21ab6c9fac80c59d401b37cb43f81ea9dde7fe34"), receipt.ContractAddress)
		require.Nil(t, receipt.To)
		require.Nil(t, receipt.Type)
		require.Equal(t, eth.TransactionTypeLegacy, receipt.TransactionType())
	}

	// And a pre-byzantine one while we're here
	old := `{
    "blockHash": "0xcd6d29f6b644e82252823053c2e051bab2461f24d3d32b7bb2e5391452f2386e",
    "blockNumber": "0x7a122",
    "contractAddress": null,
    "cumulativeGasUsed": "0x1a7a1",
    "from": "0x119058dc2c577e9c4ba6914678aa9db565300ffe",
    "gasUsed": "0x723c",
    "logs": [
      {
        "address": "0x46a9a148d617138cb5c0346de289c030856bb716",
        "blockHash": "0xcd6d29f6b644e82252823053c2e051bab2461f24d3d32b7bb2e5391452f2386e",
        "blockNumber": "0x7a122",
        "data": "0x000000000000000000000000119058dc2c577e9c4ba6914678aa9db565300ffe000000000000000000000000000000000000000000000a968163f0a57b400000",
        "logIndex": "0x1",
        "removed": false,
        "topics": [
          "0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c"
        ],
        "transactionHash": "0x45215aa0da9b7597d233d96b6f7c4ac311edaba77a99ecc6471c59663554914f",
        "transactionIndex": "0x1"
      }
    ],
    "logsBloom": "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000008000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000400000000000000000",
    "root": "0x57bd5108d8f0b8bad735ab77e2a47b80c166dcf5059b2960e0118b40562c7cf2",
    "to": "0x46a9a148d617138cb5c0346de289c030856bb716",
    "transactionHash": "0x45215aa0da9b7597d233d96b6f7c4ac311edaba77a99ecc6471c59663554914f",
    "transactionIndex": "0x1"
  }`

	{
		receipt := eth.TransactionReceipt{}
		err := json.Unmarshal([]byte(old), &receipt)
		require.NoError(t, err, "unmarshal must succeed")

		require.Equal(t, eth.MustData32("0x57bd5108d8f0b8bad735ab77e2a47b80c166dcf5059b2960e0118b40562c7cf2"), receipt.Root)
		require.Nil(t, receipt.Status)
		require.Nil(t, receipt.Type)
		require.Equal(t, eth.TransactionTypeLegacy, receipt.TransactionType())
	}

	// EIP-2718 receipts
	legacy := `{
    "type": "0x0",
    "blockHash": "0xa37f46c4692db33012c105a27b9e4c582e822ed60a54667875fb92def52fd75a",
    "blockNumber": "0x72991c",
    "contractAddress": null,
    "cumulativeGasUsed": "0x7650c2",
    "from": "0x9e44b7d42125b7bb4e809406ed5e1079ff500969",
    "gasUsed": "0x5630",
    "logs": [
      {
        "address": "0x21ab6c9fac80c59d401b37cb43f81ea9dde7fe34",
        "blockHash": "0xa37f46c4692db33012c105a27b9e4c582e822ed60a54667875fb92def52fd75a",
        "blockNumber": "0x72991c",
        "data": "0x000000000000000000000000000000000000000000000000000000070560c8c0",
        "logIndex": "0xb7",
        "removed": false,
        "topics": [
          "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
          "0x0000000000000000000000009e44b7d42125b7bb4e809406ed5e1079ff500969",
          "0x000000000000000000000000fe5854255eb1eb921525fa856a3947ed2412a1d7"
        ],
        "transactionHash": "0x9d2fb08850a9b38173044ae6a61974fde4eacca504e399ffd9d5c8af567113cc",
        "transactionIndex": "0x8a"
      }
    ],
    "logsBloom": "0x00000001000000000000000000000000000000000000000000008000000000000000000000010000000000000000000000400000000000000000000000000008000000000000000000000008000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000200000000002000000000000000000000000000000000000000000000000000000000000000008000000000000000000000010000000000000000000000000000000",
    "status": "0x1",
    "to": "0x21ab6c9fac80c59d401b37cb43f81ea9dde7fe34",
    "transactionHash": "0x9d2fb08850a9b38173044ae6a61974fde4eacca504e399ffd9d5c8af567113cc",
    "transactionIndex": "0x8a"
  }`

	{
		receipt := eth.TransactionReceipt{}
		err := json.Unmarshal([]byte(legacy), &receipt)
		require.NoError(t, err, "unmarshal must succeed")

		require.NotNil(t, receipt.Type)
		require.Equal(t, eth.TransactionTypeLegacy, receipt.Type.Int64())
		require.Equal(t, eth.TransactionTypeLegacy, receipt.TransactionType())
	}

	eip2930 := `{
		"type": "0x1",
		"blockHash": "0xc6b65d9a251257942744ba1f250df218c2db4c1ec91d54d505034af5029f5edc",
		"blockNumber": "0x45",
		"contractAddress": null,
		"cumulativeGasUsed": "0x62d4",
		"from": "0x8a8eafb1cf62bfbeb1741769dae1a9dd47996192",
		"gasUsed": "0x62d4",
		"logs": [],
		"logsBloom": "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		"status": "0x1",
		"to": "0x8a8eafb1cf62bfbeb1741769dae1a9dd47996192",
		"transactionHash": "0x5cb00d928abf074cb81fc5e54dd49ef541afa7fc014b8a53fb8c29f3ecb5cadb",
		"transactionIndex": "0x0"
	  }`

	{
		receipt := eth.TransactionReceipt{}
		err := json.Unmarshal([]byte(eip2930), &receipt)
		require.NoError(t, err, "unmarshal must succeed")

		require.NotNil(t, receipt.Type)
		require.Equal(t, eth.TransactionTypeAccessList, receipt.Type.Int64())
		require.Equal(t, eth.TransactionTypeAccessList, receipt.TransactionType())
	}
}

func TestTransactionReceipt_4844(t *testing.T) {
	// curl https://rpc.dencun-devnet-8.ethpandaops.io/ -H 'Content-Type: application/json' -d '{"method":"eth_getTransactionReceipt","params":["0x5ceec39b631763ae0b45a8fb55c373f38b8fab308336ca1dc90ecd2b3cf06d00"],"id":1,"jsonrpc":"2.0"}'
	raw := `{
		"blobGasPrice": "0x1",
		"blobGasUsed": "0x20000",
		"blockHash": "0xfc2715ff196e23ae613ed6f837abd9035329a720a1f4e8dce3b0694c867ba052",
		"blockNumber": "0x2a1cb",
		"contractAddress": null,
		"cumulativeGasUsed": "0x5208",
		"effectiveGasPrice": "0x1d1a94a201c",
		"from": "0xad01b55d7c3448b8899862eb335fbb17075d8de2",
		"gasUsed": "0x5208",
		"logs": [],
		"logsBloom": "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		"status": "0x1",
		"to": "0x000000000000000000000000000000000000f1c1",
		"transactionHash": "0x5ceec39b631763ae0b45a8fb55c373f38b8fab308336ca1dc90ecd2b3cf06d00",
		"transactionIndex": "0x0",
		"type": "0x3"
	  }`

	receipt := eth.TransactionReceipt{}
	err := json.Unmarshal([]byte(raw), &receipt)
	require.NoError(t, err, "unmarshal must succeed")

	require.NotNil(t, receipt.Type)
	require.Equal(t, "0x3", receipt.Type.String())
	require.Equal(t, "0x1", receipt.BlobGasPrice.String())
	require.Equal(t, "0x20000", receipt.BlobGasUsed.String())

	// convert back to JSON and compare
	b, err := json.Marshal(&receipt)
	require.NoError(t, err)
	require.JSONEq(t, raw, string(b))
}
