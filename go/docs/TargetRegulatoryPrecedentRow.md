# TargetRegulatoryPrecedentRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | Pointer to **NullableString** | The target&#39;s Ensembl gene id (ENSG followed by eleven digits): the entry key for the targets object type. A gene symbol is not a value here; it is a separate column. | [optional] 
**TargetSymbol** | Pointer to **NullableString** | The target&#39;s HGNC gene symbol (the id the target page address carries). | [optional] 
**DrugId** | Pointer to **NullableString** | The drug&#39;s canonical id (the same id the drug datasets use as drug_id). | [optional] 
**DrugName** | Pointer to **NullableString** | The drug&#39;s name. | [optional] 
**Authority** | Pointer to **NullableString** | The approving authority. | [optional] 
**Jurisdiction** | Pointer to **NullableString** | The authority&#39;s jurisdiction. | [optional] 
**Status** | Pointer to **NullableString** | The approval status the authority records. | [optional] 
**FirstApprovalDate** | Pointer to **NullableString** | The authority&#39;s first approval date of the drug. | [optional] 
**ApprovalId** | Pointer to **NullableString** | The authority&#39;s own approval identifier. | [optional] 
**PrimaryActionType** | Pointer to **NullableString** | The action type of the drug&#39;s curated mechanism on the target. | [optional] 
**MoaEvidenceSource** | Pointer to **NullableString** | The evidence tier of the mechanism; curated_moa on every row. | [optional] 

## Methods

### NewTargetRegulatoryPrecedentRow

`func NewTargetRegulatoryPrecedentRow() *TargetRegulatoryPrecedentRow`

NewTargetRegulatoryPrecedentRow instantiates a new TargetRegulatoryPrecedentRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetRegulatoryPrecedentRowWithDefaults

`func NewTargetRegulatoryPrecedentRowWithDefaults() *TargetRegulatoryPrecedentRow`

NewTargetRegulatoryPrecedentRowWithDefaults instantiates a new TargetRegulatoryPrecedentRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetId

`func (o *TargetRegulatoryPrecedentRow) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *TargetRegulatoryPrecedentRow) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *TargetRegulatoryPrecedentRow) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *TargetRegulatoryPrecedentRow) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *TargetRegulatoryPrecedentRow) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *TargetRegulatoryPrecedentRow) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetTargetSymbol

`func (o *TargetRegulatoryPrecedentRow) GetTargetSymbol() string`

GetTargetSymbol returns the TargetSymbol field if non-nil, zero value otherwise.

### GetTargetSymbolOk

`func (o *TargetRegulatoryPrecedentRow) GetTargetSymbolOk() (*string, bool)`

GetTargetSymbolOk returns a tuple with the TargetSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSymbol

`func (o *TargetRegulatoryPrecedentRow) SetTargetSymbol(v string)`

SetTargetSymbol sets TargetSymbol field to given value.

### HasTargetSymbol

`func (o *TargetRegulatoryPrecedentRow) HasTargetSymbol() bool`

HasTargetSymbol returns a boolean if a field has been set.

### SetTargetSymbolNil

`func (o *TargetRegulatoryPrecedentRow) SetTargetSymbolNil(b bool)`

 SetTargetSymbolNil sets the value for TargetSymbol to be an explicit nil

### UnsetTargetSymbol
`func (o *TargetRegulatoryPrecedentRow) UnsetTargetSymbol()`

UnsetTargetSymbol ensures that no value is present for TargetSymbol, not even an explicit nil
### GetDrugId

`func (o *TargetRegulatoryPrecedentRow) GetDrugId() string`

GetDrugId returns the DrugId field if non-nil, zero value otherwise.

### GetDrugIdOk

`func (o *TargetRegulatoryPrecedentRow) GetDrugIdOk() (*string, bool)`

GetDrugIdOk returns a tuple with the DrugId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugId

`func (o *TargetRegulatoryPrecedentRow) SetDrugId(v string)`

SetDrugId sets DrugId field to given value.

### HasDrugId

`func (o *TargetRegulatoryPrecedentRow) HasDrugId() bool`

HasDrugId returns a boolean if a field has been set.

### SetDrugIdNil

`func (o *TargetRegulatoryPrecedentRow) SetDrugIdNil(b bool)`

 SetDrugIdNil sets the value for DrugId to be an explicit nil

### UnsetDrugId
`func (o *TargetRegulatoryPrecedentRow) UnsetDrugId()`

UnsetDrugId ensures that no value is present for DrugId, not even an explicit nil
### GetDrugName

`func (o *TargetRegulatoryPrecedentRow) GetDrugName() string`

GetDrugName returns the DrugName field if non-nil, zero value otherwise.

### GetDrugNameOk

`func (o *TargetRegulatoryPrecedentRow) GetDrugNameOk() (*string, bool)`

GetDrugNameOk returns a tuple with the DrugName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugName

`func (o *TargetRegulatoryPrecedentRow) SetDrugName(v string)`

SetDrugName sets DrugName field to given value.

### HasDrugName

`func (o *TargetRegulatoryPrecedentRow) HasDrugName() bool`

HasDrugName returns a boolean if a field has been set.

### SetDrugNameNil

`func (o *TargetRegulatoryPrecedentRow) SetDrugNameNil(b bool)`

 SetDrugNameNil sets the value for DrugName to be an explicit nil

### UnsetDrugName
`func (o *TargetRegulatoryPrecedentRow) UnsetDrugName()`

UnsetDrugName ensures that no value is present for DrugName, not even an explicit nil
### GetAuthority

`func (o *TargetRegulatoryPrecedentRow) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *TargetRegulatoryPrecedentRow) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *TargetRegulatoryPrecedentRow) SetAuthority(v string)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *TargetRegulatoryPrecedentRow) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### SetAuthorityNil

`func (o *TargetRegulatoryPrecedentRow) SetAuthorityNil(b bool)`

 SetAuthorityNil sets the value for Authority to be an explicit nil

### UnsetAuthority
`func (o *TargetRegulatoryPrecedentRow) UnsetAuthority()`

UnsetAuthority ensures that no value is present for Authority, not even an explicit nil
### GetJurisdiction

`func (o *TargetRegulatoryPrecedentRow) GetJurisdiction() string`

GetJurisdiction returns the Jurisdiction field if non-nil, zero value otherwise.

### GetJurisdictionOk

`func (o *TargetRegulatoryPrecedentRow) GetJurisdictionOk() (*string, bool)`

GetJurisdictionOk returns a tuple with the Jurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdiction

`func (o *TargetRegulatoryPrecedentRow) SetJurisdiction(v string)`

SetJurisdiction sets Jurisdiction field to given value.

### HasJurisdiction

`func (o *TargetRegulatoryPrecedentRow) HasJurisdiction() bool`

HasJurisdiction returns a boolean if a field has been set.

### SetJurisdictionNil

`func (o *TargetRegulatoryPrecedentRow) SetJurisdictionNil(b bool)`

 SetJurisdictionNil sets the value for Jurisdiction to be an explicit nil

### UnsetJurisdiction
`func (o *TargetRegulatoryPrecedentRow) UnsetJurisdiction()`

UnsetJurisdiction ensures that no value is present for Jurisdiction, not even an explicit nil
### GetStatus

`func (o *TargetRegulatoryPrecedentRow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TargetRegulatoryPrecedentRow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TargetRegulatoryPrecedentRow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TargetRegulatoryPrecedentRow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TargetRegulatoryPrecedentRow) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TargetRegulatoryPrecedentRow) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetFirstApprovalDate

`func (o *TargetRegulatoryPrecedentRow) GetFirstApprovalDate() string`

GetFirstApprovalDate returns the FirstApprovalDate field if non-nil, zero value otherwise.

### GetFirstApprovalDateOk

`func (o *TargetRegulatoryPrecedentRow) GetFirstApprovalDateOk() (*string, bool)`

GetFirstApprovalDateOk returns a tuple with the FirstApprovalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstApprovalDate

`func (o *TargetRegulatoryPrecedentRow) SetFirstApprovalDate(v string)`

SetFirstApprovalDate sets FirstApprovalDate field to given value.

### HasFirstApprovalDate

`func (o *TargetRegulatoryPrecedentRow) HasFirstApprovalDate() bool`

HasFirstApprovalDate returns a boolean if a field has been set.

### SetFirstApprovalDateNil

`func (o *TargetRegulatoryPrecedentRow) SetFirstApprovalDateNil(b bool)`

 SetFirstApprovalDateNil sets the value for FirstApprovalDate to be an explicit nil

### UnsetFirstApprovalDate
`func (o *TargetRegulatoryPrecedentRow) UnsetFirstApprovalDate()`

UnsetFirstApprovalDate ensures that no value is present for FirstApprovalDate, not even an explicit nil
### GetApprovalId

`func (o *TargetRegulatoryPrecedentRow) GetApprovalId() string`

GetApprovalId returns the ApprovalId field if non-nil, zero value otherwise.

### GetApprovalIdOk

`func (o *TargetRegulatoryPrecedentRow) GetApprovalIdOk() (*string, bool)`

GetApprovalIdOk returns a tuple with the ApprovalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalId

`func (o *TargetRegulatoryPrecedentRow) SetApprovalId(v string)`

SetApprovalId sets ApprovalId field to given value.

### HasApprovalId

`func (o *TargetRegulatoryPrecedentRow) HasApprovalId() bool`

HasApprovalId returns a boolean if a field has been set.

### SetApprovalIdNil

`func (o *TargetRegulatoryPrecedentRow) SetApprovalIdNil(b bool)`

 SetApprovalIdNil sets the value for ApprovalId to be an explicit nil

### UnsetApprovalId
`func (o *TargetRegulatoryPrecedentRow) UnsetApprovalId()`

UnsetApprovalId ensures that no value is present for ApprovalId, not even an explicit nil
### GetPrimaryActionType

`func (o *TargetRegulatoryPrecedentRow) GetPrimaryActionType() string`

GetPrimaryActionType returns the PrimaryActionType field if non-nil, zero value otherwise.

### GetPrimaryActionTypeOk

`func (o *TargetRegulatoryPrecedentRow) GetPrimaryActionTypeOk() (*string, bool)`

GetPrimaryActionTypeOk returns a tuple with the PrimaryActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryActionType

`func (o *TargetRegulatoryPrecedentRow) SetPrimaryActionType(v string)`

SetPrimaryActionType sets PrimaryActionType field to given value.

### HasPrimaryActionType

`func (o *TargetRegulatoryPrecedentRow) HasPrimaryActionType() bool`

HasPrimaryActionType returns a boolean if a field has been set.

### SetPrimaryActionTypeNil

`func (o *TargetRegulatoryPrecedentRow) SetPrimaryActionTypeNil(b bool)`

 SetPrimaryActionTypeNil sets the value for PrimaryActionType to be an explicit nil

### UnsetPrimaryActionType
`func (o *TargetRegulatoryPrecedentRow) UnsetPrimaryActionType()`

UnsetPrimaryActionType ensures that no value is present for PrimaryActionType, not even an explicit nil
### GetMoaEvidenceSource

`func (o *TargetRegulatoryPrecedentRow) GetMoaEvidenceSource() string`

GetMoaEvidenceSource returns the MoaEvidenceSource field if non-nil, zero value otherwise.

### GetMoaEvidenceSourceOk

`func (o *TargetRegulatoryPrecedentRow) GetMoaEvidenceSourceOk() (*string, bool)`

GetMoaEvidenceSourceOk returns a tuple with the MoaEvidenceSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoaEvidenceSource

`func (o *TargetRegulatoryPrecedentRow) SetMoaEvidenceSource(v string)`

SetMoaEvidenceSource sets MoaEvidenceSource field to given value.

### HasMoaEvidenceSource

`func (o *TargetRegulatoryPrecedentRow) HasMoaEvidenceSource() bool`

HasMoaEvidenceSource returns a boolean if a field has been set.

### SetMoaEvidenceSourceNil

`func (o *TargetRegulatoryPrecedentRow) SetMoaEvidenceSourceNil(b bool)`

 SetMoaEvidenceSourceNil sets the value for MoaEvidenceSource to be an explicit nil

### UnsetMoaEvidenceSource
`func (o *TargetRegulatoryPrecedentRow) UnsetMoaEvidenceSource()`

UnsetMoaEvidenceSource ensures that no value is present for MoaEvidenceSource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


