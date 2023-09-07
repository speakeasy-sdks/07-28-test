// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/07-28-test/pkg/models/shared"
	"net/http"
)

type Pipeline1GeneralV0GeneralPostRequest struct {
	PipelineBodyV0     *shared.PipelineBodyV0 `request:"mediaType=multipart/form-data"`
	UnstructuredAPIKey *string                `header:"style=simple,explode=false,name=unstructured-api-key"`
}

func (o *Pipeline1GeneralV0GeneralPostRequest) GetPipelineBodyV0() *shared.PipelineBodyV0 {
	if o == nil {
		return nil
	}
	return o.PipelineBodyV0
}

func (o *Pipeline1GeneralV0GeneralPostRequest) GetUnstructuredAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.UnstructuredAPIKey
}

type Pipeline1GeneralV0GeneralPostResponse struct {
	ContentType string
	// Validation Error
	HTTPValidationError *shared.HTTPValidationError
	StatusCode          int
	RawResponse         *http.Response
}

func (o *Pipeline1GeneralV0GeneralPostResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *Pipeline1GeneralV0GeneralPostResponse) GetHTTPValidationError() *shared.HTTPValidationError {
	if o == nil {
		return nil
	}
	return o.HTTPValidationError
}

func (o *Pipeline1GeneralV0GeneralPostResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *Pipeline1GeneralV0GeneralPostResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}