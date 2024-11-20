package electra

import (
	"github.com/protolambda/zrnt/eth2/beacon/common"
	"github.com/protolambda/ztyp/codec"
	"github.com/protolambda/ztyp/tree"
)

type CommitteeIndicesElectra []common.ValidatorIndex

func (p *CommitteeIndicesElectra) Deserialize(spec *common.Spec, dr *codec.DecodingReader) error {
	return dr.List(func() codec.Deserializable {
		i := len(*p)
		*p = append(*p, common.ValidatorIndex(0))
		return &((*p)[i])
	}, common.ValidatorIndexType.TypeByteLength(), uint64(spec.MAX_VALIDATORS_PER_COMMITTEE_ELECTRA))
}

func (a CommitteeIndicesElectra) Serialize(_ *common.Spec, w *codec.EncodingWriter) error {
	return w.List(func(i uint64) codec.Serializable {
		return a[i]
	}, common.ValidatorIndexType.TypeByteLength(), uint64(len(a)))
}

func (a CommitteeIndicesElectra) ByteLength(*common.Spec) uint64 {
	return common.ValidatorIndexType.TypeByteLength() * uint64(len(a))
}

func (*CommitteeIndicesElectra) FixedLength(*common.Spec) uint64 {
	return 0
}

func (p CommitteeIndicesElectra) HashTreeRoot(spec *common.Spec, hFn tree.HashFn) common.Root {
	return hFn.Uint64ListHTR(func(i uint64) uint64 {
		return uint64(p[i])
	}, uint64(len(p)), uint64(spec.MAX_VALIDATORS_PER_COMMITTEE_ELECTRA))
}
