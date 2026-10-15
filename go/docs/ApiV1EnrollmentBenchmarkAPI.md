# \ApiV1EnrollmentBenchmarkAPI

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1EnrollmentBenchmarkGet**](ApiV1EnrollmentBenchmarkAPI.md#ApiV1EnrollmentBenchmarkGet) | **Get** /api/v1/enrollment_benchmark | show



## ApiV1EnrollmentBenchmarkGet

> ApiV1EnrollmentBenchmarkGet200Response ApiV1EnrollmentBenchmarkGet(ctx).ConditionId(conditionId).Phase(phase).Execute()

show

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
	phase := "PHASE2" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiV1EnrollmentBenchmarkAPI.ApiV1EnrollmentBenchmarkGet(context.Background()).ConditionId(conditionId).Phase(phase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiV1EnrollmentBenchmarkAPI.ApiV1EnrollmentBenchmarkGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1EnrollmentBenchmarkGet`: ApiV1EnrollmentBenchmarkGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ApiV1EnrollmentBenchmarkAPI.ApiV1EnrollmentBenchmarkGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1EnrollmentBenchmarkGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conditionId** | **string** |  | 
 **phase** | **string** |  | 

### Return type

[**ApiV1EnrollmentBenchmarkGet200Response**](ApiV1EnrollmentBenchmarkGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

