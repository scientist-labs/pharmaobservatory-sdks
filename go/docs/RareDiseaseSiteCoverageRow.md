# RareDiseaseSiteCoverageRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrphaCode** | Pointer to **NullableString** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] 
**ConceptType** | Pointer to **NullableString** | Whether the concept is a group of disorders, a disorder or a subtype. | [optional] 
**SnapshotDate** | Pointer to **NullableString** | The date of the registry snapshot. | [optional] 
**ExactState** | Pointer to **NullableString** | How far the exact arm reaches: no_mapping, no_trials, no_site_data or computed. | [optional] 
**ExactNConditions** | Pointer to **NullableInt32** | MeSH conditions the disease maps to under the exact arm. | [optional] 
**ExactNTrials** | Pointer to **NullableInt32** | Trials linked under the exact arm. | [optional] 
**ExactNTrialsWithSite** | Pointer to **NullableInt32** | Trials of the exact arm with at least one geocoded site. | [optional] 
**ExactNSitesTotal** | Pointer to **NullableInt32** | Distinct sites of the exact arm. | [optional] 
**ExactNSitesEmitted** | Pointer to **NullableInt32** | Sites of the exact arm published in rare_disease_trial_sites (capped at 500 per disease and arm). | [optional] 
**ExactNSitesEmittedGeocoded** | Pointer to **NullableInt32** | Published sites of the exact arm that carry a coordinate. | [optional] 
**ExactNFacilitiesUngeocoded** | Pointer to **NullableInt32** | Facility listings of the exact arm with no coordinate. | [optional] 
**UnspecifiedState** | Pointer to **NullableString** | How far the unspecified arm reaches: no_mapping, no_trials, no_site_data or computed. | [optional] 
**UnspecifiedNConditions** | Pointer to **NullableInt32** | MeSH conditions the disease maps to under the unspecified arm. | [optional] 
**UnspecifiedNTrials** | Pointer to **NullableInt32** | Trials linked under the unspecified arm. | [optional] 
**UnspecifiedNTrialsWithSite** | Pointer to **NullableInt32** | Trials of the unspecified arm with at least one geocoded site. | [optional] 
**UnspecifiedNSitesTotal** | Pointer to **NullableInt32** | Distinct sites of the unspecified arm. | [optional] 
**UnspecifiedNSitesEmitted** | Pointer to **NullableInt32** | Sites of the unspecified arm published in rare_disease_trial_sites (capped at 500 per disease and arm). | [optional] 
**UnspecifiedNSitesEmittedGeocoded** | Pointer to **NullableInt32** | Published sites of the unspecified arm that carry a coordinate. | [optional] 
**UnspecifiedNFacilitiesUngeocoded** | Pointer to **NullableInt32** | Facility listings of the unspecified arm with no coordinate. | [optional] 

## Methods

### NewRareDiseaseSiteCoverageRow

`func NewRareDiseaseSiteCoverageRow() *RareDiseaseSiteCoverageRow`

NewRareDiseaseSiteCoverageRow instantiates a new RareDiseaseSiteCoverageRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRareDiseaseSiteCoverageRowWithDefaults

`func NewRareDiseaseSiteCoverageRowWithDefaults() *RareDiseaseSiteCoverageRow`

NewRareDiseaseSiteCoverageRowWithDefaults instantiates a new RareDiseaseSiteCoverageRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrphaCode

`func (o *RareDiseaseSiteCoverageRow) GetOrphaCode() string`

GetOrphaCode returns the OrphaCode field if non-nil, zero value otherwise.

### GetOrphaCodeOk

`func (o *RareDiseaseSiteCoverageRow) GetOrphaCodeOk() (*string, bool)`

GetOrphaCodeOk returns a tuple with the OrphaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphaCode

`func (o *RareDiseaseSiteCoverageRow) SetOrphaCode(v string)`

SetOrphaCode sets OrphaCode field to given value.

### HasOrphaCode

`func (o *RareDiseaseSiteCoverageRow) HasOrphaCode() bool`

HasOrphaCode returns a boolean if a field has been set.

### SetOrphaCodeNil

`func (o *RareDiseaseSiteCoverageRow) SetOrphaCodeNil(b bool)`

 SetOrphaCodeNil sets the value for OrphaCode to be an explicit nil

### UnsetOrphaCode
`func (o *RareDiseaseSiteCoverageRow) UnsetOrphaCode()`

UnsetOrphaCode ensures that no value is present for OrphaCode, not even an explicit nil
### GetConceptType

`func (o *RareDiseaseSiteCoverageRow) GetConceptType() string`

GetConceptType returns the ConceptType field if non-nil, zero value otherwise.

### GetConceptTypeOk

`func (o *RareDiseaseSiteCoverageRow) GetConceptTypeOk() (*string, bool)`

GetConceptTypeOk returns a tuple with the ConceptType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConceptType

`func (o *RareDiseaseSiteCoverageRow) SetConceptType(v string)`

SetConceptType sets ConceptType field to given value.

### HasConceptType

`func (o *RareDiseaseSiteCoverageRow) HasConceptType() bool`

HasConceptType returns a boolean if a field has been set.

### SetConceptTypeNil

`func (o *RareDiseaseSiteCoverageRow) SetConceptTypeNil(b bool)`

 SetConceptTypeNil sets the value for ConceptType to be an explicit nil

### UnsetConceptType
`func (o *RareDiseaseSiteCoverageRow) UnsetConceptType()`

UnsetConceptType ensures that no value is present for ConceptType, not even an explicit nil
### GetSnapshotDate

`func (o *RareDiseaseSiteCoverageRow) GetSnapshotDate() string`

GetSnapshotDate returns the SnapshotDate field if non-nil, zero value otherwise.

### GetSnapshotDateOk

`func (o *RareDiseaseSiteCoverageRow) GetSnapshotDateOk() (*string, bool)`

GetSnapshotDateOk returns a tuple with the SnapshotDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotDate

`func (o *RareDiseaseSiteCoverageRow) SetSnapshotDate(v string)`

SetSnapshotDate sets SnapshotDate field to given value.

### HasSnapshotDate

`func (o *RareDiseaseSiteCoverageRow) HasSnapshotDate() bool`

HasSnapshotDate returns a boolean if a field has been set.

### SetSnapshotDateNil

`func (o *RareDiseaseSiteCoverageRow) SetSnapshotDateNil(b bool)`

 SetSnapshotDateNil sets the value for SnapshotDate to be an explicit nil

### UnsetSnapshotDate
`func (o *RareDiseaseSiteCoverageRow) UnsetSnapshotDate()`

UnsetSnapshotDate ensures that no value is present for SnapshotDate, not even an explicit nil
### GetExactState

`func (o *RareDiseaseSiteCoverageRow) GetExactState() string`

GetExactState returns the ExactState field if non-nil, zero value otherwise.

### GetExactStateOk

`func (o *RareDiseaseSiteCoverageRow) GetExactStateOk() (*string, bool)`

GetExactStateOk returns a tuple with the ExactState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactState

`func (o *RareDiseaseSiteCoverageRow) SetExactState(v string)`

SetExactState sets ExactState field to given value.

### HasExactState

`func (o *RareDiseaseSiteCoverageRow) HasExactState() bool`

HasExactState returns a boolean if a field has been set.

### SetExactStateNil

`func (o *RareDiseaseSiteCoverageRow) SetExactStateNil(b bool)`

 SetExactStateNil sets the value for ExactState to be an explicit nil

### UnsetExactState
`func (o *RareDiseaseSiteCoverageRow) UnsetExactState()`

UnsetExactState ensures that no value is present for ExactState, not even an explicit nil
### GetExactNConditions

`func (o *RareDiseaseSiteCoverageRow) GetExactNConditions() int32`

GetExactNConditions returns the ExactNConditions field if non-nil, zero value otherwise.

### GetExactNConditionsOk

`func (o *RareDiseaseSiteCoverageRow) GetExactNConditionsOk() (*int32, bool)`

GetExactNConditionsOk returns a tuple with the ExactNConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactNConditions

`func (o *RareDiseaseSiteCoverageRow) SetExactNConditions(v int32)`

SetExactNConditions sets ExactNConditions field to given value.

### HasExactNConditions

`func (o *RareDiseaseSiteCoverageRow) HasExactNConditions() bool`

HasExactNConditions returns a boolean if a field has been set.

### SetExactNConditionsNil

`func (o *RareDiseaseSiteCoverageRow) SetExactNConditionsNil(b bool)`

 SetExactNConditionsNil sets the value for ExactNConditions to be an explicit nil

### UnsetExactNConditions
`func (o *RareDiseaseSiteCoverageRow) UnsetExactNConditions()`

UnsetExactNConditions ensures that no value is present for ExactNConditions, not even an explicit nil
### GetExactNTrials

`func (o *RareDiseaseSiteCoverageRow) GetExactNTrials() int32`

GetExactNTrials returns the ExactNTrials field if non-nil, zero value otherwise.

### GetExactNTrialsOk

`func (o *RareDiseaseSiteCoverageRow) GetExactNTrialsOk() (*int32, bool)`

GetExactNTrialsOk returns a tuple with the ExactNTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactNTrials

`func (o *RareDiseaseSiteCoverageRow) SetExactNTrials(v int32)`

SetExactNTrials sets ExactNTrials field to given value.

### HasExactNTrials

`func (o *RareDiseaseSiteCoverageRow) HasExactNTrials() bool`

HasExactNTrials returns a boolean if a field has been set.

### SetExactNTrialsNil

`func (o *RareDiseaseSiteCoverageRow) SetExactNTrialsNil(b bool)`

 SetExactNTrialsNil sets the value for ExactNTrials to be an explicit nil

### UnsetExactNTrials
`func (o *RareDiseaseSiteCoverageRow) UnsetExactNTrials()`

UnsetExactNTrials ensures that no value is present for ExactNTrials, not even an explicit nil
### GetExactNTrialsWithSite

`func (o *RareDiseaseSiteCoverageRow) GetExactNTrialsWithSite() int32`

GetExactNTrialsWithSite returns the ExactNTrialsWithSite field if non-nil, zero value otherwise.

### GetExactNTrialsWithSiteOk

`func (o *RareDiseaseSiteCoverageRow) GetExactNTrialsWithSiteOk() (*int32, bool)`

GetExactNTrialsWithSiteOk returns a tuple with the ExactNTrialsWithSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactNTrialsWithSite

`func (o *RareDiseaseSiteCoverageRow) SetExactNTrialsWithSite(v int32)`

SetExactNTrialsWithSite sets ExactNTrialsWithSite field to given value.

### HasExactNTrialsWithSite

`func (o *RareDiseaseSiteCoverageRow) HasExactNTrialsWithSite() bool`

HasExactNTrialsWithSite returns a boolean if a field has been set.

### SetExactNTrialsWithSiteNil

`func (o *RareDiseaseSiteCoverageRow) SetExactNTrialsWithSiteNil(b bool)`

 SetExactNTrialsWithSiteNil sets the value for ExactNTrialsWithSite to be an explicit nil

### UnsetExactNTrialsWithSite
`func (o *RareDiseaseSiteCoverageRow) UnsetExactNTrialsWithSite()`

UnsetExactNTrialsWithSite ensures that no value is present for ExactNTrialsWithSite, not even an explicit nil
### GetExactNSitesTotal

`func (o *RareDiseaseSiteCoverageRow) GetExactNSitesTotal() int32`

GetExactNSitesTotal returns the ExactNSitesTotal field if non-nil, zero value otherwise.

### GetExactNSitesTotalOk

`func (o *RareDiseaseSiteCoverageRow) GetExactNSitesTotalOk() (*int32, bool)`

GetExactNSitesTotalOk returns a tuple with the ExactNSitesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactNSitesTotal

`func (o *RareDiseaseSiteCoverageRow) SetExactNSitesTotal(v int32)`

SetExactNSitesTotal sets ExactNSitesTotal field to given value.

### HasExactNSitesTotal

`func (o *RareDiseaseSiteCoverageRow) HasExactNSitesTotal() bool`

HasExactNSitesTotal returns a boolean if a field has been set.

### SetExactNSitesTotalNil

`func (o *RareDiseaseSiteCoverageRow) SetExactNSitesTotalNil(b bool)`

 SetExactNSitesTotalNil sets the value for ExactNSitesTotal to be an explicit nil

### UnsetExactNSitesTotal
`func (o *RareDiseaseSiteCoverageRow) UnsetExactNSitesTotal()`

UnsetExactNSitesTotal ensures that no value is present for ExactNSitesTotal, not even an explicit nil
### GetExactNSitesEmitted

`func (o *RareDiseaseSiteCoverageRow) GetExactNSitesEmitted() int32`

GetExactNSitesEmitted returns the ExactNSitesEmitted field if non-nil, zero value otherwise.

### GetExactNSitesEmittedOk

`func (o *RareDiseaseSiteCoverageRow) GetExactNSitesEmittedOk() (*int32, bool)`

GetExactNSitesEmittedOk returns a tuple with the ExactNSitesEmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactNSitesEmitted

`func (o *RareDiseaseSiteCoverageRow) SetExactNSitesEmitted(v int32)`

SetExactNSitesEmitted sets ExactNSitesEmitted field to given value.

### HasExactNSitesEmitted

`func (o *RareDiseaseSiteCoverageRow) HasExactNSitesEmitted() bool`

HasExactNSitesEmitted returns a boolean if a field has been set.

### SetExactNSitesEmittedNil

`func (o *RareDiseaseSiteCoverageRow) SetExactNSitesEmittedNil(b bool)`

 SetExactNSitesEmittedNil sets the value for ExactNSitesEmitted to be an explicit nil

### UnsetExactNSitesEmitted
`func (o *RareDiseaseSiteCoverageRow) UnsetExactNSitesEmitted()`

UnsetExactNSitesEmitted ensures that no value is present for ExactNSitesEmitted, not even an explicit nil
### GetExactNSitesEmittedGeocoded

`func (o *RareDiseaseSiteCoverageRow) GetExactNSitesEmittedGeocoded() int32`

GetExactNSitesEmittedGeocoded returns the ExactNSitesEmittedGeocoded field if non-nil, zero value otherwise.

### GetExactNSitesEmittedGeocodedOk

`func (o *RareDiseaseSiteCoverageRow) GetExactNSitesEmittedGeocodedOk() (*int32, bool)`

GetExactNSitesEmittedGeocodedOk returns a tuple with the ExactNSitesEmittedGeocoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactNSitesEmittedGeocoded

`func (o *RareDiseaseSiteCoverageRow) SetExactNSitesEmittedGeocoded(v int32)`

SetExactNSitesEmittedGeocoded sets ExactNSitesEmittedGeocoded field to given value.

### HasExactNSitesEmittedGeocoded

`func (o *RareDiseaseSiteCoverageRow) HasExactNSitesEmittedGeocoded() bool`

HasExactNSitesEmittedGeocoded returns a boolean if a field has been set.

### SetExactNSitesEmittedGeocodedNil

`func (o *RareDiseaseSiteCoverageRow) SetExactNSitesEmittedGeocodedNil(b bool)`

 SetExactNSitesEmittedGeocodedNil sets the value for ExactNSitesEmittedGeocoded to be an explicit nil

### UnsetExactNSitesEmittedGeocoded
`func (o *RareDiseaseSiteCoverageRow) UnsetExactNSitesEmittedGeocoded()`

UnsetExactNSitesEmittedGeocoded ensures that no value is present for ExactNSitesEmittedGeocoded, not even an explicit nil
### GetExactNFacilitiesUngeocoded

`func (o *RareDiseaseSiteCoverageRow) GetExactNFacilitiesUngeocoded() int32`

GetExactNFacilitiesUngeocoded returns the ExactNFacilitiesUngeocoded field if non-nil, zero value otherwise.

### GetExactNFacilitiesUngeocodedOk

`func (o *RareDiseaseSiteCoverageRow) GetExactNFacilitiesUngeocodedOk() (*int32, bool)`

GetExactNFacilitiesUngeocodedOk returns a tuple with the ExactNFacilitiesUngeocoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactNFacilitiesUngeocoded

`func (o *RareDiseaseSiteCoverageRow) SetExactNFacilitiesUngeocoded(v int32)`

SetExactNFacilitiesUngeocoded sets ExactNFacilitiesUngeocoded field to given value.

### HasExactNFacilitiesUngeocoded

`func (o *RareDiseaseSiteCoverageRow) HasExactNFacilitiesUngeocoded() bool`

HasExactNFacilitiesUngeocoded returns a boolean if a field has been set.

### SetExactNFacilitiesUngeocodedNil

`func (o *RareDiseaseSiteCoverageRow) SetExactNFacilitiesUngeocodedNil(b bool)`

 SetExactNFacilitiesUngeocodedNil sets the value for ExactNFacilitiesUngeocoded to be an explicit nil

### UnsetExactNFacilitiesUngeocoded
`func (o *RareDiseaseSiteCoverageRow) UnsetExactNFacilitiesUngeocoded()`

UnsetExactNFacilitiesUngeocoded ensures that no value is present for ExactNFacilitiesUngeocoded, not even an explicit nil
### GetUnspecifiedState

`func (o *RareDiseaseSiteCoverageRow) GetUnspecifiedState() string`

GetUnspecifiedState returns the UnspecifiedState field if non-nil, zero value otherwise.

### GetUnspecifiedStateOk

`func (o *RareDiseaseSiteCoverageRow) GetUnspecifiedStateOk() (*string, bool)`

GetUnspecifiedStateOk returns a tuple with the UnspecifiedState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnspecifiedState

`func (o *RareDiseaseSiteCoverageRow) SetUnspecifiedState(v string)`

SetUnspecifiedState sets UnspecifiedState field to given value.

### HasUnspecifiedState

`func (o *RareDiseaseSiteCoverageRow) HasUnspecifiedState() bool`

HasUnspecifiedState returns a boolean if a field has been set.

### SetUnspecifiedStateNil

`func (o *RareDiseaseSiteCoverageRow) SetUnspecifiedStateNil(b bool)`

 SetUnspecifiedStateNil sets the value for UnspecifiedState to be an explicit nil

### UnsetUnspecifiedState
`func (o *RareDiseaseSiteCoverageRow) UnsetUnspecifiedState()`

UnsetUnspecifiedState ensures that no value is present for UnspecifiedState, not even an explicit nil
### GetUnspecifiedNConditions

`func (o *RareDiseaseSiteCoverageRow) GetUnspecifiedNConditions() int32`

GetUnspecifiedNConditions returns the UnspecifiedNConditions field if non-nil, zero value otherwise.

### GetUnspecifiedNConditionsOk

`func (o *RareDiseaseSiteCoverageRow) GetUnspecifiedNConditionsOk() (*int32, bool)`

GetUnspecifiedNConditionsOk returns a tuple with the UnspecifiedNConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnspecifiedNConditions

`func (o *RareDiseaseSiteCoverageRow) SetUnspecifiedNConditions(v int32)`

SetUnspecifiedNConditions sets UnspecifiedNConditions field to given value.

### HasUnspecifiedNConditions

`func (o *RareDiseaseSiteCoverageRow) HasUnspecifiedNConditions() bool`

HasUnspecifiedNConditions returns a boolean if a field has been set.

### SetUnspecifiedNConditionsNil

`func (o *RareDiseaseSiteCoverageRow) SetUnspecifiedNConditionsNil(b bool)`

 SetUnspecifiedNConditionsNil sets the value for UnspecifiedNConditions to be an explicit nil

### UnsetUnspecifiedNConditions
`func (o *RareDiseaseSiteCoverageRow) UnsetUnspecifiedNConditions()`

UnsetUnspecifiedNConditions ensures that no value is present for UnspecifiedNConditions, not even an explicit nil
### GetUnspecifiedNTrials

`func (o *RareDiseaseSiteCoverageRow) GetUnspecifiedNTrials() int32`

GetUnspecifiedNTrials returns the UnspecifiedNTrials field if non-nil, zero value otherwise.

### GetUnspecifiedNTrialsOk

`func (o *RareDiseaseSiteCoverageRow) GetUnspecifiedNTrialsOk() (*int32, bool)`

GetUnspecifiedNTrialsOk returns a tuple with the UnspecifiedNTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnspecifiedNTrials

`func (o *RareDiseaseSiteCoverageRow) SetUnspecifiedNTrials(v int32)`

SetUnspecifiedNTrials sets UnspecifiedNTrials field to given value.

### HasUnspecifiedNTrials

`func (o *RareDiseaseSiteCoverageRow) HasUnspecifiedNTrials() bool`

HasUnspecifiedNTrials returns a boolean if a field has been set.

### SetUnspecifiedNTrialsNil

`func (o *RareDiseaseSiteCoverageRow) SetUnspecifiedNTrialsNil(b bool)`

 SetUnspecifiedNTrialsNil sets the value for UnspecifiedNTrials to be an explicit nil

### UnsetUnspecifiedNTrials
`func (o *RareDiseaseSiteCoverageRow) UnsetUnspecifiedNTrials()`

UnsetUnspecifiedNTrials ensures that no value is present for UnspecifiedNTrials, not even an explicit nil
### GetUnspecifiedNTrialsWithSite

`func (o *RareDiseaseSiteCoverageRow) GetUnspecifiedNTrialsWithSite() int32`

GetUnspecifiedNTrialsWithSite returns the UnspecifiedNTrialsWithSite field if non-nil, zero value otherwise.

### GetUnspecifiedNTrialsWithSiteOk

`func (o *RareDiseaseSiteCoverageRow) GetUnspecifiedNTrialsWithSiteOk() (*int32, bool)`

GetUnspecifiedNTrialsWithSiteOk returns a tuple with the UnspecifiedNTrialsWithSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnspecifiedNTrialsWithSite

`func (o *RareDiseaseSiteCoverageRow) SetUnspecifiedNTrialsWithSite(v int32)`

SetUnspecifiedNTrialsWithSite sets UnspecifiedNTrialsWithSite field to given value.

### HasUnspecifiedNTrialsWithSite

`func (o *RareDiseaseSiteCoverageRow) HasUnspecifiedNTrialsWithSite() bool`

HasUnspecifiedNTrialsWithSite returns a boolean if a field has been set.

### SetUnspecifiedNTrialsWithSiteNil

`func (o *RareDiseaseSiteCoverageRow) SetUnspecifiedNTrialsWithSiteNil(b bool)`

 SetUnspecifiedNTrialsWithSiteNil sets the value for UnspecifiedNTrialsWithSite to be an explicit nil

### UnsetUnspecifiedNTrialsWithSite
`func (o *RareDiseaseSiteCoverageRow) UnsetUnspecifiedNTrialsWithSite()`

UnsetUnspecifiedNTrialsWithSite ensures that no value is present for UnspecifiedNTrialsWithSite, not even an explicit nil
### GetUnspecifiedNSitesTotal

`func (o *RareDiseaseSiteCoverageRow) GetUnspecifiedNSitesTotal() int32`

GetUnspecifiedNSitesTotal returns the UnspecifiedNSitesTotal field if non-nil, zero value otherwise.

### GetUnspecifiedNSitesTotalOk

`func (o *RareDiseaseSiteCoverageRow) GetUnspecifiedNSitesTotalOk() (*int32, bool)`

GetUnspecifiedNSitesTotalOk returns a tuple with the UnspecifiedNSitesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnspecifiedNSitesTotal

`func (o *RareDiseaseSiteCoverageRow) SetUnspecifiedNSitesTotal(v int32)`

SetUnspecifiedNSitesTotal sets UnspecifiedNSitesTotal field to given value.

### HasUnspecifiedNSitesTotal

`func (o *RareDiseaseSiteCoverageRow) HasUnspecifiedNSitesTotal() bool`

HasUnspecifiedNSitesTotal returns a boolean if a field has been set.

### SetUnspecifiedNSitesTotalNil

`func (o *RareDiseaseSiteCoverageRow) SetUnspecifiedNSitesTotalNil(b bool)`

 SetUnspecifiedNSitesTotalNil sets the value for UnspecifiedNSitesTotal to be an explicit nil

### UnsetUnspecifiedNSitesTotal
`func (o *RareDiseaseSiteCoverageRow) UnsetUnspecifiedNSitesTotal()`

UnsetUnspecifiedNSitesTotal ensures that no value is present for UnspecifiedNSitesTotal, not even an explicit nil
### GetUnspecifiedNSitesEmitted

`func (o *RareDiseaseSiteCoverageRow) GetUnspecifiedNSitesEmitted() int32`

GetUnspecifiedNSitesEmitted returns the UnspecifiedNSitesEmitted field if non-nil, zero value otherwise.

### GetUnspecifiedNSitesEmittedOk

`func (o *RareDiseaseSiteCoverageRow) GetUnspecifiedNSitesEmittedOk() (*int32, bool)`

GetUnspecifiedNSitesEmittedOk returns a tuple with the UnspecifiedNSitesEmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnspecifiedNSitesEmitted

`func (o *RareDiseaseSiteCoverageRow) SetUnspecifiedNSitesEmitted(v int32)`

SetUnspecifiedNSitesEmitted sets UnspecifiedNSitesEmitted field to given value.

### HasUnspecifiedNSitesEmitted

`func (o *RareDiseaseSiteCoverageRow) HasUnspecifiedNSitesEmitted() bool`

HasUnspecifiedNSitesEmitted returns a boolean if a field has been set.

### SetUnspecifiedNSitesEmittedNil

`func (o *RareDiseaseSiteCoverageRow) SetUnspecifiedNSitesEmittedNil(b bool)`

 SetUnspecifiedNSitesEmittedNil sets the value for UnspecifiedNSitesEmitted to be an explicit nil

### UnsetUnspecifiedNSitesEmitted
`func (o *RareDiseaseSiteCoverageRow) UnsetUnspecifiedNSitesEmitted()`

UnsetUnspecifiedNSitesEmitted ensures that no value is present for UnspecifiedNSitesEmitted, not even an explicit nil
### GetUnspecifiedNSitesEmittedGeocoded

`func (o *RareDiseaseSiteCoverageRow) GetUnspecifiedNSitesEmittedGeocoded() int32`

GetUnspecifiedNSitesEmittedGeocoded returns the UnspecifiedNSitesEmittedGeocoded field if non-nil, zero value otherwise.

### GetUnspecifiedNSitesEmittedGeocodedOk

`func (o *RareDiseaseSiteCoverageRow) GetUnspecifiedNSitesEmittedGeocodedOk() (*int32, bool)`

GetUnspecifiedNSitesEmittedGeocodedOk returns a tuple with the UnspecifiedNSitesEmittedGeocoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnspecifiedNSitesEmittedGeocoded

`func (o *RareDiseaseSiteCoverageRow) SetUnspecifiedNSitesEmittedGeocoded(v int32)`

SetUnspecifiedNSitesEmittedGeocoded sets UnspecifiedNSitesEmittedGeocoded field to given value.

### HasUnspecifiedNSitesEmittedGeocoded

`func (o *RareDiseaseSiteCoverageRow) HasUnspecifiedNSitesEmittedGeocoded() bool`

HasUnspecifiedNSitesEmittedGeocoded returns a boolean if a field has been set.

### SetUnspecifiedNSitesEmittedGeocodedNil

`func (o *RareDiseaseSiteCoverageRow) SetUnspecifiedNSitesEmittedGeocodedNil(b bool)`

 SetUnspecifiedNSitesEmittedGeocodedNil sets the value for UnspecifiedNSitesEmittedGeocoded to be an explicit nil

### UnsetUnspecifiedNSitesEmittedGeocoded
`func (o *RareDiseaseSiteCoverageRow) UnsetUnspecifiedNSitesEmittedGeocoded()`

UnsetUnspecifiedNSitesEmittedGeocoded ensures that no value is present for UnspecifiedNSitesEmittedGeocoded, not even an explicit nil
### GetUnspecifiedNFacilitiesUngeocoded

`func (o *RareDiseaseSiteCoverageRow) GetUnspecifiedNFacilitiesUngeocoded() int32`

GetUnspecifiedNFacilitiesUngeocoded returns the UnspecifiedNFacilitiesUngeocoded field if non-nil, zero value otherwise.

### GetUnspecifiedNFacilitiesUngeocodedOk

`func (o *RareDiseaseSiteCoverageRow) GetUnspecifiedNFacilitiesUngeocodedOk() (*int32, bool)`

GetUnspecifiedNFacilitiesUngeocodedOk returns a tuple with the UnspecifiedNFacilitiesUngeocoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnspecifiedNFacilitiesUngeocoded

`func (o *RareDiseaseSiteCoverageRow) SetUnspecifiedNFacilitiesUngeocoded(v int32)`

SetUnspecifiedNFacilitiesUngeocoded sets UnspecifiedNFacilitiesUngeocoded field to given value.

### HasUnspecifiedNFacilitiesUngeocoded

`func (o *RareDiseaseSiteCoverageRow) HasUnspecifiedNFacilitiesUngeocoded() bool`

HasUnspecifiedNFacilitiesUngeocoded returns a boolean if a field has been set.

### SetUnspecifiedNFacilitiesUngeocodedNil

`func (o *RareDiseaseSiteCoverageRow) SetUnspecifiedNFacilitiesUngeocodedNil(b bool)`

 SetUnspecifiedNFacilitiesUngeocodedNil sets the value for UnspecifiedNFacilitiesUngeocoded to be an explicit nil

### UnsetUnspecifiedNFacilitiesUngeocoded
`func (o *RareDiseaseSiteCoverageRow) UnsetUnspecifiedNFacilitiesUngeocoded()`

UnsetUnspecifiedNFacilitiesUngeocoded ensures that no value is present for UnspecifiedNFacilitiesUngeocoded, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


