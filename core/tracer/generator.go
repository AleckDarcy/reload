package tracer

type Generator struct {
	names  []string
	ids    []int
	faults []*TFI

	combineLen int
}

type Faults struct {
	IDs    []int
	Faults []*TFI
}

func NewGenerator(names []string) *Generator {
	ids := make([]int, len(names))
	for i := 0; i < len(names); i++ {
		ids[i] = i
	}

	faults := make([]*TFI, len(names))
	nameCount := map[string]int64{}
	for i, name := range names {
		if count, ok := nameCount[name]; ok {
			faults[i] = &TFI{
				Type:  FaultType_FaultCrash,
				Name:  name,
				After: []*TFIMeta{{Name: name, Times: count}},
			}
			nameCount[name] = count + 1
		} else {
			faults[i] = &TFI{
				Type:  FaultType_FaultCrash,
				Name:  name,
				After: []*TFIMeta{},
			}
			nameCount[name] = 1
		}
	}

	return &Generator{names: names, ids: ids, faults: faults}
}

func (c *Generator) SetLen(length int) {
	c.combineLen = length
}

func (c *Generator) Generate(resultChan chan *Faults) {
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

func (c *Generator) GetFaults(arr []int) *Faults {
	faults := make([]*TFI, 0, len(arr))
	for _, i := range arr {
		faults = append(faults, c.faults[i])
	}

	return &Faults{IDs: arr, Faults: faults}
}
