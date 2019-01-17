package output

import (
	"errors"
	"fmt"
	"log"
	"strings"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/yoozoo/protoapi/generator/data"
	yii2 "github.com/yoozoo/protoapi/generator/output/phpyii2"
	"github.com/yoozoo/protoapi/util"
)

type yii2Gen struct {
	result     map[string]string
	enums      []*data.EnumData
	ModuleName string
	NameSpace  string
	bizErrors  []string
	comError   *data.MessageData
}

func (g *yii2Gen) Init(request *plugin.CodeGeneratorRequest) {
	for _, file := range request.ProtoFile {
		if file.GetName() == googleDescriptorProtoName {
			continue
		}

		opts := file.GetOptions()
		if opts == nil || opts.GetPhpNamespace() == "" {
			continue
		}

		if g.NameSpace == "" {
			g.NameSpace = opts.GetPhpNamespace()
		}
	}
}

/* generate functions */
func (g *yii2Gen) genController(methods []*data.Method) error {
	obj := yii2.NewController(g.NameSpace, methods)
	err := obj.Gen(g.result)
	if err != nil {
		return err
	}
	return nil
}

func (g *yii2Gen) genEnum(enum *data.EnumData) error {
	obj := yii2.NewEnum(enum, g.NameSpace)
	err := obj.Gen(g.result)
	if err != nil {
		return err
	}
	return nil
}

func (g *yii2Gen) genError(msg *data.MessageData) error {
	obj := yii2.NewError(msg, g.NameSpace, g.enums)

	err := obj.Gen(g.result)
	if err != nil {
		return err
	}
	return nil
}

func (g *yii2Gen) genHandler(methods []*data.Method) error {
	obj := yii2.NewHandler(methods, g.NameSpace)

	err := obj.Gen(g.result)
	if err != nil {
		return err
	}
	return nil
}

func (g *yii2Gen) genMessage(msg *data.MessageData) error {
	obj := yii2.NewMessage(msg, g.NameSpace, g.enums)

	err := obj.Gen(g.result)
	if err != nil {
		return err
	}
	return nil
}

func (g *yii2Gen) genModule(service *data.ServiceData) error {
	obj := yii2.NewModule(g.NameSpace, service)
	err := obj.Gen(g.result)
	if err != nil {
		return err
	}
	return nil
}

/* ****************** */

func (g *yii2Gen) isBizErr(msg *data.MessageData) bool {
	for _, field := range g.bizErrors {
		if field == msg.Name {
			return true
		}
	}
	return false
}

func (g *yii2Gen) isComErr(msg *data.MessageData) bool {
	for _, field := range g.comError.Fields {
		if field.DataType == msg.Name {
			return true
		}
	}
	return false
}

func (g *yii2Gen) Gen(applicationName string, packageName string, services []*data.ServiceData, messages []*data.MessageData, enums []*data.EnumData, options data.OptionMap) (result map[string]string, err error) {
	var service *data.ServiceData
	if len(services) > 1 {
		util.Die(fmt.Errorf("found %d services; only 1 service is supported now", len(services)))
	} else if len(services) == 1 {
		service = services[0]
	}

	// set enums
	g.enums = enums

	// set namespace
	if g.NameSpace == "" {
		g.NameSpace = strings.Replace(packageName, ".", "\\", -1)

		if g.NameSpace == "" {
			util.Die(fmt.Errorf("No name space given"))
		}

		log.Printf("Use proto package name for php: %v", g.NameSpace)
	}
	// make sure namespace start with app\modules
	if !strings.HasPrefix(g.NameSpace, "app\\modules\\") {
		g.NameSpace = "app\\modules\\" + g.NameSpace
	}

	g.ModuleName = applicationName
	g.result = make(map[string]string)

	// create error map
	for _, serv := range service.Methods {
		errorMsgName, found := serv.Options["error"]
		if found {
			g.bizErrors = append(g.bizErrors, errorMsgName)
		}
	}

	for i, msg := range messages {
		if msg.Name == data.ComErrMsgName {
			g.comError = msg
			messages = append(messages[:i], messages[i+1:]...)
			break
		}
	}
	if g.comError == nil {
		return nil, errors.New("Cannot find common error message")
	}

	// call genarator functions one by one
	err = g.genController(service.Methods)
	if err != nil {
		return nil, err
	}
	err = g.genModule(service)
	if err != nil {
		return nil, err
	}
	err = g.genHandler(service.Methods)
	if err != nil {
		return nil, err
	}
	for _, enum := range enums {
		err = g.genEnum(enum)
		if err != nil {
			return nil, err
		}
	}

	for _, msg := range messages {
		data.FlattenLocalPackage(msg)

		if g.isBizErr(msg) {
			err = g.genError(msg)
			if err != nil {
				return nil, err
			}
		} else if g.isComErr(msg) {
			err = g.genError(msg)
			if err != nil {
				return nil, err
			}
		} else {
			// rename 'Empty' message to 'Blank' to avoid PHP compilation error
			if "EMPTY" == strings.ToUpper(msg.Name) {
				msg.Name = "Blank"
			}

			err := g.genMessage(msg)
			if err != nil {
				return nil, err
			}
		}
	}

	return g.result, nil
}

func init() {
	data.OutputMap["yii2"] = &yii2Gen{}
}
