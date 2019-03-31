VERSION=v0.0.1
SVC=greetings
BIN=$(PWD)/bin/$(SVC)
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

run r:
	@echo "[running] Running service..."
	@go run cmd/server/main.go

build b:
	@echo "[build] Building service..."
	@cd cmd/server && $(GO) build -o $(BIN) -ldflags=$(LDFLAGS) -tags $(TAGS)

linux l:
	@echo "[build] Building for linux..."
	@cd cmd/server && GOOS=linux $(GO) build -a -o $(BIN) --ldflags $(LDFLAGS) -tags $(TAGS)
