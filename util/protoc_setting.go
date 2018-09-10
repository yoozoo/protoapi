package util

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

const (
	defaultProtocCmd = "protoc"
	// env name for protoapi home
	protoapiDirEnv = "PROTOAPI_PATH"
	//ProtocBin path for protoc binary under protoapi home
	ProtocBin = "bin/protoc"
	//ProtoapiCommonInclude protoapi common proto file directory
	ProtoapiCommonInclude = "include/protoapi/"

	protocInclude = "include"
)

//ClearDir remove all the files/dirs under a directory
func ClearDir(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

//GetProtoapiHome return protoconf home dir
func GetProtoapiHome() string {
	homedir := os.Getenv(protoapiDirEnv)
	if len(homedir) == 0 {
		if usr, err := user.Current(); err == nil {
			homedir = usr.HomeDir + "/.protoapi"
		}
	}
	if len(homedir) == 0 {
		HandleError(fmt.Errorf("Failed to find protoapi home dir. Please make sure current user has home directory"))
	}
	return homedir + "/"
}

// GetDefaultProtoc retrieve protoc executable path and protoc Include path
func GetDefaultProtoc(incPath string) (protoc string, newProtocIncPath string) {

	homedir := GetProtoapiHome()

	protoc = homedir + ProtocBin
	if runtime.GOOS == "windows" {
		protoc = protoc + ".exe"
	}
	// check existen
	if _, err := os.Stat(protoc); err != nil {
		HandleError(fmt.Errorf("Failed to find protoc. Please run protoapi init command to initialize: %s", err.Error()))
	} else {
		newProtocIncPath = homedir + protocInclude
		if _, err := os.Stat(newProtocIncPath); err != nil {
			HandleError(fmt.Errorf("Failed to find protoc include folder. Please run protoapi init command to initialize: %s", err.Error()))
		}
		if len(incPath) > 0 {
			newProtocIncPath += string(os.PathListSeparator) + incPath
		}
		return protoc, newProtocIncPath
	}
	return "", ""
}
