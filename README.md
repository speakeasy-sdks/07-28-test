# github.com/speakeasy-sdks/07-28-test

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/07-28-test
```
<!-- End SDK Installation -->

## SDK Example Usage
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

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [PipelineV0](docs/sdks/pipelinev0/README.md)

* [Build](docs/sdks/pipelinev0/README.md#build) - Pipeline 1

### [PipelineV0031](docs/sdks/pipelinev0031/README.md)

* [Build](docs/sdks/pipelinev0031/README.md#build) - Pipeline 1
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
