# RareDiseaseRarityEvidenceRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RarityEvidenceId** | Pointer to **NullableString** | The identifier of the evidence link. | [optional] 
**OrphaCode** | Pointer to **NullableString** | The ORPHA code of the disease, the digits in the web page address (/rare-diseases/558). The entry key of the rare_diseases object type. | [optional] 
**RarityAssertionId** | Pointer to **NullableString** | The assessment the evidence belongs to. | [optional] 
**PrevalenceAssertionId** | Pointer to **NullableString** | The prevalence observation compared. | [optional] 
**JurisdictionCode** | Pointer to **NullableString** | The jurisdiction whose rule was applied (GB or EU). | [optional] 
**EvidenceKind** | Pointer to **NullableString** | How the observation was compared (prevalence_class_ratio). | [optional] 
**ComparisonResult** | Pointer to **NullableString** | The result of comparing the observation with the threshold: rare, not_rare or unknown. | [optional] 
**IsComparable** | Pointer to **NullableBool** | True when the observation could be compared with the threshold. | [optional] 
**ComparisonReason** | Pointer to **NullableString** | The reason the comparison gave its result. | [optional] 
**AssessmentDate** | Pointer to **NullableString** | The date of the assessment. | [optional] 

## Methods

### NewRareDiseaseRarityEvidenceRow

`func NewRareDiseaseRarityEvidenceRow() *RareDiseaseRarityEvidenceRow`

NewRareDiseaseRarityEvidenceRow instantiates a new RareDiseaseRarityEvidenceRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRareDiseaseRarityEvidenceRowWithDefaults

`func NewRareDiseaseRarityEvidenceRowWithDefaults() *RareDiseaseRarityEvidenceRow`

NewRareDiseaseRarityEvidenceRowWithDefaults instantiates a new RareDiseaseRarityEvidenceRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRarityEvidenceId

`func (o *RareDiseaseRarityEvidenceRow) GetRarityEvidenceId() string`

GetRarityEvidenceId returns the RarityEvidenceId field if non-nil, zero value otherwise.

### GetRarityEvidenceIdOk

`func (o *RareDiseaseRarityEvidenceRow) GetRarityEvidenceIdOk() (*string, bool)`

GetRarityEvidenceIdOk returns a tuple with the RarityEvidenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRarityEvidenceId

`func (o *RareDiseaseRarityEvidenceRow) SetRarityEvidenceId(v string)`

SetRarityEvidenceId sets RarityEvidenceId field to given value.

### HasRarityEvidenceId

`func (o *RareDiseaseRarityEvidenceRow) HasRarityEvidenceId() bool`

HasRarityEvidenceId returns a boolean if a field has been set.

### SetRarityEvidenceIdNil

`func (o *RareDiseaseRarityEvidenceRow) SetRarityEvidenceIdNil(b bool)`

 SetRarityEvidenceIdNil sets the value for RarityEvidenceId to be an explicit nil

### UnsetRarityEvidenceId
`func (o *RareDiseaseRarityEvidenceRow) UnsetRarityEvidenceId()`

UnsetRarityEvidenceId ensures that no value is present for RarityEvidenceId, not even an explicit nil
### GetOrphaCode

`func (o *RareDiseaseRarityEvidenceRow) GetOrphaCode() string`

GetOrphaCode returns the OrphaCode field if non-nil, zero value otherwise.

### GetOrphaCodeOk

`func (o *RareDiseaseRarityEvidenceRow) GetOrphaCodeOk() (*string, bool)`

GetOrphaCodeOk returns a tuple with the OrphaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphaCode

`func (o *RareDiseaseRarityEvidenceRow) SetOrphaCode(v string)`

SetOrphaCode sets OrphaCode field to given value.

### HasOrphaCode

`func (o *RareDiseaseRarityEvidenceRow) HasOrphaCode() bool`

HasOrphaCode returns a boolean if a field has been set.

### SetOrphaCodeNil

`func (o *RareDiseaseRarityEvidenceRow) SetOrphaCodeNil(b bool)`

 SetOrphaCodeNil sets the value for OrphaCode to be an explicit nil

### UnsetOrphaCode
`func (o *RareDiseaseRarityEvidenceRow) UnsetOrphaCode()`

UnsetOrphaCode ensures that no value is present for OrphaCode, not even an explicit nil
### GetRarityAssertionId

`func (o *RareDiseaseRarityEvidenceRow) GetRarityAssertionId() string`

GetRarityAssertionId returns the RarityAssertionId field if non-nil, zero value otherwise.

### GetRarityAssertionIdOk

`func (o *RareDiseaseRarityEvidenceRow) GetRarityAssertionIdOk() (*string, bool)`

GetRarityAssertionIdOk returns a tuple with the RarityAssertionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRarityAssertionId

`func (o *RareDiseaseRarityEvidenceRow) SetRarityAssertionId(v string)`

SetRarityAssertionId sets RarityAssertionId field to given value.

### HasRarityAssertionId

`func (o *RareDiseaseRarityEvidenceRow) HasRarityAssertionId() bool`

HasRarityAssertionId returns a boolean if a field has been set.

### SetRarityAssertionIdNil

`func (o *RareDiseaseRarityEvidenceRow) SetRarityAssertionIdNil(b bool)`

 SetRarityAssertionIdNil sets the value for RarityAssertionId to be an explicit nil

### UnsetRarityAssertionId
`func (o *RareDiseaseRarityEvidenceRow) UnsetRarityAssertionId()`

UnsetRarityAssertionId ensures that no value is present for RarityAssertionId, not even an explicit nil
### GetPrevalenceAssertionId

`func (o *RareDiseaseRarityEvidenceRow) GetPrevalenceAssertionId() string`

GetPrevalenceAssertionId returns the PrevalenceAssertionId field if non-nil, zero value otherwise.

### GetPrevalenceAssertionIdOk

`func (o *RareDiseaseRarityEvidenceRow) GetPrevalenceAssertionIdOk() (*string, bool)`

GetPrevalenceAssertionIdOk returns a tuple with the PrevalenceAssertionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevalenceAssertionId

`func (o *RareDiseaseRarityEvidenceRow) SetPrevalenceAssertionId(v string)`

SetPrevalenceAssertionId sets PrevalenceAssertionId field to given value.

### HasPrevalenceAssertionId

`func (o *RareDiseaseRarityEvidenceRow) HasPrevalenceAssertionId() bool`

HasPrevalenceAssertionId returns a boolean if a field has been set.

### SetPrevalenceAssertionIdNil

`func (o *RareDiseaseRarityEvidenceRow) SetPrevalenceAssertionIdNil(b bool)`

 SetPrevalenceAssertionIdNil sets the value for PrevalenceAssertionId to be an explicit nil

### UnsetPrevalenceAssertionId
`func (o *RareDiseaseRarityEvidenceRow) UnsetPrevalenceAssertionId()`

UnsetPrevalenceAssertionId ensures that no value is present for PrevalenceAssertionId, not even an explicit nil
### GetJurisdictionCode

`func (o *RareDiseaseRarityEvidenceRow) GetJurisdictionCode() string`

GetJurisdictionCode returns the JurisdictionCode field if non-nil, zero value otherwise.

### GetJurisdictionCodeOk

`func (o *RareDiseaseRarityEvidenceRow) GetJurisdictionCodeOk() (*string, bool)`

GetJurisdictionCodeOk returns a tuple with the JurisdictionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdictionCode

`func (o *RareDiseaseRarityEvidenceRow) SetJurisdictionCode(v string)`

SetJurisdictionCode sets JurisdictionCode field to given value.

### HasJurisdictionCode

`func (o *RareDiseaseRarityEvidenceRow) HasJurisdictionCode() bool`

HasJurisdictionCode returns a boolean if a field has been set.

### SetJurisdictionCodeNil

`func (o *RareDiseaseRarityEvidenceRow) SetJurisdictionCodeNil(b bool)`

 SetJurisdictionCodeNil sets the value for JurisdictionCode to be an explicit nil

### UnsetJurisdictionCode
`func (o *RareDiseaseRarityEvidenceRow) UnsetJurisdictionCode()`

UnsetJurisdictionCode ensures that no value is present for JurisdictionCode, not even an explicit nil
### GetEvidenceKind

`func (o *RareDiseaseRarityEvidenceRow) GetEvidenceKind() string`

GetEvidenceKind returns the EvidenceKind field if non-nil, zero value otherwise.

### GetEvidenceKindOk

`func (o *RareDiseaseRarityEvidenceRow) GetEvidenceKindOk() (*string, bool)`

GetEvidenceKindOk returns a tuple with the EvidenceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceKind

`func (o *RareDiseaseRarityEvidenceRow) SetEvidenceKind(v string)`

SetEvidenceKind sets EvidenceKind field to given value.

### HasEvidenceKind

`func (o *RareDiseaseRarityEvidenceRow) HasEvidenceKind() bool`

HasEvidenceKind returns a boolean if a field has been set.

### SetEvidenceKindNil

`func (o *RareDiseaseRarityEvidenceRow) SetEvidenceKindNil(b bool)`

 SetEvidenceKindNil sets the value for EvidenceKind to be an explicit nil

### UnsetEvidenceKind
`func (o *RareDiseaseRarityEvidenceRow) UnsetEvidenceKind()`

UnsetEvidenceKind ensures that no value is present for EvidenceKind, not even an explicit nil
### GetComparisonResult

`func (o *RareDiseaseRarityEvidenceRow) GetComparisonResult() string`

GetComparisonResult returns the ComparisonResult field if non-nil, zero value otherwise.

### GetComparisonResultOk

`func (o *RareDiseaseRarityEvidenceRow) GetComparisonResultOk() (*string, bool)`

GetComparisonResultOk returns a tuple with the ComparisonResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonResult

`func (o *RareDiseaseRarityEvidenceRow) SetComparisonResult(v string)`

SetComparisonResult sets ComparisonResult field to given value.

### HasComparisonResult

`func (o *RareDiseaseRarityEvidenceRow) HasComparisonResult() bool`

HasComparisonResult returns a boolean if a field has been set.

### SetComparisonResultNil

`func (o *RareDiseaseRarityEvidenceRow) SetComparisonResultNil(b bool)`

 SetComparisonResultNil sets the value for ComparisonResult to be an explicit nil

### UnsetComparisonResult
`func (o *RareDiseaseRarityEvidenceRow) UnsetComparisonResult()`

UnsetComparisonResult ensures that no value is present for ComparisonResult, not even an explicit nil
### GetIsComparable

`func (o *RareDiseaseRarityEvidenceRow) GetIsComparable() bool`

GetIsComparable returns the IsComparable field if non-nil, zero value otherwise.

### GetIsComparableOk

`func (o *RareDiseaseRarityEvidenceRow) GetIsComparableOk() (*bool, bool)`

GetIsComparableOk returns a tuple with the IsComparable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsComparable

`func (o *RareDiseaseRarityEvidenceRow) SetIsComparable(v bool)`

SetIsComparable sets IsComparable field to given value.

### HasIsComparable

`func (o *RareDiseaseRarityEvidenceRow) HasIsComparable() bool`

HasIsComparable returns a boolean if a field has been set.

### SetIsComparableNil

`func (o *RareDiseaseRarityEvidenceRow) SetIsComparableNil(b bool)`

 SetIsComparableNil sets the value for IsComparable to be an explicit nil

### UnsetIsComparable
`func (o *RareDiseaseRarityEvidenceRow) UnsetIsComparable()`

UnsetIsComparable ensures that no value is present for IsComparable, not even an explicit nil
### GetComparisonReason

`func (o *RareDiseaseRarityEvidenceRow) GetComparisonReason() string`

GetComparisonReason returns the ComparisonReason field if non-nil, zero value otherwise.

### GetComparisonReasonOk

`func (o *RareDiseaseRarityEvidenceRow) GetComparisonReasonOk() (*string, bool)`

GetComparisonReasonOk returns a tuple with the ComparisonReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonReason

`func (o *RareDiseaseRarityEvidenceRow) SetComparisonReason(v string)`

SetComparisonReason sets ComparisonReason field to given value.

### HasComparisonReason

`func (o *RareDiseaseRarityEvidenceRow) HasComparisonReason() bool`

HasComparisonReason returns a boolean if a field has been set.

### SetComparisonReasonNil

`func (o *RareDiseaseRarityEvidenceRow) SetComparisonReasonNil(b bool)`

 SetComparisonReasonNil sets the value for ComparisonReason to be an explicit nil

### UnsetComparisonReason
`func (o *RareDiseaseRarityEvidenceRow) UnsetComparisonReason()`

UnsetComparisonReason ensures that no value is present for ComparisonReason, not even an explicit nil
### GetAssessmentDate

`func (o *RareDiseaseRarityEvidenceRow) GetAssessmentDate() string`

GetAssessmentDate returns the AssessmentDate field if non-nil, zero value otherwise.

### GetAssessmentDateOk

`func (o *RareDiseaseRarityEvidenceRow) GetAssessmentDateOk() (*string, bool)`

GetAssessmentDateOk returns a tuple with the AssessmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessmentDate

`func (o *RareDiseaseRarityEvidenceRow) SetAssessmentDate(v string)`

SetAssessmentDate sets AssessmentDate field to given value.

### HasAssessmentDate

`func (o *RareDiseaseRarityEvidenceRow) HasAssessmentDate() bool`

HasAssessmentDate returns a boolean if a field has been set.

### SetAssessmentDateNil

`func (o *RareDiseaseRarityEvidenceRow) SetAssessmentDateNil(b bool)`

 SetAssessmentDateNil sets the value for AssessmentDate to be an explicit nil

### UnsetAssessmentDate
`func (o *RareDiseaseRarityEvidenceRow) UnsetAssessmentDate()`

UnsetAssessmentDate ensures that no value is present for AssessmentDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


