package main

import (
	"cmp"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2024/util"
)

func main() {
	fmt.Println(solution())
}

var inputRegex = regexp.MustCompile(`([a-z]{1})([0-9]{2}): ([0-9]{1})`)
var ruleRegex = regexp.MustCompile("([a-z0-9]{3}) (AND|OR|XOR) ([a-z0-9]{3}) -> ([a-z0-9]{3})")

type rule struct {
	valA      string
	valB      string
	operation string
	output    *int
}

type outVal struct {
	name string
	val  int
}

func solution() int {
	parts, err := util.ReadStrings("24", false, "\n\n")
	if err != nil {
		panic(err)
	}

	rules := map[string]rule{}
	for part := range parts {
		if len(rules) == 0 {
			for _, row := range strings.Split(part, "\n") {
				matches := inputRegex.FindStringSubmatch(row)
				val, err := strconv.Atoi(matches[3])
				if err != nil {
					panic(err)
				}
				rules[matches[1]+matches[2]] = rule{
					output: &val,
				}
			}
		} else {
			for _, row := range strings.Split(part, "\n") {
				matches := ruleRegex.FindStringSubmatch(row)
				rules[matches[4]] = rule{
					valA:      matches[1],
					valB:      matches[3],
					operation: matches[2],
				}
			}
		}
	}

	output := []outVal{}
	for out := range rules {
		if out[0] == 'z' {
			val := recurseRules(rules, out)
			output = append(output, outVal{
				name: out,
				val:  val,
			})
		}
	}

	slices.SortFunc(output, func(a, b outVal) int {
		return cmp.Compare(b.name, a.name)
	})

	var outNum []byte
	for _, out := range output {
		outNum = append(outNum, byte(out.val)+'0')
	}

	intNum, err := strconv.ParseInt(string(outNum), 2, 64)
	if err != nil {
		panic(err)
	}

	return int(intNum)
}

func recurseRules(rules map[string]rule, out string) int {
	if rules[out].output != nil {
		return *rules[out].output
	}

	var valA int
	if rules[rules[out].valA].output == nil {
		valA = recurseRules(rules, rules[out].valA)
		rule := rules[rules[out].valA]
		rule.output = &valA
		rules[rules[out].valA] = rule
	} else {
		valA = *rules[rules[out].valA].output
	}

	var valB int
	if rules[rules[out].valB].output == nil {
		valB = recurseRules(rules, rules[out].valB)
		rule := rules[rules[out].valB]
		rule.output = &valB
		rules[rules[out].valB] = rule
	} else {
		valB = *rules[rules[out].valB].output
	}

	switch rules[out].operation {
	case "AND":
		return valA & valB
	case "OR":
		return valA | valB
	case "XOR":
		return valA ^ valB
	default:
		panic("invalid opteration")
	}

	return -1
}
