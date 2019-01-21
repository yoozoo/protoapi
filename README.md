[![Build Status](https://travis-ci.com/yoozoo/protoapi.svg?branch=master)](https://travis-ci.com/yoozoo/protoapi)

# README #

这份文档包含使用protoapi所需条件， 以及如何使用protoapi

## 项目介绍

* 本项目基于`go`的实现
* 目的：自动生成前后端API的基础代码，节省开发时间
    * 前端生成TypeScript的代码, 客戶端
    * 后端目前支持生成java (spring) 和go (echo)的代码
* 当前版本： 0.1.1

## 配置环境

安装`go 1.10`以上版本:

* [go安装](https://golang.org/doc/install)
* 需要确保$GOPATH设置正确

## 项目安装

1. 下载项目代码:
  * `go get github.com/yoozoo/protoapi`

2. 使用`protoapi`自动初始化[protoc](https://github.com/protocolbuffers/protobuf):
  * `protoapi init`

## 如何使用protoapi

* `protoapi gen --lang=[language] [output directory] [proto file]`

* 生成前端TypeScript代码: `protoapi gen --lang=ts [output_folder] [proto file path]`
* 生成前端PHP代码：`protoapi gen --lang=php [output_folder] [proto file path]`
* 生成后端Spring代码：`protoapi gen --lang=spring [output_folder] [proto file path]`
* 生成后端echo代码：`protoapi gen --lang=echo [output_folder] [proto file path]`
* 生成后端markdown代码: `protoapi gen --lang=markdown [output_folder] [proto file path]`

例如：
* 生成前端TypeScript代码： `protoapi gen --lang=ts . ./test/proto/todolist.proto`
* 生成后端Spring代码： `protoapi gen --lang=spring . ./test/proto/todolist.proto`

* 其他相关command请参考[这里](docs/protoapi_cli.md)

## 项目结构
* generator
    * data 包含数据结构
        * tpl/tpl.go
        * data.go 定义共享的数据结构
    * output 包含代码生成的具体逻辑
        * echo_xx.go 支持生成echo的代码
        * spring_xx.go 支持生成spring的代码
        * vue_ts.go 支持生成vue(使用ts)的代码
        * php.go 支持生成php的代码
        * markdown.go 支持生成markdown的代码
    * template 包含所有模板文件
        * ts
            xx.gots TS的模板
            xx.govue Vue的模板
        * echo_xx.gogo go的模板(对应echo)
        * spring_xx.gojava java的模板(对应spring)
        * php.gophp php的模板
        * markdown.gomd markdown的模板
    * generator.go 包含一些共享的代码生成函数逻辑
* test 包含所有测试代码生成的proto文件

## 如何参与项目

* 写测试 Writing test
* 代码审查 Code review
* 添加新的template

### 添加新的template

1. 在generator/template文件夹里添加新的template：

* 具体语法可参考[这里](https://golang.org/pkg/text/template/)
* 现有例子可参考generator/template里已有的template
* 新添加template文件根据生成文件命名后缀， 如生成ts文件则命名为：xxx.gots, 生成java文件则叫xxx.gojava等

2. 在generator/output文件夹里添加新的xxx.go文件或改动现有文件的逻辑

* 后端代码可参考generator/output/spring.go
* 前端代码可参考generator/output/vue_ts.go
* 例如，如果想要多生成一个ts文件：
  * 添加新的模板： generator/template/ts/example.gots
  * 在generator/output/ts.go里面

```golang
        type tsGen struct {
            // 1. 添加数据
            ...
            exampleFile string
            ...
            exampleTpl       *template.Template

        }
        ...
        /**
        * Get TEMPLATE
        */
        func (g *tsGen) loadTpl() {
            ...
            // 2. 添加输入的模板
            g.exampleTpl = g.getTpl("/generator/template/ts/example.gots")
        }

        /**
        * init filename with path
        */
        func initFiles(packageName string, service *data.ServiceData) *tsGen {
            gen := &tsGen{
                ...
                // 3. 添加生成的文件名
                // 新生成文件会命名为： example.ts, 并根据packageName指定生成于哪个文件夹
                // 例如： packageName = yoozoo.protoconf.ts的话， 文件会生成与 $output_dir/yoozoo/protoconf/ts
                // packageName 定义于proto文件内： “package yoozoo.protoconf.ts;”
                exampleFile:      genFileName(packageName, "example"),
            }
            return gen
        }

        func generateVueTsCode(applicationName string, packageName string, service *data.ServiceData, messages []*data.MessageData, enums []*data.EnumData, options []*data.Option) (map[string]string, error) {

            ...
            /**
            * combine data with template
            */
            // 4. 最后输出生成的文件
            ...
            result[gen.exampleFile] = gen.genContent(gen.exampleTpl, dataMap)
            ...
        }

```

3. 参考【建立执行文件/插件】 和 【如何使用插件】测试新加的模板
    * 可测试现有的proto样本文件：`protoapi/test/hello.proto`
    * 或自定义proto文件测试

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

### HTTP Method

* 所有API生成默认使用HTTP POST
* 特效场景下可以使用GET，但不鼓励， 因为query string无法很好的对复杂请求对象做序列化

### 错误处理

* [错误处理规范](docs/ErrorHandling.md)

### 认证处理

* [认证处理规范](docs/protoapi_auth_cn.md)

### 相关资料
1. [go的基本语法和使用](https://golang.org/doc/)
2. [protobuf(proto3)基本语法](https://developers.google.com/protocol-buffers/docs/proto3)
3. [template的基本语法](https://golang.org/pkg/text/template/)
4. [spring](https://spring.io/guides)
