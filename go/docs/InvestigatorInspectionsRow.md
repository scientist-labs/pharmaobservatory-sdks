# InvestigatorInspectionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvestigatorPersonId** | Pointer to **NullableString** | The investigator&#39;s resolved person key, the id in the web page address (/people/&lt;id&gt;). The entry key of the investigators object type. | [optional] 
**BimoRecordId** | Pointer to **NullableString** | The record id of the inspection and deficiency code. | [optional] 
**BimoInvestigatorId** | Pointer to **NullableString** | The CLIIL investigator code; two codes under one name may be different people. | [optional] 
**InspectedInvestigatorName** | Pointer to **NullableString** | The inspected investigator&#39;s name as CLIIL files it (&#39;last, first&#39;). | [optional] 
**InspectionDate** | Pointer to **NullableString** | The inspection&#39;s date. | [optional] 
**InspectionType** | Pointer to **NullableString** | The inspection type code as CLIIL files it (DA or FC). | [optional] 
**InspectionClassification** | Pointer to **NullableString** | The FDA inspection classification: NAI (no action indicated), VAI (voluntary action indicated) or OAI (official action indicated). | [optional] 
**DeficiencyCode** | Pointer to **NullableString** | The deficiency code as CLIIL codes it (null when the record carries none). | [optional] 
**HasDeficiency** | Pointer to **NullableBool** | True when the inspection cited a deficiency. | [optional] 
**SiteName** | Pointer to **NullableString** | The inspected site&#39;s name. | [optional] 
**SiteCity** | Pointer to **NullableString** | The inspected site&#39;s city. | [optional] 
**SiteState** | Pointer to **NullableString** | The inspected site&#39;s state. | [optional] 
**SiteCountry** | Pointer to **NullableString** | The inspected site&#39;s country. | [optional] 
**MatchTier** | Pointer to **NullableString** | How the CLIIL investigator was matched to the person: low is a bare-name match that reaches exactly one person (the only tier published; an ambiguous or unresolved name has no person and no row). | [optional] 
**MatchConfidence** | Pointer to **NullableFloat32** | The match&#39;s confidence (0.45 for the low tier). | [optional] 

## Methods

### NewInvestigatorInspectionsRow

`func NewInvestigatorInspectionsRow() *InvestigatorInspectionsRow`

NewInvestigatorInspectionsRow instantiates a new InvestigatorInspectionsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestigatorInspectionsRowWithDefaults

`func NewInvestigatorInspectionsRowWithDefaults() *InvestigatorInspectionsRow`

NewInvestigatorInspectionsRowWithDefaults instantiates a new InvestigatorInspectionsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvestigatorPersonId

`func (o *InvestigatorInspectionsRow) GetInvestigatorPersonId() string`

GetInvestigatorPersonId returns the InvestigatorPersonId field if non-nil, zero value otherwise.

### GetInvestigatorPersonIdOk

`func (o *InvestigatorInspectionsRow) GetInvestigatorPersonIdOk() (*string, bool)`

GetInvestigatorPersonIdOk returns a tuple with the InvestigatorPersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestigatorPersonId

`func (o *InvestigatorInspectionsRow) SetInvestigatorPersonId(v string)`

SetInvestigatorPersonId sets InvestigatorPersonId field to given value.

### HasInvestigatorPersonId

`func (o *InvestigatorInspectionsRow) HasInvestigatorPersonId() bool`

HasInvestigatorPersonId returns a boolean if a field has been set.

### SetInvestigatorPersonIdNil

`func (o *InvestigatorInspectionsRow) SetInvestigatorPersonIdNil(b bool)`

 SetInvestigatorPersonIdNil sets the value for InvestigatorPersonId to be an explicit nil

### UnsetInvestigatorPersonId
`func (o *InvestigatorInspectionsRow) UnsetInvestigatorPersonId()`

UnsetInvestigatorPersonId ensures that no value is present for InvestigatorPersonId, not even an explicit nil
### GetBimoRecordId

`func (o *InvestigatorInspectionsRow) GetBimoRecordId() string`

GetBimoRecordId returns the BimoRecordId field if non-nil, zero value otherwise.

### GetBimoRecordIdOk

`func (o *InvestigatorInspectionsRow) GetBimoRecordIdOk() (*string, bool)`

GetBimoRecordIdOk returns a tuple with the BimoRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBimoRecordId

`func (o *InvestigatorInspectionsRow) SetBimoRecordId(v string)`

SetBimoRecordId sets BimoRecordId field to given value.

### HasBimoRecordId

`func (o *InvestigatorInspectionsRow) HasBimoRecordId() bool`

HasBimoRecordId returns a boolean if a field has been set.

### SetBimoRecordIdNil

`func (o *InvestigatorInspectionsRow) SetBimoRecordIdNil(b bool)`

 SetBimoRecordIdNil sets the value for BimoRecordId to be an explicit nil

### UnsetBimoRecordId
`func (o *InvestigatorInspectionsRow) UnsetBimoRecordId()`

UnsetBimoRecordId ensures that no value is present for BimoRecordId, not even an explicit nil
### GetBimoInvestigatorId

`func (o *InvestigatorInspectionsRow) GetBimoInvestigatorId() string`

GetBimoInvestigatorId returns the BimoInvestigatorId field if non-nil, zero value otherwise.

### GetBimoInvestigatorIdOk

`func (o *InvestigatorInspectionsRow) GetBimoInvestigatorIdOk() (*string, bool)`

GetBimoInvestigatorIdOk returns a tuple with the BimoInvestigatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBimoInvestigatorId

`func (o *InvestigatorInspectionsRow) SetBimoInvestigatorId(v string)`

SetBimoInvestigatorId sets BimoInvestigatorId field to given value.

### HasBimoInvestigatorId

`func (o *InvestigatorInspectionsRow) HasBimoInvestigatorId() bool`

HasBimoInvestigatorId returns a boolean if a field has been set.

### SetBimoInvestigatorIdNil

`func (o *InvestigatorInspectionsRow) SetBimoInvestigatorIdNil(b bool)`

 SetBimoInvestigatorIdNil sets the value for BimoInvestigatorId to be an explicit nil

### UnsetBimoInvestigatorId
`func (o *InvestigatorInspectionsRow) UnsetBimoInvestigatorId()`

UnsetBimoInvestigatorId ensures that no value is present for BimoInvestigatorId, not even an explicit nil
### GetInspectedInvestigatorName

`func (o *InvestigatorInspectionsRow) GetInspectedInvestigatorName() string`

GetInspectedInvestigatorName returns the InspectedInvestigatorName field if non-nil, zero value otherwise.

### GetInspectedInvestigatorNameOk

`func (o *InvestigatorInspectionsRow) GetInspectedInvestigatorNameOk() (*string, bool)`

GetInspectedInvestigatorNameOk returns a tuple with the InspectedInvestigatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectedInvestigatorName

`func (o *InvestigatorInspectionsRow) SetInspectedInvestigatorName(v string)`

SetInspectedInvestigatorName sets InspectedInvestigatorName field to given value.

### HasInspectedInvestigatorName

`func (o *InvestigatorInspectionsRow) HasInspectedInvestigatorName() bool`

HasInspectedInvestigatorName returns a boolean if a field has been set.

### SetInspectedInvestigatorNameNil

`func (o *InvestigatorInspectionsRow) SetInspectedInvestigatorNameNil(b bool)`

 SetInspectedInvestigatorNameNil sets the value for InspectedInvestigatorName to be an explicit nil

### UnsetInspectedInvestigatorName
`func (o *InvestigatorInspectionsRow) UnsetInspectedInvestigatorName()`

UnsetInspectedInvestigatorName ensures that no value is present for InspectedInvestigatorName, not even an explicit nil
### GetInspectionDate

`func (o *InvestigatorInspectionsRow) GetInspectionDate() string`

GetInspectionDate returns the InspectionDate field if non-nil, zero value otherwise.

### GetInspectionDateOk

`func (o *InvestigatorInspectionsRow) GetInspectionDateOk() (*string, bool)`

GetInspectionDateOk returns a tuple with the InspectionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectionDate

`func (o *InvestigatorInspectionsRow) SetInspectionDate(v string)`

SetInspectionDate sets InspectionDate field to given value.

### HasInspectionDate

`func (o *InvestigatorInspectionsRow) HasInspectionDate() bool`

HasInspectionDate returns a boolean if a field has been set.

### SetInspectionDateNil

`func (o *InvestigatorInspectionsRow) SetInspectionDateNil(b bool)`

 SetInspectionDateNil sets the value for InspectionDate to be an explicit nil

### UnsetInspectionDate
`func (o *InvestigatorInspectionsRow) UnsetInspectionDate()`

UnsetInspectionDate ensures that no value is present for InspectionDate, not even an explicit nil
### GetInspectionType

`func (o *InvestigatorInspectionsRow) GetInspectionType() string`

GetInspectionType returns the InspectionType field if non-nil, zero value otherwise.

### GetInspectionTypeOk

`func (o *InvestigatorInspectionsRow) GetInspectionTypeOk() (*string, bool)`

GetInspectionTypeOk returns a tuple with the InspectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectionType

`func (o *InvestigatorInspectionsRow) SetInspectionType(v string)`

SetInspectionType sets InspectionType field to given value.

### HasInspectionType

`func (o *InvestigatorInspectionsRow) HasInspectionType() bool`

HasInspectionType returns a boolean if a field has been set.

### SetInspectionTypeNil

`func (o *InvestigatorInspectionsRow) SetInspectionTypeNil(b bool)`

 SetInspectionTypeNil sets the value for InspectionType to be an explicit nil

### UnsetInspectionType
`func (o *InvestigatorInspectionsRow) UnsetInspectionType()`

UnsetInspectionType ensures that no value is present for InspectionType, not even an explicit nil
### GetInspectionClassification

`func (o *InvestigatorInspectionsRow) GetInspectionClassification() string`

GetInspectionClassification returns the InspectionClassification field if non-nil, zero value otherwise.

### GetInspectionClassificationOk

`func (o *InvestigatorInspectionsRow) GetInspectionClassificationOk() (*string, bool)`

GetInspectionClassificationOk returns a tuple with the InspectionClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectionClassification

`func (o *InvestigatorInspectionsRow) SetInspectionClassification(v string)`

SetInspectionClassification sets InspectionClassification field to given value.

### HasInspectionClassification

`func (o *InvestigatorInspectionsRow) HasInspectionClassification() bool`

HasInspectionClassification returns a boolean if a field has been set.

### SetInspectionClassificationNil

`func (o *InvestigatorInspectionsRow) SetInspectionClassificationNil(b bool)`

 SetInspectionClassificationNil sets the value for InspectionClassification to be an explicit nil

### UnsetInspectionClassification
`func (o *InvestigatorInspectionsRow) UnsetInspectionClassification()`

UnsetInspectionClassification ensures that no value is present for InspectionClassification, not even an explicit nil
### GetDeficiencyCode

`func (o *InvestigatorInspectionsRow) GetDeficiencyCode() string`

GetDeficiencyCode returns the DeficiencyCode field if non-nil, zero value otherwise.

### GetDeficiencyCodeOk

`func (o *InvestigatorInspectionsRow) GetDeficiencyCodeOk() (*string, bool)`

GetDeficiencyCodeOk returns a tuple with the DeficiencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeficiencyCode

`func (o *InvestigatorInspectionsRow) SetDeficiencyCode(v string)`

SetDeficiencyCode sets DeficiencyCode field to given value.

### HasDeficiencyCode

`func (o *InvestigatorInspectionsRow) HasDeficiencyCode() bool`

HasDeficiencyCode returns a boolean if a field has been set.

### SetDeficiencyCodeNil

`func (o *InvestigatorInspectionsRow) SetDeficiencyCodeNil(b bool)`

 SetDeficiencyCodeNil sets the value for DeficiencyCode to be an explicit nil

### UnsetDeficiencyCode
`func (o *InvestigatorInspectionsRow) UnsetDeficiencyCode()`

UnsetDeficiencyCode ensures that no value is present for DeficiencyCode, not even an explicit nil
### GetHasDeficiency

`func (o *InvestigatorInspectionsRow) GetHasDeficiency() bool`

GetHasDeficiency returns the HasDeficiency field if non-nil, zero value otherwise.

### GetHasDeficiencyOk

`func (o *InvestigatorInspectionsRow) GetHasDeficiencyOk() (*bool, bool)`

GetHasDeficiencyOk returns a tuple with the HasDeficiency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDeficiency

`func (o *InvestigatorInspectionsRow) SetHasDeficiency(v bool)`

SetHasDeficiency sets HasDeficiency field to given value.

### HasHasDeficiency

`func (o *InvestigatorInspectionsRow) HasHasDeficiency() bool`

HasHasDeficiency returns a boolean if a field has been set.

### SetHasDeficiencyNil

`func (o *InvestigatorInspectionsRow) SetHasDeficiencyNil(b bool)`

 SetHasDeficiencyNil sets the value for HasDeficiency to be an explicit nil

### UnsetHasDeficiency
`func (o *InvestigatorInspectionsRow) UnsetHasDeficiency()`

UnsetHasDeficiency ensures that no value is present for HasDeficiency, not even an explicit nil
### GetSiteName

`func (o *InvestigatorInspectionsRow) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *InvestigatorInspectionsRow) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *InvestigatorInspectionsRow) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *InvestigatorInspectionsRow) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### SetSiteNameNil

`func (o *InvestigatorInspectionsRow) SetSiteNameNil(b bool)`

 SetSiteNameNil sets the value for SiteName to be an explicit nil

### UnsetSiteName
`func (o *InvestigatorInspectionsRow) UnsetSiteName()`

UnsetSiteName ensures that no value is present for SiteName, not even an explicit nil
### GetSiteCity

`func (o *InvestigatorInspectionsRow) GetSiteCity() string`

GetSiteCity returns the SiteCity field if non-nil, zero value otherwise.

### GetSiteCityOk

`func (o *InvestigatorInspectionsRow) GetSiteCityOk() (*string, bool)`

GetSiteCityOk returns a tuple with the SiteCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteCity

`func (o *InvestigatorInspectionsRow) SetSiteCity(v string)`

SetSiteCity sets SiteCity field to given value.

### HasSiteCity

`func (o *InvestigatorInspectionsRow) HasSiteCity() bool`

HasSiteCity returns a boolean if a field has been set.

### SetSiteCityNil

`func (o *InvestigatorInspectionsRow) SetSiteCityNil(b bool)`

 SetSiteCityNil sets the value for SiteCity to be an explicit nil

### UnsetSiteCity
`func (o *InvestigatorInspectionsRow) UnsetSiteCity()`

UnsetSiteCity ensures that no value is present for SiteCity, not even an explicit nil
### GetSiteState

`func (o *InvestigatorInspectionsRow) GetSiteState() string`

GetSiteState returns the SiteState field if non-nil, zero value otherwise.

### GetSiteStateOk

`func (o *InvestigatorInspectionsRow) GetSiteStateOk() (*string, bool)`

GetSiteStateOk returns a tuple with the SiteState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteState

`func (o *InvestigatorInspectionsRow) SetSiteState(v string)`

SetSiteState sets SiteState field to given value.

### HasSiteState

`func (o *InvestigatorInspectionsRow) HasSiteState() bool`

HasSiteState returns a boolean if a field has been set.

### SetSiteStateNil

`func (o *InvestigatorInspectionsRow) SetSiteStateNil(b bool)`

 SetSiteStateNil sets the value for SiteState to be an explicit nil

### UnsetSiteState
`func (o *InvestigatorInspectionsRow) UnsetSiteState()`

UnsetSiteState ensures that no value is present for SiteState, not even an explicit nil
### GetSiteCountry

`func (o *InvestigatorInspectionsRow) GetSiteCountry() string`

GetSiteCountry returns the SiteCountry field if non-nil, zero value otherwise.

### GetSiteCountryOk

`func (o *InvestigatorInspectionsRow) GetSiteCountryOk() (*string, bool)`

GetSiteCountryOk returns a tuple with the SiteCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteCountry

`func (o *InvestigatorInspectionsRow) SetSiteCountry(v string)`

SetSiteCountry sets SiteCountry field to given value.

### HasSiteCountry

`func (o *InvestigatorInspectionsRow) HasSiteCountry() bool`

HasSiteCountry returns a boolean if a field has been set.

### SetSiteCountryNil

`func (o *InvestigatorInspectionsRow) SetSiteCountryNil(b bool)`

 SetSiteCountryNil sets the value for SiteCountry to be an explicit nil

### UnsetSiteCountry
`func (o *InvestigatorInspectionsRow) UnsetSiteCountry()`

UnsetSiteCountry ensures that no value is present for SiteCountry, not even an explicit nil
### GetMatchTier

`func (o *InvestigatorInspectionsRow) GetMatchTier() string`

GetMatchTier returns the MatchTier field if non-nil, zero value otherwise.

### GetMatchTierOk

`func (o *InvestigatorInspectionsRow) GetMatchTierOk() (*string, bool)`

GetMatchTierOk returns a tuple with the MatchTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchTier

`func (o *InvestigatorInspectionsRow) SetMatchTier(v string)`

SetMatchTier sets MatchTier field to given value.

### HasMatchTier

`func (o *InvestigatorInspectionsRow) HasMatchTier() bool`

HasMatchTier returns a boolean if a field has been set.

### SetMatchTierNil

`func (o *InvestigatorInspectionsRow) SetMatchTierNil(b bool)`

 SetMatchTierNil sets the value for MatchTier to be an explicit nil

### UnsetMatchTier
`func (o *InvestigatorInspectionsRow) UnsetMatchTier()`

UnsetMatchTier ensures that no value is present for MatchTier, not even an explicit nil
### GetMatchConfidence

`func (o *InvestigatorInspectionsRow) GetMatchConfidence() float32`

GetMatchConfidence returns the MatchConfidence field if non-nil, zero value otherwise.

### GetMatchConfidenceOk

`func (o *InvestigatorInspectionsRow) GetMatchConfidenceOk() (*float32, bool)`

GetMatchConfidenceOk returns a tuple with the MatchConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchConfidence

`func (o *InvestigatorInspectionsRow) SetMatchConfidence(v float32)`

SetMatchConfidence sets MatchConfidence field to given value.

### HasMatchConfidence

`func (o *InvestigatorInspectionsRow) HasMatchConfidence() bool`

HasMatchConfidence returns a boolean if a field has been set.

### SetMatchConfidenceNil

`func (o *InvestigatorInspectionsRow) SetMatchConfidenceNil(b bool)`

 SetMatchConfidenceNil sets the value for MatchConfidence to be an explicit nil

### UnsetMatchConfidence
`func (o *InvestigatorInspectionsRow) UnsetMatchConfidence()`

UnsetMatchConfidence ensures that no value is present for MatchConfidence, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


