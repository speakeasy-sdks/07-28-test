<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	seven28test "github.com/speakeasy-sdks/07-28-test"
	"github.com/speakeasy-sdks/07-28-test/pkg/models/operations"
	"github.com/speakeasy-sdks/07-28-test/pkg/models/shared"
	"log"
)

func main() {
	s := seven28test.New()

	ctx := context.Background()
	res, err := s.PipelineV0.Build(ctx, operations.Pipeline1GeneralV0GeneralPostRequest{
		PipelineBodyV0: &shared.PipelineBodyV0{
			Coordinates: []string{
				"string",
			},
			Encoding: []string{
				"string",
			},
			Files: []shared.PipelineBodyV0Files{
				shared.PipelineBodyV0Files{
					Content: []byte("9G&x$kc[eA"),
					Files:   "string",
				},
			},
			HiResModelName: []string{
				"string",
			},
			OcrLanguages: []string{
				"string",
			},
			PdfInferTableStructure: []string{
				"string",
			},
			Strategy: []string{
				"string",
			},
			XMLKeepTags: []string{
				"string",
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->