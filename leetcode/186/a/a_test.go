// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	exampleIns := [][]string{{`"011101"`}, {`"00111"`}, {`"1111"`}}
	exampleOuts := [][]string{{`5`}, {`5`}, {`3`}}
	// TODO: 测试参数的下界和上界！
	// custom test cases or WA cases.
	exampleIns = append(exampleIns, []string{`"00"`})
	exampleOuts = append(exampleOuts, []string{`1`})
	targetCaseNum := 4
	if err := testutil.RunLeetCodeFuncWithCase(t, maxScore, exampleIns, exampleOuts, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-186/problems/maximum-score-after-splitting-a-string/
