while true
do
    watchman-wait -p "**/*.go" -- .
    make test
done
