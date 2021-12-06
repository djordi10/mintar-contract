.PHONY: clean
clean:
	rm -rf ./build
	rm -rf ./api/go

.PHONY: build
build:
	rm -rf ./build
	rm -rf ./api
	mkdir build
	mkdir api
	mkdir api/core
	mkdir api/sale
	npx hardhat compile
	solc @openzeppelin/=$(shell pwd)/node_modules/@openzeppelin/ contracts/*.sol --optimize --abi ./contracts/Core.sol -o build
	solc @openzeppelin/=$(shell pwd)/node_modules/@openzeppelin/ contracts/*.sol --optimize --bin ./contracts/Core.sol -o build
	abigen --abi=./build/contracts_Core_sol_MintarCore.abi --bin=./build/contracts_Core_sol_MintarCore.bin --pkg=core --out=./api/core/core.go
	abigen --abi=./build/contracts_Sale_sol_MintarSale.abi --bin=./build/contracts_Sale_sol_MintarSale.bin --pkg=sale --out=./api/sale/sale.go
