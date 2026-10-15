# \ApiV1BiomarkerAssayScreenAPI

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1BiomarkerAssayScreenGet**](ApiV1BiomarkerAssayScreenAPI.md#ApiV1BiomarkerAssayScreenGet) | **Get** /api/v1/biomarker_assay_screen | index



## ApiV1BiomarkerAssayScreenGet

> ApiV1BiomarkerAssayScreenGet200Response ApiV1BiomarkerAssayScreenGet(ctx).AssayProviderCompany(assayProviderCompany).DocumentId(documentId).Gene(gene).GeneResolutionStatus(geneResolutionStatus).Pmid(pmid).ValidationTypeMentioned(validationTypeMentioned).Execute()

index

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
	assayProviderCompany := "Roche" // string |  (optional)
	documentId := "doc-pubmed-biomarker-0002" // string |  (optional)
	gene := "BRCA2" // string |  (optional)
	geneResolutionStatus := "unresolved" // string |  (optional)
	pmid := int32(20349101) // int32 |  (optional)
	validationTypeMentioned := "clinical" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiV1BiomarkerAssayScreenAPI.ApiV1BiomarkerAssayScreenGet(context.Background()).AssayProviderCompany(assayProviderCompany).DocumentId(documentId).Gene(gene).GeneResolutionStatus(geneResolutionStatus).Pmid(pmid).ValidationTypeMentioned(validationTypeMentioned).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiV1BiomarkerAssayScreenAPI.ApiV1BiomarkerAssayScreenGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1BiomarkerAssayScreenGet`: ApiV1BiomarkerAssayScreenGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ApiV1BiomarkerAssayScreenAPI.ApiV1BiomarkerAssayScreenGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1BiomarkerAssayScreenGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assayProviderCompany** | **string** |  | 
 **documentId** | **string** |  | 
 **gene** | **string** |  | 
 **geneResolutionStatus** | **string** |  | 
 **pmid** | **int32** |  | 
 **validationTypeMentioned** | **string** |  | 

### Return type

[**ApiV1BiomarkerAssayScreenGet200Response**](ApiV1BiomarkerAssayScreenGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

