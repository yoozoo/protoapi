# protoapi CLI

protoapi command line tool

## First time use

Please run init command to initialize

```bash
protoapi init
```

## Support commands

* root
* init
* gen
* help

## root command

root command will display help information.

## init command

init command will download protoc and other required files into your home directory.

## gen command

gen command will generate user declared language/framework code including API interface and request/response.

To run gen command, use:

```bash
protoapi gen --lang=[language] [output directory] [proto file]
```
