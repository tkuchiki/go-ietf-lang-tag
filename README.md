# go-ietf-lang-tag
IETF language tag Utility for Golang

# Examples

```golang
package main

import (
	"fmt"
	"github.com/tkuchiki/go-ietf-lang-tag"
)

func main() {
	fmt.Println(ietflt.GetISO6391("af-NA"))                 // af
	fmt.Println(ietflt.GetLangTags("af"))                   // [af af-NA af-ZA]
	fmt.Println(ietflt.GetISO31661Alpha2("az-Cyrl-AZ"))     // AZ
	fmt.Println(ietflt.GetISO31661Alpha2("ca-ES-VALENCIA")) // ES
}
```
