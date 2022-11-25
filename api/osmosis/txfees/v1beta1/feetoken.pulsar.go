// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package txfeesv1beta1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_FeeToken        protoreflect.MessageDescriptor
	fd_FeeToken_denom  protoreflect.FieldDescriptor
	fd_FeeToken_poolID protoreflect.FieldDescriptor
)

func init() {
	file_osmosis_txfees_v1beta1_feetoken_proto_init()
	md_FeeToken = File_osmosis_txfees_v1beta1_feetoken_proto.Messages().ByName("FeeToken")
	fd_FeeToken_denom = md_FeeToken.Fields().ByName("denom")
	fd_FeeToken_poolID = md_FeeToken.Fields().ByName("poolID")
}

var _ protoreflect.Message = (*fastReflection_FeeToken)(nil)

type fastReflection_FeeToken FeeToken

func (x *FeeToken) ProtoReflect() protoreflect.Message {
	return (*fastReflection_FeeToken)(x)
}

func (x *FeeToken) slowProtoReflect() protoreflect.Message {
	mi := &file_osmosis_txfees_v1beta1_feetoken_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_FeeToken_messageType fastReflection_FeeToken_messageType
var _ protoreflect.MessageType = fastReflection_FeeToken_messageType{}

type fastReflection_FeeToken_messageType struct{}

func (x fastReflection_FeeToken_messageType) Zero() protoreflect.Message {
	return (*fastReflection_FeeToken)(nil)
}
func (x fastReflection_FeeToken_messageType) New() protoreflect.Message {
	return new(fastReflection_FeeToken)
}
func (x fastReflection_FeeToken_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_FeeToken
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_FeeToken) Descriptor() protoreflect.MessageDescriptor {
	return md_FeeToken
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_FeeToken) Type() protoreflect.MessageType {
	return _fastReflection_FeeToken_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_FeeToken) New() protoreflect.Message {
	return new(fastReflection_FeeToken)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_FeeToken) Interface() protoreflect.ProtoMessage {
	return (*FeeToken)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_FeeToken) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Denom != "" {
		value := protoreflect.ValueOfString(x.Denom)
		if !f(fd_FeeToken_denom, value) {
			return
		}
	}
	if x.PoolID != uint64(0) {
		value := protoreflect.ValueOfUint64(x.PoolID)
		if !f(fd_FeeToken_poolID, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_FeeToken) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "osmosis.txfees.v1beta1.FeeToken.denom":
		return x.Denom != ""
	case "osmosis.txfees.v1beta1.FeeToken.poolID":
		return x.PoolID != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: osmosis.txfees.v1beta1.FeeToken"))
		}
		panic(fmt.Errorf("message osmosis.txfees.v1beta1.FeeToken does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_FeeToken) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "osmosis.txfees.v1beta1.FeeToken.denom":
		x.Denom = ""
	case "osmosis.txfees.v1beta1.FeeToken.poolID":
		x.PoolID = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: osmosis.txfees.v1beta1.FeeToken"))
		}
		panic(fmt.Errorf("message osmosis.txfees.v1beta1.FeeToken does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_FeeToken) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "osmosis.txfees.v1beta1.FeeToken.denom":
		value := x.Denom
		return protoreflect.ValueOfString(value)
	case "osmosis.txfees.v1beta1.FeeToken.poolID":
		value := x.PoolID
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: osmosis.txfees.v1beta1.FeeToken"))
		}
		panic(fmt.Errorf("message osmosis.txfees.v1beta1.FeeToken does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_FeeToken) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "osmosis.txfees.v1beta1.FeeToken.denom":
		x.Denom = value.Interface().(string)
	case "osmosis.txfees.v1beta1.FeeToken.poolID":
		x.PoolID = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: osmosis.txfees.v1beta1.FeeToken"))
		}
		panic(fmt.Errorf("message osmosis.txfees.v1beta1.FeeToken does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_FeeToken) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "osmosis.txfees.v1beta1.FeeToken.denom":
		panic(fmt.Errorf("field denom of message osmosis.txfees.v1beta1.FeeToken is not mutable"))
	case "osmosis.txfees.v1beta1.FeeToken.poolID":
		panic(fmt.Errorf("field poolID of message osmosis.txfees.v1beta1.FeeToken is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: osmosis.txfees.v1beta1.FeeToken"))
		}
		panic(fmt.Errorf("message osmosis.txfees.v1beta1.FeeToken does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_FeeToken) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "osmosis.txfees.v1beta1.FeeToken.denom":
		return protoreflect.ValueOfString("")
	case "osmosis.txfees.v1beta1.FeeToken.poolID":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: osmosis.txfees.v1beta1.FeeToken"))
		}
		panic(fmt.Errorf("message osmosis.txfees.v1beta1.FeeToken does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_FeeToken) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in osmosis.txfees.v1beta1.FeeToken", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_FeeToken) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_FeeToken) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_FeeToken) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_FeeToken) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*FeeToken)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Denom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.PoolID != 0 {
			n += 1 + runtime.Sov(uint64(x.PoolID))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*FeeToken)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.PoolID != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.PoolID))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Denom) > 0 {
			i -= len(x.Denom)
			copy(dAtA[i:], x.Denom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Denom)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*FeeToken)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: FeeToken: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: FeeToken: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Denom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
				}
				x.PoolID = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.PoolID |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: osmosis/txfees/v1beta1/feetoken.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// FeeToken is a struct that specifies a coin denom, and pool ID pair.
// This marks the token as eligible for use as a tx fee asset in Osmosis.
// Its price in osmo is derived through looking at the provided pool ID.
// The pool ID must have osmo as one of its assets.
type FeeToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Denom  string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	PoolID uint64 `protobuf:"varint,2,opt,name=poolID,proto3" json:"poolID,omitempty"`
}

func (x *FeeToken) Reset() {
	*x = FeeToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_osmosis_txfees_v1beta1_feetoken_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeeToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeeToken) ProtoMessage() {}

// Deprecated: Use FeeToken.ProtoReflect.Descriptor instead.
func (*FeeToken) Descriptor() ([]byte, []int) {
	return file_osmosis_txfees_v1beta1_feetoken_proto_rawDescGZIP(), []int{0}
}

func (x *FeeToken) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *FeeToken) GetPoolID() uint64 {
	if x != nil {
		return x.PoolID
	}
	return 0
}

var File_osmosis_txfees_v1beta1_feetoken_proto protoreflect.FileDescriptor

var file_osmosis_txfees_v1beta1_feetoken_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73, 0x2f, 0x74, 0x78, 0x66, 0x65, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x66, 0x65, 0x65, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73,
	0x2e, 0x74, 0x78, 0x66, 0x65, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a,
	0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x08, 0x46, 0x65, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x26, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x10, 0xf2, 0xde, 0x1f, 0x0c, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x64, 0x65, 0x6e, 0x6f,
	0x6d, 0x22, 0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x2a, 0x0a, 0x06, 0x70, 0x6f, 0x6f,
	0x6c, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x12, 0xf2, 0xde, 0x1f, 0x0e, 0x79,
	0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x22, 0x52, 0x06, 0x70,
	0x6f, 0x6f, 0x6c, 0x49, 0x44, 0x3a, 0x04, 0xe8, 0xa0, 0x1f, 0x01, 0x42, 0xf3, 0x01, 0x0a, 0x1a,
	0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73, 0x2e, 0x74, 0x78, 0x66, 0x65,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0d, 0x46, 0x65, 0x65, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73, 0x2d,
	0x6c, 0x61, 0x62, 0x73, 0x2f, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x33,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73, 0x2f, 0x74, 0x78, 0x66,
	0x65, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x74, 0x78, 0x66, 0x65,
	0x65, 0x73, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x4f, 0x54, 0x58, 0xaa,
	0x02, 0x16, 0x4f, 0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73, 0x2e, 0x54, 0x78, 0x66, 0x65, 0x65, 0x73,
	0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x16, 0x4f, 0x73, 0x6d, 0x6f, 0x73,
	0x69, 0x73, 0x5c, 0x54, 0x78, 0x66, 0x65, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0xe2, 0x02, 0x22, 0x4f, 0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73, 0x5c, 0x54, 0x78, 0x66, 0x65,
	0x65, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x18, 0x4f, 0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73,
	0x3a, 0x3a, 0x54, 0x78, 0x66, 0x65, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_osmosis_txfees_v1beta1_feetoken_proto_rawDescOnce sync.Once
	file_osmosis_txfees_v1beta1_feetoken_proto_rawDescData = file_osmosis_txfees_v1beta1_feetoken_proto_rawDesc
)

func file_osmosis_txfees_v1beta1_feetoken_proto_rawDescGZIP() []byte {
	file_osmosis_txfees_v1beta1_feetoken_proto_rawDescOnce.Do(func() {
		file_osmosis_txfees_v1beta1_feetoken_proto_rawDescData = protoimpl.X.CompressGZIP(file_osmosis_txfees_v1beta1_feetoken_proto_rawDescData)
	})
	return file_osmosis_txfees_v1beta1_feetoken_proto_rawDescData
}

var file_osmosis_txfees_v1beta1_feetoken_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_osmosis_txfees_v1beta1_feetoken_proto_goTypes = []interface{}{
	(*FeeToken)(nil), // 0: osmosis.txfees.v1beta1.FeeToken
}
var file_osmosis_txfees_v1beta1_feetoken_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_osmosis_txfees_v1beta1_feetoken_proto_init() }
func file_osmosis_txfees_v1beta1_feetoken_proto_init() {
	if File_osmosis_txfees_v1beta1_feetoken_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_osmosis_txfees_v1beta1_feetoken_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeeToken); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_osmosis_txfees_v1beta1_feetoken_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_osmosis_txfees_v1beta1_feetoken_proto_goTypes,
		DependencyIndexes: file_osmosis_txfees_v1beta1_feetoken_proto_depIdxs,
		MessageInfos:      file_osmosis_txfees_v1beta1_feetoken_proto_msgTypes,
	}.Build()
	File_osmosis_txfees_v1beta1_feetoken_proto = out.File
	file_osmosis_txfees_v1beta1_feetoken_proto_rawDesc = nil
	file_osmosis_txfees_v1beta1_feetoken_proto_goTypes = nil
	file_osmosis_txfees_v1beta1_feetoken_proto_depIdxs = nil
}
