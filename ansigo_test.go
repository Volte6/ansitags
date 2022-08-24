package ansigo

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

type TestCase struct {
	Input    string `yaml:"input"`
	Expected string `yaml:"expected"`
}

func TestParse(t *testing.T) {

	testTable := loadTestFile("testdata/ansigo_test.yaml")

	for name, testCase := range testTable {

		t.Run(name, func(t *testing.T) {

			output := Parse(testCase.Input)
			assert.Equal(t, testCase.Expected, output)

			// fmt.Println(output)
			//bytes, _ := json.Marshal(output)
			//fmt.Println(string(bytes))
		})
	}

}

func loadTestFile(filename string) map[string]TestCase {

	data := make(map[string]TestCase, 100)

	if yfile, err := ioutil.ReadFile(filename); err != nil {
		panic(err)
	} else {
		if err := yaml.Unmarshal(yfile, &data); err != nil {
			panic(err)
		}
	}

	return data

}

//
// Benchmarks
// cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
//

//
// BenchmarkSprintf
// BenchmarkSprintf-16               500000               114.7 ns/op             8 B/op          1 allocs/op
//
func BenchmarkSprintf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = fmt.Sprintf("\033[%d;%dm", 8+60, 7)
	}
}

//
// BenchmarkAtoI
// BenchmarkAtoI-16                  500000                34.44 ns/op            0 B/op          0 allocs/op
//
func BenchmarkAtoI(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = "\033[" + strconv.Itoa(8+60) + ";" + strconv.Itoa(7) + "m"
	}
}

//
// BenchmarkConcat
// BenchmarkConcat-16                500000            168857 ns/op         1403048 B/op          2 allocs/op
//

func BenchmarkConcat(b *testing.B) {
	var sConcat string = ""
	for n := 0; n < b.N; n++ {
		sConcat += strconv.Itoa(n)
	}
}

//
// BenchmarkStringBuilder
// BenchmarkStringBuilder-16         500000                38.46 ns/op           41 B/op          0 allocs/op
//
func BenchmarkStringBuilder(b *testing.B) {
	var sBuilder strings.Builder
	for n := 0; n < b.N; n++ {
		sBuilder.WriteString(strconv.Itoa(n))
	}
}

// BenchmarkParseColorString
// BenchmarkParseColorString-16    	  617618	      1665 ns/op	     808 B/op	      16 allocs/op
func BenchmarkParseColorString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parse("This is a prefix <ansi fg=blue bg=red>This is color</ansi> This is a suffix")
	}
}

// BenchmarkParseColorInt
// BenchmarkParseColorInt-16    	  880810	      1393 ns/op	     712 B/op	      14 allocs/op
func BenchmarkParseColorInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parse("This is a prefix <ansi fg='9' bg='2'>This is color</ansi> This is a suffix")
	}
}
