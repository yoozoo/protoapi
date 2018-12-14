.PHONY: clean copy gen

clean:
	rm -rf expected/go/*

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
	../protoapi gen --lang=spring expected/ proto/test.proto
