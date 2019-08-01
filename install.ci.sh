PROTOC_VERSION=3.9.0
OS=linux
ARCH=x86_64
PROTOC=protoc-${PROTOC_VERSION}-${OS}-${ARCH}.zip
PROTOC_URL=https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/${PROTOC}

wget $PROTOC_URL
mkdir tmp
unzip ${PROTOC} -d tmp
go get
go get github.com/golang/protobuf/protoc-gen-go
