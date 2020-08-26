package kubectl

import (
	"bytes"
	"os/exec"

	"github.com/AleckDarcy/reload/core/log"
)

func AllReady() bool {
	buffer := &bytes.Buffer{}
	c := exec.Command("kubectl", "get", "pods")
	c.Stdout = buffer
	c.Run()

	ready := true
	for i := 0; ; i++ {
		line, err := buffer.ReadBytes('\n')
		if err != nil {
			break
		} else if i == 0 {
			continue
		}

		pod := ParsePod(string(line))
		if pod == nil {
			log.Logf("parse line: \"%s\" fail", line)
		} else if pod.Instance.Ready != pod.Instance.Total {
			ready = false
			//log.Logf("pod %s not ready, status: %+v", pod.Name, pod)
		}
	}

	return ready
}
