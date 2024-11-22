package electra

import (
	"encoding/json"

	"github.com/protolambda/zrnt/eth2/beacon/common"
	"github.com/protolambda/zrnt/eth2/beacon/phase0"
	"github.com/protolambda/ztyp/codec"
	"github.com/protolambda/ztyp/tree"
	. "github.com/protolambda/ztyp/view"
)

func BlockAttestationsElectraType(spec *common.Spec) ListTypeDef {
	return ListType(AttestationElectraType(spec), uint64(spec.MAX_ATTESTATIONS_ALPACA))
}

func AttestationElectraType(spec *common.Spec) *ContainerTypeDef {
	return ContainerType("Attestation", []FieldDef{
		{"aggregation_bits", AttestationBitsElectraType(spec)},
		{"data", phase0.AttestationDataType},
		{"signature", common.BLSSignatureType},
		{"committee_bits", CommitteeBitsType(spec)},
	})
}

type AttestationElectra struct {
	AggregationBits AttestationBitsElectra `json:"aggregation_bits" yaml:"aggregation_bits"`
	Data            phase0.AttestationData `json:"data" yaml:"data"`
	Signature       common.BLSSignature    `json:"signature" yaml:"signature"`
	CommitteeBits   CommitteeBits          `json:"committee_bits" yaml:"committee_bits"`
}

func (a *AttestationElectra) Deserialize(spec *common.Spec, dr *codec.DecodingReader) error {
	return dr.Container(spec.Wrap(&a.AggregationBits), &a.Data, &a.Signature, spec.Wrap(&a.CommitteeBits))
}

func (a *AttestationElectra) Serialize(spec *common.Spec, w *codec.EncodingWriter) error {
	return w.Container(spec.Wrap(&a.AggregationBits), &a.Data, &a.Signature, spec.Wrap(&a.CommitteeBits))
}

func (a *AttestationElectra) ByteLength(spec *common.Spec) uint64 {
	return codec.ContainerLength(spec.Wrap(&a.AggregationBits), &a.Data, &a.Signature, spec.Wrap(&a.CommitteeBits))
}

func (a *AttestationElectra) FixedLength(*common.Spec) uint64 {
	return 0
}

func (a *AttestationElectra) HashTreeRoot(spec *common.Spec, hFn tree.HashFn) common.Root {
	return hFn.HashTreeRoot(spec.Wrap(&a.AggregationBits), &a.Data, a.Signature, spec.Wrap(&a.CommitteeBits))
}

type AttestationsElectra []AttestationElectra

func (a *AttestationsElectra) Deserialize(spec *common.Spec, dr *codec.DecodingReader) error {
	return dr.List(func() codec.Deserializable {
		i := len(*a)
		*a = append(*a, AttestationElectra{})
		return spec.Wrap(&((*a)[i]))
	}, 0, uint64(spec.MAX_ATTESTATIONS_ALPACA))
}

func (a AttestationsElectra) Serialize(spec *common.Spec, w *codec.EncodingWriter) error {
	return w.List(func(i uint64) codec.Serializable {
		return spec.Wrap(&a[i])
	}, 0, uint64(len(a)))
}

func (a AttestationsElectra) ByteLength(spec *common.Spec) (out uint64) {
	for _, v := range a {
		out += v.ByteLength(spec) + codec.OFFSET_SIZE
	}
	return
}

func (a *AttestationsElectra) FixedLength(*common.Spec) uint64 {
	return 0
}

func (li AttestationsElectra) HashTreeRoot(spec *common.Spec, hFn tree.HashFn) common.Root {
	length := uint64(len(li))
	return hFn.ComplexListHTR(func(i uint64) tree.HTR {
		if i < length {
			return spec.Wrap(&li[i])
		}
		return nil
	}, length, uint64(spec.MAX_ATTESTATIONS_ALPACA))
}

func (li AttestationsElectra) MarshalJSON() ([]byte, error) {
	if li == nil {
		return json.Marshal([]AttestationElectra{}) // encode as empty list, not null
	}
	return json.Marshal([]AttestationElectra(li))
}
