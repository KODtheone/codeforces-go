// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	exampleIns := [][]string{{`[6,4,14,6,8,13,9,7,10,6,12]`, `2`}, {`[3,3,3,3,3]`, `3`}, {`[7,6,5,4,3,2,1]`, `1`}, {`[7,1,7,1,7,1]`, `2`}, {`[66]`, `1`}}
	exampleOuts := [][]string{{`4`}, {`1`}, {`7`}, {`2`}, {`1`}}
	// custom test cases or WA cases.
	exampleIns = append(exampleIns, []string{`[22,29,52,97,29,75,78,2,92,70,90,12,43,17,97,18,58,100,41,32]`,`17`})
	exampleOuts = append(exampleOuts, []string{`6`})
	if err := testutil.RunLeetCodeFuncWithCase(t, maxJumps, exampleIns, exampleOuts, 0); err != nil {
		t.Fatal(err)
	}
}
