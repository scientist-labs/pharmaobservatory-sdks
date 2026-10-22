# DrugSitesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**molecule_id** | Option<**String**> | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional]
**site_id** | Option<**String**> | The site key: one facility spelling from the registry. | [optional]
**site_name** | Option<**String**> | The facility name as registered. | [optional]
**city** | Option<**String**> | The facility's city. | [optional]
**country** | Option<**String**> | The facility's country. | [optional]
**n_trials** | Option<**i32**> | The drug's trials registered at the site. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


