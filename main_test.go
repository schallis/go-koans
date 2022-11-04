package main

import (
	"fmt"
	"testing"
)

func TestBasics(t *testing.T) {
	aboutBasics()
	aboutStrings()
	aboutArrays()
	aboutSlices()
	aboutTypes()
	aboutControlFlow()
	aboutEnumeration()
	aboutAnonymousFunctions()
	aboutVariadicFunctions()
	aboutDefer()
	aboutFiles()
	aboutInterfaces()
	aboutCommonInterfaces()
	aboutMaps()
	aboutPointers()
	aboutStructs()
	aboutAllocation()
	aboutChannels()
	aboutConcurrency()
	aboutPanics()

	fmt.Printf("\n%c[32;1mYou won life. Good job.%c[0m\n\n", 27, 27)
}
