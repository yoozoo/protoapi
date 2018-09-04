package util

import (
	"io/ioutil"
	"os"
	"strings"
)

var protoapiInc string

const embeddedDir = "/proto/"

// GetIncludePath returns the actual include path
func GetIncludePath(userPath string, filePath string) string {
	// extract common includes if needed
	path, err := ioutil.TempDir("", "protoapi_inc_")
	if err != nil {
		panic(err)
	}

	err = ExtractIncludes(path)
	if err != nil {
		panic(err)
	}
	protoapiInc = path

	// return path concatenated with user include path
	var result = protoapiInc + string(os.PathListSeparator) + filePath
	if len(userPath) > 0 {
		result += string(os.PathListSeparator) + userPath
	}
	return result
}

// CleanIncludePath cleanup include path if needed
func CleanIncludePath() {
	// make sure we remove the dir under os.TempDir only
	if len(protoapiInc) > 0 && strings.HasPrefix(protoapiInc, os.TempDir()) {
		os.RemoveAll(protoapiInc)
	}
}

// ExtractIncludes extract common protoconf includes
func ExtractIncludes(path string) error {

	needClean := true
	defer func() {
		if needClean {
			os.RemoveAll(path)
		}
	}()

	for n := range _escData {
		if strings.HasPrefix(n, embeddedDir) {
			name := strings.TrimPrefix(n, embeddedDir)
			data, err := FSByte(false, n)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(path+string(os.PathSeparator)+name, data, 0666)
			if err != nil {
				return err
			}
		}
	}
	needClean = false
	return nil
}
