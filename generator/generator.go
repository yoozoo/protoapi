// Package generator code generator module for protoconf
package generator

import (
	"fmt"
	"log"
	"path/filepath"
	"reflect"
	"strings"

	"protoapi/generator/data"
	// this is to let the output plugins initialize themselves and add to the output plugin registra
	_ "protoapi/generator/output"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

const (
	defaultPackageName = "com.yoozoo.configuration"
)

// createEnums create EnumData objects from the passed in enum discriptor
func createEnums(pkg string, enums []*descriptor.EnumDescriptorProto) []*data.EnumData {
	var result []*data.EnumData
	for _, enum := range enums {
		var enumData = new(data.EnumData)
		enumData.Name = enum.GetName()
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
		msgData.Name = message.GetName()
		msgData.File = file

		// the message itself
		fields := message.GetField()
		for _, field := range fields {
			var msgField data.MessageField
			msgField.Name = field.GetName()
			msgField.Key = field.GetName()

			switch field.GetType().String() {
			case "TYPE_STRING":
				msgField.Data = data.StringFieldType
			case "TYPE_BYTES":
				msgField.Data = data.StringFieldType
			case "TYPE_ENUM":
				msgField.Data = field.GetTypeName()
			case "TYPE_MESSAGE":
				msgField.Data = field.GetTypeName()
			case "TYPE_FLOAT":
				msgField.Data = data.DoubleFieldType
			case "TYPE_DOUBLE":
				msgField.Data = data.DoubleFieldType
			case "TYPE_BOOL":
				msgField.Data = data.BooleanFieldType
			default:
				msgField.Data = data.IntFieldType
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

// getMessages returns the flattened message and enum definitions generated from the discriptors
func getMessages(files []*descriptor.FileDescriptorProto) ([]*data.MessageData, []*data.EnumData) {
	var resultMsg []*data.MessageData
	var resultEnum []*data.EnumData
	for _, file := range files {
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
func getMethods(pkg string, methods []*descriptor.MethodDescriptorProto) []data.Method {
	var resultMtd []data.Method

	for _, mtd := range methods {
		log.Printf("mtd: %s\n", mtd)
		var mtdData = data.Method{
			Name:       mtd.GetName(),
			InputType:  mtd.GetInputType()[len(pkg)+2:],
			OutputType: mtd.GetOutputType()[len(pkg)+2:],
		}
		resultMtd = append(resultMtd, mtdData)
	}
	log.Printf("mtds: %s\n", resultMtd)
	return resultMtd
}

// createMessages create message and enum definitions from the passed in descriptor
func createServices(file string, pkg string, services []*descriptor.ServiceDescriptorProto) []*data.ServiceData {
	var resultSers []*data.ServiceData

	for _, service := range services {
		var serData = new(data.ServiceData)
		// the service itself

		// Get all mtds for the service
		mtds := getMethods(pkg, service.GetMethod())
		serData.Name = service.GetName()
		serData.Methods = mtds

		log.Printf("service: %s\n", service)
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

// getPackageName returns the package name from the .proto file on the command line
func getPackageName(request *plugin.CodeGeneratorRequest) string {
	for _, file := range request.ProtoFile {
		if strings.Compare(file.GetName(), request.FileToGenerate[0]) == 0 {
			packageName := file.Package
			if len(*packageName) != 0 {
				return *packageName
			}

		}
	}
	return defaultPackageName
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
			if !isPrimitiveType(field.Data) {
				if v, ok := rootMsgs[field.Data]; ok {
					rootMsgs[field.Data] = v + 1
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
			if !isPrimitiveType(field.Data) {
				// make sure the message does not reference itself
				if strings.Compare(field.Data, msg.Name) != 0 {
					if _, ok := msgMap[field.Data]; ok {
						if _, ok := result[field.Data]; !ok {
							pendingMsgs = append(pendingMsgs, msgMap[field.Data])
						}
						depList[field.Data] = true
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

	// message names
	for _, msg := range messages {
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
	for _, enum := range enums {
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
			if _, ok := translateTable[msg.Fields[idx].Data]; ok {
				// assign use the index access only, as the array is not a point array
				msg.Fields[idx].Data = translateTable[msg.Fields[idx].Data]
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
			if _, ok := enumMap[field.Data]; ok {
				resultEnums = append(resultEnums, enumMap[field.Data])
				// make sure we only add the enum once
				delete(enumMap, field.Data)
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
		if _, ok := msgMap[field.Data]; ok {

			tmp := createKeyList(prefix+field.Name+data.PathSeparator, msgMap[field.Data], msgMap)
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

// Generate the entry point for the code generation module
func Generate(request *plugin.CodeGeneratorRequest) (*plugin.CodeGeneratorResponse, error) {

	var outputLang = "ts"

	if request.Parameter != nil {
		parameters := strings.Split(*(request.Parameter), ":")
		for _, parameter := range parameters {
			kv := strings.Split(parameter, "=")
			if len(kv) == 2 && strings.Compare(kv[0], "lang") == 0 {
				outputLang = kv[1]
			}
		}
		log.Printf("parameter is %s\n", *request.Parameter)
	}

	applicationFile := filepath.Base(request.FileToGenerate[0])
	log.Printf("application File: %s\n", applicationFile)

	applicationName := applicationFile[0 : len(applicationFile)-len(filepath.Ext(applicationFile))]

	packageName := getPackageName(request)

	messages, enums := getMessages(request.ProtoFile)

	services := getServices(request.ProtoFile)

	if services == nil {
		return nil, nil
	}
	log.Printf("Service: %s\n", request.ProtoFile[0].Service)

	if messages == nil {
		return nil, nil
	}

	if outputFunc, ok := data.OutputMap[outputLang]; ok {
		response := new(plugin.CodeGeneratorResponse)
		results, err := outputFunc(applicationName, packageName, services[0], messages, enums)
		for file, content := range results {
			var resultFile = new(plugin.CodeGeneratorResponse_File)
			fileName := file
			resultFile.Name = &fileName
			fileContent := content
			resultFile.Content = &fileContent
			response.File = append(response.File, resultFile)
		}
		return response, err
	}
	err := fmt.Errorf("Output plugin not found for %s\nsupported languages %v", outputLang, reflect.ValueOf(data.OutputMap).MapKeys())
	return nil, err
}
