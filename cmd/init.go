package cmd

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"version.uuzu.com/Merlion/protoapi/util"
)

var protocURLMap = map[string]string{
	// https://developers.google.com/protocol-buffers/docs/proto#scalar
	"darwin":  "https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-osx-x86_64.zip",
	"windows": "https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-win32.zip",
	"linux":   "https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip",
}

// newInitCommand downloads protoc binary and required files into ./protoconf/ folder
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Downloads protoc binary and required files into ./protoapi/ folder",
	Run:   initCommandFunc,
}

var (
	forceInit = false
)

const (
	forceInitFlag = "force"
)

func initCommandFunc(cmd *cobra.Command, args []string) {
	// get protoc download url from os
	protocURL, ok := protocURLMap[runtime.GOOS]

	if !ok {
		util.Die(fmt.Errorf("unsupported system %s", runtime.GOOS))
	}

	workingDir := util.GetProtoapiHome()

	protocInstalled := false
	protocFile := workingDir + util.ProtocBin
	if runtime.GOOS == "windows" {
		protocFile += ".exe"
	}
	if _, err := os.Stat(protocFile); err == nil && !forceInit {
		// protoc alread initialized
		protocInstalled = true
	}
	if !protocInstalled {
		util.ClearDir(workingDir)
		// create working dir
		if _, err := os.Stat(workingDir); os.IsNotExist(err) {
			// path not exist
			err = os.Mkdir(workingDir, os.ModePerm)
			if err != nil {
				util.Die(fmt.Errorf("Failed to create working dir %s: %s", workingDir, err.Error()))
			}
		}

		// download protoc
		err := downloadFile(workingDir+"/protoc.zip", protocURL)
		if err != nil {
			util.Die(fmt.Errorf("Failed to download protoc from %s : %s", protocURL, err.Error()))
		}

		// unzip protoc.zip, it will create bin, include etc
		files, err := unzip(workingDir+"/protoc.zip", workingDir)
		if err != nil {
			util.Die(fmt.Errorf("Failed to unzip %s: %s", workingDir+"/protoc.zip", err.Error()))
		}
		fmt.Println("Downloaded and unzipped:\n" + strings.Join(files, "\n"))
		err = os.Remove(workingDir + "/protoc.zip")
		if err != nil {
			fmt.Printf("\n\nFailed to delete protoc.zip from %s: %s", workingDir+"/protoc.zip", err.Error())
		}
	}
	// write protoapi include file
	protoapiIncPath := workingDir + util.ProtoapiCommonInclude
	if _, err := os.Stat(protoapiIncPath); os.IsNotExist(err) {
		err = os.MkdirAll(protoapiIncPath, os.ModePerm)
		if err != nil {
			util.Die(fmt.Errorf("Failed create directory %s: %s", protoapiIncPath, err))
		}
		err = util.ExtractIncludes(protoapiIncPath)
		if err != nil {
			util.Die(fmt.Errorf("Failed to download protoapi include file into %s: %s", protoapiIncPath, err))
		}
		fmt.Println(filepath.FromSlash(protoapiIncPath + "protoapi_common.proto"))
	}
	fmt.Println("Protoapi initialized.")
}

// downloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func downloadFile(filepath string, url string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// unzip will decompress a zip archive, moving all files and folders
// within the zip file (parameter 1) to an output directory (parameter 2).
func unzip(src string, dest string) ([]string, error) {
	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}
		defer rc.Close()

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)
		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			// Make File
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return filenames, err
			}

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return filenames, err
			}

			_, err = io.Copy(outFile, rc)

			// Close the file without defer to close before next iteration of loop
			outFile.Close()

			if err != nil {
				return filenames, err
			}
		}
	}
	return filenames, nil
}

func init() {
	initCmd.Flags().BoolVar(&forceInit, forceInitFlag, false, "force protoapi initialization even if it is initialized.")
	RootCmd.AddCommand(initCmd)
}
