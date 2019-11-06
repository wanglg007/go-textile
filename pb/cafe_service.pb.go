// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cafe_service.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CafeChallenge struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeChallenge) Reset()         { *m = CafeChallenge{} }
func (m *CafeChallenge) String() string { return proto.CompactTextString(m) }
func (*CafeChallenge) ProtoMessage()    {}
func (*CafeChallenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_af259e22dc6e576e, []int{0}
}

func (m *CafeChallenge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeChallenge.Unmarshal(m, b)
}
func (m *CafeChallenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeChallenge.Marshal(b, m, deterministic)
}
func (m *CafeChallenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeChallenge.Merge(m, src)
}
func (m *CafeChallenge) XXX_Size() int {
	return xxx_messageInfo_CafeChallenge.Size(m)
}
func (m *CafeChallenge) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeChallenge.DiscardUnknown(m)
}

var xxx_messageInfo_CafeChallenge proto.InternalMessageInfo

func (m *CafeChallenge) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type CafeNonce struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeNonce) Reset()         { *m = CafeNonce{} }
func (m *CafeNonce) String() string { return proto.CompactTextString(m) }
func (*CafeNonce) ProtoMessage()    {}
func (*CafeNonce) Descriptor() ([]byte, []int) {
	return fileDescriptor_af259e22dc6e576e, []int{1}
}

func (m *CafeNonce) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeNonce.Unmarshal(m, b)
}
func (m *CafeNonce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeNonce.Marshal(b, m, deterministic)
}
func (m *CafeNonce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeNonce.Merge(m, src)
}
func (m *CafeNonce) XXX_Size() int {
	return xxx_messageInfo_CafeNonce.Size(m)
}
func (m *CafeNonce) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeNonce.DiscardUnknown(m)
}

var xxx_messageInfo_CafeNonce proto.InternalMessageInfo

func (m *CafeNonce) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type CafeRegistration struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Nonce                string   `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Sig                  []byte   `protobuf:"bytes,4,opt,name=sig,proto3" json:"sig,omitempty"`
	Token                string   `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeRegistration) Reset()         { *m = CafeRegistration{} }
func (m *CafeRegistration) String() string { return proto.CompactTextString(m) }
func (*CafeRegistration) ProtoMessage()    {}
func (*CafeRegistration) Descriptor() ([]byte, []int) {
	return fileDescriptor_af259e22dc6e576e, []int{2}
}

func (m *CafeRegistration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeRegistration.Unmarshal(m, b)
}
func (m *CafeRegistration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeRegistration.Marshal(b, m, deterministic)
}
func (m *CafeRegistration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeRegistration.Merge(m, src)
}
func (m *CafeRegistration) XXX_Size() int {
	return xxx_messageInfo_CafeRegistration.Size(m)
}
func (m *CafeRegistration) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeRegistration.DiscardUnknown(m)
}

var xxx_messageInfo_CafeRegistration proto.InternalMessageInfo

func (m *CafeRegistration) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *CafeRegistration) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *CafeRegistration) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *CafeRegistration) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

func (m *CafeRegistration) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CafeDeregistration struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeDeregistration) Reset()         { *m = CafeDeregistration{} }
func (m *CafeDeregistration) String() string { return proto.CompactTextString(m) }
func (*CafeDeregistration) ProtoMessage()    {}
func (*CafeDeregistration) Descriptor() ([]byte, []int) {
	return fileDescriptor_af259e22dc6e576e, []int{3}
}

func (m *CafeDeregistration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeDeregistration.Unmarshal(m, b)
}
func (m *CafeDeregistration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeDeregistration.Marshal(b, m, deterministic)
}
func (m *CafeDeregistration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeDeregistration.Merge(m, src)
}
func (m *CafeDeregistration) XXX_Size() int {
	return xxx_messageInfo_CafeDeregistration.Size(m)
}
func (m *CafeDeregistration) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeDeregistration.DiscardUnknown(m)
}

var xxx_messageInfo_CafeDeregistration proto.InternalMessageInfo

func (m *CafeDeregistration) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CafeDeregistrationAck struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeDeregistrationAck) Reset()         { *m = CafeDeregistrationAck{} }
func (m *CafeDeregistrationAck) String() string { return proto.CompactTextString(m) }
func (*CafeDeregistrationAck) ProtoMessage()    {}
func (*CafeDeregistrationAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_af259e22dc6e576e, []int{4}
}

func (m *CafeDeregistrationAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeDeregistrationAck.Unmarshal(m, b)
}
func (m *CafeDeregistrationAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeDeregistrationAck.Marshal(b, m, deterministic)
}
func (m *CafeDeregistrationAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeDeregistrationAck.Merge(m, src)
}
func (m *CafeDeregistrationAck) XXX_Size() int {
	return xxx_messageInfo_CafeDeregistrationAck.Size(m)
}
func (m *CafeDeregistrationAck) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeDeregistrationAck.DiscardUnknown(m)
}

var xxx_messageInfo_CafeDeregistrationAck proto.InternalMessageInfo

func (m *CafeDeregistrationAck) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CafeRefreshSession struct {
	Access               string   `protobuf:"bytes,1,opt,name=access,proto3" json:"access,omitempty"`
	Refresh              string   `protobuf:"bytes,2,opt,name=refresh,proto3" json:"refresh,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeRefreshSession) Reset()         { *m = CafeRefreshSession{} }
func (m *CafeRefreshSession) String() string { return proto.CompactTextString(m) }
func (*CafeRefreshSession) ProtoMessage()    {}
func (*CafeRefreshSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_af259e22dc6e576e, []int{5}
}

func (m *CafeRefreshSession) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeRefreshSession.Unmarshal(m, b)
}
func (m *CafeRefreshSession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeRefreshSession.Marshal(b, m, deterministic)
}
func (m *CafeRefreshSession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeRefreshSession.Merge(m, src)
}
func (m *CafeRefreshSession) XXX_Size() int {
	return xxx_messageInfo_CafeRefreshSession.Size(m)
}
func (m *CafeRefreshSession) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeRefreshSession.DiscardUnknown(m)
}

var xxx_messageInfo_CafeRefreshSession proto.InternalMessageInfo

func (m *CafeRefreshSession) GetAccess() string {
	if m != nil {
		return m.Access
	}
	return ""
}

func (m *CafeRefreshSession) GetRefresh() string {
	if m != nil {
		return m.Refresh
	}
	return ""
}

type CafePublishPeer struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Peer                 *Peer    `protobuf:"bytes,2,opt,name=peer,proto3" json:"peer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafePublishPeer) Reset()         { *m = CafePublishPeer{} }
func (m *CafePublishPeer) String() string { return proto.CompactTextString(m) }
func (*CafePublishPeer) ProtoMessage()    {}
func (*CafePublishPeer) Descriptor() ([]byte, []int) {
	return fileDescriptor_af259e22dc6e576e, []int{6}
}

func (m *CafePublishPeer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafePublishPeer.Unmarshal(m, b)
}
func (m *CafePublishPeer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafePublishPeer.Marshal(b, m, deterministic)
}
func (m *CafePublishPeer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafePublishPeer.Merge(m, src)
}
func (m *CafePublishPeer) XXX_Size() int {
	return xxx_messageInfo_CafePublishPeer.Size(m)
}
func (m *CafePublishPeer) XXX_DiscardUnknown() {
	xxx_messageInfo_CafePublishPeer.DiscardUnknown(m)
}

var xxx_messageInfo_CafePublishPeer proto.InternalMessageInfo

func (m *CafePublishPeer) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CafePublishPeer) GetPeer() *Peer {
	if m != nil {
		return m.Peer
	}
	return nil
}

type CafePublishPeerAck struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafePublishPeerAck) Reset()         { *m = CafePublishPeerAck{} }
func (m *CafePublishPeerAck) String() string { return proto.CompactTextString(m) }
func (*CafePublishPeerAck) ProtoMessage()    {}
func (*CafePublishPeerAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_af259e22dc6e576e, []int{7}
}

func (m *CafePublishPeerAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafePublishPeerAck.Unmarshal(m, b)
}
func (m *CafePublishPeerAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafePublishPeerAck.Marshal(b, m, deterministic)
}
func (m *CafePublishPeerAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafePublishPeerAck.Merge(m, src)
}
func (m *CafePublishPeerAck) XXX_Size() int {
	return xxx_messageInfo_CafePublishPeerAck.Size(m)
}
func (m *CafePublishPeerAck) XXX_DiscardUnknown() {
	xxx_messageInfo_CafePublishPeerAck.DiscardUnknown(m)
}

var xxx_messageInfo_CafePublishPeerAck proto.InternalMessageInfo

func (m *CafePublishPeerAck) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CafeStore struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Cids                 []string `protobuf:"bytes,2,rep,name=cids,proto3" json:"cids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeStore) Reset()         { *m = CafeStore{} }
func (m *CafeStore) String() string { return proto.CompactTextString(m) }
func (*CafeStore) ProtoMessage()    {}
func (*CafeStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_af259e22dc6e576e, []int{8}
}

func (m *CafeStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeStore.Unmarshal(m, b)
}
func (m *CafeStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeStore.Marshal(b, m, deterministic)
}
func (m *CafeStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeStore.Merge(m, src)
}
func (m *CafeStore) XXX_Size() int {
	return xxx_messageInfo_CafeStore.Size(m)
}
func (m *CafeStore) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeStore.DiscardUnknown(m)
}

var xxx_messageInfo_CafeStore proto.InternalMessageInfo

func (m *CafeStore) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CafeStore) GetCids() []string {
	if m != nil {
		return m.Cids
	}
	return nil
}

type CafeStoreAck struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeStoreAck) Reset()         { *m = CafeStoreAck{} }
func (m *CafeStoreAck) String() string { return proto.CompactTextString(m) }
func (*CafeStoreAck) ProtoMessage()    {}
func (*CafeStoreAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_af259e22dc6e576e, []int{9}
}

func (m *CafeStoreAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeStoreAck.Unmarshal(m, b)
}
func (m *CafeStoreAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeStoreAck.Marshal(b, m, deterministic)
}
func (m *CafeStoreAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeStoreAck.Merge(m, src)
}
func (m *CafeStoreAck) XXX_Size() int {
	return xxx_messageInfo_CafeStoreAck.Size(m)
}
func (m *CafeStoreAck) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeStoreAck.DiscardUnknown(m)
}

var xxx_messageInfo_CafeStoreAck proto.InternalMessageInfo

func (m *CafeStoreAck) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CafeUnstore struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Cids                 []string `protobuf:"bytes,2,rep,name=cids,proto3" json:"cids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeUnstore) Reset()         { *m = CafeUnstore{} }
func (m *CafeUnstore) String() string { return proto.CompactTextString(m) }
func (*CafeUnstore) ProtoMessage()    {}
func (*CafeUnstore) Descriptor() ([]byte, []int) {
	return fileDescriptor_af259e22dc6e576e, []int{10}
}

func (m *CafeUnstore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeUnstore.Unmarshal(m, b)
}
func (m *CafeUnstore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeUnstore.Marshal(b, m, deterministic)
}
func (m *CafeUnstore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeUnstore.Merge(m, src)
}
func (m *CafeUnstore) XXX_Size() int {
	return xxx_messageInfo_CafeUnstore.Size(m)
}
func (m *CafeUnstore) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeUnstore.DiscardUnknown(m)
}

var xxx_messageInfo_CafeUnstore proto.InternalMessageInfo

func (m *CafeUnstore) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CafeUnstore) GetCids() []string {
	if m != nil {
		return m.Cids
	}
	return nil
}

type CafeUnstoreAck struct {
	Cids                 []string `protobuf:"bytes,1,rep,name=cids,proto3" json:"cids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeUnstoreAck) Reset()         { *m = CafeUnstoreAck{} }
func (m *CafeUnstoreAck) String() string { return proto.CompactTextString(m) }
func (*CafeUnstoreAck) ProtoMessage()    {}
func (*CafeUnstoreAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_af259e22dc6e576e, []int{11}
}

func (m *CafeUnstoreAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeUnstoreAck.Unmarshal(m, b)
}
func (m *CafeUnstoreAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeUnstoreAck.Marshal(b, m, deterministic)
}
func (m *CafeUnstoreAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeUnstoreAck.Merge(m, src)
}
func (m *CafeUnstoreAck) XXX_Size() int {
	return xxx_messageInfo_CafeUnstoreAck.Size(m)
}
func (m *CafeUnstoreAck) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeUnstoreAck.DiscardUnknown(m)
}

var xxx_messageInfo_CafeUnstoreAck proto.InternalMessageInfo

func (m *CafeUnstoreAck) GetCids() []string {
	if m != nil {
		return m.Cids
	}
	return nil
}

type CafeObjectList struct {
	Cids                 []string `protobuf:"bytes,1,rep,name=cids,proto3" json:"cids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeObjectList) Reset()         { *m = CafeObjectList{} }
func (m *CafeObjectList) String() string { return proto.CompactTextString(m) }
func (*CafeObjectList) ProtoMessage()    {}
func (*CafeObjectList) Descriptor() ([]byte, []int) {
	return fileDescriptor_af259e22dc6e576e, []int{12}
}

func (m *CafeObjectList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeObjectList.Unmarshal(m, b)
}
func (m *CafeObjectList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeObjectList.Marshal(b, m, deterministic)
}
func (m *CafeObjectList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeObjectList.Merge(m, src)
}
func (m *CafeObjectList) XXX_Size() int {
	return xxx_messageInfo_CafeObjectList.Size(m)
}
func (m *CafeObjectList) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeObjectList.DiscardUnknown(m)
}

var xxx_messageInfo_CafeObjectList proto.InternalMessageInfo

func (m *CafeObjectList) GetCids() []string {
	if m != nil {
		return m.Cids
	}
	return nil
}

type CafeObject struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Cid                  string   `protobuf:"bytes,2,opt,name=cid,proto3" json:"cid,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Node                 []byte   `protobuf:"bytes,4,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeObject) Reset()         { *m = CafeObject{} }
func (m *CafeObject) String() string { return proto.CompactTextString(m) }
func (*CafeObject) ProtoMessage()    {}
func (*CafeObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_af259e22dc6e576e, []int{13}
}

func (m *CafeObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeObject.Unmarshal(m, b)
}
func (m *CafeObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeObject.Marshal(b, m, deterministic)
}
func (m *CafeObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeObject.Merge(m, src)
}
func (m *CafeObject) XXX_Size() int {
	return xxx_messageInfo_CafeObject.Size(m)
}
func (m *CafeObject) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeObject.DiscardUnknown(m)
}

var xxx_messageInfo_CafeObject proto.InternalMessageInfo

func (m *CafeObject) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CafeObject) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func (m *CafeObject) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *CafeObject) GetNode() []byte {
	if m != nil {
		return m.Node
	}
	return nil
}

type CafeStoreThread struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Ciphertext           []byte   `protobuf:"bytes,3,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeStoreThread) Reset()         { *m = CafeStoreThread{} }
func (m *CafeStoreThread) String() string { return proto.CompactTextString(m) }
func (*CafeStoreThread) ProtoMessage()    {}
func (*CafeStoreThread) Descriptor() ([]byte, []int) {
	return fileDescriptor_af259e22dc6e576e, []int{14}
}

func (m *CafeStoreThread) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeStoreThread.Unmarshal(m, b)
}
func (m *CafeStoreThread) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeStoreThread.Marshal(b, m, deterministic)
}
func (m *CafeStoreThread) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeStoreThread.Merge(m, src)
}
func (m *CafeStoreThread) XXX_Size() int {
	return xxx_messageInfo_CafeStoreThread.Size(m)
}
func (m *CafeStoreThread) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeStoreThread.DiscardUnknown(m)
}

var xxx_messageInfo_CafeStoreThread proto.InternalMessageInfo

func (m *CafeStoreThread) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CafeStoreThread) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CafeStoreThread) GetCiphertext() []byte {
	if m != nil {
		return m.Ciphertext
	}
	return nil
}

type CafeStoreThreadAck struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeStoreThreadAck) Reset()         { *m = CafeStoreThreadAck{} }
func (m *CafeStoreThreadAck) String() string { return proto.CompactTextString(m) }
func (*CafeStoreThreadAck) ProtoMessage()    {}
func (*CafeStoreThreadAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_af259e22dc6e576e, []int{15}
}

func (m *CafeStoreThreadAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeStoreThreadAck.Unmarshal(m, b)
}
func (m *CafeStoreThreadAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeStoreThreadAck.Marshal(b, m, deterministic)
}
func (m *CafeStoreThreadAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeStoreThreadAck.Merge(m, src)
}
func (m *CafeStoreThreadAck) XXX_Size() int {
	return xxx_messageInfo_CafeStoreThreadAck.Size(m)
}
func (m *CafeStoreThreadAck) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeStoreThreadAck.DiscardUnknown(m)
}

var xxx_messageInfo_CafeStoreThreadAck proto.InternalMessageInfo

func (m *CafeStoreThreadAck) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CafeUnstoreThread struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeUnstoreThread) Reset()         { *m = CafeUnstoreThread{} }
func (m *CafeUnstoreThread) String() string { return proto.CompactTextString(m) }
func (*CafeUnstoreThread) ProtoMessage()    {}
func (*CafeUnstoreThread) Descriptor() ([]byte, []int) {
	return fileDescriptor_af259e22dc6e576e, []int{16}
}

func (m *CafeUnstoreThread) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeUnstoreThread.Unmarshal(m, b)
}
func (m *CafeUnstoreThread) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeUnstoreThread.Marshal(b, m, deterministic)
}
func (m *CafeUnstoreThread) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeUnstoreThread.Merge(m, src)
}
func (m *CafeUnstoreThread) XXX_Size() int {
	return xxx_messageInfo_CafeUnstoreThread.Size(m)
}
func (m *CafeUnstoreThread) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeUnstoreThread.DiscardUnknown(m)
}

var xxx_messageInfo_CafeUnstoreThread proto.InternalMessageInfo

func (m *CafeUnstoreThread) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CafeUnstoreThread) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CafeUnstoreThreadAck struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeUnstoreThreadAck) Reset()         { *m = CafeUnstoreThreadAck{} }
func (m *CafeUnstoreThreadAck) String() string { return proto.CompactTextString(m) }
func (*CafeUnstoreThreadAck) ProtoMessage()    {}
func (*CafeUnstoreThreadAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_af259e22dc6e576e, []int{17}
}

func (m *CafeUnstoreThreadAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeUnstoreThreadAck.Unmarshal(m, b)
}
func (m *CafeUnstoreThreadAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeUnstoreThreadAck.Marshal(b, m, deterministic)
}
func (m *CafeUnstoreThreadAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeUnstoreThreadAck.Merge(m, src)
}
func (m *CafeUnstoreThreadAck) XXX_Size() int {
	return xxx_messageInfo_CafeUnstoreThreadAck.Size(m)
}
func (m *CafeUnstoreThreadAck) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeUnstoreThreadAck.DiscardUnknown(m)
}

var xxx_messageInfo_CafeUnstoreThreadAck proto.InternalMessageInfo

func (m *CafeUnstoreThreadAck) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CafeDeliverMessage struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Client               string   `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	Env                  []byte   `protobuf:"bytes,3,opt,name=env,proto3" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeDeliverMessage) Reset()         { *m = CafeDeliverMessage{} }
func (m *CafeDeliverMessage) String() string { return proto.CompactTextString(m) }
func (*CafeDeliverMessage) ProtoMessage()    {}
func (*CafeDeliverMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_af259e22dc6e576e, []int{18}
}

func (m *CafeDeliverMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeDeliverMessage.Unmarshal(m, b)
}
func (m *CafeDeliverMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeDeliverMessage.Marshal(b, m, deterministic)
}
func (m *CafeDeliverMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeDeliverMessage.Merge(m, src)
}
func (m *CafeDeliverMessage) XXX_Size() int {
	return xxx_messageInfo_CafeDeliverMessage.Size(m)
}
func (m *CafeDeliverMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeDeliverMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CafeDeliverMessage proto.InternalMessageInfo

func (m *CafeDeliverMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CafeDeliverMessage) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

func (m *CafeDeliverMessage) GetEnv() []byte {
	if m != nil {
		return m.Env
	}
	return nil
}

type CafeCheckMessages struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeCheckMessages) Reset()         { *m = CafeCheckMessages{} }
func (m *CafeCheckMessages) String() string { return proto.CompactTextString(m) }
func (*CafeCheckMessages) ProtoMessage()    {}
func (*CafeCheckMessages) Descriptor() ([]byte, []int) {
	return fileDescriptor_af259e22dc6e576e, []int{19}
}

func (m *CafeCheckMessages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeCheckMessages.Unmarshal(m, b)
}
func (m *CafeCheckMessages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeCheckMessages.Marshal(b, m, deterministic)
}
func (m *CafeCheckMessages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeCheckMessages.Merge(m, src)
}
func (m *CafeCheckMessages) XXX_Size() int {
	return xxx_messageInfo_CafeCheckMessages.Size(m)
}
func (m *CafeCheckMessages) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeCheckMessages.DiscardUnknown(m)
}

var xxx_messageInfo_CafeCheckMessages proto.InternalMessageInfo

func (m *CafeCheckMessages) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CafeMessages struct {
	Messages             []*CafeMessage `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CafeMessages) Reset()         { *m = CafeMessages{} }
func (m *CafeMessages) String() string { return proto.CompactTextString(m) }
func (*CafeMessages) ProtoMessage()    {}
func (*CafeMessages) Descriptor() ([]byte, []int) {
	return fileDescriptor_af259e22dc6e576e, []int{20}
}

func (m *CafeMessages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeMessages.Unmarshal(m, b)
}
func (m *CafeMessages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeMessages.Marshal(b, m, deterministic)
}
func (m *CafeMessages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeMessages.Merge(m, src)
}
func (m *CafeMessages) XXX_Size() int {
	return xxx_messageInfo_CafeMessages.Size(m)
}
func (m *CafeMessages) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeMessages.DiscardUnknown(m)
}

var xxx_messageInfo_CafeMessages proto.InternalMessageInfo

func (m *CafeMessages) GetMessages() []*CafeMessage {
	if m != nil {
		return m.Messages
	}
	return nil
}

type CafeDeleteMessages struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeDeleteMessages) Reset()         { *m = CafeDeleteMessages{} }
func (m *CafeDeleteMessages) String() string { return proto.CompactTextString(m) }
func (*CafeDeleteMessages) ProtoMessage()    {}
func (*CafeDeleteMessages) Descriptor() ([]byte, []int) {
	return fileDescriptor_af259e22dc6e576e, []int{21}
}

func (m *CafeDeleteMessages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeDeleteMessages.Unmarshal(m, b)
}
func (m *CafeDeleteMessages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeDeleteMessages.Marshal(b, m, deterministic)
}
func (m *CafeDeleteMessages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeDeleteMessages.Merge(m, src)
}
func (m *CafeDeleteMessages) XXX_Size() int {
	return xxx_messageInfo_CafeDeleteMessages.Size(m)
}
func (m *CafeDeleteMessages) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeDeleteMessages.DiscardUnknown(m)
}

var xxx_messageInfo_CafeDeleteMessages proto.InternalMessageInfo

func (m *CafeDeleteMessages) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CafeDeleteMessagesAck struct {
	More                 bool     `protobuf:"varint,1,opt,name=more,proto3" json:"more,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafeDeleteMessagesAck) Reset()         { *m = CafeDeleteMessagesAck{} }
func (m *CafeDeleteMessagesAck) String() string { return proto.CompactTextString(m) }
func (*CafeDeleteMessagesAck) ProtoMessage()    {}
func (*CafeDeleteMessagesAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_af259e22dc6e576e, []int{22}
}

func (m *CafeDeleteMessagesAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafeDeleteMessagesAck.Unmarshal(m, b)
}
func (m *CafeDeleteMessagesAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafeDeleteMessagesAck.Marshal(b, m, deterministic)
}
func (m *CafeDeleteMessagesAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafeDeleteMessagesAck.Merge(m, src)
}
func (m *CafeDeleteMessagesAck) XXX_Size() int {
	return xxx_messageInfo_CafeDeleteMessagesAck.Size(m)
}
func (m *CafeDeleteMessagesAck) XXX_DiscardUnknown() {
	xxx_messageInfo_CafeDeleteMessagesAck.DiscardUnknown(m)
}

var xxx_messageInfo_CafeDeleteMessagesAck proto.InternalMessageInfo

func (m *CafeDeleteMessagesAck) GetMore() bool {
	if m != nil {
		return m.More
	}
	return false
}

func init() {
	proto.RegisterType((*CafeChallenge)(nil), "CafeChallenge")
	proto.RegisterType((*CafeNonce)(nil), "CafeNonce")
	proto.RegisterType((*CafeRegistration)(nil), "CafeRegistration")
	proto.RegisterType((*CafeDeregistration)(nil), "CafeDeregistration")
	proto.RegisterType((*CafeDeregistrationAck)(nil), "CafeDeregistrationAck")
	proto.RegisterType((*CafeRefreshSession)(nil), "CafeRefreshSession")
	proto.RegisterType((*CafePublishPeer)(nil), "CafePublishPeer")
	proto.RegisterType((*CafePublishPeerAck)(nil), "CafePublishPeerAck")
	proto.RegisterType((*CafeStore)(nil), "CafeStore")
	proto.RegisterType((*CafeStoreAck)(nil), "CafeStoreAck")
	proto.RegisterType((*CafeUnstore)(nil), "CafeUnstore")
	proto.RegisterType((*CafeUnstoreAck)(nil), "CafeUnstoreAck")
	proto.RegisterType((*CafeObjectList)(nil), "CafeObjectList")
	proto.RegisterType((*CafeObject)(nil), "CafeObject")
	proto.RegisterType((*CafeStoreThread)(nil), "CafeStoreThread")
	proto.RegisterType((*CafeStoreThreadAck)(nil), "CafeStoreThreadAck")
	proto.RegisterType((*CafeUnstoreThread)(nil), "CafeUnstoreThread")
	proto.RegisterType((*CafeUnstoreThreadAck)(nil), "CafeUnstoreThreadAck")
	proto.RegisterType((*CafeDeliverMessage)(nil), "CafeDeliverMessage")
	proto.RegisterType((*CafeCheckMessages)(nil), "CafeCheckMessages")
	proto.RegisterType((*CafeMessages)(nil), "CafeMessages")
	proto.RegisterType((*CafeDeleteMessages)(nil), "CafeDeleteMessages")
	proto.RegisterType((*CafeDeleteMessagesAck)(nil), "CafeDeleteMessagesAck")
}

func init() { proto.RegisterFile("cafe_service.proto", fileDescriptor_af259e22dc6e576e) }

var fileDescriptor_af259e22dc6e576e = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0x95, 0x74, 0x1d, 0xeb, 0xb5, 0x1b, 0x25, 0x0c, 0x14, 0x78, 0x98, 0x8a, 0x35, 0x41,
	0x07, 0x52, 0x1f, 0x86, 0x10, 0xf0, 0xc8, 0x8a, 0x78, 0x82, 0x31, 0x65, 0x20, 0x24, 0x84, 0x84,
	0x5c, 0xe7, 0xda, 0x98, 0xa6, 0x49, 0x65, 0x7b, 0x15, 0x8f, 0x7c, 0x74, 0x74, 0xb6, 0xd3, 0x46,
	0x34, 0x15, 0xda, 0xdb, 0xdd, 0xe5, 0xe7, 0xff, 0x9d, 0xef, 0xce, 0x81, 0x48, 0xf0, 0x29, 0xfe,
	0xd4, 0xa8, 0x56, 0x52, 0xe0, 0x68, 0xa9, 0x4a, 0x53, 0x3e, 0xee, 0x2e, 0xca, 0x14, 0x73, 0xe7,
	0xb0, 0x33, 0x38, 0x1c, 0xf3, 0x29, 0x8e, 0x33, 0x9e, 0xe7, 0x58, 0xcc, 0x30, 0x8a, 0xe1, 0x0e,
	0x4f, 0x53, 0x85, 0x5a, 0xc7, 0xc1, 0x20, 0x18, 0x76, 0x92, 0xca, 0x65, 0x4f, 0xa0, 0x43, 0xe8,
	0x65, 0x59, 0x08, 0x8c, 0x8e, 0xa1, 0xbd, 0xe2, 0xf9, 0x0d, 0x7a, 0xc8, 0x39, 0xec, 0x4f, 0x00,
	0x7d, 0x62, 0x12, 0x9c, 0x49, 0x6d, 0x14, 0x37, 0xb2, 0x2c, 0x76, 0x2b, 0x6e, 0x44, 0xc2, 0x9a,
	0x08, 0x45, 0x0b, 0xca, 0x11, 0xb7, 0x5c, 0xd4, 0x3a, 0x51, 0x1f, 0x5a, 0x5a, 0xce, 0xe2, 0xbd,
	0x41, 0x30, 0xec, 0x25, 0x64, 0x12, 0x67, 0xca, 0x39, 0x16, 0x71, 0xdb, 0x71, 0xd6, 0x61, 0xcf,
	0x21, 0xa2, 0x0a, 0xde, 0xa3, 0xaa, 0xd7, 0xb0, 0x66, 0x83, 0x3a, 0xfb, 0x0c, 0x1e, 0x6c, 0xb3,
	0xef, 0xc4, 0x3c, 0x3a, 0x82, 0x50, 0xa6, 0x9e, 0x0d, 0x65, 0xca, 0x3e, 0x38, 0xd1, 0x04, 0xa7,
	0x0a, 0x75, 0x76, 0x8d, 0x5a, 0x93, 0xe8, 0x43, 0xd8, 0xe7, 0x42, 0x6c, 0xee, 0xe5, 0x3d, 0xba,
	0xb0, 0x72, 0xa4, 0xbf, 0x58, 0xe5, 0xb2, 0x0b, 0xb8, 0x4b, 0x3a, 0x57, 0x37, 0x93, 0x5c, 0xea,
	0xec, 0x0a, 0x51, 0x35, 0x57, 0x16, 0x3d, 0x82, 0xbd, 0x25, 0xa2, 0xb2, 0xe7, 0xbb, 0xe7, 0xed,
	0x11, 0xa1, 0x89, 0x0d, 0xb1, 0x53, 0x57, 0x4b, 0x4d, 0xa3, 0xa9, 0xe2, 0x57, 0x6e, 0x58, 0xd7,
	0xa6, 0x54, 0xb8, 0x23, 0x47, 0x04, 0x7b, 0x42, 0xa6, 0x3a, 0x0e, 0x07, 0xad, 0x61, 0x27, 0xb1,
	0x36, 0x3b, 0x81, 0xde, 0xfa, 0x58, 0x93, 0xec, 0x6b, 0xe8, 0xd2, 0xf7, 0xaf, 0x85, 0xbe, 0xa5,
	0xf0, 0x29, 0x1c, 0xd5, 0x0e, 0x92, 0x74, 0x45, 0x05, 0xdb, 0xd4, 0xe7, 0xc9, 0x2f, 0x14, 0xe6,
	0xa3, 0xd4, 0xa6, 0x91, 0xfa, 0x01, 0xb0, 0xa1, 0x76, 0xd4, 0xd0, 0x87, 0x96, 0x90, 0xa9, 0xef,
	0x3f, 0x99, 0xa4, 0x94, 0x72, 0xc3, 0xed, 0x56, 0xf5, 0x12, 0x6b, 0x53, 0xac, 0x28, 0x53, 0xf4,
	0x5b, 0x65, 0x6d, 0xf6, 0xcd, 0xcd, 0xc8, 0xb6, 0xe0, 0x4b, 0xa6, 0x90, 0xa7, 0x3b, 0x52, 0xb8,
	0xde, 0x84, 0x55, 0x6f, 0xa2, 0x13, 0x00, 0x21, 0x97, 0x19, 0x2a, 0x83, 0xbf, 0x8d, 0x4f, 0x53,
	0x8b, 0x54, 0x83, 0xab, 0x09, 0x37, 0x75, 0xf8, 0x2d, 0xdc, 0xab, 0x35, 0xea, 0x36, 0x05, 0xb0,
	0xa7, 0x70, 0xbc, 0x75, 0xb4, 0x29, 0xc5, 0x65, 0xf5, 0x44, 0x72, 0xb9, 0x42, 0xf5, 0x09, 0xb5,
	0xe6, 0x33, 0xfc, 0x97, 0xa2, 0xed, 0x16, 0xb9, 0xc4, 0xc2, 0xf8, 0x0c, 0xde, 0xa3, 0xce, 0x62,
	0xb1, 0xf2, 0xf7, 0x23, 0x93, 0x9d, 0xb9, 0x92, 0xc7, 0x19, 0x8a, 0xb9, 0x57, 0xd3, 0x3b, 0x5e,
	0xdc, 0x1b, 0xb7, 0x5f, 0x6b, 0x6a, 0x08, 0x07, 0x0b, 0x6f, 0xdb, 0x11, 0x77, 0xcf, 0x7b, 0xa3,
	0x1a, 0x90, 0xac, 0xbf, 0x6e, 0xde, 0x75, 0x8e, 0x06, 0xff, 0x93, 0xe5, 0x45, 0xf5, 0xae, 0xeb,
	0xac, 0xdf, 0xb9, 0x45, 0xa9, 0xdc, 0x4f, 0xeb, 0x20, 0xb1, 0xf6, 0xc5, 0x7d, 0x38, 0x94, 0xe5,
	0x88, 0x26, 0x24, 0x73, 0x1c, 0x2d, 0x27, 0xdf, 0xc3, 0xe5, 0x64, 0xb2, 0x6f, 0xff, 0x8e, 0x2f,
	0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xc5, 0x28, 0x58, 0x40, 0x05, 0x00, 0x00,
}