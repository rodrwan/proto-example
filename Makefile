VERSION=v0.0.1
SVC=greetings
PROTO_FOLDER=proto
PROTO_SVC=greetings

GO ?= go
LDFLAGS='-extldflags "static" -X main.svcVersion=$(VERSION) -X main.svcName=$(SVC)'
TAGS=netgo -installsuffix netgo

proto p:
	@echo "[proto] Generate proto for go"
	rm -f $(PROTO_FOLDER)/$(PROTO_SVC).pb.go
	protoc -I $(PROTO_FOLDER) $(PROTO_FOLDER)/*.proto \
		--go_out=plugins=grpc:$(PROTO_FOLDER)

client c:
	@echo "[running] Running client..."
	@go run cmd/client/main.go

server s:
	@echo "[running] Running server..."
	@go run cmd/server/main.go
