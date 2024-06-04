// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.6.1
// source: blockchain/xai_sepolia/xai_sepolia_index_types.proto

package xai_sepolia

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

type XaiSepoliaTransactionAccessList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address     string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	StorageKeys []string `protobuf:"bytes,2,rep,name=storage_keys,json=storageKeys,proto3" json:"storage_keys,omitempty"`
}

func (x *XaiSepoliaTransactionAccessList) Reset() {
	*x = XaiSepoliaTransactionAccessList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XaiSepoliaTransactionAccessList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XaiSepoliaTransactionAccessList) ProtoMessage() {}

func (x *XaiSepoliaTransactionAccessList) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XaiSepoliaTransactionAccessList.ProtoReflect.Descriptor instead.
func (*XaiSepoliaTransactionAccessList) Descriptor() ([]byte, []int) {
	return file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_rawDescGZIP(), []int{0}
}

func (x *XaiSepoliaTransactionAccessList) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *XaiSepoliaTransactionAccessList) GetStorageKeys() []string {
	if x != nil {
		return x.StorageKeys
	}
	return nil
}

// Represents a single transaction within a block
type XaiSepoliaTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash                 string                             `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`                                                                   // The hash of the transaction
	BlockNumber          uint64                             `protobuf:"varint,2,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`                                 // The block number the transaction is in
	FromAddress          string                             `protobuf:"bytes,3,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`                                  // The address the transaction is sent from
	ToAddress            string                             `protobuf:"bytes,4,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`                                        // The address the transaction is sent to
	Gas                  string                             `protobuf:"bytes,5,opt,name=gas,proto3" json:"gas,omitempty"`                                                                     // The gas limit of the transaction
	GasPrice             string                             `protobuf:"bytes,6,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`                                           // The gas price of the transaction
	MaxFeePerGas         string                             `protobuf:"bytes,7,opt,name=max_fee_per_gas,json=maxFeePerGas,proto3" json:"max_fee_per_gas,omitempty"`                           // Used as a field to match potential EIP-1559 transaction types
	MaxPriorityFeePerGas string                             `protobuf:"bytes,8,opt,name=max_priority_fee_per_gas,json=maxPriorityFeePerGas,proto3" json:"max_priority_fee_per_gas,omitempty"` // Used as a field to match potential EIP-1559 transaction types
	Input                string                             `protobuf:"bytes,9,opt,name=input,proto3" json:"input,omitempty"`                                                                 // The input data of the transaction
	Nonce                string                             `protobuf:"bytes,10,opt,name=nonce,proto3" json:"nonce,omitempty"`                                                                // The nonce of the transaction
	TransactionIndex     uint64                             `protobuf:"varint,11,opt,name=transaction_index,json=transactionIndex,proto3" json:"transaction_index,omitempty"`                 // The index of the transaction in the block
	TransactionType      uint64                             `protobuf:"varint,12,opt,name=transaction_type,json=transactionType,proto3" json:"transaction_type,omitempty"`                    // Field to match potential EIP-1559 transaction types
	Value                string                             `protobuf:"bytes,13,opt,name=value,proto3" json:"value,omitempty"`                                                                // The value of the transaction
	IndexedAt            uint64                             `protobuf:"varint,14,opt,name=indexed_at,json=indexedAt,proto3" json:"indexed_at,omitempty"`                                      // When the transaction was indexed by crawler
	BlockTimestamp       uint64                             `protobuf:"varint,15,opt,name=block_timestamp,json=blockTimestamp,proto3" json:"block_timestamp,omitempty"`                       // The timestamp of this block
	BlockHash            string                             `protobuf:"bytes,16,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`                                       // The hash of the block the transaction is in
	ChainId              string                             `protobuf:"bytes,17,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`                                             // Used as a field to match potential EIP-1559 transaction types
	V                    string                             `protobuf:"bytes,18,opt,name=v,proto3" json:"v,omitempty"`                                                                        // Used as a field to match potential EIP-1559 transaction types
	R                    string                             `protobuf:"bytes,19,opt,name=r,proto3" json:"r,omitempty"`                                                                        // Used as a field to match potential EIP-1559 transaction types
	S                    string                             `protobuf:"bytes,20,opt,name=s,proto3" json:"s,omitempty"`                                                                        // Used as a field to match potential EIP-1559 transaction types
	AccessList           []*XaiSepoliaTransactionAccessList `protobuf:"bytes,21,rep,name=access_list,json=accessList,proto3" json:"access_list,omitempty"`
	YParity              string                             `protobuf:"bytes,22,opt,name=y_parity,json=yParity,proto3" json:"y_parity,omitempty"` // Used as a field to match potential EIP-1559 transaction types
}

func (x *XaiSepoliaTransaction) Reset() {
	*x = XaiSepoliaTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XaiSepoliaTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XaiSepoliaTransaction) ProtoMessage() {}

func (x *XaiSepoliaTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XaiSepoliaTransaction.ProtoReflect.Descriptor instead.
func (*XaiSepoliaTransaction) Descriptor() ([]byte, []int) {
	return file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_rawDescGZIP(), []int{1}
}

func (x *XaiSepoliaTransaction) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *XaiSepoliaTransaction) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *XaiSepoliaTransaction) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *XaiSepoliaTransaction) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *XaiSepoliaTransaction) GetGas() string {
	if x != nil {
		return x.Gas
	}
	return ""
}

func (x *XaiSepoliaTransaction) GetGasPrice() string {
	if x != nil {
		return x.GasPrice
	}
	return ""
}

func (x *XaiSepoliaTransaction) GetMaxFeePerGas() string {
	if x != nil {
		return x.MaxFeePerGas
	}
	return ""
}

func (x *XaiSepoliaTransaction) GetMaxPriorityFeePerGas() string {
	if x != nil {
		return x.MaxPriorityFeePerGas
	}
	return ""
}

func (x *XaiSepoliaTransaction) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *XaiSepoliaTransaction) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *XaiSepoliaTransaction) GetTransactionIndex() uint64 {
	if x != nil {
		return x.TransactionIndex
	}
	return 0
}

func (x *XaiSepoliaTransaction) GetTransactionType() uint64 {
	if x != nil {
		return x.TransactionType
	}
	return 0
}

func (x *XaiSepoliaTransaction) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *XaiSepoliaTransaction) GetIndexedAt() uint64 {
	if x != nil {
		return x.IndexedAt
	}
	return 0
}

func (x *XaiSepoliaTransaction) GetBlockTimestamp() uint64 {
	if x != nil {
		return x.BlockTimestamp
	}
	return 0
}

func (x *XaiSepoliaTransaction) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *XaiSepoliaTransaction) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *XaiSepoliaTransaction) GetV() string {
	if x != nil {
		return x.V
	}
	return ""
}

func (x *XaiSepoliaTransaction) GetR() string {
	if x != nil {
		return x.R
	}
	return ""
}

func (x *XaiSepoliaTransaction) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

func (x *XaiSepoliaTransaction) GetAccessList() []*XaiSepoliaTransactionAccessList {
	if x != nil {
		return x.AccessList
	}
	return nil
}

func (x *XaiSepoliaTransaction) GetYParity() string {
	if x != nil {
		return x.YParity
	}
	return ""
}

// Represents a block in the blockchain
type XaiSepoliaBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumber      uint64                   `protobuf:"varint,1,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`          // The block number
	Difficulty       uint64                   `protobuf:"varint,2,opt,name=difficulty,proto3" json:"difficulty,omitempty"`                               // The difficulty of this block
	ExtraData        string                   `protobuf:"bytes,3,opt,name=extra_data,json=extraData,proto3" json:"extra_data,omitempty"`                 // Extra data included in the block
	GasLimit         uint64                   `protobuf:"varint,4,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`                   // The gas limit for this block
	GasUsed          uint64                   `protobuf:"varint,5,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`                      // The total gas used by all transactions in this block
	BaseFeePerGas    string                   `protobuf:"bytes,6,opt,name=base_fee_per_gas,json=baseFeePerGas,proto3" json:"base_fee_per_gas,omitempty"` // The base fee per gas for this block
	Hash             string                   `protobuf:"bytes,7,opt,name=hash,proto3" json:"hash,omitempty"`                                            // The hash of this block
	LogsBloom        string                   `protobuf:"bytes,8,opt,name=logs_bloom,json=logsBloom,proto3" json:"logs_bloom,omitempty"`                 // The logs bloom filter for this block
	Miner            string                   `protobuf:"bytes,9,opt,name=miner,proto3" json:"miner,omitempty"`                                          // The address of the miner who mined this block
	Nonce            string                   `protobuf:"bytes,10,opt,name=nonce,proto3" json:"nonce,omitempty"`                                         // The nonce of this block
	ParentHash       string                   `protobuf:"bytes,11,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`             // The hash of the parent block
	ReceiptsRoot     string                   `protobuf:"bytes,12,opt,name=receipts_root,json=receiptsRoot,proto3" json:"receipts_root,omitempty"`       // The root hash of the receipts trie
	Sha3Uncles       string                   `protobuf:"bytes,13,opt,name=sha3_uncles,json=sha3Uncles,proto3" json:"sha3_uncles,omitempty"`             // The SHA3 hash of the uncles data in this block
	Size             uint64                   `protobuf:"varint,14,opt,name=size,proto3" json:"size,omitempty"`                                          // The size of this block
	StateRoot        string                   `protobuf:"bytes,15,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`                // The root hash of the state trie
	Timestamp        uint64                   `protobuf:"varint,16,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	TotalDifficulty  string                   `protobuf:"bytes,17,opt,name=total_difficulty,json=totalDifficulty,proto3" json:"total_difficulty,omitempty"`    // The total difficulty of the chain until this block
	TransactionsRoot string                   `protobuf:"bytes,18,opt,name=transactions_root,json=transactionsRoot,proto3" json:"transactions_root,omitempty"` // The root hash of the transactions trie
	IndexedAt        uint64                   `protobuf:"varint,19,opt,name=indexed_at,json=indexedAt,proto3" json:"indexed_at,omitempty"`                     // When the block was indexed by crawler
	Transactions     []*XaiSepoliaTransaction `protobuf:"bytes,20,rep,name=transactions,proto3" json:"transactions,omitempty"`                                 // The transactions included in this block
	MixHash          string                   `protobuf:"bytes,21,opt,name=mix_hash,json=mixHash,proto3" json:"mix_hash,omitempty"`                            // The timestamp of this block
	SendCount        string                   `protobuf:"bytes,22,opt,name=send_count,json=sendCount,proto3" json:"send_count,omitempty"`                      // The number of sends in this block
	SendRoot         string                   `protobuf:"bytes,23,opt,name=send_root,json=sendRoot,proto3" json:"send_root,omitempty"`                         // The root hash of the sends trie
	L1BlockNumber    uint64                   `protobuf:"varint,24,opt,name=l1_block_number,json=l1BlockNumber,proto3" json:"l1_block_number,omitempty"`       // The block number of the corresponding L1 block
}

func (x *XaiSepoliaBlock) Reset() {
	*x = XaiSepoliaBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XaiSepoliaBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XaiSepoliaBlock) ProtoMessage() {}

func (x *XaiSepoliaBlock) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XaiSepoliaBlock.ProtoReflect.Descriptor instead.
func (*XaiSepoliaBlock) Descriptor() ([]byte, []int) {
	return file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_rawDescGZIP(), []int{2}
}

func (x *XaiSepoliaBlock) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *XaiSepoliaBlock) GetDifficulty() uint64 {
	if x != nil {
		return x.Difficulty
	}
	return 0
}

func (x *XaiSepoliaBlock) GetExtraData() string {
	if x != nil {
		return x.ExtraData
	}
	return ""
}

func (x *XaiSepoliaBlock) GetGasLimit() uint64 {
	if x != nil {
		return x.GasLimit
	}
	return 0
}

func (x *XaiSepoliaBlock) GetGasUsed() uint64 {
	if x != nil {
		return x.GasUsed
	}
	return 0
}

func (x *XaiSepoliaBlock) GetBaseFeePerGas() string {
	if x != nil {
		return x.BaseFeePerGas
	}
	return ""
}

func (x *XaiSepoliaBlock) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *XaiSepoliaBlock) GetLogsBloom() string {
	if x != nil {
		return x.LogsBloom
	}
	return ""
}

func (x *XaiSepoliaBlock) GetMiner() string {
	if x != nil {
		return x.Miner
	}
	return ""
}

func (x *XaiSepoliaBlock) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *XaiSepoliaBlock) GetParentHash() string {
	if x != nil {
		return x.ParentHash
	}
	return ""
}

func (x *XaiSepoliaBlock) GetReceiptsRoot() string {
	if x != nil {
		return x.ReceiptsRoot
	}
	return ""
}

func (x *XaiSepoliaBlock) GetSha3Uncles() string {
	if x != nil {
		return x.Sha3Uncles
	}
	return ""
}

func (x *XaiSepoliaBlock) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *XaiSepoliaBlock) GetStateRoot() string {
	if x != nil {
		return x.StateRoot
	}
	return ""
}

func (x *XaiSepoliaBlock) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *XaiSepoliaBlock) GetTotalDifficulty() string {
	if x != nil {
		return x.TotalDifficulty
	}
	return ""
}

func (x *XaiSepoliaBlock) GetTransactionsRoot() string {
	if x != nil {
		return x.TransactionsRoot
	}
	return ""
}

func (x *XaiSepoliaBlock) GetIndexedAt() uint64 {
	if x != nil {
		return x.IndexedAt
	}
	return 0
}

func (x *XaiSepoliaBlock) GetTransactions() []*XaiSepoliaTransaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *XaiSepoliaBlock) GetMixHash() string {
	if x != nil {
		return x.MixHash
	}
	return ""
}

func (x *XaiSepoliaBlock) GetSendCount() string {
	if x != nil {
		return x.SendCount
	}
	return ""
}

func (x *XaiSepoliaBlock) GetSendRoot() string {
	if x != nil {
		return x.SendRoot
	}
	return ""
}

func (x *XaiSepoliaBlock) GetL1BlockNumber() uint64 {
	if x != nil {
		return x.L1BlockNumber
	}
	return 0
}

type XaiSepoliaEventLog struct {
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

func (x *XaiSepoliaEventLog) Reset() {
	*x = XaiSepoliaEventLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XaiSepoliaEventLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XaiSepoliaEventLog) ProtoMessage() {}

func (x *XaiSepoliaEventLog) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XaiSepoliaEventLog.ProtoReflect.Descriptor instead.
func (*XaiSepoliaEventLog) Descriptor() ([]byte, []int) {
	return file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_rawDescGZIP(), []int{3}
}

func (x *XaiSepoliaEventLog) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *XaiSepoliaEventLog) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

func (x *XaiSepoliaEventLog) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *XaiSepoliaEventLog) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *XaiSepoliaEventLog) GetTransactionHash() string {
	if x != nil {
		return x.TransactionHash
	}
	return ""
}

func (x *XaiSepoliaEventLog) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *XaiSepoliaEventLog) GetRemoved() bool {
	if x != nil {
		return x.Removed
	}
	return false
}

func (x *XaiSepoliaEventLog) GetLogIndex() uint64 {
	if x != nil {
		return x.LogIndex
	}
	return 0
}

func (x *XaiSepoliaEventLog) GetTransactionIndex() uint64 {
	if x != nil {
		return x.TransactionIndex
	}
	return 0
}

var File_blockchain_xai_sepolia_xai_sepolia_index_types_proto protoreflect.FileDescriptor

var file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_rawDesc = []byte{
	0x0a, 0x34, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x78, 0x61, 0x69,
	0x5f, 0x73, 0x65, 0x70, 0x6f, 0x6c, 0x69, 0x61, 0x2f, 0x78, 0x61, 0x69, 0x5f, 0x73, 0x65, 0x70,
	0x6f, 0x6c, 0x69, 0x61, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x1f, 0x58, 0x61, 0x69, 0x53, 0x65, 0x70,
	0x6f, 0x6c, 0x69, 0x61, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6b,
	0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x22, 0xc2, 0x05, 0x0a, 0x15, 0x58, 0x61, 0x69, 0x53, 0x65,
	0x70, 0x6f, 0x6c, 0x69, 0x61, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
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
	0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x27, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x49, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01,
	0x76, 0x12, 0x0c, 0x0a, 0x01, 0x72, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x72, 0x12,
	0x0c, 0x0a, 0x01, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x73, 0x12, 0x41, 0x0a,
	0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x15, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x58, 0x61, 0x69, 0x53, 0x65, 0x70, 0x6f, 0x6c, 0x69, 0x61, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x79, 0x50, 0x61, 0x72, 0x69, 0x74, 0x79, 0x22, 0x9d, 0x06, 0x0a, 0x0f,
	0x58, 0x61, 0x69, 0x53, 0x65, 0x70, 0x6f, 0x6c, 0x69, 0x61, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c,
	0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x67, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x67, 0x61, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x67, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x10, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x66, 0x65, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x67, 0x61, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x46, 0x65, 0x65, 0x50, 0x65, 0x72, 0x47,
	0x61, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x62,
	0x6c, 0x6f, 0x6f, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x73,
	0x42, 0x6c, 0x6f, 0x6f, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6e,
	0x6f, 0x6e, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x5f, 0x72,
	0x6f, 0x6f, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x73, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x61, 0x33, 0x5f,
	0x75, 0x6e, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x68,
	0x61, 0x33, 0x55, 0x6e, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
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
	0x12, 0x3a, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x58, 0x61, 0x69, 0x53, 0x65, 0x70, 0x6f,
	0x6c, 0x69, 0x61, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19, 0x0a, 0x08,
	0x6d, 0x69, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x69, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x6e,
	0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x72,
	0x6f, 0x6f, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x52,
	0x6f, 0x6f, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x31, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x18, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x6c, 0x31,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xab, 0x02, 0x0a, 0x12,
	0x58, 0x61, 0x69, 0x53, 0x65, 0x70, 0x6f, 0x6c, 0x69, 0x61, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c,
	0x6f, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2b, 0x0a, 0x11,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2d, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x65, 0x72, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x78, 0x61, 0x69, 0x5f, 0x73, 0x65, 0x70, 0x6f, 0x6c, 0x69,
	0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_rawDescOnce sync.Once
	file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_rawDescData = file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_rawDesc
)

func file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_rawDescGZIP() []byte {
	file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_rawDescOnce.Do(func() {
		file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_rawDescData)
	})
	return file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_rawDescData
}

var file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_goTypes = []interface{}{
	(*XaiSepoliaTransactionAccessList)(nil), // 0: XaiSepoliaTransactionAccessList
	(*XaiSepoliaTransaction)(nil),           // 1: XaiSepoliaTransaction
	(*XaiSepoliaBlock)(nil),                 // 2: XaiSepoliaBlock
	(*XaiSepoliaEventLog)(nil),              // 3: XaiSepoliaEventLog
}
var file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_depIdxs = []int32{
	0, // 0: XaiSepoliaTransaction.access_list:type_name -> XaiSepoliaTransactionAccessList
	1, // 1: XaiSepoliaBlock.transactions:type_name -> XaiSepoliaTransaction
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_init() }
func file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_init() {
	if File_blockchain_xai_sepolia_xai_sepolia_index_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XaiSepoliaTransactionAccessList); i {
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
		file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XaiSepoliaTransaction); i {
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
		file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XaiSepoliaBlock); i {
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
		file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XaiSepoliaEventLog); i {
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
			RawDescriptor: file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_goTypes,
		DependencyIndexes: file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_depIdxs,
		MessageInfos:      file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_msgTypes,
	}.Build()
	File_blockchain_xai_sepolia_xai_sepolia_index_types_proto = out.File
	file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_rawDesc = nil
	file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_goTypes = nil
	file_blockchain_xai_sepolia_xai_sepolia_index_types_proto_depIdxs = nil
}
