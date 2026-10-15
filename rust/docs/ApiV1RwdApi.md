# \ApiV1RwdApi

All URIs are relative to *https://pharmaobservatory.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**api_v1_rwd_conductor_track_record_get**](ApiV1RwdApi.md#api_v1_rwd_conductor_track_record_get) | **GET** /api/v1/rwd_conductor_track_record | conductor_track_record
[**api_v1_rwd_study_screen_get**](ApiV1RwdApi.md#api_v1_rwd_study_screen_get) | **GET** /api/v1/rwd_study_screen | study_screen



## api_v1_rwd_conductor_track_record_get

> models::ApiV1RwdConductorTrackRecordGet200Response api_v1_rwd_conductor_track_record_get(canonical_org_id, hma_ema_id, institution_name)
conductor_track_record

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**canonical_org_id** | Option<**String**> |  |  |
**hma_ema_id** | Option<**i32**> |  |  |
**institution_name** | Option<**String**> |  |  |

### Return type

[**models::ApiV1RwdConductorTrackRecordGet200Response**](_api_v1_rwd_conductor_track_record_get_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## api_v1_rwd_study_screen_get

> models::ApiV1RwdStudyScreenGet200Response api_v1_rwd_study_screen_get(condition, mandate)
study_screen

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**condition** | Option<**String**> |  |  |
**mandate** | Option<**String**> |  |  |

### Return type

[**models::ApiV1RwdStudyScreenGet200Response**](_api_v1_rwd_study_screen_get_200_response.md)

### Authorization

[ApiBearerJwt](../README.md#ApiBearerJwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

