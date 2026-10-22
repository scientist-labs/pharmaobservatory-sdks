# RwdStudyDataSourcesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StudyId** | Pointer to **NullableString** | The register&#39;s own Study ID. The entry key of the rwd_studies object type; the EU PAS number is an attribute. | [optional] 
**ItemOrdinal** | Pointer to **NullableInt32** | The position of the item in the register&#39;s own list, from 1. | [optional] 
**DataSourceName** | Pointer to **NullableString** | The data source name as the study named it, from the register&#39;s Data source(s) list. | [optional] 
**DataSourceNameKey** | Pointer to **NullableString** | The normalised name key the study register is joined to the data source on (the study register carries no data source id). | [optional] 
**MatchedDataSourceRegister** | Pointer to **NullableBool** | Whether the name matches a registered data source. | [optional] 
**DataSourceId** | Pointer to **NullableString** | The register Data source ID. The entry key of the rwd_sources object type. On a row built from a study&#39;s data source name it is null unless the name reaches exactly one registered source (see data_source_ids). | [optional] 
**DataSourceIds** | Pointer to **[]string** | Every registered data source id the name matches (more than one when two entries share a name). | [optional] 
**RegisterRowCount** | Pointer to **NullableInt32** | How many register entries share the name. | [optional] 
**StudyStatus** | Pointer to **NullableString** | The study&#39;s status as the register words it (Planned, Ongoing, Finalised, and so on). | [optional] 
**MandateStratum** | Pointer to **NullableString** | mandated when a regulator required the study (a post-authorisation obligation), voluntary when the registrant says it did not, unknown when unanswered. Delivery timeliness is only comparable within a stratum. | [optional] 
**RequiredByRegulator** | Pointer to **NullableString** | The registrant&#39;s answer to whether a regulatory body required the study (Yes, No or Unknown). | [optional] 
**StudyStartActualOn** | Pointer to **NullableString** | Actual study start date. | [optional] 
**StartedInLast5Years** | Pointer to **NullableBool** | Whether the study started within five years of the snapshot. | [optional] 
**ConditionCount** | Pointer to **NullableInt32** | How many distinct condition terms the study names. | [optional] 
**ConductingInstitutionCount** | Pointer to **NullableInt32** | How many conducting institutions the study names. | [optional] 
**SnapshotLabel** | Pointer to **NullableString** | The label of the catalogue snapshot the row was read from (a date). | [optional] 
**AccessedMonth** | Pointer to **NullableString** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRwdStudyDataSourcesRow

`func NewRwdStudyDataSourcesRow() *RwdStudyDataSourcesRow`

NewRwdStudyDataSourcesRow instantiates a new RwdStudyDataSourcesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRwdStudyDataSourcesRowWithDefaults

`func NewRwdStudyDataSourcesRowWithDefaults() *RwdStudyDataSourcesRow`

NewRwdStudyDataSourcesRowWithDefaults instantiates a new RwdStudyDataSourcesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStudyId

`func (o *RwdStudyDataSourcesRow) GetStudyId() string`

GetStudyId returns the StudyId field if non-nil, zero value otherwise.

### GetStudyIdOk

`func (o *RwdStudyDataSourcesRow) GetStudyIdOk() (*string, bool)`

GetStudyIdOk returns a tuple with the StudyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyId

`func (o *RwdStudyDataSourcesRow) SetStudyId(v string)`

SetStudyId sets StudyId field to given value.

### HasStudyId

`func (o *RwdStudyDataSourcesRow) HasStudyId() bool`

HasStudyId returns a boolean if a field has been set.

### SetStudyIdNil

`func (o *RwdStudyDataSourcesRow) SetStudyIdNil(b bool)`

 SetStudyIdNil sets the value for StudyId to be an explicit nil

### UnsetStudyId
`func (o *RwdStudyDataSourcesRow) UnsetStudyId()`

UnsetStudyId ensures that no value is present for StudyId, not even an explicit nil
### GetItemOrdinal

`func (o *RwdStudyDataSourcesRow) GetItemOrdinal() int32`

GetItemOrdinal returns the ItemOrdinal field if non-nil, zero value otherwise.

### GetItemOrdinalOk

`func (o *RwdStudyDataSourcesRow) GetItemOrdinalOk() (*int32, bool)`

GetItemOrdinalOk returns a tuple with the ItemOrdinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemOrdinal

`func (o *RwdStudyDataSourcesRow) SetItemOrdinal(v int32)`

SetItemOrdinal sets ItemOrdinal field to given value.

### HasItemOrdinal

`func (o *RwdStudyDataSourcesRow) HasItemOrdinal() bool`

HasItemOrdinal returns a boolean if a field has been set.

### SetItemOrdinalNil

`func (o *RwdStudyDataSourcesRow) SetItemOrdinalNil(b bool)`

 SetItemOrdinalNil sets the value for ItemOrdinal to be an explicit nil

### UnsetItemOrdinal
`func (o *RwdStudyDataSourcesRow) UnsetItemOrdinal()`

UnsetItemOrdinal ensures that no value is present for ItemOrdinal, not even an explicit nil
### GetDataSourceName

`func (o *RwdStudyDataSourcesRow) GetDataSourceName() string`

GetDataSourceName returns the DataSourceName field if non-nil, zero value otherwise.

### GetDataSourceNameOk

`func (o *RwdStudyDataSourcesRow) GetDataSourceNameOk() (*string, bool)`

GetDataSourceNameOk returns a tuple with the DataSourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceName

`func (o *RwdStudyDataSourcesRow) SetDataSourceName(v string)`

SetDataSourceName sets DataSourceName field to given value.

### HasDataSourceName

`func (o *RwdStudyDataSourcesRow) HasDataSourceName() bool`

HasDataSourceName returns a boolean if a field has been set.

### SetDataSourceNameNil

`func (o *RwdStudyDataSourcesRow) SetDataSourceNameNil(b bool)`

 SetDataSourceNameNil sets the value for DataSourceName to be an explicit nil

### UnsetDataSourceName
`func (o *RwdStudyDataSourcesRow) UnsetDataSourceName()`

UnsetDataSourceName ensures that no value is present for DataSourceName, not even an explicit nil
### GetDataSourceNameKey

`func (o *RwdStudyDataSourcesRow) GetDataSourceNameKey() string`

GetDataSourceNameKey returns the DataSourceNameKey field if non-nil, zero value otherwise.

### GetDataSourceNameKeyOk

`func (o *RwdStudyDataSourcesRow) GetDataSourceNameKeyOk() (*string, bool)`

GetDataSourceNameKeyOk returns a tuple with the DataSourceNameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceNameKey

`func (o *RwdStudyDataSourcesRow) SetDataSourceNameKey(v string)`

SetDataSourceNameKey sets DataSourceNameKey field to given value.

### HasDataSourceNameKey

`func (o *RwdStudyDataSourcesRow) HasDataSourceNameKey() bool`

HasDataSourceNameKey returns a boolean if a field has been set.

### SetDataSourceNameKeyNil

`func (o *RwdStudyDataSourcesRow) SetDataSourceNameKeyNil(b bool)`

 SetDataSourceNameKeyNil sets the value for DataSourceNameKey to be an explicit nil

### UnsetDataSourceNameKey
`func (o *RwdStudyDataSourcesRow) UnsetDataSourceNameKey()`

UnsetDataSourceNameKey ensures that no value is present for DataSourceNameKey, not even an explicit nil
### GetMatchedDataSourceRegister

`func (o *RwdStudyDataSourcesRow) GetMatchedDataSourceRegister() bool`

GetMatchedDataSourceRegister returns the MatchedDataSourceRegister field if non-nil, zero value otherwise.

### GetMatchedDataSourceRegisterOk

`func (o *RwdStudyDataSourcesRow) GetMatchedDataSourceRegisterOk() (*bool, bool)`

GetMatchedDataSourceRegisterOk returns a tuple with the MatchedDataSourceRegister field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedDataSourceRegister

`func (o *RwdStudyDataSourcesRow) SetMatchedDataSourceRegister(v bool)`

SetMatchedDataSourceRegister sets MatchedDataSourceRegister field to given value.

### HasMatchedDataSourceRegister

`func (o *RwdStudyDataSourcesRow) HasMatchedDataSourceRegister() bool`

HasMatchedDataSourceRegister returns a boolean if a field has been set.

### SetMatchedDataSourceRegisterNil

`func (o *RwdStudyDataSourcesRow) SetMatchedDataSourceRegisterNil(b bool)`

 SetMatchedDataSourceRegisterNil sets the value for MatchedDataSourceRegister to be an explicit nil

### UnsetMatchedDataSourceRegister
`func (o *RwdStudyDataSourcesRow) UnsetMatchedDataSourceRegister()`

UnsetMatchedDataSourceRegister ensures that no value is present for MatchedDataSourceRegister, not even an explicit nil
### GetDataSourceId

`func (o *RwdStudyDataSourcesRow) GetDataSourceId() string`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *RwdStudyDataSourcesRow) GetDataSourceIdOk() (*string, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *RwdStudyDataSourcesRow) SetDataSourceId(v string)`

SetDataSourceId sets DataSourceId field to given value.

### HasDataSourceId

`func (o *RwdStudyDataSourcesRow) HasDataSourceId() bool`

HasDataSourceId returns a boolean if a field has been set.

### SetDataSourceIdNil

`func (o *RwdStudyDataSourcesRow) SetDataSourceIdNil(b bool)`

 SetDataSourceIdNil sets the value for DataSourceId to be an explicit nil

### UnsetDataSourceId
`func (o *RwdStudyDataSourcesRow) UnsetDataSourceId()`

UnsetDataSourceId ensures that no value is present for DataSourceId, not even an explicit nil
### GetDataSourceIds

`func (o *RwdStudyDataSourcesRow) GetDataSourceIds() []string`

GetDataSourceIds returns the DataSourceIds field if non-nil, zero value otherwise.

### GetDataSourceIdsOk

`func (o *RwdStudyDataSourcesRow) GetDataSourceIdsOk() (*[]string, bool)`

GetDataSourceIdsOk returns a tuple with the DataSourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceIds

`func (o *RwdStudyDataSourcesRow) SetDataSourceIds(v []string)`

SetDataSourceIds sets DataSourceIds field to given value.

### HasDataSourceIds

`func (o *RwdStudyDataSourcesRow) HasDataSourceIds() bool`

HasDataSourceIds returns a boolean if a field has been set.

### SetDataSourceIdsNil

`func (o *RwdStudyDataSourcesRow) SetDataSourceIdsNil(b bool)`

 SetDataSourceIdsNil sets the value for DataSourceIds to be an explicit nil

### UnsetDataSourceIds
`func (o *RwdStudyDataSourcesRow) UnsetDataSourceIds()`

UnsetDataSourceIds ensures that no value is present for DataSourceIds, not even an explicit nil
### GetRegisterRowCount

`func (o *RwdStudyDataSourcesRow) GetRegisterRowCount() int32`

GetRegisterRowCount returns the RegisterRowCount field if non-nil, zero value otherwise.

### GetRegisterRowCountOk

`func (o *RwdStudyDataSourcesRow) GetRegisterRowCountOk() (*int32, bool)`

GetRegisterRowCountOk returns a tuple with the RegisterRowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterRowCount

`func (o *RwdStudyDataSourcesRow) SetRegisterRowCount(v int32)`

SetRegisterRowCount sets RegisterRowCount field to given value.

### HasRegisterRowCount

`func (o *RwdStudyDataSourcesRow) HasRegisterRowCount() bool`

HasRegisterRowCount returns a boolean if a field has been set.

### SetRegisterRowCountNil

`func (o *RwdStudyDataSourcesRow) SetRegisterRowCountNil(b bool)`

 SetRegisterRowCountNil sets the value for RegisterRowCount to be an explicit nil

### UnsetRegisterRowCount
`func (o *RwdStudyDataSourcesRow) UnsetRegisterRowCount()`

UnsetRegisterRowCount ensures that no value is present for RegisterRowCount, not even an explicit nil
### GetStudyStatus

`func (o *RwdStudyDataSourcesRow) GetStudyStatus() string`

GetStudyStatus returns the StudyStatus field if non-nil, zero value otherwise.

### GetStudyStatusOk

`func (o *RwdStudyDataSourcesRow) GetStudyStatusOk() (*string, bool)`

GetStudyStatusOk returns a tuple with the StudyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyStatus

`func (o *RwdStudyDataSourcesRow) SetStudyStatus(v string)`

SetStudyStatus sets StudyStatus field to given value.

### HasStudyStatus

`func (o *RwdStudyDataSourcesRow) HasStudyStatus() bool`

HasStudyStatus returns a boolean if a field has been set.

### SetStudyStatusNil

`func (o *RwdStudyDataSourcesRow) SetStudyStatusNil(b bool)`

 SetStudyStatusNil sets the value for StudyStatus to be an explicit nil

### UnsetStudyStatus
`func (o *RwdStudyDataSourcesRow) UnsetStudyStatus()`

UnsetStudyStatus ensures that no value is present for StudyStatus, not even an explicit nil
### GetMandateStratum

`func (o *RwdStudyDataSourcesRow) GetMandateStratum() string`

GetMandateStratum returns the MandateStratum field if non-nil, zero value otherwise.

### GetMandateStratumOk

`func (o *RwdStudyDataSourcesRow) GetMandateStratumOk() (*string, bool)`

GetMandateStratumOk returns a tuple with the MandateStratum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateStratum

`func (o *RwdStudyDataSourcesRow) SetMandateStratum(v string)`

SetMandateStratum sets MandateStratum field to given value.

### HasMandateStratum

`func (o *RwdStudyDataSourcesRow) HasMandateStratum() bool`

HasMandateStratum returns a boolean if a field has been set.

### SetMandateStratumNil

`func (o *RwdStudyDataSourcesRow) SetMandateStratumNil(b bool)`

 SetMandateStratumNil sets the value for MandateStratum to be an explicit nil

### UnsetMandateStratum
`func (o *RwdStudyDataSourcesRow) UnsetMandateStratum()`

UnsetMandateStratum ensures that no value is present for MandateStratum, not even an explicit nil
### GetRequiredByRegulator

`func (o *RwdStudyDataSourcesRow) GetRequiredByRegulator() string`

GetRequiredByRegulator returns the RequiredByRegulator field if non-nil, zero value otherwise.

### GetRequiredByRegulatorOk

`func (o *RwdStudyDataSourcesRow) GetRequiredByRegulatorOk() (*string, bool)`

GetRequiredByRegulatorOk returns a tuple with the RequiredByRegulator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredByRegulator

`func (o *RwdStudyDataSourcesRow) SetRequiredByRegulator(v string)`

SetRequiredByRegulator sets RequiredByRegulator field to given value.

### HasRequiredByRegulator

`func (o *RwdStudyDataSourcesRow) HasRequiredByRegulator() bool`

HasRequiredByRegulator returns a boolean if a field has been set.

### SetRequiredByRegulatorNil

`func (o *RwdStudyDataSourcesRow) SetRequiredByRegulatorNil(b bool)`

 SetRequiredByRegulatorNil sets the value for RequiredByRegulator to be an explicit nil

### UnsetRequiredByRegulator
`func (o *RwdStudyDataSourcesRow) UnsetRequiredByRegulator()`

UnsetRequiredByRegulator ensures that no value is present for RequiredByRegulator, not even an explicit nil
### GetStudyStartActualOn

`func (o *RwdStudyDataSourcesRow) GetStudyStartActualOn() string`

GetStudyStartActualOn returns the StudyStartActualOn field if non-nil, zero value otherwise.

### GetStudyStartActualOnOk

`func (o *RwdStudyDataSourcesRow) GetStudyStartActualOnOk() (*string, bool)`

GetStudyStartActualOnOk returns a tuple with the StudyStartActualOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyStartActualOn

`func (o *RwdStudyDataSourcesRow) SetStudyStartActualOn(v string)`

SetStudyStartActualOn sets StudyStartActualOn field to given value.

### HasStudyStartActualOn

`func (o *RwdStudyDataSourcesRow) HasStudyStartActualOn() bool`

HasStudyStartActualOn returns a boolean if a field has been set.

### SetStudyStartActualOnNil

`func (o *RwdStudyDataSourcesRow) SetStudyStartActualOnNil(b bool)`

 SetStudyStartActualOnNil sets the value for StudyStartActualOn to be an explicit nil

### UnsetStudyStartActualOn
`func (o *RwdStudyDataSourcesRow) UnsetStudyStartActualOn()`

UnsetStudyStartActualOn ensures that no value is present for StudyStartActualOn, not even an explicit nil
### GetStartedInLast5Years

`func (o *RwdStudyDataSourcesRow) GetStartedInLast5Years() bool`

GetStartedInLast5Years returns the StartedInLast5Years field if non-nil, zero value otherwise.

### GetStartedInLast5YearsOk

`func (o *RwdStudyDataSourcesRow) GetStartedInLast5YearsOk() (*bool, bool)`

GetStartedInLast5YearsOk returns a tuple with the StartedInLast5Years field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedInLast5Years

`func (o *RwdStudyDataSourcesRow) SetStartedInLast5Years(v bool)`

SetStartedInLast5Years sets StartedInLast5Years field to given value.

### HasStartedInLast5Years

`func (o *RwdStudyDataSourcesRow) HasStartedInLast5Years() bool`

HasStartedInLast5Years returns a boolean if a field has been set.

### SetStartedInLast5YearsNil

`func (o *RwdStudyDataSourcesRow) SetStartedInLast5YearsNil(b bool)`

 SetStartedInLast5YearsNil sets the value for StartedInLast5Years to be an explicit nil

### UnsetStartedInLast5Years
`func (o *RwdStudyDataSourcesRow) UnsetStartedInLast5Years()`

UnsetStartedInLast5Years ensures that no value is present for StartedInLast5Years, not even an explicit nil
### GetConditionCount

`func (o *RwdStudyDataSourcesRow) GetConditionCount() int32`

GetConditionCount returns the ConditionCount field if non-nil, zero value otherwise.

### GetConditionCountOk

`func (o *RwdStudyDataSourcesRow) GetConditionCountOk() (*int32, bool)`

GetConditionCountOk returns a tuple with the ConditionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionCount

`func (o *RwdStudyDataSourcesRow) SetConditionCount(v int32)`

SetConditionCount sets ConditionCount field to given value.

### HasConditionCount

`func (o *RwdStudyDataSourcesRow) HasConditionCount() bool`

HasConditionCount returns a boolean if a field has been set.

### SetConditionCountNil

`func (o *RwdStudyDataSourcesRow) SetConditionCountNil(b bool)`

 SetConditionCountNil sets the value for ConditionCount to be an explicit nil

### UnsetConditionCount
`func (o *RwdStudyDataSourcesRow) UnsetConditionCount()`

UnsetConditionCount ensures that no value is present for ConditionCount, not even an explicit nil
### GetConductingInstitutionCount

`func (o *RwdStudyDataSourcesRow) GetConductingInstitutionCount() int32`

GetConductingInstitutionCount returns the ConductingInstitutionCount field if non-nil, zero value otherwise.

### GetConductingInstitutionCountOk

`func (o *RwdStudyDataSourcesRow) GetConductingInstitutionCountOk() (*int32, bool)`

GetConductingInstitutionCountOk returns a tuple with the ConductingInstitutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConductingInstitutionCount

`func (o *RwdStudyDataSourcesRow) SetConductingInstitutionCount(v int32)`

SetConductingInstitutionCount sets ConductingInstitutionCount field to given value.

### HasConductingInstitutionCount

`func (o *RwdStudyDataSourcesRow) HasConductingInstitutionCount() bool`

HasConductingInstitutionCount returns a boolean if a field has been set.

### SetConductingInstitutionCountNil

`func (o *RwdStudyDataSourcesRow) SetConductingInstitutionCountNil(b bool)`

 SetConductingInstitutionCountNil sets the value for ConductingInstitutionCount to be an explicit nil

### UnsetConductingInstitutionCount
`func (o *RwdStudyDataSourcesRow) UnsetConductingInstitutionCount()`

UnsetConductingInstitutionCount ensures that no value is present for ConductingInstitutionCount, not even an explicit nil
### GetSnapshotLabel

`func (o *RwdStudyDataSourcesRow) GetSnapshotLabel() string`

GetSnapshotLabel returns the SnapshotLabel field if non-nil, zero value otherwise.

### GetSnapshotLabelOk

`func (o *RwdStudyDataSourcesRow) GetSnapshotLabelOk() (*string, bool)`

GetSnapshotLabelOk returns a tuple with the SnapshotLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotLabel

`func (o *RwdStudyDataSourcesRow) SetSnapshotLabel(v string)`

SetSnapshotLabel sets SnapshotLabel field to given value.

### HasSnapshotLabel

`func (o *RwdStudyDataSourcesRow) HasSnapshotLabel() bool`

HasSnapshotLabel returns a boolean if a field has been set.

### SetSnapshotLabelNil

`func (o *RwdStudyDataSourcesRow) SetSnapshotLabelNil(b bool)`

 SetSnapshotLabelNil sets the value for SnapshotLabel to be an explicit nil

### UnsetSnapshotLabel
`func (o *RwdStudyDataSourcesRow) UnsetSnapshotLabel()`

UnsetSnapshotLabel ensures that no value is present for SnapshotLabel, not even an explicit nil
### GetAccessedMonth

`func (o *RwdStudyDataSourcesRow) GetAccessedMonth() string`

GetAccessedMonth returns the AccessedMonth field if non-nil, zero value otherwise.

### GetAccessedMonthOk

`func (o *RwdStudyDataSourcesRow) GetAccessedMonthOk() (*string, bool)`

GetAccessedMonthOk returns a tuple with the AccessedMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessedMonth

`func (o *RwdStudyDataSourcesRow) SetAccessedMonth(v string)`

SetAccessedMonth sets AccessedMonth field to given value.

### HasAccessedMonth

`func (o *RwdStudyDataSourcesRow) HasAccessedMonth() bool`

HasAccessedMonth returns a boolean if a field has been set.

### SetAccessedMonthNil

`func (o *RwdStudyDataSourcesRow) SetAccessedMonthNil(b bool)`

 SetAccessedMonthNil sets the value for AccessedMonth to be an explicit nil

### UnsetAccessedMonth
`func (o *RwdStudyDataSourcesRow) UnsetAccessedMonth()`

UnsetAccessedMonth ensures that no value is present for AccessedMonth, not even an explicit nil
### GetRestrictionTier

`func (o *RwdStudyDataSourcesRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *RwdStudyDataSourcesRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *RwdStudyDataSourcesRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *RwdStudyDataSourcesRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *RwdStudyDataSourcesRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *RwdStudyDataSourcesRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetAttribution

`func (o *RwdStudyDataSourcesRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RwdStudyDataSourcesRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RwdStudyDataSourcesRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RwdStudyDataSourcesRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RwdStudyDataSourcesRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RwdStudyDataSourcesRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


