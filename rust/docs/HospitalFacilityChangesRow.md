# HospitalFacilityChangesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ccn** | Option<**String**> | CMS Certification Number of the hospital the event is about. | [optional]
**as_of** | Option<[**String**](string.md)> | The quarter-end date the change was detected at. | [optional]
**change_type** | Option<**String**> | rename (same CCN, new provider name), appeared (CCN enters the series after its start), disappeared (CCN leaves it before its end), merge_candidate (a retired CCN with one successor) or split_candidate (a retired CCN with several). | [optional]
**from_value** | Option<**String**> | The provider name before the change (rename), or the last name of the retired CCN (disappeared, merge and split candidates). Null on appeared. | [optional]
**to_value** | Option<**String**> | The provider name after the change (rename), the first name of the new CCN (appeared, merge and split candidates). Null on disappeared. | [optional]
**related_ccn** | Option<**String**> | The nearest later CCN that shares the retired CCN's normalised name, city and state. Set on merge_candidate and split_candidate rows. | [optional]
**successor_ccn** | Option<**String**> | The CCN that absorbed the retired one. Set on merge_candidate rows only; null on split_candidate, where more than one CCN matches. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


