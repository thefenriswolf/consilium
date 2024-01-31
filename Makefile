##
# insomniplan - schedule planner
#
# @file
# @version 0.1
.PHONY: release clean test benchmark lint

debug:
	CGO_ENABLED=0 go build -o ./bin/iplan

release: linux windows osx

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/iplan_linux_amd64

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/iplan_win_amd64.exe

osx:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./bin/iplan_osx_amd64

lint:
	go vet .
	staticcheck .
test:
	go test -v
	go test -coverprofile=./test/coverage.out
	go tool cover -html=./test/coverage.out

#benchmark:
#	go test -bench=. -benchmem

clean:
	rm -v ./bin/iplan_* ./bin/insomniplan* ./bin/iplan*
	rm -v ./test/*

# end
