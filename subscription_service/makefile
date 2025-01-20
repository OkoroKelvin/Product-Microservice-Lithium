generate:
	@echo "Generating gRPC code..."
	@protoc \
		--go_out=. \
		--go-grpc_out=. proto/subscription.proto \


mocks:
	@echo "Generating mocks..."
	@mockgen -source=internal/repositories/subscription.repository.go -destination=internal/repositories/mocks/subscription.repository.go -package=mocks
	@mockgen -source=internal/services/subscription.service.go -destination=internal/services/mocks/subscription.service.go -package=mocks