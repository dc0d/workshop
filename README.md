# A TDD Workflow; TCR

Source: [test && commit || revert](https://medium.com/@kentbeck_7670/test-commit-revert-870bbd756864)

To start the tcr loop:

```
$ . ./tcr.sh
```

The test command that is running (inside `./scripts/test.sh`) is:

```
$ go test -v -count=1 -timeout 10s ./...
```

It is possible to adapt this to your style.

A build step is added to make sure the code is valid before running the TCR part - so here it's `build && (test && commit || revert)`.

The test code will not be reverted. Only the code will be reverted - in case it fails to fulfil the expectations (tests).

Each successful change will be committed with a `WIP` message. So it's better to work in a local branch and then rebase or merge/squash.

# Traditional TDD

Still sometimes I find myself to fall back to traditional TDD. For such occasions, it's possible to run just the tests using:

```
$ . ./watch-test.sh
```
