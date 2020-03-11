package tracer

import (
	"fmt"
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

	return m
}

func (m *Trace) JSONString() string {
	strRecords, strRlfi, strTfi := "null", "null", "null"

	if len(m.Records) != 0 {
		for i, record := range m.Records {
			if i == 0 {
				strRecords = fmt.Sprintf(`
		{"Type": "%v", "Timestamp"": "%s", "MessageName"": "%s"}`,
					record.Type.String(), time.Unix(record.Timestamp/1e9, record.Timestamp%1e9), record.MessageName)
			} else {
				strRecords += fmt.Sprintf(`,
		{"Type": "%v", "Timestamp"": "%s", "MessageName"": "%s"}`,
					record.Type.String(), time.Unix(record.Timestamp/1e9, record.Timestamp%1e9), record.MessageName)
			}
		}
	}

	if m.Rlfi != nil {
		strRlfi = fmt.Sprintf(`{
		"Type": "%v",
		"Name": "%s",
		"Delay": %d
	}`,
			m.Rlfi.Type, m.Rlfi.Name, m.Rlfi.Delay,
		)
	}

	if m.Tfi != nil {
		strAfter := "null"
		if len(m.Tfi.After) != 0 {
			for i, after := range m.Tfi.After {
				if i == 0 {
					strAfter = fmt.Sprintf(`
			"%s"`,
						after,
					)
				} else {
					strAfter = fmt.Sprintf(`,
			"%s"`,
						after,
					)
				}
			}
		}

		strTfi = fmt.Sprintf(`{
		"Type": "%v",
		"Name": "%s",
		"Delay": %d,
		"After": [
			%s
		]
	}`,
			m.Rlfi.Type, m.Rlfi.Name, m.Rlfi.Delay, strAfter,
		)
	}

	return fmt.Sprintf(`
{
	"ID": %d,
	"Records": [
		%s
	],
	"Rlfi": %s,
	"Tfi": %s
}
`,
		m.Id, strRecords, strRlfi, strTfi,
	)
}
