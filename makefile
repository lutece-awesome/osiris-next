MAIN := main
SRC := .
DIST := dist
PROTO := proto
CODEGEN := codegen
PROTOC := protoc
PROTO_SRC := $(wildcard proto/*.proto)
.PHONY: clean format fmt

# Call protoc to generate the code.
gen_proto:
	@ ${PROTOC} -I=${PROTO}/ \
				--go-grpc_out=${CODEGEN}/ \
				--go-grpc_opt=paths=source_relative \
				--go_out=${CODEGEN}/ \
				--go_opt=paths=source_relative\
				${PROTO_SRC}

fmt_go:
	@ gofmt -w -s .

fmt_proto:
	@ clang-format -style=Google -i ${PROTO_SRC}

format: fmt_go fmt_proto

# Alias of format
fmt: format

clean:
	@ rm -r ${DIST}