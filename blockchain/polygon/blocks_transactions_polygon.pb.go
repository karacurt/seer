// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.6.1
// source: blocks_transactions_polygon.proto

package polygon

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents a single transaction within a block
type SingleTransactionPolygon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash                 string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	BlockNumber          uint64 `protobuf:"varint,2,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	FromAddress          string `protobuf:"bytes,3,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	ToAddress            string `protobuf:"bytes,4,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	Gas                  string `protobuf:"bytes,5,opt,name=gas,proto3" json:"gas,omitempty"` // using string to handle big numeric values
	GasPrice             string `protobuf:"bytes,6,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	MaxFeePerGas         string `protobuf:"bytes,7,opt,name=max_fee_per_gas,json=maxFeePerGas,proto3" json:"max_fee_per_gas,omitempty"`
	MaxPriorityFeePerGas string `protobuf:"bytes,8,opt,name=max_priority_fee_per_gas,json=maxPriorityFeePerGas,proto3" json:"max_priority_fee_per_gas,omitempty"`
	Input                string `protobuf:"bytes,9,opt,name=input,proto3" json:"input,omitempty"` // could be a long text
	Nonce                string `protobuf:"bytes,10,opt,name=nonce,proto3" json:"nonce,omitempty"`
	TransactionIndex     uint64 `protobuf:"varint,11,opt,name=transaction_index,json=transactionIndex,proto3" json:"transaction_index,omitempty"`
	TransactionType      int32  `protobuf:"varint,12,opt,name=transaction_type,json=transactionType,proto3" json:"transaction_type,omitempty"`
	Value                string `protobuf:"bytes,13,opt,name=value,proto3" json:"value,omitempty"`                                          // using string to handle big numeric values
	IndexedAt            uint64 `protobuf:"varint,14,opt,name=indexed_at,json=indexedAt,proto3" json:"indexed_at,omitempty"`                // using uint64 to represent timestamp
	BlockTimestamp       uint64 `protobuf:"varint,15,opt,name=block_timestamp,json=blockTimestamp,proto3" json:"block_timestamp,omitempty"` // using uint64 to represent timestam
}

func (x *SingleTransactionPolygon) Reset() {
	*x = SingleTransactionPolygon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocks_transactions_polygon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleTransactionPolygon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleTransactionPolygon) ProtoMessage() {}

func (x *SingleTransactionPolygon) ProtoReflect() protoreflect.Message {
	mi := &file_blocks_transactions_polygon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleTransactionPolygon.ProtoReflect.Descriptor instead.
func (*SingleTransactionPolygon) Descriptor() ([]byte, []int) {
	return file_blocks_transactions_polygon_proto_rawDescGZIP(), []int{0}
}

func (x *SingleTransactionPolygon) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *SingleTransactionPolygon) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *SingleTransactionPolygon) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *SingleTransactionPolygon) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *SingleTransactionPolygon) GetGas() string {
	if x != nil {
		return x.Gas
	}
	return ""
}

func (x *SingleTransactionPolygon) GetGasPrice() string {
	if x != nil {
		return x.GasPrice
	}
	return ""
}

func (x *SingleTransactionPolygon) GetMaxFeePerGas() string {
	if x != nil {
		return x.MaxFeePerGas
	}
	return ""
}

func (x *SingleTransactionPolygon) GetMaxPriorityFeePerGas() string {
	if x != nil {
		return x.MaxPriorityFeePerGas
	}
	return ""
}

func (x *SingleTransactionPolygon) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *SingleTransactionPolygon) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *SingleTransactionPolygon) GetTransactionIndex() uint64 {
	if x != nil {
		return x.TransactionIndex
	}
	return 0
}

func (x *SingleTransactionPolygon) GetTransactionType() int32 {
	if x != nil {
		return x.TransactionType
	}
	return 0
}

func (x *SingleTransactionPolygon) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *SingleTransactionPolygon) GetIndexedAt() uint64 {
	if x != nil {
		return x.IndexedAt
	}
	return 0
}

func (x *SingleTransactionPolygon) GetBlockTimestamp() uint64 {
	if x != nil {
		return x.BlockTimestamp
	}
	return 0
}

// Represents a single blockchain block
type BlockPolygon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumber      uint64                      `protobuf:"varint,1,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	Difficulty       uint64                      `protobuf:"varint,2,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	ExtraData        string                      `protobuf:"bytes,3,opt,name=extra_data,json=extraData,proto3" json:"extra_data,omitempty"`
	GasLimit         uint64                      `protobuf:"varint,4,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	GasUsed          uint64                      `protobuf:"varint,5,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	BaseFeePerGas    string                      `protobuf:"bytes,6,opt,name=base_fee_per_gas,json=baseFeePerGas,proto3" json:"base_fee_per_gas,omitempty"` // using string to handle big numeric values
	Hash             string                      `protobuf:"bytes,7,opt,name=hash,proto3" json:"hash,omitempty"`
	LogsBloom        string                      `protobuf:"bytes,8,opt,name=logs_bloom,json=logsBloom,proto3" json:"logs_bloom,omitempty"`
	Miner            string                      `protobuf:"bytes,9,opt,name=miner,proto3" json:"miner,omitempty"`
	Nonce            string                      `protobuf:"bytes,10,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ParentHash       string                      `protobuf:"bytes,11,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	ReceiptRoot      string                      `protobuf:"bytes,12,opt,name=receipt_root,json=receiptRoot,proto3" json:"receipt_root,omitempty"`
	Uncles           string                      `protobuf:"bytes,13,opt,name=uncles,proto3" json:"uncles,omitempty"`
	Size             int32                       `protobuf:"varint,14,opt,name=size,proto3" json:"size,omitempty"`
	StateRoot        string                      `protobuf:"bytes,15,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
	Timestamp        uint64                      `protobuf:"varint,16,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	TotalDifficulty  string                      `protobuf:"bytes,17,opt,name=total_difficulty,json=totalDifficulty,proto3" json:"total_difficulty,omitempty"`
	TransactionsRoot string                      `protobuf:"bytes,18,opt,name=transactions_root,json=transactionsRoot,proto3" json:"transactions_root,omitempty"`
	IndexedAt        uint64                      `protobuf:"varint,19,opt,name=indexed_at,json=indexedAt,proto3" json:"indexed_at,omitempty"` // using uint64 to represent timestamp
	Transactions     []*SingleTransactionPolygon `protobuf:"bytes,20,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *BlockPolygon) Reset() {
	*x = BlockPolygon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocks_transactions_polygon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockPolygon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockPolygon) ProtoMessage() {}

func (x *BlockPolygon) ProtoReflect() protoreflect.Message {
	mi := &file_blocks_transactions_polygon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockPolygon.ProtoReflect.Descriptor instead.
func (*BlockPolygon) Descriptor() ([]byte, []int) {
	return file_blocks_transactions_polygon_proto_rawDescGZIP(), []int{1}
}

func (x *BlockPolygon) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *BlockPolygon) GetDifficulty() uint64 {
	if x != nil {
		return x.Difficulty
	}
	return 0
}

func (x *BlockPolygon) GetExtraData() string {
	if x != nil {
		return x.ExtraData
	}
	return ""
}

func (x *BlockPolygon) GetGasLimit() uint64 {
	if x != nil {
		return x.GasLimit
	}
	return 0
}

func (x *BlockPolygon) GetGasUsed() uint64 {
	if x != nil {
		return x.GasUsed
	}
	return 0
}

func (x *BlockPolygon) GetBaseFeePerGas() string {
	if x != nil {
		return x.BaseFeePerGas
	}
	return ""
}

func (x *BlockPolygon) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *BlockPolygon) GetLogsBloom() string {
	if x != nil {
		return x.LogsBloom
	}
	return ""
}

func (x *BlockPolygon) GetMiner() string {
	if x != nil {
		return x.Miner
	}
	return ""
}

func (x *BlockPolygon) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *BlockPolygon) GetParentHash() string {
	if x != nil {
		return x.ParentHash
	}
	return ""
}

func (x *BlockPolygon) GetReceiptRoot() string {
	if x != nil {
		return x.ReceiptRoot
	}
	return ""
}

func (x *BlockPolygon) GetUncles() string {
	if x != nil {
		return x.Uncles
	}
	return ""
}

func (x *BlockPolygon) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *BlockPolygon) GetStateRoot() string {
	if x != nil {
		return x.StateRoot
	}
	return ""
}

func (x *BlockPolygon) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *BlockPolygon) GetTotalDifficulty() string {
	if x != nil {
		return x.TotalDifficulty
	}
	return ""
}

func (x *BlockPolygon) GetTransactionsRoot() string {
	if x != nil {
		return x.TransactionsRoot
	}
	return ""
}

func (x *BlockPolygon) GetIndexedAt() uint64 {
	if x != nil {
		return x.IndexedAt
	}
	return 0
}

func (x *BlockPolygon) GetTransactions() []*SingleTransactionPolygon {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type EventLogPolygon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address          string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`                                            // The address of the contract that generated the log
	Topics           []string `protobuf:"bytes,2,rep,name=topics,proto3" json:"topics,omitempty"`                                              // Topics are indexed parameters during log generation
	Data             string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`                                                  // The data field from the log
	BlockNumber      uint64   `protobuf:"varint,4,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`                // The block number where this log was in
	TransactionHash  string   `protobuf:"bytes,5,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash,omitempty"`     // The hash of the transaction that generated this log
	BlockHash        string   `protobuf:"bytes,6,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`                       // The hash of the block where this log was in
	Removed          bool     `protobuf:"varint,7,opt,name=removed,proto3" json:"removed,omitempty"`                                           // True if the log was reverted due to a chain reorganization
	LogIndex         uint64   `protobuf:"varint,8,opt,name=log_index,json=logIndex,proto3" json:"log_index,omitempty"`                         // The index of the log in the block
	TransactionIndex uint64   `protobuf:"varint,9,opt,name=transaction_index,json=transactionIndex,proto3" json:"transaction_index,omitempty"` // The index of the transaction in the block
}

func (x *EventLogPolygon) Reset() {
	*x = EventLogPolygon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocks_transactions_polygon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventLogPolygon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventLogPolygon) ProtoMessage() {}

func (x *EventLogPolygon) ProtoReflect() protoreflect.Message {
	mi := &file_blocks_transactions_polygon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventLogPolygon.ProtoReflect.Descriptor instead.
func (*EventLogPolygon) Descriptor() ([]byte, []int) {
	return file_blocks_transactions_polygon_proto_rawDescGZIP(), []int{2}
}

func (x *EventLogPolygon) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *EventLogPolygon) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

func (x *EventLogPolygon) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *EventLogPolygon) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *EventLogPolygon) GetTransactionHash() string {
	if x != nil {
		return x.TransactionHash
	}
	return ""
}

func (x *EventLogPolygon) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *EventLogPolygon) GetRemoved() bool {
	if x != nil {
		return x.Removed
	}
	return false
}

func (x *EventLogPolygon) GetLogIndex() uint64 {
	if x != nil {
		return x.LogIndex
	}
	return 0
}

func (x *EventLogPolygon) GetTransactionIndex() uint64 {
	if x != nil {
		return x.TransactionIndex
	}
	return 0
}

var File_blocks_transactions_polygon_proto protoreflect.FileDescriptor

var file_blocks_transactions_polygon_proto_rawDesc = []byte{
	0x0a, 0x21, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x83, 0x04, 0x0a, 0x18, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66,
	0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x61, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x61, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x67,
	0x61, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x67, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f,
	0x66, 0x65, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x67, 0x61, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x46, 0x65, 0x65, 0x50, 0x65, 0x72, 0x47, 0x61, 0x73, 0x12,
	0x36, 0x0a, 0x18, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f,
	0x66, 0x65, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x67, 0x61, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x14, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x46, 0x65,
	0x65, 0x50, 0x65, 0x72, 0x47, 0x61, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f,
	0x6e, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x29, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x27, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x93, 0x05, 0x0a, 0x0c, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1e, 0x0a,
	0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x1d, 0x0a,
	0x0a, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09,
	0x67, 0x61, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x67, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x61, 0x73,
	0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x67, 0x61, 0x73,
	0x55, 0x73, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x10, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x66, 0x65, 0x65,
	0x5f, 0x70, 0x65, 0x72, 0x5f, 0x67, 0x61, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x62, 0x61, 0x73, 0x65, 0x46, 0x65, 0x65, 0x50, 0x65, 0x72, 0x47, 0x61, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x6f, 0x6d, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x73, 0x42, 0x6c, 0x6f, 0x6f, 0x6d,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x6f, 0x6f, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x6e, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x10, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63,
	0x75, 0x6c, 0x74, 0x79, 0x12, 0x2b, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x6f, 0x6f,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x3d, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f,
	0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0xa8, 0x02, 0x0a, 0x0f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x50, 0x6f, 0x6c, 0x79,
	0x67, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2b, 0x0a,
	0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blocks_transactions_polygon_proto_rawDescOnce sync.Once
	file_blocks_transactions_polygon_proto_rawDescData = file_blocks_transactions_polygon_proto_rawDesc
)

func file_blocks_transactions_polygon_proto_rawDescGZIP() []byte {
	file_blocks_transactions_polygon_proto_rawDescOnce.Do(func() {
		file_blocks_transactions_polygon_proto_rawDescData = protoimpl.X.CompressGZIP(file_blocks_transactions_polygon_proto_rawDescData)
	})
	return file_blocks_transactions_polygon_proto_rawDescData
}

var file_blocks_transactions_polygon_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_blocks_transactions_polygon_proto_goTypes = []interface{}{
	(*SingleTransactionPolygon)(nil), // 0: SingleTransactionPolygon
	(*BlockPolygon)(nil),             // 1: BlockPolygon
	(*EventLogPolygon)(nil),          // 2: EventLogPolygon
}
var file_blocks_transactions_polygon_proto_depIdxs = []int32{
	0, // 0: BlockPolygon.transactions:type_name -> SingleTransactionPolygon
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_blocks_transactions_polygon_proto_init() }
func file_blocks_transactions_polygon_proto_init() {
	if File_blocks_transactions_polygon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blocks_transactions_polygon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleTransactionPolygon); i {
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
		file_blocks_transactions_polygon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockPolygon); i {
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
		file_blocks_transactions_polygon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventLogPolygon); i {
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
			RawDescriptor: file_blocks_transactions_polygon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blocks_transactions_polygon_proto_goTypes,
		DependencyIndexes: file_blocks_transactions_polygon_proto_depIdxs,
		MessageInfos:      file_blocks_transactions_polygon_proto_msgTypes,
	}.Build()
	File_blocks_transactions_polygon_proto = out.File
	file_blocks_transactions_polygon_proto_rawDesc = nil
	file_blocks_transactions_polygon_proto_goTypes = nil
	file_blocks_transactions_polygon_proto_depIdxs = nil
}
