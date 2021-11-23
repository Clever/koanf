package main

import (
	"fmt"

	"github.com/Clever/koanf"
	"github.com/Clever/koanf/parsers/json"
	"github.com/Clever/koanf/providers/rawbytes"
)

// Global koanf instance. Use . as the key path delimiter. This can be / or anything.
var k = koanf.New(".")

func main() {
	b := []byte(`{"type": "rawbytes", "parent1": {"child1": {"type": "rawbytes"}}}`)
	k.Load(rawbytes.Provider(b), json.Parser())
	fmt.Println("type is = ", k.String("parent1.child1.type"))
}
