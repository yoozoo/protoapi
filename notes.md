1. install
    ```
    PROTOC_ZIP=protoc-3.3.0-osx-x86_64.zip
    curl -OL https://github.com/google/protobuf/releases/download/v3.3.0/$PROTOC_ZIP
    sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
    rm -f $PROTOC_ZIP
    ```
    [reference](http://google.github.io/proto-lens/installing-protoc.html)

2. Run
`protoc -I=$SRC_DIR --java_out=$DST_DIR $SRC_DIR/addressbook.proto`
e.g. `protoc -I=. --java_out=test test/test.proto`


# Issues:

1. [error-using-import-in-proto-file](https://stackoverflow.com/questions/21134066/error-using-import-in-proto-file)

2. cannot build ts
* go build -> [folder].exe -> rename to protoc-gen-[anything].exe
* protoc --[anything]_out=lang=ts:$DST_DIR $SRC_DIR/protofile

3. multiple messages: `We only support one root configuration object and we found more than one: .com.yoozoo.ts.HelloRequest, .com.yoozoo.ts.HelloResponse`

4. The Go code generator does not produce output for services by default
