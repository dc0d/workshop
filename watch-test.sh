run="(. ./scripts/build.sh) && (. ./scripts/test.sh)"

while true
do
    watchman-wait -p "**/*.go" -- .
    eval $run
done
