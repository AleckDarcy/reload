package public

import "time"

const ENV_PROFILE_INTERVAL = 5 * time.Second
const CPU_PROFILE_DURATION = time.Second
const CPU_PROFILE_DURATION_MAX = 2 * CPU_PROFILE_DURATION

const BUS_OBSERVATION_QUEUE_INTERVAL = time.Second
const EventMetadata_Timeout = 5 * time.Second

const TIME_FORMAT_DEFAULT = time.RFC3339
const TIME_FORMAT_RFC3339 = time.RFC3339
const TIME_FORMAT_RFC3339Nano = time.RFC3339Nano

const PREREQUISITE_ACCOMPLISHED = -1
const CONFIGURE_ID_DEFAULT = -1
