package main

import (
	"fmt"
	"jvmgo/classpath"
	"jvmgo/rtda/heap"
	"strings"
)

func mainCh06() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.6")
	} else {
		startJVM06(cmd)
	}
}

func startJVM06(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)

	classLoader := heap.NewClassLoader(cp)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpretCh06(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}
