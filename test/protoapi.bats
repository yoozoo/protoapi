#!/usr/bin/env bats

@test "test.proto go output" {
  skip
  ../protoapi gen --lang=go result/ proto/test.proto
  diff -I "^//.*$" -r result/apisvr/ expected/apisvr/
}

@test "test.proto ts output" {
  ../protoapi gen --lang=ts result/ts proto/test.proto
  ../protoapi gen --lang=ts-fetch result/ts/fetch proto/test.proto
  ../protoapi gen --lang=ts-axios result/ts/axios proto/test.proto
  diff -I "^//.*$" -r result/ts/ expected/ts/
}

@test "test.proto spring outout" {
  ../protoapi gen --lang=spring result/ proto/test.proto
  diff -I "^//.*$" -r result/com/yoozoo/spring/ expected/com/yoozoo/spring/
}

@test "test.proto phpclient outout" {
  ../protoapi gen --lang=phpclient result/ proto/test.proto
  diff -I "^//.*$" -r result/yoozoo.protoconf.ts/ expected/yoozoo.protoconf.ts/
}
