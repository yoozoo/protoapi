# README #

This README would normally document whatever steps are necessary to get your application up and running.

### What is this repository for? ###

* Quick summary: This repo is a plugin that auto generates boiler template Vue code with TypeScript for API services.
* Version 0.1

### Getting Started ###

1. go to the place you want to clone the repo: 
* `git clone https://zhuql@bitbucket.org/yoozoosg/protoc-gen-ts.git`
2. Install protoc-gen-go:
* go to `/usr/local` directory (same dir with .bash_profile for Mac)
* run cli: `go get -u github.com/golang/protobuf/protoc-gen-go` => installed in $GOBIN, defaulting to $GOPATH/bin
3. Make sure protoc is installed in correct folder: 
* [follow the steps here](http://google.github.io/proto-lens/installing-protoc.html)
4. Get other relevant lib:
* run cli: `go get` in the cloned repo folder
5. build the plugin: 
* run cli: `go build` in the cloned repo folder
* make sure a `protoc-gen-ts.exe` file is created in the repo root directory
6. run the plugin:
* `protoc -I=. --plugin=./protoc-gen-ts --ts_out=lang=ts:. test/hello.proto`

### Contribution guidelines ###

* Writing tests
* Code review
* Other guidelines

### Who do I talk to? ###

* Repo owner or admin - [Qinglei](ZHUQL@YOOZOO.COM)
* Other community or team contact - [WenTian](WengW@yoozoo.com)