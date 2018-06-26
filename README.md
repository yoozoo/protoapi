# README #

This README would normally document whatever steps are necessary to get your application up and running.

### What is this repository for? ###

* Quick summary: This repo is a plugin that auto generates boiler template Vue code with TypeScript for API services.
* Version 0.1

### Getting Started ###

* go to the place you want to clone the repo: 
    * `git clone https://zhuql@bitbucket.org/yoozoosg/protoc-gen-ts.git`

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
    * you will see a `[folder-name].exe` file is created in the repo root directory
    * rename the `.exe` file to `protoc-gen-ts`, in order to be able to use the ts generator plugin

* run the plugin:
    * `protoc -I=. --plugin=./protoc-gen-ts --ts_out=lang=ts:. test/hello.proto`
    * if you have other `.proto` files to test, just change the cli to: `protoc -I=. --plugin=./protoc-gen-ts --ts_out=lang=ts:. $SRC_DIR/$TEST_FILE.proto`


### Contribution guidelines ###

* Writing tests
* Code review
* Other guidelines

### Who do I talk to? ###

* Repo owner or admin - [Qinglei](ZHUQL@YOOZOO.COM)
* Other community or team contact - [WenTian](WengW@yoozoo.com)