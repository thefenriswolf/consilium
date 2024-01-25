##
# insomniplan - schedule planner
#
# @file
# @version 0.1
.PHONY: release clean test benchmark

debug:
	CGO_ENABLED=0 go build -o iplan

release: linux windows osx

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o iplan_linux_amd64

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o iplan_win_amd64.exe

osx:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o iplan_osx_amd64

test:
	go vet .
	staticcheck .

#benchmark:
#	go test -bench=. -benchmem

clean:
	rm -v iplan_linux_* iplan_osx_* iplan_win_* iplan

# end
