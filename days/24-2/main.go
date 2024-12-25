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

func solution() string {
	parts, err := util.ReadStrings("24", false, "\n\n")
	if err != nil {
		panic(err)
	}

	inputX := []outVal{}
	inputY := []outVal{}
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
				if matches[1] == "x" {
					inputX = append(inputX, outVal{
						name: matches[1] + matches[2],
						val:  val,
					})
				} else if matches[1] == "y" {
					inputY = append(inputY, outVal{
						name: matches[1] + matches[2],
						val:  val,
					})
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

	var swapList []string
	swap(rules, "kth", "z12", &swapList)
	swap(rules, "gsd", "z26", &swapList)
	swap(rules, "tbt", "z32", &swapList)
	swap(rules, "qnf", "vpm", &swapList)

	slices.Sort(swapList)

	output := []outVal{}
	for out := range rules {
		if out[0] == 'z' {
			val := recurseRules(rules, out)
			output = append(output, outVal{
				name: out,
				val:  val,
			})
			rule := rules[out]
			rule.output = &val
			rules[out] = rule
		}
	}

	inputXNum := numFromBinaryList(inputX)
	inputYNum := numFromBinaryList(inputY)
	outputNum := numFromBinaryList(output)

	if outputNum != inputXNum+inputYNum {
		panic("error")
	}

	return strings.Join(swapList, ",")
}

func swap(rules map[string]rule, swapA, swapB string, swapList *[]string) {
	ruleA := rules[swapA]
	ruleB := rules[swapB]
	rules[swapA] = ruleB
	rules[swapB] = ruleA

	*swapList = append(*swapList, swapA, swapB)
}

func formatCheck(rules map[string]rule) {
	nextCarry := "bdj"
	for i := 1; i < 45; i++ {
		var xVal string
		var yVal string
		var zVal string
		if i < 10 {
			xVal = fmt.Sprintf("x0%d", i)
			yVal = fmt.Sprintf("y0%d", i)
			zVal = fmt.Sprintf("z0%d", i)
		} else {
			xVal = fmt.Sprintf("x%d", i)
			yVal = fmt.Sprintf("y%d", i)
			zVal = fmt.Sprintf("z%d", i)
		}

		andOut := ""
		xorOut := ""
		for out, r := range rules {
			if (r.valA == yVal && r.valB == xVal) ||
				(r.valA == xVal && r.valB == yVal) {
				if r.operation == "AND" {
					andOut = out
					if xorOut != "" {
						break
					}
				} else if r.operation == "XOR" {
					xorOut = out
					if andOut != "" {
						break
					}
				}
			}
		}

		if rules[zVal].valA != nextCarry && rules[zVal].valB != nextCarry {
			fmt.Printf("unexpected next carry: %s - expected %s\n", rules[zVal].valB, nextCarry)
		}

		carryAdd := ""
		for out, r := range rules {
			if (r.valA == xorOut && r.valB == nextCarry) ||
				(r.valA == nextCarry && r.valB == xorOut) {
				if r.operation == "AND" {
					carryAdd = out
					break
				}
			}
		}

		for out, r := range rules {
			if (r.valA == carryAdd && r.valB == andOut) ||
				(r.valA == andOut && r.valB == carryAdd) {
				if r.operation == "OR" {
					nextCarry = out
					break
				}
			}
		}

		fmt.Println(i)
		fmt.Printf("1st half adder and: %s\n", andOut)
		fmt.Printf("1st half adder xor: %s\n", xorOut)
		fmt.Printf("2nd half adder and: %s\n", carryAdd)
		fmt.Printf("or carry: %s\n", nextCarry)
	}
}

func numFromBinaryList(input []outVal) int {
	slices.SortFunc(input, func(a, b outVal) int {
		return cmp.Compare(b.name, a.name)
	})

	var b []byte
	for _, in := range input {
		b = append(b, byte(in.val)+'0')
	}

	intNum, err := strconv.ParseInt(string(b), 2, 64)
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
