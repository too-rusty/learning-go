package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("data stream project")

	c := dataStream1()
	for v := range c {
		fmt.Println(string(v))
	}
	_, ok := <-c
	if !ok {
		fmt.Println("CHANNEL CLOSED")
	}
}

func dataStream1() <-chan []byte {
	// for now emits data but ultimately writes to a file (hint maybe use encoder)
	// returns a read only chan for now
	data := Data{
		Name: "abc",
		Age:  25,
	}
	c := make(chan []byte)
	go func() {
		for {
			byt, err := json.Marshal(data)
			if err != nil {
				close(c)
				break
			}
			c <- byt
			// close(c)
			// return
		}
	}()
	return c
}

func filter(data Data) (ok bool) {
	return
}

type Data struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

/*
make a data stream which emits multi format data
--  multiple data sources that emit all kinds of data and that too mixed
make a go routine that ingests the data and parses it and moves it to appropriate KV store bucket
maybe writes data to multiple binary files in some encoded format and reads data and writes
it to another file in decoded format, filtering out the useless data
*/

/*

TASKS
1. randomise the data
2. filter based on age
3. encode the data so that it needs to be decoded
4. try to make channels for everything
5. use multiple data sources for multiple data return value
	and handle them using the same go routine, ( use select )



LEARNINGS
1. closing a channel returns false and 0 value
*/
