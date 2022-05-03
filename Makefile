run:
	go run main.go

sol:
	solc --optimize --overwrite --abi ./ledger/contracts/PublishSongContract.sol -o build
	solc --optimize --overwrite --bin ./ledger/contracts/PublishSongContract.sol -o build
	abigen --abi=./build/PublishSongContract.abi --bin=./build/PublishSongContract.bin --pkg=api --out=./ledger/contracts/api/PublishSongContract.go