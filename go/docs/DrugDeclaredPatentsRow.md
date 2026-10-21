# DrugDeclaredPatentsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoleculeId** | Pointer to **NullableString** | The molecule key: the resolved drug a brand, a salt form or an investigational arm resolves to, and the entry key for the drugs object type. A member id (a salt or a folded trial arm) is not a value here; it redirects to its molecule. | [optional] 
**PatentId** | Pointer to **NullableString** | The patent id in the patent dimension. | [optional] 
**AssigneeSequence** | Pointer to **NullableInt32** | The assignee&#39;s order on the patent; -1 when there is none. | [optional] 
**PatentNumber** | Pointer to **NullableString** | The patent number. | [optional] 
**PatentTitle** | Pointer to **NullableString** | The patent title. | [optional] 
**PatentDate** | Pointer to **NullableString** | The date the patent was granted. | [optional] 
**DeclaredExpiry** | Pointer to **NullableString** | The expiry declared for the drug. | [optional] 
**ApplNos** | Pointer to **[]string** | The drug application numbers the patent is declared against. | [optional] 
**LinkBasis** | Pointer to **NullableString** | The record that ties the patent to the drug (orange_book or uspto_pte). | [optional] 
**LinkBases** | Pointer to **[]string** | Every record that ties the patent to the drug. | [optional] 
**HasPedExtension** | Pointer to **NullableBool** | True when the Orange Book lists a pediatric exclusivity extension. | [optional] 
**AssigneeOrganization** | Pointer to **NullableString** | The assignee as filed. | [optional] 
**AssigneeOrgId** | Pointer to **NullableString** | Legal-entity level: the assignee&#39;s organization record (null when unresolved). | [optional] 
**AssigneeCompanyId** | Pointer to **NullableString** | Company level: the assignee&#39;s company (null when unresolved). | [optional] 
**AssigneeCompanyName** | Pointer to **NullableString** | Display name of the assignee&#39;s company. | [optional] 
**AssigneeGroupId** | Pointer to **NullableString** | Group level: the assignee&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**AssigneeGroupName** | Pointer to **NullableString** | Display name of the assignee&#39;s group. | [optional] 
**OrgInSpine** | Pointer to **NullableBool** | True when the assignee is an organization record in the spine. | [optional] 
**ResolutionMethod** | Pointer to **NullableString** | How the assignee was matched to an organization. | [optional] 
**MatchConfidence** | Pointer to **NullableString** | The match confidence (high, med, low). | [optional] 
**OwnerBasis** | Pointer to **NullableString** | Why the assignee is, or is not, on record. | [optional] 

## Methods

### NewDrugDeclaredPatentsRow

`func NewDrugDeclaredPatentsRow() *DrugDeclaredPatentsRow`

NewDrugDeclaredPatentsRow instantiates a new DrugDeclaredPatentsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrugDeclaredPatentsRowWithDefaults

`func NewDrugDeclaredPatentsRowWithDefaults() *DrugDeclaredPatentsRow`

NewDrugDeclaredPatentsRowWithDefaults instantiates a new DrugDeclaredPatentsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoleculeId

`func (o *DrugDeclaredPatentsRow) GetMoleculeId() string`

GetMoleculeId returns the MoleculeId field if non-nil, zero value otherwise.

### GetMoleculeIdOk

`func (o *DrugDeclaredPatentsRow) GetMoleculeIdOk() (*string, bool)`

GetMoleculeIdOk returns a tuple with the MoleculeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeId

`func (o *DrugDeclaredPatentsRow) SetMoleculeId(v string)`

SetMoleculeId sets MoleculeId field to given value.

### HasMoleculeId

`func (o *DrugDeclaredPatentsRow) HasMoleculeId() bool`

HasMoleculeId returns a boolean if a field has been set.

### SetMoleculeIdNil

`func (o *DrugDeclaredPatentsRow) SetMoleculeIdNil(b bool)`

 SetMoleculeIdNil sets the value for MoleculeId to be an explicit nil

### UnsetMoleculeId
`func (o *DrugDeclaredPatentsRow) UnsetMoleculeId()`

UnsetMoleculeId ensures that no value is present for MoleculeId, not even an explicit nil
### GetPatentId

`func (o *DrugDeclaredPatentsRow) GetPatentId() string`

GetPatentId returns the PatentId field if non-nil, zero value otherwise.

### GetPatentIdOk

`func (o *DrugDeclaredPatentsRow) GetPatentIdOk() (*string, bool)`

GetPatentIdOk returns a tuple with the PatentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatentId

`func (o *DrugDeclaredPatentsRow) SetPatentId(v string)`

SetPatentId sets PatentId field to given value.

### HasPatentId

`func (o *DrugDeclaredPatentsRow) HasPatentId() bool`

HasPatentId returns a boolean if a field has been set.

### SetPatentIdNil

`func (o *DrugDeclaredPatentsRow) SetPatentIdNil(b bool)`

 SetPatentIdNil sets the value for PatentId to be an explicit nil

### UnsetPatentId
`func (o *DrugDeclaredPatentsRow) UnsetPatentId()`

UnsetPatentId ensures that no value is present for PatentId, not even an explicit nil
### GetAssigneeSequence

`func (o *DrugDeclaredPatentsRow) GetAssigneeSequence() int32`

GetAssigneeSequence returns the AssigneeSequence field if non-nil, zero value otherwise.

### GetAssigneeSequenceOk

`func (o *DrugDeclaredPatentsRow) GetAssigneeSequenceOk() (*int32, bool)`

GetAssigneeSequenceOk returns a tuple with the AssigneeSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeSequence

`func (o *DrugDeclaredPatentsRow) SetAssigneeSequence(v int32)`

SetAssigneeSequence sets AssigneeSequence field to given value.

### HasAssigneeSequence

`func (o *DrugDeclaredPatentsRow) HasAssigneeSequence() bool`

HasAssigneeSequence returns a boolean if a field has been set.

### SetAssigneeSequenceNil

`func (o *DrugDeclaredPatentsRow) SetAssigneeSequenceNil(b bool)`

 SetAssigneeSequenceNil sets the value for AssigneeSequence to be an explicit nil

### UnsetAssigneeSequence
`func (o *DrugDeclaredPatentsRow) UnsetAssigneeSequence()`

UnsetAssigneeSequence ensures that no value is present for AssigneeSequence, not even an explicit nil
### GetPatentNumber

`func (o *DrugDeclaredPatentsRow) GetPatentNumber() string`

GetPatentNumber returns the PatentNumber field if non-nil, zero value otherwise.

### GetPatentNumberOk

`func (o *DrugDeclaredPatentsRow) GetPatentNumberOk() (*string, bool)`

GetPatentNumberOk returns a tuple with the PatentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatentNumber

`func (o *DrugDeclaredPatentsRow) SetPatentNumber(v string)`

SetPatentNumber sets PatentNumber field to given value.

### HasPatentNumber

`func (o *DrugDeclaredPatentsRow) HasPatentNumber() bool`

HasPatentNumber returns a boolean if a field has been set.

### SetPatentNumberNil

`func (o *DrugDeclaredPatentsRow) SetPatentNumberNil(b bool)`

 SetPatentNumberNil sets the value for PatentNumber to be an explicit nil

### UnsetPatentNumber
`func (o *DrugDeclaredPatentsRow) UnsetPatentNumber()`

UnsetPatentNumber ensures that no value is present for PatentNumber, not even an explicit nil
### GetPatentTitle

`func (o *DrugDeclaredPatentsRow) GetPatentTitle() string`

GetPatentTitle returns the PatentTitle field if non-nil, zero value otherwise.

### GetPatentTitleOk

`func (o *DrugDeclaredPatentsRow) GetPatentTitleOk() (*string, bool)`

GetPatentTitleOk returns a tuple with the PatentTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatentTitle

`func (o *DrugDeclaredPatentsRow) SetPatentTitle(v string)`

SetPatentTitle sets PatentTitle field to given value.

### HasPatentTitle

`func (o *DrugDeclaredPatentsRow) HasPatentTitle() bool`

HasPatentTitle returns a boolean if a field has been set.

### SetPatentTitleNil

`func (o *DrugDeclaredPatentsRow) SetPatentTitleNil(b bool)`

 SetPatentTitleNil sets the value for PatentTitle to be an explicit nil

### UnsetPatentTitle
`func (o *DrugDeclaredPatentsRow) UnsetPatentTitle()`

UnsetPatentTitle ensures that no value is present for PatentTitle, not even an explicit nil
### GetPatentDate

`func (o *DrugDeclaredPatentsRow) GetPatentDate() string`

GetPatentDate returns the PatentDate field if non-nil, zero value otherwise.

### GetPatentDateOk

`func (o *DrugDeclaredPatentsRow) GetPatentDateOk() (*string, bool)`

GetPatentDateOk returns a tuple with the PatentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatentDate

`func (o *DrugDeclaredPatentsRow) SetPatentDate(v string)`

SetPatentDate sets PatentDate field to given value.

### HasPatentDate

`func (o *DrugDeclaredPatentsRow) HasPatentDate() bool`

HasPatentDate returns a boolean if a field has been set.

### SetPatentDateNil

`func (o *DrugDeclaredPatentsRow) SetPatentDateNil(b bool)`

 SetPatentDateNil sets the value for PatentDate to be an explicit nil

### UnsetPatentDate
`func (o *DrugDeclaredPatentsRow) UnsetPatentDate()`

UnsetPatentDate ensures that no value is present for PatentDate, not even an explicit nil
### GetDeclaredExpiry

`func (o *DrugDeclaredPatentsRow) GetDeclaredExpiry() string`

GetDeclaredExpiry returns the DeclaredExpiry field if non-nil, zero value otherwise.

### GetDeclaredExpiryOk

`func (o *DrugDeclaredPatentsRow) GetDeclaredExpiryOk() (*string, bool)`

GetDeclaredExpiryOk returns a tuple with the DeclaredExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredExpiry

`func (o *DrugDeclaredPatentsRow) SetDeclaredExpiry(v string)`

SetDeclaredExpiry sets DeclaredExpiry field to given value.

### HasDeclaredExpiry

`func (o *DrugDeclaredPatentsRow) HasDeclaredExpiry() bool`

HasDeclaredExpiry returns a boolean if a field has been set.

### SetDeclaredExpiryNil

`func (o *DrugDeclaredPatentsRow) SetDeclaredExpiryNil(b bool)`

 SetDeclaredExpiryNil sets the value for DeclaredExpiry to be an explicit nil

### UnsetDeclaredExpiry
`func (o *DrugDeclaredPatentsRow) UnsetDeclaredExpiry()`

UnsetDeclaredExpiry ensures that no value is present for DeclaredExpiry, not even an explicit nil
### GetApplNos

`func (o *DrugDeclaredPatentsRow) GetApplNos() []string`

GetApplNos returns the ApplNos field if non-nil, zero value otherwise.

### GetApplNosOk

`func (o *DrugDeclaredPatentsRow) GetApplNosOk() (*[]string, bool)`

GetApplNosOk returns a tuple with the ApplNos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplNos

`func (o *DrugDeclaredPatentsRow) SetApplNos(v []string)`

SetApplNos sets ApplNos field to given value.

### HasApplNos

`func (o *DrugDeclaredPatentsRow) HasApplNos() bool`

HasApplNos returns a boolean if a field has been set.

### SetApplNosNil

`func (o *DrugDeclaredPatentsRow) SetApplNosNil(b bool)`

 SetApplNosNil sets the value for ApplNos to be an explicit nil

### UnsetApplNos
`func (o *DrugDeclaredPatentsRow) UnsetApplNos()`

UnsetApplNos ensures that no value is present for ApplNos, not even an explicit nil
### GetLinkBasis

`func (o *DrugDeclaredPatentsRow) GetLinkBasis() string`

GetLinkBasis returns the LinkBasis field if non-nil, zero value otherwise.

### GetLinkBasisOk

`func (o *DrugDeclaredPatentsRow) GetLinkBasisOk() (*string, bool)`

GetLinkBasisOk returns a tuple with the LinkBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkBasis

`func (o *DrugDeclaredPatentsRow) SetLinkBasis(v string)`

SetLinkBasis sets LinkBasis field to given value.

### HasLinkBasis

`func (o *DrugDeclaredPatentsRow) HasLinkBasis() bool`

HasLinkBasis returns a boolean if a field has been set.

### SetLinkBasisNil

`func (o *DrugDeclaredPatentsRow) SetLinkBasisNil(b bool)`

 SetLinkBasisNil sets the value for LinkBasis to be an explicit nil

### UnsetLinkBasis
`func (o *DrugDeclaredPatentsRow) UnsetLinkBasis()`

UnsetLinkBasis ensures that no value is present for LinkBasis, not even an explicit nil
### GetLinkBases

`func (o *DrugDeclaredPatentsRow) GetLinkBases() []string`

GetLinkBases returns the LinkBases field if non-nil, zero value otherwise.

### GetLinkBasesOk

`func (o *DrugDeclaredPatentsRow) GetLinkBasesOk() (*[]string, bool)`

GetLinkBasesOk returns a tuple with the LinkBases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkBases

`func (o *DrugDeclaredPatentsRow) SetLinkBases(v []string)`

SetLinkBases sets LinkBases field to given value.

### HasLinkBases

`func (o *DrugDeclaredPatentsRow) HasLinkBases() bool`

HasLinkBases returns a boolean if a field has been set.

### SetLinkBasesNil

`func (o *DrugDeclaredPatentsRow) SetLinkBasesNil(b bool)`

 SetLinkBasesNil sets the value for LinkBases to be an explicit nil

### UnsetLinkBases
`func (o *DrugDeclaredPatentsRow) UnsetLinkBases()`

UnsetLinkBases ensures that no value is present for LinkBases, not even an explicit nil
### GetHasPedExtension

`func (o *DrugDeclaredPatentsRow) GetHasPedExtension() bool`

GetHasPedExtension returns the HasPedExtension field if non-nil, zero value otherwise.

### GetHasPedExtensionOk

`func (o *DrugDeclaredPatentsRow) GetHasPedExtensionOk() (*bool, bool)`

GetHasPedExtensionOk returns a tuple with the HasPedExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPedExtension

`func (o *DrugDeclaredPatentsRow) SetHasPedExtension(v bool)`

SetHasPedExtension sets HasPedExtension field to given value.

### HasHasPedExtension

`func (o *DrugDeclaredPatentsRow) HasHasPedExtension() bool`

HasHasPedExtension returns a boolean if a field has been set.

### SetHasPedExtensionNil

`func (o *DrugDeclaredPatentsRow) SetHasPedExtensionNil(b bool)`

 SetHasPedExtensionNil sets the value for HasPedExtension to be an explicit nil

### UnsetHasPedExtension
`func (o *DrugDeclaredPatentsRow) UnsetHasPedExtension()`

UnsetHasPedExtension ensures that no value is present for HasPedExtension, not even an explicit nil
### GetAssigneeOrganization

`func (o *DrugDeclaredPatentsRow) GetAssigneeOrganization() string`

GetAssigneeOrganization returns the AssigneeOrganization field if non-nil, zero value otherwise.

### GetAssigneeOrganizationOk

`func (o *DrugDeclaredPatentsRow) GetAssigneeOrganizationOk() (*string, bool)`

GetAssigneeOrganizationOk returns a tuple with the AssigneeOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeOrganization

`func (o *DrugDeclaredPatentsRow) SetAssigneeOrganization(v string)`

SetAssigneeOrganization sets AssigneeOrganization field to given value.

### HasAssigneeOrganization

`func (o *DrugDeclaredPatentsRow) HasAssigneeOrganization() bool`

HasAssigneeOrganization returns a boolean if a field has been set.

### SetAssigneeOrganizationNil

`func (o *DrugDeclaredPatentsRow) SetAssigneeOrganizationNil(b bool)`

 SetAssigneeOrganizationNil sets the value for AssigneeOrganization to be an explicit nil

### UnsetAssigneeOrganization
`func (o *DrugDeclaredPatentsRow) UnsetAssigneeOrganization()`

UnsetAssigneeOrganization ensures that no value is present for AssigneeOrganization, not even an explicit nil
### GetAssigneeOrgId

`func (o *DrugDeclaredPatentsRow) GetAssigneeOrgId() string`

GetAssigneeOrgId returns the AssigneeOrgId field if non-nil, zero value otherwise.

### GetAssigneeOrgIdOk

`func (o *DrugDeclaredPatentsRow) GetAssigneeOrgIdOk() (*string, bool)`

GetAssigneeOrgIdOk returns a tuple with the AssigneeOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeOrgId

`func (o *DrugDeclaredPatentsRow) SetAssigneeOrgId(v string)`

SetAssigneeOrgId sets AssigneeOrgId field to given value.

### HasAssigneeOrgId

`func (o *DrugDeclaredPatentsRow) HasAssigneeOrgId() bool`

HasAssigneeOrgId returns a boolean if a field has been set.

### SetAssigneeOrgIdNil

`func (o *DrugDeclaredPatentsRow) SetAssigneeOrgIdNil(b bool)`

 SetAssigneeOrgIdNil sets the value for AssigneeOrgId to be an explicit nil

### UnsetAssigneeOrgId
`func (o *DrugDeclaredPatentsRow) UnsetAssigneeOrgId()`

UnsetAssigneeOrgId ensures that no value is present for AssigneeOrgId, not even an explicit nil
### GetAssigneeCompanyId

`func (o *DrugDeclaredPatentsRow) GetAssigneeCompanyId() string`

GetAssigneeCompanyId returns the AssigneeCompanyId field if non-nil, zero value otherwise.

### GetAssigneeCompanyIdOk

`func (o *DrugDeclaredPatentsRow) GetAssigneeCompanyIdOk() (*string, bool)`

GetAssigneeCompanyIdOk returns a tuple with the AssigneeCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeCompanyId

`func (o *DrugDeclaredPatentsRow) SetAssigneeCompanyId(v string)`

SetAssigneeCompanyId sets AssigneeCompanyId field to given value.

### HasAssigneeCompanyId

`func (o *DrugDeclaredPatentsRow) HasAssigneeCompanyId() bool`

HasAssigneeCompanyId returns a boolean if a field has been set.

### SetAssigneeCompanyIdNil

`func (o *DrugDeclaredPatentsRow) SetAssigneeCompanyIdNil(b bool)`

 SetAssigneeCompanyIdNil sets the value for AssigneeCompanyId to be an explicit nil

### UnsetAssigneeCompanyId
`func (o *DrugDeclaredPatentsRow) UnsetAssigneeCompanyId()`

UnsetAssigneeCompanyId ensures that no value is present for AssigneeCompanyId, not even an explicit nil
### GetAssigneeCompanyName

`func (o *DrugDeclaredPatentsRow) GetAssigneeCompanyName() string`

GetAssigneeCompanyName returns the AssigneeCompanyName field if non-nil, zero value otherwise.

### GetAssigneeCompanyNameOk

`func (o *DrugDeclaredPatentsRow) GetAssigneeCompanyNameOk() (*string, bool)`

GetAssigneeCompanyNameOk returns a tuple with the AssigneeCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeCompanyName

`func (o *DrugDeclaredPatentsRow) SetAssigneeCompanyName(v string)`

SetAssigneeCompanyName sets AssigneeCompanyName field to given value.

### HasAssigneeCompanyName

`func (o *DrugDeclaredPatentsRow) HasAssigneeCompanyName() bool`

HasAssigneeCompanyName returns a boolean if a field has been set.

### SetAssigneeCompanyNameNil

`func (o *DrugDeclaredPatentsRow) SetAssigneeCompanyNameNil(b bool)`

 SetAssigneeCompanyNameNil sets the value for AssigneeCompanyName to be an explicit nil

### UnsetAssigneeCompanyName
`func (o *DrugDeclaredPatentsRow) UnsetAssigneeCompanyName()`

UnsetAssigneeCompanyName ensures that no value is present for AssigneeCompanyName, not even an explicit nil
### GetAssigneeGroupId

`func (o *DrugDeclaredPatentsRow) GetAssigneeGroupId() string`

GetAssigneeGroupId returns the AssigneeGroupId field if non-nil, zero value otherwise.

### GetAssigneeGroupIdOk

`func (o *DrugDeclaredPatentsRow) GetAssigneeGroupIdOk() (*string, bool)`

GetAssigneeGroupIdOk returns a tuple with the AssigneeGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeGroupId

`func (o *DrugDeclaredPatentsRow) SetAssigneeGroupId(v string)`

SetAssigneeGroupId sets AssigneeGroupId field to given value.

### HasAssigneeGroupId

`func (o *DrugDeclaredPatentsRow) HasAssigneeGroupId() bool`

HasAssigneeGroupId returns a boolean if a field has been set.

### SetAssigneeGroupIdNil

`func (o *DrugDeclaredPatentsRow) SetAssigneeGroupIdNil(b bool)`

 SetAssigneeGroupIdNil sets the value for AssigneeGroupId to be an explicit nil

### UnsetAssigneeGroupId
`func (o *DrugDeclaredPatentsRow) UnsetAssigneeGroupId()`

UnsetAssigneeGroupId ensures that no value is present for AssigneeGroupId, not even an explicit nil
### GetAssigneeGroupName

`func (o *DrugDeclaredPatentsRow) GetAssigneeGroupName() string`

GetAssigneeGroupName returns the AssigneeGroupName field if non-nil, zero value otherwise.

### GetAssigneeGroupNameOk

`func (o *DrugDeclaredPatentsRow) GetAssigneeGroupNameOk() (*string, bool)`

GetAssigneeGroupNameOk returns a tuple with the AssigneeGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeGroupName

`func (o *DrugDeclaredPatentsRow) SetAssigneeGroupName(v string)`

SetAssigneeGroupName sets AssigneeGroupName field to given value.

### HasAssigneeGroupName

`func (o *DrugDeclaredPatentsRow) HasAssigneeGroupName() bool`

HasAssigneeGroupName returns a boolean if a field has been set.

### SetAssigneeGroupNameNil

`func (o *DrugDeclaredPatentsRow) SetAssigneeGroupNameNil(b bool)`

 SetAssigneeGroupNameNil sets the value for AssigneeGroupName to be an explicit nil

### UnsetAssigneeGroupName
`func (o *DrugDeclaredPatentsRow) UnsetAssigneeGroupName()`

UnsetAssigneeGroupName ensures that no value is present for AssigneeGroupName, not even an explicit nil
### GetOrgInSpine

`func (o *DrugDeclaredPatentsRow) GetOrgInSpine() bool`

GetOrgInSpine returns the OrgInSpine field if non-nil, zero value otherwise.

### GetOrgInSpineOk

`func (o *DrugDeclaredPatentsRow) GetOrgInSpineOk() (*bool, bool)`

GetOrgInSpineOk returns a tuple with the OrgInSpine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgInSpine

`func (o *DrugDeclaredPatentsRow) SetOrgInSpine(v bool)`

SetOrgInSpine sets OrgInSpine field to given value.

### HasOrgInSpine

`func (o *DrugDeclaredPatentsRow) HasOrgInSpine() bool`

HasOrgInSpine returns a boolean if a field has been set.

### SetOrgInSpineNil

`func (o *DrugDeclaredPatentsRow) SetOrgInSpineNil(b bool)`

 SetOrgInSpineNil sets the value for OrgInSpine to be an explicit nil

### UnsetOrgInSpine
`func (o *DrugDeclaredPatentsRow) UnsetOrgInSpine()`

UnsetOrgInSpine ensures that no value is present for OrgInSpine, not even an explicit nil
### GetResolutionMethod

`func (o *DrugDeclaredPatentsRow) GetResolutionMethod() string`

GetResolutionMethod returns the ResolutionMethod field if non-nil, zero value otherwise.

### GetResolutionMethodOk

`func (o *DrugDeclaredPatentsRow) GetResolutionMethodOk() (*string, bool)`

GetResolutionMethodOk returns a tuple with the ResolutionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionMethod

`func (o *DrugDeclaredPatentsRow) SetResolutionMethod(v string)`

SetResolutionMethod sets ResolutionMethod field to given value.

### HasResolutionMethod

`func (o *DrugDeclaredPatentsRow) HasResolutionMethod() bool`

HasResolutionMethod returns a boolean if a field has been set.

### SetResolutionMethodNil

`func (o *DrugDeclaredPatentsRow) SetResolutionMethodNil(b bool)`

 SetResolutionMethodNil sets the value for ResolutionMethod to be an explicit nil

### UnsetResolutionMethod
`func (o *DrugDeclaredPatentsRow) UnsetResolutionMethod()`

UnsetResolutionMethod ensures that no value is present for ResolutionMethod, not even an explicit nil
### GetMatchConfidence

`func (o *DrugDeclaredPatentsRow) GetMatchConfidence() string`

GetMatchConfidence returns the MatchConfidence field if non-nil, zero value otherwise.

### GetMatchConfidenceOk

`func (o *DrugDeclaredPatentsRow) GetMatchConfidenceOk() (*string, bool)`

GetMatchConfidenceOk returns a tuple with the MatchConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchConfidence

`func (o *DrugDeclaredPatentsRow) SetMatchConfidence(v string)`

SetMatchConfidence sets MatchConfidence field to given value.

### HasMatchConfidence

`func (o *DrugDeclaredPatentsRow) HasMatchConfidence() bool`

HasMatchConfidence returns a boolean if a field has been set.

### SetMatchConfidenceNil

`func (o *DrugDeclaredPatentsRow) SetMatchConfidenceNil(b bool)`

 SetMatchConfidenceNil sets the value for MatchConfidence to be an explicit nil

### UnsetMatchConfidence
`func (o *DrugDeclaredPatentsRow) UnsetMatchConfidence()`

UnsetMatchConfidence ensures that no value is present for MatchConfidence, not even an explicit nil
### GetOwnerBasis

`func (o *DrugDeclaredPatentsRow) GetOwnerBasis() string`

GetOwnerBasis returns the OwnerBasis field if non-nil, zero value otherwise.

### GetOwnerBasisOk

`func (o *DrugDeclaredPatentsRow) GetOwnerBasisOk() (*string, bool)`

GetOwnerBasisOk returns a tuple with the OwnerBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerBasis

`func (o *DrugDeclaredPatentsRow) SetOwnerBasis(v string)`

SetOwnerBasis sets OwnerBasis field to given value.

### HasOwnerBasis

`func (o *DrugDeclaredPatentsRow) HasOwnerBasis() bool`

HasOwnerBasis returns a boolean if a field has been set.

### SetOwnerBasisNil

`func (o *DrugDeclaredPatentsRow) SetOwnerBasisNil(b bool)`

 SetOwnerBasisNil sets the value for OwnerBasis to be an explicit nil

### UnsetOwnerBasis
`func (o *DrugDeclaredPatentsRow) UnsetOwnerBasis()`

UnsetOwnerBasis ensures that no value is present for OwnerBasis, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


