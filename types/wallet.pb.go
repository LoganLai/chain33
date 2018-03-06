// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wallet.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WalletTxDetail struct {
	Tx         *Transaction `protobuf:"bytes,1,opt,name=tx" json:"tx,omitempty"`
	Receipt    *ReceiptData `protobuf:"bytes,2,opt,name=receipt" json:"receipt,omitempty"`
	Height     int64        `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	Index      int64        `protobuf:"varint,4,opt,name=index" json:"index,omitempty"`
	Blocktime  int64        `protobuf:"varint,5,opt,name=blocktime" json:"blocktime,omitempty"`
	Amount     int64        `protobuf:"varint,6,opt,name=amount" json:"amount,omitempty"`
	Fromaddr   string       `protobuf:"bytes,7,opt,name=fromaddr" json:"fromaddr,omitempty"`
	Txhash     []byte       `protobuf:"bytes,8,opt,name=txhash,proto3" json:"txhash,omitempty"`
	ActionName string       `protobuf:"bytes,9,opt,name=actionName" json:"actionName,omitempty"`
}

func (m *WalletTxDetail) Reset()                    { *m = WalletTxDetail{} }
func (m *WalletTxDetail) String() string            { return proto.CompactTextString(m) }
func (*WalletTxDetail) ProtoMessage()               {}
func (*WalletTxDetail) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *WalletTxDetail) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *WalletTxDetail) GetReceipt() *ReceiptData {
	if m != nil {
		return m.Receipt
	}
	return nil
}

func (m *WalletTxDetail) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *WalletTxDetail) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *WalletTxDetail) GetBlocktime() int64 {
	if m != nil {
		return m.Blocktime
	}
	return 0
}

func (m *WalletTxDetail) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *WalletTxDetail) GetFromaddr() string {
	if m != nil {
		return m.Fromaddr
	}
	return ""
}

func (m *WalletTxDetail) GetTxhash() []byte {
	if m != nil {
		return m.Txhash
	}
	return nil
}

func (m *WalletTxDetail) GetActionName() string {
	if m != nil {
		return m.ActionName
	}
	return ""
}

type WalletTxDetails struct {
	TxDetails []*WalletTxDetail `protobuf:"bytes,1,rep,name=txDetails" json:"txDetails,omitempty"`
}

func (m *WalletTxDetails) Reset()                    { *m = WalletTxDetails{} }
func (m *WalletTxDetails) String() string            { return proto.CompactTextString(m) }
func (*WalletTxDetails) ProtoMessage()               {}
func (*WalletTxDetails) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

func (m *WalletTxDetails) GetTxDetails() []*WalletTxDetail {
	if m != nil {
		return m.TxDetails
	}
	return nil
}

type WalletAccountStore struct {
	Privkey   string `protobuf:"bytes,1,opt,name=privkey" json:"privkey,omitempty"`
	Label     string `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
	Addr      string `protobuf:"bytes,3,opt,name=addr" json:"addr,omitempty"`
	TimeStamp string `protobuf:"bytes,4,opt,name=timeStamp" json:"timeStamp,omitempty"`
}

func (m *WalletAccountStore) Reset()                    { *m = WalletAccountStore{} }
func (m *WalletAccountStore) String() string            { return proto.CompactTextString(m) }
func (*WalletAccountStore) ProtoMessage()               {}
func (*WalletAccountStore) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

func (m *WalletAccountStore) GetPrivkey() string {
	if m != nil {
		return m.Privkey
	}
	return ""
}

func (m *WalletAccountStore) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *WalletAccountStore) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *WalletAccountStore) GetTimeStamp() string {
	if m != nil {
		return m.TimeStamp
	}
	return ""
}

type WalletPwHash struct {
	PwHash  []byte `protobuf:"bytes,1,opt,name=pwHash,proto3" json:"pwHash,omitempty"`
	Randstr string `protobuf:"bytes,2,opt,name=randstr" json:"randstr,omitempty"`
}

func (m *WalletPwHash) Reset()                    { *m = WalletPwHash{} }
func (m *WalletPwHash) String() string            { return proto.CompactTextString(m) }
func (*WalletPwHash) ProtoMessage()               {}
func (*WalletPwHash) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{3} }

func (m *WalletPwHash) GetPwHash() []byte {
	if m != nil {
		return m.PwHash
	}
	return nil
}

func (m *WalletPwHash) GetRandstr() string {
	if m != nil {
		return m.Randstr
	}
	return ""
}

type WalletStatus struct {
	IsLock       bool `protobuf:"varint,1,opt,name=isLock" json:"isLock,omitempty"`
	IsAutoMining bool `protobuf:"varint,2,opt,name=isAutoMining" json:"isAutoMining,omitempty"`
	HasSeed      bool `protobuf:"varint,3,opt,name=hasSeed" json:"hasSeed,omitempty"`
}

func (m *WalletStatus) Reset()                    { *m = WalletStatus{} }
func (m *WalletStatus) String() string            { return proto.CompactTextString(m) }
func (*WalletStatus) ProtoMessage()               {}
func (*WalletStatus) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{4} }

func (m *WalletStatus) GetIsLock() bool {
	if m != nil {
		return m.IsLock
	}
	return false
}

func (m *WalletStatus) GetIsAutoMining() bool {
	if m != nil {
		return m.IsAutoMining
	}
	return false
}

func (m *WalletStatus) GetHasSeed() bool {
	if m != nil {
		return m.HasSeed
	}
	return false
}

type WalletAccounts struct {
	Wallets []*WalletAccount `protobuf:"bytes,1,rep,name=wallets" json:"wallets,omitempty"`
}

func (m *WalletAccounts) Reset()                    { *m = WalletAccounts{} }
func (m *WalletAccounts) String() string            { return proto.CompactTextString(m) }
func (*WalletAccounts) ProtoMessage()               {}
func (*WalletAccounts) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{5} }

func (m *WalletAccounts) GetWallets() []*WalletAccount {
	if m != nil {
		return m.Wallets
	}
	return nil
}

type WalletAccount struct {
	Acc   *Account `protobuf:"bytes,1,opt,name=acc" json:"acc,omitempty"`
	Label string   `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
}

func (m *WalletAccount) Reset()                    { *m = WalletAccount{} }
func (m *WalletAccount) String() string            { return proto.CompactTextString(m) }
func (*WalletAccount) ProtoMessage()               {}
func (*WalletAccount) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{6} }

func (m *WalletAccount) GetAcc() *Account {
	if m != nil {
		return m.Acc
	}
	return nil
}

func (m *WalletAccount) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type WalletUnLock struct {
	Passwd  string `protobuf:"bytes,1,opt,name=passwd" json:"passwd,omitempty"`
	Timeout int64  `protobuf:"varint,2,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *WalletUnLock) Reset()                    { *m = WalletUnLock{} }
func (m *WalletUnLock) String() string            { return proto.CompactTextString(m) }
func (*WalletUnLock) ProtoMessage()               {}
func (*WalletUnLock) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{7} }

func (m *WalletUnLock) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

func (m *WalletUnLock) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

type GenSeedLang struct {
	Lang int32 `protobuf:"varint,1,opt,name=lang" json:"lang,omitempty"`
}

func (m *GenSeedLang) Reset()                    { *m = GenSeedLang{} }
func (m *GenSeedLang) String() string            { return proto.CompactTextString(m) }
func (*GenSeedLang) ProtoMessage()               {}
func (*GenSeedLang) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{8} }

func (m *GenSeedLang) GetLang() int32 {
	if m != nil {
		return m.Lang
	}
	return 0
}

type GetSeedByPw struct {
	Passwd string `protobuf:"bytes,1,opt,name=passwd" json:"passwd,omitempty"`
}

func (m *GetSeedByPw) Reset()                    { *m = GetSeedByPw{} }
func (m *GetSeedByPw) String() string            { return proto.CompactTextString(m) }
func (*GetSeedByPw) ProtoMessage()               {}
func (*GetSeedByPw) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{9} }

func (m *GetSeedByPw) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

type SaveSeedByPw struct {
	Seed   string `protobuf:"bytes,1,opt,name=seed" json:"seed,omitempty"`
	Passwd string `protobuf:"bytes,2,opt,name=passwd" json:"passwd,omitempty"`
}

func (m *SaveSeedByPw) Reset()                    { *m = SaveSeedByPw{} }
func (m *SaveSeedByPw) String() string            { return proto.CompactTextString(m) }
func (*SaveSeedByPw) ProtoMessage()               {}
func (*SaveSeedByPw) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{10} }

func (m *SaveSeedByPw) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

func (m *SaveSeedByPw) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

type ReplySeed struct {
	Seed string `protobuf:"bytes,1,opt,name=seed" json:"seed,omitempty"`
}

func (m *ReplySeed) Reset()                    { *m = ReplySeed{} }
func (m *ReplySeed) String() string            { return proto.CompactTextString(m) }
func (*ReplySeed) ProtoMessage()               {}
func (*ReplySeed) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{11} }

func (m *ReplySeed) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

type ReqWalletSetPasswd struct {
	Oldpass string `protobuf:"bytes,1,opt,name=oldpass" json:"oldpass,omitempty"`
	Newpass string `protobuf:"bytes,2,opt,name=newpass" json:"newpass,omitempty"`
}

func (m *ReqWalletSetPasswd) Reset()                    { *m = ReqWalletSetPasswd{} }
func (m *ReqWalletSetPasswd) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletSetPasswd) ProtoMessage()               {}
func (*ReqWalletSetPasswd) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{12} }

func (m *ReqWalletSetPasswd) GetOldpass() string {
	if m != nil {
		return m.Oldpass
	}
	return ""
}

func (m *ReqWalletSetPasswd) GetNewpass() string {
	if m != nil {
		return m.Newpass
	}
	return ""
}

type ReqNewAccount struct {
	Label string `protobuf:"bytes,1,opt,name=label" json:"label,omitempty"`
}

func (m *ReqNewAccount) Reset()                    { *m = ReqNewAccount{} }
func (m *ReqNewAccount) String() string            { return proto.CompactTextString(m) }
func (*ReqNewAccount) ProtoMessage()               {}
func (*ReqNewAccount) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{13} }

func (m *ReqNewAccount) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type MinerFlag struct {
	Flag int32 `protobuf:"varint,1,opt,name=flag" json:"flag,omitempty"`
}

func (m *MinerFlag) Reset()                    { *m = MinerFlag{} }
func (m *MinerFlag) String() string            { return proto.CompactTextString(m) }
func (*MinerFlag) ProtoMessage()               {}
func (*MinerFlag) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{14} }

func (m *MinerFlag) GetFlag() int32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

type ReqWalletTransactionList struct {
	FromTx    []byte `protobuf:"bytes,1,opt,name=fromTx,proto3" json:"fromTx,omitempty"`
	Count     int32  `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	Direction int32  `protobuf:"varint,3,opt,name=direction" json:"direction,omitempty"`
}

func (m *ReqWalletTransactionList) Reset()                    { *m = ReqWalletTransactionList{} }
func (m *ReqWalletTransactionList) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletTransactionList) ProtoMessage()               {}
func (*ReqWalletTransactionList) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{15} }

func (m *ReqWalletTransactionList) GetFromTx() []byte {
	if m != nil {
		return m.FromTx
	}
	return nil
}

func (m *ReqWalletTransactionList) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReqWalletTransactionList) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

type ReqWalletImportPrivKey struct {
	Privkey string `protobuf:"bytes,1,opt,name=privkey" json:"privkey,omitempty"`
	Label   string `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
}

func (m *ReqWalletImportPrivKey) Reset()                    { *m = ReqWalletImportPrivKey{} }
func (m *ReqWalletImportPrivKey) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletImportPrivKey) ProtoMessage()               {}
func (*ReqWalletImportPrivKey) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{16} }

func (m *ReqWalletImportPrivKey) GetPrivkey() string {
	if m != nil {
		return m.Privkey
	}
	return ""
}

func (m *ReqWalletImportPrivKey) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type ReqWalletSendToAddress struct {
	From   string `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	To     string `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
	Amount int64  `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
	Note   string `protobuf:"bytes,4,opt,name=note" json:"note,omitempty"`
}

func (m *ReqWalletSendToAddress) Reset()                    { *m = ReqWalletSendToAddress{} }
func (m *ReqWalletSendToAddress) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletSendToAddress) ProtoMessage()               {}
func (*ReqWalletSendToAddress) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{17} }

func (m *ReqWalletSendToAddress) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *ReqWalletSendToAddress) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *ReqWalletSendToAddress) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ReqWalletSendToAddress) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

type ReqWalletSetFee struct {
	Amount int64 `protobuf:"varint,1,opt,name=amount" json:"amount,omitempty"`
}

func (m *ReqWalletSetFee) Reset()                    { *m = ReqWalletSetFee{} }
func (m *ReqWalletSetFee) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletSetFee) ProtoMessage()               {}
func (*ReqWalletSetFee) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{18} }

func (m *ReqWalletSetFee) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type ReqWalletSetLabel struct {
	Addr  string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
}

func (m *ReqWalletSetLabel) Reset()                    { *m = ReqWalletSetLabel{} }
func (m *ReqWalletSetLabel) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletSetLabel) ProtoMessage()               {}
func (*ReqWalletSetLabel) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{19} }

func (m *ReqWalletSetLabel) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *ReqWalletSetLabel) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type ReqWalletMergeBalance struct {
	To string `protobuf:"bytes,1,opt,name=to" json:"to,omitempty"`
}

func (m *ReqWalletMergeBalance) Reset()                    { *m = ReqWalletMergeBalance{} }
func (m *ReqWalletMergeBalance) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletMergeBalance) ProtoMessage()               {}
func (*ReqWalletMergeBalance) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{20} }

func (m *ReqWalletMergeBalance) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

type ReqStr struct {
	Reqstr string `protobuf:"bytes,1,opt,name=reqstr" json:"reqstr,omitempty"`
}

func (m *ReqStr) Reset()                    { *m = ReqStr{} }
func (m *ReqStr) String() string            { return proto.CompactTextString(m) }
func (*ReqStr) ProtoMessage()               {}
func (*ReqStr) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{21} }

func (m *ReqStr) GetReqstr() string {
	if m != nil {
		return m.Reqstr
	}
	return ""
}

type ReplyStr struct {
	Replystr string `protobuf:"bytes,1,opt,name=replystr" json:"replystr,omitempty"`
}

func (m *ReplyStr) Reset()                    { *m = ReplyStr{} }
func (m *ReplyStr) String() string            { return proto.CompactTextString(m) }
func (*ReplyStr) ProtoMessage()               {}
func (*ReplyStr) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{22} }

func (m *ReplyStr) GetReplystr() string {
	if m != nil {
		return m.Replystr
	}
	return ""
}

func init() {
	proto.RegisterType((*WalletTxDetail)(nil), "types.WalletTxDetail")
	proto.RegisterType((*WalletTxDetails)(nil), "types.WalletTxDetails")
	proto.RegisterType((*WalletAccountStore)(nil), "types.WalletAccountStore")
	proto.RegisterType((*WalletPwHash)(nil), "types.WalletPwHash")
	proto.RegisterType((*WalletStatus)(nil), "types.WalletStatus")
	proto.RegisterType((*WalletAccounts)(nil), "types.WalletAccounts")
	proto.RegisterType((*WalletAccount)(nil), "types.WalletAccount")
	proto.RegisterType((*WalletUnLock)(nil), "types.WalletUnLock")
	proto.RegisterType((*GenSeedLang)(nil), "types.GenSeedLang")
	proto.RegisterType((*GetSeedByPw)(nil), "types.GetSeedByPw")
	proto.RegisterType((*SaveSeedByPw)(nil), "types.SaveSeedByPw")
	proto.RegisterType((*ReplySeed)(nil), "types.ReplySeed")
	proto.RegisterType((*ReqWalletSetPasswd)(nil), "types.ReqWalletSetPasswd")
	proto.RegisterType((*ReqNewAccount)(nil), "types.ReqNewAccount")
	proto.RegisterType((*MinerFlag)(nil), "types.MinerFlag")
	proto.RegisterType((*ReqWalletTransactionList)(nil), "types.ReqWalletTransactionList")
	proto.RegisterType((*ReqWalletImportPrivKey)(nil), "types.ReqWalletImportPrivKey")
	proto.RegisterType((*ReqWalletSendToAddress)(nil), "types.ReqWalletSendToAddress")
	proto.RegisterType((*ReqWalletSetFee)(nil), "types.ReqWalletSetFee")
	proto.RegisterType((*ReqWalletSetLabel)(nil), "types.ReqWalletSetLabel")
	proto.RegisterType((*ReqWalletMergeBalance)(nil), "types.ReqWalletMergeBalance")
	proto.RegisterType((*ReqStr)(nil), "types.ReqStr")
	proto.RegisterType((*ReplyStr)(nil), "types.ReplyStr")
}

func init() { proto.RegisterFile("wallet.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 799 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdf, 0x6f, 0xdb, 0x36,
	0x10, 0x86, 0xec, 0xd8, 0xb1, 0x2e, 0x4e, 0xba, 0x12, 0x6d, 0x20, 0x04, 0xc3, 0xea, 0x11, 0xe8,
	0xe6, 0x01, 0x43, 0x1e, 0xda, 0xb7, 0x01, 0x03, 0x9a, 0xa2, 0x48, 0x33, 0x2c, 0x29, 0x02, 0x3a,
	0xc3, 0x9e, 0x19, 0xe9, 0x62, 0x0b, 0x91, 0x45, 0x87, 0x64, 0xfc, 0xe3, 0x5f, 0xd9, 0x5f, 0x3b,
	0xf0, 0x48, 0xca, 0x12, 0x90, 0x3d, 0xec, 0xed, 0xbe, 0xe3, 0x7d, 0x77, 0xbc, 0x4f, 0xc7, 0x13,
	0x8c, 0x37, 0xb2, 0xaa, 0xd0, 0x9e, 0xaf, 0xb4, 0xb2, 0x8a, 0x0d, 0xec, 0x6e, 0x85, 0xe6, 0xec,
	0xb5, 0xd5, 0xb2, 0x36, 0x32, 0xb7, 0xa5, 0xaa, 0xfd, 0xc9, 0xd9, 0x77, 0xf7, 0x95, 0xca, 0x1f,
	0xf3, 0x85, 0x2c, 0xa3, 0xe7, 0x58, 0xe6, 0xb9, 0x7a, 0xae, 0x03, 0x95, 0xff, 0xd3, 0x83, 0x93,
	0xbf, 0x29, 0xd7, 0xdd, 0xf6, 0x0b, 0x5a, 0x59, 0x56, 0x8c, 0x43, 0xcf, 0x6e, 0xb3, 0x64, 0x92,
	0x4c, 0x8f, 0x3e, 0xb0, 0x73, 0x4a, 0x7d, 0x7e, 0xb7, 0xcf, 0x2c, 0x7a, 0x76, 0xcb, 0x7e, 0x85,
	0x43, 0x8d, 0x39, 0x96, 0x2b, 0x9b, 0xf5, 0x3a, 0x81, 0xc2, 0x7b, 0xbf, 0x48, 0x2b, 0x45, 0x0c,
	0x61, 0xa7, 0x30, 0x5c, 0x60, 0x39, 0x5f, 0xd8, 0xac, 0x3f, 0x49, 0xa6, 0x7d, 0x11, 0x10, 0x7b,
	0x03, 0x83, 0xb2, 0x2e, 0x70, 0x9b, 0x1d, 0x90, 0xdb, 0x03, 0xf6, 0x3d, 0xa4, 0x74, 0x6b, 0x5b,
	0x2e, 0x31, 0x1b, 0xd0, 0xc9, 0xde, 0xe1, 0x72, 0xc9, 0xa5, 0x6b, 0x20, 0x1b, 0xfa, 0x5c, 0x1e,
	0xb1, 0x33, 0x18, 0x3d, 0x68, 0xb5, 0x94, 0x45, 0xa1, 0xb3, 0xc3, 0x49, 0x32, 0x4d, 0x45, 0x83,
	0x1d, 0xc7, 0x6e, 0x17, 0xd2, 0x2c, 0xb2, 0xd1, 0x24, 0x99, 0x8e, 0x45, 0x40, 0xec, 0x07, 0x00,
	0xdf, 0xd3, 0x37, 0xb9, 0xc4, 0x2c, 0x25, 0x56, 0xcb, 0xc3, 0x2f, 0xe1, 0x55, 0x57, 0x1b, 0xc3,
	0x3e, 0x42, 0x6a, 0x23, 0xc8, 0x92, 0x49, 0x7f, 0x7a, 0xf4, 0xe1, 0x6d, 0x68, 0xbd, 0x1b, 0x2a,
	0xf6, 0x71, 0x7c, 0x0d, 0xcc, 0x1f, 0x5e, 0x78, 0xed, 0x67, 0x56, 0x69, 0x64, 0x19, 0x1c, 0xae,
	0x74, 0xb9, 0x7e, 0xc4, 0x1d, 0x89, 0x9d, 0x8a, 0x08, 0x9d, 0x2e, 0x95, 0xbc, 0xc7, 0x8a, 0xb4,
	0x4d, 0x85, 0x07, 0x8c, 0xc1, 0x01, 0x75, 0xd7, 0x27, 0x27, 0xd9, 0x4e, 0x2b, 0xa7, 0xca, 0xcc,
	0xca, 0xe5, 0x8a, 0x54, 0x4c, 0xc5, 0xde, 0xc1, 0x3f, 0xc1, 0xd8, 0xd7, 0xbd, 0xdd, 0x5c, 0xb9,
	0x7e, 0x4f, 0x61, 0xb8, 0x22, 0x8b, 0x0a, 0x8e, 0x45, 0x40, 0xee, 0x26, 0x5a, 0xd6, 0x85, 0xb1,
	0x3a, 0x54, 0x8c, 0x90, 0x17, 0x31, 0xc3, 0xcc, 0x4a, 0xfb, 0x6c, 0x5c, 0x86, 0xd2, 0x5c, 0xab,
	0xfc, 0x91, 0x32, 0x8c, 0x44, 0x40, 0x8c, 0xc3, 0xb8, 0x34, 0x17, 0xcf, 0x56, 0xdd, 0x94, 0x75,
	0x59, 0xcf, 0x29, 0xcd, 0x48, 0x74, 0x7c, 0xae, 0xca, 0x42, 0x9a, 0x19, 0x62, 0x41, 0x2d, 0x8c,
	0x44, 0x84, 0xfc, 0x53, 0x9c, 0xc1, 0xa0, 0x8f, 0x61, 0xe7, 0x70, 0xe8, 0x27, 0x3c, 0x8a, 0xfc,
	0xa6, 0x23, 0x72, 0x88, 0x13, 0x31, 0x88, 0x7f, 0x85, 0xe3, 0xce, 0x09, 0x9b, 0x40, 0x5f, 0xe6,
	0x79, 0x98, 0xe2, 0x93, 0x40, 0x8e, 0x34, 0x77, 0xf4, 0xb2, 0xc8, 0x7b, 0xc9, 0xfe, 0xaa, 0xa9,
	0x31, 0x27, 0x99, 0x34, 0x66, 0x53, 0x84, 0x6f, 0x14, 0x90, 0x6b, 0xc6, 0xe9, 0xac, 0x9e, 0xfd,
	0x03, 0xe8, 0x8b, 0x08, 0xf9, 0x8f, 0x70, 0xf4, 0x15, 0x6b, 0xd7, 0xd7, 0xb5, 0xac, 0xe7, 0xee,
	0xab, 0x55, 0xb2, 0x9e, 0x13, 0x7d, 0x20, 0xc8, 0xe6, 0xef, 0x5d, 0x88, 0x75, 0x21, 0x9f, 0x77,
	0xb7, 0x9b, 0xff, 0xaa, 0xc1, 0x7f, 0x83, 0xf1, 0x4c, 0xae, 0xb1, 0x89, 0x63, 0x70, 0x60, 0x9c,
	0x7a, 0x3e, 0x8a, 0xec, 0x16, 0xb7, 0xd7, 0xe1, 0xbe, 0x83, 0x54, 0xe0, 0xaa, 0xda, 0x39, 0xf2,
	0x4b, 0x44, 0x7e, 0x05, 0x4c, 0xe0, 0x53, 0xf8, 0xb8, 0x68, 0x6f, 0x9b, 0xb6, 0x54, 0x55, 0xb8,
	0x1c, 0x71, 0x26, 0x03, 0x74, 0x27, 0x35, 0x6e, 0xe8, 0x24, 0xcc, 0x48, 0x80, 0xfc, 0x3d, 0x1c,
	0x0b, 0x7c, 0xfa, 0x86, 0x9b, 0xa8, 0x7d, 0xa3, 0x6c, 0xd2, 0x56, 0xf6, 0x1d, 0xa4, 0x37, 0x65,
	0x8d, 0xfa, 0xb2, 0x92, 0xa4, 0xca, 0x43, 0x25, 0x1b, 0x55, 0x9c, 0xcd, 0x1f, 0x20, 0x6b, 0x6e,
	0xd4, 0xda, 0x37, 0xd7, 0xa5, 0xa1, 0x0d, 0xe2, 0x5e, 0xf3, 0xdd, 0x36, 0x4e, 0xae, 0x47, 0xae,
	0x14, 0xd5, 0xa4, 0x3b, 0x0d, 0x84, 0x07, 0xee, 0x55, 0x14, 0xa5, 0x46, 0xa2, 0xd3, 0xac, 0x0d,
	0xc4, 0xde, 0xc1, 0xaf, 0xe0, 0xb4, 0xa9, 0xf3, 0xc7, 0x72, 0xa5, 0xb4, 0xbd, 0xd5, 0xe5, 0xfa,
	0x4f, 0xdc, 0xfd, 0xdf, 0x17, 0xc9, 0x17, 0xad, 0x4c, 0x33, 0xac, 0x8b, 0x3b, 0x75, 0x51, 0x14,
	0x1a, 0x8d, 0xa1, 0xfe, 0xb4, 0x5a, 0x46, 0xc5, 0x9d, 0xcd, 0x4e, 0xa0, 0x67, 0x55, 0x48, 0xd0,
	0xb3, 0xaa, 0xb5, 0xc9, 0xfa, 0x9d, 0x4d, 0xc6, 0xe0, 0xa0, 0x56, 0x16, 0xc3, 0x73, 0x26, 0x9b,
	0xff, 0x02, 0xaf, 0xda, 0x5f, 0xeb, 0x12, 0xdb, 0x8b, 0x30, 0x69, 0xd3, 0xf9, 0xef, 0xf0, 0xba,
	0x1d, 0x7a, 0xdd, 0xd9, 0x1d, 0x49, 0x6b, 0x77, 0xbc, 0xdc, 0xd3, 0xcf, 0xf0, 0xb6, 0xa1, 0xdf,
	0xa0, 0x9e, 0xe3, 0x67, 0x59, 0xc9, 0x3a, 0xc7, 0x70, 0xfd, 0x24, 0x5e, 0x9f, 0x4f, 0x60, 0x28,
	0xf0, 0x69, 0x66, 0x69, 0xbd, 0x6a, 0x7c, 0x72, 0xdb, 0x23, 0xcc, 0xaf, 0x47, 0xfc, 0x27, 0x18,
	0xf9, 0x19, 0xb4, 0xda, 0xad, 0x67, 0xed, 0xec, 0x7d, 0x54, 0x83, 0xef, 0x87, 0xf4, 0x2b, 0xfa,
	0xf8, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x12, 0x45, 0x5d, 0xd5, 0x06, 0x00, 0x00,
}
