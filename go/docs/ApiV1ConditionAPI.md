# \ApiV1ConditionAPI

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1ConditionsConditionIdBurdenVsEvidenceGet**](ApiV1ConditionAPI.md#ApiV1ConditionsConditionIdBurdenVsEvidenceGet) | **Get** /api/v1/conditions/{condition_id}/burden_vs_evidence | burden_vs_evidence
[**ApiV1ConditionsConditionIdEvidenceFreshnessGet**](ApiV1ConditionAPI.md#ApiV1ConditionsConditionIdEvidenceFreshnessGet) | **Get** /api/v1/conditions/{condition_id}/evidence_freshness | evidence_freshness
[**ApiV1ConditionsConditionIdSponsorLandscapeGet**](ApiV1ConditionAPI.md#ApiV1ConditionsConditionIdSponsorLandscapeGet) | **Get** /api/v1/conditions/{condition_id}/sponsor_landscape | sponsor_landscape



## ApiV1ConditionsConditionIdBurdenVsEvidenceGet

> ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response ApiV1ConditionsConditionIdBurdenVsEvidenceGet(ctx, conditionId).Execute()

burden_vs_evidence

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/scientist-labs/pharmaobservatory-sdks/go"
)

func main() {
	conditionId := "D020521" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiV1ConditionAPI.ApiV1ConditionsConditionIdBurdenVsEvidenceGet(context.Background(), conditionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiV1ConditionAPI.ApiV1ConditionsConditionIdBurdenVsEvidenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1ConditionsConditionIdBurdenVsEvidenceGet`: ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ApiV1ConditionAPI.ApiV1ConditionsConditionIdBurdenVsEvidenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conditionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ConditionsConditionIdBurdenVsEvidenceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response**](ApiV1ConditionsConditionIdBurdenVsEvidenceGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1ConditionsConditionIdEvidenceFreshnessGet

> ApiV1ConditionsConditionIdEvidenceFreshnessGet200Response ApiV1ConditionsConditionIdEvidenceFreshnessGet(ctx, conditionId).Execute()

evidence_freshness

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/scientist-labs/pharmaobservatory-sdks/go"
)

func main() {
	conditionId := "D064419" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiV1ConditionAPI.ApiV1ConditionsConditionIdEvidenceFreshnessGet(context.Background(), conditionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiV1ConditionAPI.ApiV1ConditionsConditionIdEvidenceFreshnessGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1ConditionsConditionIdEvidenceFreshnessGet`: ApiV1ConditionsConditionIdEvidenceFreshnessGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ApiV1ConditionAPI.ApiV1ConditionsConditionIdEvidenceFreshnessGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conditionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ConditionsConditionIdEvidenceFreshnessGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiV1ConditionsConditionIdEvidenceFreshnessGet200Response**](ApiV1ConditionsConditionIdEvidenceFreshnessGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1ConditionsConditionIdSponsorLandscapeGet

> ApiV1ConditionsConditionIdSponsorLandscapeGet200Response ApiV1ConditionsConditionIdSponsorLandscapeGet(ctx, conditionId).Execute()

sponsor_landscape

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/scientist-labs/pharmaobservatory-sdks/go"
)

func main() {
	conditionId := "D064419" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiV1ConditionAPI.ApiV1ConditionsConditionIdSponsorLandscapeGet(context.Background(), conditionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiV1ConditionAPI.ApiV1ConditionsConditionIdSponsorLandscapeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1ConditionsConditionIdSponsorLandscapeGet`: ApiV1ConditionsConditionIdSponsorLandscapeGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ApiV1ConditionAPI.ApiV1ConditionsConditionIdSponsorLandscapeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conditionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ConditionsConditionIdSponsorLandscapeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiV1ConditionsConditionIdSponsorLandscapeGet200Response**](ApiV1ConditionsConditionIdSponsorLandscapeGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

