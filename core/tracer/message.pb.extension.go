package tracer

import (
	"time"

	"github.com/AleckDarcy/reload/core/errors"
)

func (m *Trace) Copy() *Trace {
	newM := *m

	return &newM
}

func (m *Trace) RLFI() error {
	if rlfi := m.Rlfi; rlfi != nil {
		lastIndex := len(m.Records) - 1

		if rlfi.Name == m.Records[lastIndex].MessageName {
			if rlfi.Type == FaultType_FaultCrash {
				return errors.ErrorFI_RLFI_Crash
			} else if rlfi.Type == FaultType_FaultDelay {
				time.Sleep(time.Duration(rlfi.Delay) * time.Millisecond)

				return errors.ErrorFI_RLFI_Delay
			} else {
				return errors.ErrorFI_RLFI_
			}
		}
	}

	return nil
}

func (m *Trace) TFI() error {
	if tfi := m.Tfi; tfi != nil {
		lastIndex := len(m.Records) - 1

		if tfi.Name == m.Records[lastIndex].MessageName {
			count := 0
			after := tfi.After

			if lastIndex < len(after) {
				return nil
			}

			for i, record := range m.Records {
				if i == lastIndex {
					break
				}

				if count == len(after) {
					if tfi.Type == FaultType_FaultCrash {
						return errors.ErrorFI_TFI_Crash
					} else if tfi.Type == FaultType_FaultDelay {
						time.Sleep(time.Duration(tfi.Delay) * time.Millisecond)

						return errors.ErrorFI_TFI_Delay
					} else {
						return errors.ErrorFI_TFI_
					}
				}

				if record.Type == RecordType_RecordReceive && record.MessageName == after[count] {
					count++
				}
			}
		}
	}

	return nil
}

func (m *Trace) AppendRecord(record *Record) *Trace {
	m.Records = append(m.Records, record)
	m.Depth = int64(len(m.Records))

	return m
}
