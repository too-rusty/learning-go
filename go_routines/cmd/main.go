package main

import (
	"fmt"
	"math/big"
	"sync"
	"time"
)

func timeIt(f func() interface{}) (int64, interface{}) {
	// time taken and the return value
	now := time.Now()
	v := f()
	var ret interface{}
	if val, ok := v.(*big.Int); ok {
		ret = interface{}(val)
	}
	return time.Since(now).Milliseconds(), ret
}

const MAX = 1000 * 100

func factorial() interface{} {
	x := big.NewInt(1)
	for i := 1; i < MAX; i++ {
		x.Mul(x, big.NewInt(int64(i)))
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func factorial_go() interface{} {
	start, offset := 1, 1000
	var arr []*big.Int = make([]*big.Int, MAX/offset+100)
	for i := 0; i < len(arr); i++ {
		arr[i] = big.NewInt(1)
	}
	var wg sync.WaitGroup

	for i := start; i < MAX; i += offset {
		wg.Add(1)
		go func(i int) {
			at := i / offset
			res := big.NewInt(1)
			for j := i; j < min(i+offset, MAX); j++ {
				res.Mul(res, big.NewInt(int64(j)))
			}
			arr[at] = res
			wg.Done()
		}(i)
	}
	wg.Wait()
	res := big.NewInt(1)
	for _, v := range arr {
		res.Mul(res, v)
	}
	return interface{}(res)
}

func main() {

	// x := big.NewInt(1)
	// x.Mul(x, big.NewInt(2))
	// fmt.Printf("%T\n", x)

	// runtime.GOMAXPROCS(runtime.NumCPU())
	tim1, val1 := timeIt(factorial)
	tim2, val2 := timeIt(factorial_go)
	fmt.Println(tim1, tim2)

	_, v1 := val1.(*big.Int)
	_, v2 := val2.(*big.Int)
	if v1 == v2 {
		fmt.Println("EQUAL!!")
	}

}
