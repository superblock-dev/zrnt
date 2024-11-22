package electra

import (
	"bytes"
	"fmt"

	"github.com/protolambda/zrnt/eth2/beacon/common"
	"github.com/protolambda/ztyp/bitfields"
	"github.com/protolambda/ztyp/codec"
	"github.com/protolambda/ztyp/conv"
	"github.com/protolambda/ztyp/tree"
	. "github.com/protolambda/ztyp/view"
)

// AttestationBits is formatted as a serialized SSZ bitlist, including the delimit bit
type AttestationBitsElectra []byte

func (li AttestationBitsElectra) View(spec *common.Spec) *AttestationBitsElectraView {
	v, _ := AttestationBitsElectraType(spec).Deserialize(codec.NewDecodingReader(bytes.NewReader(li), uint64(len(li))))
	return &AttestationBitsElectraView{v.(*BitListView)}
}

func (li *AttestationBitsElectra) Deserialize(spec *common.Spec, dr *codec.DecodingReader) error {
	return dr.BitList((*[]byte)(li), uint64(spec.MAX_ATTESTING_INDICES))
}

func (a AttestationBitsElectra) Serialize(spec *common.Spec, w *codec.EncodingWriter) error {
	return w.BitList(a[:])
}

func (a AttestationBitsElectra) ByteLength(spec *common.Spec) uint64 {
	return uint64(len(a))
}

func (a *AttestationBitsElectra) FixedLength(*common.Spec) uint64 {
	return 0
}

func (li AttestationBitsElectra) HashTreeRoot(spec *common.Spec, hFn tree.HashFn) common.Root {
	return hFn.BitListHTR(li, uint64(spec.MAX_ATTESTING_INDICES))
}

func (cb AttestationBitsElectra) MarshalText() ([]byte, error) {
	return conv.BytesMarshalText(cb[:])
}

func (cb *AttestationBitsElectra) UnmarshalText(text []byte) error {
	return conv.DynamicBytesUnmarshalText((*[]byte)(cb), text)
}

func (cb AttestationBitsElectra) String() string {
	return conv.BytesString(cb[:])
}

func (cb AttestationBitsElectra) BitLen() uint64 {
	return bitfields.BitlistLen(cb)
}

func (cb AttestationBitsElectra) GetBit(i uint64) bool {
	return bitfields.GetBit(cb, i)
}

func (cb AttestationBitsElectra) SetBit(i uint64, v bool) {
	bitfields.SetBit(cb, i, v)
}

// Sets the bits to true that are true in other. (in place)
func (cb AttestationBitsElectra) Or(other AttestationBitsElectra) {
	for i := 0; i < len(cb); i++ {
		cb[i] |= other[i]
	}
}

// In-place filters a list of committees indices to only keep the bitfield participants.
// The result is not sorted. Returns the re-sliced filtered participants list.
//
// WARNING: unsafe to use, panics if committee size does not match.
func (cb AttestationBitsElectra) FilterParticipants(committee []common.ValidatorIndex) []common.ValidatorIndex {
	out := committee[:0]
	bitLen := cb.BitLen()
	if bitLen != uint64(len(committee)) {
		panic("committee mismatch, bitfield length does not match")
	}
	for i := uint64(0); i < bitLen; i++ {
		if bitfields.GetBit(cb, i) {
			out = append(out, committee[i])
		}
	}
	return out
}

// In-place filters a list of committees indices to only keep the bitfield NON-participants.
// The result is not sorted. Returns the re-sliced filtered non-participants list.
//
// WARNING: unsafe to use, panics if committee size does not match.
func (cb AttestationBitsElectra) FilterNonParticipants(committee []common.ValidatorIndex) []common.ValidatorIndex {
	out := committee[:0]
	bitLen := cb.BitLen()
	if bitLen != uint64(len(committee)) {
		panic("committee mismatch, bitfield length does not match")
	}
	for i := uint64(0); i < bitLen; i++ {
		if !bitfields.GetBit(cb, i) {
			out = append(out, committee[i])
		}
	}
	return out
}

// Returns true if other only has bits set to 1 that this bitfield also has set to 1
func (cb AttestationBitsElectra) Covers(other AttestationBitsElectra) (bool, error) {
	if a, b := cb.BitLen(), other.BitLen(); a != b {
		return false, fmt.Errorf("bitfield length mismatch: %d <> %d", a, b)
	}
	return bitfields.Covers(cb, other)
}

func (cb AttestationBitsElectra) OnesCount() uint64 {
	return bitfields.BitlistOnesCount(cb)
}

func (cb AttestationBitsElectra) SingleParticipant(committee []common.ValidatorIndex) (common.ValidatorIndex, error) {
	bitLen := cb.BitLen()
	if bitLen != uint64(len(committee)) {
		return 0, fmt.Errorf("committee mismatch, bitfield length %d does not match committee size %d", bitLen, len(committee))
	}
	var found *common.ValidatorIndex
	for i := uint64(0); i < bitLen; i++ {
		if cb.GetBit(i) {
			if found == nil {
				found = &committee[i]
			} else {
				return 0, fmt.Errorf("found at least two participants: %d and %d", *found, committee[i])
			}
		}
	}
	if found == nil {
		return 0, fmt.Errorf("found no participants")
	}
	return *found, nil
}

func (cb AttestationBitsElectra) Copy() AttestationBitsElectra {
	// append won't find capacity, and thus put contents into new array, and then returns typed slice of it.
	return append(AttestationBitsElectra(nil), cb...)
}

func AttestationBitsElectraType(spec *common.Spec) *BitListTypeDef {
	return BitListType(uint64(spec.MAX_ATTESTING_INDICES))
}

type AttestationBitsElectraView struct {
	*BitListView
}

func AsAttestationBits(v View, err error) (*AttestationBitsElectraView, error) {
	c, err := AsBitList(v, err)
	return &AttestationBitsElectraView{c}, err
}

func (v *AttestationBitsElectraView) Raw() (AttestationBitsElectra, error) {
	bitLength, err := v.Length()
	if err != nil {
		return nil, err
	}
	// rounded up, and then an extra bit for delimiting. ((bitLength + 7 + 1)/ 8)
	byteLength := (bitLength / 8) + 1
	var buf bytes.Buffer
	if err := v.Serialize(codec.NewEncodingWriter(&buf)); err != nil {
		return nil, err
	}
	out := AttestationBitsElectra(buf.Bytes())
	if uint64(len(out)) != byteLength {
		return nil, fmt.Errorf("failed to convert attestation tree bits view to raw bits")
	}
	return out, nil
}

// AttestationBits is formatted as a serialized SSZ bitlist, including the delimit bit
type CommitteeBits []byte

func (li CommitteeBits) View(spec *common.Spec) *CommitteeBitsView {
	v, _ := CommitteeBitsType(spec).Deserialize(codec.NewDecodingReader(bytes.NewReader(li), uint64(len(li))))
	return &CommitteeBitsView{v.(*BitListView)}
}

func (li *CommitteeBits) Deserialize(spec *common.Spec, dr *codec.DecodingReader) error {
	return dr.BitList((*[]byte)(li), uint64(spec.COMMITTEE_BITS))
}

func (a CommitteeBits) Serialize(spec *common.Spec, w *codec.EncodingWriter) error {
	return w.BitList(a[:])
}

func (a CommitteeBits) ByteLength(spec *common.Spec) uint64 {
	return uint64(len(a))
}

func (a *CommitteeBits) FixedLength(*common.Spec) uint64 {
	return 0
}

func (li CommitteeBits) HashTreeRoot(spec *common.Spec, hFn tree.HashFn) common.Root {
	return hFn.BitListHTR(li, uint64(spec.COMMITTEE_BITS))
}

func (cb CommitteeBits) MarshalText() ([]byte, error) {
	return conv.BytesMarshalText(cb[:])
}

func (cb *CommitteeBits) UnmarshalText(text []byte) error {
	return conv.DynamicBytesUnmarshalText((*[]byte)(cb), text)
}

func (cb CommitteeBits) String() string {
	return conv.BytesString(cb[:])
}

func (cb CommitteeBits) BitLen() uint64 {
	return bitfields.BitlistLen(cb)
}

func (cb CommitteeBits) GetBit(i uint64) bool {
	return bitfields.GetBit(cb, i)
}

func (cb CommitteeBits) SetBit(i uint64, v bool) {
	bitfields.SetBit(cb, i, v)
}

// Sets the bits to true that are true in other. (in place)
func (cb CommitteeBits) Or(other CommitteeBits) {
	for i := 0; i < len(cb); i++ {
		cb[i] |= other[i]
	}
}

func (cb CommitteeBits) Copy() CommitteeBits {
	// append won't find capacity, and thus put contents into new array, and then returns typed slice of it.
	return append(CommitteeBits(nil), cb...)
}

func CommitteeBitsType(spec *common.Spec) *BitListTypeDef {
	return BitListType(uint64(spec.COMMITTEE_BITS))
}

type CommitteeBitsView struct {
	*BitListView
}

func AsCommitteeBits(v View, err error) (*CommitteeBitsView, error) {
	c, err := AsBitList(v, err)
	return &CommitteeBitsView{c}, err
}

func (v *CommitteeBitsView) Raw() (CommitteeBits, error) {
	bitLength, err := v.Length()
	if err != nil {
		return nil, err
	}
	// rounded up, and then an extra bit for delimiting. ((bitLength + 7 + 1)/ 8)
	byteLength := (bitLength / 8) + 1
	var buf bytes.Buffer
	if err := v.Serialize(codec.NewEncodingWriter(&buf)); err != nil {
		return nil, err
	}
	out := CommitteeBits(buf.Bytes())
	if uint64(len(out)) != byteLength {
		return nil, fmt.Errorf("failed to convert Committee tree bits view to raw bits")
	}
	return out, nil
}
