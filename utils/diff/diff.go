package diff

import (
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func TextDiff(text1, text2 string) []string {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(text1, text2, false)

	for _, diff := range diffs {
		fmt.Println(diff)

	}

	fmt.Println(diffs)
	return []string{}
}
