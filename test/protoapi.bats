#!/usr/bin/env bats

@test "test.proto echo output" {
  ../protoapi --lang=echo:result/ proto/test.proto
  diff -I -r result/yoozoo_protoconf_ts/ expected/yoozoo_protoconf_ts/
}

@test "test.proto ts output" {
  ../protoapi --lang=ts:result/ proto/test.proto
  diff -I -r result/yoozoo/protoconf/ts/ expected/yoozoo/protoconf/ts/
}

@test "test.proto spring outout" {
  ../protoapi --lang=spring:result/ proto/test.proto
  diff -I -r result/com/yoozoo/spring/ expected/com/yoozoo/spring/
}
