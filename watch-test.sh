while true
do
    watchman-wait -p "**/*.go" -- .
    go test -count=1 -timeout 30s -p 1 -cover ./...
done
