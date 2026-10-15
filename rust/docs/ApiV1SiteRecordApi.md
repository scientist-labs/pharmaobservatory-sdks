# \ApiV1SiteRecordApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_v1_site_record_get**](ApiV1SiteRecordApi.md#api_v1_site_record_get) | **GET** /api/v1/site_record | Look up a clinical trial site's full record, including hosted trials and named investigators



## api_v1_site_record_get

> models::ApiV1SiteRecordGet200Response api_v1_site_record_get(investigator_id, limit, org_id, site_id)
Look up a clinical trial site's full record, including hosted trials and named investigators

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**investigator_id** | Option<**String**> |  |  |
**limit** | Option<**i32**> |  |  |
**org_id** | Option<**String**> |  |  |
**site_id** | Option<**String**> |  |  |

### Return type

[**models::ApiV1SiteRecordGet200Response**](_api_v1_site_record_get_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

