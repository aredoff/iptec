## Installation

```bash
go get -u github.com/aredoff/iptec/plugins/asn
```

## API

See [godoc](https://godoc.org/github.com/aredoff/iptec/plugins/asn) reference.

## Example

```go
package main

import (
	"github.com/aredoff/iptec"
	"github.com/aredoff/iptec/plugins/asn"
)

func main() {
	a := iptec.New()
	defer a.Close()
	a.Use(asn.New())
	a.Activate()
	a.Collect()
}
```

## License

[MIT](https://choosealicense.com/licenses/mit/)