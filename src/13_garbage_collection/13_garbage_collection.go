package main

import (
	"fmt"
	"runtime"
	"time"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc(MB):", mem.Alloc/1000000)
	fmt.Println("mem.TotalAlloc(MB):", mem.TotalAlloc/1000000)
	fmt.Println("mem.HeapAlloc(MB):", mem.HeapAlloc/1000000)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("-----")
}

func main() {
	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 10; i++ {
		// approx 1GB of mem allocation
		s := make([]byte, 1000000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
		time.Sleep(5 * time.Second)
	}
	printStats(mem)
}

/*

[jd@jdpc go_learning]$ GODEBUG=gctrace=1 go run 13_garbage_collection/13_garbage_collection.go
...
...
mem.Alloc(MB): 0
mem.TotalAlloc(MB): 0
mem.HeapAlloc(MB): 0
mem.NumGC: 0
-----
gc 1 @0.011s 0%: 0.003+0.071+0.014 ms clock, 0.013+0.027/0.023/0.081+0.059 ms cpu, 953->953->0 MB, 954 MB goal, 4 P
gc 2 @5.417s 0%: 0.003+0.092+0.092 ms clock, 0.013+0/0.066/0+0.37 ms cpu, 953->953->0 MB, 954 MB goal, 4 P
gc 3 @10.480s 0%: 0.003+0.12+0.062 ms clock, 0.012+0/0.068/0+0.25 ms cpu, 953->953->0 MB, 954 MB goal, 4 P
gc 4 @15.545s 0%: 0.003+0.11+0.017 ms clock, 0.012+0/0.067/0+0.068 ms cpu, 953->953->0 MB, 954 MB goal, 4 P
gc 5 @20.605s 0%: 0.002+0.14+0.095 ms clock, 0.011+0/0.063/0.061+0.38 ms cpu, 953->953->0 MB, 954 MB goal, 4 P
gc 6 @25.665s 0%: 0.003+0.12+0.017 ms clock, 0.013+0/0.074/0+0.071 ms cpu, 953->953->0 MB, 954 MB goal, 4 P
gc 7 @30.701s 0%: 0.002+0.10+0.036 ms clock, 0.010+0/0.083/0+0.14 ms cpu, 953->953->0 MB, 954 MB goal, 4 P
gc 8 @35.762s 0%: 0.003+0.12+0.10 ms clock, 0.013+0/0.068/0+0.40 ms cpu, 953->953->0 MB, 954 MB goal, 4 P
gc 9 @40.824s 0%: 0.003+0.14+0.094 ms clock, 0.012+0/0.069/0+0.37 ms cpu, 953->953->0 MB, 954 MB goal, 4 P
gc 10 @45.888s 0%: 0.002+0.12+0.019 ms clock, 0.011+0/0.069/0+0.076 ms cpu, 953->953->0 MB, 954 MB goal, 4 P
mem.Alloc(MB): 0
mem.TotalAlloc(MB): 10000
mem.HeapAlloc(MB): 0
mem.NumGC: 10
-----

*/
