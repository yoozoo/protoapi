package main

type outputFunc func(applicationName string, packageName string) string

// OutputMap the registra for output code type and its associated output plugin

func getFunc(prefix string) outputFunc {
	return func(applicationName string, packageName string) string {
		return prefix + applicationName + packageName
	}
}

func main() {
	f1 := getFunc("f1")
	f2 := getFunc("f2")

	println(f1("a1", "p"))
	println(f2("a2", "p"))
	println(f1("a3", "p"))
}
