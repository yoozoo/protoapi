// Package generator code generator module for protoconf
package generator

import (
	"fmt"
	"log"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/golang/protobuf/proto"

	"github.com/yoozoo/protoapi/generator/data"
	"github.com/yoozoo/protoapi/util"

	// this is to let the output plugins initialize themselves and add to the output plugin registra
	_ "github.com/yoozoo/protoapi/generator/output"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

const (
	googleDescriptorProtoName = "google/protobuf/descriptor.proto"
)

// createEnums create EnumData objects from the passed in enum discriptor
func createEnums(pkg string, enums []*descriptor.EnumDescriptorProto) []*data.EnumData {
	var result []*data.EnumData
	for _, enum := range enums {
		var enumData = new(data.EnumData)
		enumData.Name = pkg + "." + enum.GetName()
		for _, field := range enum.GetValue() {
			var enumField data.EnumField
			enumField.Name = field.GetName()
			enumField.Value = field.GetNumber()
			enumData.Fields = append(enumData.Fields, enumField)
		}
		result = append(result, enumData)
	}
	return result
}

// createMessages create message and enum definitions from the passed in descriptor
func createMessages(file string, pkg string, messages []*descriptor.DescriptorProto) ([]*data.MessageData, []*data.EnumData) {
	var resultMsg []*data.MessageData
	var resultEnum []*data.EnumData

	for _, message := range messages {
		var msgData = new(data.MessageData)
		msgData.Name = pkg + "." + message.GetName()
		msgData.File = file

		// the message itself
		fields := message.GetField()
		for _, field := range fields {
			var msgField data.MessageField
			msgField.Name = field.GetName()
			msgField.Key = field.GetName()
			msgField.Label = field.GetLabel().String()
			msgField.Options = getFieldOptions(field)

			switch field.GetType().String() {
			case "TYPE_STRING":
				msgField.DataType = data.StringFieldType
			case "TYPE_BYTES":
				msgField.DataType = data.StringFieldType
			case "TYPE_ENUM":
				msgField.DataType = field.GetTypeName()
			case "TYPE_MESSAGE":
				msgField.DataType = field.GetTypeName()
			case "TYPE_FLOAT":
				msgField.DataType = data.DoubleFieldType
			case "TYPE_DOUBLE":
				msgField.DataType = data.DoubleFieldType
			case "TYPE_BOOL":
				msgField.DataType = data.BooleanFieldType
			case "TYPE_INT64":
				msgField.DataType = data.Int64FieldType
			default:
				msgField.DataType = data.IntFieldType
			}

			msgData.Fields = append(msgData.Fields, msgField)
		}
		resultEnum = append(resultEnum, createEnums(msgData.Name, message.GetEnumType())...)
		// msg and enum definitions from the nested messages and enums (recursively)
		msgs, enums := createMessages(file, msgData.Name, message.GetNestedType())
		resultEnum = append(resultEnum, enums...)
		resultMsg = append(resultMsg, msgs...)
		resultMsg = append(resultMsg, msgData)
	}
	return resultMsg, resultEnum
}

func parseMessageDataType(dataType string) string {
	if strings.HasPrefix(dataType, ".") {
		return dataType[1:]
	}

	return dataType
}

// getMessages returns the flattened message and enum definitions generated from the discriptors
func getMessages(files []*descriptor.FileDescriptorProto) ([]*data.MessageData, []*data.EnumData) {
	var resultMsg []*data.MessageData
	var resultEnum []*data.EnumData
	for _, file := range files {
		// exclude google protobuf descriptor proto file
		if file.GetName() == googleDescriptorProtoName {
			continue
		}

		packageName := file.GetPackage()
		// packageName for this file
		if len(packageName) > 0 {
			packageName = "." + packageName
		}

		//enums at file level
		resultEnum = append(resultEnum, createEnums(packageName, file.GetEnumType())...)
		//messages at file level
		msgs, enums := createMessages(file.GetName(), packageName, file.GetMessageType())
		resultEnum = append(resultEnum, enums...)
		resultMsg = append(resultMsg, msgs...)
	}

	return resultMsg, resultEnum
}

// map MethodDescriptorProto to Method
func getMethods(pkg string, service *descriptor.ServiceDescriptorProto) []data.Method {
	methods := service.GetMethod()
	serviceName := service.GetName()
	var resultMtd []data.Method
	log.Printf("proto pkg: %s\n", pkg)
	for _, mtd := range methods {
		var mtdData = data.Method{
			Name:       mtd.GetName(),
			InputType:  parseMessageDataType(mtd.GetInputType()),
			OutputType: parseMessageDataType(mtd.GetOutputType()),
			HttpMtd:    mapHTTPMtd(mtd.GetName()),
			URI:        serviceName + "." + mtd.GetName(),
			Options:    getMethodOptions(mtd),
		}
		resultMtd = append(resultMtd, mtdData)
	}
	return resultMtd
}

// map http method according to the method name, assume only post and get for now
func mapHTTPMtd(method string) string {
	isGet := strings.Contains(method, "Get")
	if isGet {
		return "get"
	} else {
		return "post"
	}
}

// createServices create message and enum definitions from the passed in descriptor
func createServices(file string, pkg string, services []*descriptor.ServiceDescriptorProto) []*data.ServiceData {
	var resultSers []*data.ServiceData

	for _, service := range services {
		var serData = new(data.ServiceData)
		// the service itself

		// Get all mtds for the service
		serData.Name = service.GetName()
		mtds := getMethods(pkg, service)
		serData.Methods = mtds
		serData.Service = service
		serData.Options = getServiceOptions(service)
		serData.CommonErrorType = serData.Options["common_error"]

		resultSers = append(resultSers, serData)
	}
	return resultSers
}

/**
 *	GET all the services in .proto file
 *  Returns an array of service data
 */
func getServices(files []*descriptor.FileDescriptorProto) []*data.ServiceData {
	var resultSers []*data.ServiceData
	var resultEnum []*data.EnumData

	for _, file := range files {
		packageName := file.GetPackage()
		// enums at file level
		resultEnum = append(resultEnum, createEnums(packageName, file.GetEnumType())...)
		// service at file level
		sers := createServices(file.GetName(), packageName, file.GetService())
		resultSers = append(resultSers, sers...)
	}
	return resultSers
}

// getPackageName returns the package name from the .proto file on the command line and if java package is defined, return java package name
func getPackageName(request *plugin.CodeGeneratorRequest) string {
	for _, file := range request.ProtoFile {
		if strings.Compare(file.GetName(), request.FileToGenerate[0]) == 0 {
			// get package name
			if packageName := file.GetPackage(); packageName != "" {
				return packageName
			}

		}
	}
	return ""
}

// isPrimitiveType returns if the field type is considered primitive (ie can be translated to language/built-in of the target code)
func isPrimitiveType(fieldType string) bool {
	switch fieldType {
	case data.StringFieldType:
		return true
	case data.DoubleFieldType:
		return true
	case data.IntFieldType:
		return true
	case data.BooleanFieldType:
		return true
	default:
		return false
	}
}

// findRootObject  find the root object from the command line .proto file
// root object is the message that is not depended by other message
// there should be one and only one root object ??
func findRootObject(file string, messages []*data.MessageData, msgMap map[string]*data.MessageData) *data.MessageData {

	rootMsgs := make(map[string]int)

	// get the list of messages defined in the .proto file on command line
	for _, message := range messages {
		if strings.Compare(message.File, file) == 0 {
			rootMsgs[message.Name] = 0
		}
	}

	// check how many other messages are depending on it.
	for k := range rootMsgs {
		for _, field := range msgMap[k].Fields {
			if !isPrimitiveType(field.DataType) {
				if v, ok := rootMsgs[field.DataType]; ok {
					rootMsgs[field.DataType] = v + 1
				}
			}
		}
	}

	// find root object, if find more than one, it is fatal error
	var rootMsg *data.MessageData
	for k, v := range rootMsgs {
		// log.Printf("root msg", v, k)
		if v == 0 {
			if rootMsg == nil {
				rootMsg = msgMap[k]
			}
		}
	}
	// no root object find, error also
	if rootMsg == nil {
		log.Printf("We could not find root configuration object\n")
	}
	return rootMsg
}

// buildDepGraph build the dependency graph in the format below
// node -> list of nodes it depends on
func buildDepGraph(rootMsg *data.MessageData, msgMap map[string]*data.MessageData) map[string](map[string]bool) {
	var result = make(map[string](map[string]bool))
	var pendingMsgs = []*data.MessageData{rootMsg}

	for len(pendingMsgs) > 0 {
		// array to simulate fifo of the messages to process
		msg := pendingMsgs[0]
		pendingMsgs = pendingMsgs[1:]

		var depList = make(map[string]bool)

		for _, field := range msg.Fields {
			if !isPrimitiveType(field.DataType) {
				// make sure the message does not reference itself
				if strings.Compare(field.DataType, msg.Name) != 0 {
					if _, ok := msgMap[field.DataType]; ok {
						if _, ok := result[field.DataType]; !ok {
							pendingMsgs = append(pendingMsgs, msgMap[field.DataType])
						}
						depList[field.DataType] = true
					}
				} else {
					log.Printf("Object cant reference itself: %s (in %s)\n", msg.Name, msg.File)
					return nil
				}
			}
		}
		result[msg.Name] = depList
	}
	return result
}

// findDepOrder do a topological sort to create the order of the message definitions
func findDepOrder(rootMsg *data.MessageData, msgMap map[string]*data.MessageData) []string {
	msgGraph := buildDepGraph(rootMsg, msgMap)
	if msgGraph == nil {
		return nil
	}
	var result []string
	for len(msgGraph) > 0 {
		var msgName string
		for k, v := range msgGraph {
			if len(v) == 0 {
				msgName = k
				delete(msgGraph, k)
				break
			}
		}
		// if curcular dependency found, error out
		if len(msgName) == 0 {
			log.Printf("Circular dependency found. This is not supported.\n")
			return nil
		}

		result = append(result, msgName)
		for _, v := range msgGraph {
			delete(v, msgName)
		}
	}
	return result
}

// fixMessageName make sure the names of the messages are unique when put them under the same package/namespace
func fixMessageName(messages []*data.MessageData, enums []*data.EnumData) {
	var postfix = 1
	var existName = make(map[string]bool) // existing shortended name
	var translateTable = make(map[string]string)

	// message names in reversed order
	for i := len(messages) - 1; i >= 0; i-- {
		msg := messages[i]
		name := msg.Name[strings.LastIndexByte(msg.Name, '.')+1:]
		originalName := name
		for {
			// if not unique , append "_NNN" NNN is the current postfix value
			if _, ok := existName[name]; ok {
				name = fmt.Sprintf("%s_%d", originalName, postfix)
			} else {
				break
			}
		}
		existName[name] = true
		translateTable[msg.Name] = name
	}
	//enum names
	for i := len(enums) - 1; i >= 0; i-- {
		enum := enums[i]
		name := enum.Name[strings.LastIndexByte(enum.Name, '.')+1:]
		originalName := name
		for {
			if _, ok := existName[name]; ok {
				name = fmt.Sprintf("%s_%d", originalName, postfix)
				postfix++
			} else {
				break
			}
		}
		existName[name] = true
		translateTable[enum.Name] = name
		enum.Name = name
	}

	// convert all message/enum names appeared in messages
	for _, msg := range messages {
		msg.Name = translateTable[msg.Name]
		for idx := range msg.Fields {
			if _, ok := translateTable[msg.Fields[idx].DataType]; ok {
				// assign use the index access only, as the array is not a point array
				msg.Fields[idx].DataType = translateTable[msg.Fields[idx].DataType]
			}
		}
	}
}

// filterMessages reorgnize the message/enum to the way we need
// 1. remove not used items
// 2. reorder according to dependency order
// 3. check and modify item names
func filterMessages(file string, messages []*data.MessageData, enums []*data.EnumData) ([]*data.MessageData, []*data.EnumData) {
	msgMap := make(map[string]*data.MessageData)
	enumMap := make(map[string]*data.EnumData)

	for _, message := range messages {
		msgMap[message.Name] = message
	}
	for _, enum := range enums {
		enumMap[enum.Name] = enum
	}

	rootMsg := findRootObject(file, messages, msgMap)
	if rootMsg == nil {
		return nil, nil
	}

	msgNames := findDepOrder(rootMsg, msgMap)

	if msgNames == nil {
		return nil, nil
	}

	// create the result based on the message order
	var resultMsgs []*data.MessageData
	var resultEnums []*data.EnumData
	for _, name := range msgNames {
		msg := msgMap[name]
		resultMsgs = append(resultMsgs, msg)
		for _, field := range msg.Fields {
			if _, ok := enumMap[field.DataType]; ok {
				resultEnums = append(resultEnums, enumMap[field.DataType])
				// make sure we only add the enum once
				delete(enumMap, field.DataType)
			}
		}
	}

	fixMessageName(resultMsgs, resultEnums)

	return resultMsgs, resultEnums
}

//createkeyList recursively create the key list
func createKeyList(prefix string, msg *data.MessageData, msgMap map[string]*data.MessageData) []string {
	var result []string
	log.Printf("msg: %s\n", msg.Name)
	log.Printf("msg.Fields: %s\n", msg.Fields)
	for _, field := range msg.Fields {
		if _, ok := msgMap[field.DataType]; ok {

			tmp := createKeyList(prefix+field.Name+data.PathSeparator, msgMap[field.DataType], msgMap)
			for _, v := range tmp {
				result = append(result, v)
			}
			log.Printf("tmp: %s\n", tmp)
		} else {
			result = append(result, prefix+field.Name)
		}
	}
	//log.Printf("keys: %s\n", result)
	return result
}

//generateKeyList returns a list of strings use as kv store key
func generateKeyList(messages []*data.MessageData) []string {
	var msgMap = make(map[string]*data.MessageData)
	for _, msg := range messages {
		msgMap[msg.Name] = msg
		log.Printf("msg.Name: %s\n", msg.Name)
		log.Printf("msg: %s\n", msg)
	}
	return createKeyList("", messages[len(messages)-1], msgMap)
}

// Get method options from proto file
func getFieldOptions(fieldPb *descriptor.FieldDescriptorProto) data.OptionMap {
	options := make(map[string]string)
	// create extension description
	for field, name := range data.FieldOptions {
		var extDesc = &proto.ExtensionDesc{
			ExtendedType:  (*descriptor.FieldOptions)(nil),
			ExtensionType: (*string)(nil),
			Field:         field,
			Name:          name,
			Tag:           "bytes," + string(field) + ",opt,name=" + name,
		}

		ext, err := proto.GetExtension(fieldPb.GetOptions(), extDesc)
		if err == nil {
			// add the service method option to the method data
			options[name] = *ext.(*string)
		}
	}
	return options
}

// Get method options from proto file
func getMethodOptions(method *descriptor.MethodDescriptorProto) data.OptionMap {
	options := make(map[string]string)
	// create extension description
	for field, name := range data.MethodOptions {
		var extDesc = &proto.ExtensionDesc{
			ExtendedType:  (*descriptor.MethodOptions)(nil),
			ExtensionType: (*string)(nil),
			Field:         field,
			Name:          name,
			Tag:           "bytes," + string(field) + ",opt,name=" + name,
		}

		ext, err := proto.GetExtension(method.GetOptions(), extDesc)
		if err == nil {
			// add the service method option to the method data
			options[name] = *ext.(*string)
		}
	}
	return options
}

// Get service options from proto file
func getServiceOptions(service *descriptor.ServiceDescriptorProto) data.OptionMap {
	options := make(map[string]string)
	// create extension description
	for field, name := range data.ServiceOptions {
		var extDesc = &proto.ExtensionDesc{
			ExtendedType:  (*descriptor.ServiceOptions)(nil),
			ExtensionType: (*string)(nil),
			Field:         field,
			Name:          name,
			Tag:           "bytes," + string(field) + ",opt,name=" + name,
		}

		ext, err := proto.GetExtension(service.GetOptions(), extDesc)
		if err == nil {
			// add the service method option to the method data
			options[name] = *ext.(*string)
		}
	}
	return options
}

// Get s from proto file
func getFileOptions(request *plugin.CodeGeneratorRequest) data.OptionMap {
	for _, file := range request.ProtoFile {
		if strings.Compare(file.GetName(), request.FileToGenerate[0]) == 0 {
			// check options from .proto file
			if fileOptions := file.GetOptions(); fileOptions != nil {
				options := make(map[string]string)
				// get java package from options
				if javaPackageName := fileOptions.GetJavaPackage(); javaPackageName != "" {
					options[data.JavaPackageOption] = javaPackageName
				}
				return options
			}

		}
	}
	return nil
}

// Generate the entry point for the code generation module
func Generate(input []byte) *plugin.CodeGeneratorResponse {
	request := new(plugin.CodeGeneratorRequest)

	err := proto.Unmarshal(input, request)
	if err != nil {
		util.Die(fmt.Errorf("invalid CodeGeneratorRequest: %v", err))
	}

	if len(request.FileToGenerate) != 1 {
		util.Die(fmt.Errorf("Multiple input files given: %v\nprotoapi only support one proto file", request.FileToGenerate))
	}

	var outputLang = "ts"
	var params = make(map[string]string)
	parameter := request.GetParameter()

	if parameter != "" {
		parameters := strings.Split(*(request.Parameter), ",")
		for _, parameter := range parameters {
			kv := strings.Split(parameter, "=")
			if len(kv) == 2 {
				if strings.Compare(kv[0], "lang") == 0 {
					outputLang = kv[1]
				}
				params[kv[0]] = kv[1]
			} else {
				params[kv[0]] = ""
			}
		}
	}

	applicationFile := filepath.Base(request.FileToGenerate[0])
	log.Printf("proto file: %s\n", applicationFile)
	log.Printf("code generated: %s\n", outputLang)

	applicationName := applicationFile[0 : len(applicationFile)-len(filepath.Ext(applicationFile))]

	packageName := getPackageName(request)

	options := getFileOptions(request)

	messages, enums := getMessages(request.ProtoFile)
	// Fix same message name issue
	fixMessageName(messages, enums)

	services := getServices(request.ProtoFile)

	var service *data.ServiceData
	if len(services) > 1 {
		util.Die(fmt.Errorf("found %d services; only 1 service is supported now", len(services)))
	} else if len(services) == 1 {
		service = services[0]
	}

	data.Setup(request)

	if gen, ok := data.OutputMap[outputLang]; ok {
		response := new(plugin.CodeGeneratorResponse)
		gen.Init(request)

		results, err := gen.Gen(applicationName, packageName, service, messages, enums, options)
		if err != nil {
			util.Die(err)
		}
		for file, content := range results {
			var resultFile = new(plugin.CodeGeneratorResponse_File)
			// generate the file to the specified package
			fileName := file
			resultFile.Name = &fileName
			fileContent := content
			resultFile.Content = &fileContent
			response.File = append(response.File, resultFile)
		}
		return response
	}

	err = fmt.Errorf("Output plugin not found for %s\nsupported options: %v", outputLang, reflect.ValueOf(data.OutputMap).MapKeys())
	util.Die(err)

	return nil
}
