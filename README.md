# README #

This README would normally document whatever steps are necessary to get your application up and running.

### What is this repository for? ###

* Quick summary: This repo is a plugin that auto generates boiler template Vue code with TypeScript for API services.
* Version 0.1

### Getting Started ###

* go to the place you want to clone the repo:
    * `git clone https://zhuql@bitbucket.org/yoozoosg/protoapi.git`

* Install protoc-gen-go if you have not done so:
    * go to `/usr/local` directory (same dir with .bash_profile for Mac)
    * run cli: `go get -u github.com/golang/protobuf/protoc-gen-go` => installed in $GOBIN, defaulting to $GOPATH/bin

* Make sure protoc is installed in correct folder:
    * [follow the steps here](http://google.github.io/proto-lens/installing-protoc.html)

* Get other relevant lib:
    * go to the cloned project folder: `cd protoapi`
    * run cli: `go get` in the cloned repo folder

* build the plugin:
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

### Contribution guidelines ###

* Writing tests
* Code review
* Other guidelines

### Who do I talk to? ###

* Repo owner or admin
    - [Qinglei](ZHUQL@YOOZOO.COM)
* Other community or team contact
    - [WenTian](WengW@yoozoo.com)
    - [HongBo](WuHongbo@yoozoo.com)

# Todo

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
