// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L1CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_0_0_20\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initialized\",\"offset\":20,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1002,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initializing\",\"offset\":21,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1003,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_1_0_1600\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1004,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_51_0_20\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1005,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_52_0_1568\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1006,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_101_0_1\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_bool\"},{\"astId\":1007,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_102_0_1568\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1008,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_151_0_32\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_uint256\"},{\"astId\":1009,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_152_0_1568\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1010,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_201_0_32\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1011,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_202_0_32\",\"offset\":0,\"slot\":\"202\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1012,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"successfulMessages\",\"offset\":0,\"slot\":\"203\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1013,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"xDomainMsgSender\",\"offset\":0,\"slot\":\"204\",\"type\":\"t_address\"},{\"astId\":1014,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"205\",\"type\":\"t_uint240\"},{\"astId\":1015,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"failedMessages\",\"offset\":0,\"slot\":\"206\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1016,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"otherMessenger\",\"offset\":0,\"slot\":\"207\",\"type\":\"t_contract(CrossDomainMessenger)1021\"},{\"astId\":1017,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"208\",\"type\":\"t_array(t_uint256)43_storage\"},{\"astId\":1018,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"superchainConfig\",\"offset\":0,\"slot\":\"251\",\"type\":\"t_contract(SuperchainConfig)1023\"},{\"astId\":1019,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"portal\",\"offset\":0,\"slot\":\"252\",\"type\":\"t_contract(OptimismPortal)1022\"},{\"astId\":1020,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"systemConfig\",\"offset\":0,\"slot\":\"253\",\"type\":\"t_contract(SystemConfig)1024\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)43_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[43]\",\"numberOfBytes\":\"1376\",\"base\":\"t_uint256\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_contract(CrossDomainMessenger)1021\":{\"encoding\":\"inplace\",\"label\":\"contract CrossDomainMessenger\",\"numberOfBytes\":\"20\"},\"t_contract(OptimismPortal)1022\":{\"encoding\":\"inplace\",\"label\":\"contract OptimismPortal\",\"numberOfBytes\":\"20\"},\"t_contract(SuperchainConfig)1023\":{\"encoding\":\"inplace\",\"label\":\"contract SuperchainConfig\",\"numberOfBytes\":\"20\"},\"t_contract(SystemConfig)1024\":{\"encoding\":\"inplace\",\"label\":\"contract SystemConfig\",\"numberOfBytes\":\"20\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var L1CrossDomainMessengerDeployedBin = "0x60806040526004361061018b5760003560e01c80636425666b116100d6578063b1b1b2091161007f578063d764ad0b11610059578063d764ad0b1461049b578063db505d80146104ae578063ecc70428146104db57600080fd5b8063b1b1b2091461042b578063b28ade251461045b578063c0c53b8b1461047b57600080fd5b80638cbeeef2116100b05780638cbeeef2146102d05780639fce812c146103d0578063a4e7f8bd146103fb57600080fd5b80636425666b146103775780636e296e45146103a457806383a74074146103b957600080fd5b80633dbb202b1161013857806354fd4d501161011257806354fd4d50146102e65780635644cfdf1461033c5780635c975abb1461035257600080fd5b80633dbb202b146102935780633f827a5a146102a85780634c1d6a69146102d057600080fd5b80632828d7e8116101695780632828d7e81461022457806333d7e2bd1461023957806335e80ab31461026657600080fd5b8063028f85f7146101905780630c568498146101c35780630ff754ea146101d8575b600080fd5b34801561019c57600080fd5b506101a5601081565b60405167ffffffffffffffff90911681526020015b60405180910390f35b3480156101cf57600080fd5b506101a5603f81565b3480156101e457600080fd5b5060fc5473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101ba565b34801561023057600080fd5b506101a5604081565b34801561024557600080fd5b5060fd546101ff9073ffffffffffffffffffffffffffffffffffffffff1681565b34801561027257600080fd5b5060fb546101ff9073ffffffffffffffffffffffffffffffffffffffff1681565b6102a66102a1366004611bdf565b610540565b005b3480156102b457600080fd5b506102bd600181565b60405161ffff90911681526020016101ba565b3480156102dc57600080fd5b506101a5619c4081565b3480156102f257600080fd5b5061032f6040518060400160405280600581526020017f322e342e3000000000000000000000000000000000000000000000000000000081525081565b6040516101ba9190611cb1565b34801561034857600080fd5b506101a561138881565b34801561035e57600080fd5b5061036761083d565b60405190151581526020016101ba565b34801561038357600080fd5b5060fc546101ff9073ffffffffffffffffffffffffffffffffffffffff1681565b3480156103b057600080fd5b506101ff6108d6565b3480156103c557600080fd5b506101a562030d4081565b3480156103dc57600080fd5b5060cf5473ffffffffffffffffffffffffffffffffffffffff166101ff565b34801561040757600080fd5b50610367610416366004611ccb565b60ce6020526000908152604090205460ff1681565b34801561043757600080fd5b50610367610446366004611ccb565b60cb6020526000908152604090205460ff1681565b34801561046757600080fd5b506101a5610476366004611ce4565b6109bd565b34801561048757600080fd5b506102a6610496366004611d38565b610a2b565b6102a66104a9366004611d83565b610ca2565b3480156104ba57600080fd5b5060cf546101ff9073ffffffffffffffffffffffffffffffffffffffff1681565b3480156104e757600080fd5b5061053260cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b6040519081526020016101ba565b6105486115d3565b156105e05734156105e0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603d60248201527f43726f7373446f6d61696e4d657373656e6765723a2063616e6e6f742073656e60448201527f642076616c7565207769746820637573746f6d2067617320746f6b656e00000060648201526084015b60405180910390fd5b60cf546107129073ffffffffffffffffffffffffffffffffffffffff166106088585856109bd565b347fd764ad0b0000000000000000000000000000000000000000000000000000000061067460cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338a34898c8c6040516024016106909796959493929190611e52565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611612565b8373ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a33858561079760cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b866040516107a9959493929190611eb1565b60405180910390a260405134815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a2505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff0000000000000000000000000000000000000000000000000000000000009091161790555050565b60fb54604080517f5c975abb000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff1691635c975abb9160048083019260209291908290030181865afa1580156108ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108d19190611eff565b905090565b60cc5460009073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff2153016109a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f43726f7373446f6d61696e4d657373656e6765723a2078446f6d61696e4d657360448201527f7361676553656e646572206973206e6f7420736574000000000000000000000060648201526084016105d7565b5060cc5473ffffffffffffffffffffffffffffffffffffffff1690565b6000611388619c4080603f6109d9604063ffffffff8816611f50565b6109e39190611f80565b6109ee601088611f50565b6109fb9062030d40611fce565b610a059190611fce565b610a0f9190611fce565b610a199190611fce565b610a239190611fce565b949350505050565b6000547501000000000000000000000000000000000000000000900460ff1615808015610a76575060005460017401000000000000000000000000000000000000000090910460ff16105b80610aa85750303b158015610aa8575060005474010000000000000000000000000000000000000000900460ff166001145b610b34576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016105d7565b600080547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790558015610bba57600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff1675010000000000000000000000000000000000000000001790555b60fb805473ffffffffffffffffffffffffffffffffffffffff8087167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560fc805486841690831617905560fd805492851692909116919091179055610c397342000000000000000000000000000000000000076116ab565b8015610c9c57600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b610caa61083d565b15610d11576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f43726f7373446f6d61696e4d657373656e6765723a207061757365640000000060448201526064016105d7565b60f087901c60028110610dcc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604d60248201527f43726f7373446f6d61696e4d657373656e6765723a206f6e6c7920766572736960448201527f6f6e2030206f722031206d657373616765732061726520737570706f7274656460648201527f20617420746869732074696d6500000000000000000000000000000000000000608482015260a4016105d7565b8061ffff16600003610ec1576000610e1d878986868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508f92506117e7915050565b600081815260cb602052604090205490915060ff1615610ebf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f43726f7373446f6d61696e4d657373656e6765723a206c65676163792077697460448201527f6864726177616c20616c72656164792072656c6179656400000000000000000060648201526084016105d7565b505b6000610f07898989898989898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061180692505050565b9050610f11611829565b15610f4957853414610f2557610f25611ffa565b600081815260ce602052604090205460ff1615610f4457610f44611ffa565b61109b565b3415610ffd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605060248201527f43726f7373446f6d61696e4d657373656e6765723a2076616c7565206d75737460448201527f206265207a65726f20756e6c657373206d6573736167652069732066726f6d2060648201527f612073797374656d206164647265737300000000000000000000000000000000608482015260a4016105d7565b600081815260ce602052604090205460ff1661109b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520636160448201527f6e6e6f74206265207265706c617965640000000000000000000000000000000060648201526084016105d7565b6110a487611905565b15611157576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f43726f7373446f6d61696e4d657373656e6765723a2063616e6e6f742073656e60448201527f64206d65737361676520746f20626c6f636b65642073797374656d206164647260648201527f6573730000000000000000000000000000000000000000000000000000000000608482015260a4016105d7565b600081815260cb602052604090205460ff16156111f6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520686160448201527f7320616c7265616479206265656e2072656c617965640000000000000000000060648201526084016105d7565b61121785611208611388619c40611fce565b67ffffffffffffffff1661194b565b158061123d575060cc5473ffffffffffffffffffffffffffffffffffffffff1661dead14155b1561135657600081815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555182917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff320161134f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d6573736167650000000000000000000000000000000000000060648201526084016105d7565b50506115ae565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8a1617905560006113e788619c405a6113aa9190612029565b8988888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061196992505050565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead1790559050801561149d57600082815260cb602052604090205460ff161561143a5761143a611ffa565b600082815260cb602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a26115aa565b600082815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff32016115aa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d6573736167650000000000000000000000000000000000000060648201526084016105d7565b5050505b50505050505050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b6000806115de611983565b5073ffffffffffffffffffffffffffffffffffffffff1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee141592915050565b60fc546040517fe9e05c4200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063e9e05c42908490611673908890839089906000908990600401612040565b6000604051808303818588803b15801561168c57600080fd5b505af11580156116a0573d6000803e3d6000fd5b505050505050505050565b6000547501000000000000000000000000000000000000000000900460ff16611756576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105d7565b60cc5473ffffffffffffffffffffffffffffffffffffffff166117a05760cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead1790555b60cf80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60006117f585858585611a20565b805190602001209050949350505050565b6000611816878787878787611ab9565b8051906020012090509695505050505050565b60fc5460009073ffffffffffffffffffffffffffffffffffffffff16331480156108d1575060cf5460fc54604080517f9bf62d82000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff9384169390921691639bf62d82916004808201926020929091908290030181865afa1580156118c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118e99190612098565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b600073ffffffffffffffffffffffffffffffffffffffff8216301480611945575060fc5473ffffffffffffffffffffffffffffffffffffffff8381169116145b92915050565b600080603f83619c4001026040850201603f5a021015949350505050565b600080600080845160208601878a8af19695505050505050565b60fd54604080517f4397dfef0000000000000000000000000000000000000000000000000000000081528151600093849373ffffffffffffffffffffffffffffffffffffffff90911692634397dfef92600480830193928290030181865afa1580156119f3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a1791906120b5565b90939092509050565b606084848484604051602401611a3994939291906120f5565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fcbd4ece9000000000000000000000000000000000000000000000000000000001790529050949350505050565b6060868686868686604051602401611ad69695949392919061213f565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd764ad0b0000000000000000000000000000000000000000000000000000000017905290509695505050505050565b73ffffffffffffffffffffffffffffffffffffffff81168114611b7a57600080fd5b50565b60008083601f840112611b8f57600080fd5b50813567ffffffffffffffff811115611ba757600080fd5b602083019150836020828501011115611bbf57600080fd5b9250929050565b803563ffffffff81168114611bda57600080fd5b919050565b60008060008060608587031215611bf557600080fd5b8435611c0081611b58565b9350602085013567ffffffffffffffff811115611c1c57600080fd5b611c2887828801611b7d565b9094509250611c3b905060408601611bc6565b905092959194509250565b6000815180845260005b81811015611c6c57602081850181015186830182015201611c50565b81811115611c7e576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611cc46020830184611c46565b9392505050565b600060208284031215611cdd57600080fd5b5035919050565b600080600060408486031215611cf957600080fd5b833567ffffffffffffffff811115611d1057600080fd5b611d1c86828701611b7d565b9094509250611d2f905060208501611bc6565b90509250925092565b600080600060608486031215611d4d57600080fd5b8335611d5881611b58565b92506020840135611d6881611b58565b91506040840135611d7881611b58565b809150509250925092565b600080600080600080600060c0888a031215611d9e57600080fd5b873596506020880135611db081611b58565b95506040880135611dc081611b58565b9450606088013593506080880135925060a088013567ffffffffffffffff811115611dea57600080fd5b611df68a828b01611b7d565b989b979a50959850939692959293505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b878152600073ffffffffffffffffffffffffffffffffffffffff808916602084015280881660408401525085606083015263ffffffff8516608083015260c060a0830152611ea460c083018486611e09565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff86168152608060208201526000611ee1608083018688611e09565b905083604083015263ffffffff831660608301529695505050505050565b600060208284031215611f1157600080fd5b81518015158114611cc457600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff80831681851681830481118215151615611f7757611f77611f21565b02949350505050565b600067ffffffffffffffff80841680611fc2577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910492915050565b600067ffffffffffffffff808316818516808303821115611ff157611ff1611f21565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b60008282101561203b5761203b611f21565b500390565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015267ffffffffffffffff84166040820152821515606082015260a06080820152600061208d60a0830184611c46565b979650505050505050565b6000602082840312156120aa57600080fd5b8151611cc481611b58565b600080604083850312156120c857600080fd5b82516120d381611b58565b602084015190925060ff811681146120ea57600080fd5b809150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff80871683528086166020840152506080604083015261212e6080830185611c46565b905082606083015295945050505050565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a083015261218a60c0830184611c46565b9897505050505050505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L1CrossDomainMessengerStorageLayoutJSON), L1CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1CrossDomainMessenger"] = L1CrossDomainMessengerStorageLayout
	deployedBytecodes["L1CrossDomainMessenger"] = L1CrossDomainMessengerDeployedBin
	immutableReferences["L1CrossDomainMessenger"] = false
}
