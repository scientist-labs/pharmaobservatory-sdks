# \ApiV1CandidateSiteAPI

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1CandidateSitesGet**](ApiV1CandidateSiteAPI.md#ApiV1CandidateSitesGet) | **Get** /api/v1/candidate_sites | Rank candidate clinical trial sites for a condition by composite capacity, activity, and competition scores



## ApiV1CandidateSitesGet

> ApiV1CandidateSitesGet200Response ApiV1CandidateSitesGet(ctx).ConditionId(conditionId).Country(country).ExcludeSaturated(excludeSaturated).Limit(limit).Execute()

Rank candidate clinical trial sites for a condition by composite capacity, activity, and competition scores

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
	conditionId := "D001523" // string |  (optional)
	country := "United States" // string |  (optional)
	excludeSaturated := true // bool |  (optional)
	limit := int32(500) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiV1CandidateSiteAPI.ApiV1CandidateSitesGet(context.Background()).ConditionId(conditionId).Country(country).ExcludeSaturated(excludeSaturated).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiV1CandidateSiteAPI.ApiV1CandidateSitesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1CandidateSitesGet`: ApiV1CandidateSitesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ApiV1CandidateSiteAPI.ApiV1CandidateSitesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1CandidateSitesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conditionId** | **string** |  | 
 **country** | **string** |  | 
 **excludeSaturated** | **bool** |  | 
 **limit** | **int32** |  | 

### Return type

[**ApiV1CandidateSitesGet200Response**](ApiV1CandidateSitesGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

