# DrugsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoleculeId** | Pointer to **NullableString** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] 
**MoleculeName** | Pointer to **NullableString** | The drug&#39;s display name, from its lead record. | [optional] 
**SubstanceName** | Pointer to **NullableString** | The preferred substance name in the FDA substance registry. | [optional] 
**InnName** | Pointer to **NullableString** | The International Nonproprietary Name (INN) where one is recorded. | [optional] 
**Unii** | Pointer to **NullableString** | The FDA Unique Ingredient Identifier (UNII) of the lead record. | [optional] 
**Cas** | Pointer to **NullableString** | The Chemical Abstracts Service (CAS) registry number of the lead record. | [optional] 
**ChemblId** | Pointer to **NullableString** | The ChEMBL compound id of the lead record. | [optional] 
**SubstanceType** | Pointer to **NullableString** | The registry&#39;s substance type (chemical, protein, mixture and so on). | [optional] 
**AtcCode** | Pointer to **NullableString** | The Anatomical Therapeutic Chemical (ATC) code. | [optional] 
**AtcClass** | Pointer to **NullableString** | The top-level ATC class. | [optional] 
**PharmClass** | Pointer to **NullableString** | The pharmacologic class the FDA registry gives the substance. | [optional] 
**IsBiologic** | Pointer to **NullableBool** | True when the substance is a biologic, by the registry, Drugs@FDA or the Purple Book. | [optional] 
**IsMarketed** | Pointer to **NullableBool** | True when the drug has a package in the US National Drug Code (NDC) directory. | [optional] 
**Withdrawn** | Pointer to **NullableBool** | True when ChEMBL records the drug as withdrawn somewhere. | [optional] 
**ApprovedAnywhere** | Pointer to **NullableBool** | True when the drug is approved in at least one jurisdiction. | [optional] 
**ApprovedIn** | Pointer to **[]string** | The regulators whose approval is recorded (FDA, EMA, Health Canada). | [optional] 
**FirstApprovalYear** | Pointer to **NullableInt32** | The earliest approval year ChEMBL records. | [optional] 
**MaxPhase** | Pointer to **NullableInt32** | The highest development phase reached across the folded records (4 is approved). | [optional] 
**NTrials** | Pointer to **NullableInt32** | Registered trials that study the drug or a record folded under it. | [optional] 
**NMembers** | Pointer to **NullableInt32** | Drug records folded into this drug, the lead included. | [optional] 
**MemberIds** | Pointer to **[]string** | The drug record ids folded into this drug, the lead included. The ids other datasets name drug_id. | [optional] 
**VariantCount** | Pointer to **NullableInt32** | Variant records folded under the molecule; null when the drug is not a molecule. | [optional] 
**Brands** | Pointer to **[]string** | The brand keys that resolve to the drug. | [optional] 
**BrandCount** | Pointer to **NullableInt32** | The number of brand keys. | [optional] 
**IsMolecule** | Pointer to **NullableBool** | True when the drug has a chemical-entities page. | [optional] 
**IsCombination** | Pointer to **NullableBool** | True when the drug is a combination product of two or more components. | [optional] 
**HasBoxedWarning** | Pointer to **NullableBool** | True when the drug&#39;s current label carries a boxed warning; null when no label is on record. | [optional] 
**Smiles** | Pointer to **NullableString** | The canonical SMILES string of the structure (ChEMBL). | [optional] 
**Inchikey** | Pointer to **NullableString** | The InChIKey of the structure (ChEMBL). | [optional] 
**NPatentMentions** | Pointer to **NullableInt32** | Patents that mention the drug in the SureChEMBL corpus. | [optional] 
**NPatentFamilies** | Pointer to **NullableInt32** | Distinct patent families among those mentions. | [optional] 
**FirstPatentMention** | Pointer to **NullableString** | Publication date of the earliest mentioning patent. | [optional] 
**LastPatentMention** | Pointer to **NullableString** | Publication date of the latest mentioning patent. | [optional] 

## Methods

### NewDrugsRow

`func NewDrugsRow() *DrugsRow`

NewDrugsRow instantiates a new DrugsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrugsRowWithDefaults

`func NewDrugsRowWithDefaults() *DrugsRow`

NewDrugsRowWithDefaults instantiates a new DrugsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoleculeId

`func (o *DrugsRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *DrugsRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *DrugsRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *DrugsRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *DrugsRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *DrugsRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetMoleculeName

`func (o *DrugsRow) GetMoleculeName() string`

GetMoleculeName returns the MoleculeName field if non-nil, zero value otherwise.

### GetMoleculeNameOk

`func (o *DrugsRow) GetMoleculeNameOk() (*string, bool)`

GetMoleculeNameOk returns a tuple with the MoleculeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeName

`func (o *DrugsRow) SetMoleculeName(v string)`

SetMoleculeName sets MoleculeName field to given value.

### HasMoleculeName

`func (o *DrugsRow) HasMoleculeName() bool`

HasMoleculeName returns a boolean if a field has been set.

### SetMoleculeNameNil

`func (o *DrugsRow) SetMoleculeNameNil(b bool)`

 SetMoleculeNameNil sets the value for MoleculeName to be an explicit nil

### UnsetMoleculeName
`func (o *DrugsRow) UnsetMoleculeName()`

UnsetMoleculeName ensures that no value is present for MoleculeName, not even an explicit nil
### GetSubstanceName

`func (o *DrugsRow) GetSubstanceName() string`

GetSubstanceName returns the SubstanceName field if non-nil, zero value otherwise.

### GetSubstanceNameOk

`func (o *DrugsRow) GetSubstanceNameOk() (*string, bool)`

GetSubstanceNameOk returns a tuple with the SubstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstanceName

`func (o *DrugsRow) SetSubstanceName(v string)`

SetSubstanceName sets SubstanceName field to given value.

### HasSubstanceName

`func (o *DrugsRow) HasSubstanceName() bool`

HasSubstanceName returns a boolean if a field has been set.

### SetSubstanceNameNil

`func (o *DrugsRow) SetSubstanceNameNil(b bool)`

 SetSubstanceNameNil sets the value for SubstanceName to be an explicit nil

### UnsetSubstanceName
`func (o *DrugsRow) UnsetSubstanceName()`

UnsetSubstanceName ensures that no value is present for SubstanceName, not even an explicit nil
### GetInnName

`func (o *DrugsRow) GetInnName() string`

GetInnName returns the InnName field if non-nil, zero value otherwise.

### GetInnNameOk

`func (o *DrugsRow) GetInnNameOk() (*string, bool)`

GetInnNameOk returns a tuple with the InnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnName

`func (o *DrugsRow) SetInnName(v string)`

SetInnName sets InnName field to given value.

### HasInnName

`func (o *DrugsRow) HasInnName() bool`

HasInnName returns a boolean if a field has been set.

### SetInnNameNil

`func (o *DrugsRow) SetInnNameNil(b bool)`

 SetInnNameNil sets the value for InnName to be an explicit nil

### UnsetInnName
`func (o *DrugsRow) UnsetInnName()`

UnsetInnName ensures that no value is present for InnName, not even an explicit nil
### GetUnii

`func (o *DrugsRow) GetUnii() string`

GetUnii returns the Unii field if non-nil, zero value otherwise.

### GetUniiOk

`func (o *DrugsRow) GetUniiOk() (*string, bool)`

GetUniiOk returns a tuple with the Unii field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnii

`func (o *DrugsRow) SetUnii(v string)`

SetUnii sets Unii field to given value.

### HasUnii

`func (o *DrugsRow) HasUnii() bool`

HasUnii returns a boolean if a field has been set.

### SetUniiNil

`func (o *DrugsRow) SetUniiNil(b bool)`

 SetUniiNil sets the value for Unii to be an explicit nil

### UnsetUnii
`func (o *DrugsRow) UnsetUnii()`

UnsetUnii ensures that no value is present for Unii, not even an explicit nil
### GetCas

`func (o *DrugsRow) GetCas() string`

GetCas returns the Cas field if non-nil, zero value otherwise.

### GetCasOk

`func (o *DrugsRow) GetCasOk() (*string, bool)`

GetCasOk returns a tuple with the Cas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCas

`func (o *DrugsRow) SetCas(v string)`

SetCas sets Cas field to given value.

### HasCas

`func (o *DrugsRow) HasCas() bool`

HasCas returns a boolean if a field has been set.

### SetCasNil

`func (o *DrugsRow) SetCasNil(b bool)`

 SetCasNil sets the value for Cas to be an explicit nil

### UnsetCas
`func (o *DrugsRow) UnsetCas()`

UnsetCas ensures that no value is present for Cas, not even an explicit nil
### GetChemblId

`func (o *DrugsRow) GetChemblId() string`

GetChemblId returns the ChemblId field if non-nil, zero value otherwise.

### GetChemblIdOk

`func (o *DrugsRow) GetChemblIdOk() (*string, bool)`

GetChemblIdOk returns a tuple with the ChemblId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChemblId

`func (o *DrugsRow) SetChemblId(v string)`

SetChemblId sets ChemblId field to given value.

### HasChemblId

`func (o *DrugsRow) HasChemblId() bool`

HasChemblId returns a boolean if a field has been set.

### SetChemblIdNil

`func (o *DrugsRow) SetChemblIdNil(b bool)`

 SetChemblIdNil sets the value for ChemblId to be an explicit nil

### UnsetChemblId
`func (o *DrugsRow) UnsetChemblId()`

UnsetChemblId ensures that no value is present for ChemblId, not even an explicit nil
### GetSubstanceType

`func (o *DrugsRow) GetSubstanceType() string`

GetSubstanceType returns the SubstanceType field if non-nil, zero value otherwise.

### GetSubstanceTypeOk

`func (o *DrugsRow) GetSubstanceTypeOk() (*string, bool)`

GetSubstanceTypeOk returns a tuple with the SubstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstanceType

`func (o *DrugsRow) SetSubstanceType(v string)`

SetSubstanceType sets SubstanceType field to given value.

### HasSubstanceType

`func (o *DrugsRow) HasSubstanceType() bool`

HasSubstanceType returns a boolean if a field has been set.

### SetSubstanceTypeNil

`func (o *DrugsRow) SetSubstanceTypeNil(b bool)`

 SetSubstanceTypeNil sets the value for SubstanceType to be an explicit nil

### UnsetSubstanceType
`func (o *DrugsRow) UnsetSubstanceType()`

UnsetSubstanceType ensures that no value is present for SubstanceType, not even an explicit nil
### GetAtcCode

`func (o *DrugsRow) GetAtcCode() string`

GetAtcCode returns the AtcCode field if non-nil, zero value otherwise.

### GetAtcCodeOk

`func (o *DrugsRow) GetAtcCodeOk() (*string, bool)`

GetAtcCodeOk returns a tuple with the AtcCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtcCode

`func (o *DrugsRow) SetAtcCode(v string)`

SetAtcCode sets AtcCode field to given value.

### HasAtcCode

`func (o *DrugsRow) HasAtcCode() bool`

HasAtcCode returns a boolean if a field has been set.

### SetAtcCodeNil

`func (o *DrugsRow) SetAtcCodeNil(b bool)`

 SetAtcCodeNil sets the value for AtcCode to be an explicit nil

### UnsetAtcCode
`func (o *DrugsRow) UnsetAtcCode()`

UnsetAtcCode ensures that no value is present for AtcCode, not even an explicit nil
### GetAtcClass

`func (o *DrugsRow) GetAtcClass() string`

GetAtcClass returns the AtcClass field if non-nil, zero value otherwise.

### GetAtcClassOk

`func (o *DrugsRow) GetAtcClassOk() (*string, bool)`

GetAtcClassOk returns a tuple with the AtcClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtcClass

`func (o *DrugsRow) SetAtcClass(v string)`

SetAtcClass sets AtcClass field to given value.

### HasAtcClass

`func (o *DrugsRow) HasAtcClass() bool`

HasAtcClass returns a boolean if a field has been set.

### SetAtcClassNil

`func (o *DrugsRow) SetAtcClassNil(b bool)`

 SetAtcClassNil sets the value for AtcClass to be an explicit nil

### UnsetAtcClass
`func (o *DrugsRow) UnsetAtcClass()`

UnsetAtcClass ensures that no value is present for AtcClass, not even an explicit nil
### GetPharmClass

`func (o *DrugsRow) GetPharmClass() string`

GetPharmClass returns the PharmClass field if non-nil, zero value otherwise.

### GetPharmClassOk

`func (o *DrugsRow) GetPharmClassOk() (*string, bool)`

GetPharmClassOk returns a tuple with the PharmClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPharmClass

`func (o *DrugsRow) SetPharmClass(v string)`

SetPharmClass sets PharmClass field to given value.

### HasPharmClass

`func (o *DrugsRow) HasPharmClass() bool`

HasPharmClass returns a boolean if a field has been set.

### SetPharmClassNil

`func (o *DrugsRow) SetPharmClassNil(b bool)`

 SetPharmClassNil sets the value for PharmClass to be an explicit nil

### UnsetPharmClass
`func (o *DrugsRow) UnsetPharmClass()`

UnsetPharmClass ensures that no value is present for PharmClass, not even an explicit nil
### GetIsBiologic

`func (o *DrugsRow) GetIsBiologic() bool`

GetIsBiologic returns the IsBiologic field if non-nil, zero value otherwise.

### GetIsBiologicOk

`func (o *DrugsRow) GetIsBiologicOk() (*bool, bool)`

GetIsBiologicOk returns a tuple with the IsBiologic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBiologic

`func (o *DrugsRow) SetIsBiologic(v bool)`

SetIsBiologic sets IsBiologic field to given value.

### HasIsBiologic

`func (o *DrugsRow) HasIsBiologic() bool`

HasIsBiologic returns a boolean if a field has been set.

### SetIsBiologicNil

`func (o *DrugsRow) SetIsBiologicNil(b bool)`

 SetIsBiologicNil sets the value for IsBiologic to be an explicit nil

### UnsetIsBiologic
`func (o *DrugsRow) UnsetIsBiologic()`

UnsetIsBiologic ensures that no value is present for IsBiologic, not even an explicit nil
### GetIsMarketed

`func (o *DrugsRow) GetIsMarketed() bool`

GetIsMarketed returns the IsMarketed field if non-nil, zero value otherwise.

### GetIsMarketedOk

`func (o *DrugsRow) GetIsMarketedOk() (*bool, bool)`

GetIsMarketedOk returns a tuple with the IsMarketed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarketed

`func (o *DrugsRow) SetIsMarketed(v bool)`

SetIsMarketed sets IsMarketed field to given value.

### HasIsMarketed

`func (o *DrugsRow) HasIsMarketed() bool`

HasIsMarketed returns a boolean if a field has been set.

### SetIsMarketedNil

`func (o *DrugsRow) SetIsMarketedNil(b bool)`

 SetIsMarketedNil sets the value for IsMarketed to be an explicit nil

### UnsetIsMarketed
`func (o *DrugsRow) UnsetIsMarketed()`

UnsetIsMarketed ensures that no value is present for IsMarketed, not even an explicit nil
### GetWithdrawn

`func (o *DrugsRow) GetWithdrawn() bool`

GetWithdrawn returns the Withdrawn field if non-nil, zero value otherwise.

### GetWithdrawnOk

`func (o *DrugsRow) GetWithdrawnOk() (*bool, bool)`

GetWithdrawnOk returns a tuple with the Withdrawn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawn

`func (o *DrugsRow) SetWithdrawn(v bool)`

SetWithdrawn sets Withdrawn field to given value.

### HasWithdrawn

`func (o *DrugsRow) HasWithdrawn() bool`

HasWithdrawn returns a boolean if a field has been set.

### SetWithdrawnNil

`func (o *DrugsRow) SetWithdrawnNil(b bool)`

 SetWithdrawnNil sets the value for Withdrawn to be an explicit nil

### UnsetWithdrawn
`func (o *DrugsRow) UnsetWithdrawn()`

UnsetWithdrawn ensures that no value is present for Withdrawn, not even an explicit nil
### GetApprovedAnywhere

`func (o *DrugsRow) GetApprovedAnywhere() bool`

GetApprovedAnywhere returns the ApprovedAnywhere field if non-nil, zero value otherwise.

### GetApprovedAnywhereOk

`func (o *DrugsRow) GetApprovedAnywhereOk() (*bool, bool)`

GetApprovedAnywhereOk returns a tuple with the ApprovedAnywhere field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAnywhere

`func (o *DrugsRow) SetApprovedAnywhere(v bool)`

SetApprovedAnywhere sets ApprovedAnywhere field to given value.

### HasApprovedAnywhere

`func (o *DrugsRow) HasApprovedAnywhere() bool`

HasApprovedAnywhere returns a boolean if a field has been set.

### SetApprovedAnywhereNil

`func (o *DrugsRow) SetApprovedAnywhereNil(b bool)`

 SetApprovedAnywhereNil sets the value for ApprovedAnywhere to be an explicit nil

### UnsetApprovedAnywhere
`func (o *DrugsRow) UnsetApprovedAnywhere()`

UnsetApprovedAnywhere ensures that no value is present for ApprovedAnywhere, not even an explicit nil
### GetApprovedIn

`func (o *DrugsRow) GetApprovedIn() []string`

GetApprovedIn returns the ApprovedIn field if non-nil, zero value otherwise.

### GetApprovedInOk

`func (o *DrugsRow) GetApprovedInOk() (*[]string, bool)`

GetApprovedInOk returns a tuple with the ApprovedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedIn

`func (o *DrugsRow) SetApprovedIn(v []string)`

SetApprovedIn sets ApprovedIn field to given value.

### HasApprovedIn

`func (o *DrugsRow) HasApprovedIn() bool`

HasApprovedIn returns a boolean if a field has been set.

### SetApprovedInNil

`func (o *DrugsRow) SetApprovedInNil(b bool)`

 SetApprovedInNil sets the value for ApprovedIn to be an explicit nil

### UnsetApprovedIn
`func (o *DrugsRow) UnsetApprovedIn()`

UnsetApprovedIn ensures that no value is present for ApprovedIn, not even an explicit nil
### GetFirstApprovalYear

`func (o *DrugsRow) GetFirstApprovalYear() int32`

GetFirstApprovalYear returns the FirstApprovalYear field if non-nil, zero value otherwise.

### GetFirstApprovalYearOk

`func (o *DrugsRow) GetFirstApprovalYearOk() (*int32, bool)`

GetFirstApprovalYearOk returns a tuple with the FirstApprovalYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstApprovalYear

`func (o *DrugsRow) SetFirstApprovalYear(v int32)`

SetFirstApprovalYear sets FirstApprovalYear field to given value.

### HasFirstApprovalYear

`func (o *DrugsRow) HasFirstApprovalYear() bool`

HasFirstApprovalYear returns a boolean if a field has been set.

### SetFirstApprovalYearNil

`func (o *DrugsRow) SetFirstApprovalYearNil(b bool)`

 SetFirstApprovalYearNil sets the value for FirstApprovalYear to be an explicit nil

### UnsetFirstApprovalYear
`func (o *DrugsRow) UnsetFirstApprovalYear()`

UnsetFirstApprovalYear ensures that no value is present for FirstApprovalYear, not even an explicit nil
### GetMaxPhase

`func (o *DrugsRow) GetMaxPhase() int32`

GetMaxPhase returns the MaxPhase field if non-nil, zero value otherwise.

### GetMaxPhaseOk

`func (o *DrugsRow) GetMaxPhaseOk() (*int32, bool)`

GetMaxPhaseOk returns a tuple with the MaxPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPhase

`func (o *DrugsRow) SetMaxPhase(v int32)`

SetMaxPhase sets MaxPhase field to given value.

### HasMaxPhase

`func (o *DrugsRow) HasMaxPhase() bool`

HasMaxPhase returns a boolean if a field has been set.

### SetMaxPhaseNil

`func (o *DrugsRow) SetMaxPhaseNil(b bool)`

 SetMaxPhaseNil sets the value for MaxPhase to be an explicit nil

### UnsetMaxPhase
`func (o *DrugsRow) UnsetMaxPhase()`

UnsetMaxPhase ensures that no value is present for MaxPhase, not even an explicit nil
### GetNTrials

`func (o *DrugsRow) GetNTrials() int32`

GetNTrials returns the NTrials field if non-nil, zero value otherwise.

### GetNTrialsOk

`func (o *DrugsRow) GetNTrialsOk() (*int32, bool)`

GetNTrialsOk returns a tuple with the NTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTrials

`func (o *DrugsRow) SetNTrials(v int32)`

SetNTrials sets NTrials field to given value.

### HasNTrials

`func (o *DrugsRow) HasNTrials() bool`

HasNTrials returns a boolean if a field has been set.

### SetNTrialsNil

`func (o *DrugsRow) SetNTrialsNil(b bool)`

 SetNTrialsNil sets the value for NTrials to be an explicit nil

### UnsetNTrials
`func (o *DrugsRow) UnsetNTrials()`

UnsetNTrials ensures that no value is present for NTrials, not even an explicit nil
### GetNMembers

`func (o *DrugsRow) GetNMembers() int32`

GetNMembers returns the NMembers field if non-nil, zero value otherwise.

### GetNMembersOk

`func (o *DrugsRow) GetNMembersOk() (*int32, bool)`

GetNMembersOk returns a tuple with the NMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNMembers

`func (o *DrugsRow) SetNMembers(v int32)`

SetNMembers sets NMembers field to given value.

### HasNMembers

`func (o *DrugsRow) HasNMembers() bool`

HasNMembers returns a boolean if a field has been set.

### SetNMembersNil

`func (o *DrugsRow) SetNMembersNil(b bool)`

 SetNMembersNil sets the value for NMembers to be an explicit nil

### UnsetNMembers
`func (o *DrugsRow) UnsetNMembers()`

UnsetNMembers ensures that no value is present for NMembers, not even an explicit nil
### GetMemberIds

`func (o *DrugsRow) GetMemberIds() []string`

GetMemberIds returns the MemberIds field if non-nil, zero value otherwise.

### GetMemberIdsOk

`func (o *DrugsRow) GetMemberIdsOk() (*[]string, bool)`

GetMemberIdsOk returns a tuple with the MemberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberIds

`func (o *DrugsRow) SetMemberIds(v []string)`

SetMemberIds sets MemberIds field to given value.

### HasMemberIds

`func (o *DrugsRow) HasMemberIds() bool`

HasMemberIds returns a boolean if a field has been set.

### SetMemberIdsNil

`func (o *DrugsRow) SetMemberIdsNil(b bool)`

 SetMemberIdsNil sets the value for MemberIds to be an explicit nil

### UnsetMemberIds
`func (o *DrugsRow) UnsetMemberIds()`

UnsetMemberIds ensures that no value is present for MemberIds, not even an explicit nil
### GetVariantCount

`func (o *DrugsRow) GetVariantCount() int32`

GetVariantCount returns the VariantCount field if non-nil, zero value otherwise.

### GetVariantCountOk

`func (o *DrugsRow) GetVariantCountOk() (*int32, bool)`

GetVariantCountOk returns a tuple with the VariantCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantCount

`func (o *DrugsRow) SetVariantCount(v int32)`

SetVariantCount sets VariantCount field to given value.

### HasVariantCount

`func (o *DrugsRow) HasVariantCount() bool`

HasVariantCount returns a boolean if a field has been set.

### SetVariantCountNil

`func (o *DrugsRow) SetVariantCountNil(b bool)`

 SetVariantCountNil sets the value for VariantCount to be an explicit nil

### UnsetVariantCount
`func (o *DrugsRow) UnsetVariantCount()`

UnsetVariantCount ensures that no value is present for VariantCount, not even an explicit nil
### GetBrands

`func (o *DrugsRow) GetBrands() []string`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *DrugsRow) GetBrandsOk() (*[]string, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *DrugsRow) SetBrands(v []string)`

SetBrands sets Brands field to given value.

### HasBrands

`func (o *DrugsRow) HasBrands() bool`

HasBrands returns a boolean if a field has been set.

### SetBrandsNil

`func (o *DrugsRow) SetBrandsNil(b bool)`

 SetBrandsNil sets the value for Brands to be an explicit nil

### UnsetBrands
`func (o *DrugsRow) UnsetBrands()`

UnsetBrands ensures that no value is present for Brands, not even an explicit nil
### GetBrandCount

`func (o *DrugsRow) GetBrandCount() int32`

GetBrandCount returns the BrandCount field if non-nil, zero value otherwise.

### GetBrandCountOk

`func (o *DrugsRow) GetBrandCountOk() (*int32, bool)`

GetBrandCountOk returns a tuple with the BrandCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandCount

`func (o *DrugsRow) SetBrandCount(v int32)`

SetBrandCount sets BrandCount field to given value.

### HasBrandCount

`func (o *DrugsRow) HasBrandCount() bool`

HasBrandCount returns a boolean if a field has been set.

### SetBrandCountNil

`func (o *DrugsRow) SetBrandCountNil(b bool)`

 SetBrandCountNil sets the value for BrandCount to be an explicit nil

### UnsetBrandCount
`func (o *DrugsRow) UnsetBrandCount()`

UnsetBrandCount ensures that no value is present for BrandCount, not even an explicit nil
### GetIsMolecule

`func (o *DrugsRow) GetIsMolecule() bool`

GetIsMolecule returns the IsMolecule field if non-nil, zero value otherwise.

### GetIsMoleculeOk

`func (o *DrugsRow) GetIsMoleculeOk() (*bool, bool)`

GetIsMoleculeOk returns a tuple with the IsMolecule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMolecule

`func (o *DrugsRow) SetIsMolecule(v bool)`

SetIsMolecule sets IsMolecule field to given value.

### HasIsMolecule

`func (o *DrugsRow) HasIsMolecule() bool`

HasIsMolecule returns a boolean if a field has been set.

### SetIsMoleculeNil

`func (o *DrugsRow) SetIsMoleculeNil(b bool)`

 SetIsMoleculeNil sets the value for IsMolecule to be an explicit nil

### UnsetIsMolecule
`func (o *DrugsRow) UnsetIsMolecule()`

UnsetIsMolecule ensures that no value is present for IsMolecule, not even an explicit nil
### GetIsCombination

`func (o *DrugsRow) GetIsCombination() bool`

GetIsCombination returns the IsCombination field if non-nil, zero value otherwise.

### GetIsCombinationOk

`func (o *DrugsRow) GetIsCombinationOk() (*bool, bool)`

GetIsCombinationOk returns a tuple with the IsCombination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCombination

`func (o *DrugsRow) SetIsCombination(v bool)`

SetIsCombination sets IsCombination field to given value.

### HasIsCombination

`func (o *DrugsRow) HasIsCombination() bool`

HasIsCombination returns a boolean if a field has been set.

### SetIsCombinationNil

`func (o *DrugsRow) SetIsCombinationNil(b bool)`

 SetIsCombinationNil sets the value for IsCombination to be an explicit nil

### UnsetIsCombination
`func (o *DrugsRow) UnsetIsCombination()`

UnsetIsCombination ensures that no value is present for IsCombination, not even an explicit nil
### GetHasBoxedWarning

`func (o *DrugsRow) GetHasBoxedWarning() bool`

GetHasBoxedWarning returns the HasBoxedWarning field if non-nil, zero value otherwise.

### GetHasBoxedWarningOk

`func (o *DrugsRow) GetHasBoxedWarningOk() (*bool, bool)`

GetHasBoxedWarningOk returns a tuple with the HasBoxedWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBoxedWarning

`func (o *DrugsRow) SetHasBoxedWarning(v bool)`

SetHasBoxedWarning sets HasBoxedWarning field to given value.

### HasHasBoxedWarning

`func (o *DrugsRow) HasHasBoxedWarning() bool`

HasHasBoxedWarning returns a boolean if a field has been set.

### SetHasBoxedWarningNil

`func (o *DrugsRow) SetHasBoxedWarningNil(b bool)`

 SetHasBoxedWarningNil sets the value for HasBoxedWarning to be an explicit nil

### UnsetHasBoxedWarning
`func (o *DrugsRow) UnsetHasBoxedWarning()`

UnsetHasBoxedWarning ensures that no value is present for HasBoxedWarning, not even an explicit nil
### GetSmiles

`func (o *DrugsRow) GetSmiles() string`

GetSmiles returns the Smiles field if non-nil, zero value otherwise.

### GetSmilesOk

`func (o *DrugsRow) GetSmilesOk() (*string, bool)`

GetSmilesOk returns a tuple with the Smiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmiles

`func (o *DrugsRow) SetSmiles(v string)`

SetSmiles sets Smiles field to given value.

### HasSmiles

`func (o *DrugsRow) HasSmiles() bool`

HasSmiles returns a boolean if a field has been set.

### SetSmilesNil

`func (o *DrugsRow) SetSmilesNil(b bool)`

 SetSmilesNil sets the value for Smiles to be an explicit nil

### UnsetSmiles
`func (o *DrugsRow) UnsetSmiles()`

UnsetSmiles ensures that no value is present for Smiles, not even an explicit nil
### GetInchikey

`func (o *DrugsRow) GetInchikey() string`

GetInchikey returns the Inchikey field if non-nil, zero value otherwise.

### GetInchikeyOk

`func (o *DrugsRow) GetInchikeyOk() (*string, bool)`

GetInchikeyOk returns a tuple with the Inchikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInchikey

`func (o *DrugsRow) SetInchikey(v string)`

SetInchikey sets Inchikey field to given value.

### HasInchikey

`func (o *DrugsRow) HasInchikey() bool`

HasInchikey returns a boolean if a field has been set.

### SetInchikeyNil

`func (o *DrugsRow) SetInchikeyNil(b bool)`

 SetInchikeyNil sets the value for Inchikey to be an explicit nil

### UnsetInchikey
`func (o *DrugsRow) UnsetInchikey()`

UnsetInchikey ensures that no value is present for Inchikey, not even an explicit nil
### GetNPatentMentions

`func (o *DrugsRow) GetNPatentMentions() int32`

GetNPatentMentions returns the NPatentMentions field if non-nil, zero value otherwise.

### GetNPatentMentionsOk

`func (o *DrugsRow) GetNPatentMentionsOk() (*int32, bool)`

GetNPatentMentionsOk returns a tuple with the NPatentMentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPatentMentions

`func (o *DrugsRow) SetNPatentMentions(v int32)`

SetNPatentMentions sets NPatentMentions field to given value.

### HasNPatentMentions

`func (o *DrugsRow) HasNPatentMentions() bool`

HasNPatentMentions returns a boolean if a field has been set.

### SetNPatentMentionsNil

`func (o *DrugsRow) SetNPatentMentionsNil(b bool)`

 SetNPatentMentionsNil sets the value for NPatentMentions to be an explicit nil

### UnsetNPatentMentions
`func (o *DrugsRow) UnsetNPatentMentions()`

UnsetNPatentMentions ensures that no value is present for NPatentMentions, not even an explicit nil
### GetNPatentFamilies

`func (o *DrugsRow) GetNPatentFamilies() int32`

GetNPatentFamilies returns the NPatentFamilies field if non-nil, zero value otherwise.

### GetNPatentFamiliesOk

`func (o *DrugsRow) GetNPatentFamiliesOk() (*int32, bool)`

GetNPatentFamiliesOk returns a tuple with the NPatentFamilies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPatentFamilies

`func (o *DrugsRow) SetNPatentFamilies(v int32)`

SetNPatentFamilies sets NPatentFamilies field to given value.

### HasNPatentFamilies

`func (o *DrugsRow) HasNPatentFamilies() bool`

HasNPatentFamilies returns a boolean if a field has been set.

### SetNPatentFamiliesNil

`func (o *DrugsRow) SetNPatentFamiliesNil(b bool)`

 SetNPatentFamiliesNil sets the value for NPatentFamilies to be an explicit nil

### UnsetNPatentFamilies
`func (o *DrugsRow) UnsetNPatentFamilies()`

UnsetNPatentFamilies ensures that no value is present for NPatentFamilies, not even an explicit nil
### GetFirstPatentMention

`func (o *DrugsRow) GetFirstPatentMention() string`

GetFirstPatentMention returns the FirstPatentMention field if non-nil, zero value otherwise.

### GetFirstPatentMentionOk

`func (o *DrugsRow) GetFirstPatentMentionOk() (*string, bool)`

GetFirstPatentMentionOk returns a tuple with the FirstPatentMention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPatentMention

`func (o *DrugsRow) SetFirstPatentMention(v string)`

SetFirstPatentMention sets FirstPatentMention field to given value.

### HasFirstPatentMention

`func (o *DrugsRow) HasFirstPatentMention() bool`

HasFirstPatentMention returns a boolean if a field has been set.

### SetFirstPatentMentionNil

`func (o *DrugsRow) SetFirstPatentMentionNil(b bool)`

 SetFirstPatentMentionNil sets the value for FirstPatentMention to be an explicit nil

### UnsetFirstPatentMention
`func (o *DrugsRow) UnsetFirstPatentMention()`

UnsetFirstPatentMention ensures that no value is present for FirstPatentMention, not even an explicit nil
### GetLastPatentMention

`func (o *DrugsRow) GetLastPatentMention() string`

GetLastPatentMention returns the LastPatentMention field if non-nil, zero value otherwise.

### GetLastPatentMentionOk

`func (o *DrugsRow) GetLastPatentMentionOk() (*string, bool)`

GetLastPatentMentionOk returns a tuple with the LastPatentMention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPatentMention

`func (o *DrugsRow) SetLastPatentMention(v string)`

SetLastPatentMention sets LastPatentMention field to given value.

### HasLastPatentMention

`func (o *DrugsRow) HasLastPatentMention() bool`

HasLastPatentMention returns a boolean if a field has been set.

### SetLastPatentMentionNil

`func (o *DrugsRow) SetLastPatentMentionNil(b bool)`

 SetLastPatentMentionNil sets the value for LastPatentMention to be an explicit nil

### UnsetLastPatentMention
`func (o *DrugsRow) UnsetLastPatentMention()`

UnsetLastPatentMention ensures that no value is present for LastPatentMention, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


