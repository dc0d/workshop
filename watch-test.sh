run="(. ./scripts/build.sh) && (. ./scripts/test.sh)"

while true
do
    watchman-wait -p "**/*.ex" "**/*.exs" -- .
    eval $run
done
