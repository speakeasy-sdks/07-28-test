<!-- Start SDK Example Usage -->


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
                "provident",
                "distinctio",
                "quibusdam",
            },
            Encoding: []string{
                "nulla",
                "corrupti",
                "illum",
            },
            Files: [][]byte{
                []byte("error"),
                []byte("deserunt"),
            },
            GzUncompressedContentType: testingtesting.String("suscipit"),
            HiResModelName: []string{
                "magnam",
                "debitis",
            },
            OcrLanguages: []string{
                "delectus",
            },
            OutputFormat: testingtesting.String("tempora"),
            PdfInferTableStructure: []string{
                "molestiae",
                "minus",
            },
            Strategy: []string{
                "voluptatum",
                "iusto",
                "excepturi",
                "nisi",
            },
            XMLKeepTags: []string{
                "temporibus",
                "ab",
                "quis",
                "veritatis",
            },
        },
        UnstructuredAPIKey: testingtesting.String("deserunt"),
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