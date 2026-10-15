# \ApiV1OwnershipApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_v1_investor_profile_get**](ApiV1OwnershipApi.md#api_v1_investor_profile_get) | **GET** /api/v1/investor_profile | investor_profile
[**api_v1_org_financing_and_ownership_get**](ApiV1OwnershipApi.md#api_v1_org_financing_and_ownership_get) | **GET** /api/v1/org_financing_and_ownership | org_financing_and_ownership



## api_v1_investor_profile_get

> models::ApiV1InvestorProfileGet200Response api_v1_investor_profile_get(investor_id, investor_name)
investor_profile

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**investor_id** | Option<**String**> |  |  |
**investor_name** | Option<**String**> |  |  |

### Return type

[**models::ApiV1InvestorProfileGet200Response**](_api_v1_investor_profile_get_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## api_v1_org_financing_and_ownership_get

> models::ApiV1OrgFinancingAndOwnershipGet200Response api_v1_org_financing_and_ownership_get(edgar_cik, org_id, org_name, ror_id)
org_financing_and_ownership

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**edgar_cik** | Option<**i32**> |  |  |
**org_id** | Option<**String**> |  |  |
**org_name** | Option<**String**> |  |  |
**ror_id** | Option<**String**> |  |  |

### Return type

[**models::ApiV1OrgFinancingAndOwnershipGet200Response**](_api_v1_org_financing_and_ownership_get_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

