// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package seven28test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/07-28-test/pkg/models/operations"
	"github.com/speakeasy-sdks/07-28-test/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/07-28-test/pkg/models/shared"
	"github.com/speakeasy-sdks/07-28-test/pkg/utils"
	"io"
	"net/http"
	"strings"
)

type pipelineV0 struct {
	sdkConfiguration sdkConfiguration
}

func newPipelineV0(sdkConfig sdkConfiguration) *pipelineV0 {
	return &pipelineV0{
		sdkConfiguration: sdkConfig,
	}
}

// Build - Pipeline 1
func (s *pipelineV0) Build(ctx context.Context, request operations.Pipeline1GeneralV0GeneralPostRequest) (*operations.Pipeline1GeneralV0GeneralPostResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/general/v0/general"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, true, "PipelineBodyV0", "multipart", `request:"mediaType=multipart/form-data"`)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.DefaultClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.Pipeline1GeneralV0GeneralPostResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))
	switch {
	case httpRes.StatusCode == 200:
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.HTTPValidationError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.HTTPValidationError = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}
