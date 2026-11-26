# InvestmentGraphNodesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | Pointer to **NullableString** | The node key: an investor id (adv:, agency:, lei:, investor:), a company&#39;s organization id (the resolved parent), or investor_site: followed by the normalized printed name for an unadmitted website-only company. | [optional] 
**NodeKind** | Pointer to **NullableString** | investor or company. | [optional] 
**DisplayName** | Pointer to **NullableString** | The investor&#39;s legal name, the company&#39;s resolved parent name, or for an unadmitted website-only company the lowest name as printed on the investors&#39; pages. | [optional] 
**InvestorType** | Pointer to **NullableString** | venture_capital, private_equity, hedge_fund, government_agency or other. Null for a company node. | [optional] 
**InvestorId** | Pointer to **NullableString** | The investor key for an investor node (the same as node_id), null for a company node. | [optional] 
**OrgId** | Pointer to **NullableString** | Legal-entity level: a company node&#39;s organization (the same as node_id), or an investor node&#39;s own organization resolved to its parent. Null for an unadmitted website-only company and for an investor with no organization. | [optional] 
**CompanyId** | Pointer to **NullableString** | Company level: the company record of org_id (null where unresolved). | [optional] 
**CompanyName** | Pointer to **NullableString** | Display name of the company. | [optional] 
**GroupId** | Pointer to **NullableString** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**GroupName** | Pointer to **NullableString** | Display name of the group. | [optional] 
**SameEntityNodeId** | Pointer to **NullableString** | The other node of one real-world entity: an investor that is also a company with edges points at its company node and the reverse. An annotation, not an edge. Null when there is none. | [optional] 
**NodeResolution** | Pointer to **NullableString** | spine: the id is an organization or investor id. unadmitted_listing: a website-only company the organization spine has not admitted, keyed on its printed name (same-printed-name companies merge into one node; the id will change on admission). | [optional] 
**EdgeCount** | Pointer to **NullableInt32** | Stored edge rows incident to the node, as the sum of out_edge_count and in_edge_count. A co_invested edge is stored in both directions, so a co-investor counts it once in each. | [optional] 
**OutEdgeCount** | Pointer to **NullableInt32** | Stored edge rows that start at the node. | [optional] 
**InEdgeCount** | Pointer to **NullableInt32** | Stored edge rows that end at the node. | [optional] 
**NodeSourceCount** | Pointer to **NullableInt32** | Distinct sources among the node&#39;s edges. | [optional] 
**Sources** | Pointer to **[]string** | The sources that state the node&#39;s edges. | [optional] 
**UnlinkedCompanyCount** | Pointer to **NullableInt32** | For an investor node, its portfolio rows that cannot become a node (a company known only by an SEC filer number, an organization with no resolved parent, or a listing that is page chrome, a section heading, the investor&#39;s own name or a name a curator marked as not a company); shown as a count and never guessed at. 0 for a company node. | [optional] 
**ListedNames** | Pointer to **[]string** | For a company node, the distinct names investors printed for it on their portfolio pages, sorted; empty for an investor node. A name that differs from display_name is how the investors spell a company the registry names otherwise (Scientist.com, listed for Assay Depot Inc). A label, not a statement of identity. | [optional] 

## Methods

### NewInvestmentGraphNodesRow

`func NewInvestmentGraphNodesRow() *InvestmentGraphNodesRow`

NewInvestmentGraphNodesRow instantiates a new InvestmentGraphNodesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestmentGraphNodesRowWithDefaults

`func NewInvestmentGraphNodesRowWithDefaults() *InvestmentGraphNodesRow`

NewInvestmentGraphNodesRowWithDefaults instantiates a new InvestmentGraphNodesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *InvestmentGraphNodesRow) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *InvestmentGraphNodesRow) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *InvestmentGraphNodesRow) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *InvestmentGraphNodesRow) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### SetNodeIdNil

`func (o *InvestmentGraphNodesRow) SetNodeIdNil(b bool)`

 SetNodeIdNil sets the value for NodeId to be an explicit nil

### UnsetNodeId
`func (o *InvestmentGraphNodesRow) UnsetNodeId()`

UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
### GetNodeKind

`func (o *InvestmentGraphNodesRow) GetNodeKind() string`

GetNodeKind returns the NodeKind field if non-nil, zero value otherwise.

### GetNodeKindOk

`func (o *InvestmentGraphNodesRow) GetNodeKindOk() (*string, bool)`

GetNodeKindOk returns a tuple with the NodeKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeKind

`func (o *InvestmentGraphNodesRow) SetNodeKind(v string)`

SetNodeKind sets NodeKind field to given value.

### HasNodeKind

`func (o *InvestmentGraphNodesRow) HasNodeKind() bool`

HasNodeKind returns a boolean if a field has been set.

### SetNodeKindNil

`func (o *InvestmentGraphNodesRow) SetNodeKindNil(b bool)`

 SetNodeKindNil sets the value for NodeKind to be an explicit nil

### UnsetNodeKind
`func (o *InvestmentGraphNodesRow) UnsetNodeKind()`

UnsetNodeKind ensures that no value is present for NodeKind, not even an explicit nil
### GetDisplayName

`func (o *InvestmentGraphNodesRow) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InvestmentGraphNodesRow) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InvestmentGraphNodesRow) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InvestmentGraphNodesRow) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *InvestmentGraphNodesRow) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *InvestmentGraphNodesRow) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetInvestorType

`func (o *InvestmentGraphNodesRow) GetInvestorType() string`

GetInvestorType returns the InvestorType field if non-nil, zero value otherwise.

### GetInvestorTypeOk

`func (o *InvestmentGraphNodesRow) GetInvestorTypeOk() (*string, bool)`

GetInvestorTypeOk returns a tuple with the InvestorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestorType

`func (o *InvestmentGraphNodesRow) SetInvestorType(v string)`

SetInvestorType sets InvestorType field to given value.

### HasInvestorType

`func (o *InvestmentGraphNodesRow) HasInvestorType() bool`

HasInvestorType returns a boolean if a field has been set.

### SetInvestorTypeNil

`func (o *InvestmentGraphNodesRow) SetInvestorTypeNil(b bool)`

 SetInvestorTypeNil sets the value for InvestorType to be an explicit nil

### UnsetInvestorType
`func (o *InvestmentGraphNodesRow) UnsetInvestorType()`

UnsetInvestorType ensures that no value is present for InvestorType, not even an explicit nil
### GetInvestorId

`func (o *InvestmentGraphNodesRow) GetInvestorId() string`

GetInvestorId returns the InvestorId field if non-nil, zero value otherwise.

### GetInvestorIdOk

`func (o *InvestmentGraphNodesRow) GetInvestorIdOk() (*string, bool)`

GetInvestorIdOk returns a tuple with the InvestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestorId

`func (o *InvestmentGraphNodesRow) SetInvestorId(v string)`

SetInvestorId sets InvestorId field to given value.

### HasInvestorId

`func (o *InvestmentGraphNodesRow) HasInvestorId() bool`

HasInvestorId returns a boolean if a field has been set.

### SetInvestorIdNil

`func (o *InvestmentGraphNodesRow) SetInvestorIdNil(b bool)`

 SetInvestorIdNil sets the value for InvestorId to be an explicit nil

### UnsetInvestorId
`func (o *InvestmentGraphNodesRow) UnsetInvestorId()`

UnsetInvestorId ensures that no value is present for InvestorId, not even an explicit nil
### GetOrgId

`func (o *InvestmentGraphNodesRow) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *InvestmentGraphNodesRow) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *InvestmentGraphNodesRow) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *InvestmentGraphNodesRow) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *InvestmentGraphNodesRow) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *InvestmentGraphNodesRow) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetCompanyId

`func (o *InvestmentGraphNodesRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *InvestmentGraphNodesRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *InvestmentGraphNodesRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *InvestmentGraphNodesRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *InvestmentGraphNodesRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *InvestmentGraphNodesRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *InvestmentGraphNodesRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *InvestmentGraphNodesRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *InvestmentGraphNodesRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *InvestmentGraphNodesRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *InvestmentGraphNodesRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *InvestmentGraphNodesRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetGroupId

`func (o *InvestmentGraphNodesRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *InvestmentGraphNodesRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *InvestmentGraphNodesRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *InvestmentGraphNodesRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *InvestmentGraphNodesRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *InvestmentGraphNodesRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *InvestmentGraphNodesRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *InvestmentGraphNodesRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *InvestmentGraphNodesRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *InvestmentGraphNodesRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *InvestmentGraphNodesRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *InvestmentGraphNodesRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetSameEntityNodeId

`func (o *InvestmentGraphNodesRow) GetSameEntityNodeId() string`

GetSameEntityNodeId returns the SameEntityNodeId field if non-nil, zero value otherwise.

### GetSameEntityNodeIdOk

`func (o *InvestmentGraphNodesRow) GetSameEntityNodeIdOk() (*string, bool)`

GetSameEntityNodeIdOk returns a tuple with the SameEntityNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSameEntityNodeId

`func (o *InvestmentGraphNodesRow) SetSameEntityNodeId(v string)`

SetSameEntityNodeId sets SameEntityNodeId field to given value.

### HasSameEntityNodeId

`func (o *InvestmentGraphNodesRow) HasSameEntityNodeId() bool`

HasSameEntityNodeId returns a boolean if a field has been set.

### SetSameEntityNodeIdNil

`func (o *InvestmentGraphNodesRow) SetSameEntityNodeIdNil(b bool)`

 SetSameEntityNodeIdNil sets the value for SameEntityNodeId to be an explicit nil

### UnsetSameEntityNodeId
`func (o *InvestmentGraphNodesRow) UnsetSameEntityNodeId()`

UnsetSameEntityNodeId ensures that no value is present for SameEntityNodeId, not even an explicit nil
### GetNodeResolution

`func (o *InvestmentGraphNodesRow) GetNodeResolution() string`

GetNodeResolution returns the NodeResolution field if non-nil, zero value otherwise.

### GetNodeResolutionOk

`func (o *InvestmentGraphNodesRow) GetNodeResolutionOk() (*string, bool)`

GetNodeResolutionOk returns a tuple with the NodeResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeResolution

`func (o *InvestmentGraphNodesRow) SetNodeResolution(v string)`

SetNodeResolution sets NodeResolution field to given value.

### HasNodeResolution

`func (o *InvestmentGraphNodesRow) HasNodeResolution() bool`

HasNodeResolution returns a boolean if a field has been set.

### SetNodeResolutionNil

`func (o *InvestmentGraphNodesRow) SetNodeResolutionNil(b bool)`

 SetNodeResolutionNil sets the value for NodeResolution to be an explicit nil

### UnsetNodeResolution
`func (o *InvestmentGraphNodesRow) UnsetNodeResolution()`

UnsetNodeResolution ensures that no value is present for NodeResolution, not even an explicit nil
### GetEdgeCount

`func (o *InvestmentGraphNodesRow) GetEdgeCount() int32`

GetEdgeCount returns the EdgeCount field if non-nil, zero value otherwise.

### GetEdgeCountOk

`func (o *InvestmentGraphNodesRow) GetEdgeCountOk() (*int32, bool)`

GetEdgeCountOk returns a tuple with the EdgeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeCount

`func (o *InvestmentGraphNodesRow) SetEdgeCount(v int32)`

SetEdgeCount sets EdgeCount field to given value.

### HasEdgeCount

`func (o *InvestmentGraphNodesRow) HasEdgeCount() bool`

HasEdgeCount returns a boolean if a field has been set.

### SetEdgeCountNil

`func (o *InvestmentGraphNodesRow) SetEdgeCountNil(b bool)`

 SetEdgeCountNil sets the value for EdgeCount to be an explicit nil

### UnsetEdgeCount
`func (o *InvestmentGraphNodesRow) UnsetEdgeCount()`

UnsetEdgeCount ensures that no value is present for EdgeCount, not even an explicit nil
### GetOutEdgeCount

`func (o *InvestmentGraphNodesRow) GetOutEdgeCount() int32`

GetOutEdgeCount returns the OutEdgeCount field if non-nil, zero value otherwise.

### GetOutEdgeCountOk

`func (o *InvestmentGraphNodesRow) GetOutEdgeCountOk() (*int32, bool)`

GetOutEdgeCountOk returns a tuple with the OutEdgeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutEdgeCount

`func (o *InvestmentGraphNodesRow) SetOutEdgeCount(v int32)`

SetOutEdgeCount sets OutEdgeCount field to given value.

### HasOutEdgeCount

`func (o *InvestmentGraphNodesRow) HasOutEdgeCount() bool`

HasOutEdgeCount returns a boolean if a field has been set.

### SetOutEdgeCountNil

`func (o *InvestmentGraphNodesRow) SetOutEdgeCountNil(b bool)`

 SetOutEdgeCountNil sets the value for OutEdgeCount to be an explicit nil

### UnsetOutEdgeCount
`func (o *InvestmentGraphNodesRow) UnsetOutEdgeCount()`

UnsetOutEdgeCount ensures that no value is present for OutEdgeCount, not even an explicit nil
### GetInEdgeCount

`func (o *InvestmentGraphNodesRow) GetInEdgeCount() int32`

GetInEdgeCount returns the InEdgeCount field if non-nil, zero value otherwise.

### GetInEdgeCountOk

`func (o *InvestmentGraphNodesRow) GetInEdgeCountOk() (*int32, bool)`

GetInEdgeCountOk returns a tuple with the InEdgeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInEdgeCount

`func (o *InvestmentGraphNodesRow) SetInEdgeCount(v int32)`

SetInEdgeCount sets InEdgeCount field to given value.

### HasInEdgeCount

`func (o *InvestmentGraphNodesRow) HasInEdgeCount() bool`

HasInEdgeCount returns a boolean if a field has been set.

### SetInEdgeCountNil

`func (o *InvestmentGraphNodesRow) SetInEdgeCountNil(b bool)`

 SetInEdgeCountNil sets the value for InEdgeCount to be an explicit nil

### UnsetInEdgeCount
`func (o *InvestmentGraphNodesRow) UnsetInEdgeCount()`

UnsetInEdgeCount ensures that no value is present for InEdgeCount, not even an explicit nil
### GetNodeSourceCount

`func (o *InvestmentGraphNodesRow) GetNodeSourceCount() int32`

GetNodeSourceCount returns the NodeSourceCount field if non-nil, zero value otherwise.

### GetNodeSourceCountOk

`func (o *InvestmentGraphNodesRow) GetNodeSourceCountOk() (*int32, bool)`

GetNodeSourceCountOk returns a tuple with the NodeSourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSourceCount

`func (o *InvestmentGraphNodesRow) SetNodeSourceCount(v int32)`

SetNodeSourceCount sets NodeSourceCount field to given value.

### HasNodeSourceCount

`func (o *InvestmentGraphNodesRow) HasNodeSourceCount() bool`

HasNodeSourceCount returns a boolean if a field has been set.

### SetNodeSourceCountNil

`func (o *InvestmentGraphNodesRow) SetNodeSourceCountNil(b bool)`

 SetNodeSourceCountNil sets the value for NodeSourceCount to be an explicit nil

### UnsetNodeSourceCount
`func (o *InvestmentGraphNodesRow) UnsetNodeSourceCount()`

UnsetNodeSourceCount ensures that no value is present for NodeSourceCount, not even an explicit nil
### GetSources

`func (o *InvestmentGraphNodesRow) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *InvestmentGraphNodesRow) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *InvestmentGraphNodesRow) SetSources(v []string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *InvestmentGraphNodesRow) HasSources() bool`

HasSources returns a boolean if a field has been set.

### SetSourcesNil

`func (o *InvestmentGraphNodesRow) SetSourcesNil(b bool)`

 SetSourcesNil sets the value for Sources to be an explicit nil

### UnsetSources
`func (o *InvestmentGraphNodesRow) UnsetSources()`

UnsetSources ensures that no value is present for Sources, not even an explicit nil
### GetUnlinkedCompanyCount

`func (o *InvestmentGraphNodesRow) GetUnlinkedCompanyCount() int32`

GetUnlinkedCompanyCount returns the UnlinkedCompanyCount field if non-nil, zero value otherwise.

### GetUnlinkedCompanyCountOk

`func (o *InvestmentGraphNodesRow) GetUnlinkedCompanyCountOk() (*int32, bool)`

GetUnlinkedCompanyCountOk returns a tuple with the UnlinkedCompanyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlinkedCompanyCount

`func (o *InvestmentGraphNodesRow) SetUnlinkedCompanyCount(v int32)`

SetUnlinkedCompanyCount sets UnlinkedCompanyCount field to given value.

### HasUnlinkedCompanyCount

`func (o *InvestmentGraphNodesRow) HasUnlinkedCompanyCount() bool`

HasUnlinkedCompanyCount returns a boolean if a field has been set.

### SetUnlinkedCompanyCountNil

`func (o *InvestmentGraphNodesRow) SetUnlinkedCompanyCountNil(b bool)`

 SetUnlinkedCompanyCountNil sets the value for UnlinkedCompanyCount to be an explicit nil

### UnsetUnlinkedCompanyCount
`func (o *InvestmentGraphNodesRow) UnsetUnlinkedCompanyCount()`

UnsetUnlinkedCompanyCount ensures that no value is present for UnlinkedCompanyCount, not even an explicit nil
### GetListedNames

`func (o *InvestmentGraphNodesRow) GetListedNames() []string`

GetListedNames returns the ListedNames field if non-nil, zero value otherwise.

### GetListedNamesOk

`func (o *InvestmentGraphNodesRow) GetListedNamesOk() (*[]string, bool)`

GetListedNamesOk returns a tuple with the ListedNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListedNames

`func (o *InvestmentGraphNodesRow) SetListedNames(v []string)`

SetListedNames sets ListedNames field to given value.

### HasListedNames

`func (o *InvestmentGraphNodesRow) HasListedNames() bool`

HasListedNames returns a boolean if a field has been set.

### SetListedNamesNil

`func (o *InvestmentGraphNodesRow) SetListedNamesNil(b bool)`

 SetListedNamesNil sets the value for ListedNames to be an explicit nil

### UnsetListedNames
`func (o *InvestmentGraphNodesRow) UnsetListedNames()`

UnsetListedNames ensures that no value is present for ListedNames, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


