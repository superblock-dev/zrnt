package common

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/protolambda/ztyp/codec"
	"github.com/protolambda/ztyp/tree"
	. "github.com/protolambda/ztyp/view"
)

type WithdrawalPrefix [1]byte

func (p WithdrawalPrefix) MarshalText() ([]byte, error) {
	return []byte("0x" + hex.EncodeToString(p[:])), nil
}

func (p WithdrawalPrefix) String() string {
	return "0x" + hex.EncodeToString(p[:])
}

func (p *WithdrawalPrefix) UnmarshalText(text []byte) error {
	if p == nil {
		return errors.New("cannot decode into nil WithdrawalPrefix")
	}
	if len(text) >= 2 && text[0] == '0' && (text[1] == 'x' || text[1] == 'X') {
		text = text[2:]
	}
	if len(text) != 2 {
		return fmt.Errorf("unexpected length string '%s'", string(text))
	}
	_, err := hex.Decode(p[:], text)
	return err
}

const WithdrawalIndexType = Uint64Type

type WithdrawalIndex Uint64View

func AsWithdrawalIndex(v View, err error) (WithdrawalIndex, error) {
	i, err := AsUint64(v, err)
	return WithdrawalIndex(i), err
}

func (a *WithdrawalIndex) Deserialize(dr *codec.DecodingReader) error {
	return (*Uint64View)(a).Deserialize(dr)
}

func (i WithdrawalIndex) Serialize(w *codec.EncodingWriter) error {
	return w.WriteUint64(uint64(i))
}

func (WithdrawalIndex) ByteLength() uint64 {
	return 8
}

func (WithdrawalIndex) FixedLength() uint64 {
	return 8
}

func (t WithdrawalIndex) HashTreeRoot(hFn tree.HashFn) Root {
	return Uint64View(t).HashTreeRoot(hFn)
}

func (e WithdrawalIndex) MarshalJSON() ([]byte, error) {
	return Uint64View(e).MarshalJSON()
}

func (e *WithdrawalIndex) UnmarshalJSON(b []byte) error {
	return ((*Uint64View)(e)).UnmarshalJSON(b)
}

func (e WithdrawalIndex) String() string {
	return Uint64View(e).String()
}

var WithdrawalType = ContainerType("Withdrawal", []FieldDef{
	{"index", WithdrawalIndexType},
	{"validator_index", ValidatorIndexType},
	{"address", Eth1AddressType},
	{"amount", GweiType},
})

type WithdrawalView struct {
	*ContainerView
}

func (v *WithdrawalView) Raw() (*Withdrawal, error) {
	values, err := v.FieldValues()
	if err != nil {
		return nil, err
	}
	if len(values) != 4 {
		return nil, fmt.Errorf("unexpected number of withdrawal fields: %d", len(values))
	}
	index, err := AsWithdrawalIndex(values[0], err)
	validatorIndex, err := AsValidatorIndex(values[1], err)
	address, err := AsEth1Address(values[2], err)
	amount, err := AsGwei(values[3], err)
	if err != nil {
		return nil, err
	}
	return &Withdrawal{
		Index:          index,
		ValidatorIndex: validatorIndex,
		Address:        address,
		Amount:         amount,
	}, nil
}

func (v *WithdrawalView) Index() (WithdrawalIndex, error) {
	return AsWithdrawalIndex(v.Get(0))
}

func (v *WithdrawalView) ValidatorIndex() (ValidatorIndex, error) {
	return AsValidatorIndex(v.Get(1))
}

func (v *WithdrawalView) Address() (Eth1Address, error) {
	return AsEth1Address(v.Get(2))
}

func (v *WithdrawalView) Amount() (Gwei, error) {
	return AsGwei(v.Get(3))
}

func AsWithdrawal(v View, err error) (*WithdrawalView, error) {
	c, err := AsContainer(v, err)
	return &WithdrawalView{c}, err
}

type Withdrawal struct {
	Index          WithdrawalIndex `json:"index" yaml:"index"`
	ValidatorIndex ValidatorIndex  `json:"validator_index" yaml:"validator_index"`
	Address        Eth1Address     `json:"address" yaml:"address"`
	Amount         Gwei            `json:"amount" yaml:"amount"`
}

func (s *Withdrawal) View() *WithdrawalView {
	i, vi, ad, am := s.Index, s.ValidatorIndex, s.Address, s.Amount
	v, err := AsWithdrawal(WithdrawalType.FromFields(Uint64View(i), Uint64View(vi), ad.View(), Uint64View(am)))
	if err != nil {
		panic(err)
	}
	return v
}

func (s *Withdrawal) Deserialize(dr *codec.DecodingReader) error {
	return dr.FixedLenContainer(&s.Index, &s.ValidatorIndex, &s.Address, &s.Amount)
}

func (s *Withdrawal) Serialize(w *codec.EncodingWriter) error {
	return w.FixedLenContainer(&s.Index, &s.ValidatorIndex, &s.Address, &s.Amount)
}

func (s *Withdrawal) ByteLength() uint64 {
	return Uint64Type.TypeByteLength()*3 + Eth1AddressType.TypeByteLength()
}

func (s *Withdrawal) FixedLength() uint64 {
	return Uint64Type.TypeByteLength()*3 + Eth1AddressType.TypeByteLength()
}

func (s *Withdrawal) HashTreeRoot(hFn tree.HashFn) Root {
	return hFn.HashTreeRoot(&s.Index, &s.ValidatorIndex, &s.Address, &s.Amount)
}

func WithdrawalsType(spec *Spec) ListTypeDef {
	return ListType(WithdrawalType, uint64(spec.MAX_WITHDRAWALS_PER_PAYLOAD))
}

type Withdrawals []Withdrawal

func (ws *Withdrawals) Deserialize(spec *Spec, dr *codec.DecodingReader) error {
	return dr.List(func() codec.Deserializable {
		i := len(*ws)
		*ws = append(*ws, Withdrawal{})
		return &((*ws)[i])
	}, WithdrawalType.TypeByteLength(), uint64(spec.MAX_WITHDRAWALS_PER_PAYLOAD))
}

func (ws Withdrawals) Serialize(spec *Spec, w *codec.EncodingWriter) error {
	return w.List(func(i uint64) codec.Serializable {
		return &ws[i]
	}, WithdrawalType.TypeByteLength(), uint64(len(ws)))
}

func (ws Withdrawals) ByteLength(spec *Spec) (out uint64) {
	return WithdrawalType.TypeByteLength() * uint64(len(ws))
}

func (ws *Withdrawals) FixedLength(*Spec) uint64 {
	return 0
}

func (ws Withdrawals) HashTreeRoot(spec *Spec, hFn tree.HashFn) Root {
	length := uint64(len(ws))
	return hFn.ComplexListHTR(func(i uint64) tree.HTR {
		if i < length {
			return &ws[i]
		}
		return nil
	}, length, uint64(spec.MAX_WITHDRAWALS_PER_PAYLOAD))
}

func (ws Withdrawals) MarshalJSON() ([]byte, error) {
	if ws == nil {
		return json.Marshal([]Withdrawal{}) // encode as empty list, not null
	}
	return json.Marshal([]Withdrawal(ws))
}
