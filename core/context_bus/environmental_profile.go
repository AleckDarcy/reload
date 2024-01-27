package context_bus

import (
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"


	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"

	"fmt"
	"runtime"
	"sync"
	"time"
)

var MEMSTATS = &runtime.MemStats{}

type environmentProfiler struct {
	Latest *cb.EnvironmentalProfile
	sync.RWMutex
}

var EP = &environmentProfiler{
	Latest: &cb.EnvironmentalProfile{
		Hardware: &cb.HardwareProfile{},
	},
}

func (e *environmentProfiler) GetNetProfile() *cb.NetProfile {
	n, err := net.IOCounters(false)

	if err != nil {
		fmt.Println("NetProfile", err)

		return nil
	}

	return &cb.NetProfile{
		BytesSent:   n[0].BytesSent,
		BytesRecv:   n[0].BytesRecv,
		PacketsSent: n[0].PacketsSent,
		PacketsRecv: n[0].PacketsRecv,
		Errin:       n[0].Errin,
		Errout:      n[0].Errout,
		Dropin:      n[0].Dropin,
		Dropout:     n[0].Dropout,
	}
}

func (e *environmentProfiler) GetEnvironmentProfile() *cb.EnvironmentalProfile {
	ep := &cb.EnvironmentalProfile{
		Timestamp: time.Now().UnixNano(),
		Hardware:  &cb.HardwareProfile{},
	}

	signal := make(chan *cb.CPUProfile)
	go func(signal chan *cb.CPUProfile) {
		if c, err := cpu.Percent(CPU_PROFILE_DURATION, false); err == nil {
			signal <- &cb.CPUProfile{
				Percent: c[0],
			}
		} else {
			signal <- nil
		}
	}(signal)

	if m, err := mem.VirtualMemory(); err == nil {
		ep.Hardware.Mem = &cb.MemProfile{
			Total:       m.Total,
			Available:   m.Available,
			Used:        m.Used,
			UsedPercent: m.UsedPercent,
			Free:        m.Free,
		}
	}

	np := e.GetNetProfile()
	np_prev := e.Latest.Hardware.Net
	if np_prev == nil {
		np_prev = np
	} else if np != nil {
		ep.Hardware.Net = &cb.NetProfile{
			BytesSent:   np.BytesSent - np_prev.BytesSent,
			BytesRecv:   np.BytesRecv - np_prev.BytesRecv,
			PacketsSent: np.PacketsSent - np_prev.PacketsSent,
			PacketsRecv: np.PacketsRecv - np_prev.PacketsRecv,
			Errin:       np.Errin - np_prev.Errin,
			Errout:      np.Errout - np_prev.Errout,
			Dropin:      np.Dropin - np_prev.Dropin,
			Dropout:     np.Dropout - np_prev.Dropout,
		}

		np_prev = np
	}

	runtime.ReadMemStats(MEMSTATS)
	ep.Language = &cb.LanguageProfile{
		Type: cb.LanguageType_Golang,
		Profile: &cb.LanguageProfile_Go{
			Go: &cb.LanguageGo{
				HeapSys:       MEMSTATS.HeapSys,
				HeapAlloc:     MEMSTATS.HeapAlloc,
				HeapInuse:     MEMSTATS.HeapInuse,
				StackSys:      MEMSTATS.StackSys,
				StackInuse:    MEMSTATS.StackInuse,
				MSpanInuse:    MEMSTATS.MSpanInuse,
				MSpanSys:      MEMSTATS.MSpanSys,
				MCacheInuse:   MEMSTATS.MCacheInuse,
				MCacheSys:     MEMSTATS.MCacheSys,
				LastGC:        MEMSTATS.LastGC,
				NextGC:        MEMSTATS.NextGC,
				GCCPUFraction: MEMSTATS.GCCPUFraction,
			},
		},
	}

	select {
	case c := <-signal:
		ep.Hardware.Cpu = c
	case <-time.After(CPU_PROFILE_DURATION_MAX):
		// todo: report error
	}

	e.Lock()
	e.Latest.Next = ep.Timestamp
	ep.Prev = e.Latest.Timestamp
	e.Latest = ep
	e.Unlock()

	return ep
}
