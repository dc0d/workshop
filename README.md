# Test

To run all tests:

```
go test -v -count=1 ./...
```

Mocks generated using:

```
mockgen -source file/path.go | pbcopy
```

Some test double are generated automatically from interfaces using `go generate ./...` - using moq.

# TODO

This is just a practice and there are room for improvement:

- use real software for infrastructure (something like RabbitMQ for message queue, database, etc)
- reconsiliate is any versions are missing (in views)
- propagate those/some/waht? events from event handler to inform others view has been updated?

should event publishing be handled by event store or the repository?

- at the repository we already have the event object (no need to deserialize it) and we can convert it to an event descriptor.
