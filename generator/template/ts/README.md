# Get Started

## How to Test Generated TS code using Protoconf API

1. generate code: `protoc --ts_out . .\test\protoconf\apps.proto`
2. go into the generated project folder: `cd yoozoo/protoconf/ts`
3. run `npm init`, go through package.json set up accordingly
4. run `npm install` in the generated project folder to install all the required packages, you will see `node_modules` folder, `package-lock.json` and `public/index.html` being created in the generated project folder
5. run `webpack` to generate `dist/bundle.js` (webpack & webpack-cli need to be installed)
6. right click on `public/index.html` and open in chrome, you should be able to see a list of buttons each labeled with the API function name
