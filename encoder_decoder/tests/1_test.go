package tests

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"
)

func TestMarshalInterfaceForDataType(t *testing.T) {

	b := []byte(`{"id":1,"name":"bowser","breed":"husky","born_at":1480979203}`)
	var d Dog
	err := json.Unmarshal(b, &d)
	if err != nil {
		t.Fatal(err)
	}
	var b2 []byte
	b2, err = json.Marshal(d)
	if !bytes.Equal(b, b2) {
		t.Fatal(err)
	}

}

type Dog struct {
	ID     int
	Name   string
	Breed  string
	BornAt time.Time
}

type JSONDog struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Breed  string `json:"breed"`
	BornAt int64  `json:"born_at"`
}

func NewJSONDog(dog Dog) JSONDog {
	return JSONDog{
		dog.ID,
		dog.Name,
		dog.Breed,
		dog.BornAt.Unix(),
	}
}

func (jd JSONDog) Dog() Dog {
	return Dog{
		jd.ID,
		jd.Name,
		jd.Breed,
		time.Unix(jd.BornAt, 0),
	}
}

func (d Dog) MarshalJSON() ([]byte, error) {
	return json.Marshal(NewJSONDog(d))
}

func (d *Dog) UnmarshalJSON(data []byte) error {
	var jd JSONDog
	err := json.Unmarshal(data, &jd)
	if err != nil {
		return err
	}
	*d = jd.Dog()
	return nil
}
