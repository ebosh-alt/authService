PROTO_DIR=pkg/proto
OUT_DIR=pkg/proto/gen/go/

auth_db:
	cd pkg/db_container
	docker-compose up --build -d

gen:
	protoc --go_out=$(OUT_DIR) --go-grpc_out=$(OUT_DIR) $(PROTO_DIR)/auth.proto