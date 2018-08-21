# README #

这份文档包含使用protoapi所需条件， 以及如何使用protoapi

## 项目介绍

* 本项目基于`go`的实现
* 目的：自动生成前后端API的基础代码， 
    * 前端生成TypeScript的代码
    * 后端目前支持生成java (spring) 和go的代码
* 当前版本： 0.1


## 项目安装

* 下载项目代码:
    * `git clone https://version.uuzu.com/Merlion/protoapi.git`

* 安装`protoc`:
    * [安装步骤请戳这里](http://google.github.io/proto-lens/installing-protoc.html)

* 安装 `protoc-gen-go`:
    * go to `/usr/local` directory (same dir with .bash_profile for Mac)
    * run cli: `go get -u github.com/golang/protobuf/protoc-gen-go` => installed in $GOBIN, defaulting to $GOPATH/bin

* 下载所需的第三方库:
    * go to the cloned project folder: `cd protoapi`
    * run cli: `go get` in the cloned repo folder

## 建立插件

* 生成插件:
    * run cli: `go build` in the cloned repo folder
    * you will see a `[folder-name]` file is created in the repo root directory
    * run cli  `go generate` in the cloned repo folder

* run the plugin:

    * for Mac Users:
        * `protoapi --lang=ts:[output_folder] [proto file path]`:  generate TS files
        * `protoapi --lang=spring:[output_folder] [proto file path]`: generate spring files
        * `protoapi --lang=echo:[output_folder] [proto file path]`: generate echo files

    * for Windows users:
        * generate ts files: `protoapi.exe --lang=ts:[output_folder] [proto file path]`
        * generate spring files: `protoapi.exe --lang=spring:[output_folder] [proto file path]`
        * generate echo files: `protoapi.exe --lang=echo:[output_folder] [proto file path]`
* 配置环境:
    * rename the file to `protoc-gen-ts`, in order to be able to use the ts generator plugin
    * or soft link `protoapi` with `protoc-gen-ts` in $GOBIN: `ln -s $PATH-TO-GENERATED-FILE protoc-gen-ts`

## 如何使用插件 

### Mac用户

    * 生成typescript代码： `protoc --ts_out :. test/hello.proto`
    * 生成后端代码： `protoc --ts_out=lang=spring:. ./test/hello.proto`

### Windows用户
    * if no softlink: `protoc --plugin=protoc-gen-ts=C:\Users\Admin\go\src\protoapi\protoapi.exe --ts_out=. .\test\protoconf\apps.proto`
    * if with softlink or renamed to `protoc-gen-ts`: `protoc --ts_out . .\test\protoconf\apps.proto`
    * generate spring files: `protoc --ts_out=lang=spring:. ./test/hello.proto`
    * how to make a softlink(command): `mklink <Link to be created> <Target file>`; e.g. `mklink C:\Users\Admin\go\bin\protoc-gen-ts.exe C:\Users\Admin\go\src\protoapi\protoapi.exe`

    * if you have other `.proto` files to test, just change the cli to: `protoc -I=. --ts_out=lang=ts:. $SRC_DIR/$TEST_FILE.proto`

## 如何参与项目 ###

* Writing tests
* Code review
* Other guidelines

## 项目负责人/联系人

- [Qinglei](ZHUQL@YOOZOO.COM)
- [WenTian](WengW@yoozoo.com)
- [HongBo](WuHongbo@yoozoo.com)

## spring ##

* complex data type support
  * support object data type: it is to be declared in the message.proto. Example:

  ```protobuf
  syntax = "proto3";

  message HelloRequest {
      Greeting greeting = 1;
  }

  message HelloResponse {
      string reply = 1;
  }

  message Greeting {
      string greetingMsg = 1;
  }
  ```

* java package name
  * user can declare java package name as options in the service.proto file. and the java classes will be generated in the specific packages. If no java_package_options is declared, files will be generate to the package in proto file.

  ```protobuf
  option java_package = "com.yoozoo.spring";
  ```

* service name options
