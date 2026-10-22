# RareDiseaseRarityRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RarityAssertionId** | Pointer to **NullableString** | The identifier of the assessment. | [optional] 
**OrphaCode** | Pointer to **NullableString** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] 
**JurisdictionCode** | Pointer to **NullableString** | The jurisdiction whose definition was applied (GB or EU). | [optional] 
**JurisdictionName** | Pointer to **NullableString** | The jurisdiction&#39;s name. | [optional] 
**RuleId** | Pointer to **NullableString** | The identifier of the rule applied. | [optional] 
**DefinitionText** | Pointer to **NullableString** | The rule&#39;s definition as stated by the jurisdiction. | [optional] 
**ComparisonOperator** | Pointer to **NullableString** | The comparison the rule applies to the threshold (&lt; or &lt;&#x3D;). | [optional] 
**ThresholdValue** | Pointer to **NullableFloat32** | The threshold figure, to be read over population_denominator (for example 5 per 10000). | [optional] 
**PopulationDenominator** | Pointer to **NullableFloat32** | The denominator of the threshold figure. | [optional] 
**PopulationUnit** | Pointer to **NullableString** | The unit of the denominator. | [optional] 
**EffectiveDate** | Pointer to **NullableString** | The date the rule took effect. | [optional] 
**EndDate** | Pointer to **NullableString** | The date the rule ended, where it has. | [optional] 
**AssessmentDate** | Pointer to **NullableString** | The date of the assessment. | [optional] 
**RarityStatus** | Pointer to **NullableString** | rare, not_rare, unknown or conflicting, as a result of applying the rule to the evidence. | [optional] 
**EvidenceBasis** | Pointer to **NullableString** | What the assessment rests on: the bounds of a source prevalence class, or no evidence in the jurisdiction&#39;s geography. | [optional] 
**EvidenceCount** | Pointer to **NullableInt64** | Prevalence observations compared in the assessment. | [optional] 
**OfficialSourceUrl** | Pointer to **NullableString** | The jurisdiction&#39;s official document the rule comes from. | [optional] 
**LicenseName** | Pointer to **NullableString** | The licence of the jurisdiction&#39;s document. | [optional] 
**Attribution** | Pointer to **NullableString** | The attribution line the source requires; quote it with any figure taken from the row. | [optional] 

## Methods

### NewRareDiseaseRarityRow

`func NewRareDiseaseRarityRow() *RareDiseaseRarityRow`

NewRareDiseaseRarityRow instantiates a new RareDiseaseRarityRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRareDiseaseRarityRowWithDefaults

`func NewRareDiseaseRarityRowWithDefaults() *RareDiseaseRarityRow`

NewRareDiseaseRarityRowWithDefaults instantiates a new RareDiseaseRarityRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRarityAssertionId

`func (o *RareDiseaseRarityRow) GetRarityAssertionId() string`

GetRarityAssertionId returns the RarityAssertionId field if non-nil, zero value otherwise.

### GetRarityAssertionIdOk

`func (o *RareDiseaseRarityRow) GetRarityAssertionIdOk() (*string, bool)`

GetRarityAssertionIdOk returns a tuple with the RarityAssertionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRarityAssertionId

`func (o *RareDiseaseRarityRow) SetRarityAssertionId(v string)`

SetRarityAssertionId sets RarityAssertionId field to given value.

### HasRarityAssertionId

`func (o *RareDiseaseRarityRow) HasRarityAssertionId() bool`

HasRarityAssertionId returns a boolean if a field has been set.

### SetRarityAssertionIdNil

`func (o *RareDiseaseRarityRow) SetRarityAssertionIdNil(b bool)`

 SetRarityAssertionIdNil sets the value for RarityAssertionId to be an explicit nil

### UnsetRarityAssertionId
`func (o *RareDiseaseRarityRow) UnsetRarityAssertionId()`

UnsetRarityAssertionId ensures that no value is present for RarityAssertionId, not even an explicit nil
### GetOrphaCode

`func (o *RareDiseaseRarityRow) GetOrphaCode() string`

GetOrphaCode returns the OrphaCode field if non-nil, zero value otherwise.

### GetOrphaCodeOk

`func (o *RareDiseaseRarityRow) GetOrphaCodeOk() (*string, bool)`

GetOrphaCodeOk returns a tuple with the OrphaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphaCode

`func (o *RareDiseaseRarityRow) SetOrphaCode(v string)`

SetOrphaCode sets OrphaCode field to given value.

### HasOrphaCode

`func (o *RareDiseaseRarityRow) HasOrphaCode() bool`

HasOrphaCode returns a boolean if a field has been set.

### SetOrphaCodeNil

`func (o *RareDiseaseRarityRow) SetOrphaCodeNil(b bool)`

 SetOrphaCodeNil sets the value for OrphaCode to be an explicit nil

### UnsetOrphaCode
`func (o *RareDiseaseRarityRow) UnsetOrphaCode()`

UnsetOrphaCode ensures that no value is present for OrphaCode, not even an explicit nil
### GetJurisdictionCode

`func (o *RareDiseaseRarityRow) GetJurisdictionCode() string`

GetJurisdictionCode returns the JurisdictionCode field if non-nil, zero value otherwise.

### GetJurisdictionCodeOk

`func (o *RareDiseaseRarityRow) GetJurisdictionCodeOk() (*string, bool)`

GetJurisdictionCodeOk returns a tuple with the JurisdictionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdictionCode

`func (o *RareDiseaseRarityRow) SetJurisdictionCode(v string)`

SetJurisdictionCode sets JurisdictionCode field to given value.

### HasJurisdictionCode

`func (o *RareDiseaseRarityRow) HasJurisdictionCode() bool`

HasJurisdictionCode returns a boolean if a field has been set.

### SetJurisdictionCodeNil

`func (o *RareDiseaseRarityRow) SetJurisdictionCodeNil(b bool)`

 SetJurisdictionCodeNil sets the value for JurisdictionCode to be an explicit nil

### UnsetJurisdictionCode
`func (o *RareDiseaseRarityRow) UnsetJurisdictionCode()`

UnsetJurisdictionCode ensures that no value is present for JurisdictionCode, not even an explicit nil
### GetJurisdictionName

`func (o *RareDiseaseRarityRow) GetJurisdictionName() string`

GetJurisdictionName returns the JurisdictionName field if non-nil, zero value otherwise.

### GetJurisdictionNameOk

`func (o *RareDiseaseRarityRow) GetJurisdictionNameOk() (*string, bool)`

GetJurisdictionNameOk returns a tuple with the JurisdictionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdictionName

`func (o *RareDiseaseRarityRow) SetJurisdictionName(v string)`

SetJurisdictionName sets JurisdictionName field to given value.

### HasJurisdictionName

`func (o *RareDiseaseRarityRow) HasJurisdictionName() bool`

HasJurisdictionName returns a boolean if a field has been set.

### SetJurisdictionNameNil

`func (o *RareDiseaseRarityRow) SetJurisdictionNameNil(b bool)`

 SetJurisdictionNameNil sets the value for JurisdictionName to be an explicit nil

### UnsetJurisdictionName
`func (o *RareDiseaseRarityRow) UnsetJurisdictionName()`

UnsetJurisdictionName ensures that no value is present for JurisdictionName, not even an explicit nil
### GetRuleId

`func (o *RareDiseaseRarityRow) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *RareDiseaseRarityRow) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *RareDiseaseRarityRow) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *RareDiseaseRarityRow) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### SetRuleIdNil

`func (o *RareDiseaseRarityRow) SetRuleIdNil(b bool)`

 SetRuleIdNil sets the value for RuleId to be an explicit nil

### UnsetRuleId
`func (o *RareDiseaseRarityRow) UnsetRuleId()`

UnsetRuleId ensures that no value is present for RuleId, not even an explicit nil
### GetDefinitionText

`func (o *RareDiseaseRarityRow) GetDefinitionText() string`

GetDefinitionText returns the DefinitionText field if non-nil, zero value otherwise.

### GetDefinitionTextOk

`func (o *RareDiseaseRarityRow) GetDefinitionTextOk() (*string, bool)`

GetDefinitionTextOk returns a tuple with the DefinitionText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionText

`func (o *RareDiseaseRarityRow) SetDefinitionText(v string)`

SetDefinitionText sets DefinitionText field to given value.

### HasDefinitionText

`func (o *RareDiseaseRarityRow) HasDefinitionText() bool`

HasDefinitionText returns a boolean if a field has been set.

### SetDefinitionTextNil

`func (o *RareDiseaseRarityRow) SetDefinitionTextNil(b bool)`

 SetDefinitionTextNil sets the value for DefinitionText to be an explicit nil

### UnsetDefinitionText
`func (o *RareDiseaseRarityRow) UnsetDefinitionText()`

UnsetDefinitionText ensures that no value is present for DefinitionText, not even an explicit nil
### GetComparisonOperator

`func (o *RareDiseaseRarityRow) GetComparisonOperator() string`

GetComparisonOperator returns the ComparisonOperator field if non-nil, zero value otherwise.

### GetComparisonOperatorOk

`func (o *RareDiseaseRarityRow) GetComparisonOperatorOk() (*string, bool)`

GetComparisonOperatorOk returns a tuple with the ComparisonOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonOperator

`func (o *RareDiseaseRarityRow) SetComparisonOperator(v string)`

SetComparisonOperator sets ComparisonOperator field to given value.

### HasComparisonOperator

`func (o *RareDiseaseRarityRow) HasComparisonOperator() bool`

HasComparisonOperator returns a boolean if a field has been set.

### SetComparisonOperatorNil

`func (o *RareDiseaseRarityRow) SetComparisonOperatorNil(b bool)`

 SetComparisonOperatorNil sets the value for ComparisonOperator to be an explicit nil

### UnsetComparisonOperator
`func (o *RareDiseaseRarityRow) UnsetComparisonOperator()`

UnsetComparisonOperator ensures that no value is present for ComparisonOperator, not even an explicit nil
### GetThresholdValue

`func (o *RareDiseaseRarityRow) GetThresholdValue() float32`

GetThresholdValue returns the ThresholdValue field if non-nil, zero value otherwise.

### GetThresholdValueOk

`func (o *RareDiseaseRarityRow) GetThresholdValueOk() (*float32, bool)`

GetThresholdValueOk returns a tuple with the ThresholdValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdValue

`func (o *RareDiseaseRarityRow) SetThresholdValue(v float32)`

SetThresholdValue sets ThresholdValue field to given value.

### HasThresholdValue

`func (o *RareDiseaseRarityRow) HasThresholdValue() bool`

HasThresholdValue returns a boolean if a field has been set.

### SetThresholdValueNil

`func (o *RareDiseaseRarityRow) SetThresholdValueNil(b bool)`

 SetThresholdValueNil sets the value for ThresholdValue to be an explicit nil

### UnsetThresholdValue
`func (o *RareDiseaseRarityRow) UnsetThresholdValue()`

UnsetThresholdValue ensures that no value is present for ThresholdValue, not even an explicit nil
### GetPopulationDenominator

`func (o *RareDiseaseRarityRow) GetPopulationDenominator() float32`

GetPopulationDenominator returns the PopulationDenominator field if non-nil, zero value otherwise.

### GetPopulationDenominatorOk

`func (o *RareDiseaseRarityRow) GetPopulationDenominatorOk() (*float32, bool)`

GetPopulationDenominatorOk returns a tuple with the PopulationDenominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulationDenominator

`func (o *RareDiseaseRarityRow) SetPopulationDenominator(v float32)`

SetPopulationDenominator sets PopulationDenominator field to given value.

### HasPopulationDenominator

`func (o *RareDiseaseRarityRow) HasPopulationDenominator() bool`

HasPopulationDenominator returns a boolean if a field has been set.

### SetPopulationDenominatorNil

`func (o *RareDiseaseRarityRow) SetPopulationDenominatorNil(b bool)`

 SetPopulationDenominatorNil sets the value for PopulationDenominator to be an explicit nil

### UnsetPopulationDenominator
`func (o *RareDiseaseRarityRow) UnsetPopulationDenominator()`

UnsetPopulationDenominator ensures that no value is present for PopulationDenominator, not even an explicit nil
### GetPopulationUnit

`func (o *RareDiseaseRarityRow) GetPopulationUnit() string`

GetPopulationUnit returns the PopulationUnit field if non-nil, zero value otherwise.

### GetPopulationUnitOk

`func (o *RareDiseaseRarityRow) GetPopulationUnitOk() (*string, bool)`

GetPopulationUnitOk returns a tuple with the PopulationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulationUnit

`func (o *RareDiseaseRarityRow) SetPopulationUnit(v string)`

SetPopulationUnit sets PopulationUnit field to given value.

### HasPopulationUnit

`func (o *RareDiseaseRarityRow) HasPopulationUnit() bool`

HasPopulationUnit returns a boolean if a field has been set.

### SetPopulationUnitNil

`func (o *RareDiseaseRarityRow) SetPopulationUnitNil(b bool)`

 SetPopulationUnitNil sets the value for PopulationUnit to be an explicit nil

### UnsetPopulationUnit
`func (o *RareDiseaseRarityRow) UnsetPopulationUnit()`

UnsetPopulationUnit ensures that no value is present for PopulationUnit, not even an explicit nil
### GetEffectiveDate

`func (o *RareDiseaseRarityRow) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *RareDiseaseRarityRow) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *RareDiseaseRarityRow) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *RareDiseaseRarityRow) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### SetEffectiveDateNil

`func (o *RareDiseaseRarityRow) SetEffectiveDateNil(b bool)`

 SetEffectiveDateNil sets the value for EffectiveDate to be an explicit nil

### UnsetEffectiveDate
`func (o *RareDiseaseRarityRow) UnsetEffectiveDate()`

UnsetEffectiveDate ensures that no value is present for EffectiveDate, not even an explicit nil
### GetEndDate

`func (o *RareDiseaseRarityRow) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *RareDiseaseRarityRow) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *RareDiseaseRarityRow) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *RareDiseaseRarityRow) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *RareDiseaseRarityRow) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *RareDiseaseRarityRow) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetAssessmentDate

`func (o *RareDiseaseRarityRow) GetAssessmentDate() string`

GetAssessmentDate returns the AssessmentDate field if non-nil, zero value otherwise.

### GetAssessmentDateOk

`func (o *RareDiseaseRarityRow) GetAssessmentDateOk() (*string, bool)`

GetAssessmentDateOk returns a tuple with the AssessmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessmentDate

`func (o *RareDiseaseRarityRow) SetAssessmentDate(v string)`

SetAssessmentDate sets AssessmentDate field to given value.

### HasAssessmentDate

`func (o *RareDiseaseRarityRow) HasAssessmentDate() bool`

HasAssessmentDate returns a boolean if a field has been set.

### SetAssessmentDateNil

`func (o *RareDiseaseRarityRow) SetAssessmentDateNil(b bool)`

 SetAssessmentDateNil sets the value for AssessmentDate to be an explicit nil

### UnsetAssessmentDate
`func (o *RareDiseaseRarityRow) UnsetAssessmentDate()`

UnsetAssessmentDate ensures that no value is present for AssessmentDate, not even an explicit nil
### GetRarityStatus

`func (o *RareDiseaseRarityRow) GetRarityStatus() string`

GetRarityStatus returns the RarityStatus field if non-nil, zero value otherwise.

### GetRarityStatusOk

`func (o *RareDiseaseRarityRow) GetRarityStatusOk() (*string, bool)`

GetRarityStatusOk returns a tuple with the RarityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRarityStatus

`func (o *RareDiseaseRarityRow) SetRarityStatus(v string)`

SetRarityStatus sets RarityStatus field to given value.

### HasRarityStatus

`func (o *RareDiseaseRarityRow) HasRarityStatus() bool`

HasRarityStatus returns a boolean if a field has been set.

### SetRarityStatusNil

`func (o *RareDiseaseRarityRow) SetRarityStatusNil(b bool)`

 SetRarityStatusNil sets the value for RarityStatus to be an explicit nil

### UnsetRarityStatus
`func (o *RareDiseaseRarityRow) UnsetRarityStatus()`

UnsetRarityStatus ensures that no value is present for RarityStatus, not even an explicit nil
### GetEvidenceBasis

`func (o *RareDiseaseRarityRow) GetEvidenceBasis() string`

GetEvidenceBasis returns the EvidenceBasis field if non-nil, zero value otherwise.

### GetEvidenceBasisOk

`func (o *RareDiseaseRarityRow) GetEvidenceBasisOk() (*string, bool)`

GetEvidenceBasisOk returns a tuple with the EvidenceBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceBasis

`func (o *RareDiseaseRarityRow) SetEvidenceBasis(v string)`

SetEvidenceBasis sets EvidenceBasis field to given value.

### HasEvidenceBasis

`func (o *RareDiseaseRarityRow) HasEvidenceBasis() bool`

HasEvidenceBasis returns a boolean if a field has been set.

### SetEvidenceBasisNil

`func (o *RareDiseaseRarityRow) SetEvidenceBasisNil(b bool)`

 SetEvidenceBasisNil sets the value for EvidenceBasis to be an explicit nil

### UnsetEvidenceBasis
`func (o *RareDiseaseRarityRow) UnsetEvidenceBasis()`

UnsetEvidenceBasis ensures that no value is present for EvidenceBasis, not even an explicit nil
### GetEvidenceCount

`func (o *RareDiseaseRarityRow) GetEvidenceCount() int64`

GetEvidenceCount returns the EvidenceCount field if non-nil, zero value otherwise.

### GetEvidenceCountOk

`func (o *RareDiseaseRarityRow) GetEvidenceCountOk() (*int64, bool)`

GetEvidenceCountOk returns a tuple with the EvidenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceCount

`func (o *RareDiseaseRarityRow) SetEvidenceCount(v int64)`

SetEvidenceCount sets EvidenceCount field to given value.

### HasEvidenceCount

`func (o *RareDiseaseRarityRow) HasEvidenceCount() bool`

HasEvidenceCount returns a boolean if a field has been set.

### SetEvidenceCountNil

`func (o *RareDiseaseRarityRow) SetEvidenceCountNil(b bool)`

 SetEvidenceCountNil sets the value for EvidenceCount to be an explicit nil

### UnsetEvidenceCount
`func (o *RareDiseaseRarityRow) UnsetEvidenceCount()`

UnsetEvidenceCount ensures that no value is present for EvidenceCount, not even an explicit nil
### GetOfficialSourceUrl

`func (o *RareDiseaseRarityRow) GetOfficialSourceUrl() string`

GetOfficialSourceUrl returns the OfficialSourceUrl field if non-nil, zero value otherwise.

### GetOfficialSourceUrlOk

`func (o *RareDiseaseRarityRow) GetOfficialSourceUrlOk() (*string, bool)`

GetOfficialSourceUrlOk returns a tuple with the OfficialSourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialSourceUrl

`func (o *RareDiseaseRarityRow) SetOfficialSourceUrl(v string)`

SetOfficialSourceUrl sets OfficialSourceUrl field to given value.

### HasOfficialSourceUrl

`func (o *RareDiseaseRarityRow) HasOfficialSourceUrl() bool`

HasOfficialSourceUrl returns a boolean if a field has been set.

### SetOfficialSourceUrlNil

`func (o *RareDiseaseRarityRow) SetOfficialSourceUrlNil(b bool)`

 SetOfficialSourceUrlNil sets the value for OfficialSourceUrl to be an explicit nil

### UnsetOfficialSourceUrl
`func (o *RareDiseaseRarityRow) UnsetOfficialSourceUrl()`

UnsetOfficialSourceUrl ensures that no value is present for OfficialSourceUrl, not even an explicit nil
### GetLicenseName

`func (o *RareDiseaseRarityRow) GetLicenseName() string`

GetLicenseName returns the LicenseName field if non-nil, zero value otherwise.

### GetLicenseNameOk

`func (o *RareDiseaseRarityRow) GetLicenseNameOk() (*string, bool)`

GetLicenseNameOk returns a tuple with the LicenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseName

`func (o *RareDiseaseRarityRow) SetLicenseName(v string)`

SetLicenseName sets LicenseName field to given value.

### HasLicenseName

`func (o *RareDiseaseRarityRow) HasLicenseName() bool`

HasLicenseName returns a boolean if a field has been set.

### SetLicenseNameNil

`func (o *RareDiseaseRarityRow) SetLicenseNameNil(b bool)`

 SetLicenseNameNil sets the value for LicenseName to be an explicit nil

### UnsetLicenseName
`func (o *RareDiseaseRarityRow) UnsetLicenseName()`

UnsetLicenseName ensures that no value is present for LicenseName, not even an explicit nil
### GetAttribution

`func (o *RareDiseaseRarityRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *RareDiseaseRarityRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *RareDiseaseRarityRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *RareDiseaseRarityRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *RareDiseaseRarityRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *RareDiseaseRarityRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


