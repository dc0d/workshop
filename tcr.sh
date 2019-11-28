run="(. ./scripts/build.sh) && ((. ./scripts/test.sh) && (. ./scripts/commit.sh) || (. ./scripts/revert.sh))"

while true
do
    watchman-wait -p "**/*.go" -- .
    eval $run
done
