.PHONY: test
test:
	clear
	go test -count=1 -timeout 30s -p 1 -cover ./...

.PHONY: cover
cover:
	go test -count=1 -timeout 60s -p 1 -coverprofile=./all-profile.out -coverpkg=./... ./...

cover-html:
	clear
	go test -count=1 -timeout 60s -p 1 -coverprofile=./all-profile.out -coverpkg=./... ./...
	go tool cover -html=./all-profile.out -o ./all-coverage.html

lint:
	golangci-lint run ./...
