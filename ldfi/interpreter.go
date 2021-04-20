package ldfi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/AleckDarcy/reload/core/client/data"
	"github.com/AleckDarcy/reload/core/tracer"
	"github.com/hashicorp/terraform/dag"
)

var service2Requests = map[string][]string{
	"cartservice": {"AddItemRequest", "EmptyCartRequest", "GetCartRequest"},
	"recommendationservice": {"ListRecommendationsRequest"},
	"productcatalogservice": {"ListProductsRequest", "GetProductRequest", "SearchProductsRequest"},
	"shippingservice": {"GetQuoteRequest", "ShipOrderRequest"},
	"currencyservice": {"GetSupportedCurrenciesRequest", "CurrencyConversionRequest"},
	"paymentservice": {"ChargeRequest"},
	"emailservice": {"EmailService", "SendOrderConfirmationRequest"},
	"checkoutservice": {"PlaceOrderRequest"},
	"adservice": {"AdRequest"},
}

type Interpreter struct {
	rounds int
	crashes int
	goalDAG *dag.Graph
}

func (interpreter *Interpreter) handleTraceProcessing(resp *data.Response) map[string][]*tracer.Record {
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

	return recordMap
}

func (interpreter *Interpreter) createDAG(data map[string][]*tracer.Record) *dag.Graph {
	/*
		1. Create new DAG
	*/
	nodeNeighbors := make(map[string][]string)
	for _, val := range data {
		node := val[0]
		neighbor := val[1]

		nodeNeighbors[node.Service] = append(nodeNeighbors[node.Service], neighbor.Service)
	}

	d :=  &dag.Graph{}

	uniqueKeys := make(map[string]int)
	for key, val := range nodeNeighbors {
		if _, ok := uniqueKeys[key]; ok {
			// key exists so do nothing
		} else {
			//v := dag.NewVertex(key, nil)
			uniqueKeys[key] = 1
			d.Add(key)
		}

		for _, neighbor := range val {
			if _, ok := uniqueKeys[neighbor]; ok {
				// key exists so do nothing
			} else {
				uniqueKeys[neighbor] = 1
				d.Add(neighbor)
			}
			//d.AddEdge(uniqueKeys[key], uniqueKeys[neighbor])
			// edge := dag.BasicEdge(uniqueKeys[key], uniqueKeys[neighbor])
			edge := dag.BasicEdge(key, neighbor)
			d.Connect(edge)
		}
	}

	return d
}

/*
	Method used to take edges from a trace and produce their corresponding
	Fault Injections (Fault Crashes)
 */
func (interpreter *Interpreter) edgesToFaults(d *dag.AcyclicGraph) []*tracer.Trace {
	/*
		1. We have our DAG, now we need to convert to possible
		crash faults that we can inject into the system
	 */

	return []*tracer.Trace{}
}


func (interpreter *Interpreter) forwardStep(reqs *data.Requests, resp *data.Response) *data.Requests {
	/*
		1. Process Trace information
	 */
	recordMap := interpreter.handleTraceProcessing(resp)

	/*
		2. Analyze processed data to create DAG
	 */
	d := interpreter.createDAG(recordMap)

	interpreter.goalDAG = d

	/*
		3. Feed DAG to SAT Solver to determine possible faults
	 */
	fmt.Println(d.String())

	return &data.Requests{}
}