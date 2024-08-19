package protoraw

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
)

var (
	messageFile  = proto.String("empty_message.proto")
	messageName  = "EmptyMessage"
	protoFactory protoreflect.FileDescriptor
	initErr      error
)

// Decode a protobuf message and return a string representation
func Decode(raw []byte) (string, error) {
	if initErr != nil {
		return "", fmt.Errorf("protoraw: init error: %w", initErr)
	}
	message := dynamicpb.NewMessage(
		protoFactory.Messages().ByName(
			protoreflect.Name(messageName),
		),
	)

	err := proto.Unmarshal(raw, message)
	if err != nil {
		return "", err
	}
	return message.String(), nil
}

func init() {
	fdp := &descriptorpb.FileDescriptorProto{
		Name: messageFile,
		MessageType: []*descriptorpb.DescriptorProto{
			&descriptorpb.DescriptorProto{
				Name: proto.String(messageName),
			},
		},
	}
	protoFactory, initErr = protodesc.NewFile(fdp, &protoregistry.Files{})
}
