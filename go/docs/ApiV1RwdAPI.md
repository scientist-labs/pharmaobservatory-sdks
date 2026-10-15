# \ApiV1RwdAPI

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1RwdConductorTrackRecordGet**](ApiV1RwdAPI.md#ApiV1RwdConductorTrackRecordGet) | **Get** /api/v1/rwd_conductor_track_record | conductor_track_record
[**ApiV1RwdStudyScreenGet**](ApiV1RwdAPI.md#ApiV1RwdStudyScreenGet) | **Get** /api/v1/rwd_study_screen | study_screen



## ApiV1RwdConductorTrackRecordGet

> ApiV1RwdConductorTrackRecordGet200Response ApiV1RwdConductorTrackRecordGet(ctx).CanonicalOrgId(canonicalOrgId).HmaEmaId(hmaEmaId).InstitutionName(institutionName).Execute()

conductor_track_record

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
	canonicalOrgId := "ror:rareres01" // string |  (optional)
	hmaEmaId := int32(900000005) // int32 |  (optional)
	institutionName := "Seedmark Consulting" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiV1RwdAPI.ApiV1RwdConductorTrackRecordGet(context.Background()).CanonicalOrgId(canonicalOrgId).HmaEmaId(hmaEmaId).InstitutionName(institutionName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiV1RwdAPI.ApiV1RwdConductorTrackRecordGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1RwdConductorTrackRecordGet`: ApiV1RwdConductorTrackRecordGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ApiV1RwdAPI.ApiV1RwdConductorTrackRecordGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RwdConductorTrackRecordGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **canonicalOrgId** | **string** |  | 
 **hmaEmaId** | **int32** |  | 
 **institutionName** | **string** |  | 

### Return type

[**ApiV1RwdConductorTrackRecordGet200Response**](ApiV1RwdConductorTrackRecordGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1RwdStudyScreenGet

> ApiV1RwdStudyScreenGet200Response ApiV1RwdStudyScreenGet(ctx).Condition(condition).Mandate(mandate).Execute()

study_screen

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
	condition := "Cataract" // string |  (optional)
	mandate := "mandated" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiV1RwdAPI.ApiV1RwdStudyScreenGet(context.Background()).Condition(condition).Mandate(mandate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiV1RwdAPI.ApiV1RwdStudyScreenGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1RwdStudyScreenGet`: ApiV1RwdStudyScreenGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ApiV1RwdAPI.ApiV1RwdStudyScreenGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RwdStudyScreenGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **condition** | **string** |  | 
 **mandate** | **string** |  | 

### Return type

[**ApiV1RwdStudyScreenGet200Response**](ApiV1RwdStudyScreenGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

