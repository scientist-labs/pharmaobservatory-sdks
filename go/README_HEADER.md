# pharmaobservatory (Go)

Go client for the [Pharma Observatory](https://pharmaobservatory.com) public REST API. Generated from the
[OpenAPI document](https://pharmaobservatory.com/api/v1/openapi.yaml); see the
[API reference](https://pharmaobservatory.com/api-reference) for what each endpoint returns.

## Install

```shell
go get github.com/scientist-labs/pharmaobservatory-sdks/go
```

The package name is `pharmaobservatory`. Versions are tagged `go/vX.Y.Z` because the module lives in the `go/` directory.

## A minimal authenticated call

The API takes a bearer token (see the `ApiBearerJwt` security scheme in the OpenAPI document for how to obtain one).

```go
package main

import (
	"context"
	"fmt"
	"os"

	pharmaobservatory "github.com/scientist-labs/pharmaobservatory-sdks/go"
)

func main() {
	ctx := context.WithValue(context.Background(), pharmaobservatory.ContextAccessToken, os.Getenv("PHARMAOBSERVATORY_TOKEN"))
	client := pharmaobservatory.NewAPIClient(pharmaobservatory.NewConfiguration())

	who, _, err := client.ApiV1BaseAPI.ApiV1WhoamiGet(ctx).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", who)
}
```

The default host is `https://pharmaobservatory.com`. The generated reference follows.

---

