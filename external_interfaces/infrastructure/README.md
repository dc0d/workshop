mockgen -source ./../model/event_store.go -destination mock_event_store_test.go -package infrastructure_test

mockgen -source ./../model/event_publisher.go -destination mock_event_publisher_test.go -package infrastructure_test
