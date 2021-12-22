GO_PATH=`go env GOPATH`

init:
	cd ../
	git clone git@github.com:pashifika/fyne.git
	git checkout origin/release/fix_v2.1.2 -b release/fix_v2.1.2
	cd ../raku-redmine/
	go mod download

go-test:
	go test ./...

release-win:
	fyne package --release --tags=release --target=windows

release-mac:
	fyne package --release --tags=release --target=darwin

## show help
help:
	@make2help $(MAKEFILE_LIST)
