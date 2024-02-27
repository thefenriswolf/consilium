##
# consilium - schedule planner
#
# @file
# @version 0.1
.PHONY: build-release build-debug clean test benchmark lint

run:
	nixGL go run -tags=ebitenginedebug main.go database.go resources.go shapes.go

build-debug:
	go build -o ./bin/cons

build-release: linux windows

build-run: build-debug
	nixGL ./bin/cons

linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/consilium_linux_amd64

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags='-H windowsgui' -o ./bin/consilium_win_amd64.exe

lint:
	go vet .
	staticcheck .
	revive
	govulncheck .

test:
	go test -v
	go test -coverprofile=./test/coverage.out
	go tool cover -html=./test/coverage.out

benchmark:
	go test -bench=. -benchmem -cpu=2 #-benchtime=10s

clean:
	rm -v ./bin/iplan_* ./bin/consilium* ./bin/cons*
	rm -v ./test/*
	rm *.db
# end
