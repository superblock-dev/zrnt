package common

import (
	"encoding/json"

	"github.com/protolambda/ztyp/codec"
	"github.com/protolambda/ztyp/tree"
	. "github.com/protolambda/ztyp/view"
	"gopkg.in/yaml.v3"
)

const TARGET_AGGREGATORS_PER_COMMITTEE = 16
const RANDOM_SUBNETS_PER_VALIDATOR = 1
const EPOCHS_PER_RANDOM_SUBNET_SUBSCRIPTION = 256
const BLS_WITHDRAWAL_PREFIX = 0
const ETH1_ADDRESS_WITHDRAWAL_PREFIX = 1
const SYNC_COMMITTEE_SUBNET_COUNT = 4
const TARGET_AGGREGATORS_PER_SYNC_SUBCOMMITTEE = 16

// Phase0
var DOMAIN_BEACON_PROPOSER = BLSDomainType{0x00, 0x00, 0x00, 0x00}
var DOMAIN_BEACON_ATTESTER = BLSDomainType{0x01, 0x00, 0x00, 0x00}
var DOMAIN_RANDAO = BLSDomainType{0x02, 0x00, 0x00, 0x00}
var DOMAIN_DEPOSIT = BLSDomainType{0x03, 0x00, 0x00, 0x00}
var DOMAIN_VOLUNTARY_EXIT = BLSDomainType{0x04, 0x00, 0x00, 0x00}
var DOMAIN_SELECTION_PROOF = BLSDomainType{0x05, 0x00, 0x00, 0x00}
var DOMAIN_AGGREGATE_AND_PROOF = BLSDomainType{0x06, 0x00, 0x00, 0x00}

// Altair
var DOMAIN_SYNC_COMMITTEE = BLSDomainType{0x07, 0x00, 0x00, 0x00}
var DOMAIN_SYNC_COMMITTEE_SELECTION_PROOF = BLSDomainType{0x08, 0x00, 0x00, 0x00}
var DOMAIN_CONTRIBUTION_AND_PROOF = BLSDomainType{0x09, 0x00, 0x00, 0x00}

// Capella
var DOMAIN_BLS_TO_EXECUTION_CHANGE = BLSDomainType{0x0A, 0x00, 0x00, 0x00}

// Deneb
const BLOB_TX_TYPE = 0x03
const VERSIONED_HASH_VERSION_KZG = 0x01

type Phase0Preset struct {
	// Misc.
	MAX_COMMITTEES_PER_SLOT      Uint64View `yaml:"MAX_COMMITTEES_PER_SLOT" json:"MAX_COMMITTEES_PER_SLOT"`
	TARGET_COMMITTEE_SIZE        Uint64View `yaml:"TARGET_COMMITTEE_SIZE" json:"TARGET_COMMITTEE_SIZE"`
	MAX_VALIDATORS_PER_COMMITTEE Uint64View `yaml:"MAX_VALIDATORS_PER_COMMITTEE" json:"MAX_VALIDATORS_PER_COMMITTEE"`
	SHUFFLE_ROUND_COUNT          Uint8View  `yaml:"SHUFFLE_ROUND_COUNT" json:"SHUFFLE_ROUND_COUNT"`

	// Balance math
	HYSTERESIS_QUOTIENT            Uint64View `yaml:"HYSTERESIS_QUOTIENT" json:"HYSTERESIS_QUOTIENT"`
	HYSTERESIS_DOWNWARD_MULTIPLIER Uint64View `yaml:"HYSTERESIS_DOWNWARD_MULTIPLIER" json:"HYSTERESIS_DOWNWARD_MULTIPLIER"`
	HYSTERESIS_UPWARD_MULTIPLIER   Uint64View `yaml:"HYSTERESIS_UPWARD_MULTIPLIER" json:"HYSTERESIS_UPWARD_MULTIPLIER"`

	// Gwei values
	MIN_DEPOSIT_AMOUNT          Gwei `yaml:"MIN_DEPOSIT_AMOUNT" json:"MIN_DEPOSIT_AMOUNT"`
	MAX_EFFECTIVE_BALANCE       Gwei `yaml:"MAX_EFFECTIVE_BALANCE" json:"MAX_EFFECTIVE_BALANCE"`
	EFFECTIVE_BALANCE_INCREMENT Gwei `yaml:"EFFECTIVE_BALANCE_INCREMENT" json:"EFFECTIVE_BALANCE_INCREMENT"`

	// Time parameters
	MIN_ATTESTATION_INCLUSION_DELAY  Slot  `yaml:"MIN_ATTESTATION_INCLUSION_DELAY" json:"MIN_ATTESTATION_INCLUSION_DELAY"`
	SLOTS_PER_EPOCH                  Slot  `yaml:"SLOTS_PER_EPOCH" json:"SLOTS_PER_EPOCH"`
	MIN_SEED_LOOKAHEAD               Epoch `yaml:"MIN_SEED_LOOKAHEAD" json:"MIN_SEED_LOOKAHEAD"`
	MAX_SEED_LOOKAHEAD               Epoch `yaml:"MAX_SEED_LOOKAHEAD" json:"MAX_SEED_LOOKAHEAD"`
	SLOTS_PER_HISTORICAL_ROOT        Slot  `yaml:"SLOTS_PER_HISTORICAL_ROOT" json:"SLOTS_PER_HISTORICAL_ROOT"`
	MIN_EPOCHS_TO_INACTIVITY_PENALTY Epoch `yaml:"MIN_EPOCHS_TO_INACTIVITY_PENALTY" json:"MIN_EPOCHS_TO_INACTIVITY_PENALTY"`

	// State vector lengths
	EPOCHS_PER_HISTORICAL_VECTOR    Epoch      `yaml:"EPOCHS_PER_HISTORICAL_VECTOR" json:"EPOCHS_PER_HISTORICAL_VECTOR"`
	MIN_SLASHING_WITHDRAWABLE_DELAY Epoch      `yaml:"MIN_SLASHING_WITHDRAWABLE_DELAY" json:"MIN_SLASHING_WITHDRAWABLE_DELAY"`
	HISTORICAL_ROOTS_LIMIT          Uint64View `yaml:"HISTORICAL_ROOTS_LIMIT" json:"HISTORICAL_ROOTS_LIMIT"`
	VALIDATOR_REGISTRY_LIMIT        Uint64View `yaml:"VALIDATOR_REGISTRY_LIMIT" json:"VALIDATOR_REGISTRY_LIMIT"`

	// Reward and penalty quotients
	BASE_REWARD_FACTOR            Uint64View `yaml:"BASE_REWARD_FACTOR" json:"BASE_REWARD_FACTOR"`
	WHISTLEBLOWER_REWARD_QUOTIENT Uint64View `yaml:"WHISTLEBLOWER_REWARD_QUOTIENT" json:"WHISTLEBLOWER_REWARD_QUOTIENT"`
	PROPOSER_REWARD_QUOTIENT      Uint64View `yaml:"PROPOSER_REWARD_QUOTIENT" json:"PROPOSER_REWARD_QUOTIENT"`
	INACTIVITY_PENALTY_QUOTIENT   Uint64View `yaml:"INACTIVITY_PENALTY_QUOTIENT" json:"INACTIVITY_PENALTY_QUOTIENT"`
	MIN_SLASHING_PENALTY_QUOTIENT Uint64View `yaml:"MIN_SLASHING_PENALTY_QUOTIENT" json:"MIN_SLASHING_PENALTY_QUOTIENT"`

	// Max operations per block
	MAX_PROPOSER_SLASHINGS Uint64View `yaml:"MAX_PROPOSER_SLASHINGS" json:"MAX_PROPOSER_SLASHINGS"`
	MAX_ATTESTER_SLASHINGS Uint64View `yaml:"MAX_ATTESTER_SLASHINGS" json:"MAX_ATTESTER_SLASHINGS"`
	MAX_ATTESTATIONS       Uint64View `yaml:"MAX_ATTESTATIONS" json:"MAX_ATTESTATIONS"`
	MAX_DEPOSITS           Uint64View `yaml:"MAX_DEPOSITS" json:"MAX_DEPOSITS"`
	MAX_VOLUNTARY_EXITS    Uint64View `yaml:"MAX_VOLUNTARY_EXITS" json:"MAX_VOLUNTARY_EXITS"`
}

type AltairPreset struct {
	// Updated penalty values
	INACTIVITY_PENALTY_QUOTIENT_ALTAIR   Uint64View `yaml:"INACTIVITY_PENALTY_QUOTIENT_ALTAIR" json:"INACTIVITY_PENALTY_QUOTIENT_ALTAIR"`
	MIN_SLASHING_PENALTY_QUOTIENT_ALTAIR Uint64View `yaml:"MIN_SLASHING_PENALTY_QUOTIENT_ALTAIR" json:"MIN_SLASHING_PENALTY_QUOTIENT_ALTAIR"`

	// Sync committee
	SYNC_COMMITTEE_SIZE              Uint64View `yaml:"SYNC_COMMITTEE_SIZE" json:"SYNC_COMMITTEE_SIZE"`
	EPOCHS_PER_SYNC_COMMITTEE_PERIOD Epoch      `yaml:"EPOCHS_PER_SYNC_COMMITTEE_PERIOD" json:"EPOCHS_PER_SYNC_COMMITTEE_PERIOD"`

	// Sync committees and light clients
	MIN_SYNC_COMMITTEE_PARTICIPANTS Uint64View `yaml:"MIN_SYNC_COMMITTEE_PARTICIPANTS" json:"MIN_SYNC_COMMITTEE_PARTICIPANTS"`
}

type BellatrixPreset struct {
	INACTIVITY_PENALTY_QUOTIENT_BELLATRIX   Uint64View `yaml:"INACTIVITY_PENALTY_QUOTIENT_BELLATRIX" json:"INACTIVITY_PENALTY_QUOTIENT_BELLATRIX"`
	MIN_SLASHING_PENALTY_QUOTIENT_BELLATRIX Uint64View `yaml:"MIN_SLASHING_PENALTY_QUOTIENT_BELLATRIX" json:"MIN_SLASHING_PENALTY_QUOTIENT_BELLATRIX"`
	MAX_BYTES_PER_TRANSACTION               Uint64View `yaml:"MAX_BYTES_PER_TRANSACTION" json:"MAX_BYTES_PER_TRANSACTION"`
	MAX_TRANSACTIONS_PER_PAYLOAD            Uint64View `yaml:"MAX_TRANSACTIONS_PER_PAYLOAD" json:"MAX_TRANSACTIONS_PER_PAYLOAD"`
	BYTES_PER_LOGS_BLOOM                    Uint64View `yaml:"BYTES_PER_LOGS_BLOOM" json:"BYTES_PER_LOGS_BLOOM"`
	MAX_EXTRA_DATA_BYTES                    Uint64View `yaml:"MAX_EXTRA_DATA_BYTES" json:"MAX_EXTRA_DATA_BYTES"`
}

type CapellaPreset struct {
	MAX_BLS_TO_EXECUTION_CHANGES         Uint64View `yaml:"MAX_BLS_TO_EXECUTION_CHANGES" json:"MAX_BLS_TO_EXECUTION_CHANGES"`
	MAX_WITHDRAWALS_PER_PAYLOAD          Uint64View `yaml:"MAX_WITHDRAWALS_PER_PAYLOAD" json:"MAX_WITHDRAWALS_PER_PAYLOAD"`
	MAX_VALIDATORS_PER_WITHDRAWALS_SWEEP Uint64View `yaml:"MAX_VALIDATORS_PER_WITHDRAWALS_SWEEP" json:"MAX_VALIDATORS_PER_WITHDRAWALS_SWEEP"`
}

type DenebPreset struct {
	FIELD_ELEMENTS_PER_BLOB              Uint64View `yaml:"FIELD_ELEMENTS_PER_BLOB" json:"FIELD_ELEMENTS_PER_BLOB"`
	MAX_BLOB_COMMITMENTS_PER_BLOCK       Uint64View `yaml:"MAX_BLOB_COMMITMENTS_PER_BLOCK" json:"MAX_BLOB_COMMITMENTS_PER_BLOCK"`
	MAX_BLOBS_PER_BLOCK                  Uint64View `yaml:"MAX_BLOBS_PER_BLOCK" json:"MAX_BLOBS_PER_BLOCK"`
	KZG_COMMITMENT_INCLUSION_PROOF_DEPTH Uint64View `yaml:"KZG_COMMITMENT_INCLUSION_PROOF_DEPTH" json:"KZG_COMMITMENT_INCLUSION_PROOF_DEPTH"`
}

type ElectraPreset struct {
	CONFIG_TO_ADD Uint64View `yaml:"CONFIG_TO_ADD" json:"CONFIG_TO_ADD"`
}

type Config struct {
	PRESET_BASE string `yaml:"PRESET_BASE" json:"PRESET_BASE"`

	CONFIG_NAME string `yaml:"CONFIG_NAME" json:"CONFIG_NAME"`

	// Transition
	TERMINAL_TOTAL_DIFFICULTY            Uint256View `yaml:"TERMINAL_TOTAL_DIFFICULTY" json:"TERMINAL_TOTAL_DIFFICULTY"`
	TERMINAL_BLOCK_HASH                  Hash32      `yaml:"TERMINAL_BLOCK_HASH" json:"TERMINAL_BLOCK_HASH"`
	TERMINAL_BLOCK_HASH_ACTIVATION_EPOCH Epoch       `yaml:"TERMINAL_BLOCK_HASH_ACTIVATION_EPOCH" json:"TERMINAL_BLOCK_HASH_ACTIVATION_EPOCH"`

	// Genesis.
	MIN_GENESIS_ACTIVE_VALIDATOR_COUNT Uint64View `yaml:"MIN_GENESIS_ACTIVE_VALIDATOR_COUNT" json:"MIN_GENESIS_ACTIVE_VALIDATOR_COUNT"`
	MIN_GENESIS_TIME                   Timestamp  `yaml:"MIN_GENESIS_TIME" json:"MIN_GENESIS_TIME"`
	GENESIS_FORK_VERSION               Version    `yaml:"GENESIS_FORK_VERSION" json:"GENESIS_FORK_VERSION"`
	GENESIS_DELAY                      Timestamp  `yaml:"GENESIS_DELAY" json:"GENESIS_DELAY"`

	// Altair
	ALTAIR_FORK_VERSION Version `yaml:"ALTAIR_FORK_VERSION" json:"ALTAIR_FORK_VERSION"`
	ALTAIR_FORK_EPOCH   Epoch   `yaml:"ALTAIR_FORK_EPOCH" json:"ALTAIR_FORK_EPOCH"`

	// Bellatrix
	BELLATRIX_FORK_VERSION Version `yaml:"BELLATRIX_FORK_VERSION" json:"BELLATRIX_FORK_VERSION"`
	BELLATRIX_FORK_EPOCH   Epoch   `yaml:"BELLATRIX_FORK_EPOCH" json:"BELLATRIX_FORK_EPOCH"`

	// Capella
	CAPELLA_FORK_VERSION Version `yaml:"CAPELLA_FORK_VERSION" json:"CAPELLA_FORK_VERSION"`
	CAPELLA_FORK_EPOCH   Epoch   `yaml:"CAPELLA_FORK_EPOCH" json:"CAPELLA_FORK_EPOCH"`

	// Deneb
	DENEB_FORK_VERSION Version `yaml:"DENEB_FORK_VERSION" json:"DENEB_FORK_VERSION"`
	DENEB_FORK_EPOCH   Epoch   `yaml:"DENEB_FORK_EPOCH" json:"DENEB_FORK_EPOCH"`

	// Electra
	ELECTRA_FORK_VERSION Version `yaml:"ELECTRA_FORK_VERSION" json:"ELECTRA_FORK_VERSION"`
	ELECTRA_FORK_EPOCH   Epoch   `yaml:"ELECTRA_FORK_EPOCH" json:"ELECTRA_FORK_EPOCH"`

	// EIP6110
	EIP6110_FORK_VERSION Version `yaml:"EIP6110_FORK_VERSION" json:"EIP6110_FORK_VERSION"`
	EIP6110_FORK_EPOCH   Epoch   `yaml:"EIP6110_FORK_EPOCH" json:"EIP6110_FORK_EPOCH"`

	// EIP7002
	EIP7002_FORK_VERSION Version `yaml:"EIP7002_FORK_VERSION" json:"EIP7002_FORK_VERSION"`
	EIP7002_FORK_EPOCH   Epoch   `yaml:"EIP7002_FORK_EPOCH" json:"EIP7002_FORK_EPOCH"`

	// WHISK
	WHISK_FORK_VERSION Version `yaml:"WHISK_FORK_VERSION" json:"WHISK_FORK_VERSION"`
	WHISK_FORK_EPOCH   Epoch   `yaml:"WHISK_FORK_EPOCH" json:"WHISK_FORK_EPOCH"`

	// Time parameters
	SECONDS_PER_SLOT                    Timestamp  `yaml:"SECONDS_PER_SLOT" json:"SECONDS_PER_SLOT"`
	SECONDS_PER_ETH1_BLOCK              Uint64View `yaml:"SECONDS_PER_ETH1_BLOCK" json:"SECONDS_PER_ETH1_BLOCK"`
	MIN_VALIDATOR_WITHDRAWABILITY_DELAY Epoch      `yaml:"MIN_VALIDATOR_WITHDRAWABILITY_DELAY" json:"MIN_VALIDATOR_WITHDRAWABILITY_DELAY"`
	SHARD_COMMITTEE_PERIOD              Epoch      `yaml:"SHARD_COMMITTEE_PERIOD" json:"SHARD_COMMITTEE_PERIOD"`

	// Validator cycle
	INACTIVITY_SCORE_BIAS          Uint64View `yaml:"INACTIVITY_SCORE_BIAS" json:"INACTIVITY_SCORE_BIAS"`
	INACTIVITY_SCORE_RECOVERY_RATE Uint64View `yaml:"INACTIVITY_SCORE_RECOVERY_RATE" json:"INACTIVITY_SCORE_RECOVERY_RATE"`
	EJECTION_BALANCE               Gwei       `yaml:"EJECTION_BALANCE" json:"EJECTION_BALANCE"`
	MIN_PER_EPOCH_CHURN_LIMIT      Uint64View `yaml:"MIN_PER_EPOCH_CHURN_LIMIT" json:"MIN_PER_EPOCH_CHURN_LIMIT"`
	CHURN_LIMIT_QUOTIENT           Uint64View `yaml:"CHURN_LIMIT_QUOTIENT" json:"CHURN_LIMIT_QUOTIENT"`
	// New in Deneb:EIP7514
	MAX_PER_EPOCH_ACTIVATION_CHURN_LIMIT Uint64View `yaml:"MAX_PER_EPOCH_ACTIVATION_CHURN_LIMIT" json:"MAX_PER_EPOCH_ACTIVATION_CHURN_LIMIT"`

	// Fork choice
	PROPOSER_SCORE_BOOST                Uint64View `yaml:"PROPOSER_SCORE_BOOST" json:"PROPOSER_SCORE_BOOST"`
	REORG_HEAD_WEIGHT_THRESHOLD         Uint64View `yaml:"REORG_HEAD_WEIGHT_THRESHOLD" json:"REORG_HEAD_WEIGHT_THRESHOLD"`
	REORG_PARENT_WEIGHT_THRESHOLD       Uint64View `yaml:"REORG_PARENT_WEIGHT_THRESHOLD" json:"REORG_PARENT_WEIGHT_THRESHOLD"`
	REORG_MAX_EPOCHS_SINCE_FINALIZATION Epoch      `yaml:"REORG_MAX_EPOCHS_SINCE_FINALIZATION" json:"REORG_MAX_EPOCHS_SINCE_FINALIZATION"`

	// Deposit contract
	DEPOSIT_CHAIN_ID   Uint64View `yaml:"DEPOSIT_CHAIN_ID" json:"DEPOSIT_CHAIN_ID"`
	DEPOSIT_NETWORK_ID Uint64View `yaml:"DEPOSIT_NETWORK_ID" json:"DEPOSIT_NETWORK_ID"`

	// Networking
	GOSSIP_MAX_SIZE                    Uint64View           `yaml:"GOSSIP_MAX_SIZE" json:"GOSSIP_MAX_SIZE"`
	MAX_REQUEST_BLOCKS                 Uint64View           `yaml:"MAX_REQUEST_BLOCKS" json:"MAX_REQUEST_BLOCKS"`
	EPOCHS_PER_SUBNET_SUBSCRIPTION     Uint64View           `yaml:"EPOCHS_PER_SUBNET_SUBSCRIPTION" json:"EPOCHS_PER_SUBNET_SUBSCRIPTION"`
	MIN_EPOCHS_FOR_BLOCK_REQUESTS      Uint64View           `yaml:"MIN_EPOCHS_FOR_BLOCK_REQUESTS" json:"MIN_EPOCHS_FOR_BLOCK_REQUESTS"`
	MAX_CHUNK_SIZE                     Uint64View           `yaml:"MAX_CHUNK_SIZE" json:"MAX_CHUNK_SIZE"`
	TTFB_TIMEOUT                       Uint64View           `yaml:"TTFB_TIMEOUT" json:"TTFB_TIMEOUT"`
	RESP_TIMEOUT                       Uint64View           `yaml:"RESP_TIMEOUT" json:"RESP_TIMEOUT"`
	ATTESTATION_PROPAGATION_SLOT_RANGE Uint64View           `yaml:"ATTESTATION_PROPAGATION_SLOT_RANGE" json:"ATTESTATION_PROPAGATION_SLOT_RANGE"`
	MAXIMUM_GOSSIP_CLOCK_DISPARITY     Uint64View           `yaml:"MAXIMUM_GOSSIP_CLOCK_DISPARITY" json:"MAXIMUM_GOSSIP_CLOCK_DISPARITY"`
	MESSAGE_DOMAIN_INVALID_SNAPPY      NetworkMessageDomain `yaml:"MESSAGE_DOMAIN_INVALID_SNAPPY" json:"MESSAGE_DOMAIN_INVALID_SNAPPY"`
	MESSAGE_DOMAIN_VALID_SNAPPY        NetworkMessageDomain `yaml:"MESSAGE_DOMAIN_VALID_SNAPPY" json:"MESSAGE_DOMAIN_VALID_SNAPPY"`
	SUBNETS_PER_NODE                   Uint64View           `yaml:"SUBNETS_PER_NODE" json:"SUBNETS_PER_NODE"`
	ATTESTATION_SUBNET_COUNT           Uint64View           `yaml:"ATTESTATION_SUBNET_COUNT" json:"ATTESTATION_SUBNET_COUNT"`
	ATTESTATION_SUBNET_EXTRA_BITS      Uint64View           `yaml:"ATTESTATION_SUBNET_EXTRA_BITS" json:"ATTESTATION_SUBNET_EXTRA_BITS"`
	ATTESTATION_SUBNET_PREFIX_BITS     Uint64View           `yaml:"ATTESTATION_SUBNET_PREFIX_BITS" json:"ATTESTATION_SUBNET_PREFIX_BITS"`

	// Deneb
	MAX_REQUEST_BLOCKS_DENEB              Uint64View `yaml:"MAX_REQUEST_BLOCKS_DENEB" json:"MAX_REQUEST_BLOCKS_DENEB"`
	MAX_REQUEST_BLOB_SIDECARS             Uint64View `yaml:"MAX_REQUEST_BLOB_SIDECARS" json:"MAX_REQUEST_BLOB_SIDECARS"`
	MIN_EPOCHS_FOR_BLOB_SIDECARS_REQUESTS Uint64View `yaml:"MIN_EPOCHS_FOR_BLOB_SIDECARS_REQUESTS" json:"MIN_EPOCHS_FOR_BLOB_SIDECARS_REQUESTS"`
	BLOB_SIDECAR_SUBNET_COUNT             Uint64View `yaml:"BLOB_SIDECAR_SUBNET_COUNT" json:"BLOB_SIDECAR_SUBNET_COUNT"`

	// Whish
	WHISK_EPOCHS_PER_SHUFFLING_PHASE Uint64View `yaml:"WHISK_EPOCHS_PER_SHUFFLING_PHASE" json:"WHISK_EPOCHS_PER_SHUFFLING_PHASE"`
	WHISK_PROPOSER_SELECTION_GAP     Uint64View `yaml:"WHISK_PROPOSER_SELECTION_GAP" json:"WHISK_PROPOSER_SELECTION_GAP"`

	// EIP7594
	EIP7594_FORK_VERSION Version `yaml:"EIP7594_FORK_VERSION" json:"EIP7594_FORK_VERSION"`
	EIP7594_FORK_EPOCH   Epoch   `yaml:"EIP7594_FORK_EPOCH" json:"EIP7594_FORK_EPOCH"`
}

type SpecObj interface {
	Deserialize(spec *Spec, dr *codec.DecodingReader) error
	Serialize(spec *Spec, w *codec.EncodingWriter) error
	ByteLength(spec *Spec) uint64
	HashTreeRoot(spec *Spec, h tree.HashFn) Root
	FixedLength(spec *Spec) uint64
}

type SSZObj interface {
	codec.Serializable
	codec.Deserializable
	codec.FixedLength
	tree.HTR
}

type WrappedSpecObj interface {
	SSZObj
	Unwrap() (*Spec, SpecObj)
}

type specObj struct {
	spec *Spec
	des  SpecObj
}

func (s *specObj) Deserialize(dr *codec.DecodingReader) error {
	return s.des.Deserialize(s.spec, dr)
}

func (s *specObj) Serialize(w *codec.EncodingWriter) error {
	return s.des.Serialize(s.spec, w)
}

func (s *specObj) ByteLength() uint64 {
	return s.des.ByteLength(s.spec)
}

func (s *specObj) HashTreeRoot(h tree.HashFn) Root {
	return s.des.HashTreeRoot(s.spec, h)
}

func (s *specObj) FixedLength() uint64 {
	return s.des.FixedLength(s.spec)
}

func (s *specObj) Unwrap() (*Spec, SpecObj) {
	return s.spec, s.des
}

func (s *specObj) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, s.des)
}

func (s *specObj) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.des)
}

func (s *specObj) UnmarshalYAML(value *yaml.Node) error {
	return value.Decode(s.des)
}

func (s *specObj) MarshalYAML() (interface{}, error) {
	return s.des, nil
}

type Spec struct {
	Phase0Preset    `json:",inline" yaml:",inline"`
	AltairPreset    `json:",inline" yaml:",inline"`
	BellatrixPreset `json:",inline" yaml:",inline"`
	CapellaPreset   `json:",inline" yaml:",inline"`
	DenebPreset     `json:",inline" yaml:",inline"`
	ElectraPreset   `json:",inline" yaml:",inline"`
	Config          `json:",inline" yaml:",inline"`

	ExecutionEngine `json:"-" yaml:"-"`
}

// Wraps the object to parametrize with given spec. JSON and YAML functionality is proxied to the inner value.
func (spec *Spec) Wrap(des SpecObj) SSZObj {
	return &specObj{spec, des}
}

func (spec *Spec) ForkVersion(slot Slot) Version {
	epoch := spec.SlotToEpoch(slot)
	if epoch < spec.ALTAIR_FORK_EPOCH {
		return spec.GENESIS_FORK_VERSION
	} else if epoch < spec.BELLATRIX_FORK_EPOCH {
		return spec.ALTAIR_FORK_VERSION
	} else if epoch < spec.CAPELLA_FORK_EPOCH {
		return spec.BELLATRIX_FORK_VERSION
	} else {
		return spec.CAPELLA_FORK_VERSION
	}
}
