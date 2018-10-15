# README #

This README would normally document whatever steps are necessary to get your application up and running.
This file is not updated, please refer to README.md for latest version

### What is this repository for?

* Quick summary: This repo is a plugin that auto generates boiler template frontend (Vue code with TypeScript) and backend code (java and go) for API services.
  
* Version 0.1


### Getting Started ###

* go to the place you want to clone the repo:
    * `git clone https://github.com/yoozoo/protoapi.git`

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

* set up env:
    * rename the file to `protoc-gen-ts`, in order to be able to use the ts generator plugin
    * or soft link `protoapi` with `protoc-gen-ts` in $GOBIN: `ln -s $PATH-TO-GENERATED-FILE protoc-gen-ts`

* run the plugin:

    * for Mac Users:
        * `protoc --ts_out :. test/hello.proto`:  generate TS files
        * `protoc --ts_out=lang=spring:. ./test/hello.proto`: generate spring files

    * for Windows users:
        * if no softlink: `protoc --plugin=protoc-gen-ts=C:\Users\Admin\go\src\protoapi\protoapi.exe --ts_out=. .\test\protoconf\apps.proto`
        * if with softlink or renamed to `protoc-gen-ts`: `protoc --ts_out . .\test\protoconf\apps.proto`
        * generate spring files: `protoc --ts_out=lang=spring:. ./test/hello.proto`
        * how to make a softlink(command): `mklink <Link to be created> <Target file>`; e.g. `mklink C:\Users\Admin\go\bin\protoc-gen-ts.exe C:\Users\Admin\go\src\protoapi\protoapi.exe`

    * if you have other `.proto` files to test, just change the cli to: `protoc -I=. --ts_out=lang=ts:. $SRC_DIR/$TEST_FILE.proto`

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
