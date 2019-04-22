package base

import "jvmgo/rtda"

func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().PC()
	frame.SetNextPC(pc + offset)
}
