package zerorpc

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *EventHeader) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "message_id":
			z.Id, err = dc.ReadString()
			if err != nil {
				return
			}
		case "v":
			z.Version, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "response_to":
			z.ResponseTo, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z EventHeader) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "message_id"
	err = en.Append(0x83, 0xaa, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Id)
	if err != nil {
		return
	}
	// write "v"
	err = en.Append(0xa1, 0x76)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.Version)
	if err != nil {
		return
	}
	// write "response_to"
	err = en.Append(0xab, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x6f)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ResponseTo)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z EventHeader) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "message_id"
	o = append(o, 0x83, 0xaa, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64)
	o = msgp.AppendString(o, z.Id)
	// string "v"
	o = append(o, 0xa1, 0x76)
	o = msgp.AppendInt(o, z.Version)
	// string "response_to"
	o = append(o, 0xab, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x6f)
	o = msgp.AppendString(o, z.ResponseTo)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *EventHeader) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "message_id":
			z.Id, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "v":
			z.Version, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "response_to":
			z.ResponseTo, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z EventHeader) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.Id) + 2 + msgp.IntSize + 12 + msgp.StringPrefixSize + len(z.ResponseTo)
	return
}
