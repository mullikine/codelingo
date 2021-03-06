// Code generated by protoc-gen-go.
// source: github.com/codelingo/codelingo/flows/codelingo/rewrite/rpc/rewrite.proto
// DO NOT EDIT!

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	github.com/codelingo/codelingo/flows/codelingo/rewrite/rpc/rewrite.proto

It has these top-level messages:
	Request
	Hunk
*/
package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the files or directories to review.
type Request struct {
	Host             string   `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
	Hostname         string   `protobuf:"bytes,2,opt,name=hostname" json:"hostname,omitempty"`
	GenerateComments bool     `protobuf:"varint,3,opt,name=generateComments" json:"generateComments,omitempty"`
	Repo             string   `protobuf:"bytes,4,opt,name=repo" json:"repo,omitempty"`
	Sha              string   `protobuf:"bytes,5,opt,name=sha" json:"sha,omitempty"`
	Patches          []string `protobuf:"bytes,6,rep,name=Patches" json:"Patches,omitempty"`
	IsPullRequest    bool     `protobuf:"varint,7,opt,name=isPullRequest" json:"isPullRequest,omitempty"`
	PullRequestID    int64    `protobuf:"varint,8,opt,name=pullRequestID" json:"pullRequestID,omitempty"`
	Vcs              string   `protobuf:"bytes,9,opt,name=vcs" json:"vcs,omitempty"`
	Dotlingo         string   `protobuf:"bytes,10,opt,name=dotlingo" json:"dotlingo,omitempty"`
	Dir              string   `protobuf:"bytes,11,opt,name=dir" json:"dir,omitempty"`
	// Types that are valid to be assigned to OwnerOrDepot:
	//	*Request_Owner
	//	*Request_Depot
	OwnerOrDepot isRequest_OwnerOrDepot `protobuf_oneof:"ownerOrDepot"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isRequest_OwnerOrDepot interface {
	isRequest_OwnerOrDepot()
}

type Request_Owner struct {
	Owner string `protobuf:"bytes,12,opt,name=owner,oneof"`
}
type Request_Depot struct {
	Depot string `protobuf:"bytes,13,opt,name=depot,oneof"`
}

func (*Request_Owner) isRequest_OwnerOrDepot() {}
func (*Request_Depot) isRequest_OwnerOrDepot() {}

func (m *Request) GetOwnerOrDepot() isRequest_OwnerOrDepot {
	if m != nil {
		return m.OwnerOrDepot
	}
	return nil
}

func (m *Request) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Request) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Request) GetGenerateComments() bool {
	if m != nil {
		return m.GenerateComments
	}
	return false
}

func (m *Request) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

func (m *Request) GetSha() string {
	if m != nil {
		return m.Sha
	}
	return ""
}

func (m *Request) GetPatches() []string {
	if m != nil {
		return m.Patches
	}
	return nil
}

func (m *Request) GetIsPullRequest() bool {
	if m != nil {
		return m.IsPullRequest
	}
	return false
}

func (m *Request) GetPullRequestID() int64 {
	if m != nil {
		return m.PullRequestID
	}
	return 0
}

func (m *Request) GetVcs() string {
	if m != nil {
		return m.Vcs
	}
	return ""
}

func (m *Request) GetDotlingo() string {
	if m != nil {
		return m.Dotlingo
	}
	return ""
}

func (m *Request) GetDir() string {
	if m != nil {
		return m.Dir
	}
	return ""
}

func (m *Request) GetOwner() string {
	if x, ok := m.GetOwnerOrDepot().(*Request_Owner); ok {
		return x.Owner
	}
	return ""
}

func (m *Request) GetDepot() string {
	if x, ok := m.GetOwnerOrDepot().(*Request_Depot); ok {
		return x.Depot
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Request) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Request_OneofMarshaler, _Request_OneofUnmarshaler, _Request_OneofSizer, []interface{}{
		(*Request_Owner)(nil),
		(*Request_Depot)(nil),
	}
}

func _Request_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Request)
	// ownerOrDepot
	switch x := m.OwnerOrDepot.(type) {
	case *Request_Owner:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Owner)
	case *Request_Depot:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Depot)
	case nil:
	default:
		return fmt.Errorf("Request.OwnerOrDepot has unexpected type %T", x)
	}
	return nil
}

func _Request_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Request)
	switch tag {
	case 12: // ownerOrDepot.owner
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.OwnerOrDepot = &Request_Owner{x}
		return true, err
	case 13: // ownerOrDepot.depot
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.OwnerOrDepot = &Request_Depot{x}
		return true, err
	default:
		return false, nil
	}
}

func _Request_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Request)
	// ownerOrDepot
	switch x := m.OwnerOrDepot.(type) {
	case *Request_Owner:
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Owner)))
		n += len(x.Owner)
	case *Request_Depot:
		n += proto.SizeVarint(13<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Depot)))
		n += len(x.Depot)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Hunk struct {
	Filename    string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	StartOffset int32  `protobuf:"varint,2,opt,name=startOffset" json:"startOffset,omitempty"`
	EndOffset   int32  `protobuf:"varint,3,opt,name=endOffset" json:"endOffset,omitempty"`
	SRC         string `protobuf:"bytes,4,opt,name=SRC" json:"SRC,omitempty"`
	Comment     string `protobuf:"bytes,5,opt,name=comment" json:"comment,omitempty"`
	// Hack to pass flag from CLIApp to writer
	CommentOutputFile string `protobuf:"bytes,6,opt,name=commentOutputFile" json:"commentOutputFile,omitempty"`
	DecoratorOptions  string `protobuf:"bytes,7,opt,name=decoratorOptions" json:"decoratorOptions,omitempty"`
}

func (m *Hunk) Reset()                    { *m = Hunk{} }
func (m *Hunk) String() string            { return proto.CompactTextString(m) }
func (*Hunk) ProtoMessage()               {}
func (*Hunk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Hunk) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *Hunk) GetStartOffset() int32 {
	if m != nil {
		return m.StartOffset
	}
	return 0
}

func (m *Hunk) GetEndOffset() int32 {
	if m != nil {
		return m.EndOffset
	}
	return 0
}

func (m *Hunk) GetSRC() string {
	if m != nil {
		return m.SRC
	}
	return ""
}

func (m *Hunk) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *Hunk) GetCommentOutputFile() string {
	if m != nil {
		return m.CommentOutputFile
	}
	return ""
}

func (m *Hunk) GetDecoratorOptions() string {
	if m != nil {
		return m.DecoratorOptions
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "rpc.Request")
	proto.RegisterType((*Hunk)(nil), "rpc.Hunk")
}

func init() {
	proto.RegisterFile("github.com/codelingo/codelingo/flows/codelingo/rewrite/rpc/rewrite.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x92, 0xcf, 0x8a, 0xdb, 0x30,
	0x10, 0xc6, 0xeb, 0x38, 0xff, 0xac, 0xfd, 0xc3, 0x56, 0x94, 0x22, 0x4a, 0x0f, 0x66, 0xe9, 0x21,
	0x94, 0xe2, 0x3d, 0xf4, 0x0d, 0x92, 0xa5, 0xa4, 0xb0, 0x90, 0xe0, 0x1e, 0x7a, 0xf6, 0xca, 0x93,
	0x44, 0xd4, 0xf1, 0xa8, 0xd2, 0x78, 0xf3, 0x3e, 0x7d, 0xba, 0x5e, 0xfb, 0x06, 0x65, 0x64, 0x3b,
	0x9b, 0x90, 0x93, 0xbf, 0xef, 0xa7, 0xc1, 0x23, 0x7d, 0x33, 0x62, 0xb9, 0x35, 0xb4, 0x6b, 0x9e,
	0x33, 0x8d, 0xfb, 0x07, 0x8d, 0x25, 0x54, 0xa6, 0xde, 0xe2, 0x89, 0xda, 0x54, 0x78, 0xf0, 0x27,
	0xde, 0xc1, 0xc1, 0x19, 0x82, 0x07, 0x67, 0x75, 0xaf, 0x33, 0xeb, 0x90, 0x50, 0xc6, 0xce, 0xea,
	0xfb, 0x7f, 0x03, 0x31, 0xc9, 0xe1, 0x77, 0x03, 0x9e, 0xa4, 0x14, 0xc3, 0x1d, 0x7a, 0x52, 0x51,
	0x1a, 0xcd, 0x92, 0x3c, 0x68, 0xf9, 0x41, 0x4c, 0xf9, 0x5b, 0x17, 0x7b, 0x50, 0x83, 0xc0, 0x8f,
	0x5e, 0x7e, 0x16, 0x77, 0x5b, 0xa8, 0xc1, 0x15, 0x04, 0x0b, 0xdc, 0xef, 0xa1, 0x26, 0xaf, 0xe2,
	0x34, 0x9a, 0x4d, 0xf3, 0x0b, 0xce, 0xff, 0x76, 0x60, 0x51, 0x0d, 0xdb, 0x7f, 0xb3, 0x96, 0x77,
	0x22, 0xf6, 0xbb, 0x42, 0x8d, 0x02, 0x62, 0x29, 0x95, 0x98, 0xac, 0x0b, 0xd2, 0x3b, 0xf0, 0x6a,
	0x9c, 0xc6, 0xb3, 0x24, 0xef, 0xad, 0xfc, 0x24, 0x6e, 0x8c, 0x5f, 0x37, 0x55, 0xd5, 0x5d, 0x56,
	0x4d, 0x42, 0xa3, 0x73, 0xc8, 0x55, 0xf6, 0xd5, 0x7e, 0x7f, 0x54, 0xd3, 0x34, 0x9a, 0xc5, 0xf9,
	0x39, 0xe4, 0xbe, 0x2f, 0xda, 0xab, 0xa4, 0xed, 0xfb, 0xa2, 0x3d, 0xbf, 0xb2, 0x44, 0x0a, 0x81,
	0x29, 0xd1, 0xbe, 0xb2, 0xf7, 0x5c, 0x5d, 0x1a, 0xa7, 0xae, 0xda, 0xea, 0xd2, 0x38, 0xf9, 0x5e,
	0x8c, 0xf0, 0x50, 0x83, 0x53, 0xd7, 0xcc, 0x96, 0x6f, 0xf2, 0xd6, 0x32, 0x2f, 0xc1, 0x22, 0xa9,
	0x9b, 0x9e, 0x07, 0x3b, 0xbf, 0x15, 0xd7, 0xa1, 0x60, 0xe5, 0x1e, 0xd9, 0xdf, 0xff, 0x8d, 0xc4,
	0x70, 0xd9, 0xd4, 0xbf, 0xb8, 0xed, 0xc6, 0x54, 0x10, 0xc2, 0x6d, 0x43, 0x3f, 0x7a, 0x99, 0x8a,
	0x2b, 0x4f, 0x85, 0xa3, 0xd5, 0x66, 0xe3, 0x81, 0x42, 0xf6, 0xa3, 0xfc, 0x14, 0xc9, 0x8f, 0x22,
	0x81, 0xba, 0xec, 0xce, 0xe3, 0x70, 0xfe, 0x0a, 0xf8, 0xda, 0x3f, 0xf2, 0x45, 0x97, 0x37, 0x4b,
	0x0e, 0x57, 0xb7, 0xe3, 0xe8, 0x22, 0xef, 0xad, 0xfc, 0x22, 0xde, 0x76, 0x72, 0xd5, 0x90, 0x6d,
	0xe8, 0x9b, 0xa9, 0x40, 0x8d, 0x43, 0xcd, 0xe5, 0x01, 0x8f, 0xbd, 0x04, 0x8d, 0xae, 0x20, 0x74,
	0x2b, 0x4b, 0x06, 0x6b, 0x1f, 0xa6, 0x91, 0xe4, 0x17, 0x7c, 0x9e, 0x89, 0x77, 0x06, 0xb3, 0xe3,
	0x36, 0x66, 0x5b, 0x67, 0x75, 0xa6, 0xab, 0xf9, 0xed, 0x02, 0x4b, 0x78, 0x62, 0xb4, 0xe6, 0x5d,
	0x5c, 0x47, 0x7f, 0x06, 0xf1, 0xe2, 0xe9, 0xe7, 0xf3, 0x38, 0xac, 0xe6, 0xd7, 0xff, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x61, 0x2f, 0x9f, 0xcd, 0xe6, 0x02, 0x00, 0x00,
}
