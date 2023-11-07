# PipelineV0031
(*.PipelineV0031*)

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
	seven28test "github.com/speakeasy-sdks/07-28-test/v2"
	"github.com/speakeasy-sdks/07-28-test/v2/pkg/models/operations"
	"github.com/speakeasy-sdks/07-28-test/v2/pkg/models/shared"
)

func main() {
    s := seven28test.New()

    ctx := context.Background()
    res, err := s.PipelineV0031.Build(ctx, operations.Pipeline1GeneralV0031GeneralPostRequest{
        PipelineBodyV0031: &shared.PipelineBodyV0031{
            Coordinates: []string{
                "string",
            },
            Encoding: []string{
                "string",
            },
            Files: []shared.PipelineBodyV0031Files{
                shared.PipelineBodyV0031Files{
                    Content: []byte("0x591E0BfdA7"),
                    FileName: "cab_touring_henry.mpg4",
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

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.Pipeline1GeneralV0031GeneralPostRequest](../../models/operations/pipeline1generalv0031generalpostrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.Pipeline1GeneralV0031GeneralPostResponse](../../models/operations/pipeline1generalv0031generalpostresponse.md), error**

