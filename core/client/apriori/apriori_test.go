package apriori

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"
)

var tfiNames = map[string]string{
	"0":  "Currency",
	"1":  "Product",
	"2":  "Cart",
	"3":  "Currency_TFI",
	"4":  "Currency_TFI",
	"5":  "Currency_TFI",
	"6":  "Currency_TFI",
	"7":  "Currency_TFI",
	"8":  "Currency_TFI",
	"9":  "Currency_TFI",
	"10": "Currency_TFI",
	"11": "Currency_TFI",
	"12": "Ad",
	"13": "Currency_TFI",
}

var rlfiNames = map[string]string{
	"0": "Currency",
	"1": "Currency",
	"2": "Product",
	"3": "Cart",
	"4": "Ad",
	"5": "5",
	"6": "6",
	"7": "7",
	"8": "8",
	"9": "9",
}

func helper(t *testing.T, path string, nameMap map[string]string) {
	f, _ := os.Open(path)
	r := bufio.NewReader(f)
	in := [][]string{}

	for {
		if line, _, err := r.ReadLine(); err != nil {
			break
		} else {
			str := string(line)
			idsStr := str[strings.Index(str, "[")+1 : strings.Index(str, "]")]
			ids := strings.Split(idsStr, " ")

			dup := false // Currency_TFI, Currency
			names := make([]string, 0, len(ids))
			for _, idStr := range ids {
				name := nameMap[idStr]
				if strings.Contains(name, "TFI") || name == "Currency" {
					if dup {
						continue
					}

					dup = true
				}

				names = append(names, name)
			}

			in = append(in, names)
		}
	}

	a := NewApriori(in)
	out := Sortable(a.Calculate(NewOptions(0.1, 0.5, 0.0, 0)))

	inMap := map[string]struct{}{}
	for _, names := range in {
		sort.Strings(names)
		inMap[fmt.Sprintf("%v", names)] = struct{}{}
	}

	//fmt.Println(inMap)
	sort.Sort(out)
	for i := len(out) - 1; i >= 0; i-- {
		//fmt.Println(fmt.Sprintf("%v", out[i].supportRecord.items))
		sort.Strings(out[i].supportRecord.items)
		if _, ok := inMap[fmt.Sprintf("%v", out[i].supportRecord.items)]; ok {
			t.Log(fmt.Sprintf("%+v", out[i]))
		}
		_ = i
		//if i == 10 {
		//	break
		//}
	}
}

type Sortable []RelationRecord

func (s Sortable) Len() int {
	return len(s)
}

func (s Sortable) Less(i, j int) bool {
	return s[i].supportRecord.support < s[j].supportRecord.support
}

func (s Sortable) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func TestName(t *testing.T) {
	t.Log("RLFI")
	helper(t, "chaos.log", rlfiNames)

	t.Log("TFI")
	helper(t, "tfi.log", tfiNames)
}

func TestApriori_Calculate(t *testing.T) {
	provider := []struct {
		in  [][]string
		out string
	}{
		{[][]string{
			{"beer", "nuts", "cheese"},
			{"beer", "nuts", "jam"},
			{"beer", "butter"},
			{"nuts", "cheese"},
			{"beer", "nuts", "cheese", "jam"},
			{"butter"},
			{"beer", "nuts", "jam", "butter"},
			{"jam"},
		},
			"[{{[beer] 0.625} [{[] [beer] 0.625 1}]} {{[jam] 0.5} [{[] [jam] 0.5 1}]} {{[nuts] 0.625} [{[] [nuts] 0.625 1}]} {{[beer butter] 0.25} [{[butter] [beer] 0.6666666666666666 1.0666666666666667}]} {{[beer cheese] 0.25} [{[cheese] [beer] 0.6666666666666666 1.0666666666666667}]} {{[beer jam] 0.375} [{[beer] [jam] 0.6 1.2} {[jam] [beer] 0.75 1.2}]} {{[beer nuts] 0.5} [{[beer] [nuts] 0.8 1.28} {[nuts] [beer] 0.8 1.28}]} {{[cheese nuts] 0.375} [{[cheese] [nuts] 1 1.6} {[nuts] [cheese] 0.6 1.5999999999999999}]} {{[jam nuts] 0.375} [{[jam] [nuts] 0.75 1.2} {[nuts] [jam] 0.6 1.2}]} {{[beer butter jam] 0.125} [{[beer butter] [jam] 0.5 1} {[butter jam] [beer] 1 1.6}]} {{[beer butter nuts] 0.125} [{[beer butter] [nuts] 0.5 0.8} {[butter nuts] [beer] 1 1.6}]} {{[beer cheese jam] 0.125} [{[beer cheese] [jam] 0.5 1} {[cheese jam] [beer] 1 1.6}]} {{[beer cheese nuts] 0.25} [{[beer cheese] [nuts] 1 1.6} {[beer nuts] [cheese] 0.5 1.3333333333333333} {[cheese nuts] [beer] 0.6666666666666666 1.0666666666666667}]} {{[beer jam nuts] 0.375} [{[beer jam] [nuts] 1 1.6} {[beer nuts] [jam] 0.75 1.5} {[jam nuts] [beer] 1 1.6}]} {{[butter jam nuts] 0.125} [{[butter jam] [nuts] 1 1.6} {[butter nuts] [jam] 1 2}]} {{[cheese jam nuts] 0.125} [{[cheese jam] [nuts] 1 1.6}]} {{[beer butter jam nuts] 0.125} [{[beer butter jam] [nuts] 1 1.6} {[beer butter nuts] [jam] 1 2} {[butter jam nuts] [beer] 1 1.6}]} {{[beer cheese jam nuts] 0.125} [{[beer cheese jam] [nuts] 1 1.6} {[beer cheese nuts] [jam] 0.5 1} {[cheese jam nuts] [beer] 1 1.6}]}]"},
		{[][]string{
			{"beer", "nuts", "cheese", "pizza", "cola", "lays"},
			{"beer", "nuts", "jam"},
			{"beer", "butter", "cola", "lays"},
			{"nuts", "cheese", "pizza", "cola"},
			{"beer", "nuts", "lays", "cheese", "jam"},
			{"butter"},
			{"beer", "nuts", "pizza", "cola", "jam", "butter"},
			{"jam"},
		},
			"[{{[beer] 0.625} [{[] [beer] 0.625 1}]} {{[cola] 0.5} [{[] [cola] 0.5 1}]} {{[jam] 0.5} [{[] [jam] 0.5 1}]} {{[nuts] 0.625} [{[] [nuts] 0.625 1}]} {{[beer butter] 0.25} [{[butter] [beer] 0.6666666666666666 1.0666666666666667}]} {{[beer cheese] 0.25} [{[cheese] [beer] 0.6666666666666666 1.0666666666666667}]} {{[beer cola] 0.375} [{[beer] [cola] 0.6 1.2} {[cola] [beer] 0.75 1.2}]} {{[beer jam] 0.375} [{[beer] [jam] 0.6 1.2} {[jam] [beer] 0.75 1.2}]} {{[beer lays] 0.375} [{[beer] [lays] 0.6 1.5999999999999999} {[lays] [beer] 1 1.6}]} {{[beer nuts] 0.5} [{[beer] [nuts] 0.8 1.28} {[nuts] [beer] 0.8 1.28}]} {{[beer pizza] 0.25} [{[pizza] [beer] 0.6666666666666666 1.0666666666666667}]} {{[butter cola] 0.25} [{[butter] [cola] 0.6666666666666666 1.3333333333333333} {[cola] [butter] 0.5 1.3333333333333333}]} {{[cheese cola] 0.25} [{[cheese] [cola] 0.6666666666666666 1.3333333333333333} {[cola] [cheese] 0.5 1.3333333333333333}]} {{[cheese lays] 0.25} [{[cheese] [lays] 0.6666666666666666 1.7777777777777777} {[lays] [cheese] 0.6666666666666666 1.7777777777777777}]} {{[cheese nuts] 0.375} [{[cheese] [nuts] 1 1.6} {[nuts] [cheese] 0.6 1.5999999999999999}]} {{[cheese pizza] 0.25} [{[cheese] [pizza] 0.6666666666666666 1.7777777777777777} {[pizza] [cheese] 0.6666666666666666 1.7777777777777777}]} {{[cola lays] 0.25} [{[cola] [lays] 0.5 1.3333333333333333} {[lays] [cola] 0.6666666666666666 1.3333333333333333}]} {{[cola nuts] 0.375} [{[cola] [nuts] 0.75 1.2} {[nuts] [cola] 0.6 1.2}]} {{[cola pizza] 0.375} [{[cola] [pizza] 0.75 2} {[pizza] [cola] 1 2}]} {{[jam nuts] 0.375} [{[jam] [nuts] 0.75 1.2} {[nuts] [jam] 0.6 1.2}]} {{[lays nuts] 0.25} [{[lays] [nuts] 0.6666666666666666 1.0666666666666667}]} {{[nuts pizza] 0.375} [{[nuts] [pizza] 0.6 1.5999999999999999} {[pizza] [nuts] 1 1.6}]} {{[beer butter cola] 0.25} [{[beer butter] [cola] 1 2} {[beer cola] [butter] 0.6666666666666666 1.7777777777777777} {[butter cola] [beer] 1 1.6}]} {{[beer butter jam] 0.125} [{[beer butter] [jam] 0.5 1} {[butter jam] [beer] 1 1.6}]} {{[beer butter lays] 0.125} [{[beer butter] [lays] 0.5 1.3333333333333333} {[butter lays] [beer] 1 1.6}]} {{[beer butter nuts] 0.125} [{[beer butter] [nuts] 0.5 0.8} {[butter nuts] [beer] 1 1.6}]} {{[beer butter pizza] 0.125} [{[beer butter] [pizza] 0.5 1.3333333333333333} {[beer pizza] [butter] 0.5 1.3333333333333333} {[butter pizza] [beer] 1 1.6}]} {{[beer cheese cola] 0.125} [{[beer cheese] [cola] 0.5 1} {[cheese cola] [beer] 0.5 0.8}]} {{[beer cheese jam] 0.125} [{[beer cheese] [jam] 0.5 1} {[cheese jam] [beer] 1 1.6}]} {{[beer cheese lays] 0.25} [{[beer cheese] [lays] 1 2.6666666666666665} {[beer lays] [cheese] 0.6666666666666666 1.7777777777777777} {[cheese lays] [beer] 1 1.6}]} {{[beer cheese nuts] 0.25} [{[beer cheese] [nuts] 1 1.6} {[beer nuts] [cheese] 0.5 1.3333333333333333} {[cheese nuts] [beer] 0.6666666666666666 1.0666666666666667}]} {{[beer cheese pizza] 0.125} [{[beer cheese] [pizza] 0.5 1.3333333333333333} {[beer pizza] [cheese] 0.5 1.3333333333333333} {[cheese pizza] [beer] 0.5 0.8}]} {{[beer cola jam] 0.125} [{[cola jam] [beer] 1 1.6}]} {{[beer cola lays] 0.25} [{[beer cola] [lays] 0.6666666666666666 1.7777777777777777} {[beer lays] [cola] 0.6666666666666666 1.3333333333333333} {[cola lays] [beer] 1 1.6}]} {{[beer cola nuts] 0.25} [{[beer cola] [nuts] 0.6666666666666666 1.0666666666666667} {[beer nuts] [cola] 0.5 1} {[cola nuts] [beer] 0.6666666666666666 1.0666666666666667}]} {{[beer cola pizza] 0.25} [{[beer cola] [pizza] 0.6666666666666666 1.7777777777777777} {[beer pizza] [cola] 1 2} {[cola pizza] [beer] 0.6666666666666666 1.0666666666666667}]} {{[beer jam lays] 0.125} [{[jam lays] [beer] 1 1.6}]} {{[beer jam nuts] 0.375} [{[beer jam] [nuts] 1 1.6} {[beer nuts] [jam] 0.75 1.5} {[jam nuts] [beer] 1 1.6}]} {{[beer jam pizza] 0.125} [{[beer pizza] [jam] 0.5 1} {[jam pizza] [beer] 1 1.6}]} {{[beer lays nuts] 0.25} [{[beer lays] [nuts] 0.6666666666666666 1.0666666666666667} {[beer nuts] [lays] 0.5 1.3333333333333333} {[lays nuts] [beer] 1 1.6}]} {{[beer lays pizza] 0.125} [{[beer pizza] [lays] 0.5 1.3333333333333333} {[lays pizza] [beer] 1 1.6}]} {{[beer nuts pizza] 0.25} [{[beer nuts] [pizza] 0.5 1.3333333333333333} {[beer pizza] [nuts] 1 1.6} {[nuts pizza] [beer] 0.6666666666666666 1.0666666666666667}]} {{[butter cola jam] 0.125} [{[butter cola] [jam] 0.5 1} {[butter jam] [cola] 1 2} {[cola jam] [butter] 1 2.6666666666666665}]} {{[butter cola lays] 0.125} [{[butter cola] [lays] 0.5 1.3333333333333333} {[butter lays] [cola] 1 2} {[cola lays] [butter] 0.5 1.3333333333333333}]} {{[butter cola nuts] 0.125} [{[butter cola] [nuts] 0.5 0.8} {[butter nuts] [cola] 1 2}]} {{[butter cola pizza] 0.125} [{[butter cola] [pizza] 0.5 1.3333333333333333} {[butter pizza] [cola] 1 2}]} {{[butter jam nuts] 0.125} [{[butter jam] [nuts] 1 1.6} {[butter nuts] [jam] 1 2}]} {{[butter jam pizza] 0.125} [{[butter jam] [pizza] 1 2.6666666666666665} {[butter pizza] [jam] 1 2} {[jam pizza] [butter] 1 2.6666666666666665}]} {{[butter nuts pizza] 0.125} [{[butter nuts] [pizza] 1 2.6666666666666665} {[butter pizza] [nuts] 1 1.6}]} {{[cheese cola lays] 0.125} [{[cheese cola] [lays] 0.5 1.3333333333333333} {[cheese lays] [cola] 0.5 1} {[cola lays] [cheese] 0.5 1.3333333333333333}]} {{[cheese cola nuts] 0.25} [{[cheese cola] [nuts] 1 1.6} {[cheese nuts] [cola] 0.6666666666666666 1.3333333333333333} {[cola nuts] [cheese] 0.6666666666666666 1.7777777777777777}]} {{[cheese cola pizza] 0.25} [{[cheese cola] [pizza] 1 2.6666666666666665} {[cheese pizza] [cola] 1 2} {[cola pizza] [cheese] 0.6666666666666666 1.7777777777777777}]} {{[cheese jam lays] 0.125} [{[cheese jam] [lays] 1 2.6666666666666665} {[cheese lays] [jam] 0.5 1} {[jam lays] [cheese] 1 2.6666666666666665}]} {{[cheese jam nuts] 0.125} [{[cheese jam] [nuts] 1 1.6}]} {{[cheese lays nuts] 0.25} [{[cheese lays] [nuts] 1 1.6} {[cheese nuts] [lays] 0.6666666666666666 1.7777777777777777} {[lays nuts] [cheese] 1 2.6666666666666665}]} {{[cheese lays pizza] 0.125} [{[cheese lays] [pizza] 0.5 1.3333333333333333} {[cheese pizza] [lays] 0.5 1.3333333333333333} {[lays pizza] [cheese] 1 2.6666666666666665}]} {{[cheese nuts pizza] 0.25} [{[cheese nuts] [pizza] 0.6666666666666666 1.7777777777777777} {[cheese pizza] [nuts] 1 1.6} {[nuts pizza] [cheese] 0.6666666666666666 1.7777777777777777}]} {{[cola jam nuts] 0.125} [{[cola jam] [nuts] 1 1.6}]} {{[cola jam pizza] 0.125} [{[cola jam] [pizza] 1 2.6666666666666665} {[jam pizza] [cola] 1 2}]} {{[cola lays nuts] 0.125} [{[cola lays] [nuts] 0.5 0.8} {[lays nuts] [cola] 0.5 1}]} {{[cola lays pizza] 0.125} [{[cola lays] [pizza] 0.5 1.3333333333333333} {[lays pizza] [cola] 1 2}]} {{[cola nuts pizza] 0.375} [{[cola nuts] [pizza] 1 2.6666666666666665} {[cola pizza] [nuts] 1 1.6} {[nuts pizza] [cola] 1 2}]} {{[jam lays nuts] 0.125} [{[jam lays] [nuts] 1 1.6} {[lays nuts] [jam] 0.5 1}]} {{[jam nuts pizza] 0.125} [{[jam pizza] [nuts] 1 1.6}]} {{[lays nuts pizza] 0.125} [{[lays nuts] [pizza] 0.5 1.3333333333333333} {[lays pizza] [nuts] 1 1.6}]} {{[beer butter cola jam] 0.125} [{[beer butter cola] [jam] 0.5 1} {[beer butter jam] [cola] 1 2} {[beer cola jam] [butter] 1 2.6666666666666665} {[butter cola jam] [beer] 1 1.6}]} {{[beer butter cola lays] 0.125} [{[beer butter cola] [lays] 0.5 1.3333333333333333} {[beer butter lays] [cola] 1 2} {[beer cola lays] [butter] 0.5 1.3333333333333333} {[butter cola lays] [beer] 1 1.6}]} {{[beer butter cola nuts] 0.125} [{[beer butter cola] [nuts] 0.5 0.8} {[beer butter nuts] [cola] 1 2} {[beer cola nuts] [butter] 0.5 1.3333333333333333} {[butter cola nuts] [beer] 1 1.6}]} {{[beer butter cola pizza] 0.125} [{[beer butter cola] [pizza] 0.5 1.3333333333333333} {[beer butter pizza] [cola] 1 2} {[beer cola pizza] [butter] 0.5 1.3333333333333333} {[butter cola pizza] [beer] 1 1.6}]} {{[beer butter jam nuts] 0.125} [{[beer butter jam] [nuts] 1 1.6} {[beer butter nuts] [jam] 1 2} {[butter jam nuts] [beer] 1 1.6}]} {{[beer butter jam pizza] 0.125} [{[beer butter jam] [pizza] 1 2.6666666666666665} {[beer butter pizza] [jam] 1 2} {[beer jam pizza] [butter] 1 2.6666666666666665} {[butter jam pizza] [beer] 1 1.6}]} {{[beer butter nuts pizza] 0.125} [{[beer butter nuts] [pizza] 1 2.6666666666666665} {[beer butter pizza] [nuts] 1 1.6} {[beer nuts pizza] [butter] 0.5 1.3333333333333333} {[butter nuts pizza] [beer] 1 1.6}]} {{[beer cheese cola lays] 0.125} [{[beer cheese cola] [lays] 1 2.6666666666666665} {[beer cheese lays] [cola] 0.5 1} {[beer cola lays] [cheese] 0.5 1.3333333333333333} {[cheese cola lays] [beer] 1 1.6}]} {{[beer cheese cola nuts] 0.125} [{[beer cheese cola] [nuts] 1 1.6} {[beer cheese nuts] [cola] 0.5 1} {[beer cola nuts] [cheese] 0.5 1.3333333333333333} {[cheese cola nuts] [beer] 0.5 0.8}]} {{[beer cheese cola pizza] 0.125} [{[beer cheese cola] [pizza] 1 2.6666666666666665} {[beer cheese pizza] [cola] 1 2} {[beer cola pizza] [cheese] 0.5 1.3333333333333333} {[cheese cola pizza] [beer] 0.5 0.8}]} {{[beer cheese jam lays] 0.125} [{[beer cheese jam] [lays] 1 2.6666666666666665} {[beer cheese lays] [jam] 0.5 1} {[beer jam lays] [cheese] 1 2.6666666666666665} {[cheese jam lays] [beer] 1 1.6}]} {{[beer cheese jam nuts] 0.125} [{[beer cheese jam] [nuts] 1 1.6} {[beer cheese nuts] [jam] 0.5 1} {[cheese jam nuts] [beer] 1 1.6}]} {{[beer cheese lays nuts] 0.25} [{[beer cheese lays] [nuts] 1 1.6} {[beer cheese nuts] [lays] 1 2.6666666666666665} {[beer lays nuts] [cheese] 1 2.6666666666666665} {[cheese lays nuts] [beer] 1 1.6}]} {{[beer cheese lays pizza] 0.125} [{[beer cheese lays] [pizza] 0.5 1.3333333333333333} {[beer cheese pizza] [lays] 1 2.6666666666666665} {[beer lays pizza] [cheese] 1 2.6666666666666665} {[cheese lays pizza] [beer] 1 1.6}]} {{[beer cheese nuts pizza] 0.125} [{[beer cheese nuts] [pizza] 0.5 1.3333333333333333} {[beer cheese pizza] [nuts] 1 1.6} {[beer nuts pizza] [cheese] 0.5 1.3333333333333333} {[cheese nuts pizza] [beer] 0.5 0.8}]} {{[beer cola jam nuts] 0.125} [{[beer cola jam] [nuts] 1 1.6} {[beer cola nuts] [jam] 0.5 1} {[cola jam nuts] [beer] 1 1.6}]} {{[beer cola jam pizza] 0.125} [{[beer cola jam] [pizza] 1 2.6666666666666665} {[beer cola pizza] [jam] 0.5 1} {[beer jam pizza] [cola] 1 2} {[cola jam pizza] [beer] 1 1.6}]} {{[beer cola lays nuts] 0.125} [{[beer cola lays] [nuts] 0.5 0.8} {[beer cola nuts] [lays] 0.5 1.3333333333333333} {[beer lays nuts] [cola] 0.5 1} {[cola lays nuts] [beer] 1 1.6}]} {{[beer cola lays pizza] 0.125} [{[beer cola lays] [pizza] 0.5 1.3333333333333333} {[beer cola pizza] [lays] 0.5 1.3333333333333333} {[beer lays pizza] [cola] 1 2} {[cola lays pizza] [beer] 1 1.6}]} {{[beer cola nuts pizza] 0.25} [{[beer cola nuts] [pizza] 1 2.6666666666666665} {[beer cola pizza] [nuts] 1 1.6} {[beer nuts pizza] [cola] 1 2} {[cola nuts pizza] [beer] 0.6666666666666666 1.0666666666666667}]} {{[beer jam lays nuts] 0.125} [{[beer jam lays] [nuts] 1 1.6} {[beer lays nuts] [jam] 0.5 1} {[jam lays nuts] [beer] 1 1.6}]} {{[beer jam nuts pizza] 0.125} [{[beer jam pizza] [nuts] 1 1.6} {[beer nuts pizza] [jam] 0.5 1} {[jam nuts pizza] [beer] 1 1.6}]} {{[beer lays nuts pizza] 0.125} [{[beer lays nuts] [pizza] 0.5 1.3333333333333333} {[beer lays pizza] [nuts] 1 1.6} {[beer nuts pizza] [lays] 0.5 1.3333333333333333} {[lays nuts pizza] [beer] 1 1.6}]} {{[butter cola jam nuts] 0.125} [{[butter cola jam] [nuts] 1 1.6} {[butter cola nuts] [jam] 1 2} {[butter jam nuts] [cola] 1 2} {[cola jam nuts] [butter] 1 2.6666666666666665}]} {{[butter cola jam pizza] 0.125} [{[butter cola jam] [pizza] 1 2.6666666666666665} {[butter cola pizza] [jam] 1 2} {[butter jam pizza] [cola] 1 2} {[cola jam pizza] [butter] 1 2.6666666666666665}]} {{[butter cola nuts pizza] 0.125} [{[butter cola nuts] [pizza] 1 2.6666666666666665} {[butter cola pizza] [nuts] 1 1.6} {[butter nuts pizza] [cola] 1 2}]} {{[butter jam nuts pizza] 0.125} [{[butter jam nuts] [pizza] 1 2.6666666666666665} {[butter jam pizza] [nuts] 1 1.6} {[butter nuts pizza] [jam] 1 2} {[jam nuts pizza] [butter] 1 2.6666666666666665}]} {{[cheese cola lays nuts] 0.125} [{[cheese cola lays] [nuts] 1 1.6} {[cheese cola nuts] [lays] 0.5 1.3333333333333333} {[cheese lays nuts] [cola] 0.5 1} {[cola lays nuts] [cheese] 1 2.6666666666666665}]} {{[cheese cola lays pizza] 0.125} [{[cheese cola lays] [pizza] 1 2.6666666666666665} {[cheese cola pizza] [lays] 0.5 1.3333333333333333} {[cheese lays pizza] [cola] 1 2} {[cola lays pizza] [cheese] 1 2.6666666666666665}]} {{[cheese cola nuts pizza] 0.25} [{[cheese cola nuts] [pizza] 1 2.6666666666666665} {[cheese cola pizza] [nuts] 1 1.6} {[cheese nuts pizza] [cola] 1 2} {[cola nuts pizza] [cheese] 0.6666666666666666 1.7777777777777777}]} {{[cheese jam lays nuts] 0.125} [{[cheese jam lays] [nuts] 1 1.6} {[cheese jam nuts] [lays] 1 2.6666666666666665} {[cheese lays nuts] [jam] 0.5 1} {[jam lays nuts] [cheese] 1 2.6666666666666665}]} {{[cheese lays nuts pizza] 0.125} [{[cheese lays nuts] [pizza] 0.5 1.3333333333333333} {[cheese lays pizza] [nuts] 1 1.6} {[cheese nuts pizza] [lays] 0.5 1.3333333333333333} {[lays nuts pizza] [cheese] 1 2.6666666666666665}]} {{[cola jam nuts pizza] 0.125} [{[cola jam nuts] [pizza] 1 2.6666666666666665} {[cola jam pizza] [nuts] 1 1.6} {[jam nuts pizza] [cola] 1 2}]} {{[cola lays nuts pizza] 0.125} [{[cola lays nuts] [pizza] 1 2.6666666666666665} {[cola lays pizza] [nuts] 1 1.6} {[lays nuts pizza] [cola] 1 2}]} {{[beer butter cola jam nuts] 0.125} [{[beer butter cola jam] [nuts] 1 1.6} {[beer butter cola nuts] [jam] 1 2} {[beer butter jam nuts] [cola] 1 2} {[beer cola jam nuts] [butter] 1 2.6666666666666665} {[butter cola jam nuts] [beer] 1 1.6}]} {{[beer butter cola jam pizza] 0.125} [{[beer butter cola jam] [pizza] 1 2.6666666666666665} {[beer butter cola pizza] [jam] 1 2} {[beer butter jam pizza] [cola] 1 2} {[beer cola jam pizza] [butter] 1 2.6666666666666665} {[butter cola jam pizza] [beer] 1 1.6}]} {{[beer butter cola nuts pizza] 0.125} [{[beer butter cola nuts] [pizza] 1 2.6666666666666665} {[beer butter cola pizza] [nuts] 1 1.6} {[beer butter nuts pizza] [cola] 1 2} {[beer cola nuts pizza] [butter] 0.5 1.3333333333333333} {[butter cola nuts pizza] [beer] 1 1.6}]} {{[beer butter jam nuts pizza] 0.125} [{[beer butter jam nuts] [pizza] 1 2.6666666666666665} {[beer butter jam pizza] [nuts] 1 1.6} {[beer butter nuts pizza] [jam] 1 2} {[beer jam nuts pizza] [butter] 1 2.6666666666666665} {[butter jam nuts pizza] [beer] 1 1.6}]} {{[beer cheese cola lays nuts] 0.125} [{[beer cheese cola lays] [nuts] 1 1.6} {[beer cheese cola nuts] [lays] 1 2.6666666666666665} {[beer cheese lays nuts] [cola] 0.5 1} {[beer cola lays nuts] [cheese] 1 2.6666666666666665} {[cheese cola lays nuts] [beer] 1 1.6}]} {{[beer cheese cola lays pizza] 0.125} [{[beer cheese cola lays] [pizza] 1 2.6666666666666665} {[beer cheese cola pizza] [lays] 1 2.6666666666666665} {[beer cheese lays pizza] [cola] 1 2} {[beer cola lays pizza] [cheese] 1 2.6666666666666665} {[cheese cola lays pizza] [beer] 1 1.6}]} {{[beer cheese cola nuts pizza] 0.125} [{[beer cheese cola nuts] [pizza] 1 2.6666666666666665} {[beer cheese cola pizza] [nuts] 1 1.6} {[beer cheese nuts pizza] [cola] 1 2} {[beer cola nuts pizza] [cheese] 0.5 1.3333333333333333} {[cheese cola nuts pizza] [beer] 0.5 0.8}]} {{[beer cheese jam lays nuts] 0.125} [{[beer cheese jam lays] [nuts] 1 1.6} {[beer cheese jam nuts] [lays] 1 2.6666666666666665} {[beer cheese lays nuts] [jam] 0.5 1} {[beer jam lays nuts] [cheese] 1 2.6666666666666665} {[cheese jam lays nuts] [beer] 1 1.6}]} {{[beer cheese lays nuts pizza] 0.125} [{[beer cheese lays nuts] [pizza] 0.5 1.3333333333333333} {[beer cheese lays pizza] [nuts] 1 1.6} {[beer cheese nuts pizza] [lays] 1 2.6666666666666665} {[beer lays nuts pizza] [cheese] 1 2.6666666666666665} {[cheese lays nuts pizza] [beer] 1 1.6}]} {{[beer cola jam nuts pizza] 0.125} [{[beer cola jam nuts] [pizza] 1 2.6666666666666665} {[beer cola jam pizza] [nuts] 1 1.6} {[beer cola nuts pizza] [jam] 0.5 1} {[beer jam nuts pizza] [cola] 1 2} {[cola jam nuts pizza] [beer] 1 1.6}]} {{[beer cola lays nuts pizza] 0.125} [{[beer cola lays nuts] [pizza] 1 2.6666666666666665} {[beer cola lays pizza] [nuts] 1 1.6} {[beer cola nuts pizza] [lays] 0.5 1.3333333333333333} {[beer lays nuts pizza] [cola] 1 2} {[cola lays nuts pizza] [beer] 1 1.6}]} {{[butter cola jam nuts pizza] 0.125} [{[butter cola jam nuts] [pizza] 1 2.6666666666666665} {[butter cola jam pizza] [nuts] 1 1.6} {[butter cola nuts pizza] [jam] 1 2} {[butter jam nuts pizza] [cola] 1 2} {[cola jam nuts pizza] [butter] 1 2.6666666666666665}]} {{[cheese cola lays nuts pizza] 0.125} [{[cheese cola lays nuts] [pizza] 1 2.6666666666666665} {[cheese cola lays pizza] [nuts] 1 1.6} {[cheese cola nuts pizza] [lays] 0.5 1.3333333333333333} {[cheese lays nuts pizza] [cola] 1 2} {[cola lays nuts pizza] [cheese] 1 2.6666666666666665}]} {{[beer butter cola jam nuts pizza] 0.125} [{[beer butter cola jam nuts] [pizza] 1 2.6666666666666665} {[beer butter cola jam pizza] [nuts] 1 1.6} {[beer butter cola nuts pizza] [jam] 1 2} {[beer butter jam nuts pizza] [cola] 1 2} {[beer cola jam nuts pizza] [butter] 1 2.6666666666666665} {[butter cola jam nuts pizza] [beer] 1 1.6}]} {{[beer cheese cola lays nuts pizza] 0.125} [{[beer cheese cola lays nuts] [pizza] 1 2.6666666666666665} {[beer cheese cola lays pizza] [nuts] 1 1.6} {[beer cheese cola nuts pizza] [lays] 1 2.6666666666666665} {[beer cheese lays nuts pizza] [cola] 1 2} {[beer cola lays nuts pizza] [cheese] 1 2.6666666666666665} {[cheese cola lays nuts pizza] [beer] 1 1.6}]}]"},
	}

	for _, data := range provider {
		a := NewApriori(data.in)
		out := a.Calculate(NewOptions(0.1, 0.5, 0.0, 0))

		assert(data.out == fmt.Sprint(out), "Expected output not equal to actual output")
		fmt.Printf("%+v\n", out)
	}
}

func assert(b bool, s string) {
	if !b {
		println(s)
		os.Exit(1)
	}
}
