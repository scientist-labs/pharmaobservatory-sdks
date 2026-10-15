# \ApiV1CandidateSiteApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_v1_candidate_sites_get**](ApiV1CandidateSiteApi.md#api_v1_candidate_sites_get) | **GET** /api/v1/candidate_sites | Rank candidate clinical trial sites for a condition by composite capacity, activity, and competition scores



## api_v1_candidate_sites_get

> models::ApiV1CandidateSitesGet200Response api_v1_candidate_sites_get(condition_id, country, exclude_saturated, limit)
Rank candidate clinical trial sites for a condition by composite capacity, activity, and competition scores

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**condition_id** | Option<**String**> |  |  |
**country** | Option<**String**> |  |  |
**exclude_saturated** | Option<**bool**> |  |  |
**limit** | Option<**i32**> |  |  |

### Return type

[**models::ApiV1CandidateSitesGet200Response**](_api_v1_candidate_sites_get_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

