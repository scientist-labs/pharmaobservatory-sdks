# OwnershipSourceCoverageRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceClass** | Pointer to **NullableString** | The ownership source class (for example sec_def_14a for proxy statements, sec_form_adv for Form ADV). | [optional] 
**SubjectSourceRef** | Pointer to **NullableString** | The subject as the source names it: the source&#39;s own reference for the owned company. | [optional] 
**SubjectName** | Pointer to **NullableString** | The subject&#39;s name as filed. | [optional] 
**SubjectOrgId** | Pointer to **NullableString** | Legal-entity level: the organization record of the subject. | [optional] 
**SubjectCompanyId** | Pointer to **NullableString** | Company level: the company the subject belongs to (null when unresolved). | [optional] 
**SubjectCompanyName** | Pointer to **NullableString** | Display name of the subject&#39;s company. | [optional] 
**SubjectGroupId** | Pointer to **NullableString** | Group level: the subject&#39;s current ultimate parent company; a root company&#39;s group is itself. | [optional] 
**SubjectGroupName** | Pointer to **NullableString** | Display name of the subject&#39;s group. | [optional] 
**SourceSystem** | Pointer to **NullableString** | The system the source class is read from (for example sec_edgar). | [optional] 
**SourceLoadState** | Pointer to **NullableString** | The whole-source state: loaded, served_by_dedicated_mart (13F is served by the institutional holdings datasets), no_documents_ingested or source_not_ingested. | [optional] 
**SourceDocumentsIngested** | Pointer to **NullableInt64** | Documents the whole source has ingested. | [optional] 
**CoverageState** | Pointer to **NullableString** | What the source read about this subject: reported, no_reportable_owner, extraction_abstained or read_no_owner. | [optional] 
**DocumentsIngested** | Pointer to **NullableInt64** | Documents the source ingested about this subject. | [optional] 
**Observations** | Pointer to **NullableInt64** | Ownership observations the source recorded about this subject. | [optional] 
**CurrentEdges** | Pointer to **NullableInt64** | Ownership relationships that currently stand. | [optional] 
**OwnersDistinct** | Pointer to **NullableInt64** | Distinct owners the source named for this subject. | [optional] 
**OwnersUnresolved** | Pointer to **NullableInt64** | Named owners that did not resolve to an organization or investor. | [optional] 
**NoReportableOwnerStatements** | Pointer to **NullableInt64** | Filings that positively stated the subject has no reportable owner. | [optional] 
**ExtractionAbstained** | Pointer to **NullableInt64** | Documents on which the extractor abstained. | [optional] 
**LastFilingDate** | Pointer to **NullableString** | The latest filing date the source holds for this subject. | [optional] 
**MaxAsOfDate** | Pointer to **NullableString** | The latest as-of date among the subject&#39;s observations. | [optional] 
**SourceLicense** | Pointer to **NullableString** | The licence of the source. | [optional] 
**Attribution** | Pointer to **NullableString** | The credit line for the source. | [optional] 

## Methods

### NewOwnershipSourceCoverageRow

`func NewOwnershipSourceCoverageRow() *OwnershipSourceCoverageRow`

NewOwnershipSourceCoverageRow instantiates a new OwnershipSourceCoverageRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOwnershipSourceCoverageRowWithDefaults

`func NewOwnershipSourceCoverageRowWithDefaults() *OwnershipSourceCoverageRow`

NewOwnershipSourceCoverageRowWithDefaults instantiates a new OwnershipSourceCoverageRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceClass

`func (o *OwnershipSourceCoverageRow) GetSourceClass() string`

GetSourceClass returns the SourceClass field if non-nil, zero value otherwise.

### GetSourceClassOk

`func (o *OwnershipSourceCoverageRow) GetSourceClassOk() (*string, bool)`

GetSourceClassOk returns a tuple with the SourceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceClass

`func (o *OwnershipSourceCoverageRow) SetSourceClass(v string)`

SetSourceClass sets SourceClass field to given value.

### HasSourceClass

`func (o *OwnershipSourceCoverageRow) HasSourceClass() bool`

HasSourceClass returns a boolean if a field has been set.

### SetSourceClassNil

`func (o *OwnershipSourceCoverageRow) SetSourceClassNil(b bool)`

 SetSourceClassNil sets the value for SourceClass to be an explicit nil

### UnsetSourceClass
`func (o *OwnershipSourceCoverageRow) UnsetSourceClass()`

UnsetSourceClass ensures that no value is present for SourceClass, not even an explicit nil
### GetSubjectSourceRef

`func (o *OwnershipSourceCoverageRow) GetSubjectSourceRef() string`

GetSubjectSourceRef returns the SubjectSourceRef field if non-nil, zero value otherwise.

### GetSubjectSourceRefOk

`func (o *OwnershipSourceCoverageRow) GetSubjectSourceRefOk() (*string, bool)`

GetSubjectSourceRefOk returns a tuple with the SubjectSourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectSourceRef

`func (o *OwnershipSourceCoverageRow) SetSubjectSourceRef(v string)`

SetSubjectSourceRef sets SubjectSourceRef field to given value.

### HasSubjectSourceRef

`func (o *OwnershipSourceCoverageRow) HasSubjectSourceRef() bool`

HasSubjectSourceRef returns a boolean if a field has been set.

### SetSubjectSourceRefNil

`func (o *OwnershipSourceCoverageRow) SetSubjectSourceRefNil(b bool)`

 SetSubjectSourceRefNil sets the value for SubjectSourceRef to be an explicit nil

### UnsetSubjectSourceRef
`func (o *OwnershipSourceCoverageRow) UnsetSubjectSourceRef()`

UnsetSubjectSourceRef ensures that no value is present for SubjectSourceRef, not even an explicit nil
### GetSubjectName

`func (o *OwnershipSourceCoverageRow) GetSubjectName() string`

GetSubjectName returns the SubjectName field if non-nil, zero value otherwise.

### GetSubjectNameOk

`func (o *OwnershipSourceCoverageRow) GetSubjectNameOk() (*string, bool)`

GetSubjectNameOk returns a tuple with the SubjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectName

`func (o *OwnershipSourceCoverageRow) SetSubjectName(v string)`

SetSubjectName sets SubjectName field to given value.

### HasSubjectName

`func (o *OwnershipSourceCoverageRow) HasSubjectName() bool`

HasSubjectName returns a boolean if a field has been set.

### SetSubjectNameNil

`func (o *OwnershipSourceCoverageRow) SetSubjectNameNil(b bool)`

 SetSubjectNameNil sets the value for SubjectName to be an explicit nil

### UnsetSubjectName
`func (o *OwnershipSourceCoverageRow) UnsetSubjectName()`

UnsetSubjectName ensures that no value is present for SubjectName, not even an explicit nil
### GetSubjectOrgId

`func (o *OwnershipSourceCoverageRow) GetSubjectOrgId() string`

GetSubjectOrgId returns the SubjectOrgId field if non-nil, zero value otherwise.

### GetSubjectOrgIdOk

`func (o *OwnershipSourceCoverageRow) GetSubjectOrgIdOk() (*string, bool)`

GetSubjectOrgIdOk returns a tuple with the SubjectOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectOrgId

`func (o *OwnershipSourceCoverageRow) SetSubjectOrgId(v string)`

SetSubjectOrgId sets SubjectOrgId field to given value.

### HasSubjectOrgId

`func (o *OwnershipSourceCoverageRow) HasSubjectOrgId() bool`

HasSubjectOrgId returns a boolean if a field has been set.

### SetSubjectOrgIdNil

`func (o *OwnershipSourceCoverageRow) SetSubjectOrgIdNil(b bool)`

 SetSubjectOrgIdNil sets the value for SubjectOrgId to be an explicit nil

### UnsetSubjectOrgId
`func (o *OwnershipSourceCoverageRow) UnsetSubjectOrgId()`

UnsetSubjectOrgId ensures that no value is present for SubjectOrgId, not even an explicit nil
### GetSubjectCompanyId

`func (o *OwnershipSourceCoverageRow) GetSubjectCompanyId() string`

GetSubjectCompanyId returns the SubjectCompanyId field if non-nil, zero value otherwise.

### GetSubjectCompanyIdOk

`func (o *OwnershipSourceCoverageRow) GetSubjectCompanyIdOk() (*string, bool)`

GetSubjectCompanyIdOk returns a tuple with the SubjectCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectCompanyId

`func (o *OwnershipSourceCoverageRow) SetSubjectCompanyId(v string)`

SetSubjectCompanyId sets SubjectCompanyId field to given value.

### HasSubjectCompanyId

`func (o *OwnershipSourceCoverageRow) HasSubjectCompanyId() bool`

HasSubjectCompanyId returns a boolean if a field has been set.

### SetSubjectCompanyIdNil

`func (o *OwnershipSourceCoverageRow) SetSubjectCompanyIdNil(b bool)`

 SetSubjectCompanyIdNil sets the value for SubjectCompanyId to be an explicit nil

### UnsetSubjectCompanyId
`func (o *OwnershipSourceCoverageRow) UnsetSubjectCompanyId()`

UnsetSubjectCompanyId ensures that no value is present for SubjectCompanyId, not even an explicit nil
### GetSubjectCompanyName

`func (o *OwnershipSourceCoverageRow) GetSubjectCompanyName() string`

GetSubjectCompanyName returns the SubjectCompanyName field if non-nil, zero value otherwise.

### GetSubjectCompanyNameOk

`func (o *OwnershipSourceCoverageRow) GetSubjectCompanyNameOk() (*string, bool)`

GetSubjectCompanyNameOk returns a tuple with the SubjectCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectCompanyName

`func (o *OwnershipSourceCoverageRow) SetSubjectCompanyName(v string)`

SetSubjectCompanyName sets SubjectCompanyName field to given value.

### HasSubjectCompanyName

`func (o *OwnershipSourceCoverageRow) HasSubjectCompanyName() bool`

HasSubjectCompanyName returns a boolean if a field has been set.

### SetSubjectCompanyNameNil

`func (o *OwnershipSourceCoverageRow) SetSubjectCompanyNameNil(b bool)`

 SetSubjectCompanyNameNil sets the value for SubjectCompanyName to be an explicit nil

### UnsetSubjectCompanyName
`func (o *OwnershipSourceCoverageRow) UnsetSubjectCompanyName()`

UnsetSubjectCompanyName ensures that no value is present for SubjectCompanyName, not even an explicit nil
### GetSubjectGroupId

`func (o *OwnershipSourceCoverageRow) GetSubjectGroupId() string`

GetSubjectGroupId returns the SubjectGroupId field if non-nil, zero value otherwise.

### GetSubjectGroupIdOk

`func (o *OwnershipSourceCoverageRow) GetSubjectGroupIdOk() (*string, bool)`

GetSubjectGroupIdOk returns a tuple with the SubjectGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectGroupId

`func (o *OwnershipSourceCoverageRow) SetSubjectGroupId(v string)`

SetSubjectGroupId sets SubjectGroupId field to given value.

### HasSubjectGroupId

`func (o *OwnershipSourceCoverageRow) HasSubjectGroupId() bool`

HasSubjectGroupId returns a boolean if a field has been set.

### SetSubjectGroupIdNil

`func (o *OwnershipSourceCoverageRow) SetSubjectGroupIdNil(b bool)`

 SetSubjectGroupIdNil sets the value for SubjectGroupId to be an explicit nil

### UnsetSubjectGroupId
`func (o *OwnershipSourceCoverageRow) UnsetSubjectGroupId()`

UnsetSubjectGroupId ensures that no value is present for SubjectGroupId, not even an explicit nil
### GetSubjectGroupName

`func (o *OwnershipSourceCoverageRow) GetSubjectGroupName() string`

GetSubjectGroupName returns the SubjectGroupName field if non-nil, zero value otherwise.

### GetSubjectGroupNameOk

`func (o *OwnershipSourceCoverageRow) GetSubjectGroupNameOk() (*string, bool)`

GetSubjectGroupNameOk returns a tuple with the SubjectGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectGroupName

`func (o *OwnershipSourceCoverageRow) SetSubjectGroupName(v string)`

SetSubjectGroupName sets SubjectGroupName field to given value.

### HasSubjectGroupName

`func (o *OwnershipSourceCoverageRow) HasSubjectGroupName() bool`

HasSubjectGroupName returns a boolean if a field has been set.

### SetSubjectGroupNameNil

`func (o *OwnershipSourceCoverageRow) SetSubjectGroupNameNil(b bool)`

 SetSubjectGroupNameNil sets the value for SubjectGroupName to be an explicit nil

### UnsetSubjectGroupName
`func (o *OwnershipSourceCoverageRow) UnsetSubjectGroupName()`

UnsetSubjectGroupName ensures that no value is present for SubjectGroupName, not even an explicit nil
### GetSourceSystem

`func (o *OwnershipSourceCoverageRow) GetSourceSystem() string`

GetSourceSystem returns the SourceSystem field if non-nil, zero value otherwise.

### GetSourceSystemOk

`func (o *OwnershipSourceCoverageRow) GetSourceSystemOk() (*string, bool)`

GetSourceSystemOk returns a tuple with the SourceSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSystem

`func (o *OwnershipSourceCoverageRow) SetSourceSystem(v string)`

SetSourceSystem sets SourceSystem field to given value.

### HasSourceSystem

`func (o *OwnershipSourceCoverageRow) HasSourceSystem() bool`

HasSourceSystem returns a boolean if a field has been set.

### SetSourceSystemNil

`func (o *OwnershipSourceCoverageRow) SetSourceSystemNil(b bool)`

 SetSourceSystemNil sets the value for SourceSystem to be an explicit nil

### UnsetSourceSystem
`func (o *OwnershipSourceCoverageRow) UnsetSourceSystem()`

UnsetSourceSystem ensures that no value is present for SourceSystem, not even an explicit nil
### GetSourceLoadState

`func (o *OwnershipSourceCoverageRow) GetSourceLoadState() string`

GetSourceLoadState returns the SourceLoadState field if non-nil, zero value otherwise.

### GetSourceLoadStateOk

`func (o *OwnershipSourceCoverageRow) GetSourceLoadStateOk() (*string, bool)`

GetSourceLoadStateOk returns a tuple with the SourceLoadState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLoadState

`func (o *OwnershipSourceCoverageRow) SetSourceLoadState(v string)`

SetSourceLoadState sets SourceLoadState field to given value.

### HasSourceLoadState

`func (o *OwnershipSourceCoverageRow) HasSourceLoadState() bool`

HasSourceLoadState returns a boolean if a field has been set.

### SetSourceLoadStateNil

`func (o *OwnershipSourceCoverageRow) SetSourceLoadStateNil(b bool)`

 SetSourceLoadStateNil sets the value for SourceLoadState to be an explicit nil

### UnsetSourceLoadState
`func (o *OwnershipSourceCoverageRow) UnsetSourceLoadState()`

UnsetSourceLoadState ensures that no value is present for SourceLoadState, not even an explicit nil
### GetSourceDocumentsIngested

`func (o *OwnershipSourceCoverageRow) GetSourceDocumentsIngested() int64`

GetSourceDocumentsIngested returns the SourceDocumentsIngested field if non-nil, zero value otherwise.

### GetSourceDocumentsIngestedOk

`func (o *OwnershipSourceCoverageRow) GetSourceDocumentsIngestedOk() (*int64, bool)`

GetSourceDocumentsIngestedOk returns a tuple with the SourceDocumentsIngested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocumentsIngested

`func (o *OwnershipSourceCoverageRow) SetSourceDocumentsIngested(v int64)`

SetSourceDocumentsIngested sets SourceDocumentsIngested field to given value.

### HasSourceDocumentsIngested

`func (o *OwnershipSourceCoverageRow) HasSourceDocumentsIngested() bool`

HasSourceDocumentsIngested returns a boolean if a field has been set.

### SetSourceDocumentsIngestedNil

`func (o *OwnershipSourceCoverageRow) SetSourceDocumentsIngestedNil(b bool)`

 SetSourceDocumentsIngestedNil sets the value for SourceDocumentsIngested to be an explicit nil

### UnsetSourceDocumentsIngested
`func (o *OwnershipSourceCoverageRow) UnsetSourceDocumentsIngested()`

UnsetSourceDocumentsIngested ensures that no value is present for SourceDocumentsIngested, not even an explicit nil
### GetCoverageState

`func (o *OwnershipSourceCoverageRow) GetCoverageState() string`

GetCoverageState returns the CoverageState field if non-nil, zero value otherwise.

### GetCoverageStateOk

`func (o *OwnershipSourceCoverageRow) GetCoverageStateOk() (*string, bool)`

GetCoverageStateOk returns a tuple with the CoverageState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageState

`func (o *OwnershipSourceCoverageRow) SetCoverageState(v string)`

SetCoverageState sets CoverageState field to given value.

### HasCoverageState

`func (o *OwnershipSourceCoverageRow) HasCoverageState() bool`

HasCoverageState returns a boolean if a field has been set.

### SetCoverageStateNil

`func (o *OwnershipSourceCoverageRow) SetCoverageStateNil(b bool)`

 SetCoverageStateNil sets the value for CoverageState to be an explicit nil

### UnsetCoverageState
`func (o *OwnershipSourceCoverageRow) UnsetCoverageState()`

UnsetCoverageState ensures that no value is present for CoverageState, not even an explicit nil
### GetDocumentsIngested

`func (o *OwnershipSourceCoverageRow) GetDocumentsIngested() int64`

GetDocumentsIngested returns the DocumentsIngested field if non-nil, zero value otherwise.

### GetDocumentsIngestedOk

`func (o *OwnershipSourceCoverageRow) GetDocumentsIngestedOk() (*int64, bool)`

GetDocumentsIngestedOk returns a tuple with the DocumentsIngested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentsIngested

`func (o *OwnershipSourceCoverageRow) SetDocumentsIngested(v int64)`

SetDocumentsIngested sets DocumentsIngested field to given value.

### HasDocumentsIngested

`func (o *OwnershipSourceCoverageRow) HasDocumentsIngested() bool`

HasDocumentsIngested returns a boolean if a field has been set.

### SetDocumentsIngestedNil

`func (o *OwnershipSourceCoverageRow) SetDocumentsIngestedNil(b bool)`

 SetDocumentsIngestedNil sets the value for DocumentsIngested to be an explicit nil

### UnsetDocumentsIngested
`func (o *OwnershipSourceCoverageRow) UnsetDocumentsIngested()`

UnsetDocumentsIngested ensures that no value is present for DocumentsIngested, not even an explicit nil
### GetObservations

`func (o *OwnershipSourceCoverageRow) GetObservations() int64`

GetObservations returns the Observations field if non-nil, zero value otherwise.

### GetObservationsOk

`func (o *OwnershipSourceCoverageRow) GetObservationsOk() (*int64, bool)`

GetObservationsOk returns a tuple with the Observations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservations

`func (o *OwnershipSourceCoverageRow) SetObservations(v int64)`

SetObservations sets Observations field to given value.

### HasObservations

`func (o *OwnershipSourceCoverageRow) HasObservations() bool`

HasObservations returns a boolean if a field has been set.

### SetObservationsNil

`func (o *OwnershipSourceCoverageRow) SetObservationsNil(b bool)`

 SetObservationsNil sets the value for Observations to be an explicit nil

### UnsetObservations
`func (o *OwnershipSourceCoverageRow) UnsetObservations()`

UnsetObservations ensures that no value is present for Observations, not even an explicit nil
### GetCurrentEdges

`func (o *OwnershipSourceCoverageRow) GetCurrentEdges() int64`

GetCurrentEdges returns the CurrentEdges field if non-nil, zero value otherwise.

### GetCurrentEdgesOk

`func (o *OwnershipSourceCoverageRow) GetCurrentEdgesOk() (*int64, bool)`

GetCurrentEdgesOk returns a tuple with the CurrentEdges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentEdges

`func (o *OwnershipSourceCoverageRow) SetCurrentEdges(v int64)`

SetCurrentEdges sets CurrentEdges field to given value.

### HasCurrentEdges

`func (o *OwnershipSourceCoverageRow) HasCurrentEdges() bool`

HasCurrentEdges returns a boolean if a field has been set.

### SetCurrentEdgesNil

`func (o *OwnershipSourceCoverageRow) SetCurrentEdgesNil(b bool)`

 SetCurrentEdgesNil sets the value for CurrentEdges to be an explicit nil

### UnsetCurrentEdges
`func (o *OwnershipSourceCoverageRow) UnsetCurrentEdges()`

UnsetCurrentEdges ensures that no value is present for CurrentEdges, not even an explicit nil
### GetOwnersDistinct

`func (o *OwnershipSourceCoverageRow) GetOwnersDistinct() int64`

GetOwnersDistinct returns the OwnersDistinct field if non-nil, zero value otherwise.

### GetOwnersDistinctOk

`func (o *OwnershipSourceCoverageRow) GetOwnersDistinctOk() (*int64, bool)`

GetOwnersDistinctOk returns a tuple with the OwnersDistinct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnersDistinct

`func (o *OwnershipSourceCoverageRow) SetOwnersDistinct(v int64)`

SetOwnersDistinct sets OwnersDistinct field to given value.

### HasOwnersDistinct

`func (o *OwnershipSourceCoverageRow) HasOwnersDistinct() bool`

HasOwnersDistinct returns a boolean if a field has been set.

### SetOwnersDistinctNil

`func (o *OwnershipSourceCoverageRow) SetOwnersDistinctNil(b bool)`

 SetOwnersDistinctNil sets the value for OwnersDistinct to be an explicit nil

### UnsetOwnersDistinct
`func (o *OwnershipSourceCoverageRow) UnsetOwnersDistinct()`

UnsetOwnersDistinct ensures that no value is present for OwnersDistinct, not even an explicit nil
### GetOwnersUnresolved

`func (o *OwnershipSourceCoverageRow) GetOwnersUnresolved() int64`

GetOwnersUnresolved returns the OwnersUnresolved field if non-nil, zero value otherwise.

### GetOwnersUnresolvedOk

`func (o *OwnershipSourceCoverageRow) GetOwnersUnresolvedOk() (*int64, bool)`

GetOwnersUnresolvedOk returns a tuple with the OwnersUnresolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnersUnresolved

`func (o *OwnershipSourceCoverageRow) SetOwnersUnresolved(v int64)`

SetOwnersUnresolved sets OwnersUnresolved field to given value.

### HasOwnersUnresolved

`func (o *OwnershipSourceCoverageRow) HasOwnersUnresolved() bool`

HasOwnersUnresolved returns a boolean if a field has been set.

### SetOwnersUnresolvedNil

`func (o *OwnershipSourceCoverageRow) SetOwnersUnresolvedNil(b bool)`

 SetOwnersUnresolvedNil sets the value for OwnersUnresolved to be an explicit nil

### UnsetOwnersUnresolved
`func (o *OwnershipSourceCoverageRow) UnsetOwnersUnresolved()`

UnsetOwnersUnresolved ensures that no value is present for OwnersUnresolved, not even an explicit nil
### GetNoReportableOwnerStatements

`func (o *OwnershipSourceCoverageRow) GetNoReportableOwnerStatements() int64`

GetNoReportableOwnerStatements returns the NoReportableOwnerStatements field if non-nil, zero value otherwise.

### GetNoReportableOwnerStatementsOk

`func (o *OwnershipSourceCoverageRow) GetNoReportableOwnerStatementsOk() (*int64, bool)`

GetNoReportableOwnerStatementsOk returns a tuple with the NoReportableOwnerStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoReportableOwnerStatements

`func (o *OwnershipSourceCoverageRow) SetNoReportableOwnerStatements(v int64)`

SetNoReportableOwnerStatements sets NoReportableOwnerStatements field to given value.

### HasNoReportableOwnerStatements

`func (o *OwnershipSourceCoverageRow) HasNoReportableOwnerStatements() bool`

HasNoReportableOwnerStatements returns a boolean if a field has been set.

### SetNoReportableOwnerStatementsNil

`func (o *OwnershipSourceCoverageRow) SetNoReportableOwnerStatementsNil(b bool)`

 SetNoReportableOwnerStatementsNil sets the value for NoReportableOwnerStatements to be an explicit nil

### UnsetNoReportableOwnerStatements
`func (o *OwnershipSourceCoverageRow) UnsetNoReportableOwnerStatements()`

UnsetNoReportableOwnerStatements ensures that no value is present for NoReportableOwnerStatements, not even an explicit nil
### GetExtractionAbstained

`func (o *OwnershipSourceCoverageRow) GetExtractionAbstained() int64`

GetExtractionAbstained returns the ExtractionAbstained field if non-nil, zero value otherwise.

### GetExtractionAbstainedOk

`func (o *OwnershipSourceCoverageRow) GetExtractionAbstainedOk() (*int64, bool)`

GetExtractionAbstainedOk returns a tuple with the ExtractionAbstained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionAbstained

`func (o *OwnershipSourceCoverageRow) SetExtractionAbstained(v int64)`

SetExtractionAbstained sets ExtractionAbstained field to given value.

### HasExtractionAbstained

`func (o *OwnershipSourceCoverageRow) HasExtractionAbstained() bool`

HasExtractionAbstained returns a boolean if a field has been set.

### SetExtractionAbstainedNil

`func (o *OwnershipSourceCoverageRow) SetExtractionAbstainedNil(b bool)`

 SetExtractionAbstainedNil sets the value for ExtractionAbstained to be an explicit nil

### UnsetExtractionAbstained
`func (o *OwnershipSourceCoverageRow) UnsetExtractionAbstained()`

UnsetExtractionAbstained ensures that no value is present for ExtractionAbstained, not even an explicit nil
### GetLastFilingDate

`func (o *OwnershipSourceCoverageRow) GetLastFilingDate() string`

GetLastFilingDate returns the LastFilingDate field if non-nil, zero value otherwise.

### GetLastFilingDateOk

`func (o *OwnershipSourceCoverageRow) GetLastFilingDateOk() (*string, bool)`

GetLastFilingDateOk returns a tuple with the LastFilingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFilingDate

`func (o *OwnershipSourceCoverageRow) SetLastFilingDate(v string)`

SetLastFilingDate sets LastFilingDate field to given value.

### HasLastFilingDate

`func (o *OwnershipSourceCoverageRow) HasLastFilingDate() bool`

HasLastFilingDate returns a boolean if a field has been set.

### SetLastFilingDateNil

`func (o *OwnershipSourceCoverageRow) SetLastFilingDateNil(b bool)`

 SetLastFilingDateNil sets the value for LastFilingDate to be an explicit nil

### UnsetLastFilingDate
`func (o *OwnershipSourceCoverageRow) UnsetLastFilingDate()`

UnsetLastFilingDate ensures that no value is present for LastFilingDate, not even an explicit nil
### GetMaxAsOfDate

`func (o *OwnershipSourceCoverageRow) GetMaxAsOfDate() string`

GetMaxAsOfDate returns the MaxAsOfDate field if non-nil, zero value otherwise.

### GetMaxAsOfDateOk

`func (o *OwnershipSourceCoverageRow) GetMaxAsOfDateOk() (*string, bool)`

GetMaxAsOfDateOk returns a tuple with the MaxAsOfDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAsOfDate

`func (o *OwnershipSourceCoverageRow) SetMaxAsOfDate(v string)`

SetMaxAsOfDate sets MaxAsOfDate field to given value.

### HasMaxAsOfDate

`func (o *OwnershipSourceCoverageRow) HasMaxAsOfDate() bool`

HasMaxAsOfDate returns a boolean if a field has been set.

### SetMaxAsOfDateNil

`func (o *OwnershipSourceCoverageRow) SetMaxAsOfDateNil(b bool)`

 SetMaxAsOfDateNil sets the value for MaxAsOfDate to be an explicit nil

### UnsetMaxAsOfDate
`func (o *OwnershipSourceCoverageRow) UnsetMaxAsOfDate()`

UnsetMaxAsOfDate ensures that no value is present for MaxAsOfDate, not even an explicit nil
### GetSourceLicense

`func (o *OwnershipSourceCoverageRow) GetSourceLicense() string`

GetSourceLicense returns the SourceLicense field if non-nil, zero value otherwise.

### GetSourceLicenseOk

`func (o *OwnershipSourceCoverageRow) GetSourceLicenseOk() (*string, bool)`

GetSourceLicenseOk returns a tuple with the SourceLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLicense

`func (o *OwnershipSourceCoverageRow) SetSourceLicense(v string)`

SetSourceLicense sets SourceLicense field to given value.

### HasSourceLicense

`func (o *OwnershipSourceCoverageRow) HasSourceLicense() bool`

HasSourceLicense returns a boolean if a field has been set.

### SetSourceLicenseNil

`func (o *OwnershipSourceCoverageRow) SetSourceLicenseNil(b bool)`

 SetSourceLicenseNil sets the value for SourceLicense to be an explicit nil

### UnsetSourceLicense
`func (o *OwnershipSourceCoverageRow) UnsetSourceLicense()`

UnsetSourceLicense ensures that no value is present for SourceLicense, not even an explicit nil
### GetAttribution

`func (o *OwnershipSourceCoverageRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *OwnershipSourceCoverageRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *OwnershipSourceCoverageRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *OwnershipSourceCoverageRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *OwnershipSourceCoverageRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *OwnershipSourceCoverageRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


