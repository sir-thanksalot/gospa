# GoSpa

Zero-dependency, standard library implementation of [`http.Handler`](https://golang.org/pkg/net/http/#Handler) for SPAs (pushState/client side routing).  
What it does:

- Serve any existing files under `rootPath`
- Serve `indexFile` otherwise

## Example

```golang
package main

import (
	"net/http"

	spa "github.com/sir-thanksalot/gospa"
)

func main() {
	spaHandler := spa.NewSpaHandlerAt("/", "dist", "index.html")
	http.Handle("/", spaHandler)
	http.ListenAndServe(":5000", nil)
}
```
