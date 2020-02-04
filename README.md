# Test

## Tools

`moq` (for generating Mocks) and `wire` (for Dependency Injection) must be present. To get them:

```
$ go get -u -v github.com/google/wire/cmd/wire
$ go get -u -v github.com/matryer/moq
```

## Run The Tests

To run all tests:

```
go test -v -count=1 ./...
```

Mocks generated using:

```
mockgen -source file/path.go | pbcopy
```

Some test double are generated automatically from interfaces using `go generate ./...` - using moq.

# Tooling

To avoid being forced by default linter, `golint`, it's possible to use a drop-in replacement, [`revive`](https://github.com/mgechev/revive).

Go to home directory (`$ cd`) and create this config file for `revive`:

```
ignoreGeneratedHeader = false
severity = "warning"
confidence = 0.8
errorCode = 0
warningCode = 0

[rule.blank-imports]
[rule.context-as-argument]
[rule.context-keys-type]
[rule.dot-imports]
[rule.error-return]
[rule.error-strings]
[rule.error-naming]
# [rule.exported]
[rule.if-return]
[rule.increment-decrement]
[rule.var-naming]
[rule.var-declaration]
[rule.package-comments]
[rule.range]
[rule.receiver-naming]
[rule.time-naming]
[rule.unexported-return]
[rule.indent-error-flow]
[rule.errorf]
[rule.empty-block]
[rule.superfluous-else]
[rule.unused-parameter]
[rule.unreachable-code]
[rule.redefines-builtin-id]
```

Notice the commented line, which stops the linter from complaining on missing documentation.

In VSCode provide this setting (for Go extension):

```
{
    "go.lintTool": "revive",
    "go.lintFlags": ["--config=~/.golint_revive_config.toml"]
}
```

# TODO

This is just a practice and there are room for improvement:

- use real software for infrastructure (something like RabbitMQ for message queue, database, etc)
- reconsiliate is any versions are missing (in views)
- propagate those/some/waht? events from event handler to inform others view has been updated?

should event publishing be handled by event store or the repository?

- at the repository we already have the event object (no need to deserialize it) and we can convert it to an event descriptor.
