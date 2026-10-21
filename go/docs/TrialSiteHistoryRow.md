# TrialSiteHistoryRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialId** | Pointer to **NullableString** | The trial&#39;s ClinicalTrials.gov registry id (NCT, National Clinical Trial, number). | [optional] 
**AsOf** | Pointer to **NullableString** | The registry snapshot date this row describes. | [optional] 
**PrevAsOf** | Pointer to **NullableString** | The trial&#39;s previous snapshot date (null for its first snapshot). | [optional] 
**FacilitiesCaptured** | Pointer to **NullableBool** | True when the snapshot carried the site list; false means the site counts are unknown, not zero. | [optional] 
**StatusCaptured** | Pointer to **NullableBool** | True when the snapshot carried a recruitment status on the site rows. | [optional] 
**Comparable** | Pointer to **NullableBool** | True when the previous snapshot also carried the site list, so the add and drop counts are defined. | [optional] 
**SitesListed** | Pointer to **NullableInt32** | Distinct sites the trial listed at the snapshot (null where facilities_captured is false). | [optional] 
**SitesAdded** | Pointer to **NullableInt32** | Sites present at this snapshot and absent at the previous one (null where comparable is false). | [optional] 
**SitesDropped** | Pointer to **NullableInt32** | Sites present at the previous snapshot and absent at this one (null where comparable is false). | [optional] 
**SitesStatusPopulated** | Pointer to **NullableInt32** | Sites whose row carried a recruitment status at the snapshot. | [optional] 
**SitesRecruiting** | Pointer to **NullableInt32** | Sites with status RECRUITING at the snapshot. | [optional] 
**SitesNotYetRecruiting** | Pointer to **NullableInt32** | Sites with status NOT_YET_RECRUITING at the snapshot. | [optional] 
**SitesActiveNotRecruiting** | Pointer to **NullableInt32** | Sites with status ACTIVE_NOT_RECRUITING at the snapshot. | [optional] 
**SitesCompleted** | Pointer to **NullableInt32** | Sites with status COMPLETED at the snapshot. | [optional] 
**SitesOtherStatus** | Pointer to **NullableInt32** | Sites with a status other than the four counted above. | [optional] 
**CountriesListed** | Pointer to **NullableInt32** | Countries the listed sites span at the snapshot. | [optional] 
**CapturedColumns** | Pointer to **[]string** | The registry study fields the snapshot&#39;s archive carried (null when no capture record exists for the snapshot; an empty list means it carried none of them). | [optional] 

## Methods

### NewTrialSiteHistoryRow

`func NewTrialSiteHistoryRow() *TrialSiteHistoryRow`

NewTrialSiteHistoryRow instantiates a new TrialSiteHistoryRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialSiteHistoryRowWithDefaults

`func NewTrialSiteHistoryRowWithDefaults() *TrialSiteHistoryRow`

NewTrialSiteHistoryRowWithDefaults instantiates a new TrialSiteHistoryRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialId

`func (o *TrialSiteHistoryRow) GetTrialId() string`

GetTrialId returns the TrialId field if non-nil, zero value otherwise.

### GetTrialIdOk

`func (o *TrialSiteHistoryRow) GetTrialIdOk() (*string, bool)`

GetTrialIdOk returns a tuple with the TrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialId

`func (o *TrialSiteHistoryRow) SetTrialId(v string)`

SetTrialId sets TrialId field to given value.

### HasTrialId

`func (o *TrialSiteHistoryRow) HasTrialId() bool`

HasTrialId returns a boolean if a field has been set.

### SetTrialIdNil

`func (o *TrialSiteHistoryRow) SetTrialIdNil(b bool)`

 SetTrialIdNil sets the value for TrialId to be an explicit nil

### UnsetTrialId
`func (o *TrialSiteHistoryRow) UnsetTrialId()`

UnsetTrialId ensures that no value is present for TrialId, not even an explicit nil
### GetAsOf

`func (o *TrialSiteHistoryRow) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *TrialSiteHistoryRow) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *TrialSiteHistoryRow) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *TrialSiteHistoryRow) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### SetAsOfNil

`func (o *TrialSiteHistoryRow) SetAsOfNil(b bool)`

 SetAsOfNil sets the value for AsOf to be an explicit nil

### UnsetAsOf
`func (o *TrialSiteHistoryRow) UnsetAsOf()`

UnsetAsOf ensures that no value is present for AsOf, not even an explicit nil
### GetPrevAsOf

`func (o *TrialSiteHistoryRow) GetPrevAsOf() string`

GetPrevAsOf returns the PrevAsOf field if non-nil, zero value otherwise.

### GetPrevAsOfOk

`func (o *TrialSiteHistoryRow) GetPrevAsOfOk() (*string, bool)`

GetPrevAsOfOk returns a tuple with the PrevAsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevAsOf

`func (o *TrialSiteHistoryRow) SetPrevAsOf(v string)`

SetPrevAsOf sets PrevAsOf field to given value.

### HasPrevAsOf

`func (o *TrialSiteHistoryRow) HasPrevAsOf() bool`

HasPrevAsOf returns a boolean if a field has been set.

### SetPrevAsOfNil

`func (o *TrialSiteHistoryRow) SetPrevAsOfNil(b bool)`

 SetPrevAsOfNil sets the value for PrevAsOf to be an explicit nil

### UnsetPrevAsOf
`func (o *TrialSiteHistoryRow) UnsetPrevAsOf()`

UnsetPrevAsOf ensures that no value is present for PrevAsOf, not even an explicit nil
### GetFacilitiesCaptured

`func (o *TrialSiteHistoryRow) GetFacilitiesCaptured() bool`

GetFacilitiesCaptured returns the FacilitiesCaptured field if non-nil, zero value otherwise.

### GetFacilitiesCapturedOk

`func (o *TrialSiteHistoryRow) GetFacilitiesCapturedOk() (*bool, bool)`

GetFacilitiesCapturedOk returns a tuple with the FacilitiesCaptured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilitiesCaptured

`func (o *TrialSiteHistoryRow) SetFacilitiesCaptured(v bool)`

SetFacilitiesCaptured sets FacilitiesCaptured field to given value.

### HasFacilitiesCaptured

`func (o *TrialSiteHistoryRow) HasFacilitiesCaptured() bool`

HasFacilitiesCaptured returns a boolean if a field has been set.

### SetFacilitiesCapturedNil

`func (o *TrialSiteHistoryRow) SetFacilitiesCapturedNil(b bool)`

 SetFacilitiesCapturedNil sets the value for FacilitiesCaptured to be an explicit nil

### UnsetFacilitiesCaptured
`func (o *TrialSiteHistoryRow) UnsetFacilitiesCaptured()`

UnsetFacilitiesCaptured ensures that no value is present for FacilitiesCaptured, not even an explicit nil
### GetStatusCaptured

`func (o *TrialSiteHistoryRow) GetStatusCaptured() bool`

GetStatusCaptured returns the StatusCaptured field if non-nil, zero value otherwise.

### GetStatusCapturedOk

`func (o *TrialSiteHistoryRow) GetStatusCapturedOk() (*bool, bool)`

GetStatusCapturedOk returns a tuple with the StatusCaptured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCaptured

`func (o *TrialSiteHistoryRow) SetStatusCaptured(v bool)`

SetStatusCaptured sets StatusCaptured field to given value.

### HasStatusCaptured

`func (o *TrialSiteHistoryRow) HasStatusCaptured() bool`

HasStatusCaptured returns a boolean if a field has been set.

### SetStatusCapturedNil

`func (o *TrialSiteHistoryRow) SetStatusCapturedNil(b bool)`

 SetStatusCapturedNil sets the value for StatusCaptured to be an explicit nil

### UnsetStatusCaptured
`func (o *TrialSiteHistoryRow) UnsetStatusCaptured()`

UnsetStatusCaptured ensures that no value is present for StatusCaptured, not even an explicit nil
### GetComparable

`func (o *TrialSiteHistoryRow) GetComparable() bool`

GetComparable returns the Comparable field if non-nil, zero value otherwise.

### GetComparableOk

`func (o *TrialSiteHistoryRow) GetComparableOk() (*bool, bool)`

GetComparableOk returns a tuple with the Comparable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparable

`func (o *TrialSiteHistoryRow) SetComparable(v bool)`

SetComparable sets Comparable field to given value.

### HasComparable

`func (o *TrialSiteHistoryRow) HasComparable() bool`

HasComparable returns a boolean if a field has been set.

### SetComparableNil

`func (o *TrialSiteHistoryRow) SetComparableNil(b bool)`

 SetComparableNil sets the value for Comparable to be an explicit nil

### UnsetComparable
`func (o *TrialSiteHistoryRow) UnsetComparable()`

UnsetComparable ensures that no value is present for Comparable, not even an explicit nil
### GetSitesListed

`func (o *TrialSiteHistoryRow) GetSitesListed() int32`

GetSitesListed returns the SitesListed field if non-nil, zero value otherwise.

### GetSitesListedOk

`func (o *TrialSiteHistoryRow) GetSitesListedOk() (*int32, bool)`

GetSitesListedOk returns a tuple with the SitesListed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitesListed

`func (o *TrialSiteHistoryRow) SetSitesListed(v int32)`

SetSitesListed sets SitesListed field to given value.

### HasSitesListed

`func (o *TrialSiteHistoryRow) HasSitesListed() bool`

HasSitesListed returns a boolean if a field has been set.

### SetSitesListedNil

`func (o *TrialSiteHistoryRow) SetSitesListedNil(b bool)`

 SetSitesListedNil sets the value for SitesListed to be an explicit nil

### UnsetSitesListed
`func (o *TrialSiteHistoryRow) UnsetSitesListed()`

UnsetSitesListed ensures that no value is present for SitesListed, not even an explicit nil
### GetSitesAdded

`func (o *TrialSiteHistoryRow) GetSitesAdded() int32`

GetSitesAdded returns the SitesAdded field if non-nil, zero value otherwise.

### GetSitesAddedOk

`func (o *TrialSiteHistoryRow) GetSitesAddedOk() (*int32, bool)`

GetSitesAddedOk returns a tuple with the SitesAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitesAdded

`func (o *TrialSiteHistoryRow) SetSitesAdded(v int32)`

SetSitesAdded sets SitesAdded field to given value.

### HasSitesAdded

`func (o *TrialSiteHistoryRow) HasSitesAdded() bool`

HasSitesAdded returns a boolean if a field has been set.

### SetSitesAddedNil

`func (o *TrialSiteHistoryRow) SetSitesAddedNil(b bool)`

 SetSitesAddedNil sets the value for SitesAdded to be an explicit nil

### UnsetSitesAdded
`func (o *TrialSiteHistoryRow) UnsetSitesAdded()`

UnsetSitesAdded ensures that no value is present for SitesAdded, not even an explicit nil
### GetSitesDropped

`func (o *TrialSiteHistoryRow) GetSitesDropped() int32`

GetSitesDropped returns the SitesDropped field if non-nil, zero value otherwise.

### GetSitesDroppedOk

`func (o *TrialSiteHistoryRow) GetSitesDroppedOk() (*int32, bool)`

GetSitesDroppedOk returns a tuple with the SitesDropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitesDropped

`func (o *TrialSiteHistoryRow) SetSitesDropped(v int32)`

SetSitesDropped sets SitesDropped field to given value.

### HasSitesDropped

`func (o *TrialSiteHistoryRow) HasSitesDropped() bool`

HasSitesDropped returns a boolean if a field has been set.

### SetSitesDroppedNil

`func (o *TrialSiteHistoryRow) SetSitesDroppedNil(b bool)`

 SetSitesDroppedNil sets the value for SitesDropped to be an explicit nil

### UnsetSitesDropped
`func (o *TrialSiteHistoryRow) UnsetSitesDropped()`

UnsetSitesDropped ensures that no value is present for SitesDropped, not even an explicit nil
### GetSitesStatusPopulated

`func (o *TrialSiteHistoryRow) GetSitesStatusPopulated() int32`

GetSitesStatusPopulated returns the SitesStatusPopulated field if non-nil, zero value otherwise.

### GetSitesStatusPopulatedOk

`func (o *TrialSiteHistoryRow) GetSitesStatusPopulatedOk() (*int32, bool)`

GetSitesStatusPopulatedOk returns a tuple with the SitesStatusPopulated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitesStatusPopulated

`func (o *TrialSiteHistoryRow) SetSitesStatusPopulated(v int32)`

SetSitesStatusPopulated sets SitesStatusPopulated field to given value.

### HasSitesStatusPopulated

`func (o *TrialSiteHistoryRow) HasSitesStatusPopulated() bool`

HasSitesStatusPopulated returns a boolean if a field has been set.

### SetSitesStatusPopulatedNil

`func (o *TrialSiteHistoryRow) SetSitesStatusPopulatedNil(b bool)`

 SetSitesStatusPopulatedNil sets the value for SitesStatusPopulated to be an explicit nil

### UnsetSitesStatusPopulated
`func (o *TrialSiteHistoryRow) UnsetSitesStatusPopulated()`

UnsetSitesStatusPopulated ensures that no value is present for SitesStatusPopulated, not even an explicit nil
### GetSitesRecruiting

`func (o *TrialSiteHistoryRow) GetSitesRecruiting() int32`

GetSitesRecruiting returns the SitesRecruiting field if non-nil, zero value otherwise.

### GetSitesRecruitingOk

`func (o *TrialSiteHistoryRow) GetSitesRecruitingOk() (*int32, bool)`

GetSitesRecruitingOk returns a tuple with the SitesRecruiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitesRecruiting

`func (o *TrialSiteHistoryRow) SetSitesRecruiting(v int32)`

SetSitesRecruiting sets SitesRecruiting field to given value.

### HasSitesRecruiting

`func (o *TrialSiteHistoryRow) HasSitesRecruiting() bool`

HasSitesRecruiting returns a boolean if a field has been set.

### SetSitesRecruitingNil

`func (o *TrialSiteHistoryRow) SetSitesRecruitingNil(b bool)`

 SetSitesRecruitingNil sets the value for SitesRecruiting to be an explicit nil

### UnsetSitesRecruiting
`func (o *TrialSiteHistoryRow) UnsetSitesRecruiting()`

UnsetSitesRecruiting ensures that no value is present for SitesRecruiting, not even an explicit nil
### GetSitesNotYetRecruiting

`func (o *TrialSiteHistoryRow) GetSitesNotYetRecruiting() int32`

GetSitesNotYetRecruiting returns the SitesNotYetRecruiting field if non-nil, zero value otherwise.

### GetSitesNotYetRecruitingOk

`func (o *TrialSiteHistoryRow) GetSitesNotYetRecruitingOk() (*int32, bool)`

GetSitesNotYetRecruitingOk returns a tuple with the SitesNotYetRecruiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitesNotYetRecruiting

`func (o *TrialSiteHistoryRow) SetSitesNotYetRecruiting(v int32)`

SetSitesNotYetRecruiting sets SitesNotYetRecruiting field to given value.

### HasSitesNotYetRecruiting

`func (o *TrialSiteHistoryRow) HasSitesNotYetRecruiting() bool`

HasSitesNotYetRecruiting returns a boolean if a field has been set.

### SetSitesNotYetRecruitingNil

`func (o *TrialSiteHistoryRow) SetSitesNotYetRecruitingNil(b bool)`

 SetSitesNotYetRecruitingNil sets the value for SitesNotYetRecruiting to be an explicit nil

### UnsetSitesNotYetRecruiting
`func (o *TrialSiteHistoryRow) UnsetSitesNotYetRecruiting()`

UnsetSitesNotYetRecruiting ensures that no value is present for SitesNotYetRecruiting, not even an explicit nil
### GetSitesActiveNotRecruiting

`func (o *TrialSiteHistoryRow) GetSitesActiveNotRecruiting() int32`

GetSitesActiveNotRecruiting returns the SitesActiveNotRecruiting field if non-nil, zero value otherwise.

### GetSitesActiveNotRecruitingOk

`func (o *TrialSiteHistoryRow) GetSitesActiveNotRecruitingOk() (*int32, bool)`

GetSitesActiveNotRecruitingOk returns a tuple with the SitesActiveNotRecruiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitesActiveNotRecruiting

`func (o *TrialSiteHistoryRow) SetSitesActiveNotRecruiting(v int32)`

SetSitesActiveNotRecruiting sets SitesActiveNotRecruiting field to given value.

### HasSitesActiveNotRecruiting

`func (o *TrialSiteHistoryRow) HasSitesActiveNotRecruiting() bool`

HasSitesActiveNotRecruiting returns a boolean if a field has been set.

### SetSitesActiveNotRecruitingNil

`func (o *TrialSiteHistoryRow) SetSitesActiveNotRecruitingNil(b bool)`

 SetSitesActiveNotRecruitingNil sets the value for SitesActiveNotRecruiting to be an explicit nil

### UnsetSitesActiveNotRecruiting
`func (o *TrialSiteHistoryRow) UnsetSitesActiveNotRecruiting()`

UnsetSitesActiveNotRecruiting ensures that no value is present for SitesActiveNotRecruiting, not even an explicit nil
### GetSitesCompleted

`func (o *TrialSiteHistoryRow) GetSitesCompleted() int32`

GetSitesCompleted returns the SitesCompleted field if non-nil, zero value otherwise.

### GetSitesCompletedOk

`func (o *TrialSiteHistoryRow) GetSitesCompletedOk() (*int32, bool)`

GetSitesCompletedOk returns a tuple with the SitesCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitesCompleted

`func (o *TrialSiteHistoryRow) SetSitesCompleted(v int32)`

SetSitesCompleted sets SitesCompleted field to given value.

### HasSitesCompleted

`func (o *TrialSiteHistoryRow) HasSitesCompleted() bool`

HasSitesCompleted returns a boolean if a field has been set.

### SetSitesCompletedNil

`func (o *TrialSiteHistoryRow) SetSitesCompletedNil(b bool)`

 SetSitesCompletedNil sets the value for SitesCompleted to be an explicit nil

### UnsetSitesCompleted
`func (o *TrialSiteHistoryRow) UnsetSitesCompleted()`

UnsetSitesCompleted ensures that no value is present for SitesCompleted, not even an explicit nil
### GetSitesOtherStatus

`func (o *TrialSiteHistoryRow) GetSitesOtherStatus() int32`

GetSitesOtherStatus returns the SitesOtherStatus field if non-nil, zero value otherwise.

### GetSitesOtherStatusOk

`func (o *TrialSiteHistoryRow) GetSitesOtherStatusOk() (*int32, bool)`

GetSitesOtherStatusOk returns a tuple with the SitesOtherStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitesOtherStatus

`func (o *TrialSiteHistoryRow) SetSitesOtherStatus(v int32)`

SetSitesOtherStatus sets SitesOtherStatus field to given value.

### HasSitesOtherStatus

`func (o *TrialSiteHistoryRow) HasSitesOtherStatus() bool`

HasSitesOtherStatus returns a boolean if a field has been set.

### SetSitesOtherStatusNil

`func (o *TrialSiteHistoryRow) SetSitesOtherStatusNil(b bool)`

 SetSitesOtherStatusNil sets the value for SitesOtherStatus to be an explicit nil

### UnsetSitesOtherStatus
`func (o *TrialSiteHistoryRow) UnsetSitesOtherStatus()`

UnsetSitesOtherStatus ensures that no value is present for SitesOtherStatus, not even an explicit nil
### GetCountriesListed

`func (o *TrialSiteHistoryRow) GetCountriesListed() int32`

GetCountriesListed returns the CountriesListed field if non-nil, zero value otherwise.

### GetCountriesListedOk

`func (o *TrialSiteHistoryRow) GetCountriesListedOk() (*int32, bool)`

GetCountriesListedOk returns a tuple with the CountriesListed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountriesListed

`func (o *TrialSiteHistoryRow) SetCountriesListed(v int32)`

SetCountriesListed sets CountriesListed field to given value.

### HasCountriesListed

`func (o *TrialSiteHistoryRow) HasCountriesListed() bool`

HasCountriesListed returns a boolean if a field has been set.

### SetCountriesListedNil

`func (o *TrialSiteHistoryRow) SetCountriesListedNil(b bool)`

 SetCountriesListedNil sets the value for CountriesListed to be an explicit nil

### UnsetCountriesListed
`func (o *TrialSiteHistoryRow) UnsetCountriesListed()`

UnsetCountriesListed ensures that no value is present for CountriesListed, not even an explicit nil
### GetCapturedColumns

`func (o *TrialSiteHistoryRow) GetCapturedColumns() []string`

GetCapturedColumns returns the CapturedColumns field if non-nil, zero value otherwise.

### GetCapturedColumnsOk

`func (o *TrialSiteHistoryRow) GetCapturedColumnsOk() (*[]string, bool)`

GetCapturedColumnsOk returns a tuple with the CapturedColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedColumns

`func (o *TrialSiteHistoryRow) SetCapturedColumns(v []string)`

SetCapturedColumns sets CapturedColumns field to given value.

### HasCapturedColumns

`func (o *TrialSiteHistoryRow) HasCapturedColumns() bool`

HasCapturedColumns returns a boolean if a field has been set.

### SetCapturedColumnsNil

`func (o *TrialSiteHistoryRow) SetCapturedColumnsNil(b bool)`

 SetCapturedColumnsNil sets the value for CapturedColumns to be an explicit nil

### UnsetCapturedColumns
`func (o *TrialSiteHistoryRow) UnsetCapturedColumns()`

UnsetCapturedColumns ensures that no value is present for CapturedColumns, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


