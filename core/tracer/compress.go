package tracer

import (
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"
)

var compressFlag = true

const pageSize = 4096
const pageCap = 256
const bit16Size = 2

var entrySize = unsafe.Sizeof(entry{})

type compress struct {
	pagePool *sync.Pool

	messagePool     *sync.Pool
	messageKey      sync.RWMutex
	messageNameMap  map[string]uint8
	messageNameList atomic.Value

	servicePool     *sync.Pool
	serviceKey      sync.RWMutex
	serviceNameMap  map[string]uint8
	serviceNameList atomic.Value
}

var Compress = &compress{
	pagePool: &sync.Pool{
		New: func() interface{} {
			bytes := &[pageSize]byte{}
			page := (*nodePage)(unsafe.Pointer(bytes))
			runtime.KeepAlive(bytes)

			return page
		},
	},
	messagePool: &sync.Pool{
		New: func() interface{} {
			return map[string]uint8{}
		},
	},
	servicePool: &sync.Pool{
		New: func() interface{} {
			return map[string]uint8{}
		},
	},

	messageNameMap: map[string]uint8{},
	serviceNameMap: map[string]uint8{},
}

func init() {
	Compress.messageNameList.Store([]string{})
	Compress.serviceNameList.Store([]string{})
}

type headerPage struct {
	traceID int64

	tfis     []*TFI
	uuidMap  map[string]uint16
	uuidList []string

	messageNameMap map[string]uint8
	serviceNameMap map[string]uint8

	pages []*page
}

type page struct {
	entryNum int
	nodePage *nodePage
}

type nodePage struct {
	entries [pageCap]entry
}

type entry struct {
	type_     uint32
	messageID byte
	serviceID byte
	uuidID    uint16
	timestamp int64
}

func (c *compress) Compress(t *Trace) *headerPage {
	//messageNameMap := c.messagePool.Get().(map[string]uint8)
	//serviceNameMap := c.servicePool.Get().(map[string]uint8)
	messageNameMap := map[string]uint8{}
	serviceNameMap := map[string]uint8{}
	pageNum := (len(t.Records)-1)/pageCap + 1
	h := &headerPage{
		uuidMap:        map[string]uint16{},
		messageNameMap: messageNameMap,
		serviceNameMap: serviceNameMap,
		pages:          make([]*page, 0, pageNum),
	}

	node := c.pagePool.Get().(*nodePage)
	p := &page{nodePage: node}

	ok := false
	for i, record := range t.Records {
		e := &node.entries[p.entryNum]

		e.type_ = uint32(RecordType_RecordSend)

		if e.messageID, ok = messageNameMap[record.MessageName]; !ok {
			c.messageKey.RLock()
			e.messageID, ok = c.messageNameMap[record.MessageName]
			c.messageKey.RUnlock()
			if !ok {
				c.messageKey.Lock()
				if _, ok = c.messageNameMap[record.MessageName]; !ok {
					e.messageID = uint8(len(c.messageNameMap))
					c.messageNameMap[record.MessageName] = e.messageID
					c.messageNameList.Store(append(c.messageNameList.Load().([]string), record.MessageName))
				}
				c.messageKey.Unlock()
			}

			messageNameMap[record.MessageName] = e.messageID
		}

		if e.serviceID, ok = serviceNameMap[record.Service]; !ok {
			c.serviceKey.RLock()
			e.serviceID, ok = c.serviceNameMap[record.Service]
			c.serviceKey.RUnlock()
			if !ok {
				c.serviceKey.Lock()
				if e.serviceID, ok = serviceNameMap[record.Service]; !ok {
					e.serviceID = uint8(len(c.serviceNameMap))
					c.serviceNameMap[record.Service] = e.serviceID
					c.serviceNameList.Store(append(c.serviceNameList.Load().([]string), record.Service))
				}
				c.serviceKey.Unlock()
			}

			serviceNameMap[record.Service] = e.messageID
		}

		if e.uuidID, ok = h.uuidMap[record.Uuid]; !ok {
			e.uuidID = uint16(len(h.uuidMap))
			h.uuidMap[record.Uuid] = e.uuidID
			h.uuidList = append(h.uuidList, record.Uuid)
		}

		e.timestamp = record.Timestamp

		if p.entryNum++; p.entryNum == pageCap {
			h.pages = append(h.pages, p)

			if i != len(t.Records)-1 {
				node = c.pagePool.Get().(*nodePage)
				p = &page{nodePage: node}
			}
		}
		//fmt.Printf("%+v\n", e)
	}

	return h
}

func (c *compress) Recycle(h *headerPage) {
	//c.messagePool.Put(h.messageNameMap)
	//c.servicePool.Put(h.serviceNameMap)

	for _, page := range h.pages {
		c.pagePool.Put(page.nodePage)
	}
}

func (c *compress) Decompress(h *headerPage) *Trace {
	entryNum := (len(h.pages)-1)*pageCap + h.pages[len(h.pages)-1].entryNum

	trace := &Trace{
		Id:      h.traceID,
		Records: make([]*Record, 0, entryNum),
	}

	for _, page := range h.pages {
		node := page.nodePage
		for j, e := range node.entries {
			if j == page.entryNum {
				break
			}

			trace.Records = append(trace.Records, &Record{
				Type:        RecordType(e.type_),
				Timestamp:   e.timestamp,
				MessageName: c.messageNameList.Load().([]string)[e.messageID],
				Uuid:        h.uuidList[e.uuidID],
				Service:     c.serviceNameList.Load().([]string)[e.serviceID],
			})
		}
	}

	return trace
}

type eachFunc func(i int, record *Record)

func (h *headerPage) IterateRecords(eachFunc eachFunc) {
	if len(h.pages) == 0 {
		return
	}

	messageNameList := Compress.messageNameList.Load().([]string)
	serviceNameList := Compress.serviceNameList.Load().([]string)

	i := 0
	record := &Record{}
	e := (*entry)(nil)

	for _, page := range h.pages {
		entryNum := page.entryNum
		node := page.nodePage
		for j := 0; j < entryNum; j++ {
			e = &node.entries[j]

			record.Type = RecordType(e.type_)
			record.MessageName = messageNameList[e.messageID]
			record.Service = serviceNameList[e.serviceID]
			record.Uuid = h.uuidList[e.uuidID]
			record.Timestamp = e.timestamp

			eachFunc(i, record)

			i++
		}
	}
}
