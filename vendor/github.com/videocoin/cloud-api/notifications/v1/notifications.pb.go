// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: notifications/v1/notifications.proto

package v1

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type NotificationTarget int32

const (
	NotificationTarget_NULL  NotificationTarget = 0
	NotificationTarget_EMAIL NotificationTarget = 1
	NotificationTarget_WEB   NotificationTarget = 2
)

var NotificationTarget_name = map[int32]string{
	0: "NULL",
	1: "EMAIL",
	2: "WEB",
}

var NotificationTarget_value = map[string]int32{
	"NULL":  0,
	"EMAIL": 1,
	"WEB":   2,
}

func (x NotificationTarget) String() string {
	return proto.EnumName(NotificationTarget_name, int32(x))
}

func (NotificationTarget) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bd21247c5a3de394, []int{0}
}

type Notification struct {
	Target               NotificationTarget `protobuf:"varint,1,opt,name=target,proto3,enum=cloud.api.notifications.v1.NotificationTarget" json:"target,omitempty"`
	Template             string             `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
	Params               map[string]string  `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd21247c5a3de394, []int{0}
}
func (m *Notification) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return m.Size()
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetTarget() NotificationTarget {
	if m != nil {
		return m.Target
	}
	return NotificationTarget_NULL
}

func (m *Notification) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

func (m *Notification) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

func (*Notification) XXX_MessageName() string {
	return "cloud.api.notifications.v1.Notification"
}
func init() {
	proto.RegisterEnum("cloud.api.notifications.v1.NotificationTarget", NotificationTarget_name, NotificationTarget_value)
	golang_proto.RegisterEnum("cloud.api.notifications.v1.NotificationTarget", NotificationTarget_name, NotificationTarget_value)
	proto.RegisterType((*Notification)(nil), "cloud.api.notifications.v1.Notification")
	golang_proto.RegisterType((*Notification)(nil), "cloud.api.notifications.v1.Notification")
	proto.RegisterMapType((map[string]string)(nil), "cloud.api.notifications.v1.Notification.ParamsEntry")
	golang_proto.RegisterMapType((map[string]string)(nil), "cloud.api.notifications.v1.Notification.ParamsEntry")
}

func init() {
	proto.RegisterFile("notifications/v1/notifications.proto", fileDescriptor_bd21247c5a3de394)
}
func init() {
	golang_proto.RegisterFile("notifications/v1/notifications.proto", fileDescriptor_bd21247c5a3de394)
}

var fileDescriptor_bd21247c5a3de394 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc9, 0xcb, 0x2f, 0xc9,
	0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x2f, 0x33, 0xd4, 0x47, 0x11, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4a, 0xce, 0xc9, 0x2f, 0x4d, 0xd1, 0x4b, 0x2c, 0xc8,
	0xd4, 0x43, 0x95, 0x2e, 0x33, 0x94, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce,
	0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x6b, 0x49, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c,
	0x30, 0x0b, 0x62, 0x94, 0xd2, 0x6f, 0x46, 0x2e, 0x1e, 0x3f, 0x24, 0x33, 0x84, 0xdc, 0xb8, 0xd8,
	0x4a, 0x12, 0x8b, 0xd2, 0x53, 0x4b, 0x24, 0x18, 0x15, 0x18, 0x35, 0xf8, 0x8c, 0xf4, 0xf4, 0x70,
	0x5b, 0xa6, 0x87, 0xac, 0x33, 0x04, 0xac, 0x2b, 0x08, 0xaa, 0x5b, 0x48, 0x8a, 0x8b, 0xa3, 0x24,
	0x35, 0xb7, 0x20, 0x27, 0xb1, 0x24, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83, 0x33, 0x08, 0xce, 0x17,
	0xf2, 0xe1, 0x62, 0x2b, 0x48, 0x2c, 0x4a, 0xcc, 0x2d, 0x96, 0x60, 0x56, 0x60, 0xd6, 0xe0, 0x36,
	0x32, 0x21, 0xd6, 0x0e, 0xbd, 0x00, 0xb0, 0x36, 0xd7, 0xbc, 0x92, 0xa2, 0xca, 0x20, 0xa8, 0x19,
	0x52, 0x96, 0x5c, 0xdc, 0x48, 0xc2, 0x42, 0x02, 0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x60, 0xd7, 0x73,
	0x06, 0x81, 0x98, 0x42, 0x22, 0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0x30, 0x77, 0x40, 0x38, 0x56,
	0x4c, 0x16, 0x8c, 0x5a, 0x46, 0x5c, 0x42, 0x98, 0x5e, 0x10, 0xe2, 0xe0, 0x62, 0xf1, 0x0b, 0xf5,
	0xf1, 0x11, 0x60, 0x10, 0xe2, 0xe4, 0x62, 0x75, 0xf5, 0x75, 0xf4, 0xf4, 0x11, 0x60, 0x14, 0x62,
	0xe7, 0x62, 0x0e, 0x77, 0x75, 0x12, 0x60, 0x72, 0x92, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23,
	0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x0f, 0x3c, 0x96, 0x63, 0x3c, 0xf1, 0x58, 0x8e, 0x31, 0x8a,
	0xa9, 0xcc, 0x30, 0x89, 0x0d, 0x1c, 0xa4, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x65,
	0x68, 0x18, 0xc5, 0x01, 0x00, 0x00,
}

func (m *Notification) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Notification) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Target != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintNotifications(dAtA, i, uint64(m.Target))
	}
	if len(m.Template) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintNotifications(dAtA, i, uint64(len(m.Template)))
		i += copy(dAtA[i:], m.Template)
	}
	if len(m.Params) > 0 {
		for k, _ := range m.Params {
			dAtA[i] = 0x1a
			i++
			v := m.Params[k]
			mapSize := 1 + len(k) + sovNotifications(uint64(len(k))) + 1 + len(v) + sovNotifications(uint64(len(v)))
			i = encodeVarintNotifications(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintNotifications(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintNotifications(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintNotifications(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Notification) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Target != 0 {
		n += 1 + sovNotifications(uint64(m.Target))
	}
	l = len(m.Template)
	if l > 0 {
		n += 1 + l + sovNotifications(uint64(l))
	}
	if len(m.Params) > 0 {
		for k, v := range m.Params {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovNotifications(uint64(len(k))) + 1 + len(v) + sovNotifications(uint64(len(v)))
			n += mapEntrySize + 1 + sovNotifications(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovNotifications(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNotifications(x uint64) (n int) {
	return sovNotifications(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Notification) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNotifications
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Notification: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Notification: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
			}
			m.Target = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotifications
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Target |= NotificationTarget(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Template", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotifications
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNotifications
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNotifications
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Template = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotifications
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNotifications
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNotifications
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Params == nil {
				m.Params = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowNotifications
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNotifications
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthNotifications
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthNotifications
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNotifications
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthNotifications
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthNotifications
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipNotifications(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthNotifications
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Params[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNotifications(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNotifications
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNotifications
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipNotifications(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNotifications
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNotifications
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNotifications
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthNotifications
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthNotifications
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNotifications
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipNotifications(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthNotifications
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthNotifications = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNotifications   = fmt.Errorf("proto: integer overflow")
)
