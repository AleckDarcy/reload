package ldfi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/AleckDarcy/reload/core/client/data"
	"github.com/AleckDarcy/reload/core/tracer"
)

type Interpreter struct {
	rounds int
}

func (i *Interpreter) handleTrace(reqs *data.Requests, resp *data.Response) *data.Requests {
	fmt.Println("HERE")
	recordMap := make(map[string][]*tracer.Record)
	/*
		1. Write trace to JSON File
	*/
	file, _ := json.MarshalIndent(resp, "", " ")
	_ = ioutil.WriteFile("test.json", file, 0644)

	/*
		2. Analyze trace in order to create DAG
	*/
	for i := 0; i < len(resp.Trace.Records); i++ {
		currRecord := resp.Trace.Records[i]

		recordMap[currRecord.Uuid] = append(recordMap[currRecord.Uuid], currRecord)
		fmt.Println(currRecord)

	}

	/*
		3. Sort trace's based on timestamp
	*/
	for k, val := range recordMap {
		sort.Slice(val, func(i, j int) bool {
			return val[i].Timestamp <  val[j].Timestamp
		})

		recordMap[k] = val
	}

	/*
		4. We can prune out all requests that were not correctly handled
	 */
	for k, val := range recordMap {
		if len(val) < 4 {
			delete(recordMap, k)
		}
	}


	/*
		5. Prune out the responses from the UUid -> records mapping
		they are not needed for the formation of the dag
	 */
	for k, val := range recordMap {
		val = val[:2]
		recordMap[k] = val
	}

	/*
		6. Create new DAG
	 */
	nodeNeighbors := make(map[string][]string)
	for _, val := range recordMap {
		node := val[0]
		neighbor := val[1]

		nodeNeighbors[node.Service] = append(nodeNeighbors[node.Service], neighbor.Service)
	}

	//d :=  dag.NewDAG()

	uniqueKeys := make(map[string]int)
	for key, val := range nodeNeighbors {
		if _, ok := uniqueKeys[key]; ok {
			// key exists so do nothing
		} else {
			uniqueKeys[key] = 1
			//d.NewVertex(key)
		}

		for _, neighbor := range val {
			if _, ok := uniqueKeys[neighbor]; ok {
				// key exists so do nothing
			} else {
				uniqueKeys[neighbor] = 1

			}

		}
	}
	/*
	uniqueKeys := make(map[string]int)

	for _, val := range recordMap {

		for _, record := range val {
			if _, ok := uniqueKeys[record.Service]; ok {
				continue
			} else {
				uniqueKeys[record.Service] = 1
			}

		}
	}
	*/


	return &data.Requests{}
}