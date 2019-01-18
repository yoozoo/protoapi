// +build debug

package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/yoozoo/protoapi/generator"

	"github.com/spf13/cobra"
)

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:   "code <file>",
	Short: "read binary protobuf ast data and print generated code to stdout",
	Long: `a command to assist in testing code generation. It reads binary protobuf ast data, generate
	code accordingly and print the generated file content with file name to stdout. The binary data file can be
	created by protoc-gen-dump. When dumping, u need to supply proper parameters on the protoc command line which
	is then being embeded into the data file for this command to use`,
	Args: cobra.ExactArgs(1),
	Run:  outputCode,
}

func outputCode(cmd *cobra.Command, args []string) {

	input, err := ioutil.ReadFile(args[0])
	if err != nil {
		panic(fmt.Errorf("reading file %s error: %s", args[0], err))
	}

	response := generator.Generate(input)

	for _, file := range response.File {
		fmt.Println(*file.Name)
		fmt.Println("-----------------------------------------------------")
		fmt.Println(*file.Content)
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
	}
}

func init() {
	RootCmd.AddCommand(codeCmd)
}
