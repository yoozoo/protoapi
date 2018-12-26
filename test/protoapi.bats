#!/usr/bin/env bats

@test "test.proto go output" {
  ../protoapi gen --lang=go result/go proto/test.proto
  ../protoapi gen --lang=go result/go proto/echo.proto
  ../protoapi gen --lang=go result/go proto/calc.proto
  ../protoapi gen --lang=go result/go proto/todolist.proto

  diff -I "^//.*$" -r result/go/ expected/go/
}

@test "packagetest.proto go output" {
  ../protoapi gen --lang=go result/package/go proto/package/common.proto
  ../protoapi gen --lang=go result/package/go proto/package/gopackage_addReqFull.proto
  ../protoapi gen --lang=go result/package/go proto/package/gopackage_addReq.proto
  ../protoapi gen --lang=go result/package/go proto/package/gopackage_calcFull.proto
  ../protoapi gen --lang=go result/package/go proto/package/gopackage_calc.proto
  ../protoapi gen --lang=go result/package/go proto/package/gopackage_calc_warn.proto
  ../protoapi gen --lang=go result/package/go proto/package/mixpackage_addReq.proto
  ../protoapi gen --lang=go result/package/go proto/package/mixpackage_calc.proto
  ../protoapi gen --lang=go result/package/go proto/package/nopackage_calc.proto
  ../protoapi gen --lang=go result/package/go proto/package/nopackage_calc_warn.proto
  ../protoapi gen --lang=go result/package/go proto/package/package_addReq.proto
  ../protoapi gen --lang=go result/package/go proto/package/package_calc_commonerror.proto
  ../protoapi gen --lang=go result/package/go proto/package/package_calc.proto
  ../protoapi gen --lang=go result/package/go proto/package/package_calc._without_commonerror.proto

  diff -r result/package/go/ expected/package/go/
}

@test "test.proto ts output" {
  ../protoapi gen --lang=ts result/ts proto/test.proto
  ../protoapi gen --lang=ts-fetch result/ts/fetch proto/test.proto
  ../protoapi gen --lang=ts-axios result/ts/axios proto/test.proto
  diff -I "^//.*$" -r result/ts/ expected/ts/
}

@test "test.proto spring output" {
  ../protoapi gen --lang=spring result/ proto/test.proto
  diff -I "^//.*$" -r result/com/yoozoo/spring/ expected/com/yoozoo/spring/
}

@test "test.proto phpclient output" {
  ../protoapi gen --lang=phpclient result/ proto/test.proto
  diff -I "^//.*$" -r result/yoozoo.protoconf.ts/ expected/yoozoo.protoconf.ts/
}

@test "todolist.proto php yii2 output" {
  ../protoapi gen --lang=yii2 result/ proto/todolist.proto
  diff -I "^//.*$" -r result/app/ expected/app/
}
