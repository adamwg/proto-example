bins: proto
	go build ./cmd/reader
	go build ./cmd/writer

proto:
	protoc --go_out=. model/*.proto

clean:
	-rm reader writer saved-user saved-post

.PHONY: proto clean
