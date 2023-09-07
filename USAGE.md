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
                "corrupti",
            },
            Encoding: []string{
                "provident",
            },
            Files: [][]byte{
                []byte("distinctio"),
            },
            GzUncompressedContentType: testingtesting.String("quibusdam"),
            HiResModelName: []string{
                "unde",
            },
            OcrLanguages: []string{
                "nulla",
            },
            OutputFormat: testingtesting.String("corrupti"),
            PdfInferTableStructure: []string{
                "illum",
            },
            Strategy: []string{
                "vel",
            },
            XMLKeepTags: []string{
                "error",
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