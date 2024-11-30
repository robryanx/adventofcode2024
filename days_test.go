package test

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var expectations = map[string]string{}

func TestDays(t *testing.T) {
	for day, expect := range expectations {
		expect := expect
		day := day
		t.Run(day, func(t *testing.T) {
			t.Parallel()
			runCmd := exec.Command("go", "run", fmt.Sprintf("days/%s/main.go", day))
			output, err := runCmd.CombinedOutput()
			if err != nil {
				fmt.Println(output)
			}

			assert.NoError(t, err)
			assert.Equal(t, expect, strings.TrimRight(string(output), "\n"), fmt.Sprintf("Day %s", day))
		})
	}
}
