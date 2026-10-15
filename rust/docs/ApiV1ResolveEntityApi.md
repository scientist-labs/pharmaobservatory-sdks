# \ApiV1ResolveEntityApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_v1_resolve_entity_get**](ApiV1ResolveEntityApi.md#api_v1_resolve_entity_get) | **GET** /api/v1/resolve_entity | Resolve a free-text or coded entity mention to a canonical Pharma Observatory ID



## api_v1_resolve_entity_get

> models::ApiV1ResolveEntityGet200Response api_v1_resolve_entity_get(query, raw_key, source_type, spine)
Resolve a free-text or coded entity mention to a canonical Pharma Observatory ID

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**query** | Option<**String**> |  |  |
**raw_key** | Option<**String**> |  |  |
**source_type** | Option<**String**> |  |  |
**spine** | Option<**String**> |  |  |

### Return type

[**models::ApiV1ResolveEntityGet200Response**](_api_v1_resolve_entity_get_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

