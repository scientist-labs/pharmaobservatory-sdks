# \ApiV1ProgramCatalystAPI

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1LicensingCandidatesGet**](ApiV1ProgramCatalystAPI.md#ApiV1LicensingCandidatesGet) | **Get** /api/v1/licensing_candidates | licensing_candidates
[**ApiV1OrgPipelineStagesGet**](ApiV1ProgramCatalystAPI.md#ApiV1OrgPipelineStagesGet) | **Get** /api/v1/org_pipeline_stages | pipeline_stages
[**ApiV1PatentCliffExposureGet**](ApiV1ProgramCatalystAPI.md#ApiV1PatentCliffExposureGet) | **Get** /api/v1/patent_cliff_exposure | patent_cliff_exposure
[**ApiV1ProgramCatalystsGet**](ApiV1ProgramCatalystAPI.md#ApiV1ProgramCatalystsGet) | **Get** /api/v1/program_catalysts | catalysts



## ApiV1LicensingCandidatesGet

> ApiV1LicensingCandidatesGet200Response ApiV1LicensingCandidatesGet(ctx).ConditionId(conditionId).Execute()

licensing_candidates

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
	conditionId := "D003920" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiV1ProgramCatalystAPI.ApiV1LicensingCandidatesGet(context.Background()).ConditionId(conditionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiV1ProgramCatalystAPI.ApiV1LicensingCandidatesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1LicensingCandidatesGet`: ApiV1LicensingCandidatesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ApiV1ProgramCatalystAPI.ApiV1LicensingCandidatesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1LicensingCandidatesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conditionId** | **string** |  | 

### Return type

[**ApiV1LicensingCandidatesGet200Response**](ApiV1LicensingCandidatesGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1OrgPipelineStagesGet

> ApiV1OrgPipelineStagesGet200Response ApiV1OrgPipelineStagesGet(ctx).CompanyId(companyId).Execute()

pipeline_stages

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
	companyId := "edgar:0001048477" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiV1ProgramCatalystAPI.ApiV1OrgPipelineStagesGet(context.Background()).CompanyId(companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiV1ProgramCatalystAPI.ApiV1OrgPipelineStagesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1OrgPipelineStagesGet`: ApiV1OrgPipelineStagesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ApiV1ProgramCatalystAPI.ApiV1OrgPipelineStagesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1OrgPipelineStagesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** |  | 

### Return type

[**ApiV1OrgPipelineStagesGet200Response**](ApiV1OrgPipelineStagesGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1PatentCliffExposureGet

> ApiV1PatentCliffExposureGet200Response ApiV1PatentCliffExposureGet(ctx).CompanyIds(companyIds).DrugIds(drugIds).Execute()

patent_cliff_exposure

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
	companyIds := []string{"Inner_example"} // []string |  (optional)
	drugIds := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiV1ProgramCatalystAPI.ApiV1PatentCliffExposureGet(context.Background()).CompanyIds(companyIds).DrugIds(drugIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiV1ProgramCatalystAPI.ApiV1PatentCliffExposureGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1PatentCliffExposureGet`: ApiV1PatentCliffExposureGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ApiV1ProgramCatalystAPI.ApiV1PatentCliffExposureGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1PatentCliffExposureGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyIds** | **[]string** |  | 
 **drugIds** | **[]string** |  | 

### Return type

[**ApiV1PatentCliffExposureGet200Response**](ApiV1PatentCliffExposureGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1ProgramCatalystsGet

> ApiV1ProgramCatalystsGet200Response ApiV1ProgramCatalystsGet(ctx).CompanyIds(companyIds).DrugIds(drugIds).Execute()

catalysts

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
	companyIds := []string{"Inner_example"} // []string |  (optional)
	drugIds := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiV1ProgramCatalystAPI.ApiV1ProgramCatalystsGet(context.Background()).CompanyIds(companyIds).DrugIds(drugIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiV1ProgramCatalystAPI.ApiV1ProgramCatalystsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1ProgramCatalystsGet`: ApiV1ProgramCatalystsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ApiV1ProgramCatalystAPI.ApiV1ProgramCatalystsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ProgramCatalystsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyIds** | **[]string** |  | 
 **drugIds** | **[]string** |  | 

### Return type

[**ApiV1ProgramCatalystsGet200Response**](ApiV1ProgramCatalystsGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

