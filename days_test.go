package test

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var expectations = map[string]string{
	"1-1":  "1222801",
	"1-2":  "22545250",
	"2-1":  "334",
	"2-2":  "400",
	"3-1":  "166630675",
	"3-2":  "93465710",
	"4-1":  "2639",
	"4-2":  "2005",
	"5-1":  "6951",
	"5-2":  "4121",
	"6-1":  "5177",
	"6-2":  "1686",
	"7-1":  "6231007345478",
	"7-2":  "333027885676693",
	"8-1":  "269",
	"8-2":  "949",
	"9-1":  "6283170117911",
	"9-2":  "6307653242596",
	"10-1": "538",
	"10-2": "1110",
	"11-1": "186203",
	"11-2": "221291560078593",
	"12-1": "1431440",
	"12-2": "869070",
	"13-1": "29522",
	"13-2": "101214869433312",
	"14-1": "232589280",
	"14-2": "7569",
	"15-1": "1514333",
	"15-2": "1528453",
}

func TestDays(t *testing.T) {
	for day, expect := range expectations {
		t.Run(day, func(t *testing.T) {
			t.Parallel()
			buildCmd := exec.Command("go", "build", "-o", fmt.Sprintf("day-%s", day), fmt.Sprintf("days/%s/main.go", day))
			output, err := buildCmd.CombinedOutput()
			if err != nil {
				fmt.Println(output)
			}

			runCmd := exec.Command(fmt.Sprintf("./day-%s", day))
			output, err = runCmd.CombinedOutput()
			if err != nil {
				fmt.Println(output)
			}

			assert.NoError(t, err)
			assert.Equal(t, expect, strings.TrimRight(string(output), "\n"), fmt.Sprintf("Day %s", day))

			rmCmd := exec.Command("rm", fmt.Sprintf("day-%s", day))
			output, err = rmCmd.CombinedOutput()
			if err != nil {
				fmt.Println(output)
			}
		})
	}
}
