package kubectl

import "encoding/json"

type Service struct {
	Name      string
	NameSpace string
	Replicas  int
}

func (s *Service) String() string {
	jsonBytes, _ := json.Marshal(s)

	return string(jsonBytes)
}

func GetServiceNames(services []*Service) string {
	result := ""

	for i, service := range services {
		if i == 0 {
			result = service.Name
		} else {
			result += ", " + service.Name
		}
	}

	return result
}

type Instance struct {
	Ready int
	Total int
}

type Pod struct {
	Type      string
	Name      string
	NameSpace string
	Instance  Instance
	Status    string
	Restarts  int
	Age       string
}

func (p *Pod) Ready() bool {
	return p.Instance.Total == p.Instance.Ready
}

type Pods struct {
	Pods []*Pod
}

func (p *Pods) String() string {
	jsonBytes, _ := json.Marshal(p)

	return string(jsonBytes)
}
