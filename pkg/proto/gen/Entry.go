// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package gen

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Entry struct {
	_tab flatbuffers.Table
}

func GetRootAsEntry(buf []byte, offset flatbuffers.UOffsetT) *Entry {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Entry{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsEntry(buf []byte, offset flatbuffers.UOffsetT) *Entry {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Entry{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Entry) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Entry) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Entry) Id() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Entry) Title() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Entry) Password() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Entry) CreatedAt() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Entry) MutateCreatedAt(n int64) bool {
	return rcv._tab.MutateInt64Slot(10, n)
}

func (rcv *Entry) UpdatedAt() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Entry) MutateUpdatedAt(n int64) bool {
	return rcv._tab.MutateInt64Slot(12, n)
}

func EntryStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func EntryAddId(builder *flatbuffers.Builder, id flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(id), 0)
}
func EntryAddTitle(builder *flatbuffers.Builder, title flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(title), 0)
}
func EntryAddPassword(builder *flatbuffers.Builder, password flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(password), 0)
}
func EntryAddCreatedAt(builder *flatbuffers.Builder, createdAt int64) {
	builder.PrependInt64Slot(3, createdAt, 0)
}
func EntryAddUpdatedAt(builder *flatbuffers.Builder, updatedAt int64) {
	builder.PrependInt64Slot(4, updatedAt, 0)
}
func EntryEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
