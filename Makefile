GO_PATH=`go env GOPATH`

init:
	cd ../
	git clone git@github.com:pashifika/fyne.git
	cd fyne
	git checkout origin/release/fix_v2.1.2 -b release/fix_v2.1.2
	cd ../raku-redmine/
	go mod download
	cp ./config.go ./config_dev.go
	cp ./utils/database/config.go ./utils/database/config_dev.go
	echo "Please edit the:"
	echo "    ./utils/database/config_dev.go"
	echo "    ./utils/database/config_dev.go"
	echo "  file to you local path to debug"

go-test:
	go test ./...

release-win:
	fyne package --release --tags=release --target=windows

release-mac:
	fyne package --release --tags=release --target=darwin

release-mac-x64:
	GOARCH=amd64 fyne package --release --tags=release --target=darwin

## show help
help:
	@make2help $(MAKEFILE_LIST)
