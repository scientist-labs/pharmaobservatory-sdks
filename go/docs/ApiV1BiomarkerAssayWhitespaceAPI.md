# \ApiV1BiomarkerAssayWhitespaceAPI

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1BiomarkerAssayWhitespaceGet**](ApiV1BiomarkerAssayWhitespaceAPI.md#ApiV1BiomarkerAssayWhitespaceGet) | **Get** /api/v1/biomarker_assay_whitespace | index



## ApiV1BiomarkerAssayWhitespaceGet

> ApiV1BiomarkerAssayWhitespaceGet200Response ApiV1BiomarkerAssayWhitespaceGet(ctx).Bucket(bucket).Gene(gene).GeneEvidenceTier(geneEvidenceTier).GeneSupplyStatus(geneSupplyStatus).Limit(limit).Q(q).Execute()

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
	bucket := "served" // string |  (optional)
	gene := "A1BG" // string |  (optional)
	geneEvidenceTier := "literature_only" // string |  (optional)
	geneSupplyStatus := "declared_positive" // string |  (optional)
	limit := int32(700) // int32 |  (optional)
	q := "EGFR" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiV1BiomarkerAssayWhitespaceAPI.ApiV1BiomarkerAssayWhitespaceGet(context.Background()).Bucket(bucket).Gene(gene).GeneEvidenceTier(geneEvidenceTier).GeneSupplyStatus(geneSupplyStatus).Limit(limit).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiV1BiomarkerAssayWhitespaceAPI.ApiV1BiomarkerAssayWhitespaceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1BiomarkerAssayWhitespaceGet`: ApiV1BiomarkerAssayWhitespaceGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ApiV1BiomarkerAssayWhitespaceAPI.ApiV1BiomarkerAssayWhitespaceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1BiomarkerAssayWhitespaceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucket** | **string** |  | 
 **gene** | **string** |  | 
 **geneEvidenceTier** | **string** |  | 
 **geneSupplyStatus** | **string** |  | 
 **limit** | **int32** |  | 
 **q** | **string** |  | 

### Return type

[**ApiV1BiomarkerAssayWhitespaceGet200Response**](ApiV1BiomarkerAssayWhitespaceGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

