package generator

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/yoozoo/protoapi/util"
)

// define file/message/field structs to be used in language generators
// wrap protoc-gen-go/descriptor to provide helper methods

// GenerateReq is the code-gen request struct passed to generators
type GenerateReq struct {
	Files map[string]*ProtoFile
}

func NewGenerateReq(request *plugin.CodeGeneratorRequest) *GenerateReq {
	result := &GenerateReq{}
	result.Files = make(map[string]*ProtoFile)

	for _, file := range request.ProtoFile {
		pf := NewProtoFile(file)
		result.Files[file.GetName()] = pf

		if util.IsStrInSlice(file.GetName(), request.FileToGenerate) {
			pf.IsFileToGenerate = true
		}
	}

	return result
}

// ProtoFile is a thin wrapper around descriptor.FileDescriptorProto
type ProtoFile struct {
	IsFileToGenerate bool
	Proto            *descriptor.FileDescriptorProto
	Options          map[string]*ProtoOption
	Enums            map[string]*ProtoEnum
	Messages         map[string]*ProtoMessage
	Services         map[string]*ProtoService
}

// NewProtoFile create ProtoFile from descriptor.FileDescriptorProto
func NewProtoFile(proto *descriptor.FileDescriptorProto) *ProtoFile {
	p := &ProtoFile{
		Proto: proto,
	}

	p.Messages = make(map[string]*ProtoMessage)
	for _, msg := range proto.MessageType {
		p.Messages[msg.GetName()] = NewProtoMessage(msg)
	}

	p.Services = make(map[string]*ProtoService)
	for _, svr := range proto.Service {
		p.Services[svr.GetName()] = NewProtoService(svr)
	}

	p.Enums = make(map[string]*ProtoEnum)
	for _, obj := range proto.EnumType {
		p.Enums[obj.GetName()] = NewProtoEnum(obj)
	}

	return p
}

// ProtoOption is a thin wrapper around descriptor.OptionDescriptorProto
type ProtoOption struct {
}

// ProtoEnum is a thin wrapper around descriptor.EnumDescriptorProto
type ProtoEnum struct {
	Proto *descriptor.EnumDescriptorProto
}

// NewProtoEnum create ProtoEnum from descriptor.EnumDescriptorProto
func NewProtoEnum(proto *descriptor.EnumDescriptorProto) *ProtoEnum {
	return &ProtoEnum{
		Proto: proto,
	}
}

// ProtoMessage is a thin wrapper around descriptor.DescriptorProto (Message descriptor)
type ProtoMessage struct {
	Proto   *descriptor.DescriptorProto
	Options map[string]*ProtoOption
	Fields  map[string]*ProtoField
}

// NewProtoMessage create ProtoMessage from descriptor.DescriptorProto
func NewProtoMessage(proto *descriptor.DescriptorProto) *ProtoMessage {
	return &ProtoMessage{
		Proto: proto,
	}
}

// ProtoField is a thin wrapper around descriptor.FieldDescriptorProto
type ProtoField struct {
	Proto   *descriptor.FieldDescriptorProto
	Options map[string]*ProtoOption
}

// NewProtoField create ProtoField from descriptor.FieldDescriptorProto
func NewProtoField(proto *descriptor.FieldDescriptorProto) *ProtoField {
	return &ProtoField{
		Proto: proto,
	}
}

// ProtoMethod is a thin wrapper around descriptor.MethodDescriptorProto
type ProtoMethod struct {
	Proto   *descriptor.MethodDescriptorProto
	Options map[string]*ProtoOption
}

// NewProtoMethod create ProtoMethod from descriptor.MethodDescriptorProto
func NewProtoMethod(proto *descriptor.MethodDescriptorProto) *ProtoMethod {
	return &ProtoMethod{
		Proto: proto,
	}
}

// ProtoService is a thin wrapper around descriptor.ServiceDescriptorProto
type ProtoService struct {
	Proto   *descriptor.ServiceDescriptorProto
	Options map[string]*ProtoOption
	Methods map[string]*ProtoMethod
}

// NewProtoService create ProtoService from descriptor.ServiceDescriptorProto
func NewProtoService(proto *descriptor.ServiceDescriptorProto) *ProtoService {
	return &ProtoService{
		Proto: proto,
	}
}
