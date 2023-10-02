# PipelineV0
(*PipelineV0*)

### Available Operations

* [Build](#build) - Pipeline 1

## Build

Pipeline 1

### Example Usage

```go
package main

import(
	"context"
	"log"
	seven28test "github.com/speakeasy-sdks/07-28-test"
	"github.com/speakeasy-sdks/07-28-test/pkg/models/operations"
	"github.com/speakeasy-sdks/07-28-test/pkg/models/shared"
)

func main() {
    s := seven28test.New()

    ctx := context.Background()
    res, err := s.PipelineV0.Build(ctx, operations.Pipeline1GeneralV0GeneralPostRequest{
        PipelineBodyV0: &shared.PipelineBodyV0{
            Coordinates: []string{
                "Practical",
            },
            Encoding: []string{
                "Gasoline",
            },
            Files: [][]byte{
                []byte("[eAhpDJhn'"),
            },
            GzUncompressedContentType: seven28test.String("Shoes"),
            HiResModelName: []string{
                "what",
            },
            OcrLanguages: []string{
                "whether",
            },
            OutputFormat: seven28test.String("Direct"),
            PdfInferTableStructure: []string{
                "gee",
            },
            Strategy: []string{
                "Metal",
            },
            XMLKeepTags: []string{
                "Avon",
            },
        },
        UnstructuredAPIKey: seven28test.String("Intranet synthesizing HTTP"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.Pipeline1GeneralV0GeneralPostRequest](../../models/operations/pipeline1generalv0generalpostrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.Pipeline1GeneralV0GeneralPostResponse](../../models/operations/pipeline1generalv0generalpostresponse.md), error**

