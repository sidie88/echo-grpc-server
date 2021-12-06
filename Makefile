protogen:
	protoc -I./proto --go_opt=paths=source_relative --go_out=plugins=grpc:./grpc proto/*.proto
build-run:
	go build -o server && ./server