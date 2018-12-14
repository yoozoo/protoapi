.PHONY: clean copy gen

clean:
	rm -rf expected/go/*
	rm -rf expected/package/go/*

copy:
	rm -rf expected/*
	cp -r result/* expected

gen:
	../protoapi gen --lang=go expected/go proto/calc.proto
	../protoapi gen --lang=go expected/go proto/test.proto
	../protoapi gen --lang=go expected/go proto/echo.proto
	../protoapi gen --lang=go expected/go proto/todolist.proto
	../protoapi gen --lang=yii2 expected/ proto/todolist.proto
	../protoapi gen --lang=ts expected/ts proto/test.proto
	../protoapi gen --lang=ts-fetch expected/ts/fetch proto/test.proto
	../protoapi gen --lang=ts-axios expected/ts/axios proto/test.proto
	../protoapi gen --lang=phpclient expected/ proto/test.proto
	../protoapi gen --lang=spring expected/ proto/test.proto

pkg:
	../protoapi gen --lang=go expected/package/go proto/package/common.proto
	../protoapi gen --lang=go expected/package/go proto/package/gopackage_addReqFull.proto
	../protoapi gen --lang=go expected/package/go proto/package/gopackage_addReq.proto
	../protoapi gen --lang=go expected/package/go proto/package/gopackage_calcFull.proto
	../protoapi gen --lang=go expected/package/go proto/package/gopackage_calc.proto
	../protoapi gen --lang=go expected/package/go proto/package/gopackage_calc_warn.proto
	../protoapi gen --lang=go expected/package/go proto/package/mixpackage_addReq.proto
	../protoapi gen --lang=go expected/package/go proto/package/mixpackage_calc.proto
	../protoapi gen --lang=go expected/package/go proto/package/nopackage_calc.proto
	../protoapi gen --lang=go expected/package/go proto/package/nopackage_calc_warn.proto
	../protoapi gen --lang=go expected/package/go proto/package/package_addReq.proto
	../protoapi gen --lang=go expected/package/go proto/package/package_calc_commonerror.proto
	../protoapi gen --lang=go expected/package/go proto/package/package_calc.proto
	../protoapi gen --lang=go expected/package/go proto/package/package_calc._without_commonerror.proto
