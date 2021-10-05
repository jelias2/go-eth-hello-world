# Notes


##### To access solidity contract from golang
1. Get the contract source code in solidity
1. Generate the ABI ( application binary interface) using abigen
  `solc --abi Store.sol -o build`
1. Convert the abi into an importable go file
  `abigen --abi=./build/Store.abi --pkg=store --out=Store.go`
 
##### To Deploy solidity from golang
 1. In order to deploy a smart contract from Go, we also need to compile the solidity smart contract to EVM bytecode. The EVM bytecode is what will be sent in the data field of the transaction. The bin file is required for generating the deploy methods on the Go contract file.
 `solc --bin Store.sol -o build`
 1. Now we compile the Go contract file which will include the deploy methods because we includes the bin file.
 `abigen --bin=./build/Store.bin --abi=./build/Store.abi --pkg=store --out=Store.go` 
