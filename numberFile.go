package main

import "fmt"

/*
    This is not a comprehensive generator for math functions.
    Its purpose is to provide an example of generating functions
    for up to all 17 number types from just one generator function.
    You're not likely to need functions for uintptr,
    and if you don't, you can just adjust the numbers slice.
    To generate different functions, adjust the funcSlice.
    A full file's worth of content is written to standard output.
*/

func typeToFuncString(typeName string) string {
	typeFunc := []byte(typeName)
	typeFunc[0] -= 32
	return string(typeFunc)
}

func makeSlicedFuncs(types []string, funcSlice []func (string, string) string) {
	for i := 0; i < len(funcSlice); i++ {
		for j := 0; j < len(types); j++ {
			typeName := types[j]
			typeFunc := typeToFuncString(typeName)
			funcString := funcSlice[i](typeFunc, typeName)
			fmt.Print(funcString)
		}
	}
}

func makeImportsString(packageName string, imports []string) string {
	importString := "package " + packageName + "\n\nimport ("
	for i := 0; i < len(imports); i++ {
		importString += "\n\t\"" + imports[i] + "\""
	}
	importString += "\n)\n"
	return importString
}

func isComplex(typeName string) bool {
	return ([]byte(typeName))[0] == 'c'
}

func isUnsigned(typeName string) bool {
	return ([]byte(typeName))[0] == 'u' || ([]byte(typeName))[0] == 'b'
}

func makeFunction(funcNameStart, funcNameType, funcParameters, funcReturn, funcContents string) string {
	return fmt.Sprintf("func %s%s(%s) %s {%s\n}\n\n", funcNameStart, funcNameType, funcParameters, funcReturn, funcContents)
}

func makeAbsFunc(typeFunc, typeName string) string {
	if isComplex(typeName) {
		return makeFunction("Abs", typeFunc, "x " + typeName, "float64", "return cmplx.Abs(complex128(x))")
	} else if isUnsigned(typeName) {
		return ""
	} else {
		return makeFunction("Abs", typeFunc, "x " + typeName, typeName, 
		// Don't mistake the next lines for this file's code; notice the backtick
		`
	if x < 0 {
		return -1 * x
	} else {
		return x
	}`	)  // End of generated function's content
	}
}

func makeMaxFunc(typeFunc, typeName string) string {
	if isComplex(typeName) {
		return ""
	} else {
		return makeFunction("Max", typeFunc, "x, y " + typeName, typeName, 
		// Don't mistake the next lines for this file's code; notice the backtick
		`
	if x > y {
		return x
	} else {
		return y
	}`	) // End of generated function's content
	}
}

func makeMinFunc(typeFunc, typeName string) string {
	if isComplex(typeName) {
		return ""
	} else {
		return makeFunction("Min", typeFunc, "x, y " + typeName, typeName, 
		// Don't mistake the next lines for this file's code; notice the backtick
		`
	if x < y {
		return x
	} else {
		return y
	}`	) // End of generated function's content
	}
}

func main() {
	numbers := []string{"uint8", "byte", "uint16", "uint32", "uint64", "uint", "uintptr", "int8", "int16", "int32", "rune", "int64", "int", "float32", "float64", "complex64", "complex128"}
	funcSlice := []func (string, string) string{makeAbsFunc, makeMaxFunc, makeMinFunc}
	imports := []string {"math/cmplx", "fmt"}
	fmt.Println(makeImportsString("main", imports))
	makeSlicedFuncs(numbers, funcSlice)
	// Just tests the function
	fmt.Println("func main() {\n\tfmt.Println(AbsInt16(-146))\n}")
}