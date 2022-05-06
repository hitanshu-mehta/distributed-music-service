run:
	go run main.go

sol:
	solc --optimize --overwrite --abi ./ledger/contracts/*.sol -o build
	solc --optimize --overwrite --bin ./ledger/contracts/*.sol -o build
	abigen --abi=./build/PublishSongContract.abi --bin=./build/PublishSongContract.bin --pkg=publishsongapi --out=./ledger/contracts/publishsongapi/PublishSongContract.go
	abigen --abi=./build/MusicToken.abi --bin=./build/MusicToken.bin --pkg=tokenapi --out=./ledger/contracts/tokenapi/MusicToken.go
	abigen --abi=./build/MusicTokenCrowdSale.abi --bin=./build/MusicTokenCrowdSale.bin --pkg=tokencrowdsaleapi --out=./ledger/contracts/tokencrowdsaleapi/MusicTokenCrowdSale.go