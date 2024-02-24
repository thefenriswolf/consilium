##
# insomniplan - schedule planner
#
# @file
# @version 0.1
.PHONY: build-release build-debug clean test benchmark lint

run:
	nixGL go run -tags=ebitenginedebug main.go database.go resources.go

build-debug:
	go build -o ./bin/iplan

build-release: linux windows

build-run: build-debug
	nixGL ./bin/iplan

linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/iplan_linux_amd64

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags='-H windowsgui' -o ./bin/iplan_win_amd64.exe

lint:
	go vet .
	staticcheck .
	govulncheck .
test:
	go test -v
	go test -coverprofile=./test/coverage.out
	go tool cover -html=./test/coverage.out

#benchmark:
#	go test -bench=. -benchmem

clean:
	rm -v ./bin/iplan_* ./bin/insomniplan* ./bin/iplan*
	rm -v ./test/*
	rm *.db
# end
