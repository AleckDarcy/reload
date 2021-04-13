package main

import (
	"fmt"

	Dag "github.com/heimdalr/dag"
)

func main() {
	fmt.Println("Creating New DAG")
	d := Dag.NewDAG()

	v1, _ := d.AddVertex(1)
	v2, _ := d.AddVertex(2)
	v3, _ := d.AddVertex(struct{a string; b string}{a: "foo", b: "bar"})

	_ = d.AddEdge(v1, v2)
	_ = d.AddEdge(v1, v3)

	fmt.Print(d.String())
}
