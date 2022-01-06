package configs

import (
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
		SAFE_SLOTS_TO_UPDATE_JUSTIFIED:   8,
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
		EPOCHS_PER_SLASHINGS_VECTOR:      1 << 13,
		HISTORICAL_ROOTS_LIMIT:           1 << 24,
		VALIDATOR_REGISTRY_LIMIT:         1 << 40,
		BASE_REWARD_FACTOR:               64,
		PROPORTIONAL_SLASHING_MULTIPLIER: 1,
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
		PROPORTIONAL_SLASHING_MULTIPLIER_ALTAIR: 2,
		SYNC_COMMITTEE_SIZE:                     512,
		EPOCHS_PER_SYNC_COMMITTEE_PERIOD:        256,
		MIN_SYNC_COMMITTEE_PARTICIPANTS:         1,
	},
	MergePreset: common.MergePreset{
		INACTIVITY_PENALTY_QUOTIENT_MERGE:      16777216,
		MIN_SLASHING_PENALTY_QUOTIENT_MERGE:    32,
		PROPORTIONAL_SLASHING_MULTIPLIER_MERGE: 3,
		MAX_BYTES_PER_TRANSACTION:              1073741824,
		MAX_TRANSACTIONS_PER_PAYLOAD:           1048576,
		BYTES_PER_LOGS_BLOOM:                   256,
		MAX_EXTRA_DATA_BYTES:                   32,
	},
	ShardingPreset: common.ShardingPreset{
		MAX_SHARDS:                          1024,
		INITIAL_ACTIVE_SHARDS:               64,
		SAMPLE_PRICE_ADJUSTMENT_COEFFICIENT: 8,
		MAX_SHARD_PROPOSER_SLASHINGS:        16,
		MAX_SHARD_HEADERS_PER_SHARD:         4,
		SHARD_STATE_MEMORY_SLOTS:            256,
		BLOB_BUILDER_REGISTRY_LIMIT:         1 << 40,
		MAX_SAMPLES_PER_BLOCK:               2048,
		TARGET_SAMPLES_PER_BLOCK:            1024,
		MAX_SAMPLE_PRICE:                    1 << 33,
		MIN_SAMPLE_PRICE:                    8,
	},
	Config: common.Config{
		PRESET_BASE:                          "mainnet",
		MIN_GENESIS_ACTIVE_VALIDATOR_COUNT:   1 << 14,
		MIN_GENESIS_TIME:                     1606824000,
		GENESIS_FORK_VERSION:                 common.Version{0x00, 0x00, 0x00, 0x00},
		GENESIS_DELAY:                        604800,
		ALTAIR_FORK_VERSION:                  common.Version{0x01, 0x00, 0x00, 0x00},
		ALTAIR_FORK_EPOCH:                    common.Epoch(74240),
		MERGE_FORK_VERSION:                   common.Version{0x02, 0x00, 0x00, 0x00},
		MERGE_FORK_EPOCH:                     ^common.Epoch(0),
		SHARDING_FORK_VERSION:                common.Version{0x03, 0x00, 0x00, 0x00},
		SHARDING_FORK_EPOCH:                  ^common.Epoch(0),
		TERMINAL_TOTAL_DIFFICULTY:            common.MustU256("115792089237316195423570985008687907853269984665640564039457584007913129638912"),
		TERMINAL_BLOCK_HASH:                  common.Bytes32{},
		TERMINAL_BLOCK_HASH_ACTIVATION_EPOCH: ^common.Timestamp(0),
		SECONDS_PER_SLOT:                     12,
		SECONDS_PER_ETH1_BLOCK:               14,
		MIN_VALIDATOR_WITHDRAWABILITY_DELAY:  256,
		SHARD_COMMITTEE_PERIOD:               256,
		ETH1_FOLLOW_DISTANCE:                 2048,
		INACTIVITY_SCORE_BIAS:                4,
		INACTIVITY_SCORE_RECOVERY_RATE:       16,
		EJECTION_BALANCE:                     16_000_000_000,
		MIN_PER_EPOCH_CHURN_LIMIT:            4,
		CHURN_LIMIT_QUOTIENT:                 1 << 16,
		DEPOSIT_CHAIN_ID:                     1,
		DEPOSIT_NETWORK_ID:                   1,
		DEPOSIT_CONTRACT_ADDRESS:             [20]byte{0x00, 0x00, 0x00, 0x00, 0x21, 0x9a, 0xb5, 0x40, 0x35, 0x6c, 0xBB, 0x83, 0x9C, 0xbe, 0x05, 0x30, 0x3d, 0x77, 0x05, 0xFa},
	},
	ExecutionEngine: nil,
}
