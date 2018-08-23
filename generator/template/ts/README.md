# Get Started

## How to Test Generated TS code using Protoapi

0. assume you have `protoc` installed and properly set up in your PC
1. generate code: `protoc --ts_out . .\test\protoconf\apps.proto`
2. go into the generated project folder: `cd yoozoo/protoconf/ts/src`
3. copy and paste the API code into your project for use