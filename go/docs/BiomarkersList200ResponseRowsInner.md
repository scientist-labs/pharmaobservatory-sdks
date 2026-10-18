# BiomarkersList200ResponseRowsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouteKey** | **string** |  | 
**DisplayName** | **string** |  | 
**BiomarkerKind** | **string** |  | 
**IdentityStatus** | **string** |  | 
**GeneSymbol** | **NullableString** |  | 
**AminoAcidChange** | **interface{}** |  | 
**NFdaCdxAssertions** | **int32** |  | 
**NFdaBqpAssertions** | **int32** |  | 
**NCancerActionableCgi** | **int32** |  | 
**NOtAssociation** | **int32** |  | 
**NTrialMention** | **int32** |  | 
**NPublicationVariantMentions** | **int32** |  | 
**KindLabel** | **string** |  | 
**EvidenceTier** | **int32** |  | 
**EvidenceTierLabel** | **string** |  | 
**HasClinicalValidationEvidence** | **NullableBool** |  | 
**HasAnalyticalValidationEvidence** | **NullableBool** |  | 
**NAssayProviderPapers** | **NullableInt32** |  | 
**NAssayProviderSuppliers** | **NullableInt32** |  | 
**NDiagnosticProviderSuppliers** | **NullableInt32** |  | 

## Methods

### NewBiomarkersList200ResponseRowsInner

`func NewBiomarkersList200ResponseRowsInner(routeKey string, displayName string, biomarkerKind string, identityStatus string, geneSymbol NullableString, aminoAcidChange interface{}, nFdaCdxAssertions int32, nFdaBqpAssertions int32, nCancerActionableCgi int32, nOtAssociation int32, nTrialMention int32, nPublicationVariantMentions int32, kindLabel string, evidenceTier int32, evidenceTierLabel string, hasClinicalValidationEvidence NullableBool, hasAnalyticalValidationEvidence NullableBool, nAssayProviderPapers NullableInt32, nAssayProviderSuppliers NullableInt32, nDiagnosticProviderSuppliers NullableInt32, ) *BiomarkersList200ResponseRowsInner`

NewBiomarkersList200ResponseRowsInner instantiates a new BiomarkersList200ResponseRowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiomarkersList200ResponseRowsInnerWithDefaults

`func NewBiomarkersList200ResponseRowsInnerWithDefaults() *BiomarkersList200ResponseRowsInner`

NewBiomarkersList200ResponseRowsInnerWithDefaults instantiates a new BiomarkersList200ResponseRowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouteKey

`func (o *BiomarkersList200ResponseRowsInner) GetRouteKey() string`

GetRouteKey returns the RouteKey field if non-nil, zero value otherwise.

### GetRouteKeyOk

`func (o *BiomarkersList200ResponseRowsInner) GetRouteKeyOk() (*string, bool)`

GetRouteKeyOk returns a tuple with the RouteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteKey

`func (o *BiomarkersList200ResponseRowsInner) SetRouteKey(v string)`

SetRouteKey sets RouteKey field to given value.


### GetDisplayName

`func (o *BiomarkersList200ResponseRowsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BiomarkersList200ResponseRowsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BiomarkersList200ResponseRowsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetBiomarkerKind

`func (o *BiomarkersList200ResponseRowsInner) GetBiomarkerKind() string`

GetBiomarkerKind returns the BiomarkerKind field if non-nil, zero value otherwise.

### GetBiomarkerKindOk

`func (o *BiomarkersList200ResponseRowsInner) GetBiomarkerKindOk() (*string, bool)`

GetBiomarkerKindOk returns a tuple with the BiomarkerKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiomarkerKind

`func (o *BiomarkersList200ResponseRowsInner) SetBiomarkerKind(v string)`

SetBiomarkerKind sets BiomarkerKind field to given value.


### GetIdentityStatus

`func (o *BiomarkersList200ResponseRowsInner) GetIdentityStatus() string`

GetIdentityStatus returns the IdentityStatus field if non-nil, zero value otherwise.

### GetIdentityStatusOk

`func (o *BiomarkersList200ResponseRowsInner) GetIdentityStatusOk() (*string, bool)`

GetIdentityStatusOk returns a tuple with the IdentityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityStatus

`func (o *BiomarkersList200ResponseRowsInner) SetIdentityStatus(v string)`

SetIdentityStatus sets IdentityStatus field to given value.


### GetGeneSymbol

`func (o *BiomarkersList200ResponseRowsInner) GetGeneSymbol() string`

GetGeneSymbol returns the GeneSymbol field if non-nil, zero value otherwise.

### GetGeneSymbolOk

`func (o *BiomarkersList200ResponseRowsInner) GetGeneSymbolOk() (*string, bool)`

GetGeneSymbolOk returns a tuple with the GeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSymbol

`func (o *BiomarkersList200ResponseRowsInner) SetGeneSymbol(v string)`

SetGeneSymbol sets GeneSymbol field to given value.


### SetGeneSymbolNil

`func (o *BiomarkersList200ResponseRowsInner) SetGeneSymbolNil(b bool)`

 SetGeneSymbolNil sets the value for GeneSymbol to be an explicit nil

### UnsetGeneSymbol
`func (o *BiomarkersList200ResponseRowsInner) UnsetGeneSymbol()`

UnsetGeneSymbol ensures that no value is present for GeneSymbol, not even an explicit nil
### GetAminoAcidChange

`func (o *BiomarkersList200ResponseRowsInner) GetAminoAcidChange() interface{}`

GetAminoAcidChange returns the AminoAcidChange field if non-nil, zero value otherwise.

### GetAminoAcidChangeOk

`func (o *BiomarkersList200ResponseRowsInner) GetAminoAcidChangeOk() (*interface{}, bool)`

GetAminoAcidChangeOk returns a tuple with the AminoAcidChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAminoAcidChange

`func (o *BiomarkersList200ResponseRowsInner) SetAminoAcidChange(v interface{})`

SetAminoAcidChange sets AminoAcidChange field to given value.


### SetAminoAcidChangeNil

`func (o *BiomarkersList200ResponseRowsInner) SetAminoAcidChangeNil(b bool)`

 SetAminoAcidChangeNil sets the value for AminoAcidChange to be an explicit nil

### UnsetAminoAcidChange
`func (o *BiomarkersList200ResponseRowsInner) UnsetAminoAcidChange()`

UnsetAminoAcidChange ensures that no value is present for AminoAcidChange, not even an explicit nil
### GetNFdaCdxAssertions

`func (o *BiomarkersList200ResponseRowsInner) GetNFdaCdxAssertions() int32`

GetNFdaCdxAssertions returns the NFdaCdxAssertions field if non-nil, zero value otherwise.

### GetNFdaCdxAssertionsOk

`func (o *BiomarkersList200ResponseRowsInner) GetNFdaCdxAssertionsOk() (*int32, bool)`

GetNFdaCdxAssertionsOk returns a tuple with the NFdaCdxAssertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNFdaCdxAssertions

`func (o *BiomarkersList200ResponseRowsInner) SetNFdaCdxAssertions(v int32)`

SetNFdaCdxAssertions sets NFdaCdxAssertions field to given value.


### GetNFdaBqpAssertions

`func (o *BiomarkersList200ResponseRowsInner) GetNFdaBqpAssertions() int32`

GetNFdaBqpAssertions returns the NFdaBqpAssertions field if non-nil, zero value otherwise.

### GetNFdaBqpAssertionsOk

`func (o *BiomarkersList200ResponseRowsInner) GetNFdaBqpAssertionsOk() (*int32, bool)`

GetNFdaBqpAssertionsOk returns a tuple with the NFdaBqpAssertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNFdaBqpAssertions

`func (o *BiomarkersList200ResponseRowsInner) SetNFdaBqpAssertions(v int32)`

SetNFdaBqpAssertions sets NFdaBqpAssertions field to given value.


### GetNCancerActionableCgi

`func (o *BiomarkersList200ResponseRowsInner) GetNCancerActionableCgi() int32`

GetNCancerActionableCgi returns the NCancerActionableCgi field if non-nil, zero value otherwise.

### GetNCancerActionableCgiOk

`func (o *BiomarkersList200ResponseRowsInner) GetNCancerActionableCgiOk() (*int32, bool)`

GetNCancerActionableCgiOk returns a tuple with the NCancerActionableCgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCancerActionableCgi

`func (o *BiomarkersList200ResponseRowsInner) SetNCancerActionableCgi(v int32)`

SetNCancerActionableCgi sets NCancerActionableCgi field to given value.


### GetNOtAssociation

`func (o *BiomarkersList200ResponseRowsInner) GetNOtAssociation() int32`

GetNOtAssociation returns the NOtAssociation field if non-nil, zero value otherwise.

### GetNOtAssociationOk

`func (o *BiomarkersList200ResponseRowsInner) GetNOtAssociationOk() (*int32, bool)`

GetNOtAssociationOk returns a tuple with the NOtAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNOtAssociation

`func (o *BiomarkersList200ResponseRowsInner) SetNOtAssociation(v int32)`

SetNOtAssociation sets NOtAssociation field to given value.


### GetNTrialMention

`func (o *BiomarkersList200ResponseRowsInner) GetNTrialMention() int32`

GetNTrialMention returns the NTrialMention field if non-nil, zero value otherwise.

### GetNTrialMentionOk

`func (o *BiomarkersList200ResponseRowsInner) GetNTrialMentionOk() (*int32, bool)`

GetNTrialMentionOk returns a tuple with the NTrialMention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrialMention

`func (o *BiomarkersList200ResponseRowsInner) SetNTrialMention(v int32)`

SetNTrialMention sets NTrialMention field to given value.


### GetNPublicationVariantMentions

`func (o *BiomarkersList200ResponseRowsInner) GetNPublicationVariantMentions() int32`

GetNPublicationVariantMentions returns the NPublicationVariantMentions field if non-nil, zero value otherwise.

### GetNPublicationVariantMentionsOk

`func (o *BiomarkersList200ResponseRowsInner) GetNPublicationVariantMentionsOk() (*int32, bool)`

GetNPublicationVariantMentionsOk returns a tuple with the NPublicationVariantMentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPublicationVariantMentions

`func (o *BiomarkersList200ResponseRowsInner) SetNPublicationVariantMentions(v int32)`

SetNPublicationVariantMentions sets NPublicationVariantMentions field to given value.


### GetKindLabel

`func (o *BiomarkersList200ResponseRowsInner) GetKindLabel() string`

GetKindLabel returns the KindLabel field if non-nil, zero value otherwise.

### GetKindLabelOk

`func (o *BiomarkersList200ResponseRowsInner) GetKindLabelOk() (*string, bool)`

GetKindLabelOk returns a tuple with the KindLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKindLabel

`func (o *BiomarkersList200ResponseRowsInner) SetKindLabel(v string)`

SetKindLabel sets KindLabel field to given value.


### GetEvidenceTier

`func (o *BiomarkersList200ResponseRowsInner) GetEvidenceTier() int32`

GetEvidenceTier returns the EvidenceTier field if non-nil, zero value otherwise.

### GetEvidenceTierOk

`func (o *BiomarkersList200ResponseRowsInner) GetEvidenceTierOk() (*int32, bool)`

GetEvidenceTierOk returns a tuple with the EvidenceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceTier

`func (o *BiomarkersList200ResponseRowsInner) SetEvidenceTier(v int32)`

SetEvidenceTier sets EvidenceTier field to given value.


### GetEvidenceTierLabel

`func (o *BiomarkersList200ResponseRowsInner) GetEvidenceTierLabel() string`

GetEvidenceTierLabel returns the EvidenceTierLabel field if non-nil, zero value otherwise.

### GetEvidenceTierLabelOk

`func (o *BiomarkersList200ResponseRowsInner) GetEvidenceTierLabelOk() (*string, bool)`

GetEvidenceTierLabelOk returns a tuple with the EvidenceTierLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceTierLabel

`func (o *BiomarkersList200ResponseRowsInner) SetEvidenceTierLabel(v string)`

SetEvidenceTierLabel sets EvidenceTierLabel field to given value.


### GetHasClinicalValidationEvidence

`func (o *BiomarkersList200ResponseRowsInner) GetHasClinicalValidationEvidence() bool`

GetHasClinicalValidationEvidence returns the HasClinicalValidationEvidence field if non-nil, zero value otherwise.

### GetHasClinicalValidationEvidenceOk

`func (o *BiomarkersList200ResponseRowsInner) GetHasClinicalValidationEvidenceOk() (*bool, bool)`

GetHasClinicalValidationEvidenceOk returns a tuple with the HasClinicalValidationEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasClinicalValidationEvidence

`func (o *BiomarkersList200ResponseRowsInner) SetHasClinicalValidationEvidence(v bool)`

SetHasClinicalValidationEvidence sets HasClinicalValidationEvidence field to given value.


### SetHasClinicalValidationEvidenceNil

`func (o *BiomarkersList200ResponseRowsInner) SetHasClinicalValidationEvidenceNil(b bool)`

 SetHasClinicalValidationEvidenceNil sets the value for HasClinicalValidationEvidence to be an explicit nil

### UnsetHasClinicalValidationEvidence
`func (o *BiomarkersList200ResponseRowsInner) UnsetHasClinicalValidationEvidence()`

UnsetHasClinicalValidationEvidence ensures that no value is present for HasClinicalValidationEvidence, not even an explicit nil
### GetHasAnalyticalValidationEvidence

`func (o *BiomarkersList200ResponseRowsInner) GetHasAnalyticalValidationEvidence() bool`

GetHasAnalyticalValidationEvidence returns the HasAnalyticalValidationEvidence field if non-nil, zero value otherwise.

### GetHasAnalyticalValidationEvidenceOk

`func (o *BiomarkersList200ResponseRowsInner) GetHasAnalyticalValidationEvidenceOk() (*bool, bool)`

GetHasAnalyticalValidationEvidenceOk returns a tuple with the HasAnalyticalValidationEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAnalyticalValidationEvidence

`func (o *BiomarkersList200ResponseRowsInner) SetHasAnalyticalValidationEvidence(v bool)`

SetHasAnalyticalValidationEvidence sets HasAnalyticalValidationEvidence field to given value.


### SetHasAnalyticalValidationEvidenceNil

`func (o *BiomarkersList200ResponseRowsInner) SetHasAnalyticalValidationEvidenceNil(b bool)`

 SetHasAnalyticalValidationEvidenceNil sets the value for HasAnalyticalValidationEvidence to be an explicit nil

### UnsetHasAnalyticalValidationEvidence
`func (o *BiomarkersList200ResponseRowsInner) UnsetHasAnalyticalValidationEvidence()`

UnsetHasAnalyticalValidationEvidence ensures that no value is present for HasAnalyticalValidationEvidence, not even an explicit nil
### GetNAssayProviderPapers

`func (o *BiomarkersList200ResponseRowsInner) GetNAssayProviderPapers() int32`

GetNAssayProviderPapers returns the NAssayProviderPapers field if non-nil, zero value otherwise.

### GetNAssayProviderPapersOk

`func (o *BiomarkersList200ResponseRowsInner) GetNAssayProviderPapersOk() (*int32, bool)`

GetNAssayProviderPapersOk returns a tuple with the NAssayProviderPapers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNAssayProviderPapers

`func (o *BiomarkersList200ResponseRowsInner) SetNAssayProviderPapers(v int32)`

SetNAssayProviderPapers sets NAssayProviderPapers field to given value.


### SetNAssayProviderPapersNil

`func (o *BiomarkersList200ResponseRowsInner) SetNAssayProviderPapersNil(b bool)`

 SetNAssayProviderPapersNil sets the value for NAssayProviderPapers to be an explicit nil

### UnsetNAssayProviderPapers
`func (o *BiomarkersList200ResponseRowsInner) UnsetNAssayProviderPapers()`

UnsetNAssayProviderPapers ensures that no value is present for NAssayProviderPapers, not even an explicit nil
### GetNAssayProviderSuppliers

`func (o *BiomarkersList200ResponseRowsInner) GetNAssayProviderSuppliers() int32`

GetNAssayProviderSuppliers returns the NAssayProviderSuppliers field if non-nil, zero value otherwise.

### GetNAssayProviderSuppliersOk

`func (o *BiomarkersList200ResponseRowsInner) GetNAssayProviderSuppliersOk() (*int32, bool)`

GetNAssayProviderSuppliersOk returns a tuple with the NAssayProviderSuppliers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNAssayProviderSuppliers

`func (o *BiomarkersList200ResponseRowsInner) SetNAssayProviderSuppliers(v int32)`

SetNAssayProviderSuppliers sets NAssayProviderSuppliers field to given value.


### SetNAssayProviderSuppliersNil

`func (o *BiomarkersList200ResponseRowsInner) SetNAssayProviderSuppliersNil(b bool)`

 SetNAssayProviderSuppliersNil sets the value for NAssayProviderSuppliers to be an explicit nil

### UnsetNAssayProviderSuppliers
`func (o *BiomarkersList200ResponseRowsInner) UnsetNAssayProviderSuppliers()`

UnsetNAssayProviderSuppliers ensures that no value is present for NAssayProviderSuppliers, not even an explicit nil
### GetNDiagnosticProviderSuppliers

`func (o *BiomarkersList200ResponseRowsInner) GetNDiagnosticProviderSuppliers() int32`

GetNDiagnosticProviderSuppliers returns the NDiagnosticProviderSuppliers field if non-nil, zero value otherwise.

### GetNDiagnosticProviderSuppliersOk

`func (o *BiomarkersList200ResponseRowsInner) GetNDiagnosticProviderSuppliersOk() (*int32, bool)`

GetNDiagnosticProviderSuppliersOk returns a tuple with the NDiagnosticProviderSuppliers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNDiagnosticProviderSuppliers

`func (o *BiomarkersList200ResponseRowsInner) SetNDiagnosticProviderSuppliers(v int32)`

SetNDiagnosticProviderSuppliers sets NDiagnosticProviderSuppliers field to given value.


### SetNDiagnosticProviderSuppliersNil

`func (o *BiomarkersList200ResponseRowsInner) SetNDiagnosticProviderSuppliersNil(b bool)`

 SetNDiagnosticProviderSuppliersNil sets the value for NDiagnosticProviderSuppliers to be an explicit nil

### UnsetNDiagnosticProviderSuppliers
`func (o *BiomarkersList200ResponseRowsInner) UnsetNDiagnosticProviderSuppliers()`

UnsetNDiagnosticProviderSuppliers ensures that no value is present for NDiagnosticProviderSuppliers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


