package kubectl

import (
	"bytes"
	"io"
	"os/exec"
	"strconv"
	"strings"

	"github.com/AleckDarcy/reload/core/log"
)

func ParsePod(line string) *Pod {
	attrs := strings.Fields(line)
	if len(attrs) != 5 {
		return nil
	}

	instances := strings.Split(attrs[1], "/")
	if len(instances) != 2 {
		return nil
	}

	ready, _ := strconv.ParseInt(instances[0], 10, 64)
	total, _ := strconv.ParseInt(instances[1], 10, 64)
	restarts, _ := strconv.ParseInt(attrs[3], 10, 64)

	return &Pod{
		Type:     strings.Split(attrs[0], "-")[0],
		Name:     attrs[0],
		Instance: Instance{Ready: int(ready), Total: int(total)},
		Status:   attrs[2],
		Restarts: int(restarts),
		Age:      attrs[4],
	}
}

func GetPods() *Pods {
	buffer := &bytes.Buffer{}
	c := exec.Command("kubectl", "get", "pods")
	c.Stdout = buffer
	c.Run()

	pods := &Pods{}
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
		}

		pods.Pods = append(pods.Pods, pod)
	}

	return pods
}

func GetCrashLoopBackOff() *Pods {
	res := &Pods{}
	pods := GetPods()

	for i := range pods.Pods {
		pod := pods.Pods[i]
		if pod.Status == "CrashLoopBackOff" {
			res.Pods = append(res.Pods, pod)
		}
	}

	return res
}

func GetPodsByName(name string) *Pods {
	c1 := exec.Command("kubectl", "get", "pods")
	c2 := exec.Command("grep", name)

	c2.Stdin, c1.Stdout = io.Pipe()

	buffer := &bytes.Buffer{}
	c2.Stdout = buffer
	c1.Start()
	c2.Start()
	c1.Wait()
	c1.Stdout.(*io.PipeWriter).Close()
	c2.Wait()

	pods := &Pods{}
	for {
		line, err := buffer.ReadBytes('\n')
		if err != nil {
			break
		}

		pod := ParsePod(string(line))
		if pod == nil {
			log.Logf("parse line: \"%s\" fail", line)
		}

		pods.Pods = append(pods.Pods, pod)
	}

	return pods
}
