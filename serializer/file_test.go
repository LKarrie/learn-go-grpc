package serializer_test

import (
	"testing"

	"github.com/LKarrie/learn-go-grpc/pb"
	"github.com/LKarrie/learn-go-grpc/sample"
	"github.com/LKarrie/learn-go-grpc/serializer"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../serializeout/laptop.bin"
	jsonFile := "../serializeout/laptop.json"

	laptop1 := sample.NewLaptop()

	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)

	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtoBufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)
}
