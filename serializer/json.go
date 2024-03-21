package serializer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

// ProtobufToJSON converts protocol buffer message to JSON string
func ProtobufToJSON(message proto.Message) (string, error) {
	marshaler := jsonpb.Marshaler{
		// enum Panel {
		// 	UNKNOW = 0;
		// 	IPS = 1;
		// 	OLED = 2;
		// }
		// if true enum IPS will change to 1
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
		// if false name will be lowerCamelCase
		OrigName: false,
	}

	return marshaler.MarshalToString(message)
}

// JSONToProtobufMessage converts JSON string to protocol buffer message
func JSONToProtobufMessage(data string, message proto.Message) error {
	return jsonpb.UnmarshalString(data, message)
}
