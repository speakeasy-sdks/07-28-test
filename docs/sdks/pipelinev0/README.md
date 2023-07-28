# PipelineV0

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
	"github.com/speakeasy-sdks/07-28-test"
	"github.com/speakeasy-sdks/07-28-test/pkg/models/operations"
	"github.com/speakeasy-sdks/07-28-test/pkg/models/shared"
)

func main() {
    s := testingtesting.New()

    ctx := context.Background()
    res, err := s.PipelineV0.Build(ctx, operations.Pipeline1GeneralV0GeneralPostRequest{
        PipelineBodyV0: &shared.PipelineBodyV0{
            Coordinates: []string{
                "ipsam",
            },
            Encoding: []string{
                "sapiente",
                "quo",
                "odit",
                "at",
            },
            Files: [][]byte{
                []byte("maiores"),
                []byte("molestiae"),
                []byte("quod"),
                []byte("quod"),
            },
            GzUncompressedContentType: testingtesting.String("esse"),
            HiResModelName: []string{
                "porro",
                "dolorum",
                "dicta",
            },
            OcrLanguages: []string{
                "officia",
                "occaecati",
                "fugit",
            },
            OutputFormat: testingtesting.String("deleniti"),
            PdfInferTableStructure: []string{
                "optio",
                "totam",
                "beatae",
                "commodi",
            },
            Strategy: []string{
                "modi",
                "qui",
            },
            XMLKeepTags: []string{
                "cum",
                "esse",
                "ipsum",
                "excepturi",
            },
        },
        UnstructuredAPIKey: testingtesting.String("aspernatur"),
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

