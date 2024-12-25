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
	"16-1": "75416",
	"16-2": "476",
	"17-1": "7,1,2,3,2,6,7,2,5",
	"17-2": "202356708354602",
	"18-1": "298",
	"18-2": "52,32",
	"19-1": "353",
	"19-2": "880877787214477",
	"20-1": "1399",
	"20-2": "994807",
	"21-1": "164960",
	"21-2": "205620604017764",
	"22-1": "16299144133",
	"22-2": "1896",
	"23-1": "1077",
	"23-2": "bc,bf,do,dw,dx,ll,ol,qd,sc,ua,xc,yu,zt",
	"24-1": "55544677167336",
	"24-2": "gsd,kth,qnf,tbt,vpm,z12,z26,z32",
	"25-1": "3338",
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
