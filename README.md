# learn-go-grpc
learn-go-grpc

## vscode-proto3 ext
"protoc": {
    "path": "D:/Tools/protoc/bin/protoc",
    "options": [
        "--proto_path=proto"
    ]
}

## vscode auto save setting 
"go.formatTool": "gofmt",
"editor.formatOnType": true,
"editor.formatOnSave": true,

## init
go mod init github.com/LKarrie/learn-go-grpc

## get
go get github.com/google/uuid
go get github.com/golang/protobuf
go get google.golang.org/protobuf
go get github.com/stretchr/testify

## test
### -race
在 Go 语言中，go test -race 命令用于运行带有数据竞争检测（data race detection）的测试。这里的 race 指的是“数据竞争”或“竞态条件”（race condition）。

数据竞争是当两个或更多的 goroutine 并发访问同一内存位置，且至少有一个是写操作，并且没有同步机制来协调这些访问时发生的情况。这可能导致未定义的行为和程序错误。

go test -race 会启用 Go 运行时（runtime）的数据竞争检测器，该检测器会在运行时监控所有内存访问，以查找可能的数据竞争。如果它发现任何数据竞争，它会在测试输出中报告它们。

这种检测对于确保并发代码的正确性非常有用，因为它可以帮助你找到那些可能难以手动检测的问题。然而，由于检测器会增加额外的开销，所以运行带有 -race 标志的测试可能会比普通的测试慢得多。

注意：要使用 -race 标志，你需要使用支持该功能的 Go 版本，并且你的系统需要满足一些特定的要求（例如，某些平台可能需要特殊的构建标志或配置）。

## error
1. cannot use message (variable of type protoreflect.ProtoMessage) as protoiface.MessageV1 value in argument to ProtobufToJSON: protoreflect.ProtoMessage does not implement protoiface.MessageV1 (missing method ProtoMessage)
"google.golang.org/protobuf/proto" => "github.com/golang/protobuf/proto"

2. go test -cover -race ./...
go: -race requires cgo; enable cgo by setting CGO_ENABLED=1
go env -w CGO_ENABLED=1

3. cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in %PATH%

no C compiler
get MinGW https://sourceforge.net/projects/mingw-w64/files/mingw-w64/ x86_64-win32-seh
and set bin path 

change go version to 1.22.1 
-race is not supported on windows/386

