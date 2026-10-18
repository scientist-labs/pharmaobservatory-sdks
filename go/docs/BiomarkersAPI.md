# \BiomarkersAPI

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BiomarkersAssayProviderLedger**](BiomarkersAPI.md#BiomarkersAssayProviderLedger) | **Get** /api/v1/biomarkers/{id}/assay_provider_ledger | Assay-provider ledger
[**BiomarkersCgi**](BiomarkersAPI.md#BiomarkersCgi) | **Get** /api/v1/biomarkers/{id}/cgi | Actionable biomarkers
[**BiomarkersCoverageByTraversal**](BiomarkersAPI.md#BiomarkersCoverageByTraversal) | **Get** /api/v1/biomarkers/{id}/coverage_by_traversal | Coverage by traversal
[**BiomarkersEvidenceOverTime**](BiomarkersAPI.md#BiomarkersEvidenceOverTime) | **Get** /api/v1/biomarkers/{id}/evidence_over_time | Evidence over time
[**BiomarkersFdaBqp**](BiomarkersAPI.md#BiomarkersFdaBqp) | **Get** /api/v1/biomarkers/{id}/fda_bqp | FDA biomarker qualification projects
[**BiomarkersFdaCdx**](BiomarkersAPI.md#BiomarkersFdaCdx) | **Get** /api/v1/biomarkers/{id}/fda_cdx | FDA companion diagnostic assertions
[**BiomarkersGet**](BiomarkersAPI.md#BiomarkersGet) | **Get** /api/v1/biomarkers/{id} | One biomarkers record
[**BiomarkersHubTierTables**](BiomarkersAPI.md#BiomarkersHubTierTables) | **Get** /api/v1/biomarkers/hub_tier_tables | Biomarker evidence, labeled by type
[**BiomarkersIdentity**](BiomarkersAPI.md#BiomarkersIdentity) | **Get** /api/v1/biomarkers/{id}/identity | Biomarker identity
[**BiomarkersList**](BiomarkersAPI.md#BiomarkersList) | **Get** /api/v1/biomarkers | Biomarkers
[**BiomarkersLiterature**](BiomarkersAPI.md#BiomarkersLiterature) | **Get** /api/v1/biomarkers/{id}/literature | Trial-cited literature mentions
[**BiomarkersOtAssociation**](BiomarkersAPI.md#BiomarkersOtAssociation) | **Get** /api/v1/biomarkers/{id}/ot_association | Target-disease associations
[**BiomarkersPopulationAlleleFrequency**](BiomarkersAPI.md#BiomarkersPopulationAlleleFrequency) | **Get** /api/v1/biomarkers/{id}/population_allele_frequency | Population allele frequency
[**BiomarkersPublishingOrg**](BiomarkersAPI.md#BiomarkersPublishingOrg) | **Get** /api/v1/biomarkers/{id}/publishing_org | Organizations publishing
[**BiomarkersRelatedBiomarkers**](BiomarkersAPI.md#BiomarkersRelatedBiomarkers) | **Get** /api/v1/biomarkers/{id}/related_biomarkers | Related biomarkers
[**BiomarkersRelatedConditions**](BiomarkersAPI.md#BiomarkersRelatedConditions) | **Get** /api/v1/biomarkers/{id}/related_conditions | Related conditions
[**BiomarkersRelatedOrganizations**](BiomarkersAPI.md#BiomarkersRelatedOrganizations) | **Get** /api/v1/biomarkers/{id}/related_organizations | Related organisations
[**BiomarkersRelations**](BiomarkersAPI.md#BiomarkersRelations) | **Get** /api/v1/biomarkers/{id}/relations | Literature-extracted relations
[**BiomarkersSampleScreening**](BiomarkersAPI.md#BiomarkersSampleScreening) | **Get** /api/v1/biomarkers/{id}/sample_screening | Human sample screening reference
[**BiomarkersTrialMention**](BiomarkersAPI.md#BiomarkersTrialMention) | **Get** /api/v1/biomarkers/{id}/trial_mention | Trial mentions
[**BiomarkersVariantMentions**](BiomarkersAPI.md#BiomarkersVariantMentions) | **Get** /api/v1/biomarkers/{id}/variant_mentions | Protein variant mentions
[**BiomarkersVendorSourcing**](BiomarkersAPI.md#BiomarkersVendorSourcing) | **Get** /api/v1/biomarkers/{id}/vendor_sourcing | Vendor sourcing
[**BiomarkersWhereItAppears**](BiomarkersAPI.md#BiomarkersWhereItAppears) | **Get** /api/v1/biomarkers/{id}/where_it_appears | Where it appears



## BiomarkersAssayProviderLedger

> BiomarkersAssayProviderLedger200Response BiomarkersAssayProviderLedger(ctx, id).Execute()

Assay-provider ledger

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
	id := "ALK" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersAssayProviderLedger(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersAssayProviderLedger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersAssayProviderLedger`: BiomarkersAssayProviderLedger200Response
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersAssayProviderLedger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersAssayProviderLedgerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BiomarkersAssayProviderLedger200Response**](BiomarkersAssayProviderLedger200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersCgi

> BiomarkersCgi200Response BiomarkersCgi(ctx, id).Execute()

Actionable biomarkers

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
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersCgi(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersCgi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersCgi`: BiomarkersCgi200Response
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersCgi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersCgiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BiomarkersCgi200Response**](BiomarkersCgi200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersCoverageByTraversal

> BiomarkersCoverageByTraversal200Response BiomarkersCoverageByTraversal(ctx, id).Execute()

Coverage by traversal

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
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersCoverageByTraversal(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersCoverageByTraversal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersCoverageByTraversal`: BiomarkersCoverageByTraversal200Response
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersCoverageByTraversal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersCoverageByTraversalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BiomarkersCoverageByTraversal200Response**](BiomarkersCoverageByTraversal200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersEvidenceOverTime

> BiomarkersEvidenceOverTime200Response BiomarkersEvidenceOverTime(ctx, id).Execute()

Evidence over time

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
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersEvidenceOverTime(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersEvidenceOverTime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersEvidenceOverTime`: BiomarkersEvidenceOverTime200Response
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersEvidenceOverTime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersEvidenceOverTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BiomarkersEvidenceOverTime200Response**](BiomarkersEvidenceOverTime200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersFdaBqp

> BiomarkersFdaBqp200Response BiomarkersFdaBqp(ctx, id).Execute()

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
	id := "bqp.safety-biomarker-to-be-used-with-traditional-ind.2516104f1f" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersFdaBqp(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersFdaBqp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersFdaBqp`: BiomarkersFdaBqp200Response
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersFdaBqp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersFdaBqpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BiomarkersFdaBqp200Response**](BiomarkersFdaBqp200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersFdaCdx

> BiomarkersFdaCdx200Response BiomarkersFdaCdx(ctx, id).Execute()

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
	id := "f.pd-l1" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersFdaCdx(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersFdaCdx``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersFdaCdx`: BiomarkersFdaCdx200Response
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersFdaCdx`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersFdaCdxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BiomarkersFdaCdx200Response**](BiomarkersFdaCdx200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersGet

> BiomarkersGet200Response BiomarkersGet(ctx, id).Execute()

One biomarkers record

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
	// response from `BiomarkersGet`: BiomarkersGet200Response
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

[**BiomarkersGet200Response**](BiomarkersGet200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersHubTierTables

> BiomarkersHubTierTables200Response BiomarkersHubTierTables(ctx).Execute()

Biomarker evidence, labeled by type

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersHubTierTables(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersHubTierTables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersHubTierTables`: BiomarkersHubTierTables200Response
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersHubTierTables`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersHubTierTablesRequest struct via the builder pattern


### Return type

[**BiomarkersHubTierTables200Response**](BiomarkersHubTierTables200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersIdentity

> BiomarkersIdentity200Response BiomarkersIdentity(ctx, id).Execute()

Biomarker identity

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
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersIdentity(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersIdentity`: BiomarkersIdentity200Response
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BiomarkersIdentity200Response**](BiomarkersIdentity200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersList

> BiomarkersList200Response BiomarkersList(ctx).Cursor(cursor).Dir(dir).Evidence(evidence).Facets(facets).Identity(identity).Kind(kind).Limit(limit).Q(q).Sort(sort).Validated(validated).Execute()

Biomarkers

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
	cursor := "eyJwIjoyLCJsIjozLCJmIjoic3BlYy1maW5nZXJwcmludCJ9" // string |  (optional)
	dir := "asc" // string |  (optional)
	evidence := int32(0) // int32 |  (optional)
	facets := true // bool |  (optional)
	identity := "resolved" // string |  (optional)
	kind := "gene" // string |  (optional)
	limit := int32(1) // int32 |  (optional)
	q := "BRAF" // string |  (optional)
	sort := "name" // string |  (optional)
	validated := int32(1) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersList(context.Background()).Cursor(cursor).Dir(dir).Evidence(evidence).Facets(facets).Identity(identity).Kind(kind).Limit(limit).Q(q).Sort(sort).Validated(validated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersList`: BiomarkersList200Response
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** |  | 
 **dir** | **string** |  | 
 **evidence** | **int32** |  | 
 **facets** | **bool** |  | 
 **identity** | **string** |  | 
 **kind** | **string** |  | 
 **limit** | **int32** |  | 
 **q** | **string** |  | 
 **sort** | **string** |  | 
 **validated** | **int32** |  | 

### Return type

[**BiomarkersList200Response**](BiomarkersList200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersLiterature

> BiomarkersLiterature200Response BiomarkersLiterature(ctx, id).Execute()

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
	id := "ERBB2" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersLiterature(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersLiterature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersLiterature`: BiomarkersLiterature200Response
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersLiterature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersLiteratureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BiomarkersLiterature200Response**](BiomarkersLiterature200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersOtAssociation

> BiomarkersOtAssociation200Response BiomarkersOtAssociation(ctx, id).Execute()

Target-disease associations

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
	id := "MMP9" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersOtAssociation(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersOtAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersOtAssociation`: BiomarkersOtAssociation200Response
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersOtAssociation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersOtAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BiomarkersOtAssociation200Response**](BiomarkersOtAssociation200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersPopulationAlleleFrequency

> BiomarkersPopulationAlleleFrequency200Response BiomarkersPopulationAlleleFrequency(ctx, id).Execute()

Population allele frequency

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
	id := "v.a1bg.p-a170t.a6cb074848" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersPopulationAlleleFrequency(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersPopulationAlleleFrequency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersPopulationAlleleFrequency`: BiomarkersPopulationAlleleFrequency200Response
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersPopulationAlleleFrequency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersPopulationAlleleFrequencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BiomarkersPopulationAlleleFrequency200Response**](BiomarkersPopulationAlleleFrequency200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersPublishingOrg

> BiomarkersPublishingOrg200Response BiomarkersPublishingOrg(ctx, id).Execute()

Organizations publishing

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
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersPublishingOrg(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersPublishingOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersPublishingOrg`: BiomarkersPublishingOrg200Response
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersPublishingOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersPublishingOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BiomarkersPublishingOrg200Response**](BiomarkersPublishingOrg200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersRelatedBiomarkers

> BiomarkersRelatedBiomarkers200Response BiomarkersRelatedBiomarkers(ctx, id).Execute()

Related biomarkers

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
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersRelatedBiomarkers(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersRelatedBiomarkers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersRelatedBiomarkers`: BiomarkersRelatedBiomarkers200Response
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersRelatedBiomarkers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersRelatedBiomarkersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BiomarkersRelatedBiomarkers200Response**](BiomarkersRelatedBiomarkers200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersRelatedConditions

> BiomarkersRelatedConditions200Response BiomarkersRelatedConditions(ctx, id).Execute()

Related conditions

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
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersRelatedConditions(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersRelatedConditions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersRelatedConditions`: BiomarkersRelatedConditions200Response
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersRelatedConditions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersRelatedConditionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BiomarkersRelatedConditions200Response**](BiomarkersRelatedConditions200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersRelatedOrganizations

> BiomarkersRelatedOrganizations200Response BiomarkersRelatedOrganizations(ctx, id).Execute()

Related organisations

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
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersRelatedOrganizations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersRelatedOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersRelatedOrganizations`: BiomarkersRelatedOrganizations200Response
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersRelatedOrganizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersRelatedOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BiomarkersRelatedOrganizations200Response**](BiomarkersRelatedOrganizations200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersRelations

> BiomarkersRelations200Response BiomarkersRelations(ctx, id).Chemicals(chemicals).Execute()

Literature-extracted relations

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
	id := "BRAF" // string | 
	chemicals := "x" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersRelations(context.Background(), id).Chemicals(chemicals).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersRelations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersRelations`: BiomarkersRelations200Response
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersRelations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersRelationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chemicals** | **string** |  | 

### Return type

[**BiomarkersRelations200Response**](BiomarkersRelations200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersSampleScreening

> BiomarkersSampleScreening200Response BiomarkersSampleScreening(ctx, id).PositiveSamples(positiveSamples).Execute()

Human sample screening reference

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
	id := "BRAF" // string | 
	positiveSamples := int32(20) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersSampleScreening(context.Background(), id).PositiveSamples(positiveSamples).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersSampleScreening``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersSampleScreening`: BiomarkersSampleScreening200Response
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersSampleScreening`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersSampleScreeningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **positiveSamples** | **int32** |  | 

### Return type

[**BiomarkersSampleScreening200Response**](BiomarkersSampleScreening200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersTrialMention

> BiomarkersTrialMention200Response BiomarkersTrialMention(ctx, id).Execute()

Trial mentions

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
	id := "ALB" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersTrialMention(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersTrialMention``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersTrialMention`: BiomarkersTrialMention200Response
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersTrialMention`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersTrialMentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BiomarkersTrialMention200Response**](BiomarkersTrialMention200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersVariantMentions

> BiomarkersVariantMentions200Response BiomarkersVariantMentions(ctx, id).Execute()

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
	id := "BRAF" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersVariantMentions(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersVariantMentions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersVariantMentions`: BiomarkersVariantMentions200Response
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersVariantMentions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersVariantMentionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BiomarkersVariantMentions200Response**](BiomarkersVariantMentions200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersVendorSourcing

> BiomarkersVendorSourcing200Response BiomarkersVendorSourcing(ctx, id).Execute()

Vendor sourcing

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
	id := "ALB" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersVendorSourcing(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersVendorSourcing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersVendorSourcing`: BiomarkersVendorSourcing200Response
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersVendorSourcing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersVendorSourcingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BiomarkersVendorSourcing200Response**](BiomarkersVendorSourcing200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BiomarkersWhereItAppears

> BiomarkersWhereItAppears200Response BiomarkersWhereItAppears(ctx, id).Execute()

Where it appears

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
	resp, r, err := apiClient.BiomarkersAPI.BiomarkersWhereItAppears(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiomarkersAPI.BiomarkersWhereItAppears``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BiomarkersWhereItAppears`: BiomarkersWhereItAppears200Response
	fmt.Fprintf(os.Stdout, "Response from `BiomarkersAPI.BiomarkersWhereItAppears`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBiomarkersWhereItAppearsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BiomarkersWhereItAppears200Response**](BiomarkersWhereItAppears200Response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

