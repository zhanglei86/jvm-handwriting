package main

import (
	"fmt"
	"jvmgo/classfile"
	"jvmgo/classpath"
	"strings"
)

func mainCh05() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.5")
	} else {
		startJVM05(cmd)
	}
}

func startJVM05(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}
