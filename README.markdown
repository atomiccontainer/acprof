# acprof

Container profiling & introspection.

```Go
import "github.com/atomiccontainer/acprof"
```

aclog detects information about the running container and its runtime. The aim is to provide an inventory of the executing container and its runtime for debugging purposes.

## Usage

```Go
package main

import (
	"fmt"

	"github.com/atomiccontainer/acprof"
)

func main() {
	i := acprof.New()

	fmt.Println(i.JSON())
}
```
