<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	seven28test "github.com/speakeasy-sdks/07-28-test/v3"
	"github.com/speakeasy-sdks/07-28-test/v3/pkg/models/operations"
	"log"
)

func main() {
	s := seven28test.New()

	ctx := context.Background()
	res, err := s.PipelineV0031.Build(ctx, operations.Pipeline1GeneralV0031GeneralPostRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->