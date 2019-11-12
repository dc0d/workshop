.PHONY: test
test:
	clear
	go test -v -count=1 -timeout 10s -coverprofile=./cover/profile.out -covermode=atomic -coverpkg ./...
	go tool cover -html=./cover/profile.out -o ./cover/coverage.html
