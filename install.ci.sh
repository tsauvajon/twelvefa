PROTOC_VERSION=3.9.0
OS=linux
ARCH=x86_64
PROTOC_URL=https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-${OS}-${ARCH}.zip

wget $PROTOC_URL
unzip protoc-${PROTOC_VERSION}-${OS}-${ARCH}.zip -d /usr/local
go get
go get github.com/golang/protobuf/protoc-gen-go
