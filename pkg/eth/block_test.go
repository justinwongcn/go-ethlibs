package eth_test

import (
	"encoding/json"
	"github.com/INFURA/eth/pkg/eth"
	"github.com/stretchr/testify/require"
	"testing"
)

func RequireEqualJSON(t *testing.T, expected, actual []byte, msgAndArgs ...interface{}) {
	var exp interface{}
	var act interface{}

	var err error
	err = json.Unmarshal([]byte(expected), &exp)
	require.NoError(t, err, msgAndArgs...)
	err = json.Unmarshal([]byte(actual), &act)
	require.NoError(t, err, msgAndArgs...)
	require.Equal(t, exp, act, msgAndArgs...)
}

func TestMainnetGethBlocks(t *testing.T) {

	partial := `{
    "difficulty": "0xbfabcdbd93dda",
    "extraData": "0x737061726b706f6f6c2d636e2d6e6f64652d3132",
    "gasLimit": "0x79f39e",
    "gasUsed": "0x79ccd3",
    "hash": "0xb3b20624f8f0f86eb50dd04688409e5cea4bd02d700bf6e79e9384d47d6a5a35",
    "logsBloom": "0x4848112002a2020aaa0812180045840210020005281600c80104264300080008000491220144461026015300100000128005018401002090a824a4150015410020140400d808440106689b29d0280b1005200007480ca950b15b010908814e01911000054202a020b05880b914642a0000300003010044044082075290283516be82504082003008c4d8d14462a8800c2990c88002a030140180036c220205201860402001014040180002006860810ec0a1100a14144148408118608200060461821802c081000042d0810104a8004510020211c088200420822a082040e10104c00d010064004c122692020c408a1aa2348020445403814002c800888208b1",
    "miner": "0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c",
    "mixHash": "0x3d1fdd16f15aeab72e7db1013b9f034ee33641d92f71c0736beab4e67d34c7a7",
    "nonce": "0x4db7a1c01d8a8072",
    "number": "0x5bad55",
    "parentHash": "0x61a8ad530a8a43e3583f8ec163f773ad370329b2375d66433eb82f005e1d6202",
    "receiptsRoot": "0x5eced534b3d84d3d732ddbc714f5fd51d98a941b28182b6efe6df3a0fe90004b",
    "sha3Uncles": "0x8a562e7634774d3e3a36698ac4915e37fc84a2cd0044cb84fa5d80263d2af4f6",
    "size": "0x41c7",
    "stateRoot": "0xf5208fffa2ba5a3f3a2f64ebd5ca3d098978bedd75f335f56b705d8715ee2305",
    "timestamp": "0x5b541449",
    "totalDifficulty": "0x12ac11391a2f3872fcd",
    "transactions": [
      "0x8784d99762bccd03b2086eabccee0d77f14d05463281e121a62abfebcf0d2d5f",
      "0x311be6a9b58748717ac0f70eb801d29973661aaf1365960d159e4ec4f4aa2d7f",
      "0xe42b0256058b7cad8a14b136a0364acda0b4c36f5b02dea7e69bfd82cef252a2"
    ],
    "transactionsRoot": "0xf98631e290e88f58a46b7032f025969039aa9b5696498efc76baf436fa69b262",
    "uncles": [
      "0x824cce7c7c2ec6874b9fa9a9a898eb5f27cbaf3991dfa81084c3af60d1db618c"
    ]
  }`

	var block eth.Block

	err := json.Unmarshal([]byte(partial), &block)
	require.NoError(t, err, "mainnet partial block should deserialize")

	require.Equal(t, 3, len(block.Transactions))
	require.Equal(t, "0x5bad55", block.Number.String())
	require.Equal(t, uint64(0x79f39e), block.GasLimit.UInt64())

	full := `{
    "difficulty": "0x742a575f662",
    "extraData": "0xd783010302844765746887676f312e352e31856c696e7578",
    "gasLimit": "0x2fefd8",
    "gasUsed": "0x5208",
    "hash": "0x648509915efa19b169ccab758492c7525b8498747678b894befd9ff78ad05519",
    "logsBloom": "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
    "miner": "0x2a65aca4d5fc5b5c859090a6c34d164135398226",
    "mixHash": "0x47e7eab7d034cf4b8b1501ebfc98edf715ee62f56283bf1a22a5423990600dff",
    "nonce": "0xeacef1c5a2ca3a49",
    "number": "0x99999",
    "parentHash": "0xffa241fbb914038a429c90daeeb54885f31e431d05b12fe87de8007853a1f278",
    "receiptsRoot": "0xb46f767bd3f69c0d7830eae6717f77560ee2ace0ea701d9e95fd41eb39a619ab",
    "sha3Uncles": "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
    "size": "0x290",
    "stateRoot": "0x93e74cf453c3327075b7e252deeb2d115cf2fdb204ba89806cebbd32afdedaa8",
    "timestamp": "0x565eafba",
    "totalDifficulty": "0x336f973a0249a1e9",
    "transactions": [
      {
        "blockHash": "0x648509915efa19b169ccab758492c7525b8498747678b894befd9ff78ad05519",
        "blockNumber": "0x99999",
        "from": "0x4bb96091ee9d802ed039c4d1a5f6216f90f81b01",
        "gas": "0xa028",
        "gasPrice": "0xba43b7400",
        "hash": "0xb4c724bf1f01a5371c513389d5758d531b729f15c8c6af8f74a100585d2cf33f",
        "input": "0x",
        "nonce": "0x461e",
        "r": "0xd5ee485b95d5992a4ca7d210ff28d540aea3f4031ce39203298ae266bcdb3485",
        "s": "0x71ecb17bdbbae8c57681649a95e8c7e22b90adac2e19c314de3b74ecfb5f8ce1",
        "to": "0x86d3856ad0105b9d4199936c1fd203664ba325dc",
        "transactionIndex": "0x0",
        "v": "0x1b",
        "value": "0x44b1eec6162f0000"
      }
    ],
    "transactionsRoot": "0x237e46a0a93850f7979546c717ffccce6715a6b2cb0bdb0d59a9c559a0d74f07",
    "uncles": []
  }`

	err = json.Unmarshal([]byte(full), &block)
	require.NoError(t, err, "mainnet full block should deserialize")

	require.Equal(t, 1, len(block.Transactions))
	require.Equal(t, "0x99999", block.Number.String())
	require.Equal(t, block.Hash, block.Transactions[0].BlockHash)
	require.Equal(t, eth.Data("0xd783010302844765746887676f312e352e31856c696e7578"), block.ExtraData)
	require.Equal(t, true, block.Transactions[0].Populated)
	require.Equal(t, int64(0), block.Transactions[0].Index.Int64())
	require.Equal(t, eth.Data("0x"), block.Transactions[0].Input)

	j, err := json.Marshal(&block)
	require.NoError(t, err)

	RequireEqualJSON(t, []byte(full), j)
}

func TestMainnetParityBlocks(t *testing.T) {
	full := `{
    "author": "0x2a65aca4d5fc5b5c859090a6c34d164135398226",
    "difficulty": "0x742a575f662",
    "extraData": "0xd783010302844765746887676f312e352e31856c696e7578",
    "gasLimit": "0x2fefd8",
    "gasUsed": "0x5208",
    "hash": "0x648509915efa19b169ccab758492c7525b8498747678b894befd9ff78ad05519",
    "logsBloom": "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
    "miner": "0x2a65aca4d5fc5b5c859090a6c34d164135398226",
    "mixHash": "0x47e7eab7d034cf4b8b1501ebfc98edf715ee62f56283bf1a22a5423990600dff",
    "nonce": "0xeacef1c5a2ca3a49",
    "number": "0x99999",
    "parentHash": "0xffa241fbb914038a429c90daeeb54885f31e431d05b12fe87de8007853a1f278",
    "receiptsRoot": "0xb46f767bd3f69c0d7830eae6717f77560ee2ace0ea701d9e95fd41eb39a619ab",
    "sealFields": [
      "0xa047e7eab7d034cf4b8b1501ebfc98edf715ee62f56283bf1a22a5423990600dff",
      "0x88eacef1c5a2ca3a49"
    ],
    "sha3Uncles": "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
    "size": "0x290",
    "stateRoot": "0x93e74cf453c3327075b7e252deeb2d115cf2fdb204ba89806cebbd32afdedaa8",
    "timestamp": "0x565eafba",
    "totalDifficulty": "0x336f973a0249a1e9",
    "transactions": [
      {
        "blockHash": "0x648509915efa19b169ccab758492c7525b8498747678b894befd9ff78ad05519",
        "blockNumber": "0x99999",
        "chainId": null,
        "condition": null,
        "creates": null,
        "from": "0x4bb96091ee9d802ed039c4d1a5f6216f90f81b01",
        "gas": "0xa028",
        "gasPrice": "0xba43b7400",
        "hash": "0xb4c724bf1f01a5371c513389d5758d531b729f15c8c6af8f74a100585d2cf33f",
        "input": "0x",
        "nonce": "0x461e",
        "publicKey": "0xa9177f27b99a4ad938359d77e0dca4b64e7ce3722c835d8087d4eecb27c8a54d59e2917e6b31ec12e44b1064d102d35815f9707af9571f15e92d1b6fbcd207e9",
        "r": "0xd5ee485b95d5992a4ca7d210ff28d540aea3f4031ce39203298ae266bcdb3485",
        "raw": "0xf86e82461e850ba43b740082a0289486d3856ad0105b9d4199936c1fd203664ba325dc8844b1eec6162f0000801ba0d5ee485b95d5992a4ca7d210ff28d540aea3f4031ce39203298ae266bcdb3485a071ecb17bdbbae8c57681649a95e8c7e22b90adac2e19c314de3b74ecfb5f8ce1",
        "s": "0x71ecb17bdbbae8c57681649a95e8c7e22b90adac2e19c314de3b74ecfb5f8ce1",
        "standardV": "0x0",
        "to": "0x86d3856ad0105b9d4199936c1fd203664ba325dc",
        "transactionIndex": "0x0",
        "v": "0x1b",
        "value": "0x44b1eec6162f0000"
      }
    ],
    "transactionsRoot": "0x237e46a0a93850f7979546c717ffccce6715a6b2cb0bdb0d59a9c559a0d74f07",
    "uncles": []
  }`

	var block eth.Block
	err := json.Unmarshal([]byte(full), &block)
	require.NoError(t, err, "mainnet full block should deserialize")

	require.Equal(t, 1, len(block.Transactions))
	require.Equal(t, "0x99999", block.Number.String())
	require.Equal(t, block.Hash, block.Transactions[0].BlockHash)
	require.Equal(t, eth.Data("0xd783010302844765746887676f312e352e31856c696e7578"), block.ExtraData)
	require.Equal(t, true, block.Transactions[0].Populated)
	require.Equal(t, int64(0), block.Transactions[0].Index.Int64())
	require.Equal(t, eth.Data("0x"), block.Transactions[0].Input)

	j, err := json.Marshal(&block)
	require.NoError(t, err)

	RequireEqualJSON(t, []byte(full), j)
}

func TestRinkebyGethBlocks(t *testing.T) {
	partial := `{
  "difficulty": "0x2",
  "extraData": "0xd883010900846765746888676f312e31312e31856c696e75780000000000000066c9aee037b7614e79d9abfc4422ced94a6ca6310e8f17be4a491ac9aaa1ebbd16e9996ad19244964781f806580459829876ab49382e8c1ed37444f17cbaffcc00",
  "gasLimit": "0x6ac95b",
  "gasUsed": "0x2b3ddc",
  "hash": "0x85a52d6d985f5c3221d2df50313950571e4772be23afca8ce410f65d4fbae66a",
  "logsBloom": "0x000c000020080000000000862100000020180008000008080000004921001000001010020014020001000000000200204000000128002000201000001000000004c00088004000810000000800002010000808001000000000000000446000010000000002000800200000000002080000000280080000000000009000000020100010c000008000000084000201000020800001044004000100280000000000020000000040000010000000400000200002000800000000000000000020011000400022000008000010814002040000080000050000200000000000080030000400002000000000820042000002800200082000100081001080080000000080",
  "miner": "0x0000000000000000000000000000000000000000",
  "mixHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "nonce": "0x0000000000000000",
  "number": "0x3ae0c1",
  "parentHash": "0x9b6a1fb3d50c9dd6ffe870d7828220aaede60e18aa0bc418bd3a68e06dda6ff2",
  "receiptsRoot": "0xc21fa0b7d69cc9570308066fed4cb13800d2c0799c292e3365d1824c2460d63e",
  "sha3Uncles": "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
  "size": "0x1c94",
  "stateRoot": "0x480f9a5649b844a69e75c545283cca9443cefaf0365aacd3bb4da4fb5061da97",
  "timestamp": "0x5c634043",
  "totalDifficulty": "0x6cb2df",
  "transactions": [
    "0x9490b0ccdb311bee95336ba35ce13bdbe853968672920d2e855be8467ca04484",
    "0x04871c1493f799184df5bf39f2f25b3eddb59bf91149e54f635859ffb587cdce",
    "0x79e78799b19b27dc9efada94a478728c9045f7dd239fa70f5f2f6e9ec0f53dd9",
    "0x607586f61f92dc3fc792d901a246079b78b7dd006f6e5633701ec7fd530c4e10",
    "0x0e4d4d6963f3ce12224ecdbe2f3018d3ebdda015ec65b1f74b4aea397107478b",
    "0xe5da1d2e03cdaed87fe7aa50f3456b6d230abda06b039f247b1f4d52927df4f8",
    "0xd9f393d9dbb2162db038fd84ba46155e1aa9e397dd62b955375119435aa924fe",
    "0x01a918ea5b384dd45499bd99cc5c08ec52559c3ccc12c2fbf0c1336490c78a35",
    "0xa7c78a933a9f8b9fdb48c11f53a655a7b9b1bbd14140b52b364fcabde1e1aaff",
    "0x5dbba0cd86cc68f7abeaae4dfad6ffc4ed4b4ce3e816129c938dc777c7a54cca",
    "0x20057225e2f70084397c60b34a1ee7840a6ca7d3ad363702f3df5156b3b5e051",
    "0xb037f13701cebde1e0e8e2624a1752f2f8384ca71ffe4cb591c371b715706cb5",
    "0xa6748e4441ed7dcd689c90521f2447749265c11e2be3660370f4cc6fc573ce47",
    "0x6180fbe39d0ca0044c3c4c92a62f5d70aae26afbcc4b9e494ae521b293f62a31",
    "0x7916f314e97f79217ca28eb79c573c185cf0540fe168eae16d8117ffb4e63c4a"
  ],
  "transactionsRoot": "0x8b0849114e5928a281a6a7e1d760db1fd829fe2cfb9bb98eab3fc1e386af2a9a",
  "uncles": []
}`

	var block eth.Block
	err := json.Unmarshal([]byte(partial), &block)
	require.NoError(t, err, "rinkeby partial block should deserialize")

	require.Equal(t, 15, len(block.Transactions))
	require.Equal(t, eth.MustQuantity("0x3ae0c1"), block.Number)
	require.Equal(t, uint64(0x2b3ddc), block.GasUsed.UInt64())

	j, err := json.Marshal(&block)
	require.NoError(t, err)

	RequireEqualJSON(t, []byte(partial), j)
}

func TestKovanParityBlocks(t *testing.T) {
	full := `{
    "author": "0x007733a1fe69cf3f2cf989f81c7b4cac1693387a",
    "difficulty": "0xfffffffffffffffffffffffffffffffe",
    "extraData": "0xde830203028f5061726974792d457468657265756d86312e33312e31826c69",
    "gasLimit": "0x7a1200",
    "gasUsed": "0x2276e2",
    "hash": "0x0c58244f5d538e1f5840c6751c3b3a2c9fdf9583245918b224c6aad088de6e5a",
    "logsBloom": "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
    "miner": "0x007733a1fe69cf3f2cf989f81c7b4cac1693387a",
    "number": "0x9de762",
    "parentHash": "0x7a3cce2e8ad99c92171adc9437d41eddd9e425f33b55b850528b6a9901ef0f9e",
    "receiptsRoot": "0x80add7251de396e7f6dd684e4d1be6823bdb867e96a9586c65a6f6c03659fad0",
    "sealFields": [
      "0x841718d07d",
      "0xb841e3d53556d6af10a059dd0d0d4bad2b10e3389ae08c83318b053cf04fac5c69676e96bee369cb3c11a926bbaeb25aa3ce80a101bd2d9db6f6f014b3ccbbfd7c4e00"
    ],
    "sha3Uncles": "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
    "signature": "e3d53556d6af10a059dd0d0d4bad2b10e3389ae08c83318b053cf04fac5c69676e96bee369cb3c11a926bbaeb25aa3ce80a101bd2d9db6f6f014b3ccbbfd7c4e00",
    "size": "0x2bb9",
    "stateRoot": "0x4d301d25429a2114965ce829e2c8b96b69ab826aca09aafa7fe71428d5545651",
    "step": "387502205",
    "timestamp": "0x5c6341f4",
    "totalDifficulty": "0x9ba45300000000000000000000000484a0394a",
    "transactions": [
      {
        "blockHash": "0x0c58244f5d538e1f5840c6751c3b3a2c9fdf9583245918b224c6aad088de6e5a",
        "blockNumber": "0x9de762",
        "chainId": "0x2a",
        "condition": null,
        "creates": null,
        "from": "0xd22abf44e2e2b3a9da3b84383c894f936925333c",
        "gas": "0x7a1200",
        "gasPrice": "0x3b9aca00",
        "hash": "0xd329c3a8ed5e59d649db90a1ff6c4fb632f2f5b925aedd1baa4d84b6f6b9e2f2",
        "input": "0x1a4813d700000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000120000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000001a000000000000000000000000000000000000000000000000000000000009de76000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002097035c1f04f1a223413cbe3132300000000000000000000000000000000000062cb0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002f69acd420fb6e6de34ec361ceddd000000000000000000000000000000000000a637000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001",
        "nonce": "0x2501b",
        "publicKey": "0xd1c9ae85381ad59eb19d55a0df58e5c0d3599d3a3766347526fd3902ad98146e89694746668c36d59378d72b4710e7f35389bd8c931390a4b908d0d84a1949bb",
        "r": "0x1686aacc780ad5a1805738bd87c709f3c0fbd5c8283be6f454aabefead5d5910",
        "raw": "0xf902cd8302501b843b9aca00837a1200945717adf502fd8830456bd5dc26801a4db394e6b280b902641a4813d700000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000120000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000001a000000000000000000000000000000000000000000000000000000000009de76000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002097035c1f04f1a223413cbe3132300000000000000000000000000000000000062cb0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002f69acd420fb6e6de34ec361ceddd000000000000000000000000000000000000a63700000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000177a01686aacc780ad5a1805738bd87c709f3c0fbd5c8283be6f454aabefead5d5910a002d0e9cee44425e63560ce16768d8a52552ae4b386a957bdbf5e518b438e2641",
        "s": "0x2d0e9cee44425e63560ce16768d8a52552ae4b386a957bdbf5e518b438e2641",
        "standardV": "0x0",
        "to": "0x5717adf502fd8830456bd5dc26801a4db394e6b2",
        "transactionIndex": "0x0",
        "v": "0x77",
        "value": "0x0"
      },
      {
        "blockHash": "0x0c58244f5d538e1f5840c6751c3b3a2c9fdf9583245918b224c6aad088de6e5a",
        "blockNumber": "0x9de762",
        "chainId": null,
        "condition": null,
        "creates": null,
        "from": "0xb3683b4de1fc502807464b55d151e8e2d2c19cb5",
        "gas": "0xf4240",
        "gasPrice": "0x3b9aca00",
        "hash": "0x9535a683436191e42f5a4371c1c3fe6be07964e858f52d876d9e5c06e04233f5",
        "input": "0x28fbdf0d000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000040376634613461663166336262326163646366313062656339636463336337636639313061303030336365323739313866366565323964376436626130353964360000000000000000000000000000000000000000000000000000000000000015332c31353530303038383134393939306131362c320000000000000000000000",
        "nonce": "0x2187c",
        "publicKey": "0x829019ea9bc42bbc7a182173191216d0386cbbf98cd35d2aa0126b8ccc8f94572dad6b218539e3a1692945ab905d34151ca8d1eea5c077104725ebff15277857",
        "r": "0xdede49403450d50a2db23a0f8d776fa94a6cae81ac8d803a29c47b4e6d63e3c",
        "raw": "0xf9014c8302187c843b9aca00830f424094cdbf1d1c64faad6d8484fd3dd5be2b4ea57f5f4c80b8e428fbdf0d000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000040376634613461663166336262326163646366313062656339636463336337636639313061303030336365323739313866366565323964376436626130353964360000000000000000000000000000000000000000000000000000000000000015332c31353530303038383134393939306131362c3200000000000000000000001ba00dede49403450d50a2db23a0f8d776fa94a6cae81ac8d803a29c47b4e6d63e3ca06361f099d66bdd992a4d85c9482423a9022165b40935d5623718c82c0d897f66",
        "s": "0x6361f099d66bdd992a4d85c9482423a9022165b40935d5623718c82c0d897f66",
        "standardV": "0x0",
        "to": "0xcdbf1d1c64faad6d8484fd3dd5be2b4ea57f5f4c",
        "transactionIndex": "0x1",
        "v": "0x1b",
        "value": "0x0"
      }
    ],
    "transactionsRoot": "0x6959042602c4847d263901a7d049b5565cbd511989a7d211b319fe2a96df4ed5",
    "uncles": []
  }`

	var block eth.Block
	err := json.Unmarshal([]byte(full), &block)
	require.NoError(t, err, "kovan full parity block should deserialize")

	require.Equal(t, "0x9de762", block.Number.String())
	require.Equal(t, 2, len(block.Transactions))
	require.Equal(t, uint64(0x2276e2), block.GasUsed.UInt64())

	j, err := json.Marshal(&block)
	require.NoError(t, err)

	RequireEqualJSON(t, []byte(full), j)
}

func TestGoerliBlocks(t *testing.T) {
	geth := `{
    "difficulty": "0x2",
    "extraData": "0x00000000000000000000000000000000000000000000000000000000000000003a69a356f8082954c675f1bb634c9ef5db6f1077b498a8c0a2a17afd0e7f8c5072ef2f3548bd013ae758a13b901f3b834c713c2391a2a6844c26a6ff386177a300",
    "gasLimit": "0x7a1200",
    "gasUsed": "0x146dd",
    "hash": "0x2b27fe2bbc8ce01ac7ae8bf74f793a197cf7edbe82727588811fa9a2c4776f81",
    "logsBloom": "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
    "miner": "0x0000000000000000000000000000000000000000",
    "mixHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
    "nonce": "0x0000000000000000",
    "number": "0x12b1d",
    "parentHash": "0x68393bd212eb3dec1b97f3a7e4e84a9cb0d632309276b485dbdf5cb383ab25cf",
    "receiptsRoot": "0x1ed4e3c3370adadd3eb3712fe87bb136e035ad2071d1a820de1ee1931ab677cf",
    "sha3Uncles": "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
    "size": "0x464",
    "stateRoot": "0x0a31f0562a092a79e00680cdbaf089b65d76e720c2a358901fa696a9cb5dd3d8",
    "timestamp": "0x5c64a101",
    "totalDifficulty": "0x20422",
    "transactions": [
      {
        "blockHash": "0x2b27fe2bbc8ce01ac7ae8bf74f793a197cf7edbe82727588811fa9a2c4776f81",
        "blockNumber": "0x12b1d",
        "from": "0x441a4060b5a1bf4ead6dd13acdcb7f83e4c374ca",
        "gas": "0x6d5f",
        "gasPrice": "0x3b9aca00",
        "hash": "0x8b142669ff0013c85898f471bc7cc01440b5c23aa3a0f2e1db8b2daa3e3b6974",
        "input": "0x9c0e3f7a0000000000000000000000000000000000000000000000068155a43676e00000000000000000000000000000000000000000000000000000000000000000002a",
        "nonce": "0x19",
        "to": "0x2d6a9044a88e8b6e175f31b71117dbb344c5892c",
        "transactionIndex": "0x0",
        "value": "0x0",
        "v": "0x1c",
        "r": "0x8b1d784d17627beb4c71269aa344ad1b49c2763f5062ec58e26dc3d1bd0e3fd3",
        "s": "0x59756df23b630a12f6e6d262b44b16c4170f35a7ac802df7419bb48045495275"
      },
      {
        "blockHash": "0x2b27fe2bbc8ce01ac7ae8bf74f793a197cf7edbe82727588811fa9a2c4776f81",
        "blockNumber": "0x12b1d",
        "from": "0xddff50398536a8ba7f3840581e662f1c9bd21505",
        "gas": "0x6d9f",
        "gasPrice": "0x3b9aca00",
        "hash": "0xe25fdf96c393dbd853dfb8a421b54c6bab645ac5037a3d90b679ebe7a3028d89",
        "input": "0x9c0e3f7a00000000000000000000000000000000000000000000000004d498fb9e757a68000000000000000000000000000000000000000000000000000000000000002a",
        "nonce": "0x19",
        "to": "0x4e10a95f0bb2fec6ec1c4296a16420a018a5f9fe",
        "transactionIndex": "0x1",
        "value": "0x0",
        "v": "0x1b",
        "r": "0xd12898a69e49e3c0b97e28db161b97f5411169ee7934e51eec410a2b25d6bf14",
        "s": "0x809983bda5c292d4eb5434bf82c6cd00b5c9d615dd097c8c959cf35c7d49b7a"
      },
      {
        "blockHash": "0x2b27fe2bbc8ce01ac7ae8bf74f793a197cf7edbe82727588811fa9a2c4776f81",
        "blockNumber": "0x12b1d",
        "from": "0x2b371c0262ceab27face32fbb5270ddc6aa01ba4",
        "gas": "0x6bdf",
        "gasPrice": "0x3b9aca00",
        "hash": "0xbddbb685774d8a3df036ed9fb920b48f876090a57e9e90ee60921e0510ef7090",
        "input": "0x9c0e3f7a0000000000000000000000000000000000000000000000000000000000000078000000000000000000000000000000000000000000000000000000000000002a",
        "nonce": "0x1c",
        "to": "0x8e730df7c70d33118d9e5f79ab81aed0be6f6635",
        "transactionIndex": "0x2",
        "value": "0x0",
        "v": "0x1b",
        "r": "0x2a98c51c2782f664d3ce571fef0491b48f5ebbc5845fa513192e6e6b24ecdaa1",
        "s": "0x29b8e0c67aa9c11327e16556c591dc84a7aac2f6fc57c7f93901be8ee867aebc"
      }
    ],
    "transactionsRoot": "0xa0548c58d46cc59673308c31437c9b3bd718a1f44236b208eb3bc7991ff19040",
    "uncles": []
  }`

	parity := `{
    "author": "0x0000000000000000000000000000000000000000",
    "difficulty": "0x2",
    "extraData": "0x00000000000000000000000000000000000000000000000000000000000000003a69a356f8082954c675f1bb634c9ef5db6f1077b498a8c0a2a17afd0e7f8c5072ef2f3548bd013ae758a13b901f3b834c713c2391a2a6844c26a6ff386177a300",
    "gasLimit": "0x7a1200",
    "gasUsed": "0x146dd",
    "hash": "0x2b27fe2bbc8ce01ac7ae8bf74f793a197cf7edbe82727588811fa9a2c4776f81",
    "logsBloom": "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
    "miner": "0x0000000000000000000000000000000000000000",
    "number": "0x12b1d",
    "parentHash": "0x68393bd212eb3dec1b97f3a7e4e84a9cb0d632309276b485dbdf5cb383ab25cf",
    "receiptsRoot": "0x1ed4e3c3370adadd3eb3712fe87bb136e035ad2071d1a820de1ee1931ab677cf",
    "sealFields": [
      "0xa00000000000000000000000000000000000000000000000000000000000000000",
      "0x880000000000000000"
    ],
    "sha3Uncles": "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
    "size": "0x464",
    "stateRoot": "0x0a31f0562a092a79e00680cdbaf089b65d76e720c2a358901fa696a9cb5dd3d8",
    "timestamp": "0x5c64a101",
    "totalDifficulty": "0x20422",
    "transactions": [
      {
        "blockHash": "0x2b27fe2bbc8ce01ac7ae8bf74f793a197cf7edbe82727588811fa9a2c4776f81",
        "blockNumber": "0x12b1d",
        "chainId": null,
        "condition": null,
        "creates": null,
        "from": "0x441a4060b5a1bf4ead6dd13acdcb7f83e4c374ca",
        "gas": "0x6d5f",
        "gasPrice": "0x3b9aca00",
        "hash": "0x8b142669ff0013c85898f471bc7cc01440b5c23aa3a0f2e1db8b2daa3e3b6974",
        "input": "0x9c0e3f7a0000000000000000000000000000000000000000000000068155a43676e00000000000000000000000000000000000000000000000000000000000000000002a",
        "nonce": "0x19",
        "publicKey": "0xc2423312402f1b4a75ef71dc1a2f43700c22c663e799e13776c81d9a18ad332504dd6668088edb02f03968025fcb25d90e164f184dc00fd43649307ee5f23309",
        "r": "0x8b1d784d17627beb4c71269aa344ad1b49c2763f5062ec58e26dc3d1bd0e3fd3",
        "raw": "0xf8a819843b9aca00826d5f942d6a9044a88e8b6e175f31b71117dbb344c5892c80b8449c0e3f7a0000000000000000000000000000000000000000000000068155a43676e00000000000000000000000000000000000000000000000000000000000000000002a1ca08b1d784d17627beb4c71269aa344ad1b49c2763f5062ec58e26dc3d1bd0e3fd3a059756df23b630a12f6e6d262b44b16c4170f35a7ac802df7419bb48045495275",
        "s": "0x59756df23b630a12f6e6d262b44b16c4170f35a7ac802df7419bb48045495275",
        "standardV": "0x1",
        "to": "0x2d6a9044a88e8b6e175f31b71117dbb344c5892c",
        "transactionIndex": "0x0",
        "v": "0x1c",
        "value": "0x0"
      },
      {
        "blockHash": "0x2b27fe2bbc8ce01ac7ae8bf74f793a197cf7edbe82727588811fa9a2c4776f81",
        "blockNumber": "0x12b1d",
        "chainId": null,
        "condition": null,
        "creates": null,
        "from": "0xddff50398536a8ba7f3840581e662f1c9bd21505",
        "gas": "0x6d9f",
        "gasPrice": "0x3b9aca00",
        "hash": "0xe25fdf96c393dbd853dfb8a421b54c6bab645ac5037a3d90b679ebe7a3028d89",
        "input": "0x9c0e3f7a00000000000000000000000000000000000000000000000004d498fb9e757a68000000000000000000000000000000000000000000000000000000000000002a",
        "nonce": "0x19",
        "publicKey": "0xc699e4a9b97ca35f973b2cf00760f895eb1d8b0aadbf65a3cf40ae39ad4481a50cd45bbd543dbbfba831dc55b3a16337e805e830045ff70551d4ea9f77f89efc",
        "r": "0xd12898a69e49e3c0b97e28db161b97f5411169ee7934e51eec410a2b25d6bf14",
        "raw": "0xf8a819843b9aca00826d9f944e10a95f0bb2fec6ec1c4296a16420a018a5f9fe80b8449c0e3f7a00000000000000000000000000000000000000000000000004d498fb9e757a68000000000000000000000000000000000000000000000000000000000000002a1ba0d12898a69e49e3c0b97e28db161b97f5411169ee7934e51eec410a2b25d6bf14a00809983bda5c292d4eb5434bf82c6cd00b5c9d615dd097c8c959cf35c7d49b7a",
        "s": "0x809983bda5c292d4eb5434bf82c6cd00b5c9d615dd097c8c959cf35c7d49b7a",
        "standardV": "0x0",
        "to": "0x4e10a95f0bb2fec6ec1c4296a16420a018a5f9fe",
        "transactionIndex": "0x1",
        "v": "0x1b",
        "value": "0x0"
      },
      {
        "blockHash": "0x2b27fe2bbc8ce01ac7ae8bf74f793a197cf7edbe82727588811fa9a2c4776f81",
        "blockNumber": "0x12b1d",
        "chainId": null,
        "condition": null,
        "creates": null,
        "from": "0x2b371c0262ceab27face32fbb5270ddc6aa01ba4",
        "gas": "0x6bdf",
        "gasPrice": "0x3b9aca00",
        "hash": "0xbddbb685774d8a3df036ed9fb920b48f876090a57e9e90ee60921e0510ef7090",
        "input": "0x9c0e3f7a0000000000000000000000000000000000000000000000000000000000000078000000000000000000000000000000000000000000000000000000000000002a",
        "nonce": "0x1c",
        "publicKey": "0xc4e6c6c19ab79c1dcc7bbd2c76f52cb110ac90ee56728e588cda4b6eb763848203ce9e3a12322a9128b0c3572f99f5786d6569946785b1d1b0b97ade516cad03",
        "r": "0x2a98c51c2782f664d3ce571fef0491b48f5ebbc5845fa513192e6e6b24ecdaa1",
        "raw": "0xf8a81c843b9aca00826bdf948e730df7c70d33118d9e5f79ab81aed0be6f663580b8449c0e3f7a0000000000000000000000000000000000000000000000000000000000000078000000000000000000000000000000000000000000000000000000000000002a1ba02a98c51c2782f664d3ce571fef0491b48f5ebbc5845fa513192e6e6b24ecdaa1a029b8e0c67aa9c11327e16556c591dc84a7aac2f6fc57c7f93901be8ee867aebc",
        "s": "0x29b8e0c67aa9c11327e16556c591dc84a7aac2f6fc57c7f93901be8ee867aebc",
        "standardV": "0x0",
        "to": "0x8e730df7c70d33118d9e5f79ab81aed0be6f6635",
        "transactionIndex": "0x2",
        "v": "0x1b",
        "value": "0x0"
      }
    ],
    "transactionsRoot": "0xa0548c58d46cc59673308c31437c9b3bd718a1f44236b208eb3bc7991ff19040",
    "uncles": []
  }`

	for i, payload := range []string{geth, parity} {
		msg := ""
		switch i {
		case 0:
			msg = "geth goerli block"
		case 1:
			msg = "parity goerli block"
		}
		var block eth.Block
		err := json.Unmarshal([]byte(payload), &block)
		require.NoError(t, err, msg)

		require.Equal(t, "0x12b1d", block.Number.String(), msg)
		require.Equal(t, 3, len(block.Transactions), msg)
		require.Equal(t, "0x59756df23b630a12f6e6d262b44b16c4170f35a7ac802df7419bb48045495275", block.Transactions[0].S.String(), msg)
		require.Equal(t, uint64(0x146dd), block.GasUsed.UInt64(), msg)

		j, err := json.Marshal(&block)
		require.NoError(t, err, msg)

		RequireEqualJSON(t, []byte(payload), j, msg)
	}
}
