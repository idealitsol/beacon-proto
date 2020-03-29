#!/usr/bin/env sh

# Install proto3 from source
#  brew install autoconf automake libtool
#  git clone https://github.com/google/protobuf
#  ./autogen.sh ; ./configure ; make ; make install
#
# Update protoc Go bindings via
#  go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
#
# See also
#  https://github.com/grpc/grpc-go/tree/master/examples
export PATH=$PATH:$GOPATH/bin

rm -rf pbx/*.pb.go grpc/*.pb.go
for filename in pbx/*.proto; do
  echo $filename
  protoc "$filename" --go_out=plugins=grpc:.
done

protoc "grpc/health.proto" --go_out=plugins=grpc:.

protolint lint -fix .