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

type RequestFault struct {
	source string
	target string
	request string
}

type Interpreter struct {
	rounds int
	crashes int
	goalDAG *dag.Graph
}

func (interpreter *Interpreter) handleTraceProcessing(resp *data.Response) (map[string][]*tracer.Record, map[string]string) {
	recordMap := make(map[string][]*tracer.Record)
	uuid2Request := make(map[string]string)

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

		// Create uuid mapping to the request name
		if _, ok := recordMap[currRecord.Uuid]; !ok {
			uuid2Request[currRecord.Uuid] = currRecord.MessageName
		}

		recordMap[currRecord.Uuid] = append(recordMap[currRecord.Uuid], currRecord)
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

	return recordMap, uuid2Request
}

func (interpreter *Interpreter) createDAG(data map[string][]*tracer.Record, uuid2Request map[string]string) (*dag.Graph, []*RequestFault) {
	faults := make([]*RequestFault, 0)
	/*
		1. Create new DAG
	*/
	nodeNeighbors := make(map[string][]string)
	for _, val := range data {
		node := val[0]
		neighbor := val[1]

		nodeNeighbors[node.Service] = append(nodeNeighbors[node.Service], neighbor.Service)
		fault := &RequestFault{source: node.Service, target: neighbor.Service, request: uuid2Request[node.Uuid]}
		faults = append(faults, fault)

	}

	d :=  &dag.Graph{}

	uniqueKeys := make(map[string]int)
	for key, val := range nodeNeighbors {
		if _, ok := uniqueKeys[key]; ok {
			// key exists so do nothing
		} else {
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

	return d, faults
}

/*
	Method used to take edges from a trace and produce their corresponding
	Fault Injections (Fault Crashes)
 */
func (interpreter *Interpreter) edgesToFaults(d *dag.Graph, faults []*RequestFault) []*tracer.TFI {
	tfis := make([]*tracer.TFI ,0)
	edges := d.Edges()

	/*
			1. We have our DAG, and list of possible crash faults
		       Now we need to create the corresponding TFI's
	*/
	for _, fault := range faults {
		tfi := &tracer.TFI{
			Type: tracer.FaultType_FaultCrash,
			Name: []string{fault.request},
		}

		tfis = append(tfis, tfi)
	}

	return tfis
}


func (interpreter *Interpreter) forwardStep(reqs *data.Requests, resp *data.Response) *data.Requests {
	/*
		1. Process Trace information
	 */
	recordMap, uuid2Request := interpreter.handleTraceProcessing(resp)

	/*
		2. Analyze processed data to create DAG
	 */
	d, faults := interpreter.createDAG(recordMap, uuid2Request)

	interpreter.goalDAG = d

	/*
		3. Determine possible faults, initially start with
		   support of crash faults
	 */
	tfis := interpreter.edgesToFaults(d, faults)

	/*
		4. Now we need to select a TFI for the list of possibilities
	       and return new Request suggestion
	 */

	return &data.Requests{}
}