package kubectl

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/AleckDarcy/reload/core/log"
)

func ScaleDeployment(name, namespace string, replicas int) bool {
	pods := GetPodsByName(name)
	if len(pods.Pods) == 0 && replicas == 0 {
		log.Logf("no pods available named: %s", name)
	} else if len(pods.Pods) == replicas {
		return true
	}

	err := exec.Command("kubectl", "scale", "deployment", name, fmt.Sprintf("--replicas=%d", replicas), "-n", namespace).Run()
	if err != nil {
		log.Logf("scale deployment %s --replicas=%d -n %s err: %v", name, replicas, namespace, err)
	}
	retry := 0
	for ; ; retry++ {
		time.Sleep(time.Second)
		if pods := GetPodsByName(name); len(pods.Pods) == replicas {
			ready := true
			for _, pod := range pods.Pods {
				if !pod.Ready() {
					ready = false
					break
				}
			}

			if ready {
				return true
			}
		}
	}

	log.Logf("retried for 60s, fail")

	return false
}
