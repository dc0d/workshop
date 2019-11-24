mockgen -source ./../domain_model/event_store.go -destination mock_event_store_test.go -package infrastructure_test

mockgen -source ./../domain_model/event_publisher.go -destination mock_event_publisher_test.go -package infrastructure_test
