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

// DecodeMsg implements msgp.Decodable
func (z *ServerRequest) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Header":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Header = nil
			} else {
				if z.Header == nil {
					z.Header = new(EventHeader)
				}
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
						z.Header.Id, err = dc.ReadString()
						if err != nil {
							return
						}
					case "v":
						z.Header.Version, err = dc.ReadInt()
						if err != nil {
							return
						}
					case "response_to":
						z.Header.ResponseTo, err = dc.ReadString()
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
			}
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "args":
			var xsz uint32
			xsz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Args) >= int(xsz) {
				z.Args = z.Args[:xsz]
			} else {
				z.Args = make([]interface{}, xsz)
			}
			for xvk := range z.Args {
				z.Args[xvk], err = dc.ReadIntf()
				if err != nil {
					return
				}
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
func (z *ServerRequest) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Header"
	err = en.Append(0x83, 0xa6, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72)
	if err != nil {
		return err
	}
	if z.Header == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		// map header, size 3
		// write "message_id"
		err = en.Append(0x83, 0xaa, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Header.Id)
		if err != nil {
			return
		}
		// write "v"
		err = en.Append(0xa1, 0x76)
		if err != nil {
			return err
		}
		err = en.WriteInt(z.Header.Version)
		if err != nil {
			return
		}
		// write "response_to"
		err = en.Append(0xab, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x6f)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Header.ResponseTo)
		if err != nil {
			return
		}
	}
	// write "name"
	err = en.Append(0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "args"
	err = en.Append(0xa4, 0x61, 0x72, 0x67, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Args)))
	if err != nil {
		return
	}
	for xvk := range z.Args {
		err = en.WriteIntf(z.Args[xvk])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ServerRequest) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Header"
	o = append(o, 0x83, 0xa6, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72)
	if z.Header == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 3
		// string "message_id"
		o = append(o, 0x83, 0xaa, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64)
		o = msgp.AppendString(o, z.Header.Id)
		// string "v"
		o = append(o, 0xa1, 0x76)
		o = msgp.AppendInt(o, z.Header.Version)
		// string "response_to"
		o = append(o, 0xab, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x6f)
		o = msgp.AppendString(o, z.Header.ResponseTo)
	}
	// string "name"
	o = append(o, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "args"
	o = append(o, 0xa4, 0x61, 0x72, 0x67, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Args)))
	for xvk := range z.Args {
		o, err = msgp.AppendIntf(o, z.Args[xvk])
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ServerRequest) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Header":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Header = nil
			} else {
				if z.Header == nil {
					z.Header = new(EventHeader)
				}
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
						z.Header.Id, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "v":
						z.Header.Version, bts, err = msgp.ReadIntBytes(bts)
						if err != nil {
							return
						}
					case "response_to":
						z.Header.ResponseTo, bts, err = msgp.ReadStringBytes(bts)
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
			}
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "args":
			var xsz uint32
			xsz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Args) >= int(xsz) {
				z.Args = z.Args[:xsz]
			} else {
				z.Args = make([]interface{}, xsz)
			}
			for xvk := range z.Args {
				z.Args[xvk], bts, err = msgp.ReadIntfBytes(bts)
				if err != nil {
					return
				}
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

func (z *ServerRequest) Msgsize() (s int) {
	s = 1 + 7
	if z.Header == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Header.Id) + 2 + msgp.IntSize + 12 + msgp.StringPrefixSize + len(z.Header.ResponseTo)
	}
	s += 5 + msgp.StringPrefixSize + len(z.Name) + 5 + msgp.ArrayHeaderSize
	for xvk := range z.Args {
		s += msgp.GuessSize(z.Args[xvk])
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ServerResponse) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Header":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Header = nil
			} else {
				if z.Header == nil {
					z.Header = new(EventHeader)
				}
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
						z.Header.Id, err = dc.ReadString()
						if err != nil {
							return
						}
					case "v":
						z.Header.Version, err = dc.ReadInt()
						if err != nil {
							return
						}
					case "response_to":
						z.Header.ResponseTo, err = dc.ReadString()
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
			}
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "args":
			var xsz uint32
			xsz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Args) >= int(xsz) {
				z.Args = z.Args[:xsz]
			} else {
				z.Args = make([]interface{}, xsz)
			}
			for bzg := range z.Args {
				z.Args[bzg], err = dc.ReadIntf()
				if err != nil {
					return
				}
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
func (z *ServerResponse) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Header"
	err = en.Append(0x83, 0xa6, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72)
	if err != nil {
		return err
	}
	if z.Header == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		// map header, size 3
		// write "message_id"
		err = en.Append(0x83, 0xaa, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Header.Id)
		if err != nil {
			return
		}
		// write "v"
		err = en.Append(0xa1, 0x76)
		if err != nil {
			return err
		}
		err = en.WriteInt(z.Header.Version)
		if err != nil {
			return
		}
		// write "response_to"
		err = en.Append(0xab, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x6f)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Header.ResponseTo)
		if err != nil {
			return
		}
	}
	// write "name"
	err = en.Append(0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "args"
	err = en.Append(0xa4, 0x61, 0x72, 0x67, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Args)))
	if err != nil {
		return
	}
	for bzg := range z.Args {
		err = en.WriteIntf(z.Args[bzg])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ServerResponse) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Header"
	o = append(o, 0x83, 0xa6, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72)
	if z.Header == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 3
		// string "message_id"
		o = append(o, 0x83, 0xaa, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64)
		o = msgp.AppendString(o, z.Header.Id)
		// string "v"
		o = append(o, 0xa1, 0x76)
		o = msgp.AppendInt(o, z.Header.Version)
		// string "response_to"
		o = append(o, 0xab, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x6f)
		o = msgp.AppendString(o, z.Header.ResponseTo)
	}
	// string "name"
	o = append(o, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "args"
	o = append(o, 0xa4, 0x61, 0x72, 0x67, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Args)))
	for bzg := range z.Args {
		o, err = msgp.AppendIntf(o, z.Args[bzg])
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ServerResponse) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Header":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Header = nil
			} else {
				if z.Header == nil {
					z.Header = new(EventHeader)
				}
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
						z.Header.Id, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "v":
						z.Header.Version, bts, err = msgp.ReadIntBytes(bts)
						if err != nil {
							return
						}
					case "response_to":
						z.Header.ResponseTo, bts, err = msgp.ReadStringBytes(bts)
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
			}
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "args":
			var xsz uint32
			xsz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Args) >= int(xsz) {
				z.Args = z.Args[:xsz]
			} else {
				z.Args = make([]interface{}, xsz)
			}
			for bzg := range z.Args {
				z.Args[bzg], bts, err = msgp.ReadIntfBytes(bts)
				if err != nil {
					return
				}
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

func (z *ServerResponse) Msgsize() (s int) {
	s = 1 + 7
	if z.Header == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Header.Id) + 2 + msgp.IntSize + 12 + msgp.StringPrefixSize + len(z.Header.ResponseTo)
	}
	s += 5 + msgp.StringPrefixSize + len(z.Name) + 5 + msgp.ArrayHeaderSize
	for bzg := range z.Args {
		s += msgp.GuessSize(z.Args[bzg])
	}
	return
}
