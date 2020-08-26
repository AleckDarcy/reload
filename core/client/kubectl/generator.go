package kubectl

type Generator struct {
	names []string
	ids   []int

	combineLen int
}

func NewGenerator(names []string) *Generator {
	ids := make([]int, len(names))
	for i := 0; i < len(names); i++ {
		ids[i] = i
	}

	return &Generator{
		names: names,
		ids:   ids,
	}
}

func (c *Generator) SetLen(length int) {
	c.combineLen = length
}

func (c *Generator) GenerateCombinartorial(resultChan chan []int) {
	arrLen := len(c.ids) - c.combineLen
	for i := 0; i <= arrLen; i++ {
		result := make([]int, c.combineLen)
		result[0] = c.ids[i]
		if c.combineLen == 1 {
			resultChan <- result
		} else {
			c.doProcess(resultChan, result, i, 1)
		}
	}
}

func (c *Generator) doProcess(resultChan chan []int, result []int, rawIndex int, curIndex int) {
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
			tmp := make([]int, len(tResult))
			copy(tmp, tResult)
			resultChan <- tmp
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
