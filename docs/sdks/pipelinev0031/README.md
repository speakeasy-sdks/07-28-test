# PipelineV0031

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
    res, err := s.PipelineV0031.Build(ctx, operations.Pipeline1GeneralV0031GeneralPostRequest{
        PipelineBodyV0031: &shared.PipelineBodyV0031{
            Coordinates: []string{
                "ipsam",
            },
            Encoding: []string{
                "repellendus",
            },
            Files: [][]byte{
                []byte("sapiente"),
            },
            GzUncompressedContentType: seven28test.String("quo"),
            HiResModelName: []string{
                "odit",
            },
            OcrLanguages: []string{
                "at",
            },
            OutputFormat: seven28test.String("at"),
            PdfInferTableStructure: []string{
                "maiores",
            },
            Strategy: []string{
                "molestiae",
            },
            XMLKeepTags: []string{
                "quod",
            },
        },
        UnstructuredAPIKey: seven28test.String("quod"),
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.Pipeline1GeneralV0031GeneralPostRequest](../../models/operations/pipeline1generalv0031generalpostrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.Pipeline1GeneralV0031GeneralPostResponse](../../models/operations/pipeline1generalv0031generalpostresponse.md), error**

