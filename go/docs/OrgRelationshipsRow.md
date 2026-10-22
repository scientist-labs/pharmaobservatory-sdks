# OrgRelationshipsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectOrgId** | Pointer to **NullableString** | Legal-entity level: the subject&#39;s organization record. | [optional] 
**SubjectName** | Pointer to **NullableString** | The organization&#39;s name as filed. | [optional] 
**SubjectCompanyId** | Pointer to **NullableString** | Company level: the subject&#39;s company (null when unresolved). | [optional] 
**SubjectCompanyName** | Pointer to **NullableString** | Display name of the subject&#39;s company. | [optional] 
**SubjectGroupId** | Pointer to **NullableString** | Group level: the subject&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**SubjectGroupName** | Pointer to **NullableString** | Display name of the subject&#39;s group. | [optional] 
**ObjectOrgId** | Pointer to **NullableString** | Legal-entity level: the object&#39;s organization record. | [optional] 
**ObjectName** | Pointer to **NullableString** | The organization&#39;s name as filed. | [optional] 
**ObjectCompanyId** | Pointer to **NullableString** | Company level: the object&#39;s company (null when unresolved). | [optional] 
**ObjectCompanyName** | Pointer to **NullableString** | Display name of the object&#39;s company. | [optional] 
**ObjectGroupId** | Pointer to **NullableString** | Group level: the object&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**ObjectGroupName** | Pointer to **NullableString** | Display name of the object&#39;s group. | [optional] 
**Predicate** | Pointer to **NullableString** | The directed relationship from subject to object: parent_of and ultimate_parent_of (corporate ownership, the subject is the parent), linked_charity_of (a charity linked to the main charity), uses_vendor (a sponsor and the contract research organization inferred for its trials), co_sponsors_trial_with (two sponsors of one trial), project_partner_of and awards_to (UK research grants). | [optional] 
**EdgeFamily** | Pointer to **NullableString** | The family of the predicate: corporate_structure (parent_of, ultimate_parent_of, linked_charity_of), contract_vendor (uses_vendor), co_sponsorship (co_sponsors_trial_with) or research_funding (project_partner_of, awards_to). | [optional] 
**EdgeScope** | Pointer to **NullableString** | general for an edge from a source that covers organizations at large (GLEIF, Charity Commission, the contract research organization (CRO) inference); rare_disease_scoped for an edge from a source read only for organizations in the rare-disease evidence set (ROR, co-sponsorship, UK research grants). A rare_disease_scoped edge is a true edge and an incomplete picture: an organization with none of them has not been shown to have none. | [optional] 
**SourceKey** | Pointer to **NullableString** | The source that asserted the edge: gleif (Global Legal Entity Identifier Foundation Golden Copy), charity_commission, aact_cro_inference (inferred from ClinicalTrials.gov filings against a curated CRO list), ror (Research Organization Registry), aact (ClinicalTrials.gov) or ukri_gtr (UK Research and Innovation Gateway to Research). | [optional] 
**EvidenceType** | Pointer to **NullableString** | The kind of evidence behind the edge: registry (a register entry), trial, project or grant. | [optional] 
**AuthorityClass** | Pointer to **NullableString** | authoritative when the source itself filed the relationship; derived when the edge is an inference (the CRO edge is a curated-list name match, not a filing). | [optional] 
**EvidenceCount** | Pointer to **NullableInt64** | The source records behind the edge. Counted within one source, so it does not add across sources. | [optional] 
**ContextCount** | Pointer to **NullableInt64** | Distinct source record contexts among them (for example distinct trials). | [optional] 
**Confidence** | Pointer to **NullableFloat32** | The highest confidence among the edge&#39;s records, 0 to 1, where the source states one. Null for sources that state none. A GLEIF ownership-resolution score and a CRO name-match score are different measures; do not compare them. | [optional] 
**FirstObservedAt** | Pointer to **NullableTime** | When the earliest of the edge&#39;s records was observed in the source. | [optional] 
**LastObservedAt** | Pointer to **NullableTime** | When the latest of the edge&#39;s records was observed in the source. | [optional] 
**LicenseName** | Pointer to **NullableString** | The licence of the source behind the edge. | [optional] 
**Attribution** | Pointer to **NullableString** | The credit line the source requires for the edge. Show it wherever the edge is shown: the Charity Commission and UK Research and Innovation edges require it under the Open Government Licence v3.0. | [optional] 
**RestrictionTier** | Pointer to **NullableString** | The tier vocabulary is none or licensable. Every row here is none: every source behind these edges is open, and a test on the source assertions reports any restricted source that is ever added. | [optional] 

## Methods

### NewOrgRelationshipsRow

`func NewOrgRelationshipsRow() *OrgRelationshipsRow`

NewOrgRelationshipsRow instantiates a new OrgRelationshipsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgRelationshipsRowWithDefaults

`func NewOrgRelationshipsRowWithDefaults() *OrgRelationshipsRow`

NewOrgRelationshipsRowWithDefaults instantiates a new OrgRelationshipsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectOrgId

`func (o *OrgRelationshipsRow) GetSubjectOrgId() string`

GetSubjectOrgId returns the SubjectOrgId field if non-nil, zero value otherwise.

### GetSubjectOrgIdOk

`func (o *OrgRelationshipsRow) GetSubjectOrgIdOk() (*string, bool)`

GetSubjectOrgIdOk returns a tuple with the SubjectOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectOrgId

`func (o *OrgRelationshipsRow) SetSubjectOrgId(v string)`

SetSubjectOrgId sets SubjectOrgId field to given value.

### HasSubjectOrgId

`func (o *OrgRelationshipsRow) HasSubjectOrgId() bool`

HasSubjectOrgId returns a boolean if a field has been set.

### SetSubjectOrgIdNil

`func (o *OrgRelationshipsRow) SetSubjectOrgIdNil(b bool)`

 SetSubjectOrgIdNil sets the value for SubjectOrgId to be an explicit nil

### UnsetSubjectOrgId
`func (o *OrgRelationshipsRow) UnsetSubjectOrgId()`

UnsetSubjectOrgId ensures that no value is present for SubjectOrgId, not even an explicit nil
### GetSubjectName

`func (o *OrgRelationshipsRow) GetSubjectName() string`

GetSubjectName returns the SubjectName field if non-nil, zero value otherwise.

### GetSubjectNameOk

`func (o *OrgRelationshipsRow) GetSubjectNameOk() (*string, bool)`

GetSubjectNameOk returns a tuple with the SubjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectName

`func (o *OrgRelationshipsRow) SetSubjectName(v string)`

SetSubjectName sets SubjectName field to given value.

### HasSubjectName

`func (o *OrgRelationshipsRow) HasSubjectName() bool`

HasSubjectName returns a boolean if a field has been set.

### SetSubjectNameNil

`func (o *OrgRelationshipsRow) SetSubjectNameNil(b bool)`

 SetSubjectNameNil sets the value for SubjectName to be an explicit nil

### UnsetSubjectName
`func (o *OrgRelationshipsRow) UnsetSubjectName()`

UnsetSubjectName ensures that no value is present for SubjectName, not even an explicit nil
### GetSubjectCompanyId

`func (o *OrgRelationshipsRow) GetSubjectCompanyId() string`

GetSubjectCompanyId returns the SubjectCompanyId field if non-nil, zero value otherwise.

### GetSubjectCompanyIdOk

`func (o *OrgRelationshipsRow) GetSubjectCompanyIdOk() (*string, bool)`

GetSubjectCompanyIdOk returns a tuple with the SubjectCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectCompanyId

`func (o *OrgRelationshipsRow) SetSubjectCompanyId(v string)`

SetSubjectCompanyId sets SubjectCompanyId field to given value.

### HasSubjectCompanyId

`func (o *OrgRelationshipsRow) HasSubjectCompanyId() bool`

HasSubjectCompanyId returns a boolean if a field has been set.

### SetSubjectCompanyIdNil

`func (o *OrgRelationshipsRow) SetSubjectCompanyIdNil(b bool)`

 SetSubjectCompanyIdNil sets the value for SubjectCompanyId to be an explicit nil

### UnsetSubjectCompanyId
`func (o *OrgRelationshipsRow) UnsetSubjectCompanyId()`

UnsetSubjectCompanyId ensures that no value is present for SubjectCompanyId, not even an explicit nil
### GetSubjectCompanyName

`func (o *OrgRelationshipsRow) GetSubjectCompanyName() string`

GetSubjectCompanyName returns the SubjectCompanyName field if non-nil, zero value otherwise.

### GetSubjectCompanyNameOk

`func (o *OrgRelationshipsRow) GetSubjectCompanyNameOk() (*string, bool)`

GetSubjectCompanyNameOk returns a tuple with the SubjectCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectCompanyName

`func (o *OrgRelationshipsRow) SetSubjectCompanyName(v string)`

SetSubjectCompanyName sets SubjectCompanyName field to given value.

### HasSubjectCompanyName

`func (o *OrgRelationshipsRow) HasSubjectCompanyName() bool`

HasSubjectCompanyName returns a boolean if a field has been set.

### SetSubjectCompanyNameNil

`func (o *OrgRelationshipsRow) SetSubjectCompanyNameNil(b bool)`

 SetSubjectCompanyNameNil sets the value for SubjectCompanyName to be an explicit nil

### UnsetSubjectCompanyName
`func (o *OrgRelationshipsRow) UnsetSubjectCompanyName()`

UnsetSubjectCompanyName ensures that no value is present for SubjectCompanyName, not even an explicit nil
### GetSubjectGroupId

`func (o *OrgRelationshipsRow) GetSubjectGroupId() string`

GetSubjectGroupId returns the SubjectGroupId field if non-nil, zero value otherwise.

### GetSubjectGroupIdOk

`func (o *OrgRelationshipsRow) GetSubjectGroupIdOk() (*string, bool)`

GetSubjectGroupIdOk returns a tuple with the SubjectGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectGroupId

`func (o *OrgRelationshipsRow) SetSubjectGroupId(v string)`

SetSubjectGroupId sets SubjectGroupId field to given value.

### HasSubjectGroupId

`func (o *OrgRelationshipsRow) HasSubjectGroupId() bool`

HasSubjectGroupId returns a boolean if a field has been set.

### SetSubjectGroupIdNil

`func (o *OrgRelationshipsRow) SetSubjectGroupIdNil(b bool)`

 SetSubjectGroupIdNil sets the value for SubjectGroupId to be an explicit nil

### UnsetSubjectGroupId
`func (o *OrgRelationshipsRow) UnsetSubjectGroupId()`

UnsetSubjectGroupId ensures that no value is present for SubjectGroupId, not even an explicit nil
### GetSubjectGroupName

`func (o *OrgRelationshipsRow) GetSubjectGroupName() string`

GetSubjectGroupName returns the SubjectGroupName field if non-nil, zero value otherwise.

### GetSubjectGroupNameOk

`func (o *OrgRelationshipsRow) GetSubjectGroupNameOk() (*string, bool)`

GetSubjectGroupNameOk returns a tuple with the SubjectGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectGroupName

`func (o *OrgRelationshipsRow) SetSubjectGroupName(v string)`

SetSubjectGroupName sets SubjectGroupName field to given value.

### HasSubjectGroupName

`func (o *OrgRelationshipsRow) HasSubjectGroupName() bool`

HasSubjectGroupName returns a boolean if a field has been set.

### SetSubjectGroupNameNil

`func (o *OrgRelationshipsRow) SetSubjectGroupNameNil(b bool)`

 SetSubjectGroupNameNil sets the value for SubjectGroupName to be an explicit nil

### UnsetSubjectGroupName
`func (o *OrgRelationshipsRow) UnsetSubjectGroupName()`

UnsetSubjectGroupName ensures that no value is present for SubjectGroupName, not even an explicit nil
### GetObjectOrgId

`func (o *OrgRelationshipsRow) GetObjectOrgId() string`

GetObjectOrgId returns the ObjectOrgId field if non-nil, zero value otherwise.

### GetObjectOrgIdOk

`func (o *OrgRelationshipsRow) GetObjectOrgIdOk() (*string, bool)`

GetObjectOrgIdOk returns a tuple with the ObjectOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectOrgId

`func (o *OrgRelationshipsRow) SetObjectOrgId(v string)`

SetObjectOrgId sets ObjectOrgId field to given value.

### HasObjectOrgId

`func (o *OrgRelationshipsRow) HasObjectOrgId() bool`

HasObjectOrgId returns a boolean if a field has been set.

### SetObjectOrgIdNil

`func (o *OrgRelationshipsRow) SetObjectOrgIdNil(b bool)`

 SetObjectOrgIdNil sets the value for ObjectOrgId to be an explicit nil

### UnsetObjectOrgId
`func (o *OrgRelationshipsRow) UnsetObjectOrgId()`

UnsetObjectOrgId ensures that no value is present for ObjectOrgId, not even an explicit nil
### GetObjectName

`func (o *OrgRelationshipsRow) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *OrgRelationshipsRow) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *OrgRelationshipsRow) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *OrgRelationshipsRow) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### SetObjectNameNil

`func (o *OrgRelationshipsRow) SetObjectNameNil(b bool)`

 SetObjectNameNil sets the value for ObjectName to be an explicit nil

### UnsetObjectName
`func (o *OrgRelationshipsRow) UnsetObjectName()`

UnsetObjectName ensures that no value is present for ObjectName, not even an explicit nil
### GetObjectCompanyId

`func (o *OrgRelationshipsRow) GetObjectCompanyId() string`

GetObjectCompanyId returns the ObjectCompanyId field if non-nil, zero value otherwise.

### GetObjectCompanyIdOk

`func (o *OrgRelationshipsRow) GetObjectCompanyIdOk() (*string, bool)`

GetObjectCompanyIdOk returns a tuple with the ObjectCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectCompanyId

`func (o *OrgRelationshipsRow) SetObjectCompanyId(v string)`

SetObjectCompanyId sets ObjectCompanyId field to given value.

### HasObjectCompanyId

`func (o *OrgRelationshipsRow) HasObjectCompanyId() bool`

HasObjectCompanyId returns a boolean if a field has been set.

### SetObjectCompanyIdNil

`func (o *OrgRelationshipsRow) SetObjectCompanyIdNil(b bool)`

 SetObjectCompanyIdNil sets the value for ObjectCompanyId to be an explicit nil

### UnsetObjectCompanyId
`func (o *OrgRelationshipsRow) UnsetObjectCompanyId()`

UnsetObjectCompanyId ensures that no value is present for ObjectCompanyId, not even an explicit nil
### GetObjectCompanyName

`func (o *OrgRelationshipsRow) GetObjectCompanyName() string`

GetObjectCompanyName returns the ObjectCompanyName field if non-nil, zero value otherwise.

### GetObjectCompanyNameOk

`func (o *OrgRelationshipsRow) GetObjectCompanyNameOk() (*string, bool)`

GetObjectCompanyNameOk returns a tuple with the ObjectCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectCompanyName

`func (o *OrgRelationshipsRow) SetObjectCompanyName(v string)`

SetObjectCompanyName sets ObjectCompanyName field to given value.

### HasObjectCompanyName

`func (o *OrgRelationshipsRow) HasObjectCompanyName() bool`

HasObjectCompanyName returns a boolean if a field has been set.

### SetObjectCompanyNameNil

`func (o *OrgRelationshipsRow) SetObjectCompanyNameNil(b bool)`

 SetObjectCompanyNameNil sets the value for ObjectCompanyName to be an explicit nil

### UnsetObjectCompanyName
`func (o *OrgRelationshipsRow) UnsetObjectCompanyName()`

UnsetObjectCompanyName ensures that no value is present for ObjectCompanyName, not even an explicit nil
### GetObjectGroupId

`func (o *OrgRelationshipsRow) GetObjectGroupId() string`

GetObjectGroupId returns the ObjectGroupId field if non-nil, zero value otherwise.

### GetObjectGroupIdOk

`func (o *OrgRelationshipsRow) GetObjectGroupIdOk() (*string, bool)`

GetObjectGroupIdOk returns a tuple with the ObjectGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectGroupId

`func (o *OrgRelationshipsRow) SetObjectGroupId(v string)`

SetObjectGroupId sets ObjectGroupId field to given value.

### HasObjectGroupId

`func (o *OrgRelationshipsRow) HasObjectGroupId() bool`

HasObjectGroupId returns a boolean if a field has been set.

### SetObjectGroupIdNil

`func (o *OrgRelationshipsRow) SetObjectGroupIdNil(b bool)`

 SetObjectGroupIdNil sets the value for ObjectGroupId to be an explicit nil

### UnsetObjectGroupId
`func (o *OrgRelationshipsRow) UnsetObjectGroupId()`

UnsetObjectGroupId ensures that no value is present for ObjectGroupId, not even an explicit nil
### GetObjectGroupName

`func (o *OrgRelationshipsRow) GetObjectGroupName() string`

GetObjectGroupName returns the ObjectGroupName field if non-nil, zero value otherwise.

### GetObjectGroupNameOk

`func (o *OrgRelationshipsRow) GetObjectGroupNameOk() (*string, bool)`

GetObjectGroupNameOk returns a tuple with the ObjectGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectGroupName

`func (o *OrgRelationshipsRow) SetObjectGroupName(v string)`

SetObjectGroupName sets ObjectGroupName field to given value.

### HasObjectGroupName

`func (o *OrgRelationshipsRow) HasObjectGroupName() bool`

HasObjectGroupName returns a boolean if a field has been set.

### SetObjectGroupNameNil

`func (o *OrgRelationshipsRow) SetObjectGroupNameNil(b bool)`

 SetObjectGroupNameNil sets the value for ObjectGroupName to be an explicit nil

### UnsetObjectGroupName
`func (o *OrgRelationshipsRow) UnsetObjectGroupName()`

UnsetObjectGroupName ensures that no value is present for ObjectGroupName, not even an explicit nil
### GetPredicate

`func (o *OrgRelationshipsRow) GetPredicate() string`

GetPredicate returns the Predicate field if non-nil, zero value otherwise.

### GetPredicateOk

`func (o *OrgRelationshipsRow) GetPredicateOk() (*string, bool)`

GetPredicateOk returns a tuple with the Predicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicate

`func (o *OrgRelationshipsRow) SetPredicate(v string)`

SetPredicate sets Predicate field to given value.

### HasPredicate

`func (o *OrgRelationshipsRow) HasPredicate() bool`

HasPredicate returns a boolean if a field has been set.

### SetPredicateNil

`func (o *OrgRelationshipsRow) SetPredicateNil(b bool)`

 SetPredicateNil sets the value for Predicate to be an explicit nil

### UnsetPredicate
`func (o *OrgRelationshipsRow) UnsetPredicate()`

UnsetPredicate ensures that no value is present for Predicate, not even an explicit nil
### GetEdgeFamily

`func (o *OrgRelationshipsRow) GetEdgeFamily() string`

GetEdgeFamily returns the EdgeFamily field if non-nil, zero value otherwise.

### GetEdgeFamilyOk

`func (o *OrgRelationshipsRow) GetEdgeFamilyOk() (*string, bool)`

GetEdgeFamilyOk returns a tuple with the EdgeFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFamily

`func (o *OrgRelationshipsRow) SetEdgeFamily(v string)`

SetEdgeFamily sets EdgeFamily field to given value.

### HasEdgeFamily

`func (o *OrgRelationshipsRow) HasEdgeFamily() bool`

HasEdgeFamily returns a boolean if a field has been set.

### SetEdgeFamilyNil

`func (o *OrgRelationshipsRow) SetEdgeFamilyNil(b bool)`

 SetEdgeFamilyNil sets the value for EdgeFamily to be an explicit nil

### UnsetEdgeFamily
`func (o *OrgRelationshipsRow) UnsetEdgeFamily()`

UnsetEdgeFamily ensures that no value is present for EdgeFamily, not even an explicit nil
### GetEdgeScope

`func (o *OrgRelationshipsRow) GetEdgeScope() string`

GetEdgeScope returns the EdgeScope field if non-nil, zero value otherwise.

### GetEdgeScopeOk

`func (o *OrgRelationshipsRow) GetEdgeScopeOk() (*string, bool)`

GetEdgeScopeOk returns a tuple with the EdgeScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeScope

`func (o *OrgRelationshipsRow) SetEdgeScope(v string)`

SetEdgeScope sets EdgeScope field to given value.

### HasEdgeScope

`func (o *OrgRelationshipsRow) HasEdgeScope() bool`

HasEdgeScope returns a boolean if a field has been set.

### SetEdgeScopeNil

`func (o *OrgRelationshipsRow) SetEdgeScopeNil(b bool)`

 SetEdgeScopeNil sets the value for EdgeScope to be an explicit nil

### UnsetEdgeScope
`func (o *OrgRelationshipsRow) UnsetEdgeScope()`

UnsetEdgeScope ensures that no value is present for EdgeScope, not even an explicit nil
### GetSourceKey

`func (o *OrgRelationshipsRow) GetSourceKey() string`

GetSourceKey returns the SourceKey field if non-nil, zero value otherwise.

### GetSourceKeyOk

`func (o *OrgRelationshipsRow) GetSourceKeyOk() (*string, bool)`

GetSourceKeyOk returns a tuple with the SourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceKey

`func (o *OrgRelationshipsRow) SetSourceKey(v string)`

SetSourceKey sets SourceKey field to given value.

### HasSourceKey

`func (o *OrgRelationshipsRow) HasSourceKey() bool`

HasSourceKey returns a boolean if a field has been set.

### SetSourceKeyNil

`func (o *OrgRelationshipsRow) SetSourceKeyNil(b bool)`

 SetSourceKeyNil sets the value for SourceKey to be an explicit nil

### UnsetSourceKey
`func (o *OrgRelationshipsRow) UnsetSourceKey()`

UnsetSourceKey ensures that no value is present for SourceKey, not even an explicit nil
### GetEvidenceType

`func (o *OrgRelationshipsRow) GetEvidenceType() string`

GetEvidenceType returns the EvidenceType field if non-nil, zero value otherwise.

### GetEvidenceTypeOk

`func (o *OrgRelationshipsRow) GetEvidenceTypeOk() (*string, bool)`

GetEvidenceTypeOk returns a tuple with the EvidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceType

`func (o *OrgRelationshipsRow) SetEvidenceType(v string)`

SetEvidenceType sets EvidenceType field to given value.

### HasEvidenceType

`func (o *OrgRelationshipsRow) HasEvidenceType() bool`

HasEvidenceType returns a boolean if a field has been set.

### SetEvidenceTypeNil

`func (o *OrgRelationshipsRow) SetEvidenceTypeNil(b bool)`

 SetEvidenceTypeNil sets the value for EvidenceType to be an explicit nil

### UnsetEvidenceType
`func (o *OrgRelationshipsRow) UnsetEvidenceType()`

UnsetEvidenceType ensures that no value is present for EvidenceType, not even an explicit nil
### GetAuthorityClass

`func (o *OrgRelationshipsRow) GetAuthorityClass() string`

GetAuthorityClass returns the AuthorityClass field if non-nil, zero value otherwise.

### GetAuthorityClassOk

`func (o *OrgRelationshipsRow) GetAuthorityClassOk() (*string, bool)`

GetAuthorityClassOk returns a tuple with the AuthorityClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityClass

`func (o *OrgRelationshipsRow) SetAuthorityClass(v string)`

SetAuthorityClass sets AuthorityClass field to given value.

### HasAuthorityClass

`func (o *OrgRelationshipsRow) HasAuthorityClass() bool`

HasAuthorityClass returns a boolean if a field has been set.

### SetAuthorityClassNil

`func (o *OrgRelationshipsRow) SetAuthorityClassNil(b bool)`

 SetAuthorityClassNil sets the value for AuthorityClass to be an explicit nil

### UnsetAuthorityClass
`func (o *OrgRelationshipsRow) UnsetAuthorityClass()`

UnsetAuthorityClass ensures that no value is present for AuthorityClass, not even an explicit nil
### GetEvidenceCount

`func (o *OrgRelationshipsRow) GetEvidenceCount() int64`

GetEvidenceCount returns the EvidenceCount field if non-nil, zero value otherwise.

### GetEvidenceCountOk

`func (o *OrgRelationshipsRow) GetEvidenceCountOk() (*int64, bool)`

GetEvidenceCountOk returns a tuple with the EvidenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceCount

`func (o *OrgRelationshipsRow) SetEvidenceCount(v int64)`

SetEvidenceCount sets EvidenceCount field to given value.

### HasEvidenceCount

`func (o *OrgRelationshipsRow) HasEvidenceCount() bool`

HasEvidenceCount returns a boolean if a field has been set.

### SetEvidenceCountNil

`func (o *OrgRelationshipsRow) SetEvidenceCountNil(b bool)`

 SetEvidenceCountNil sets the value for EvidenceCount to be an explicit nil

### UnsetEvidenceCount
`func (o *OrgRelationshipsRow) UnsetEvidenceCount()`

UnsetEvidenceCount ensures that no value is present for EvidenceCount, not even an explicit nil
### GetContextCount

`func (o *OrgRelationshipsRow) GetContextCount() int64`

GetContextCount returns the ContextCount field if non-nil, zero value otherwise.

### GetContextCountOk

`func (o *OrgRelationshipsRow) GetContextCountOk() (*int64, bool)`

GetContextCountOk returns a tuple with the ContextCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextCount

`func (o *OrgRelationshipsRow) SetContextCount(v int64)`

SetContextCount sets ContextCount field to given value.

### HasContextCount

`func (o *OrgRelationshipsRow) HasContextCount() bool`

HasContextCount returns a boolean if a field has been set.

### SetContextCountNil

`func (o *OrgRelationshipsRow) SetContextCountNil(b bool)`

 SetContextCountNil sets the value for ContextCount to be an explicit nil

### UnsetContextCount
`func (o *OrgRelationshipsRow) UnsetContextCount()`

UnsetContextCount ensures that no value is present for ContextCount, not even an explicit nil
### GetConfidence

`func (o *OrgRelationshipsRow) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *OrgRelationshipsRow) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *OrgRelationshipsRow) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *OrgRelationshipsRow) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### SetConfidenceNil

`func (o *OrgRelationshipsRow) SetConfidenceNil(b bool)`

 SetConfidenceNil sets the value for Confidence to be an explicit nil

### UnsetConfidence
`func (o *OrgRelationshipsRow) UnsetConfidence()`

UnsetConfidence ensures that no value is present for Confidence, not even an explicit nil
### GetFirstObservedAt

`func (o *OrgRelationshipsRow) GetFirstObservedAt() time.Time`

GetFirstObservedAt returns the FirstObservedAt field if non-nil, zero value otherwise.

### GetFirstObservedAtOk

`func (o *OrgRelationshipsRow) GetFirstObservedAtOk() (*time.Time, bool)`

GetFirstObservedAtOk returns a tuple with the FirstObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstObservedAt

`func (o *OrgRelationshipsRow) SetFirstObservedAt(v time.Time)`

SetFirstObservedAt sets FirstObservedAt field to given value.

### HasFirstObservedAt

`func (o *OrgRelationshipsRow) HasFirstObservedAt() bool`

HasFirstObservedAt returns a boolean if a field has been set.

### SetFirstObservedAtNil

`func (o *OrgRelationshipsRow) SetFirstObservedAtNil(b bool)`

 SetFirstObservedAtNil sets the value for FirstObservedAt to be an explicit nil

### UnsetFirstObservedAt
`func (o *OrgRelationshipsRow) UnsetFirstObservedAt()`

UnsetFirstObservedAt ensures that no value is present for FirstObservedAt, not even an explicit nil
### GetLastObservedAt

`func (o *OrgRelationshipsRow) GetLastObservedAt() time.Time`

GetLastObservedAt returns the LastObservedAt field if non-nil, zero value otherwise.

### GetLastObservedAtOk

`func (o *OrgRelationshipsRow) GetLastObservedAtOk() (*time.Time, bool)`

GetLastObservedAtOk returns a tuple with the LastObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastObservedAt

`func (o *OrgRelationshipsRow) SetLastObservedAt(v time.Time)`

SetLastObservedAt sets LastObservedAt field to given value.

### HasLastObservedAt

`func (o *OrgRelationshipsRow) HasLastObservedAt() bool`

HasLastObservedAt returns a boolean if a field has been set.

### SetLastObservedAtNil

`func (o *OrgRelationshipsRow) SetLastObservedAtNil(b bool)`

 SetLastObservedAtNil sets the value for LastObservedAt to be an explicit nil

### UnsetLastObservedAt
`func (o *OrgRelationshipsRow) UnsetLastObservedAt()`

UnsetLastObservedAt ensures that no value is present for LastObservedAt, not even an explicit nil
### GetLicenseName

`func (o *OrgRelationshipsRow) GetLicenseName() string`

GetLicenseName returns the LicenseName field if non-nil, zero value otherwise.

### GetLicenseNameOk

`func (o *OrgRelationshipsRow) GetLicenseNameOk() (*string, bool)`

GetLicenseNameOk returns a tuple with the LicenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseName

`func (o *OrgRelationshipsRow) SetLicenseName(v string)`

SetLicenseName sets LicenseName field to given value.

### HasLicenseName

`func (o *OrgRelationshipsRow) HasLicenseName() bool`

HasLicenseName returns a boolean if a field has been set.

### SetLicenseNameNil

`func (o *OrgRelationshipsRow) SetLicenseNameNil(b bool)`

 SetLicenseNameNil sets the value for LicenseName to be an explicit nil

### UnsetLicenseName
`func (o *OrgRelationshipsRow) UnsetLicenseName()`

UnsetLicenseName ensures that no value is present for LicenseName, not even an explicit nil
### GetAttribution

`func (o *OrgRelationshipsRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *OrgRelationshipsRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *OrgRelationshipsRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *OrgRelationshipsRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *OrgRelationshipsRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *OrgRelationshipsRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil
### GetRestrictionTier

`func (o *OrgRelationshipsRow) GetRestrictionTier() string`

GetRestrictionTier returns the RestrictionTier field if non-nil, zero value otherwise.

### GetRestrictionTierOk

`func (o *OrgRelationshipsRow) GetRestrictionTierOk() (*string, bool)`

GetRestrictionTierOk returns a tuple with the RestrictionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionTier

`func (o *OrgRelationshipsRow) SetRestrictionTier(v string)`

SetRestrictionTier sets RestrictionTier field to given value.

### HasRestrictionTier

`func (o *OrgRelationshipsRow) HasRestrictionTier() bool`

HasRestrictionTier returns a boolean if a field has been set.

### SetRestrictionTierNil

`func (o *OrgRelationshipsRow) SetRestrictionTierNil(b bool)`

 SetRestrictionTierNil sets the value for RestrictionTier to be an explicit nil

### UnsetRestrictionTier
`func (o *OrgRelationshipsRow) UnsetRestrictionTier()`

UnsetRestrictionTier ensures that no value is present for RestrictionTier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


