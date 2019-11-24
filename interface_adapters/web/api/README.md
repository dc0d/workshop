mockgen -source ./../../../domain_model/usecase_bank_statement.go -destination mock_bank_statement_test.go -package api

mockgen -source ./../../../domain_model/usecase_handle_transaction.go -destination mock_handle_transaction_test.go -package api

