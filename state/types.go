package state

import (
	"time"

	"github.com/0xPolygonHermez/zkevm-node/state/runtime/instrumentation"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ProcessBatchResponse represents the response of a batch process.
type ProcessBatchResponse struct {
	NewStateRoot        common.Hash
	NewAccInputHash     common.Hash
	NewLocalExitRoot    common.Hash
	NewBatchNumber      uint64
	CntKeccakHashes     uint32
	CntPoseidonHashes   uint32
	CntPoseidonPaddings uint32
	CntMemAligns        uint32
	CntArithmetics      uint32
	CntBinaries         uint32
	CntSteps            uint32
	CumulativeGasUsed   uint64
	Responses           []*ProcessTransactionResponse
	Error               error
	IsBatchProcessed    bool
}

// ProcessTransactionResponse represents the response of a tx process.
type ProcessTransactionResponse struct {
	// TxHash is the hash of the transaction
	TxHash common.Hash
	// Type indicates legacy transaction
	// It will be always 0 (legacy) in the executor
	Type uint32
	// ReturnValue is the returned data from the runtime (function result or data supplied with revert opcode)
	ReturnValue []byte
	// GasLeft is the total gas left as result of execution
	GasLeft uint64
	// GasUsed is the total gas used as result of execution or gas estimation
	GasUsed uint64
	// GasRefunded is the total gas refunded as result of execution
	GasRefunded uint64
	// Error represents any error encountered during the execution
	Error error
	// CreateAddress is the new SC Address in case of SC creation
	CreateAddress common.Address
	// StateRoot is the State Root
	StateRoot common.Hash
	// Logs emitted by LOG opcode
	Logs []*types.Log
	// IsProcessed indicates if this tx didn't fit into the batch
	IsProcessed bool
	// Tx is the whole transaction object
	Tx types.Transaction
	// ExecutionTrace contains the traces produced in the execution
	ExecutionTrace []instrumentation.StructLog
	// CallTrace contains the call trace.
	CallTrace instrumentation.ExecutorTrace
}

const (
	// DebugInfoErrorType_ROM_OOC indicates a not handled OOC error by the ROM
	DebugInfoErrorType_ROM_OOC = "ROM OOC"
)

// DebugInfo allows handling runtime debug info
type DebugInfo struct {
	ErrorType string
	Timestamp time.Time
	Payload   string
}
