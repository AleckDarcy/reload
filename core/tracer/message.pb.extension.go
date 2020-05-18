package tracer

import (
	"fmt"
	"time"

	"github.com/AleckDarcy/reload/core/errors"
)

func (m *Trace) Copy() *Trace {
	newM := *m
	newM.Records = make([]*Record, len(m.Records))
	copy(newM.Records, m.Records)

	return &newM
}

func (m *Trace) CalFI(records []*Record) {
	for _, tfi := range m.Tfis {
		for _, record := range records {
			if record.Type == RecordType_RecordReceive {
				for _, after := range tfi.After {
					if record.MessageName == after.Name {
						after.Already++
					}
				}
			}
		}
	}
}

func (m *Trace) DoFI(name string) error {
	if err := m.RLFI(name); err != nil {
		return err
	} else if err = m.TFI(name); err != nil {
		return err
	}

	return nil
}

func (m *Trace) RLFI(name string) error {
	for _, rlfi := range m.Rlfis {
		if rlfi.Name == name {
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

func (m *Trace) TFI(name string) error {
	for _, tfi := range m.Tfis {
		for _, after := range tfi.After {
			if after.Name == name {
				after.Already++
				if after.Already != after.Times+1 {
					return nil
				}
			} else if after.Already < after.Times {
				return nil
			}
		}

		if tfi.Name[0] == name {
			if tfi.Type == FaultType_FaultCrash {
				return errors.ErrorFI_TFI_Crash
			} else if tfi.Type == FaultType_FaultDelay {
				time.Sleep(time.Duration(tfi.Delay) * time.Millisecond)

				return errors.ErrorFI_TFI_Delay
			} else {
				return errors.ErrorFI_TFI_
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
				strRecords = fmt.Sprintf(`{"type": %d, "timestamp": %d, "uuid": "%s", messageName": "%s"}`,
					record.Type, record.Timestamp, record.Uuid, record.MessageName)
			} else {
				strRecords += fmt.Sprintf(`,
		{"type": %d, "timestamp": %d, "uuid": "%s", "messageName": "%s"}`,
					record.Type, record.Timestamp, record.Uuid, record.MessageName)
			}
		}
	}

	if len(m.Rlfis) != 0 {
		for i, rlfi := range m.Rlfis {
			if i == 0 {
				strRlfi = fmt.Sprintf(`{
		"type": "%v",
		"name": "%s",
		"delay": %d
	}`,
					rlfi.Type, rlfi.Name, rlfi.Delay,
				)
			} else {
				strRlfi = fmt.Sprintf(`,
	{
		"type": "%v",
		"name": "%s",
		"delay": %d
	}`,
					rlfi.Type, rlfi.Name, rlfi.Delay,
				)
			}
		}
	}

	if len(m.Tfis) != 0 {
		for i, tfi := range m.Tfis {
			strAfter := "null"
			if len(tfi.After) != 0 {
				for i, after := range tfi.After {
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

			if i == 0 {
				strTfi = fmt.Sprintf(`{
		"type": "%v",
		"name": "%s",
		"delay": %d,
		"after": [
			%s
		]
	}`,
					tfi.Type, tfi.Name, tfi.Delay, strAfter,
				)
			} else {
				strTfi = fmt.Sprintf(`,
	{
		"type": "%v",
		"name": "%s",
		"delay": %d,
		"after": [
			%s
		]
	}`,
					tfi.Type, tfi.Name, tfi.Delay, strAfter,
				)
			}
		}
	}

	return fmt.Sprintf(`
{
	"id": %d,
	"records": [
		%s
	],
	"rlfi": [
		%s
	],
	"tfi": [
		%s
	]
}
`,
		m.Id, strRecords, strRlfi, strTfi,
	)
}
