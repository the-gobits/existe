package existe

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

func TestExiste(t *testing.T) {
	var v any
	LoadJson("./data/a.json", &v)

	type EXTest struct {
		V        any
		Expected bool
		Key      string
	}

	type StructWithPtr struct {
		A *EXTest
	}

	type StructWithDeepPtr struct {
		A **EXTest
	}

	emptyTest := &EXTest{}

	tests := []EXTest{
		{V: v, Key: "algo", Expected: true},
		{V: v, Key: "algo_random", Expected: false},
		{V: v, Key: "algo.0", Expected: true},
		{V: v, Key: "algo.-1", Expected: false},
		{V: v, Key: "algo.0.metadata", Expected: true},
		{V: v, Key: "algo.0.falso", Expected: false},
		{V: *emptyTest, Key: "Expected", Expected: true},
		{V: *emptyTest, Key: "Falso", Expected: false},
		{V: StructWithPtr{A: emptyTest}, Key: "A.V", Expected: true},
		{V: StructWithDeepPtr{A: &emptyTest}, Key: "A.V", Expected: true},
	}

	for _, test := range tests {
		got := Existe(test.V, test.Key)
		if got != test.Expected {
			t.Logf("Expected %v, but got %v, for Existe(%v, %q)", test.Expected, got, test.V, test.Key)
			t.Fail()
		}
	}
}

func LoadJson(filename string, v *any) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(f)

	err = decoder.Decode(v)
	if err != nil {
		log.Fatal(err)
	}
}
