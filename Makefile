bins: proto
	go build ./cmd/reader
	go build ./cmd/writer
	go build ./cmd/posts-server
	go build ./cmd/posts-client
	go build ./cmd/posts-web

proto:
	protoc --go_out=plugins=grpc:. model/*.proto

clean:
	-rm reader writer saved-user saved-post posts-server posts-client posts-web

.PHONY: proto clean
