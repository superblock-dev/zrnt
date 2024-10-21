package configs

import (
	"github.com/protolambda/ztyp/view"

	"github.com/protolambda/zrnt/eth2/beacon/common"
)

var Mainnet = &common.Spec{
	Phase0Preset: common.Phase0Preset{
		MAX_COMMITTEES_PER_SLOT:          64,
		TARGET_COMMITTEE_SIZE:            128,
		MAX_VALIDATORS_PER_COMMITTEE:     2048,
		SHUFFLE_ROUND_COUNT:              90,
		HYSTERESIS_QUOTIENT:              4,
		HYSTERESIS_DOWNWARD_MULTIPLIER:   1,
		HYSTERESIS_UPWARD_MULTIPLIER:     5,
		MIN_DEPOSIT_AMOUNT:               1000_000_000,
		MAX_EFFECTIVE_BALANCE:            32_000_000_000,
		EFFECTIVE_BALANCE_INCREMENT:      1_000_000_000,
		MIN_ATTESTATION_INCLUSION_DELAY:  1,
		SLOTS_PER_EPOCH:                  32,
		MIN_SEED_LOOKAHEAD:               1,
		MAX_SEED_LOOKAHEAD:               4,
		EPOCHS_PER_ETH1_VOTING_PERIOD:    64,
		SLOTS_PER_HISTORICAL_ROOT:        8192,
		MIN_EPOCHS_TO_INACTIVITY_PENALTY: 4,
		EPOCHS_PER_HISTORICAL_VECTOR:     1 << 16,
		MIN_SLASHING_WITHDRAWABLE_DELAY:      1 << 13,
		HISTORICAL_ROOTS_LIMIT:           1 << 24,
		VALIDATOR_REGISTRY_LIMIT:         1 << 40,
		BASE_REWARD_FACTOR:               64,
		WHISTLEBLOWER_REWARD_QUOTIENT:    512,
		PROPOSER_REWARD_QUOTIENT:         8,
		INACTIVITY_PENALTY_QUOTIENT:      1 << 26,
		MIN_SLASHING_PENALTY_QUOTIENT:    128,
		MAX_PROPOSER_SLASHINGS:           16,
		MAX_ATTESTER_SLASHINGS:           2,
		MAX_ATTESTATIONS:                 128,
		MAX_DEPOSITS:                     16,
		MAX_VOLUNTARY_EXITS:              16,
	},
	AltairPreset: common.AltairPreset{
		INACTIVITY_PENALTY_QUOTIENT_ALTAIR:      3 * (1 << 24),
		MIN_SLASHING_PENALTY_QUOTIENT_ALTAIR:    64,
		SYNC_COMMITTEE_SIZE:                     512,
		EPOCHS_PER_SYNC_COMMITTEE_PERIOD:        256,
		MIN_SYNC_COMMITTEE_PARTICIPANTS:         1,
	},
	BellatrixPreset: common.BellatrixPreset{
		INACTIVITY_PENALTY_QUOTIENT_BELLATRIX:      16777216,
		MIN_SLASHING_PENALTY_QUOTIENT_BELLATRIX:    32,
		MAX_BYTES_PER_TRANSACTION:                  1073741824,
		MAX_TRANSACTIONS_PER_PAYLOAD:               1048576,
		BYTES_PER_LOGS_BLOOM:                       256,
		MAX_EXTRA_DATA_BYTES:                       32,
	},
	CapellaPreset: common.CapellaPreset{
		MAX_VALIDATORS_PER_WITHDRAWALS_SWEEP: 16384,
		MAX_BLS_TO_EXECUTION_CHANGES:         16,
		MAX_WITHDRAWALS_PER_PAYLOAD:          16,
	},
	DenebPreset: common.DenebPreset{
		FIELD_ELEMENTS_PER_BLOB:              4096,
		MAX_BLOB_COMMITMENTS_PER_BLOCK:       4096,
		MAX_BLOBS_PER_BLOCK:                  6,
		KZG_COMMITMENT_INCLUSION_PROOF_DEPTH: 17,
	},
	Config: common.Config{
		PRESET_BASE:                           "mainnet",
		CONFIG_NAME:                           "mainnet",
		TERMINAL_TOTAL_DIFFICULTY:             view.MustUint256("58750000000000000000000"),
		TERMINAL_BLOCK_HASH:                   common.Bytes32{},
		TERMINAL_BLOCK_HASH_ACTIVATION_EPOCH:  ^common.Epoch(0),
		MIN_GENESIS_ACTIVE_VALIDATOR_COUNT:    1 << 14,
		MIN_GENESIS_TIME:                      1606824000,
		GENESIS_FORK_VERSION:                  common.Version{0x00, 0x00, 0x00, 0x00},
		GENESIS_DELAY:                         604800,
		ALTAIR_FORK_VERSION:                   common.Version{0x01, 0x00, 0x00, 0x00},
		ALTAIR_FORK_EPOCH:                     common.Epoch(74240),
		BELLATRIX_FORK_VERSION:                common.Version{0x02, 0x00, 0x00, 0x00},
		BELLATRIX_FORK_EPOCH:                  common.Epoch(144896),
		CAPELLA_FORK_VERSION:                  common.Version{0x03, 0x00, 0x00, 0x00},
		CAPELLA_FORK_EPOCH:                    common.Epoch(194048),
		DENEB_FORK_VERSION:                    common.Version{0x04, 0x00, 0x00, 0x00},
		DENEB_FORK_EPOCH:                      common.Epoch(269568),
		EIP6110_FORK_VERSION:                  common.Version{0x05, 0x00, 0x00, 0x00},
		EIP6110_FORK_EPOCH:                    ^common.Epoch(0),
		EIP7002_FORK_VERSION:                  common.Version{0x05, 0x00, 0x00, 0x00},
		EIP7002_FORK_EPOCH:                    ^common.Epoch(0),
		WHISK_FORK_VERSION:                    common.Version{0x06, 0x00, 0x00, 0x00},
		WHISK_FORK_EPOCH:                      ^common.Epoch(0),
		SECONDS_PER_SLOT:                      12,
		SECONDS_PER_ETH1_BLOCK:                14,
		MIN_VALIDATOR_WITHDRAWABILITY_DELAY:   256,
		SHARD_COMMITTEE_PERIOD:                256,
		ETH1_FOLLOW_DISTANCE:                  2048,
		INACTIVITY_SCORE_BIAS:                 4,
		INACTIVITY_SCORE_RECOVERY_RATE:        16,
		EJECTION_BALANCE:                      16_000_000_000,
		MIN_PER_EPOCH_CHURN_LIMIT:             4,
		CHURN_LIMIT_QUOTIENT:                  1 << 16,
		MAX_PER_EPOCH_ACTIVATION_CHURN_LIMIT:  8,
		PROPOSER_SCORE_BOOST:                  40,
		REORG_HEAD_WEIGHT_THRESHOLD:           20,
		REORG_PARENT_WEIGHT_THRESHOLD:         160,
		REORG_MAX_EPOCHS_SINCE_FINALIZATION:   2,
		DEPOSIT_CHAIN_ID:                      1,
		DEPOSIT_NETWORK_ID:                    1,
		DEPOSIT_CONTRACT_ADDRESS:              [20]byte{0x00, 0x00, 0x00, 0x00, 0x21, 0x9a, 0xb5, 0x40, 0x35, 0x6c, 0xBB, 0x83, 0x9C, 0xbe, 0x05, 0x30, 0x3d, 0x77, 0x05, 0xFa},
		GOSSIP_MAX_SIZE:                       10 * (1 << 20),
		MAX_REQUEST_BLOCKS:                    1024,
		EPOCHS_PER_SUBNET_SUBSCRIPTION:        256,
		MIN_EPOCHS_FOR_BLOCK_REQUESTS:         33024,
		MAX_CHUNK_SIZE:                        10485760,
		TTFB_TIMEOUT:                          5,
		RESP_TIMEOUT:                          10,
		ATTESTATION_PROPAGATION_SLOT_RANGE:    32,
		MAXIMUM_GOSSIP_CLOCK_DISPARITY:        500,
		MESSAGE_DOMAIN_INVALID_SNAPPY:         common.NetworkMessageDomain{0, 0, 0, 0},
		MESSAGE_DOMAIN_VALID_SNAPPY:           common.NetworkMessageDomain{1, 0, 0, 0},
		SUBNETS_PER_NODE:                      2,
		ATTESTATION_SUBNET_COUNT:              64,
		ATTESTATION_SUBNET_EXTRA_BITS:         0,
		ATTESTATION_SUBNET_PREFIX_BITS:        6,
		MAX_REQUEST_BLOCKS_DENEB:              128,
		MAX_REQUEST_BLOB_SIDECARS:             768,
		MIN_EPOCHS_FOR_BLOB_SIDECARS_REQUESTS: 4096,
		BLOB_SIDECAR_SUBNET_COUNT:             6,
		WHISK_EPOCHS_PER_SHUFFLING_PHASE:      256,
		WHISK_PROPOSER_SELECTION_GAP:          2,
		EIP7594_FORK_VERSION:                  common.Version{6, 0, 0, 1},
		EIP7594_FORK_EPOCH:                    ^common.Epoch(0),
	},
	ExecutionEngine: nil,
}
