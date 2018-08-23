# README #

这份文档包含使用protoapi所需条件， 以及如何使用protoapi

## 项目介绍

* 本项目基于`go`的实现
* 目的：自动生成前后端API的基础代码，节省开发时间
    * 前端生成TypeScript的代码
    * 后端目前支持生成java (spring) 和go (echo)的代码
* 当前版本： 0.1
* [版本发布](ChangeLog.md)

## 项目安装 （还在修改）

1. 下载项目代码:
    * `git clone https://version.uuzu.com/Merlion/protoapi.git`

2. 安装`protoc`:
    * [安装步骤请戳这里](http://google.github.io/proto-lens/installing-protoc.html)

3. 下载所需的第三方库:
    * 进入下载下来的项目内: `cd protoapi`
    * 下载所需第三方库: `go get -u`

## 建立执行文件/插件

* 如果是第一次使用， 或有改代码/template， 需要重新生成执行文件， 在项目路径里跑：
    * `go generate` => 引入新的template到tpl.go
    * `go build` => 重新生成执行文件 - `protoapi.exe`
* 如果项目路径内已有`protoapi.exe`，也没有更改任何代码和template， 可跳过此步

## 如何使用插件

#### Mac用户

* 生成前端TypeScript代码: `protoapi --lang=ts:[output_folder] [proto file path]`
* 生成后端Spring代码：`protoapi --lang=spring:[output_folder] [proto file path]`
* 生成后端echo代码：`protoapi --lang=echo:[output_folder] [proto file path]`

例如：
* 生成前端TypeScript代码： `protoapi --lang=ts:. ./test/hello.proto`
* 生成后端Spring代码： `protoapi --lang=spring:. ./test/hello.proto`
* 生成新的文件夹yoozoo/protoconf/ts,包含新生成的TS文件； 文件夹yoozoo/protoconf/spring里包含了新生成的spring文件

#### Windows用户

* 在项目路径里跑： `go get`
* 生成前端TypeScript代码：`protoapi.exe --lang=ts:[output_folder] [proto file path]`
* 生成后端Spring代码： `protoapi.exe --lang=spring:[output_folder] [proto file path]`
* 生成后端echo代码：`protoapi.exe --lang=echo:[output_folder] [proto file path]`

例如：
* `go get`
* `protoapi.exe --lang=ts:. .\test\hello.proto`
* `protoapi.exe --lang=spring:. .\test\hello.proto`
* 生成新的文件夹yoozoo/protoconf/ts,包含新生成的TS文件； 文件夹yoozoo/protoconf/spring里包含了新生成的spring文件

## 项目结构
* generator
    * data 包含数据结构
        * tpl/tpl.go
        * data.go 定义共享的数据结构
    * output 包含代码生成的具体逻辑
        * echo_xx.go 支持生成echo的代码
        * spring_xx.go 支持生成spring的代码
        * vue_ts.go 支持生成vue(使用ts)的代码
    * template 包含所有模板文件
        * ts
            xx.gots TS的模板
            xx.govue Vue的模板
        * echo_xx.gogo go的模板(对应echo)
        * spring_xx.gojava java的模板(对应spring)
    * generator.go 包含一些共享的代码生成函数逻辑
* test 包含所有测试代码生成的proto文件

## 如何参与项目

* 写测试 Writing test
* 代码审查 Code review
* 添加新的template

### 添加新的template

1. 在generator/template文件夹里添加新的template， 具体语法可参考[这里](https://golang.org/pkg/text/template/)
2. 在generator/output文件夹里添加新的xxx.go文件，包含代码生成的逻辑， 后端代码可参考generator/output/spring.go， 前端代码可参考generator/output/vue_ts.go
3. 参考【建立执行文件/插件】 和 【如何使用插件】测试新加的模板
    * 测试proto文件：protoapi/test/hello.proto
    * 或自定义proto文件

### Proto文件举例

```
syntax = "proto3"; // 使用proto3的语法

option java_package = "com.yoozoo.spring"; //生成的java文件在[output_folder]/com/yoozoo/spring文件夹内

package com.yoozoo.ts; //生成的ts文件在[output_folder]/com/yoozoo/ts文件夹内

import "test/messages.proto"; //引用其他proto文件

// 定义service
service HelloService {
    rpc SayHello (HelloRequest) returns (HelloResponse);
}

```

### 相关资料
1. [go的基本语法和使用](https://golang.org/doc/)
2. [protobuf(proto3)基本语法](https://developers.google.com/protocol-buffers/docs/proto3)
3. [template的基本语法](https://golang.org/pkg/text/template/)
4. [spring](https://spring.io/guides)


## 项目负责人/联系人

- [Qinglei](ZHUQL@YOOZOO.COM)
- [WenTian](WengW@yoozoo.com)
- [HongBo](WuHongbo@yoozoo.com)