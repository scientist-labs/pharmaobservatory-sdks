# OrgLiteratureRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **NullableString** | Legal-entity level: the resolved parent organization record the byline names (edgar:..., ror:... or sponsor:...), the entry key for the legal_entities object type. | [optional] 
**Pmid** | Pointer to **NullableString** | The PubMed id. | [optional] 
**Title** | Pointer to **NullableString** | The paper&#39;s title. | [optional] 
**JournalTitle** | Pointer to **NullableString** | The journal&#39;s title. | [optional] 
**PubYear** | Pointer to **NullableInt32** | The publication year (2016 or later: the dataset holds recent MeSH-indexed literature). | [optional] 
**EvidenceType** | Pointer to **NullableString** | The paper&#39;s study-design label, one normalized publication type by study-design priority. | [optional] 
**IsTrialLinked** | Pointer to **NullableBool** | True when a registered clinical trial cites the paper; false for literature no registered trial cites. | [optional] 
**OrgArm** | Pointer to **NullableString** | company when the organization is company-typed (matched on a legal-form token in the affiliation); institution for every other organization. | [optional] 
**MatchTier** | Pointer to **NullableString** | The strongest tier that attributed a byline to the organization: legal_form or both (company arm), high or medium (institution arm). | [optional] 
**NBylines** | Pointer to **NullableInt32** | The number of distinct author positions on the paper whose affiliation named the organization. | [optional] 
**CitationCount** | Pointer to **NullableInt32** | Citations where iCite (the National Institutes of Health citation-metrics service) carries them; null is not measured, never zero. | [optional] 
**RelativeCitationRatio** | Pointer to **NullableFloat32** | The iCite Relative Citation Ratio (RCR) where measured; null is not measured. | [optional] 

## Methods

### NewOrgLiteratureRow

`func NewOrgLiteratureRow() *OrgLiteratureRow`

NewOrgLiteratureRow instantiates a new OrgLiteratureRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgLiteratureRowWithDefaults

`func NewOrgLiteratureRowWithDefaults() *OrgLiteratureRow`

NewOrgLiteratureRowWithDefaults instantiates a new OrgLiteratureRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *OrgLiteratureRow) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *OrgLiteratureRow) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *OrgLiteratureRow) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *OrgLiteratureRow) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *OrgLiteratureRow) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *OrgLiteratureRow) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetPmid

`func (o *OrgLiteratureRow) GetPmid() string`

GetPmid returns the Pmid field if non-nil, zero value otherwise.

### GetPmidOk

`func (o *OrgLiteratureRow) GetPmidOk() (*string, bool)`

GetPmidOk returns a tuple with the Pmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmid

`func (o *OrgLiteratureRow) SetPmid(v string)`

SetPmid sets Pmid field to given value.

### HasPmid

`func (o *OrgLiteratureRow) HasPmid() bool`

HasPmid returns a boolean if a field has been set.

### SetPmidNil

`func (o *OrgLiteratureRow) SetPmidNil(b bool)`

 SetPmidNil sets the value for Pmid to be an explicit nil

### UnsetPmid
`func (o *OrgLiteratureRow) UnsetPmid()`

UnsetPmid ensures that no value is present for Pmid, not even an explicit nil
### GetTitle

`func (o *OrgLiteratureRow) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OrgLiteratureRow) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OrgLiteratureRow) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *OrgLiteratureRow) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *OrgLiteratureRow) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *OrgLiteratureRow) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetJournalTitle

`func (o *OrgLiteratureRow) GetJournalTitle() string`

GetJournalTitle returns the JournalTitle field if non-nil, zero value otherwise.

### GetJournalTitleOk

`func (o *OrgLiteratureRow) GetJournalTitleOk() (*string, bool)`

GetJournalTitleOk returns a tuple with the JournalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalTitle

`func (o *OrgLiteratureRow) SetJournalTitle(v string)`

SetJournalTitle sets JournalTitle field to given value.

### HasJournalTitle

`func (o *OrgLiteratureRow) HasJournalTitle() bool`

HasJournalTitle returns a boolean if a field has been set.

### SetJournalTitleNil

`func (o *OrgLiteratureRow) SetJournalTitleNil(b bool)`

 SetJournalTitleNil sets the value for JournalTitle to be an explicit nil

### UnsetJournalTitle
`func (o *OrgLiteratureRow) UnsetJournalTitle()`

UnsetJournalTitle ensures that no value is present for JournalTitle, not even an explicit nil
### GetPubYear

`func (o *OrgLiteratureRow) GetPubYear() int32`

GetPubYear returns the PubYear field if non-nil, zero value otherwise.

### GetPubYearOk

`func (o *OrgLiteratureRow) GetPubYearOk() (*int32, bool)`

GetPubYearOk returns a tuple with the PubYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubYear

`func (o *OrgLiteratureRow) SetPubYear(v int32)`

SetPubYear sets PubYear field to given value.

### HasPubYear

`func (o *OrgLiteratureRow) HasPubYear() bool`

HasPubYear returns a boolean if a field has been set.

### SetPubYearNil

`func (o *OrgLiteratureRow) SetPubYearNil(b bool)`

 SetPubYearNil sets the value for PubYear to be an explicit nil

### UnsetPubYear
`func (o *OrgLiteratureRow) UnsetPubYear()`

UnsetPubYear ensures that no value is present for PubYear, not even an explicit nil
### GetEvidenceType

`func (o *OrgLiteratureRow) GetEvidenceType() string`

GetEvidenceType returns the EvidenceType field if non-nil, zero value otherwise.

### GetEvidenceTypeOk

`func (o *OrgLiteratureRow) GetEvidenceTypeOk() (*string, bool)`

GetEvidenceTypeOk returns a tuple with the EvidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceType

`func (o *OrgLiteratureRow) SetEvidenceType(v string)`

SetEvidenceType sets EvidenceType field to given value.

### HasEvidenceType

`func (o *OrgLiteratureRow) HasEvidenceType() bool`

HasEvidenceType returns a boolean if a field has been set.

### SetEvidenceTypeNil

`func (o *OrgLiteratureRow) SetEvidenceTypeNil(b bool)`

 SetEvidenceTypeNil sets the value for EvidenceType to be an explicit nil

### UnsetEvidenceType
`func (o *OrgLiteratureRow) UnsetEvidenceType()`

UnsetEvidenceType ensures that no value is present for EvidenceType, not even an explicit nil
### GetIsTrialLinked

`func (o *OrgLiteratureRow) GetIsTrialLinked() bool`

GetIsTrialLinked returns the IsTrialLinked field if non-nil, zero value otherwise.

### GetIsTrialLinkedOk

`func (o *OrgLiteratureRow) GetIsTrialLinkedOk() (*bool, bool)`

GetIsTrialLinkedOk returns a tuple with the IsTrialLinked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrialLinked

`func (o *OrgLiteratureRow) SetIsTrialLinked(v bool)`

SetIsTrialLinked sets IsTrialLinked field to given value.

### HasIsTrialLinked

`func (o *OrgLiteratureRow) HasIsTrialLinked() bool`

HasIsTrialLinked returns a boolean if a field has been set.

### SetIsTrialLinkedNil

`func (o *OrgLiteratureRow) SetIsTrialLinkedNil(b bool)`

 SetIsTrialLinkedNil sets the value for IsTrialLinked to be an explicit nil

### UnsetIsTrialLinked
`func (o *OrgLiteratureRow) UnsetIsTrialLinked()`

UnsetIsTrialLinked ensures that no value is present for IsTrialLinked, not even an explicit nil
### GetOrgArm

`func (o *OrgLiteratureRow) GetOrgArm() string`

GetOrgArm returns the OrgArm field if non-nil, zero value otherwise.

### GetOrgArmOk

`func (o *OrgLiteratureRow) GetOrgArmOk() (*string, bool)`

GetOrgArmOk returns a tuple with the OrgArm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgArm

`func (o *OrgLiteratureRow) SetOrgArm(v string)`

SetOrgArm sets OrgArm field to given value.

### HasOrgArm

`func (o *OrgLiteratureRow) HasOrgArm() bool`

HasOrgArm returns a boolean if a field has been set.

### SetOrgArmNil

`func (o *OrgLiteratureRow) SetOrgArmNil(b bool)`

 SetOrgArmNil sets the value for OrgArm to be an explicit nil

### UnsetOrgArm
`func (o *OrgLiteratureRow) UnsetOrgArm()`

UnsetOrgArm ensures that no value is present for OrgArm, not even an explicit nil
### GetMatchTier

`func (o *OrgLiteratureRow) GetMatchTier() string`

GetMatchTier returns the MatchTier field if non-nil, zero value otherwise.

### GetMatchTierOk

`func (o *OrgLiteratureRow) GetMatchTierOk() (*string, bool)`

GetMatchTierOk returns a tuple with the MatchTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchTier

`func (o *OrgLiteratureRow) SetMatchTier(v string)`

SetMatchTier sets MatchTier field to given value.

### HasMatchTier

`func (o *OrgLiteratureRow) HasMatchTier() bool`

HasMatchTier returns a boolean if a field has been set.

### SetMatchTierNil

`func (o *OrgLiteratureRow) SetMatchTierNil(b bool)`

 SetMatchTierNil sets the value for MatchTier to be an explicit nil

### UnsetMatchTier
`func (o *OrgLiteratureRow) UnsetMatchTier()`

UnsetMatchTier ensures that no value is present for MatchTier, not even an explicit nil
### GetNBylines

`func (o *OrgLiteratureRow) GetNBylines() int32`

GetNBylines returns the NBylines field if non-nil, zero value otherwise.

### GetNBylinesOk

`func (o *OrgLiteratureRow) GetNBylinesOk() (*int32, bool)`

GetNBylinesOk returns a tuple with the NBylines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNBylines

`func (o *OrgLiteratureRow) SetNBylines(v int32)`

SetNBylines sets NBylines field to given value.

### HasNBylines

`func (o *OrgLiteratureRow) HasNBylines() bool`

HasNBylines returns a boolean if a field has been set.

### SetNBylinesNil

`func (o *OrgLiteratureRow) SetNBylinesNil(b bool)`

 SetNBylinesNil sets the value for NBylines to be an explicit nil

### UnsetNBylines
`func (o *OrgLiteratureRow) UnsetNBylines()`

UnsetNBylines ensures that no value is present for NBylines, not even an explicit nil
### GetCitationCount

`func (o *OrgLiteratureRow) GetCitationCount() int32`

GetCitationCount returns the CitationCount field if non-nil, zero value otherwise.

### GetCitationCountOk

`func (o *OrgLiteratureRow) GetCitationCountOk() (*int32, bool)`

GetCitationCountOk returns a tuple with the CitationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitationCount

`func (o *OrgLiteratureRow) SetCitationCount(v int32)`

SetCitationCount sets CitationCount field to given value.

### HasCitationCount

`func (o *OrgLiteratureRow) HasCitationCount() bool`

HasCitationCount returns a boolean if a field has been set.

### SetCitationCountNil

`func (o *OrgLiteratureRow) SetCitationCountNil(b bool)`

 SetCitationCountNil sets the value for CitationCount to be an explicit nil

### UnsetCitationCount
`func (o *OrgLiteratureRow) UnsetCitationCount()`

UnsetCitationCount ensures that no value is present for CitationCount, not even an explicit nil
### GetRelativeCitationRatio

`func (o *OrgLiteratureRow) GetRelativeCitationRatio() float32`

GetRelativeCitationRatio returns the RelativeCitationRatio field if non-nil, zero value otherwise.

### GetRelativeCitationRatioOk

`func (o *OrgLiteratureRow) GetRelativeCitationRatioOk() (*float32, bool)`

GetRelativeCitationRatioOk returns a tuple with the RelativeCitationRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeCitationRatio

`func (o *OrgLiteratureRow) SetRelativeCitationRatio(v float32)`

SetRelativeCitationRatio sets RelativeCitationRatio field to given value.

### HasRelativeCitationRatio

`func (o *OrgLiteratureRow) HasRelativeCitationRatio() bool`

HasRelativeCitationRatio returns a boolean if a field has been set.

### SetRelativeCitationRatioNil

`func (o *OrgLiteratureRow) SetRelativeCitationRatioNil(b bool)`

 SetRelativeCitationRatioNil sets the value for RelativeCitationRatio to be an explicit nil

### UnsetRelativeCitationRatio
`func (o *OrgLiteratureRow) UnsetRelativeCitationRatio()`

UnsetRelativeCitationRatio ensures that no value is present for RelativeCitationRatio, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


