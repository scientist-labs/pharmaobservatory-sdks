# \ApiV1TrialAPI

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1TrialsDiscontinuationReasonsGet**](ApiV1TrialAPI.md#ApiV1TrialsDiscontinuationReasonsGet) | **Get** /api/v1/trials/discontinuation-reasons | discontinuation_reasons
[**ApiV1TrialsEligibilityGeneralizabilityScanGet**](ApiV1TrialAPI.md#ApiV1TrialsEligibilityGeneralizabilityScanGet) | **Get** /api/v1/trials/eligibility-generalizability-scan | eligibility_generalizability_scan
[**ApiV1TrialsHistoryAndPredictionGet**](ApiV1TrialAPI.md#ApiV1TrialsHistoryAndPredictionGet) | **Get** /api/v1/trials/history-and-prediction | history_and_prediction
[**ApiV1TrialsRecordsGet**](ApiV1TrialAPI.md#ApiV1TrialsRecordsGet) | **Get** /api/v1/trials/records | records
[**ApiV1TrialsReportingCompletenessGet**](ApiV1TrialAPI.md#ApiV1TrialsReportingCompletenessGet) | **Get** /api/v1/trials/reporting-completeness | reporting_completeness



## ApiV1TrialsDiscontinuationReasonsGet

> ApiV1TrialsDiscontinuationReasonsGet200Response ApiV1TrialsDiscontinuationReasonsGet(ctx).ConditionId(conditionId).Execute()

discontinuation_reasons

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
	conditionId := "D064419" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiV1TrialAPI.ApiV1TrialsDiscontinuationReasonsGet(context.Background()).ConditionId(conditionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiV1TrialAPI.ApiV1TrialsDiscontinuationReasonsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TrialsDiscontinuationReasonsGet`: ApiV1TrialsDiscontinuationReasonsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ApiV1TrialAPI.ApiV1TrialsDiscontinuationReasonsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TrialsDiscontinuationReasonsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conditionId** | **string** |  | 

### Return type

[**ApiV1TrialsDiscontinuationReasonsGet200Response**](ApiV1TrialsDiscontinuationReasonsGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1TrialsEligibilityGeneralizabilityScanGet

> ApiV1TrialsEligibilityGeneralizabilityScanGet200Response ApiV1TrialsEligibilityGeneralizabilityScanGet(ctx).ConditionId(conditionId).Limit(limit).NctIds(nctIds).Execute()

eligibility_generalizability_scan

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
	conditionId := "D064419" // string |  (optional)
	limit := int32(5) // int32 |  (optional)
	nctIds := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiV1TrialAPI.ApiV1TrialsEligibilityGeneralizabilityScanGet(context.Background()).ConditionId(conditionId).Limit(limit).NctIds(nctIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiV1TrialAPI.ApiV1TrialsEligibilityGeneralizabilityScanGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TrialsEligibilityGeneralizabilityScanGet`: ApiV1TrialsEligibilityGeneralizabilityScanGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ApiV1TrialAPI.ApiV1TrialsEligibilityGeneralizabilityScanGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TrialsEligibilityGeneralizabilityScanGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conditionId** | **string** |  | 
 **limit** | **int32** |  | 
 **nctIds** | **[]string** |  | 

### Return type

[**ApiV1TrialsEligibilityGeneralizabilityScanGet200Response**](ApiV1TrialsEligibilityGeneralizabilityScanGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1TrialsHistoryAndPredictionGet

> ApiV1TrialsHistoryAndPredictionGet200Response ApiV1TrialsHistoryAndPredictionGet(ctx).NctId(nctId).Execute()

history_and_prediction

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
	nctId := "NCT05734196" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiV1TrialAPI.ApiV1TrialsHistoryAndPredictionGet(context.Background()).NctId(nctId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiV1TrialAPI.ApiV1TrialsHistoryAndPredictionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TrialsHistoryAndPredictionGet`: ApiV1TrialsHistoryAndPredictionGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ApiV1TrialAPI.ApiV1TrialsHistoryAndPredictionGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TrialsHistoryAndPredictionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nctId** | **string** |  | 

### Return type

[**ApiV1TrialsHistoryAndPredictionGet200Response**](ApiV1TrialsHistoryAndPredictionGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1TrialsRecordsGet

> ApiV1TrialsRecordsGet200Response ApiV1TrialsRecordsGet(ctx).IncludeExtracted(includeExtracted).NctIds(nctIds).Execute()

records

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
	includeExtracted := true // bool |  (optional)
	nctIds := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiV1TrialAPI.ApiV1TrialsRecordsGet(context.Background()).IncludeExtracted(includeExtracted).NctIds(nctIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiV1TrialAPI.ApiV1TrialsRecordsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TrialsRecordsGet`: ApiV1TrialsRecordsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ApiV1TrialAPI.ApiV1TrialsRecordsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TrialsRecordsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeExtracted** | **bool** |  | 
 **nctIds** | **[]string** |  | 

### Return type

[**ApiV1TrialsRecordsGet200Response**](ApiV1TrialsRecordsGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1TrialsReportingCompletenessGet

> ApiV1TrialsReportingCompletenessGet200Response ApiV1TrialsReportingCompletenessGet(ctx).ConditionId(conditionId).Execute()

reporting_completeness

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
	conditionId := "D010661" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiV1TrialAPI.ApiV1TrialsReportingCompletenessGet(context.Background()).ConditionId(conditionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiV1TrialAPI.ApiV1TrialsReportingCompletenessGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TrialsReportingCompletenessGet`: ApiV1TrialsReportingCompletenessGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ApiV1TrialAPI.ApiV1TrialsReportingCompletenessGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TrialsReportingCompletenessGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conditionId** | **string** |  | 

### Return type

[**ApiV1TrialsReportingCompletenessGet200Response**](ApiV1TrialsReportingCompletenessGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

