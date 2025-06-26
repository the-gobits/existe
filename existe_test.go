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
	t.Log(v)

	type EXTest struct {
		V        any
		Expected bool
		Key      string
	}

	tests := []EXTest{
		{V: v, Key: "algo", Expected: true},
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
