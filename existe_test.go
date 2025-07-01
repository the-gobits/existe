package existe

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"testing"
	"text/template"
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
		{V: nil, Key: "algo", Expected: false},
		{V: 1, Key: "algo", Expected: false},
		{V: "12", Key: "algo", Expected: false},
		{V: v, Key: "algo", Expected: true},
		{V: v, Key: "algo_random", Expected: false},
		{V: v, Key: "algo.0", Expected: true},
		{V: v, Key: "algo.-1", Expected: false},
		{V: [3]int{1, 2, 3}, Key: "1", Expected: true},
		{V: [3]int{1, 2, 3}, Key: "3", Expected: false},
		{V: v, Key: "algo.0.metadata", Expected: true},
		{V: v, Key: "algo.0.falso", Expected: false},
		{V: *emptyTest, Key: "Expected", Expected: true},
		{V: *emptyTest, Key: "Falso", Expected: false},
		{V: StructWithPtr{A: emptyTest}, Key: "A.V", Expected: true},
		{V: StructWithPtr{A: nil}, Key: "A.V", Expected: false},
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

func TestExisteComTemplate(t *testing.T) {
	type EXTTest struct {
		V        any
		Expected string
		Template string
	}

	type P struct {
		AI string
	}

	type T struct {
		P *P
	}

	type A struct {
		T T
	}

	tt := []EXTTest{
		{V: A{T: T{P: &P{AI: "asd"}}}, Expected: "true", Template: `{{ existe . "T.P" }}`},
		{V: A{T: T{P: &P{AI: "asd"}}}, Expected: "true", Template: `{{ existe . "T.P.AI" }}`},
		{V: A{T: T{}}, Expected: "false", Template: `{{ existe . "T.P.AI" }}`},
	}

	for _, ttu := range tt {
		temp, err := template.New("base").Funcs(template.FuncMap{"existe": Existe}).Parse(ttu.Template)
		if err != nil {
			t.Error(err)
		}
		sb := &strings.Builder{}

		temp.Execute(sb, ttu.V)
		if s := sb.String(); s != ttu.Expected {
			t.Errorf(`Expected %q, but got %q`, ttu.Expected, s)
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
