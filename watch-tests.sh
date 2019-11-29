run="(. ./scripts/test.sh)"

while true
do
    watchman-wait -p "**/*.go" -- .
    eval $run
done
