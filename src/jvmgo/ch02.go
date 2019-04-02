package main

import (
	"fmt"
	"jvmgo/classpath"
	"strings"
)

func mainCh02() {
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println("version 0.0.2")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM02(cmd)
	}
}

func startJVM02(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n", cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}

	fmt.Printf("class data:%v\n", classData)
}
