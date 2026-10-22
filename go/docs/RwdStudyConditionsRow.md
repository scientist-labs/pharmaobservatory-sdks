# RwdStudyConditionsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StudyId** | Pointer to **NullableString** | The register&#39;s own Study ID. The entry key of the rwd_studies object type; the EU PAS number is an attribute. | [optional] 
**ConditionRole** | Pointer to **NullableString** | primary when the term is the study&#39;s medicinal condition to be studied, additional when it is in the additional conditions list. | [optional] 
**ConditionTerm** | Pointer to **NullableString** | The condition term as the register has it. | [optional] 
**MentionCount** | Pointer to **NullableInt32** | How many times the term is mentioned in the study. | [optional] 
**ConditionId** | Pointer to **NullableString** | The condition&#39;s MeSH descriptor id (D followed by six to nine digits), the id of the conditions object type. | [optional] 
**MeshTerm** | Pointer to **NullableString** | The MeSH heading the term resolved to; null when it did not resolve. | [optional] 
**MeshTreeCategory** | Pointer to **NullableString** | The MeSH tree category the condition sits under. | [optional] 
**IsDiseaseDescriptor** | Pointer to **NullableBool** | Whether the MeSH descriptor is a disease descriptor. | [optional] 
**MeshMatchRung** | Pointer to **NullableString** | How the term resolved to MeSH (curated, entry_term and so on, or unresolved). | [optional] 
**IsMeshResolved** | Pointer to **NullableBool** | Whether the term resolved to a MeSH descriptor. | [optional] 
**StudyStatus** | Pointer to **NullableString** | The study&#39;s status as the register words it (Planned, Ongoing, Finalised, and so on). | [optional] 
**MandateStratum** | Pointer to **NullableString** | mandated when a regulator required the study (a post-authorisation obligation), voluntary when the registrant says it did not, unknown when unanswered. Delivery timeliness is only comparable within a stratum. | [optional] 
**RequiredByRegulator** | Pointer to **NullableString** | The registrant&#39;s answer to whether a regulatory body required the study (Yes, No or Unknown). | [optional] 
**DataSourceCount** | Pointer to **NullableInt32** | How many data source items the study names. | [optional] 
**ConductingInstitutionCount** | Pointer to **NullableInt32** | How many conducting institutions the study names. | [optional] 
**SnapshotLabel** | Pointer to **NullableString** | The label of the catalogue snapshot the row was read from (a date). | [optional] 
**AccessedMonth** | Pointer to **NullableString** | The month and year the catalogue was accessed (YYYY-MM), which EMA&#39;s site asks to be cited with its URL. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRwdStudyConditionsRow

`func NewRwdStudyConditionsRow() *RwdStudyConditionsRow`

NewRwdStudyConditionsRow instantiates a new RwdStudyConditionsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRwdStudyConditionsRowWithDefaults

`func NewRwdStudyConditionsRowWithDefaults() *RwdStudyConditionsRow`

NewRwdStudyConditionsRowWithDefaults instantiates a new RwdStudyConditionsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStudyId

`func (o *RwdStudyConditionsRow) GetStudyId() string`

GetStudyId returns the StudyId field if non-nil, zero value otherwise.

### GetStudyIdOk

`func (o *RwdStudyConditionsRow) GetStudyIdOk() (*string, bool)`

GetStudyIdOk returns a tuple with the StudyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyId

`func (o *RwdStudyConditionsRow) SetStudyId(v string)`

SetStudyId sets StudyId field to given value.

### HasStudyId

`func (o *RwdStudyConditionsRow) HasStudyId() bool`

HasStudyId returns a boolean if a field has been set.

### SetStudyIdNil

`func (o *RwdStudyConditionsRow) SetStudyIdNil(b bool)`

 SetStudyIdNil sets the value for StudyId to be an explicit nil

### UnsetStudyId
`func (o *RwdStudyConditionsRow) UnsetStudyId()`

UnsetStudyId ensures that no value is present for StudyId, not even an explicit nil
### GetConditionRole

`func (o *RwdStudyConditionsRow) GetConditionRole() string`

GetConditionRole returns the ConditionRole field if non-nil, zero value otherwise.

### GetConditionRoleOk

`func (o *RwdStudyConditionsRow) GetConditionRoleOk() (*string, bool)`

GetConditionRoleOk returns a tuple with the ConditionRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionRole

`func (o *RwdStudyConditionsRow) SetConditionRole(v string)`

SetConditionRole sets ConditionRole field to given value.

### HasConditionRole

`func (o *RwdStudyConditionsRow) HasConditionRole() bool`

HasConditionRole returns a boolean if a field has been set.

### SetConditionRoleNil

`func (o *RwdStudyConditionsRow) SetConditionRoleNil(b bool)`

 SetConditionRoleNil sets the value for ConditionRole to be an explicit nil

### UnsetConditionRole
`func (o *RwdStudyConditionsRow) UnsetConditionRole()`

UnsetConditionRole ensures that no value is present for ConditionRole, not even an explicit nil
### GetConditionTerm

`func (o *RwdStudyConditionsRow) GetConditionTerm() string`

GetConditionTerm returns the ConditionTerm field if non-nil, zero value otherwise.

### GetConditionTermOk

`func (o *RwdStudyConditionsRow) GetConditionTermOk() (*string, bool)`

GetConditionTermOk returns a tuple with the ConditionTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionTerm

`func (o *RwdStudyConditionsRow) SetConditionTerm(v string)`

SetConditionTerm sets ConditionTerm field to given value.

### HasConditionTerm

`func (o *RwdStudyConditionsRow) HasConditionTerm() bool`

HasConditionTerm returns a boolean if a field has been set.

### SetConditionTermNil

`func (o *RwdStudyConditionsRow) SetConditionTermNil(b bool)`

 SetConditionTermNil sets the value for ConditionTerm to be an explicit nil

### UnsetConditionTerm
`func (o *RwdStudyConditionsRow) UnsetConditionTerm()`

UnsetConditionTerm ensures that no value is present for ConditionTerm, not even an explicit nil
### GetMentionCount

`func (o *RwdStudyConditionsRow) GetMentionCount() int32`

GetMentionCount returns the MentionCount field if non-nil, zero value otherwise.

### GetMentionCountOk

`func (o *RwdStudyConditionsRow) GetMentionCountOk() (*int32, bool)`

GetMentionCountOk returns a tuple with the MentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionCount

`func (o *RwdStudyConditionsRow) SetMentionCount(v int32)`

SetMentionCount sets MentionCount field to given value.

### HasMentionCount

`func (o *RwdStudyConditionsRow) HasMentionCount() bool`

HasMentionCount returns a boolean if a field has been set.

### SetMentionCountNil

`func (o *RwdStudyConditionsRow) SetMentionCountNil(b bool)`

 SetMentionCountNil sets the value for MentionCount to be an explicit nil

### UnsetMentionCount
`func (o *RwdStudyConditionsRow) UnsetMentionCount()`

UnsetMentionCount ensures that no value is present for MentionCount, not even an explicit nil
### GetConditionId

`func (o *RwdStudyConditionsRow) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *RwdStudyConditionsRow) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *RwdStudyConditionsRow) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *RwdStudyConditionsRow) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### SetConditionIdNil

`func (o *RwdStudyConditionsRow) SetConditionIdNil(b bool)`

 SetConditionIdNil sets the value for ConditionId to be an explicit nil

### UnsetConditionId
`func (o *RwdStudyConditionsRow) UnsetConditionId()`

UnsetConditionId ensures that no value is present for ConditionId, not even an explicit nil
### GetMeshTerm

`func (o *RwdStudyConditionsRow) GetMeshTerm() string`

GetMeshTerm returns the MeshTerm field if non-nil, zero value otherwise.

### GetMeshTermOk

`func (o *RwdStudyConditionsRow) GetMeshTermOk() (*string, bool)`

GetMeshTermOk returns a tuple with the MeshTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshTerm

`func (o *RwdStudyConditionsRow) SetMeshTerm(v string)`

SetMeshTerm sets MeshTerm field to given value.

### HasMeshTerm

`func (o *RwdStudyConditionsRow) HasMeshTerm() bool`

HasMeshTerm returns a boolean if a field has been set.

### SetMeshTermNil

`func (o *RwdStudyConditionsRow) SetMeshTermNil(b bool)`

 SetMeshTermNil sets the value for MeshTerm to be an explicit nil

### UnsetMeshTerm
`func (o *RwdStudyConditionsRow) UnsetMeshTerm()`

UnsetMeshTerm ensures that no value is present for MeshTerm, not even an explicit nil
### GetMeshTreeCategory

`func (o *RwdStudyConditionsRow) GetMeshTreeCategory() string`

GetMeshTreeCategory returns the MeshTreeCategory field if non-nil, zero value otherwise.

### GetMeshTreeCategoryOk

`func (o *RwdStudyConditionsRow) GetMeshTreeCategoryOk() (*string, bool)`

GetMeshTreeCategoryOk returns a tuple with the MeshTreeCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshTreeCategory

`func (o *RwdStudyConditionsRow) SetMeshTreeCategory(v string)`

SetMeshTreeCategory sets MeshTreeCategory field to given value.

### HasMeshTreeCategory

`func (o *RwdStudyConditionsRow) HasMeshTreeCategory() bool`

HasMeshTreeCategory returns a boolean if a field has been set.

### SetMeshTreeCategoryNil

`func (o *RwdStudyConditionsRow) SetMeshTreeCategoryNil(b bool)`

 SetMeshTreeCategoryNil sets the value for MeshTreeCategory to be an explicit nil

### UnsetMeshTreeCategory
`func (o *RwdStudyConditionsRow) UnsetMeshTreeCategory()`

UnsetMeshTreeCategory ensures that no value is present for MeshTreeCategory, not even an explicit nil
### GetIsDiseaseDescriptor

`func (o *RwdStudyConditionsRow) GetIsDiseaseDescriptor() bool`

GetIsDiseaseDescriptor returns the IsDiseaseDescriptor field if non-nil, zero value otherwise.

### GetIsDiseaseDescriptorOk

`func (o *RwdStudyConditionsRow) GetIsDiseaseDescriptorOk() (*bool, bool)`

GetIsDiseaseDescriptorOk returns a tuple with the IsDiseaseDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDiseaseDescriptor

`func (o *RwdStudyConditionsRow) SetIsDiseaseDescriptor(v bool)`

SetIsDiseaseDescriptor sets IsDiseaseDescriptor field to given value.

### HasIsDiseaseDescriptor

`func (o *RwdStudyConditionsRow) HasIsDiseaseDescriptor() bool`

HasIsDiseaseDescriptor returns a boolean if a field has been set.

### SetIsDiseaseDescriptorNil

`func (o *RwdStudyConditionsRow) SetIsDiseaseDescriptorNil(b bool)`

 SetIsDiseaseDescriptorNil sets the value for IsDiseaseDescriptor to be an explicit nil

### UnsetIsDiseaseDescriptor
`func (o *RwdStudyConditionsRow) UnsetIsDiseaseDescriptor()`

UnsetIsDiseaseDescriptor ensures that no value is present for IsDiseaseDescriptor, not even an explicit nil
### GetMeshMatchRung

`func (o *RwdStudyConditionsRow) GetMeshMatchRung() string`

GetMeshMatchRung returns the MeshMatchRung field if non-nil, zero value otherwise.

### GetMeshMatchRungOk

`func (o *RwdStudyConditionsRow) GetMeshMatchRungOk() (*string, bool)`

GetMeshMatchRungOk returns a tuple with the MeshMatchRung field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshMatchRung

`func (o *RwdStudyConditionsRow) SetMeshMatchRung(v string)`

SetMeshMatchRung sets MeshMatchRung field to given value.

### HasMeshMatchRung

`func (o *RwdStudyConditionsRow) HasMeshMatchRung() bool`

HasMeshMatchRung returns a boolean if a field has been set.

### SetMeshMatchRungNil

`func (o *RwdStudyConditionsRow) SetMeshMatchRungNil(b bool)`

 SetMeshMatchRungNil sets the value for MeshMatchRung to be an explicit nil

### UnsetMeshMatchRung
`func (o *RwdStudyConditionsRow) UnsetMeshMatchRung()`

UnsetMeshMatchRung ensures that no value is present for MeshMatchRung, not even an explicit nil
### GetIsMeshResolved

`func (o *RwdStudyConditionsRow) GetIsMeshResolved() bool`

GetIsMeshResolved returns the IsMeshResolved field if non-nil, zero value otherwise.

### GetIsMeshResolvedOk

`func (o *RwdStudyConditionsRow) GetIsMeshResolvedOk() (*bool, bool)`

GetIsMeshResolvedOk returns a tuple with the IsMeshResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMeshResolved

`func (o *RwdStudyConditionsRow) SetIsMeshResolved(v bool)`

SetIsMeshResolved sets IsMeshResolved field to given value.

### HasIsMeshResolved

`func (o *RwdStudyConditionsRow) HasIsMeshResolved() bool`

HasIsMeshResolved returns a boolean if a field has been set.

### SetIsMeshResolvedNil

`func (o *RwdStudyConditionsRow) SetIsMeshResolvedNil(b bool)`

 SetIsMeshResolvedNil sets the value for IsMeshResolved to be an explicit nil

### UnsetIsMeshResolved
`func (o *RwdStudyConditionsRow) UnsetIsMeshResolved()`

UnsetIsMeshResolved ensures that no value is present for IsMeshResolved, not even an explicit nil
### GetStudyStatus

`func (o *RwdStudyConditionsRow) GetStudyStatus() string`

GetStudyStatus returns the StudyStatus field if non-nil, zero value otherwise.

### GetStudyStatusOk

`func (o *RwdStudyConditionsRow) GetStudyStatusOk() (*string, bool)`

GetStudyStatusOk returns a tuple with the StudyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyStatus

`func (o *RwdStudyConditionsRow) SetStudyStatus(v string)`

SetStudyStatus sets StudyStatus field to given value.

### HasStudyStatus

`func (o *RwdStudyConditionsRow) HasStudyStatus() bool`

HasStudyStatus returns a boolean if a field has been set.

### SetStudyStatusNil

`func (o *RwdStudyConditionsRow) SetStudyStatusNil(b bool)`

 SetStudyStatusNil sets the value for StudyStatus to be an explicit nil

### UnsetStudyStatus
`func (o *RwdStudyConditionsRow) UnsetStudyStatus()`

UnsetStudyStatus ensures that no value is present for StudyStatus, not even an explicit nil
### GetMandateStratum

`func (o *RwdStudyConditionsRow) GetMandateStratum() string`

GetMandateStratum returns the MandateStratum field if non-nil, zero value otherwise.

### GetMandateStratumOk

`func (o *RwdStudyConditionsRow) GetMandateStratumOk() (*string, bool)`

GetMandateStratumOk returns a tuple with the MandateStratum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateStratum

`func (o *RwdStudyConditionsRow) SetMandateStratum(v string)`

SetMandateStratum sets MandateStratum field to given value.

### HasMandateStratum

`func (o *RwdStudyConditionsRow) HasMandateStratum() bool`

HasMandateStratum returns a boolean if a field has been set.

### SetMandateStratumNil

`func (o *RwdStudyConditionsRow) SetMandateStratumNil(b bool)`

 SetMandateStratumNil sets the value for MandateStratum to be an explicit nil

### UnsetMandateStratum
`func (o *RwdStudyConditionsRow) UnsetMandateStratum()`

UnsetMandateStratum ensures that no value is present for MandateStratum, not even an explicit nil
### GetRequiredByRegulator

`func (o *RwdStudyConditionsRow) GetRequiredByRegulator() string`

GetRequiredByRegulator returns the RequiredByRegulator field if non-nil, zero value otherwise.

### GetRequiredByRegulatorOk

`func (o *RwdStudyConditionsRow) GetRequiredByRegulatorOk() (*string, bool)`

GetRequiredByRegulatorOk returns a tuple with the RequiredByRegulator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredByRegulator

`func (o *RwdStudyConditionsRow) SetRequiredByRegulator(v string)`

SetRequiredByRegulator sets RequiredByRegulator field to given value.

### HasRequiredByRegulator

`func (o *RwdStudyConditionsRow) HasRequiredByRegulator() bool`

HasRequiredByRegulator returns a boolean if a field has been set.

### SetRequiredByRegulatorNil

`func (o *RwdStudyConditionsRow) SetRequiredByRegulatorNil(b bool)`

 SetRequiredByRegulatorNil sets the value for RequiredByRegulator to be an explicit nil

### UnsetRequiredByRegulator
`func (o *RwdStudyConditionsRow) UnsetRequiredByRegulator()`

UnsetRequiredByRegulator ensures that no value is present for RequiredByRegulator, not even an explicit nil
### GetDataSourceCount

`func (o *RwdStudyConditionsRow) GetDataSourceCount() int32`

GetDataSourceCount returns the DataSourceCount field if non-nil, zero value otherwise.

### GetDataSourceCountOk

`func (o *RwdStudyConditionsRow) GetDataSourceCountOk() (*int32, bool)`

GetDataSourceCountOk returns a tuple with the DataSourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceCount

`func (o *RwdStudyConditionsRow) SetDataSourceCount(v int32)`

SetDataSourceCount sets DataSourceCount field to given value.

### HasDataSourceCount

`func (o *RwdStudyConditionsRow) HasDataSourceCount() bool`

HasDataSourceCount returns a boolean if a field has been set.

### SetDataSourceCountNil

`func (o *RwdStudyConditionsRow) SetDataSourceCountNil(b bool)`

 SetDataSourceCountNil sets the value for DataSourceCount to be an explicit nil

### UnsetDataSourceCount
`func (o *RwdStudyConditionsRow) UnsetDataSourceCount()`

UnsetDataSourceCount ensures that no value is present for DataSourceCount, not even an explicit nil
### GetConductingInstitutionCount

`func (o *RwdStudyConditionsRow) GetConductingInstitutionCount() int32`

GetConductingInstitutionCount returns the ConductingInstitutionCount field if non-nil, zero value otherwise.

### GetConductingInstitutionCountOk

`func (o *RwdStudyConditionsRow) GetConductingInstitutionCountOk() (*int32, bool)`

GetConductingInstitutionCountOk returns a tuple with the ConductingInstitutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConductingInstitutionCount

`func (o *RwdStudyConditionsRow) SetConductingInstitutionCount(v int32)`

SetConductingInstitutionCount sets ConductingInstitutionCount field to given value.

### HasConductingInstitutionCount

`func (o *RwdStudyConditionsRow) HasConductingInstitutionCount() bool`

HasConductingInstitutionCount returns a boolean if a field has been set.

### SetConductingInstitutionCountNil

`func (o *RwdStudyConditionsRow) SetConductingInstitutionCountNil(b bool)`

 SetConductingInstitutionCountNil sets the value for ConductingInstitutionCount to be an explicit nil

### UnsetConductingInstitutionCount
`func (o *RwdStudyConditionsRow) UnsetConductingInstitutionCount()`

UnsetConductingInstitutionCount ensures that no value is present for ConductingInstitutionCount, not even an explicit nil
### GetSnapshotLabel

`func (o *RwdStudyConditionsRow) GetSnapshotLabel() string`

GetSnapshotLabel returns the SnapshotLabel field if non-nil, zero value otherwise.

### GetSnapshotLabelOk

`func (o *RwdStudyConditionsRow) GetSnapshotLabelOk() (*string, bool)`

GetSnapshotLabelOk returns a tuple with the SnapshotLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotLabel

`func (o *RwdStudyConditionsRow) SetSnapshotLabel(v string)`

SetSnapshotLabel sets SnapshotLabel field to given value.

### HasSnapshotLabel

`func (o *RwdStudyConditionsRow) HasSnapshotLabel() bool`

HasSnapshotLabel returns a boolean if a field has been set.

### SetSnapshotLabelNil

`func (o *RwdStudyConditionsRow) SetSnapshotLabelNil(b bool)`

 SetSnapshotLabelNil sets the value for SnapshotLabel to be an explicit nil

### UnsetSnapshotLabel
`func (o *RwdStudyConditionsRow) UnsetSnapshotLabel()`

UnsetSnapshotLabel ensures that no value is present for SnapshotLabel, not even an explicit nil
### GetAccessedMonth

`func (o *RwdStudyConditionsRow) GetAccessedMonth() string`

GetAccessedMonth returns the AccessedMonth field if non-nil, zero value otherwise.

### GetAccessedMonthOk

`func (o *RwdStudyConditionsRow) GetAccessedMonthOk() (*string, bool)`

GetAccessedMonthOk returns a tuple with the AccessedMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessedMonth

`func (o *RwdStudyConditionsRow) SetAccessedMonth(v string)`

SetAccessedMonth sets AccessedMonth field to given value.

### HasAccessedMonth

`func (o *RwdStudyConditionsRow) HasAccessedMonth() bool`

HasAccessedMonth returns a boolean if a field has been set.

### SetAccessedMonthNil

`func (o *RwdStudyConditionsRow) SetAccessedMonthNil(b bool)`

 SetAccessedMonthNil sets the value for AccessedMonth to be an explicit nil

### UnsetAccessedMonth
`func (o *RwdStudyConditionsRow) UnsetAccessedMonth()`

UnsetAccessedMonth ensures that no value is present for AccessedMonth, not even an explicit nil
### GetRestrictionTier

`func (o *RwdStudyConditionsRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *RwdStudyConditionsRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *RwdStudyConditionsRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *RwdStudyConditionsRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *RwdStudyConditionsRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *RwdStudyConditionsRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetAttribution

`func (o *RwdStudyConditionsRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RwdStudyConditionsRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RwdStudyConditionsRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RwdStudyConditionsRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RwdStudyConditionsRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RwdStudyConditionsRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


