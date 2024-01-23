// Code generated by fastssz. DO NOT EDIT.
// Hash: 0861ebfee0a62ce9e10d8b677c0f8aa1f12c99ee69d3a864018818270a981c57
package enginev1

import (
	ssz "github.com/prysmaticlabs/fastssz"
	github_com_prysmaticlabs_prysm_v3_consensus_types_primitives "github.com/prysmaticlabs/prysm/v3/consensus-types/primitives"
)

// MarshalSSZ ssz marshals the ExecutionPayload object
func (e *ExecutionPayload) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(e)
}

// MarshalSSZTo ssz marshals the ExecutionPayload object to a target array
func (e *ExecutionPayload) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(508)

	// Field (0) 'ParentHash'
	if size := len(e.ParentHash); size != 32 {
		err = ssz.ErrBytesLengthFn("--.ParentHash", size, 32)
		return
	}
	dst = append(dst, e.ParentHash...)

	// Field (1) 'FeeRecipient'
	if size := len(e.FeeRecipient); size != 20 {
		err = ssz.ErrBytesLengthFn("--.FeeRecipient", size, 20)
		return
	}
	dst = append(dst, e.FeeRecipient...)

	// Field (2) 'StateRoot'
	if size := len(e.StateRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.StateRoot", size, 32)
		return
	}
	dst = append(dst, e.StateRoot...)

	// Field (3) 'ReceiptsRoot'
	if size := len(e.ReceiptsRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.ReceiptsRoot", size, 32)
		return
	}
	dst = append(dst, e.ReceiptsRoot...)

	// Field (4) 'LogsBloom'
	if size := len(e.LogsBloom); size != 256 {
		err = ssz.ErrBytesLengthFn("--.LogsBloom", size, 256)
		return
	}
	dst = append(dst, e.LogsBloom...)

	// Field (5) 'PrevRandao'
	if size := len(e.PrevRandao); size != 32 {
		err = ssz.ErrBytesLengthFn("--.PrevRandao", size, 32)
		return
	}
	dst = append(dst, e.PrevRandao...)

	// Field (6) 'BlockNumber'
	dst = ssz.MarshalUint64(dst, e.BlockNumber)

	// Field (7) 'GasLimit'
	dst = ssz.MarshalUint64(dst, e.GasLimit)

	// Field (8) 'GasUsed'
	dst = ssz.MarshalUint64(dst, e.GasUsed)

	// Field (9) 'Timestamp'
	dst = ssz.MarshalUint64(dst, e.Timestamp)

	// Offset (10) 'ExtraData'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(e.ExtraData)

	// Field (11) 'BaseFeePerGas'
	if size := len(e.BaseFeePerGas); size != 32 {
		err = ssz.ErrBytesLengthFn("--.BaseFeePerGas", size, 32)
		return
	}
	dst = append(dst, e.BaseFeePerGas...)

	// Field (12) 'BlockHash'
	if size := len(e.BlockHash); size != 32 {
		err = ssz.ErrBytesLengthFn("--.BlockHash", size, 32)
		return
	}
	dst = append(dst, e.BlockHash...)

	// Offset (13) 'Transactions'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(e.Transactions); ii++ {
		offset += 4
		offset += len(e.Transactions[ii])
	}

	// Field (10) 'ExtraData'
	if size := len(e.ExtraData); size > 32 {
		err = ssz.ErrBytesLengthFn("--.ExtraData", size, 32)
		return
	}
	dst = append(dst, e.ExtraData...)

	// Field (13) 'Transactions'
	if size := len(e.Transactions); size > 1048576 {
		err = ssz.ErrListTooBigFn("--.Transactions", size, 1048576)
		return
	}
	{
		offset = 4 * len(e.Transactions)
		for ii := 0; ii < len(e.Transactions); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += len(e.Transactions[ii])
		}
	}
	for ii := 0; ii < len(e.Transactions); ii++ {
		if size := len(e.Transactions[ii]); size > 1073741824 {
			err = ssz.ErrBytesLengthFn("--.Transactions[ii]", size, 1073741824)
			return
		}
		dst = append(dst, e.Transactions[ii]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the ExecutionPayload object
func (e *ExecutionPayload) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 508 {
		return ssz.ErrSize
	}

	tail := buf
	var o10, o13 uint64

	// Field (0) 'ParentHash'
	if cap(e.ParentHash) == 0 {
		e.ParentHash = make([]byte, 0, len(buf[0:32]))
	}
	e.ParentHash = append(e.ParentHash, buf[0:32]...)

	// Field (1) 'FeeRecipient'
	if cap(e.FeeRecipient) == 0 {
		e.FeeRecipient = make([]byte, 0, len(buf[32:52]))
	}
	e.FeeRecipient = append(e.FeeRecipient, buf[32:52]...)

	// Field (2) 'StateRoot'
	if cap(e.StateRoot) == 0 {
		e.StateRoot = make([]byte, 0, len(buf[52:84]))
	}
	e.StateRoot = append(e.StateRoot, buf[52:84]...)

	// Field (3) 'ReceiptsRoot'
	if cap(e.ReceiptsRoot) == 0 {
		e.ReceiptsRoot = make([]byte, 0, len(buf[84:116]))
	}
	e.ReceiptsRoot = append(e.ReceiptsRoot, buf[84:116]...)

	// Field (4) 'LogsBloom'
	if cap(e.LogsBloom) == 0 {
		e.LogsBloom = make([]byte, 0, len(buf[116:372]))
	}
	e.LogsBloom = append(e.LogsBloom, buf[116:372]...)

	// Field (5) 'PrevRandao'
	if cap(e.PrevRandao) == 0 {
		e.PrevRandao = make([]byte, 0, len(buf[372:404]))
	}
	e.PrevRandao = append(e.PrevRandao, buf[372:404]...)

	// Field (6) 'BlockNumber'
	e.BlockNumber = ssz.UnmarshallUint64(buf[404:412])

	// Field (7) 'GasLimit'
	e.GasLimit = ssz.UnmarshallUint64(buf[412:420])

	// Field (8) 'GasUsed'
	e.GasUsed = ssz.UnmarshallUint64(buf[420:428])

	// Field (9) 'Timestamp'
	e.Timestamp = ssz.UnmarshallUint64(buf[428:436])

	// Offset (10) 'ExtraData'
	if o10 = ssz.ReadOffset(buf[436:440]); o10 > size {
		return ssz.ErrOffset
	}

	if o10 < 508 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (11) 'BaseFeePerGas'
	if cap(e.BaseFeePerGas) == 0 {
		e.BaseFeePerGas = make([]byte, 0, len(buf[440:472]))
	}
	e.BaseFeePerGas = append(e.BaseFeePerGas, buf[440:472]...)

	// Field (12) 'BlockHash'
	if cap(e.BlockHash) == 0 {
		e.BlockHash = make([]byte, 0, len(buf[472:504]))
	}
	e.BlockHash = append(e.BlockHash, buf[472:504]...)

	// Offset (13) 'Transactions'
	if o13 = ssz.ReadOffset(buf[504:508]); o13 > size || o10 > o13 {
		return ssz.ErrOffset
	}

	// Field (10) 'ExtraData'
	{
		buf = tail[o10:o13]
		if len(buf) > 32 {
			return ssz.ErrBytesLength
		}
		if cap(e.ExtraData) == 0 {
			e.ExtraData = make([]byte, 0, len(buf))
		}
		e.ExtraData = append(e.ExtraData, buf...)
	}

	// Field (13) 'Transactions'
	{
		buf = tail[o13:]
		num, err := ssz.DecodeDynamicLength(buf, 1048576)
		if err != nil {
			return err
		}
		e.Transactions = make([][]byte, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if len(buf) > 1073741824 {
				return ssz.ErrBytesLength
			}
			if cap(e.Transactions[indx]) == 0 {
				e.Transactions[indx] = make([]byte, 0, len(buf))
			}
			e.Transactions[indx] = append(e.Transactions[indx], buf...)
			return nil
		})
		if err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ExecutionPayload object
func (e *ExecutionPayload) SizeSSZ() (size int) {
	size = 508

	// Field (10) 'ExtraData'
	size += len(e.ExtraData)

	// Field (13) 'Transactions'
	for ii := 0; ii < len(e.Transactions); ii++ {
		size += 4
		size += len(e.Transactions[ii])
	}

	return
}

// HashTreeRoot ssz hashes the ExecutionPayload object
func (e *ExecutionPayload) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(e)
}

// HashTreeRootWith ssz hashes the ExecutionPayload object with a hasher
func (e *ExecutionPayload) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'ParentHash'
	if size := len(e.ParentHash); size != 32 {
		err = ssz.ErrBytesLengthFn("--.ParentHash", size, 32)
		return
	}
	hh.PutBytes(e.ParentHash)

	// Field (1) 'FeeRecipient'
	if size := len(e.FeeRecipient); size != 20 {
		err = ssz.ErrBytesLengthFn("--.FeeRecipient", size, 20)
		return
	}
	hh.PutBytes(e.FeeRecipient)

	// Field (2) 'StateRoot'
	if size := len(e.StateRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.StateRoot", size, 32)
		return
	}
	hh.PutBytes(e.StateRoot)

	// Field (3) 'ReceiptsRoot'
	if size := len(e.ReceiptsRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.ReceiptsRoot", size, 32)
		return
	}
	hh.PutBytes(e.ReceiptsRoot)

	// Field (4) 'LogsBloom'
	if size := len(e.LogsBloom); size != 256 {
		err = ssz.ErrBytesLengthFn("--.LogsBloom", size, 256)
		return
	}
	hh.PutBytes(e.LogsBloom)

	// Field (5) 'PrevRandao'
	if size := len(e.PrevRandao); size != 32 {
		err = ssz.ErrBytesLengthFn("--.PrevRandao", size, 32)
		return
	}
	hh.PutBytes(e.PrevRandao)

	// Field (6) 'BlockNumber'
	hh.PutUint64(e.BlockNumber)

	// Field (7) 'GasLimit'
	hh.PutUint64(e.GasLimit)

	// Field (8) 'GasUsed'
	hh.PutUint64(e.GasUsed)

	// Field (9) 'Timestamp'
	hh.PutUint64(e.Timestamp)

	// Field (10) 'ExtraData'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(e.ExtraData))
		if byteLen > 32 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.PutBytes(e.ExtraData)
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(elemIndx, byteLen, (32+31)/32)
		} else {
			hh.MerkleizeWithMixin(elemIndx, byteLen, (32+31)/32)
		}
	}

	// Field (11) 'BaseFeePerGas'
	if size := len(e.BaseFeePerGas); size != 32 {
		err = ssz.ErrBytesLengthFn("--.BaseFeePerGas", size, 32)
		return
	}
	hh.PutBytes(e.BaseFeePerGas)

	// Field (12) 'BlockHash'
	if size := len(e.BlockHash); size != 32 {
		err = ssz.ErrBytesLengthFn("--.BlockHash", size, 32)
		return
	}
	hh.PutBytes(e.BlockHash)

	// Field (13) 'Transactions'
	{
		subIndx := hh.Index()
		num := uint64(len(e.Transactions))
		if num > 1048576 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range e.Transactions {
			{
				elemIndx := hh.Index()
				byteLen := uint64(len(elem))
				if byteLen > 1073741824 {
					err = ssz.ErrIncorrectListSize
					return
				}
				hh.AppendBytes32(elem)
				if ssz.EnableVectorizedHTR {
					hh.MerkleizeWithMixinVectorizedHTR(elemIndx, byteLen, (1073741824+31)/32)
				} else {
					hh.MerkleizeWithMixin(elemIndx, byteLen, (1073741824+31)/32)
				}
			}
		}
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(subIndx, num, 1048576)
		} else {
			hh.MerkleizeWithMixin(subIndx, num, 1048576)
		}
	}

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the ExecutionPayloadCapella object
func (e *ExecutionPayloadCapella) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(e)
}

// MarshalSSZTo ssz marshals the ExecutionPayloadCapella object to a target array
func (e *ExecutionPayloadCapella) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(512)

	// Field (0) 'ParentHash'
	if size := len(e.ParentHash); size != 32 {
		err = ssz.ErrBytesLengthFn("--.ParentHash", size, 32)
		return
	}
	dst = append(dst, e.ParentHash...)

	// Field (1) 'FeeRecipient'
	if size := len(e.FeeRecipient); size != 20 {
		err = ssz.ErrBytesLengthFn("--.FeeRecipient", size, 20)
		return
	}
	dst = append(dst, e.FeeRecipient...)

	// Field (2) 'StateRoot'
	if size := len(e.StateRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.StateRoot", size, 32)
		return
	}
	dst = append(dst, e.StateRoot...)

	// Field (3) 'ReceiptsRoot'
	if size := len(e.ReceiptsRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.ReceiptsRoot", size, 32)
		return
	}
	dst = append(dst, e.ReceiptsRoot...)

	// Field (4) 'LogsBloom'
	if size := len(e.LogsBloom); size != 256 {
		err = ssz.ErrBytesLengthFn("--.LogsBloom", size, 256)
		return
	}
	dst = append(dst, e.LogsBloom...)

	// Field (5) 'PrevRandao'
	if size := len(e.PrevRandao); size != 32 {
		err = ssz.ErrBytesLengthFn("--.PrevRandao", size, 32)
		return
	}
	dst = append(dst, e.PrevRandao...)

	// Field (6) 'BlockNumber'
	dst = ssz.MarshalUint64(dst, e.BlockNumber)

	// Field (7) 'GasLimit'
	dst = ssz.MarshalUint64(dst, e.GasLimit)

	// Field (8) 'GasUsed'
	dst = ssz.MarshalUint64(dst, e.GasUsed)

	// Field (9) 'Timestamp'
	dst = ssz.MarshalUint64(dst, e.Timestamp)

	// Offset (10) 'ExtraData'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(e.ExtraData)

	// Field (11) 'BaseFeePerGas'
	if size := len(e.BaseFeePerGas); size != 32 {
		err = ssz.ErrBytesLengthFn("--.BaseFeePerGas", size, 32)
		return
	}
	dst = append(dst, e.BaseFeePerGas...)

	// Field (12) 'BlockHash'
	if size := len(e.BlockHash); size != 32 {
		err = ssz.ErrBytesLengthFn("--.BlockHash", size, 32)
		return
	}
	dst = append(dst, e.BlockHash...)

	// Offset (13) 'Transactions'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(e.Transactions); ii++ {
		offset += 4
		offset += len(e.Transactions[ii])
	}

	// Offset (14) 'Withdrawals'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(e.Withdrawals) * 44

	// Field (10) 'ExtraData'
	if size := len(e.ExtraData); size > 32 {
		err = ssz.ErrBytesLengthFn("--.ExtraData", size, 32)
		return
	}
	dst = append(dst, e.ExtraData...)

	// Field (13) 'Transactions'
	if size := len(e.Transactions); size > 1048576 {
		err = ssz.ErrListTooBigFn("--.Transactions", size, 1048576)
		return
	}
	{
		offset = 4 * len(e.Transactions)
		for ii := 0; ii < len(e.Transactions); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += len(e.Transactions[ii])
		}
	}
	for ii := 0; ii < len(e.Transactions); ii++ {
		if size := len(e.Transactions[ii]); size > 1073741824 {
			err = ssz.ErrBytesLengthFn("--.Transactions[ii]", size, 1073741824)
			return
		}
		dst = append(dst, e.Transactions[ii]...)
	}

	// Field (14) 'Withdrawals'
	if size := len(e.Withdrawals); size > 16 {
		err = ssz.ErrListTooBigFn("--.Withdrawals", size, 16)
		return
	}
	for ii := 0; ii < len(e.Withdrawals); ii++ {
		if dst, err = e.Withdrawals[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the ExecutionPayloadCapella object
func (e *ExecutionPayloadCapella) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 512 {
		return ssz.ErrSize
	}

	tail := buf
	var o10, o13, o14 uint64

	// Field (0) 'ParentHash'
	if cap(e.ParentHash) == 0 {
		e.ParentHash = make([]byte, 0, len(buf[0:32]))
	}
	e.ParentHash = append(e.ParentHash, buf[0:32]...)

	// Field (1) 'FeeRecipient'
	if cap(e.FeeRecipient) == 0 {
		e.FeeRecipient = make([]byte, 0, len(buf[32:52]))
	}
	e.FeeRecipient = append(e.FeeRecipient, buf[32:52]...)

	// Field (2) 'StateRoot'
	if cap(e.StateRoot) == 0 {
		e.StateRoot = make([]byte, 0, len(buf[52:84]))
	}
	e.StateRoot = append(e.StateRoot, buf[52:84]...)

	// Field (3) 'ReceiptsRoot'
	if cap(e.ReceiptsRoot) == 0 {
		e.ReceiptsRoot = make([]byte, 0, len(buf[84:116]))
	}
	e.ReceiptsRoot = append(e.ReceiptsRoot, buf[84:116]...)

	// Field (4) 'LogsBloom'
	if cap(e.LogsBloom) == 0 {
		e.LogsBloom = make([]byte, 0, len(buf[116:372]))
	}
	e.LogsBloom = append(e.LogsBloom, buf[116:372]...)

	// Field (5) 'PrevRandao'
	if cap(e.PrevRandao) == 0 {
		e.PrevRandao = make([]byte, 0, len(buf[372:404]))
	}
	e.PrevRandao = append(e.PrevRandao, buf[372:404]...)

	// Field (6) 'BlockNumber'
	e.BlockNumber = ssz.UnmarshallUint64(buf[404:412])

	// Field (7) 'GasLimit'
	e.GasLimit = ssz.UnmarshallUint64(buf[412:420])

	// Field (8) 'GasUsed'
	e.GasUsed = ssz.UnmarshallUint64(buf[420:428])

	// Field (9) 'Timestamp'
	e.Timestamp = ssz.UnmarshallUint64(buf[428:436])

	// Offset (10) 'ExtraData'
	if o10 = ssz.ReadOffset(buf[436:440]); o10 > size {
		return ssz.ErrOffset
	}

	if o10 < 512 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (11) 'BaseFeePerGas'
	if cap(e.BaseFeePerGas) == 0 {
		e.BaseFeePerGas = make([]byte, 0, len(buf[440:472]))
	}
	e.BaseFeePerGas = append(e.BaseFeePerGas, buf[440:472]...)

	// Field (12) 'BlockHash'
	if cap(e.BlockHash) == 0 {
		e.BlockHash = make([]byte, 0, len(buf[472:504]))
	}
	e.BlockHash = append(e.BlockHash, buf[472:504]...)

	// Offset (13) 'Transactions'
	if o13 = ssz.ReadOffset(buf[504:508]); o13 > size || o10 > o13 {
		return ssz.ErrOffset
	}

	// Offset (14) 'Withdrawals'
	if o14 = ssz.ReadOffset(buf[508:512]); o14 > size || o13 > o14 {
		return ssz.ErrOffset
	}

	// Field (10) 'ExtraData'
	{
		buf = tail[o10:o13]
		if len(buf) > 32 {
			return ssz.ErrBytesLength
		}
		if cap(e.ExtraData) == 0 {
			e.ExtraData = make([]byte, 0, len(buf))
		}
		e.ExtraData = append(e.ExtraData, buf...)
	}

	// Field (13) 'Transactions'
	{
		buf = tail[o13:o14]
		num, err := ssz.DecodeDynamicLength(buf, 1048576)
		if err != nil {
			return err
		}
		e.Transactions = make([][]byte, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if len(buf) > 1073741824 {
				return ssz.ErrBytesLength
			}
			if cap(e.Transactions[indx]) == 0 {
				e.Transactions[indx] = make([]byte, 0, len(buf))
			}
			e.Transactions[indx] = append(e.Transactions[indx], buf...)
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (14) 'Withdrawals'
	{
		buf = tail[o14:]
		num, err := ssz.DivideInt2(len(buf), 44, 16)
		if err != nil {
			return err
		}
		e.Withdrawals = make([]*Withdrawal, num)
		for ii := 0; ii < num; ii++ {
			if e.Withdrawals[ii] == nil {
				e.Withdrawals[ii] = new(Withdrawal)
			}
			if err = e.Withdrawals[ii].UnmarshalSSZ(buf[ii*44 : (ii+1)*44]); err != nil {
				return err
			}
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ExecutionPayloadCapella object
func (e *ExecutionPayloadCapella) SizeSSZ() (size int) {
	size = 512

	// Field (10) 'ExtraData'
	size += len(e.ExtraData)

	// Field (13) 'Transactions'
	for ii := 0; ii < len(e.Transactions); ii++ {
		size += 4
		size += len(e.Transactions[ii])
	}

	// Field (14) 'Withdrawals'
	size += len(e.Withdrawals) * 44

	return
}

// HashTreeRoot ssz hashes the ExecutionPayloadCapella object
func (e *ExecutionPayloadCapella) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(e)
}

// HashTreeRootWith ssz hashes the ExecutionPayloadCapella object with a hasher
func (e *ExecutionPayloadCapella) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'ParentHash'
	if size := len(e.ParentHash); size != 32 {
		err = ssz.ErrBytesLengthFn("--.ParentHash", size, 32)
		return
	}
	hh.PutBytes(e.ParentHash)

	// Field (1) 'FeeRecipient'
	if size := len(e.FeeRecipient); size != 20 {
		err = ssz.ErrBytesLengthFn("--.FeeRecipient", size, 20)
		return
	}
	hh.PutBytes(e.FeeRecipient)

	// Field (2) 'StateRoot'
	if size := len(e.StateRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.StateRoot", size, 32)
		return
	}
	hh.PutBytes(e.StateRoot)

	// Field (3) 'ReceiptsRoot'
	if size := len(e.ReceiptsRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.ReceiptsRoot", size, 32)
		return
	}
	hh.PutBytes(e.ReceiptsRoot)

	// Field (4) 'LogsBloom'
	if size := len(e.LogsBloom); size != 256 {
		err = ssz.ErrBytesLengthFn("--.LogsBloom", size, 256)
		return
	}
	hh.PutBytes(e.LogsBloom)

	// Field (5) 'PrevRandao'
	if size := len(e.PrevRandao); size != 32 {
		err = ssz.ErrBytesLengthFn("--.PrevRandao", size, 32)
		return
	}
	hh.PutBytes(e.PrevRandao)

	// Field (6) 'BlockNumber'
	hh.PutUint64(e.BlockNumber)

	// Field (7) 'GasLimit'
	hh.PutUint64(e.GasLimit)

	// Field (8) 'GasUsed'
	hh.PutUint64(e.GasUsed)

	// Field (9) 'Timestamp'
	hh.PutUint64(e.Timestamp)

	// Field (10) 'ExtraData'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(e.ExtraData))
		if byteLen > 32 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.PutBytes(e.ExtraData)
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(elemIndx, byteLen, (32+31)/32)
		} else {
			hh.MerkleizeWithMixin(elemIndx, byteLen, (32+31)/32)
		}
	}

	// Field (11) 'BaseFeePerGas'
	if size := len(e.BaseFeePerGas); size != 32 {
		err = ssz.ErrBytesLengthFn("--.BaseFeePerGas", size, 32)
		return
	}
	hh.PutBytes(e.BaseFeePerGas)

	// Field (12) 'BlockHash'
	if size := len(e.BlockHash); size != 32 {
		err = ssz.ErrBytesLengthFn("--.BlockHash", size, 32)
		return
	}
	hh.PutBytes(e.BlockHash)

	// Field (13) 'Transactions'
	{
		subIndx := hh.Index()
		num := uint64(len(e.Transactions))
		if num > 1048576 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range e.Transactions {
			{
				elemIndx := hh.Index()
				byteLen := uint64(len(elem))
				if byteLen > 1073741824 {
					err = ssz.ErrIncorrectListSize
					return
				}
				hh.AppendBytes32(elem)
				if ssz.EnableVectorizedHTR {
					hh.MerkleizeWithMixinVectorizedHTR(elemIndx, byteLen, (1073741824+31)/32)
				} else {
					hh.MerkleizeWithMixin(elemIndx, byteLen, (1073741824+31)/32)
				}
			}
		}
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(subIndx, num, 1048576)
		} else {
			hh.MerkleizeWithMixin(subIndx, num, 1048576)
		}
	}

	// Field (14) 'Withdrawals'
	{
		subIndx := hh.Index()
		num := uint64(len(e.Withdrawals))
		if num > 16 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range e.Withdrawals {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(subIndx, num, 16)
		} else {
			hh.MerkleizeWithMixin(subIndx, num, 16)
		}
	}

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the ExecutionPayloadHeader object
func (e *ExecutionPayloadHeader) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(e)
}

// MarshalSSZTo ssz marshals the ExecutionPayloadHeader object to a target array
func (e *ExecutionPayloadHeader) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(536)

	// Field (0) 'ParentHash'
	if size := len(e.ParentHash); size != 32 {
		err = ssz.ErrBytesLengthFn("--.ParentHash", size, 32)
		return
	}
	dst = append(dst, e.ParentHash...)

	// Field (1) 'FeeRecipient'
	if size := len(e.FeeRecipient); size != 20 {
		err = ssz.ErrBytesLengthFn("--.FeeRecipient", size, 20)
		return
	}
	dst = append(dst, e.FeeRecipient...)

	// Field (2) 'StateRoot'
	if size := len(e.StateRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.StateRoot", size, 32)
		return
	}
	dst = append(dst, e.StateRoot...)

	// Field (3) 'ReceiptsRoot'
	if size := len(e.ReceiptsRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.ReceiptsRoot", size, 32)
		return
	}
	dst = append(dst, e.ReceiptsRoot...)

	// Field (4) 'LogsBloom'
	if size := len(e.LogsBloom); size != 256 {
		err = ssz.ErrBytesLengthFn("--.LogsBloom", size, 256)
		return
	}
	dst = append(dst, e.LogsBloom...)

	// Field (5) 'PrevRandao'
	if size := len(e.PrevRandao); size != 32 {
		err = ssz.ErrBytesLengthFn("--.PrevRandao", size, 32)
		return
	}
	dst = append(dst, e.PrevRandao...)

	// Field (6) 'BlockNumber'
	dst = ssz.MarshalUint64(dst, e.BlockNumber)

	// Field (7) 'GasLimit'
	dst = ssz.MarshalUint64(dst, e.GasLimit)

	// Field (8) 'GasUsed'
	dst = ssz.MarshalUint64(dst, e.GasUsed)

	// Field (9) 'Timestamp'
	dst = ssz.MarshalUint64(dst, e.Timestamp)

	// Offset (10) 'ExtraData'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(e.ExtraData)

	// Field (11) 'BaseFeePerGas'
	if size := len(e.BaseFeePerGas); size != 32 {
		err = ssz.ErrBytesLengthFn("--.BaseFeePerGas", size, 32)
		return
	}
	dst = append(dst, e.BaseFeePerGas...)

	// Field (12) 'BlockHash'
	if size := len(e.BlockHash); size != 32 {
		err = ssz.ErrBytesLengthFn("--.BlockHash", size, 32)
		return
	}
	dst = append(dst, e.BlockHash...)

	// Field (13) 'TransactionsRoot'
	if size := len(e.TransactionsRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.TransactionsRoot", size, 32)
		return
	}
	dst = append(dst, e.TransactionsRoot...)

	// Field (10) 'ExtraData'
	if size := len(e.ExtraData); size > 32 {
		err = ssz.ErrBytesLengthFn("--.ExtraData", size, 32)
		return
	}
	dst = append(dst, e.ExtraData...)

	return
}

// UnmarshalSSZ ssz unmarshals the ExecutionPayloadHeader object
func (e *ExecutionPayloadHeader) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 536 {
		return ssz.ErrSize
	}

	tail := buf
	var o10 uint64

	// Field (0) 'ParentHash'
	if cap(e.ParentHash) == 0 {
		e.ParentHash = make([]byte, 0, len(buf[0:32]))
	}
	e.ParentHash = append(e.ParentHash, buf[0:32]...)

	// Field (1) 'FeeRecipient'
	if cap(e.FeeRecipient) == 0 {
		e.FeeRecipient = make([]byte, 0, len(buf[32:52]))
	}
	e.FeeRecipient = append(e.FeeRecipient, buf[32:52]...)

	// Field (2) 'StateRoot'
	if cap(e.StateRoot) == 0 {
		e.StateRoot = make([]byte, 0, len(buf[52:84]))
	}
	e.StateRoot = append(e.StateRoot, buf[52:84]...)

	// Field (3) 'ReceiptsRoot'
	if cap(e.ReceiptsRoot) == 0 {
		e.ReceiptsRoot = make([]byte, 0, len(buf[84:116]))
	}
	e.ReceiptsRoot = append(e.ReceiptsRoot, buf[84:116]...)

	// Field (4) 'LogsBloom'
	if cap(e.LogsBloom) == 0 {
		e.LogsBloom = make([]byte, 0, len(buf[116:372]))
	}
	e.LogsBloom = append(e.LogsBloom, buf[116:372]...)

	// Field (5) 'PrevRandao'
	if cap(e.PrevRandao) == 0 {
		e.PrevRandao = make([]byte, 0, len(buf[372:404]))
	}
	e.PrevRandao = append(e.PrevRandao, buf[372:404]...)

	// Field (6) 'BlockNumber'
	e.BlockNumber = ssz.UnmarshallUint64(buf[404:412])

	// Field (7) 'GasLimit'
	e.GasLimit = ssz.UnmarshallUint64(buf[412:420])

	// Field (8) 'GasUsed'
	e.GasUsed = ssz.UnmarshallUint64(buf[420:428])

	// Field (9) 'Timestamp'
	e.Timestamp = ssz.UnmarshallUint64(buf[428:436])

	// Offset (10) 'ExtraData'
	if o10 = ssz.ReadOffset(buf[436:440]); o10 > size {
		return ssz.ErrOffset
	}

	if o10 < 536 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (11) 'BaseFeePerGas'
	if cap(e.BaseFeePerGas) == 0 {
		e.BaseFeePerGas = make([]byte, 0, len(buf[440:472]))
	}
	e.BaseFeePerGas = append(e.BaseFeePerGas, buf[440:472]...)

	// Field (12) 'BlockHash'
	if cap(e.BlockHash) == 0 {
		e.BlockHash = make([]byte, 0, len(buf[472:504]))
	}
	e.BlockHash = append(e.BlockHash, buf[472:504]...)

	// Field (13) 'TransactionsRoot'
	if cap(e.TransactionsRoot) == 0 {
		e.TransactionsRoot = make([]byte, 0, len(buf[504:536]))
	}
	e.TransactionsRoot = append(e.TransactionsRoot, buf[504:536]...)

	// Field (10) 'ExtraData'
	{
		buf = tail[o10:]
		if len(buf) > 32 {
			return ssz.ErrBytesLength
		}
		if cap(e.ExtraData) == 0 {
			e.ExtraData = make([]byte, 0, len(buf))
		}
		e.ExtraData = append(e.ExtraData, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ExecutionPayloadHeader object
func (e *ExecutionPayloadHeader) SizeSSZ() (size int) {
	size = 536

	// Field (10) 'ExtraData'
	size += len(e.ExtraData)

	return
}

// HashTreeRoot ssz hashes the ExecutionPayloadHeader object
func (e *ExecutionPayloadHeader) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(e)
}

// HashTreeRootWith ssz hashes the ExecutionPayloadHeader object with a hasher
func (e *ExecutionPayloadHeader) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'ParentHash'
	if size := len(e.ParentHash); size != 32 {
		err = ssz.ErrBytesLengthFn("--.ParentHash", size, 32)
		return
	}
	hh.PutBytes(e.ParentHash)

	// Field (1) 'FeeRecipient'
	if size := len(e.FeeRecipient); size != 20 {
		err = ssz.ErrBytesLengthFn("--.FeeRecipient", size, 20)
		return
	}
	hh.PutBytes(e.FeeRecipient)

	// Field (2) 'StateRoot'
	if size := len(e.StateRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.StateRoot", size, 32)
		return
	}
	hh.PutBytes(e.StateRoot)

	// Field (3) 'ReceiptsRoot'
	if size := len(e.ReceiptsRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.ReceiptsRoot", size, 32)
		return
	}
	hh.PutBytes(e.ReceiptsRoot)

	// Field (4) 'LogsBloom'
	if size := len(e.LogsBloom); size != 256 {
		err = ssz.ErrBytesLengthFn("--.LogsBloom", size, 256)
		return
	}
	hh.PutBytes(e.LogsBloom)

	// Field (5) 'PrevRandao'
	if size := len(e.PrevRandao); size != 32 {
		err = ssz.ErrBytesLengthFn("--.PrevRandao", size, 32)
		return
	}
	hh.PutBytes(e.PrevRandao)

	// Field (6) 'BlockNumber'
	hh.PutUint64(e.BlockNumber)

	// Field (7) 'GasLimit'
	hh.PutUint64(e.GasLimit)

	// Field (8) 'GasUsed'
	hh.PutUint64(e.GasUsed)

	// Field (9) 'Timestamp'
	hh.PutUint64(e.Timestamp)

	// Field (10) 'ExtraData'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(e.ExtraData))
		if byteLen > 32 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.PutBytes(e.ExtraData)
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(elemIndx, byteLen, (32+31)/32)
		} else {
			hh.MerkleizeWithMixin(elemIndx, byteLen, (32+31)/32)
		}
	}

	// Field (11) 'BaseFeePerGas'
	if size := len(e.BaseFeePerGas); size != 32 {
		err = ssz.ErrBytesLengthFn("--.BaseFeePerGas", size, 32)
		return
	}
	hh.PutBytes(e.BaseFeePerGas)

	// Field (12) 'BlockHash'
	if size := len(e.BlockHash); size != 32 {
		err = ssz.ErrBytesLengthFn("--.BlockHash", size, 32)
		return
	}
	hh.PutBytes(e.BlockHash)

	// Field (13) 'TransactionsRoot'
	if size := len(e.TransactionsRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.TransactionsRoot", size, 32)
		return
	}
	hh.PutBytes(e.TransactionsRoot)

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the ExecutionPayloadHeaderCapella object
func (e *ExecutionPayloadHeaderCapella) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(e)
}

// MarshalSSZTo ssz marshals the ExecutionPayloadHeaderCapella object to a target array
func (e *ExecutionPayloadHeaderCapella) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(568)

	// Field (0) 'ParentHash'
	if size := len(e.ParentHash); size != 32 {
		err = ssz.ErrBytesLengthFn("--.ParentHash", size, 32)
		return
	}
	dst = append(dst, e.ParentHash...)

	// Field (1) 'FeeRecipient'
	if size := len(e.FeeRecipient); size != 20 {
		err = ssz.ErrBytesLengthFn("--.FeeRecipient", size, 20)
		return
	}
	dst = append(dst, e.FeeRecipient...)

	// Field (2) 'StateRoot'
	if size := len(e.StateRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.StateRoot", size, 32)
		return
	}
	dst = append(dst, e.StateRoot...)

	// Field (3) 'ReceiptsRoot'
	if size := len(e.ReceiptsRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.ReceiptsRoot", size, 32)
		return
	}
	dst = append(dst, e.ReceiptsRoot...)

	// Field (4) 'LogsBloom'
	if size := len(e.LogsBloom); size != 256 {
		err = ssz.ErrBytesLengthFn("--.LogsBloom", size, 256)
		return
	}
	dst = append(dst, e.LogsBloom...)

	// Field (5) 'PrevRandao'
	if size := len(e.PrevRandao); size != 32 {
		err = ssz.ErrBytesLengthFn("--.PrevRandao", size, 32)
		return
	}
	dst = append(dst, e.PrevRandao...)

	// Field (6) 'BlockNumber'
	dst = ssz.MarshalUint64(dst, e.BlockNumber)

	// Field (7) 'GasLimit'
	dst = ssz.MarshalUint64(dst, e.GasLimit)

	// Field (8) 'GasUsed'
	dst = ssz.MarshalUint64(dst, e.GasUsed)

	// Field (9) 'Timestamp'
	dst = ssz.MarshalUint64(dst, e.Timestamp)

	// Offset (10) 'ExtraData'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(e.ExtraData)

	// Field (11) 'BaseFeePerGas'
	if size := len(e.BaseFeePerGas); size != 32 {
		err = ssz.ErrBytesLengthFn("--.BaseFeePerGas", size, 32)
		return
	}
	dst = append(dst, e.BaseFeePerGas...)

	// Field (12) 'BlockHash'
	if size := len(e.BlockHash); size != 32 {
		err = ssz.ErrBytesLengthFn("--.BlockHash", size, 32)
		return
	}
	dst = append(dst, e.BlockHash...)

	// Field (13) 'TransactionsRoot'
	if size := len(e.TransactionsRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.TransactionsRoot", size, 32)
		return
	}
	dst = append(dst, e.TransactionsRoot...)

	// Field (14) 'WithdrawalsRoot'
	if size := len(e.WithdrawalsRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.WithdrawalsRoot", size, 32)
		return
	}
	dst = append(dst, e.WithdrawalsRoot...)

	// Field (10) 'ExtraData'
	if size := len(e.ExtraData); size > 32 {
		err = ssz.ErrBytesLengthFn("--.ExtraData", size, 32)
		return
	}
	dst = append(dst, e.ExtraData...)

	return
}

// UnmarshalSSZ ssz unmarshals the ExecutionPayloadHeaderCapella object
func (e *ExecutionPayloadHeaderCapella) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 568 {
		return ssz.ErrSize
	}

	tail := buf
	var o10 uint64

	// Field (0) 'ParentHash'
	if cap(e.ParentHash) == 0 {
		e.ParentHash = make([]byte, 0, len(buf[0:32]))
	}
	e.ParentHash = append(e.ParentHash, buf[0:32]...)

	// Field (1) 'FeeRecipient'
	if cap(e.FeeRecipient) == 0 {
		e.FeeRecipient = make([]byte, 0, len(buf[32:52]))
	}
	e.FeeRecipient = append(e.FeeRecipient, buf[32:52]...)

	// Field (2) 'StateRoot'
	if cap(e.StateRoot) == 0 {
		e.StateRoot = make([]byte, 0, len(buf[52:84]))
	}
	e.StateRoot = append(e.StateRoot, buf[52:84]...)

	// Field (3) 'ReceiptsRoot'
	if cap(e.ReceiptsRoot) == 0 {
		e.ReceiptsRoot = make([]byte, 0, len(buf[84:116]))
	}
	e.ReceiptsRoot = append(e.ReceiptsRoot, buf[84:116]...)

	// Field (4) 'LogsBloom'
	if cap(e.LogsBloom) == 0 {
		e.LogsBloom = make([]byte, 0, len(buf[116:372]))
	}
	e.LogsBloom = append(e.LogsBloom, buf[116:372]...)

	// Field (5) 'PrevRandao'
	if cap(e.PrevRandao) == 0 {
		e.PrevRandao = make([]byte, 0, len(buf[372:404]))
	}
	e.PrevRandao = append(e.PrevRandao, buf[372:404]...)

	// Field (6) 'BlockNumber'
	e.BlockNumber = ssz.UnmarshallUint64(buf[404:412])

	// Field (7) 'GasLimit'
	e.GasLimit = ssz.UnmarshallUint64(buf[412:420])

	// Field (8) 'GasUsed'
	e.GasUsed = ssz.UnmarshallUint64(buf[420:428])

	// Field (9) 'Timestamp'
	e.Timestamp = ssz.UnmarshallUint64(buf[428:436])

	// Offset (10) 'ExtraData'
	if o10 = ssz.ReadOffset(buf[436:440]); o10 > size {
		return ssz.ErrOffset
	}

	if o10 < 568 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (11) 'BaseFeePerGas'
	if cap(e.BaseFeePerGas) == 0 {
		e.BaseFeePerGas = make([]byte, 0, len(buf[440:472]))
	}
	e.BaseFeePerGas = append(e.BaseFeePerGas, buf[440:472]...)

	// Field (12) 'BlockHash'
	if cap(e.BlockHash) == 0 {
		e.BlockHash = make([]byte, 0, len(buf[472:504]))
	}
	e.BlockHash = append(e.BlockHash, buf[472:504]...)

	// Field (13) 'TransactionsRoot'
	if cap(e.TransactionsRoot) == 0 {
		e.TransactionsRoot = make([]byte, 0, len(buf[504:536]))
	}
	e.TransactionsRoot = append(e.TransactionsRoot, buf[504:536]...)

	// Field (14) 'WithdrawalsRoot'
	if cap(e.WithdrawalsRoot) == 0 {
		e.WithdrawalsRoot = make([]byte, 0, len(buf[536:568]))
	}
	e.WithdrawalsRoot = append(e.WithdrawalsRoot, buf[536:568]...)

	// Field (10) 'ExtraData'
	{
		buf = tail[o10:]
		if len(buf) > 32 {
			return ssz.ErrBytesLength
		}
		if cap(e.ExtraData) == 0 {
			e.ExtraData = make([]byte, 0, len(buf))
		}
		e.ExtraData = append(e.ExtraData, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ExecutionPayloadHeaderCapella object
func (e *ExecutionPayloadHeaderCapella) SizeSSZ() (size int) {
	size = 568

	// Field (10) 'ExtraData'
	size += len(e.ExtraData)

	return
}

// HashTreeRoot ssz hashes the ExecutionPayloadHeaderCapella object
func (e *ExecutionPayloadHeaderCapella) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(e)
}

// HashTreeRootWith ssz hashes the ExecutionPayloadHeaderCapella object with a hasher
func (e *ExecutionPayloadHeaderCapella) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'ParentHash'
	if size := len(e.ParentHash); size != 32 {
		err = ssz.ErrBytesLengthFn("--.ParentHash", size, 32)
		return
	}
	hh.PutBytes(e.ParentHash)

	// Field (1) 'FeeRecipient'
	if size := len(e.FeeRecipient); size != 20 {
		err = ssz.ErrBytesLengthFn("--.FeeRecipient", size, 20)
		return
	}
	hh.PutBytes(e.FeeRecipient)

	// Field (2) 'StateRoot'
	if size := len(e.StateRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.StateRoot", size, 32)
		return
	}
	hh.PutBytes(e.StateRoot)

	// Field (3) 'ReceiptsRoot'
	if size := len(e.ReceiptsRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.ReceiptsRoot", size, 32)
		return
	}
	hh.PutBytes(e.ReceiptsRoot)

	// Field (4) 'LogsBloom'
	if size := len(e.LogsBloom); size != 256 {
		err = ssz.ErrBytesLengthFn("--.LogsBloom", size, 256)
		return
	}
	hh.PutBytes(e.LogsBloom)

	// Field (5) 'PrevRandao'
	if size := len(e.PrevRandao); size != 32 {
		err = ssz.ErrBytesLengthFn("--.PrevRandao", size, 32)
		return
	}
	hh.PutBytes(e.PrevRandao)

	// Field (6) 'BlockNumber'
	hh.PutUint64(e.BlockNumber)

	// Field (7) 'GasLimit'
	hh.PutUint64(e.GasLimit)

	// Field (8) 'GasUsed'
	hh.PutUint64(e.GasUsed)

	// Field (9) 'Timestamp'
	hh.PutUint64(e.Timestamp)

	// Field (10) 'ExtraData'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(e.ExtraData))
		if byteLen > 32 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.PutBytes(e.ExtraData)
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(elemIndx, byteLen, (32+31)/32)
		} else {
			hh.MerkleizeWithMixin(elemIndx, byteLen, (32+31)/32)
		}
	}

	// Field (11) 'BaseFeePerGas'
	if size := len(e.BaseFeePerGas); size != 32 {
		err = ssz.ErrBytesLengthFn("--.BaseFeePerGas", size, 32)
		return
	}
	hh.PutBytes(e.BaseFeePerGas)

	// Field (12) 'BlockHash'
	if size := len(e.BlockHash); size != 32 {
		err = ssz.ErrBytesLengthFn("--.BlockHash", size, 32)
		return
	}
	hh.PutBytes(e.BlockHash)

	// Field (13) 'TransactionsRoot'
	if size := len(e.TransactionsRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.TransactionsRoot", size, 32)
		return
	}
	hh.PutBytes(e.TransactionsRoot)

	// Field (14) 'WithdrawalsRoot'
	if size := len(e.WithdrawalsRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.WithdrawalsRoot", size, 32)
		return
	}
	hh.PutBytes(e.WithdrawalsRoot)

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the Withdrawal object
func (w *Withdrawal) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(w)
}

// MarshalSSZTo ssz marshals the Withdrawal object to a target array
func (w *Withdrawal) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'WithdrawalIndex'
	dst = ssz.MarshalUint64(dst, w.WithdrawalIndex)

	// Field (1) 'ValidatorIndex'
	dst = ssz.MarshalUint64(dst, uint64(w.ValidatorIndex))

	// Field (2) 'ExecutionAddress'
	if size := len(w.ExecutionAddress); size != 20 {
		err = ssz.ErrBytesLengthFn("--.ExecutionAddress", size, 20)
		return
	}
	dst = append(dst, w.ExecutionAddress...)

	// Field (3) 'Amount'
	dst = ssz.MarshalUint64(dst, w.Amount)

	return
}

// UnmarshalSSZ ssz unmarshals the Withdrawal object
func (w *Withdrawal) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 44 {
		return ssz.ErrSize
	}

	// Field (0) 'WithdrawalIndex'
	w.WithdrawalIndex = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'ValidatorIndex'
	w.ValidatorIndex = github_com_prysmaticlabs_prysm_v3_consensus_types_primitives.ValidatorIndex(ssz.UnmarshallUint64(buf[8:16]))

	// Field (2) 'ExecutionAddress'
	if cap(w.ExecutionAddress) == 0 {
		w.ExecutionAddress = make([]byte, 0, len(buf[16:36]))
	}
	w.ExecutionAddress = append(w.ExecutionAddress, buf[16:36]...)

	// Field (3) 'Amount'
	w.Amount = ssz.UnmarshallUint64(buf[36:44])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Withdrawal object
func (w *Withdrawal) SizeSSZ() (size int) {
	size = 44
	return
}

// HashTreeRoot ssz hashes the Withdrawal object
func (w *Withdrawal) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(w)
}

// HashTreeRootWith ssz hashes the Withdrawal object with a hasher
func (w *Withdrawal) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'WithdrawalIndex'
	hh.PutUint64(w.WithdrawalIndex)

	// Field (1) 'ValidatorIndex'
	hh.PutUint64(uint64(w.ValidatorIndex))

	// Field (2) 'ExecutionAddress'
	if size := len(w.ExecutionAddress); size != 20 {
		err = ssz.ErrBytesLengthFn("--.ExecutionAddress", size, 20)
		return
	}
	hh.PutBytes(w.ExecutionAddress)

	// Field (3) 'Amount'
	hh.PutUint64(w.Amount)

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}