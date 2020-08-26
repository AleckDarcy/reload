package kubectl

import (
	"testing"
)

func TestGetPod(t *testing.T) {
	//t.Log(AllReady())
	//t.Logf("%v", GetPods())
	//t.Logf("%v", GetPodsByName("frontend"))

	//t1 := time.Now()
	//ScaleDeployment("adservice", "default", 0)
	//t2 := time.Now()
	//ScaleDeployment("adservice", "default", 1)
	//t3 := time.Now()
	//
	//t.Logf("scale in time: %v", t2.Sub(t1))
	//t.Logf("scale out time: %v", t3.Sub(t2))

	ScaleDeployment("productcatalogservice", "default", 0)
	ScaleDeployment("productcatalogservice", "default", 1)
}
