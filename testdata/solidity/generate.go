//go:generate solc --abi ./demo/Demo.sol -o ./demo --overwrite
//go:generate solc --bin ./demo/Demo.sol -o ./demo --overwrite
//go:generate solc --asm ./demo/Demo.sol -o ./demo --overwrite
//go:generate abigen --bin ./demo/Demo.bin --abi ./demo/Demo.abi --pkg demo --out ./demo/demo.go
package solidity
