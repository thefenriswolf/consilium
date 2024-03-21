##
# consilium - schedule planner
#
# @file
# @version 0.1

.PHONY: build-release build-debug clean test benchmark lint

build-release: linux windows

run: build-release
#	nixGL ./bin/cons_linux
	wine64 bin/consilium_win_amd64.exe

linux:
	GOOS=linux GOARCH=amd64 go build -race -o ./bin/consilium_linux_amd64

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags='-H windowsgui' -o ./bin/consilium_win_amd64.exe

lint:
	go vet .
	staticcheck .
	revive .
	govulncheck .

test:
	go test -v
	go test -coverprofile=./test/coverage.out
	go tool cover -html=./test/coverage.out

benchmark:
	go test -bench=. -benchmem -cpu=2 #-benchtime=10s

clean:
	rm -v ./bin/consilium* ./bin/cons*
	rm -v ./test/*
	rm *.db
# end
