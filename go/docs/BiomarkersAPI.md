# \BiomarkersAPI

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BiomarkersBiomarkerBqpAssertions**](BiomarkersAPI.md#BiomarkersBiomarkerBqpAssertions) | **Get** /api/v1/biomarkers/{id}/biomarker_bqp_assertions | FDA biomarker qualification projects
[**BiomarkersBiomarkerCdxAssertions**](BiomarkersAPI.md#BiomarkersBiomarkerCdxAssertions) | **Get** /api/v1/biomarkers/{id}/biomarker_cdx_assertions | FDA companion diagnostic assertions
[**BiomarkersBiomarkerEvidence**](BiomarkersAPI.md#BiomarkersBiomarkerEvidence) | **Get** /api/v1/biomarkers/{id}/biomarker_evidence | Biomarker evidence
[**BiomarkersBiomarkerLiterature**](BiomarkersAPI.md#BiomarkersBiomarkerLiterature) | **Get** /api/v1/biomarkers/{id}/biomarker_literature | Trial-cited literature mentions
[**BiomarkersBiomarkerPublishingCompanies**](BiomarkersAPI.md#BiomarkersBiomarkerPublishingCompanies) | **Get** /api/v1/biomarkers/{id}/biomarker_publishing_companies | Companies publishing on a biomarker
[**BiomarkersBiomarkerVariantMentions**](BiomarkersAPI.md#BiomarkersBiomarkerVariantMentions) | **Get** /api/v1/biomarkers/{id}/biomarker_variant_mentions | Protein variant mentions
[**BiomarkersGet**](BiomarkersAPI.md#BiomarkersGet) | **Get** /api/v1/biomarkers/{id} | One biomarker
[**BiomarkersList**](BiomarkersAPI.md#BiomarkersList) | **Get** /api/v1/biomarkers | List biomarkers



## BiomarkersBiomarkerBqpAssertions

> BiomarkerBqpAssertionsResponse BiomarkersBiomarkerBqpAssertions(ctx, id).Limit(limit).Sort(sort).Execute()

FDA biomarker qualification projects

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
	id := "ESR1" // string | 
	limit := int32(2) // int32 |  (optional)
	sort := "route_key" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersBiomarkerBqpAssertions(context.Background(), id).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersBiomarkerBqpAssertions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersBiomarkerBqpAssertions`: BiomarkerBqpAssertionsResponse
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersBiomarkerBqpAssertions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersBiomarkerBqpAssertionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **sort** | **string** |  | 

### Return type

[**BiomarkerBqpAssertionsResponse**](BiomarkerBqpAssertionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersBiomarkerCdxAssertions

> BiomarkerCdxAssertionsResponse BiomarkersBiomarkerCdxAssertions(ctx, id).Limit(limit).Sort(sort).Execute()

FDA companion diagnostic assertions

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
	id := "ESR1" // string | 
	limit := int32(2) // int32 |  (optional)
	sort := "biomarker_id" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersBiomarkerCdxAssertions(context.Background(), id).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersBiomarkerCdxAssertions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersBiomarkerCdxAssertions`: BiomarkerCdxAssertionsResponse
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersBiomarkerCdxAssertions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersBiomarkerCdxAssertionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **sort** | **string** |  | 

### Return type

[**BiomarkerCdxAssertionsResponse**](BiomarkerCdxAssertionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersBiomarkerEvidence

> BiomarkerEvidenceResponse BiomarkersBiomarkerEvidence(ctx, id).Limit(limit).Sort(sort).Execute()

Biomarker evidence

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
	id := "ESR1" // string | 
	limit := int32(2) // int32 |  (optional)
	sort := "gene_symbol" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersBiomarkerEvidence(context.Background(), id).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersBiomarkerEvidence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersBiomarkerEvidence`: BiomarkerEvidenceResponse
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersBiomarkerEvidence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersBiomarkerEvidenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **sort** | **string** |  | 

### Return type

[**BiomarkerEvidenceResponse**](BiomarkerEvidenceResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersBiomarkerLiterature

> BiomarkerLiteratureResponse BiomarkersBiomarkerLiterature(ctx, id).Limit(limit).Sort(sort).Execute()

Trial-cited literature mentions

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
	id := "ESR1" // string | 
	limit := int32(2) // int32 |  (optional)
	sort := "pub_year" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersBiomarkerLiterature(context.Background(), id).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersBiomarkerLiterature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersBiomarkerLiterature`: BiomarkerLiteratureResponse
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersBiomarkerLiterature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersBiomarkerLiteratureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **sort** | **string** |  | 

### Return type

[**BiomarkerLiteratureResponse**](BiomarkerLiteratureResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersBiomarkerPublishingCompanies

> BiomarkerPublishingCompaniesResponse BiomarkersBiomarkerPublishingCompanies(ctx, id).Limit(limit).Sort(sort).Execute()

Companies publishing on a biomarker

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
	id := "ESR1" // string | 
	limit := int32(2) // int32 |  (optional)
	sort := "n_publications" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersBiomarkerPublishingCompanies(context.Background(), id).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersBiomarkerPublishingCompanies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersBiomarkerPublishingCompanies`: BiomarkerPublishingCompaniesResponse
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersBiomarkerPublishingCompanies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersBiomarkerPublishingCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **sort** | **string** |  | 

### Return type

[**BiomarkerPublishingCompaniesResponse**](BiomarkerPublishingCompaniesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersBiomarkerVariantMentions

> BiomarkerVariantMentionsResponse BiomarkersBiomarkerVariantMentions(ctx, id).Limit(limit).Sort(sort).Execute()

Protein variant mentions

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
	id := "ESR1" // string | 
	limit := int32(2) // int32 |  (optional)
	sort := "n_papers" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersBiomarkerVariantMentions(context.Background(), id).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersBiomarkerVariantMentions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersBiomarkerVariantMentions`: BiomarkerVariantMentionsResponse
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersBiomarkerVariantMentions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersBiomarkerVariantMentionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **sort** | **string** |  | 

### Return type

[**BiomarkerVariantMentionsResponse**](BiomarkerVariantMentionsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersGet

> BiomarkersRecord BiomarkersGet(ctx, id).Execute()

One biomarker

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
	id := "ESR1" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersGet`: BiomarkersRecord
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BiomarkersRecord**](BiomarkersRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersList

> BiomarkersResponse BiomarkersList(ctx).Limit(limit).Sort(sort).Execute()

List biomarkers

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
	limit := int32(2) // int32 |  (optional)
	sort := "evidence_tier" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersList(context.Background()).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersList`: BiomarkersResponse
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **sort** | **string** |  | 

### Return type

[**BiomarkersResponse**](BiomarkersResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

