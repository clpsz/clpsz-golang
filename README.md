# clpsz-golang
my golang library


## how to use

```shell
go mot init <youmodname>
go get github.com/clpsz/clpsz-golang
```

```go
package main

import (
	"fmt"
	"github.com/clpsz/clpsz-golang/pkg"
)

func main() {
	fmt.Println(pkg.GetRandomString(10))
}
```
