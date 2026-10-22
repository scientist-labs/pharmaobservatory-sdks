# RwdStudyDrugsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StudyId** | Pointer to **NullableString** | The register&#39;s own Study ID. The entry key of the rwd_studies object type; the EU PAS number is an attribute. | [optional] 
**MentionSource** | Pointer to **NullableString** | Which study field the drug mention came from: inn or atc_code. | [optional] 
**ItemOrdinal** | Pointer to **NullableInt32** | The position of the item in the register&#39;s own list, from 1. | [optional] 
**DrugMention** | Pointer to **NullableString** | The drug mention as typed. | [optional] 
**SourceList** | Pointer to **NullableString** | The whole list the mention was split from. | [optional] 
**AtcCode** | Pointer to **NullableString** | The ATC code of the mention. | [optional] 
**AtcLevel5Code** | Pointer to **NullableString** | The full ATC level 5 code (seven characters, a chemical substance), when the mention carries one. | [optional] 
**AtcLevel1Code** | Pointer to **NullableString** | The ATC level 1 code (one letter, the anatomical main group). | [optional] 
**AtcLabel** | Pointer to **NullableString** | The substance name the mention carries beside the code. | [optional] 
**InnName** | Pointer to **NullableString** | The INN of the mention as typed. | [optional] 
**InnNameKey** | Pointer to **NullableString** | The normalised INN key. | [optional] 
**MatchRung** | Pointer to **NullableString** | How the mention resolved to a canonical drug: atc_level5 or inn_name when it reached exactly one, ambiguous when several tied, unresolved when none, not_applicable when the mention has nothing to match. | [optional] 
**DrugId** | Pointer to **NullableString** | The canonical drug id (for example unii:53AXN4NNHX) when the mention reached exactly one drug; null when unresolved or ambiguous (candidate_drug_ids holds the tie). | [optional] 
**CandidateDrugIds** | Pointer to **[]string** | The candidate canonical drug ids when the mention is ambiguous. | [optional] 
**CandidateDrugCount** | Pointer to **NullableInt32** | How many candidates. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The license restriction tier of the row&#39;s source. Every row of this dataset is licensable (PROVISIONAL, legal read pending). | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line EMA requires, including the catalogue URL and the access month; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRwdStudyDrugsRow

`func NewRwdStudyDrugsRow() *RwdStudyDrugsRow`

NewRwdStudyDrugsRow instantiates a new RwdStudyDrugsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRwdStudyDrugsRowWithDefaults

`func NewRwdStudyDrugsRowWithDefaults() *RwdStudyDrugsRow`

NewRwdStudyDrugsRowWithDefaults instantiates a new RwdStudyDrugsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStudyId

`func (o *RwdStudyDrugsRow) GetStudyId() string`

GetStudyId returns the StudyId field if non-nil, zero value otherwise.

### GetStudyIdOk

`func (o *RwdStudyDrugsRow) GetStudyIdOk() (*string, bool)`

GetStudyIdOk returns a tuple with the StudyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudyId

`func (o *RwdStudyDrugsRow) SetStudyId(v string)`

SetStudyId sets StudyId field to given value.

### HasStudyId

`func (o *RwdStudyDrugsRow) HasStudyId() bool`

HasStudyId returns a boolean if a field has been set.

### SetStudyIdNil

`func (o *RwdStudyDrugsRow) SetStudyIdNil(b bool)`

 SetStudyIdNil sets the value for StudyId to be an explicit nil

### UnsetStudyId
`func (o *RwdStudyDrugsRow) UnsetStudyId()`

UnsetStudyId ensures that no value is present for StudyId, not even an explicit nil
### GetMentionSource

`func (o *RwdStudyDrugsRow) GetMentionSource() string`

GetMentionSource returns the MentionSource field if non-nil, zero value otherwise.

### GetMentionSourceOk

`func (o *RwdStudyDrugsRow) GetMentionSourceOk() (*string, bool)`

GetMentionSourceOk returns a tuple with the MentionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionSource

`func (o *RwdStudyDrugsRow) SetMentionSource(v string)`

SetMentionSource sets MentionSource field to given value.

### HasMentionSource

`func (o *RwdStudyDrugsRow) HasMentionSource() bool`

HasMentionSource returns a boolean if a field has been set.

### SetMentionSourceNil

`func (o *RwdStudyDrugsRow) SetMentionSourceNil(b bool)`

 SetMentionSourceNil sets the value for MentionSource to be an explicit nil

### UnsetMentionSource
`func (o *RwdStudyDrugsRow) UnsetMentionSource()`

UnsetMentionSource ensures that no value is present for MentionSource, not even an explicit nil
### GetItemOrdinal

`func (o *RwdStudyDrugsRow) GetItemOrdinal() int32`

GetItemOrdinal returns the ItemOrdinal field if non-nil, zero value otherwise.

### GetItemOrdinalOk

`func (o *RwdStudyDrugsRow) GetItemOrdinalOk() (*int32, bool)`

GetItemOrdinalOk returns a tuple with the ItemOrdinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemOrdinal

`func (o *RwdStudyDrugsRow) SetItemOrdinal(v int32)`

SetItemOrdinal sets ItemOrdinal field to given value.

### HasItemOrdinal

`func (o *RwdStudyDrugsRow) HasItemOrdinal() bool`

HasItemOrdinal returns a boolean if a field has been set.

### SetItemOrdinalNil

`func (o *RwdStudyDrugsRow) SetItemOrdinalNil(b bool)`

 SetItemOrdinalNil sets the value for ItemOrdinal to be an explicit nil

### UnsetItemOrdinal
`func (o *RwdStudyDrugsRow) UnsetItemOrdinal()`

UnsetItemOrdinal ensures that no value is present for ItemOrdinal, not even an explicit nil
### GetDrugMention

`func (o *RwdStudyDrugsRow) GetDrugMention() string`

GetDrugMention returns the DrugMention field if non-nil, zero value otherwise.

### GetDrugMentionOk

`func (o *RwdStudyDrugsRow) GetDrugMentionOk() (*string, bool)`

GetDrugMentionOk returns a tuple with the DrugMention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugMention

`func (o *RwdStudyDrugsRow) SetDrugMention(v string)`

SetDrugMention sets DrugMention field to given value.

### HasDrugMention

`func (o *RwdStudyDrugsRow) HasDrugMention() bool`

HasDrugMention returns a boolean if a field has been set.

### SetDrugMentionNil

`func (o *RwdStudyDrugsRow) SetDrugMentionNil(b bool)`

 SetDrugMentionNil sets the value for DrugMention to be an explicit nil

### UnsetDrugMention
`func (o *RwdStudyDrugsRow) UnsetDrugMention()`

UnsetDrugMention ensures that no value is present for DrugMention, not even an explicit nil
### GetSourceList

`func (o *RwdStudyDrugsRow) GetSourceList() string`

GetSourceList returns the SourceList field if non-nil, zero value otherwise.

### GetSourceListOk

`func (o *RwdStudyDrugsRow) GetSourceListOk() (*string, bool)`

GetSourceListOk returns a tuple with the SourceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceList

`func (o *RwdStudyDrugsRow) SetSourceList(v string)`

SetSourceList sets SourceList field to given value.

### HasSourceList

`func (o *RwdStudyDrugsRow) HasSourceList() bool`

HasSourceList returns a boolean if a field has been set.

### SetSourceListNil

`func (o *RwdStudyDrugsRow) SetSourceListNil(b bool)`

 SetSourceListNil sets the value for SourceList to be an explicit nil

### UnsetSourceList
`func (o *RwdStudyDrugsRow) UnsetSourceList()`

UnsetSourceList ensures that no value is present for SourceList, not even an explicit nil
### GetAtcCode

`func (o *RwdStudyDrugsRow) GetAtcCode() string`

GetAtcCode returns the AtcCode field if non-nil, zero value otherwise.

### GetAtcCodeOk

`func (o *RwdStudyDrugsRow) GetAtcCodeOk() (*string, bool)`

GetAtcCodeOk returns a tuple with the AtcCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtcCode

`func (o *RwdStudyDrugsRow) SetAtcCode(v string)`

SetAtcCode sets AtcCode field to given value.

### HasAtcCode

`func (o *RwdStudyDrugsRow) HasAtcCode() bool`

HasAtcCode returns a boolean if a field has been set.

### SetAtcCodeNil

`func (o *RwdStudyDrugsRow) SetAtcCodeNil(b bool)`

 SetAtcCodeNil sets the value for AtcCode to be an explicit nil

### UnsetAtcCode
`func (o *RwdStudyDrugsRow) UnsetAtcCode()`

UnsetAtcCode ensures that no value is present for AtcCode, not even an explicit nil
### GetAtcLevel5Code

`func (o *RwdStudyDrugsRow) GetAtcLevel5Code() string`

GetAtcLevel5Code returns the AtcLevel5Code field if non-nil, zero value otherwise.

### GetAtcLevel5CodeOk

`func (o *RwdStudyDrugsRow) GetAtcLevel5CodeOk() (*string, bool)`

GetAtcLevel5CodeOk returns a tuple with the AtcLevel5Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtcLevel5Code

`func (o *RwdStudyDrugsRow) SetAtcLevel5Code(v string)`

SetAtcLevel5Code sets AtcLevel5Code field to given value.

### HasAtcLevel5Code

`func (o *RwdStudyDrugsRow) HasAtcLevel5Code() bool`

HasAtcLevel5Code returns a boolean if a field has been set.

### SetAtcLevel5CodeNil

`func (o *RwdStudyDrugsRow) SetAtcLevel5CodeNil(b bool)`

 SetAtcLevel5CodeNil sets the value for AtcLevel5Code to be an explicit nil

### UnsetAtcLevel5Code
`func (o *RwdStudyDrugsRow) UnsetAtcLevel5Code()`

UnsetAtcLevel5Code ensures that no value is present for AtcLevel5Code, not even an explicit nil
### GetAtcLevel1Code

`func (o *RwdStudyDrugsRow) GetAtcLevel1Code() string`

GetAtcLevel1Code returns the AtcLevel1Code field if non-nil, zero value otherwise.

### GetAtcLevel1CodeOk

`func (o *RwdStudyDrugsRow) GetAtcLevel1CodeOk() (*string, bool)`

GetAtcLevel1CodeOk returns a tuple with the AtcLevel1Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtcLevel1Code

`func (o *RwdStudyDrugsRow) SetAtcLevel1Code(v string)`

SetAtcLevel1Code sets AtcLevel1Code field to given value.

### HasAtcLevel1Code

`func (o *RwdStudyDrugsRow) HasAtcLevel1Code() bool`

HasAtcLevel1Code returns a boolean if a field has been set.

### SetAtcLevel1CodeNil

`func (o *RwdStudyDrugsRow) SetAtcLevel1CodeNil(b bool)`

 SetAtcLevel1CodeNil sets the value for AtcLevel1Code to be an explicit nil

### UnsetAtcLevel1Code
`func (o *RwdStudyDrugsRow) UnsetAtcLevel1Code()`

UnsetAtcLevel1Code ensures that no value is present for AtcLevel1Code, not even an explicit nil
### GetAtcLabel

`func (o *RwdStudyDrugsRow) GetAtcLabel() string`

GetAtcLabel returns the AtcLabel field if non-nil, zero value otherwise.

### GetAtcLabelOk

`func (o *RwdStudyDrugsRow) GetAtcLabelOk() (*string, bool)`

GetAtcLabelOk returns a tuple with the AtcLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtcLabel

`func (o *RwdStudyDrugsRow) SetAtcLabel(v string)`

SetAtcLabel sets AtcLabel field to given value.

### HasAtcLabel

`func (o *RwdStudyDrugsRow) HasAtcLabel() bool`

HasAtcLabel returns a boolean if a field has been set.

### SetAtcLabelNil

`func (o *RwdStudyDrugsRow) SetAtcLabelNil(b bool)`

 SetAtcLabelNil sets the value for AtcLabel to be an explicit nil

### UnsetAtcLabel
`func (o *RwdStudyDrugsRow) UnsetAtcLabel()`

UnsetAtcLabel ensures that no value is present for AtcLabel, not even an explicit nil
### GetInnName

`func (o *RwdStudyDrugsRow) GetInnName() string`

GetInnName returns the InnName field if non-nil, zero value otherwise.

### GetInnNameOk

`func (o *RwdStudyDrugsRow) GetInnNameOk() (*string, bool)`

GetInnNameOk returns a tuple with the InnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnName

`func (o *RwdStudyDrugsRow) SetInnName(v string)`

SetInnName sets InnName field to given value.

### HasInnName

`func (o *RwdStudyDrugsRow) HasInnName() bool`

HasInnName returns a boolean if a field has been set.

### SetInnNameNil

`func (o *RwdStudyDrugsRow) SetInnNameNil(b bool)`

 SetInnNameNil sets the value for InnName to be an explicit nil

### UnsetInnName
`func (o *RwdStudyDrugsRow) UnsetInnName()`

UnsetInnName ensures that no value is present for InnName, not even an explicit nil
### GetInnNameKey

`func (o *RwdStudyDrugsRow) GetInnNameKey() string`

GetInnNameKey returns the InnNameKey field if non-nil, zero value otherwise.

### GetInnNameKeyOk

`func (o *RwdStudyDrugsRow) GetInnNameKeyOk() (*string, bool)`

GetInnNameKeyOk returns a tuple with the InnNameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnNameKey

`func (o *RwdStudyDrugsRow) SetInnNameKey(v string)`

SetInnNameKey sets InnNameKey field to given value.

### HasInnNameKey

`func (o *RwdStudyDrugsRow) HasInnNameKey() bool`

HasInnNameKey returns a boolean if a field has been set.

### SetInnNameKeyNil

`func (o *RwdStudyDrugsRow) SetInnNameKeyNil(b bool)`

 SetInnNameKeyNil sets the value for InnNameKey to be an explicit nil

### UnsetInnNameKey
`func (o *RwdStudyDrugsRow) UnsetInnNameKey()`

UnsetInnNameKey ensures that no value is present for InnNameKey, not even an explicit nil
### GetMatchRung

`func (o *RwdStudyDrugsRow) GetMatchRung() string`

GetMatchRung returns the MatchRung field if non-nil, zero value otherwise.

### GetMatchRungOk

`func (o *RwdStudyDrugsRow) GetMatchRungOk() (*string, bool)`

GetMatchRungOk returns a tuple with the MatchRung field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchRung

`func (o *RwdStudyDrugsRow) SetMatchRung(v string)`

SetMatchRung sets MatchRung field to given value.

### HasMatchRung

`func (o *RwdStudyDrugsRow) HasMatchRung() bool`

HasMatchRung returns a boolean if a field has been set.

### SetMatchRungNil

`func (o *RwdStudyDrugsRow) SetMatchRungNil(b bool)`

 SetMatchRungNil sets the value for MatchRung to be an explicit nil

### UnsetMatchRung
`func (o *RwdStudyDrugsRow) UnsetMatchRung()`

UnsetMatchRung ensures that no value is present for MatchRung, not even an explicit nil
### GetDrugId

`func (o *RwdStudyDrugsRow) GetDrugId() string`

GetDrugId returns the DrugId field if non-nil, zero value otherwise.

### GetDrugIdOk

`func (o *RwdStudyDrugsRow) GetDrugIdOk() (*string, bool)`

GetDrugIdOk returns a tuple with the DrugId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugId

`func (o *RwdStudyDrugsRow) SetDrugId(v string)`

SetDrugId sets DrugId field to given value.

### HasDrugId

`func (o *RwdStudyDrugsRow) HasDrugId() bool`

HasDrugId returns a boolean if a field has been set.

### SetDrugIdNil

`func (o *RwdStudyDrugsRow) SetDrugIdNil(b bool)`

 SetDrugIdNil sets the value for DrugId to be an explicit nil

### UnsetDrugId
`func (o *RwdStudyDrugsRow) UnsetDrugId()`

UnsetDrugId ensures that no value is present for DrugId, not even an explicit nil
### GetCandidateDrugIds

`func (o *RwdStudyDrugsRow) GetCandidateDrugIds() []string`

GetCandidateDrugIds returns the CandidateDrugIds field if non-nil, zero value otherwise.

### GetCandidateDrugIdsOk

`func (o *RwdStudyDrugsRow) GetCandidateDrugIdsOk() (*[]string, bool)`

GetCandidateDrugIdsOk returns a tuple with the CandidateDrugIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidateDrugIds

`func (o *RwdStudyDrugsRow) SetCandidateDrugIds(v []string)`

SetCandidateDrugIds sets CandidateDrugIds field to given value.

### HasCandidateDrugIds

`func (o *RwdStudyDrugsRow) HasCandidateDrugIds() bool`

HasCandidateDrugIds returns a boolean if a field has been set.

### SetCandidateDrugIdsNil

`func (o *RwdStudyDrugsRow) SetCandidateDrugIdsNil(b bool)`

 SetCandidateDrugIdsNil sets the value for CandidateDrugIds to be an explicit nil

### UnsetCandidateDrugIds
`func (o *RwdStudyDrugsRow) UnsetCandidateDrugIds()`

UnsetCandidateDrugIds ensures that no value is present for CandidateDrugIds, not even an explicit nil
### GetCandidateDrugCount

`func (o *RwdStudyDrugsRow) GetCandidateDrugCount() int32`

GetCandidateDrugCount returns the CandidateDrugCount field if non-nil, zero value otherwise.

### GetCandidateDrugCountOk

`func (o *RwdStudyDrugsRow) GetCandidateDrugCountOk() (*int32, bool)`

GetCandidateDrugCountOk returns a tuple with the CandidateDrugCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidateDrugCount

`func (o *RwdStudyDrugsRow) SetCandidateDrugCount(v int32)`

SetCandidateDrugCount sets CandidateDrugCount field to given value.

### HasCandidateDrugCount

`func (o *RwdStudyDrugsRow) HasCandidateDrugCount() bool`

HasCandidateDrugCount returns a boolean if a field has been set.

### SetCandidateDrugCountNil

`func (o *RwdStudyDrugsRow) SetCandidateDrugCountNil(b bool)`

 SetCandidateDrugCountNil sets the value for CandidateDrugCount to be an explicit nil

### UnsetCandidateDrugCount
`func (o *RwdStudyDrugsRow) UnsetCandidateDrugCount()`

UnsetCandidateDrugCount ensures that no value is present for CandidateDrugCount, not even an explicit nil
### GetRestrictionTier

`func (o *RwdStudyDrugsRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *RwdStudyDrugsRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *RwdStudyDrugsRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *RwdStudyDrugsRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *RwdStudyDrugsRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *RwdStudyDrugsRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil
### GetAttribution

`func (o *RwdStudyDrugsRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RwdStudyDrugsRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RwdStudyDrugsRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RwdStudyDrugsRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RwdStudyDrugsRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RwdStudyDrugsRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


