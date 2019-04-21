package heap

import "math"

// 本项目中rtda包已经依赖了heap包，而go语言的包又不能互相依赖，所以本heap包无法导入rtda包，否则编辑器会报错。
// FIXME 临时解决go的互相依赖/引用问题
// 拷贝自rtda的slot和localVars的相关代码。
type SlotCopy struct {
	num int32
	ref *Object
}

type Slots []SlotCopy

func newSlots(slotCount uint) Slots {
	if slotCount > 0 {
		return make([]SlotCopy, slotCount)
	}
	return nil
}

func (self Slots) SetInt(index uint, val int32) {
	self[index].num = val
}
func (self Slots) GetInt(index uint) int32 {
	return self[index].num
}

func (self Slots) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}
func (self Slots) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}

// long consumes two slots
func (self Slots) SetLong(index uint, val int64) {
	self[index].num = int32(val)
	self[index+1].num = int32(val >> 32)
}
func (self Slots) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index+1].num)
	return int64(high)<<32 | int64(low)
}

// double consumes two slots
func (self Slots) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}
func (self Slots) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

func (self Slots) SetRef(index uint, ref *Object) {
	self[index].ref = ref
}
func (self Slots) GetRef(index uint) *Object {
	return self[index].ref
}
