package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"

	"version.uuzu.com/Merlion/protoapi/generator/data"
	"version.uuzu.com/Merlion/protoapi/util"

	"github.com/spf13/cobra"
)

const (
	defaultProtocCmd     = "protoc"
	protocFlag           = "protoc"
	langFlag             = "lang"
	protoPathFlag        = "proto_path"
	protoCustomParamFlag = "custom_params"
)

type genFlagData struct {
	langValue        string
	protocPath       string
	protoIncPath     string
	protoCustomParam string
}

func (g *genFlagData) reset() {
	g.langValue = ""
	g.protocPath = ""
	g.protoIncPath = ""
	g.protoCustomParam = ""
}

var genFlagValue genFlagData

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen <output dir> <proto file>",
	Short: "generate code from proto file",
	Long: `This command will read the input proto file and generate
	code of the requested language to the output directory.`,
	Args: cobra.ExactArgs(2),
	Run:  generateCode,
}

func generateCode(cmd *cobra.Command, args []string) {
	defer func() {
		genFlagValue.reset()
	}()

	executable, _ := os.Executable()

	var params = make(map[string]string)
	params[langFlag] = genFlagValue.langValue

	if _, ok := data.OutputMap[genFlagValue.langValue]; !ok {
		err := fmt.Errorf("Output plugin not found for %s\nsupported options: %v",
			genFlagValue.langValue, reflect.ValueOf(data.OutputMap).MapKeys())
		util.Die(err)
	}

	protoc := genFlagValue.protocPath

	if len(protoc) == 0 {
		protoc, genFlagValue.protoIncPath = util.GetDefaultProtoc(genFlagValue.protoIncPath)
	}
	protoc = filepath.FromSlash(protoc)

	cmdParam := ""
	for name, value := range params {
		if len(value) > 0 {
			if len(cmdParam) > 0 {
				cmdParam += ","
			}
			cmdParam += name + "=" + value
		}
	}

	if len(genFlagValue.protoCustomParam) > 0 {
		cmdParam += "," + genFlagValue.protoCustomParam
	}

	protoFile := filepath.FromSlash(args[1])
	stat, err := os.Stat(protoFile)
	if err != nil {
		util.Die(fmt.Errorf("Input %s is not accessible : %s", protoFile, err.Error()))
	}
	if stat.IsDir() {
		util.Die(fmt.Errorf("Input %s is not a file", protoFile))
	}
	outputDir := filepath.FromSlash(args[0])
	stat, err = os.Stat(outputDir)
	if err != nil || !stat.IsDir() {
		util.Die(fmt.Errorf("Output directory %s is not accessible", outputDir))
	}

	var arglist []string

	protoIncPath := util.GetIncludePath(filepath.FromSlash(genFlagValue.protoIncPath), filepath.Dir(protoFile))
	arglist = append(arglist, "--"+protoPathFlag+"="+protoIncPath)
	arglist = append(arglist, "--plugin=protoc-gen-custom="+executable)
	arglist = append(arglist, "--custom_out="+cmdParam+":"+outputDir)
	arglist = append(arglist, protoFile)
	protoCmd := exec.Command(protoc, arglist...)

	protoCmd.Stderr = os.Stderr
	err = protoCmd.Run()

	if err != nil {
		util.Die(fmt.Errorf("Error to execute protoc: %s", err))
	}
}

func init() {
	RootCmd.AddCommand(genCmd)

	genCmd.Flags().StringVar(&genFlagValue.langValue, langFlag, "", "language of the generated code, default is ts.")
	genCmd.Flags().StringVar(&genFlagValue.protoIncPath, protoPathFlag, "", "extra proto file import paths, seperated by ':'(unix) or ';'(windows)")
	genCmd.Flags().StringVar(&genFlagValue.protoCustomParam, protoCustomParamFlag, "", "custom parameters to the specific plugin, <key>=<value> separated by ',' ")
}
