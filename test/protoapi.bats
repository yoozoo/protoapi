#!/usr/bin/env bats

@test "test.proto echo output" {
  ../protoapi gen --lang=echo result/ proto/test.proto
  diff -I -r result/yoozoo_protoconf_ts/ expected/yoozoo_protoconf_ts/
}

@test "test.proto ts output" {
  ../protoapi gen --lang=ts result/ proto/test.proto
  diff -I -r result/yoozoo/protoconf/ts/ expected/yoozoo/protoconf/ts/
}

@test "test.proto spring outout" {
  ../protoapi gen --lang=spring result/ proto/test.proto
  diff -I -r result/com/yoozoo/spring/ expected/com/yoozoo/spring/
}

@test "test.proto php outout" {
  ../protoapi gen --lang=php result/ proto/test.proto
  diff -I -r result/yoozoo.protoconf.ts/ expected/yoozoo.protoconf.ts/
}
