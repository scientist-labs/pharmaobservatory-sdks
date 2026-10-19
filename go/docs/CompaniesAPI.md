# \CompaniesAPI

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompaniesBiomarkerBqpAssertions**](CompaniesAPI.md#CompaniesBiomarkerBqpAssertions) | **Get** /api/v1/companies/{id}/biomarker_bqp_assertions | FDA biomarker qualification projects
[**CompaniesBiomarkerCdxAssertions**](CompaniesAPI.md#CompaniesBiomarkerCdxAssertions) | **Get** /api/v1/companies/{id}/biomarker_cdx_assertions | FDA companion diagnostic assertions
[**CompaniesBiomarkerPublishingCompanies**](CompaniesAPI.md#CompaniesBiomarkerPublishingCompanies) | **Get** /api/v1/companies/{id}/biomarker_publishing_companies | Companies publishing on a biomarker
[**CompaniesCompanyMembers**](CompaniesAPI.md#CompaniesCompanyMembers) | **Get** /api/v1/companies/{id}/company_members | Company members
[**CompaniesCompanyPrograms**](CompaniesAPI.md#CompaniesCompanyPrograms) | **Get** /api/v1/companies/{id}/company_programs | Company programs
[**CompaniesCompanyTargets**](CompaniesAPI.md#CompaniesCompanyTargets) | **Get** /api/v1/companies/{id}/company_targets | Targets a company pursues
[**CompaniesDrugPatentCliff**](CompaniesAPI.md#CompaniesDrugPatentCliff) | **Get** /api/v1/companies/{id}/drug_patent_cliff | Patent-cliff exposure
[**CompaniesGet**](CompaniesAPI.md#CompaniesGet) | **Get** /api/v1/companies/{id} | One company
[**CompaniesInstitutionalHoldings**](CompaniesAPI.md#CompaniesInstitutionalHoldings) | **Get** /api/v1/companies/{id}/institutional_holdings | Institutional holdings (SEC Form 13F)
[**CompaniesLegalEntities**](CompaniesAPI.md#CompaniesLegalEntities) | **Get** /api/v1/companies/{id}/legal_entities | Legal entities
[**CompaniesList**](CompaniesAPI.md#CompaniesList) | **Get** /api/v1/companies | List companies
[**CompaniesPdufaDates**](CompaniesAPI.md#CompaniesPdufaDates) | **Get** /api/v1/companies/{id}/pdufa_dates | FDA decision dates (PDUFA) and catalysts
[**CompaniesSponsorRecruitmentByPhase**](CompaniesAPI.md#CompaniesSponsorRecruitmentByPhase) | **Get** /api/v1/companies/{id}/sponsor_recruitment_by_phase | Trial recruitment duration by sponsor and phase
[**CompaniesSponsorRecruitmentDuration**](CompaniesAPI.md#CompaniesSponsorRecruitmentDuration) | **Get** /api/v1/companies/{id}/sponsor_recruitment_duration | Trial recruitment duration by sponsor
[**CompaniesSponsorTrialStartForecast**](CompaniesAPI.md#CompaniesSponsorTrialStartForecast) | **Get** /api/v1/companies/{id}/sponsor_trial_start_forecast | Trial-start forecast by sponsor
[**CompaniesSponsorTrials**](CompaniesAPI.md#CompaniesSponsorTrials) | **Get** /api/v1/companies/{id}/sponsor_trials | Sponsor trials



## CompaniesBiomarkerBqpAssertions

> BiomarkerBqpAssertionsResponse CompaniesBiomarkerBqpAssertions(ctx, id).Limit(limit).Scope(scope).Sort(sort).Execute()

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
	id := "company:c0185cf6c39a2536ad436cf6255494c0" // string | 
	limit := int32(2) // int32 |  (optional)
	scope := "company" // string |  (optional)
	sort := "route_key" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompaniesAPI.CompaniesBiomarkerBqpAssertions(context.Background(), id).Limit(limit).Scope(scope).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CompaniesBiomarkerBqpAssertions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesBiomarkerBqpAssertions`: BiomarkerBqpAssertionsResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CompaniesBiomarkerBqpAssertions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesBiomarkerBqpAssertionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **scope** | **string** |  | 
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


## CompaniesBiomarkerCdxAssertions

> BiomarkerCdxAssertionsResponse CompaniesBiomarkerCdxAssertions(ctx, id).Limit(limit).Scope(scope).Sort(sort).Execute()

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
	id := "company:c0185cf6c39a2536ad436cf6255494c0" // string | 
	limit := int32(2) // int32 |  (optional)
	scope := "company" // string |  (optional)
	sort := "biomarker_id" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompaniesAPI.CompaniesBiomarkerCdxAssertions(context.Background(), id).Limit(limit).Scope(scope).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CompaniesBiomarkerCdxAssertions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesBiomarkerCdxAssertions`: BiomarkerCdxAssertionsResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CompaniesBiomarkerCdxAssertions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesBiomarkerCdxAssertionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **scope** | **string** |  | 
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


## CompaniesBiomarkerPublishingCompanies

> BiomarkerPublishingCompaniesResponse CompaniesBiomarkerPublishingCompanies(ctx, id).Limit(limit).Scope(scope).Sort(sort).Execute()

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
	id := "company:c0185cf6c39a2536ad436cf6255494c0" // string | 
	limit := int32(2) // int32 |  (optional)
	scope := "company" // string |  (optional)
	sort := "n_publications" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompaniesAPI.CompaniesBiomarkerPublishingCompanies(context.Background(), id).Limit(limit).Scope(scope).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CompaniesBiomarkerPublishingCompanies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesBiomarkerPublishingCompanies`: BiomarkerPublishingCompaniesResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CompaniesBiomarkerPublishingCompanies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesBiomarkerPublishingCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **scope** | **string** |  | 
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


## CompaniesCompanyMembers

> CompanyMembersResponse CompaniesCompanyMembers(ctx, id).Limit(limit).Scope(scope).Sort(sort).Execute()

Company members

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
	id := "company:c0185cf6c39a2536ad436cf6255494c0" // string | 
	limit := int32(2) // int32 |  (optional)
	scope := "company" // string |  (optional)
	sort := "company_id" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompaniesAPI.CompaniesCompanyMembers(context.Background(), id).Limit(limit).Scope(scope).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CompaniesCompanyMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyMembers`: CompanyMembersResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CompaniesCompanyMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **scope** | **string** |  | 
 **sort** | **string** |  | 

### Return type

[**CompanyMembersResponse**](CompanyMembersResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyPrograms

> CompanyProgramsResponse CompaniesCompanyPrograms(ctx, id).Limit(limit).Scope(scope).Sort(sort).Execute()

Company programs

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
	id := "company:c0185cf6c39a2536ad436cf6255494c0" // string | 
	limit := int32(2) // int32 |  (optional)
	scope := "company" // string |  (optional)
	sort := "max_phase" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompaniesAPI.CompaniesCompanyPrograms(context.Background(), id).Limit(limit).Scope(scope).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CompaniesCompanyPrograms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyPrograms`: CompanyProgramsResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CompaniesCompanyPrograms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyProgramsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **scope** | **string** |  | 
 **sort** | **string** |  | 

### Return type

[**CompanyProgramsResponse**](CompanyProgramsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyTargets

> CompanyTargetsResponse CompaniesCompanyTargets(ctx, id).Limit(limit).Scope(scope).Sort(sort).Execute()

Targets a company pursues

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
	id := "company:c0185cf6c39a2536ad436cf6255494c0" // string | 
	limit := int32(2) // int32 |  (optional)
	scope := "company" // string |  (optional)
	sort := "max_phase" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompaniesAPI.CompaniesCompanyTargets(context.Background(), id).Limit(limit).Scope(scope).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CompaniesCompanyTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyTargets`: CompanyTargetsResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CompaniesCompanyTargets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **scope** | **string** |  | 
 **sort** | **string** |  | 

### Return type

[**CompanyTargetsResponse**](CompanyTargetsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesDrugPatentCliff

> DrugPatentCliffResponse CompaniesDrugPatentCliff(ctx, id).Limit(limit).Scope(scope).Sort(sort).Execute()

Patent-cliff exposure

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
	id := "company:c0185cf6c39a2536ad436cf6255494c0" // string | 
	limit := int32(2) // int32 |  (optional)
	scope := "company" // string |  (optional)
	sort := "cliff_date" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompaniesAPI.CompaniesDrugPatentCliff(context.Background(), id).Limit(limit).Scope(scope).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CompaniesDrugPatentCliff``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesDrugPatentCliff`: DrugPatentCliffResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CompaniesDrugPatentCliff`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesDrugPatentCliffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **scope** | **string** |  | 
 **sort** | **string** |  | 

### Return type

[**DrugPatentCliffResponse**](DrugPatentCliffResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesGet

> CompaniesRecord CompaniesGet(ctx, id).Execute()

One company

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
	id := "company:c0185cf6c39a2536ad436cf6255494c0" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompaniesAPI.CompaniesGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CompaniesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesGet`: CompaniesRecord
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CompaniesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CompaniesRecord**](CompaniesRecord.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesInstitutionalHoldings

> InstitutionalHoldingsResponse CompaniesInstitutionalHoldings(ctx, id).Limit(limit).Scope(scope).Sort(sort).Execute()

Institutional holdings (SEC Form 13F)

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
	id := "company:c0185cf6c39a2536ad436cf6255494c0" // string | 
	limit := int32(2) // int32 |  (optional)
	scope := "company" // string |  (optional)
	sort := "reported_value_amount" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompaniesAPI.CompaniesInstitutionalHoldings(context.Background(), id).Limit(limit).Scope(scope).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CompaniesInstitutionalHoldings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesInstitutionalHoldings`: InstitutionalHoldingsResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CompaniesInstitutionalHoldings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesInstitutionalHoldingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **scope** | **string** |  | 
 **sort** | **string** |  | 

### Return type

[**InstitutionalHoldingsResponse**](InstitutionalHoldingsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesLegalEntities

> LegalEntitiesResponse CompaniesLegalEntities(ctx, id).Limit(limit).Scope(scope).Sort(sort).Execute()

Legal entities

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
	id := "company:c0185cf6c39a2536ad436cf6255494c0" // string | 
	limit := int32(2) // int32 |  (optional)
	scope := "company" // string |  (optional)
	sort := "activity_score" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompaniesAPI.CompaniesLegalEntities(context.Background(), id).Limit(limit).Scope(scope).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CompaniesLegalEntities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesLegalEntities`: LegalEntitiesResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CompaniesLegalEntities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesLegalEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **scope** | **string** |  | 
 **sort** | **string** |  | 

### Return type

[**LegalEntitiesResponse**](LegalEntitiesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesList

> CompaniesResponse CompaniesList(ctx).Limit(limit).Sort(sort).Execute()

List companies

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
	sort := "activity_score" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompaniesAPI.CompaniesList(context.Background()).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CompaniesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesList`: CompaniesResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CompaniesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **sort** | **string** |  | 

### Return type

[**CompaniesResponse**](CompaniesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesPdufaDates

> PdufaDatesResponse CompaniesPdufaDates(ctx, id).Limit(limit).Scope(scope).Sort(sort).Execute()

FDA decision dates (PDUFA) and catalysts

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
	id := "company:c0185cf6c39a2536ad436cf6255494c0" // string | 
	limit := int32(2) // int32 |  (optional)
	scope := "company" // string |  (optional)
	sort := "decision_date" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompaniesAPI.CompaniesPdufaDates(context.Background(), id).Limit(limit).Scope(scope).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CompaniesPdufaDates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesPdufaDates`: PdufaDatesResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CompaniesPdufaDates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesPdufaDatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **scope** | **string** |  | 
 **sort** | **string** |  | 

### Return type

[**PdufaDatesResponse**](PdufaDatesResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesSponsorRecruitmentByPhase

> SponsorRecruitmentByPhaseResponse CompaniesSponsorRecruitmentByPhase(ctx, id).Limit(limit).Scope(scope).Sort(sort).Execute()

Trial recruitment duration by sponsor and phase

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
	id := "company:c0185cf6c39a2536ad436cf6255494c0" // string | 
	limit := int32(2) // int32 |  (optional)
	scope := "company" // string |  (optional)
	sort := "sponsor_org_id" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompaniesAPI.CompaniesSponsorRecruitmentByPhase(context.Background(), id).Limit(limit).Scope(scope).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CompaniesSponsorRecruitmentByPhase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesSponsorRecruitmentByPhase`: SponsorRecruitmentByPhaseResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CompaniesSponsorRecruitmentByPhase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesSponsorRecruitmentByPhaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **scope** | **string** |  | 
 **sort** | **string** |  | 

### Return type

[**SponsorRecruitmentByPhaseResponse**](SponsorRecruitmentByPhaseResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesSponsorRecruitmentDuration

> SponsorRecruitmentDurationResponse CompaniesSponsorRecruitmentDuration(ctx, id).Limit(limit).Scope(scope).Sort(sort).Execute()

Trial recruitment duration by sponsor

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
	id := "company:c0185cf6c39a2536ad436cf6255494c0" // string | 
	limit := int32(2) // int32 |  (optional)
	scope := "company" // string |  (optional)
	sort := "sponsor_org_id" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompaniesAPI.CompaniesSponsorRecruitmentDuration(context.Background(), id).Limit(limit).Scope(scope).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CompaniesSponsorRecruitmentDuration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesSponsorRecruitmentDuration`: SponsorRecruitmentDurationResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CompaniesSponsorRecruitmentDuration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesSponsorRecruitmentDurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **scope** | **string** |  | 
 **sort** | **string** |  | 

### Return type

[**SponsorRecruitmentDurationResponse**](SponsorRecruitmentDurationResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesSponsorTrialStartForecast

> SponsorTrialStartForecastResponse CompaniesSponsorTrialStartForecast(ctx, id).Limit(limit).Scope(scope).Sort(sort).Execute()

Trial-start forecast by sponsor

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
	id := "company:c0185cf6c39a2536ad436cf6255494c0" // string | 
	limit := int32(2) // int32 |  (optional)
	scope := "company" // string |  (optional)
	sort := "sponsor_org_id" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompaniesAPI.CompaniesSponsorTrialStartForecast(context.Background(), id).Limit(limit).Scope(scope).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CompaniesSponsorTrialStartForecast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesSponsorTrialStartForecast`: SponsorTrialStartForecastResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CompaniesSponsorTrialStartForecast`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesSponsorTrialStartForecastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **scope** | **string** |  | 
 **sort** | **string** |  | 

### Return type

[**SponsorTrialStartForecastResponse**](SponsorTrialStartForecastResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesSponsorTrials

> SponsorTrialsResponse CompaniesSponsorTrials(ctx, id).Limit(limit).Scope(scope).Sort(sort).Execute()

Sponsor trials

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
	id := "company:c0185cf6c39a2536ad436cf6255494c0" // string | 
	limit := int32(2) // int32 |  (optional)
	scope := "company" // string |  (optional)
	sort := "start_date" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompaniesAPI.CompaniesSponsorTrials(context.Background(), id).Limit(limit).Scope(scope).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CompaniesSponsorTrials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesSponsorTrials`: SponsorTrialsResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CompaniesSponsorTrials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesSponsorTrialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **scope** | **string** |  | 
 **sort** | **string** |  | 

### Return type

[**SponsorTrialsResponse**](SponsorTrialsResponse.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

