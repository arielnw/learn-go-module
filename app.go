package learngomodule

import "fmt"

func PrintSliceOfString(slc []string) {
	for _, item := range slc {
		fmt.Println(item)
	}
}

// step setelah push :
// 1. create tangging
// git tag v1.0.0
// git push origin v1.0.0

// cara get package
// go get -u github.com/google/uuid

// https://semver.org/spec/v2.0.0.html
// semua package yang resmi masuk sini ya https://pkg.go.dev
