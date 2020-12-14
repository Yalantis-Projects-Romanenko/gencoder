package encodings

import (
	"fmt"
	"github.com/fdistorted/gencoder/encodings/numbers"
	"github.com/fdistorted/gencoder/encodings/piglatin"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"
)

func TestNumbers(t *testing.T) {
	assert(numbers.Encode("hello World!") == "h2ll4 W4rld!")
	assert(numbers.Decode("h2ll4 W4rld!") == "hello World!")
}

func TestPiglatin(t *testing.T) {
	//fmt.Printf("string %q\n",piglatin.Encode("Hello Mr. Fox! How are you?" )  )
	assert(piglatin.Encode("Hello Mr. Fox! How are you?") == "elloHay Mray. oxFay! owHay areyay ouyay?")
}

func assert(o bool) {
	if !o {
		fmt.Printf("\n%c[35m%s%c[0m\n\n", 27, __getRecentLine(), 27)
		os.Exit(1)
	}
}

func __getRecentLine() string {
	_, file, line, _ := runtime.Caller(2)
	buf, _ := ioutil.ReadFile(file)
	code := strings.TrimSpace(strings.Split(string(buf), "\n")[line-1])
	return fmt.Sprintf("%v:%d\n%s", path.Base(file), line, code)
}
