package tracer

import "fmt"

type NameMeta struct {
	Service string
	Name    string
}

type Generator struct {
	names           []*NameMeta
	ids             []int
	faults          []*TFI
	serviceCount    map[string]int64
	request2Service map[string]string

	combineLen int
}

type FIType int64

const (
	FI_TFI FIType = iota
	FI_RLFI
)

type Faults struct {
	FIType FIType
	IDs    []int
	Faults []*TFI
}

func NewGenerator(nameLists []*NameMeta) *Generator {
	ids := make([]int, len(nameLists))
	for i := 0; i < len(nameLists); i++ {
		ids[i] = i
	}

	serviceCount := map[string]int64{}
	request2Service := map[string]string{}
	for _, names := range nameLists {
		serviceCount[names.Service] = serviceCount[names.Service] + 1
		request2Service[names.Name] = names.Service
	}

	faults := make([]*TFI, len(nameLists))
	serviceNameCount := map[string]map[string]int64{}
	for i, names := range nameLists {
		tfi := &TFI{
			Type:  FaultType_FaultCrash,
			Name:  []string{names.Name},
			After: []*TFIMeta{},
		}

		nameCount, ok := serviceNameCount[names.Service]
		if !ok {
			nameCount = map[string]int64{}
			serviceNameCount[names.Service] = nameCount
		}

		for name, count := range nameCount {
			tfi.After = append(tfi.After, &TFIMeta{Name: name, Times: count})
		}

		nameCount[names.Name] = nameCount[names.Name] + 1

		faults[i] = tfi

		fmt.Println(i, tfi)
	}

	fmt.Println(serviceCount)

	return &Generator{
		names:           nameLists,
		ids:             ids,
		faults:          faults,
		serviceCount:    serviceCount,
		request2Service: request2Service,
	}
}

func (c *Generator) SetLen(length int) {
	c.combineLen = length
}

func (c *Generator) GenerateCombinartorial(resultChan chan *Faults) {
	arrLen := len(c.ids) - c.combineLen
	for i := 0; i <= arrLen; i++ {
		result := make([]int, c.combineLen)
		result[0] = c.ids[i]
		if c.combineLen == 1 {
			resultChan <- c.GetFaults(result)
		} else {
			c.doProcess(resultChan, result, i, 1)
		}
	}
}

func (c *Generator) doProcess(resultChan chan *Faults, result []int, rawIndex int, curIndex int) {
	var choice = len(c.ids) - rawIndex + curIndex - c.combineLen
	//fmt.Printf("Choice: %d, rawLen: %d, rawIndex : %d, curIndex : %d \r\n", choice, rawLen, rawIndex, curIndex)
	var tResult []int
	for i := 0; i < choice; i++ {
		if i != 0 {
			tResult := make([]int, c.combineLen)
			copyArr(result, tResult)
		} else {
			tResult = result
		}
		//fmt.Println(curIndex)
		tResult[curIndex] = c.ids[i+1+rawIndex]

		if curIndex+1 == c.combineLen {
			resultChan <- c.GetFaults(tResult)
			continue
		} else {
			c.doProcess(resultChan, tResult, rawIndex+i+1, curIndex+1)
		}

	}
}

func copyArr(rawArr []int, target []int) {
	for i := 0; i < len(rawArr); i++ {
		target[i] = rawArr[i]
	}
}

func (c *Generator) GetFaults(ids []int) *Faults {
	arr := make([]int, len(ids))
	copy(arr, ids)

	faults := make([]*TFI, 0, len(arr))

	tmpServiceCount := map[string]int64{}
	serviceNameCount := map[string]map[string]int64{}

	for _, i := range arr {
		fault := c.faults[i]

		service := c.request2Service[fault.Name[0]]
		tmpServiceCount[service] = tmpServiceCount[service] + 1
		nameCount, ok := serviceNameCount[service]
		if !ok {
			nameCount = map[string]int64{}
			serviceNameCount[service] = nameCount
		}

		nameCount[fault.Name[0]] = nameCount[fault.Name[0]] + 1
	}

	serviceRIFI := map[string]struct{}{}

	fiType := FI_RLFI
	for _, i := range arr {
		fault := c.faults[i]

		service := c.request2Service[fault.Name[0]]
		if tmpServiceCount[service] != c.serviceCount[service] {
			fiType = FI_TFI
			faults = append(faults, c.faults[i])
		} else {
			if _, ok := serviceRIFI[service]; !ok {
				for name := range serviceNameCount[service] {
					faults = append(faults, &TFI{
						Type: FaultType_FaultCrash,
						Name: []string{name},
					})
				}

				serviceRIFI[service] = struct{}{}
			}
		}
	}

	return &Faults{FIType: fiType, IDs: arr, Faults: faults}
}
