package mantle_sepolia

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	seer_common "github.com/moonstream-to/seer/blockchain/common"
	"github.com/moonstream-to/seer/indexer"
	"google.golang.org/protobuf/proto"
)

func NewClient(url string) (*Client, error) {
	rpcClient, err := rpc.DialContext(context.Background(), url)
	if err != nil {
		return nil, err
	}
	return &Client{rpcClient: rpcClient}, nil
}

// Client is a wrapper around the Ethereum JSON-RPC client.

type Client struct {
	rpcClient *rpc.Client
}

// Client common

// ChainType returns the chain type.
func (c *Client) ChainType() string {
	return "mantle_sepolia"
}

// Close closes the underlying RPC client.
func (c *Client) Close() {
	c.rpcClient.Close()
}

// GetLatestBlockNumber returns the latest block number.
func (c *Client) GetLatestBlockNumber() (*big.Int, error) {
	var result string
	if err := c.rpcClient.CallContext(context.Background(), &result, "eth_blockNumber"); err != nil {
		return nil, err
	}

	// Convert the hex string to a big.Int
	blockNumber, ok := new(big.Int).SetString(result, 0) // The 0 base lets the function infer the base from the string prefix.
	if !ok {
		return nil, fmt.Errorf("invalid block number format: %s", result)
	}

	return blockNumber, nil
}

// BlockByNumber returns the block with the given number.
func (c *Client) GetBlockByNumber(ctx context.Context, number *big.Int) (*seer_common.BlockJson, error) {

	var rawResponse json.RawMessage // Use RawMessage to capture the entire JSON response
	err := c.rpcClient.CallContext(ctx, &rawResponse, "eth_getBlockByNumber", "0x"+number.Text(16), true)
	if err != nil {
		fmt.Println("Error calling eth_getBlockByNumber: ", err)
		return nil, err
	}

	var response_json map[string]interface{}

	err = json.Unmarshal(rawResponse, &response_json)

	delete(response_json, "transactions")

	var block *seer_common.BlockJson
	err = c.rpcClient.CallContext(ctx, &block, "eth_getBlockByNumber", "0x"+number.Text(16), true) // true to include transactions
	return block, err
}

// BlockByHash returns the block with the given hash.
func (c *Client) BlockByHash(ctx context.Context, hash common.Hash) (*seer_common.BlockJson, error) {
	var block *seer_common.BlockJson
	err := c.rpcClient.CallContext(ctx, &block, "eth_getBlockByHash", hash, true) // true to include transactions
	return block, err
}

// TransactionReceipt returns the receipt of a transaction by transaction hash.
func (c *Client) TransactionReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	var receipt *types.Receipt
	err := c.rpcClient.CallContext(ctx, &receipt, "eth_getTransactionReceipt", hash)
	return receipt, err
}

func (c *Client) ClientFilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]*seer_common.EventJson, error) {
	var logs []*seer_common.EventJson
	fromBlock := q.FromBlock
	toBlock := q.ToBlock
	batchStep := new(big.Int).Sub(toBlock, fromBlock) // Calculate initial batch step

	for {
		// Calculate the next "lastBlock" within the batch step or adjust to "toBlock" if exceeding
		nextBlock := new(big.Int).Add(fromBlock, batchStep)
		if nextBlock.Cmp(toBlock) > 0 {
			nextBlock = new(big.Int).Set(toBlock)
		}

		var result []*seer_common.EventJson
		err := c.rpcClient.CallContext(ctx, &result, "eth_getLogs", struct {
			FromBlock string           `json:"fromBlock"`
			ToBlock   string           `json:"toBlock"`
			Addresses []common.Address `json:"addresses"`
			Topics    [][]common.Hash  `json:"topics"`
		}{
			FromBlock: toHex(fromBlock),
			ToBlock:   toHex(nextBlock),
			Addresses: q.Addresses,
			Topics:    q.Topics,
		})

		if err != nil {
			if strings.Contains(err.Error(), "query returned more than 10000 results") {
				// Halve the batch step if too many results and retry
				batchStep.Div(batchStep, big.NewInt(2))
				if batchStep.Cmp(big.NewInt(1)) < 0 {
					// If the batch step is too small we will skip that block
					fromBlock = new(big.Int).Add(nextBlock, big.NewInt(1))
					if fromBlock.Cmp(toBlock) > 0 {
						break
					}
					continue
				}
				continue
			} else {
				// For any other error, return immediately
				return nil, err
			}
		}

		// Append the results and adjust "fromBlock" for the next batch
		logs = append(logs, result...)
		fromBlock = new(big.Int).Add(nextBlock, big.NewInt(1))

		// Break the loop if we've reached or exceeded "toBlock"
		if fromBlock.Cmp(toBlock) > 0 {
			break
		}
	}

	return logs, nil
}

// fetchBlocks returns the blocks for a given range.
func (c *Client) fetchBlocks(ctx context.Context, from, to *big.Int) ([]*seer_common.BlockJson, error) {
	var blocks []*seer_common.BlockJson

	for i := from; i.Cmp(to) <= 0; i.Add(i, big.NewInt(1)) {
		block, err := c.GetBlockByNumber(ctx, i)
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, block)
	}
	return blocks, nil
}

// Utility function to convert big.Int to its hexadecimal representation.
func toHex(number *big.Int) string {
	return fmt.Sprintf("0x%x", number)
}

func fromHex(hex string) *big.Int {
	number := new(big.Int)
	number.SetString(hex, 0)
	return number
}

// FetchBlocksInRange fetches blocks within a specified range.
// This could be useful for batch processing or analysis.
func (c *Client) FetchBlocksInRange(from, to *big.Int) ([]*seer_common.BlockJson, error) {
	var blocks []*seer_common.BlockJson
	ctx := context.Background() // For simplicity, using a background context; consider timeouts for production.

	for i := new(big.Int).Set(from); i.Cmp(to) <= 0; i.Add(i, big.NewInt(1)) {
		block, err := c.GetBlockByNumber(ctx, i)
		fmt.Println("Block number: ", i)
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, block)
	}

	return blocks, nil
}

// ParseBlocksAndTransactions parses blocks and their transactions into custom data structures.
// This method showcases how to handle and transform detailed block and transaction data.
func (c *Client) ParseBlocksAndTransactions(from, to *big.Int) ([]*MantleSepoliaBlock, []*MantleSepoliaTransaction, error) {
	blocksJson, err := c.FetchBlocksInRange(from, to)
	if err != nil {
		return nil, nil, err
	}

	var parsedBlocks []*MantleSepoliaBlock
	var parsedTransactions []*MantleSepoliaTransaction
	for _, blockJson := range blocksJson {
		// Convert BlockJson to Block and Transactions as required.
		parsedBlock := ToProtoSingleBlock(blockJson)

		// Example: Parsing transactions within the block
		for _, txJson := range blockJson.Transactions {

			txJson.BlockTimestamp = blockJson.Timestamp

			parsedTransaction := ToProtoSingleTransaction(&txJson)
			parsedTransactions = append(parsedTransactions, parsedTransaction)
		}

		parsedBlocks = append(parsedBlocks, parsedBlock)
	}

	return parsedBlocks, parsedTransactions, nil
}

func (c *Client) ParseEvents(from, to *big.Int, blocksCache map[uint64]indexer.BlockCache) ([]*MantleSepoliaEventLog, error) {

	logs, err := c.ClientFilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: from,
		ToBlock:   to,
	})

	if err != nil {
		fmt.Println("Error fetching logs: ", err)
		return nil, err
	}

	var parsedEvents []*MantleSepoliaEventLog
	for _, log := range logs {
		parsedEvent := ToProtoSingleEventLog(log)
		parsedEvents = append(parsedEvents, parsedEvent)
	}

	return parsedEvents, nil
}

func (c *Client) FetchAsProtoBlocks(from, to *big.Int) ([]proto.Message, []proto.Message, []indexer.BlockIndex, []indexer.TransactionIndex, map[uint64]indexer.BlockCache, error) {
	parsedBlocks, parsedTransactions, err := c.ParseBlocksAndTransactions(from, to)

	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	blocksCache := make(map[uint64]indexer.BlockCache)

	var blocksProto []proto.Message
	var blockIndex []indexer.BlockIndex
	for index, block := range parsedBlocks {
		blocksProto = append(blocksProto, block) // Assuming block is already a proto.Message
		blocksCache[block.BlockNumber] = indexer.BlockCache{
			BlockNumber:    block.BlockNumber,
			BlockHash:      block.Hash,
			BlockTimestamp: block.Timestamp,
		} // Assuming block.BlockNumber is int64 and block.Hash is string
		blockIndex = append(blockIndex, indexer.NewBlockIndex("mantle_sepolia",
			block.BlockNumber,
			block.Hash,
			block.Timestamp,
			block.ParentHash,
			uint64(index),
			"",
			0,
		))
	}

	var transactionsProto []proto.Message
	var transactionIndex []indexer.TransactionIndex
	for index, transaction := range parsedTransactions {
		transactionsProto = append(transactionsProto, transaction) // Assuming transaction is also a proto.Message

		selector := "0x"

		if len(transaction.Input) > 10 {
			selector = transaction.Input[:10]
		}

		transactionIndex = append(transactionIndex, indexer.TransactionIndex{
			BlockNumber:      transaction.BlockNumber,
			BlockHash:        transaction.BlockHash,
			BlockTimestamp:   transaction.BlockTimestamp,
			FromAddress:      transaction.FromAddress,
			ToAddress:        transaction.ToAddress,
			RowID:            uint64(index),
			Selector:         selector, // First 10 characters of the input data 0x + 4 bytes of the function signature
			TransactionHash:  transaction.Hash,
			TransactionIndex: transaction.TransactionIndex,
			Type:             transaction.TransactionType,
			Path:             "",
		})
	}

	return blocksProto, transactionsProto, blockIndex, transactionIndex, blocksCache, nil
}

func (c *Client) FetchAsProtoEvents(from, to *big.Int, blocksCahche map[uint64]indexer.BlockCache) ([]proto.Message, []indexer.LogIndex, error) {

	parsedEvents, err := c.ParseEvents(from, to, blocksCahche)

	if err != nil {
		return nil, nil, err
	}

	var eventsProto []proto.Message
	var eventsIndex []indexer.LogIndex
	for index, event := range parsedEvents {
		eventsProto = append(eventsProto, event) // Assuming event is already a proto.Message

		var topic0, topic1, topic2 *string

		if len(event.Topics) == 0 {
			fmt.Println("No topics found for event: ", event)
		} else {
			topic0 = &event.Topics[0] // First topic
		}

		// Assign topics based on availability
		if len(event.Topics) > 1 {
			topic1 = &event.Topics[1] // Second topic, if present
		}
		if len(event.Topics) > 2 {
			topic2 = &event.Topics[2] // Third topic, if present
		}

		eventsIndex = append(eventsIndex, indexer.LogIndex{
			Address:         event.Address,
			BlockNumber:     event.BlockNumber,
			BlockHash:       event.BlockHash,
			BlockTimestamp:  blocksCahche[event.BlockNumber].BlockTimestamp,
			TransactionHash: event.TransactionHash,
			Selector:        topic0, // First topic
			Topic1:          topic1,
			Topic2:          topic2,
			RowID:           uint64(index),
			LogIndex:        event.LogIndex,
			Path:            "",
		})
	}

	return eventsProto, eventsIndex, nil

}
func ToProtoSingleBlock(obj *seer_common.BlockJson) *MantleSepoliaBlock {
	return &MantleSepoliaBlock{
		BlockNumber:      fromHex(obj.BlockNumber).Uint64(),
		Difficulty:       fromHex(obj.Difficulty).Uint64(),
		ExtraData:        obj.ExtraData,
		GasLimit:         fromHex(obj.GasLimit).Uint64(),
		GasUsed:          fromHex(obj.GasUsed).Uint64(),
		BaseFeePerGas:    obj.BaseFeePerGas,
		Hash:             obj.Hash,
		LogsBloom:        obj.LogsBloom,
		Miner:            obj.Miner,
		Nonce:            obj.Nonce,
		ParentHash:       obj.ParentHash,
		ReceiptsRoot:     obj.ReceiptsRoot,
		Sha3Uncles:       obj.Sha3Uncles,
		Size:             fromHex(obj.Size).Uint64(),
		StateRoot:        obj.StateRoot,
		Timestamp:        fromHex(obj.Timestamp).Uint64(),
		TotalDifficulty:  obj.TotalDifficulty,
		TransactionsRoot: obj.TransactionsRoot,
		IndexedAt:        fromHex(obj.IndexedAt).Uint64(),
	}
}

func ToProtoSingleTransaction(obj *seer_common.TransactionJson) *MantleSepoliaTransaction {
	var accessList []*MantleSepoliaTransactionAccessList
	for _, al := range obj.AccessList {
		accessList = append(accessList, &MantleSepoliaTransactionAccessList{
			Address:     al.Address,
			StorageKeys: al.StorageKeys,
		})
	}

	return &MantleSepoliaTransaction{
		Hash:                 obj.Hash,
		BlockNumber:          fromHex(obj.BlockNumber).Uint64(),
		BlockHash:            obj.BlockHash,
		FromAddress:          obj.FromAddress,
		ToAddress:            obj.ToAddress,
		Gas:                  obj.Gas,
		GasPrice:             obj.GasPrice,
		MaxFeePerGas:         obj.MaxFeePerGas,
		MaxPriorityFeePerGas: obj.MaxPriorityFeePerGas,
		Input:                obj.Input,
		Nonce:                obj.Nonce,
		TransactionIndex:     fromHex(obj.TransactionIndex).Uint64(),
		TransactionType:      fromHex(obj.TransactionType).Uint64(),
		Value:                obj.Value,
		IndexedAt:            fromHex(obj.IndexedAt).Uint64(),
		BlockTimestamp:       fromHex(obj.BlockTimestamp).Uint64(),

		ChainId: obj.ChainId,
		V:       obj.V,
		R:       obj.R,
		S:       obj.S,

		AccessList: accessList,
		YParity:    obj.YParity,
	}
}

func ToProtoSingleEventLog(obj *seer_common.EventJson) *MantleSepoliaEventLog {

	return &MantleSepoliaEventLog{
		Address:         obj.Address,
		Topics:          obj.Topics,
		Data:            obj.Data,
		BlockNumber:     fromHex(obj.BlockNumber).Uint64(),
		TransactionHash: obj.TransactionHash,
		LogIndex:        fromHex(obj.LogIndex).Uint64(),
		BlockHash:       obj.BlockHash,
		Removed:         obj.Removed,
	}
}

func (c *Client) DecodeProtoEventLogs(data []string) ([]*MantleSepoliaEventLog, error) {
	var events []*MantleSepoliaEventLog
	for _, d := range data {
		var event MantleSepoliaEventLog
		base64Decoded, err := base64.StdEncoding.DecodeString(d)
		if err != nil {
			return nil, err
		}
		if err := proto.Unmarshal(base64Decoded, &event); err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	return events, nil
}

func (c *Client) DecodeProtoTransactions(data []string) ([]*MantleSepoliaTransaction, error) {
	var transactions []*MantleSepoliaTransaction
	for _, d := range data {
		var transaction MantleSepoliaTransaction
		base64Decoded, err := base64.StdEncoding.DecodeString(d)
		if err != nil {
			return nil, err
		}
		if err := proto.Unmarshal(base64Decoded, &transaction); err != nil {
			return nil, err
		}
		transactions = append(transactions, &transaction)
	}
	return transactions, nil
}

func (c *Client) DecodeProtoBlocks(data []string) ([]*MantleSepoliaBlock, error) {
	var blocks []*MantleSepoliaBlock
	for _, d := range data {
		var block MantleSepoliaBlock
		base64Decoded, err := base64.StdEncoding.DecodeString(d)
		if err != nil {
			return nil, err
		}
		if err := proto.Unmarshal(base64Decoded, &block); err != nil {
			return nil, err
		}
		blocks = append(blocks, &block)
	}
	return blocks, nil
}

func (c *Client) DecodeProtoEventsToLabels(events []string, blocksCache map[uint64]uint64, abiMap map[string]map[string]map[string]string) ([]indexer.EventLabel, error) {

	decodedEvents, err := c.DecodeProtoEventLogs(events)

	if err != nil {
		return nil, err
	}

	var labels []indexer.EventLabel

	for _, event := range decodedEvents {

		var topicSelector string

		if len(event.Topics) > 0 {
			topicSelector = event.Topics[0][:10]
		} else {
			continue
		}

		checksumAddress := common.HexToAddress(event.Address).Hex()

		// Get the ABI string
		contractAbi, err := abi.JSON(strings.NewReader(abiMap[checksumAddress][topicSelector]["abi"]))
		if err != nil {
			fmt.Println("Error initializing contract ABI: ", err)
			return nil, err
		}

		// Decode the event data
		decodedArgs, err := seer_common.DecodeLogArgsToLabelData(&contractAbi, event.Topics, event.Data)

		if err != nil {
			fmt.Println("Error decoding event data: ", err)
			return nil, err
		}

		// Convert decodedArgs map to JSON
		labelDataBytes, err := json.Marshal(decodedArgs)
		if err != nil {
			return nil, err
		}

		// Convert JSON byte slice to string
		labelDataString := string(labelDataBytes)

		// Convert event to label
		eventLabel := indexer.EventLabel{
			Label:           indexer.SeerCrawlerLabel,
			LabelName:       abiMap[checksumAddress][topicSelector]["abi_name"],
			LabelType:       "event",
			BlockNumber:     event.BlockNumber,
			BlockHash:       event.BlockHash,
			Address:         event.Address,
			TransactionHash: event.TransactionHash,
			LabelData:       labelDataString,
			BlockTimestamp:  blocksCache[event.BlockNumber],
			LogIndex:        event.LogIndex,
		}

		labels = append(labels, eventLabel)

	}

	return labels, nil
}

func (c *Client) DecodeProtoTransactionsToLabels(transactions []string, blocksCache map[uint64]uint64, abiMap map[string]map[string]map[string]string) ([]indexer.TransactionLabel, error) {

	decodedTransactions, err := c.DecodeProtoTransactions(transactions)

	if err != nil {
		return nil, err
	}

	var labels []indexer.TransactionLabel

	for _, transaction := range decodedTransactions {

		selector := transaction.Input[:10]

		// To checksum address
		checksumAddress := common.HexToAddress(transaction.ToAddress).Hex()

		contractAbi, err := abi.JSON(strings.NewReader(abiMap[checksumAddress][selector]["abi"]))

		if err != nil {
			return nil, err
		}

		inputData, err := hex.DecodeString(transaction.Input[2:])
		if err != nil {
			fmt.Println("Error decoding input data: ", err)
			return nil, err
		}

		decodedArgs, err := seer_common.DecodeTransactionInputDataToInterface(&contractAbi, inputData)

		if err != nil {
			fmt.Println("Error decoding transaction input data: ", err)
			return nil, err
		}

		labelDataBytes, err := json.Marshal(decodedArgs)
		if err != nil {
			return nil, err
		}

		// Convert JSON byte slice to string
		labelDataString := string(labelDataBytes)

		// Convert transaction to label
		transactionLabel := indexer.TransactionLabel{
			Address:         transaction.ToAddress,
			BlockNumber:     transaction.BlockNumber,
			BlockHash:       transaction.BlockHash,
			CallerAddress:   transaction.FromAddress,
			LabelName:       abiMap[checksumAddress][selector]["abi_name"],
			LabelType:       "tx_call",
			OriginAddress:   transaction.FromAddress,
			Label:           indexer.SeerCrawlerLabel,
			TransactionHash: transaction.Hash,
			LabelData:       labelDataString,
			BlockTimestamp:  blocksCache[transaction.BlockNumber],
		}

		labels = append(labels, transactionLabel)

	}

	return labels, nil
}
