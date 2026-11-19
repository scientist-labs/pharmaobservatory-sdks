# InvestmentGraphNodeEdgesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | Pointer to **NullableString** | The endpoint this row is read from: the from node when direction is out, the to node when it is in. | [optional] 
**Direction** | Pointer to **NullableString** | out when node_id is the edge&#39;s from node, in when it is the to node. | [optional] 
**EdgeKey** | Pointer to **NullableString** | The stable handle of the edge, the same on its two rows. | [optional] 
**NodeKind** | Pointer to **NullableString** | investor or company: the kind of node_id. | [optional] 
**InvestorId** | Pointer to **NullableString** | The investor key when node_id is an investor, null otherwise. | [optional] 
**NodeName** | Pointer to **NullableString** | Display name of node_id. | [optional] 
**OrgId** | Pointer to **NullableString** | Legal-entity level: node_id&#39;s organization (null for an investor with none and for an unadmitted website-only company). | [optional] 
**CompanyId** | Pointer to **NullableString** | Company level: node_id&#39;s company record (null where unresolved). | [optional] 
**CompanyName** | Pointer to **NullableString** | Display name of the company. | [optional] 
**GroupId** | Pointer to **NullableString** | Group level: the company&#39;s current ultimate parent; a root&#39;s group is itself. | [optional] 
**GroupName** | Pointer to **NullableString** | Display name of the group. | [optional] 
**CounterpartNodeId** | Pointer to **NullableString** | The node at the other end of the edge. | [optional] 
**CounterpartKind** | Pointer to **NullableString** | investor or company: the kind of counterpart_node_id. | [optional] 
**CounterpartName** | Pointer to **NullableString** | Display name of the counterpart node. | [optional] 
**FromNodeId** | Pointer to **NullableString** | The node the edge starts at: an investor id or a company id (the resolved parent organization, or investor_site: and the normalized name for a website-only company the organization spine has not admitted). | [optional] 
**ToNodeId** | Pointer to **NullableString** | The node the edge ends at. Same id forms as from_node_id. invested_in goes from an investor to a company; co_invested is stored in both directions between investors. | [optional] 
**EdgeType** | Pointer to **NullableString** | invested_in, co_invested, acquired, merged_with or licensed. acquired, merged_with and licensed have no rows until the M&amp;A events dataset lands. There is no exited type: a listing that disappears from a website is never an end. | [optional] 
**RelationshipType** | Pointer to **NullableString** | The lane&#39;s own word for the relationship, so a public-market holding is never read as a venture round: portfolio_listed (investor website), holder_5pct and selling_stockholder (offering holder tables), institutional_position (13F), beneficial_owner (13D/G), co_appearance (co_invested). | [optional] 
**Source** | Pointer to **NullableString** | The lane that states the edge: investor_site, sec_edgar_holder_table (also the lane of the derived co_invested edges), sec_13f, sec_13dg, sec_form_d, sec_8k, company_press_release or reconciled. sec_form_d, sec_8k and company_press_release have no rows yet; reconciled carries the press-release and SEC deal metadata for the pairs it enriches and is the preferred lane of such a pair. | [optional] 
**SourceRef** | Pointer to **NullableString** | The evidence key within the lane: the portfolio company key for the website and holder-table lanes, the ownership observation id for 13F and 13D/G, and the literal co_invested_pairs for a derived co-investment. | [optional] 
**SourceUrl** | Pointer to **NullableString** | The page the edge was read from. Null for a derived co_invested edge, which has no page of its own. | [optional] 
**EventDate** | Pointer to **NullableString** | The date the lane states, of the kind in event_date_kind. Null exactly when event_date_kind is unknown. | [optional] 
**EventDateKind** | Pointer to **NullableString** | What event_date is: transaction (a date the lane states for the deal itself), announcement, filing (the filing date), as_of (a report period end, as 13F states), form_d_filing (the filing date of a Form D offering by the company, never an investment date), first_observed (the first crawl that saw a website listing, never an investment date) or unknown (no date, event_date null). | [optional] 
**ObservedThrough** | Pointer to **NullableString** | The latest date the relationship was still observed in the lane (the last crawl, the last filing). Null when the lane gives none. Not an end date. | [optional] 
**EventStage** | Pointer to **NullableString** | Where a deal stands (announced, signed, closed, completed or terminated), so an announced acquisition is never drawn as a completed one. Null for every lane today. | [optional] 
**RoundLabel** | Pointer to **NullableString** | The funding round named by the company&#39;s own press release (reconciled lane); null for every lane that does not state one. | [optional] 
**AmountValue** | Pointer to **NullableFloat32** | The ROUND amount in amount_currency, never the investor&#39;s own cheque (reconciled lane): the press-release round total, else the amount sold in a Form D offering. Null for every lane that does not state one. | [optional] 
**AmountCurrency** | Pointer to **NullableString** | The currency of amount_value as the source states it, never defaulted to USD. Null while amount_value is null. | [optional] 
**IsLead** | Pointer to **NullableBool** | Whether the investor led or co-led the round per the company&#39;s press release (reconciled lane): true for lead or co-lead, false for a named participant, null where the role is unstated or the lane does not state it. | [optional] 
**SharedCompanyCount** | Pointer to **NullableInt32** | For a derived co_invested edge, the companies the two investors are both named against in offering filings. Null for every other edge. | [optional] 
**IsDerived** | Pointer to **NullableBool** | True for an edge computed from other rows (co_invested), false for an edge a source states. | [optional] 
**PairKey** | Pointer to **NullableString** | from_node_id, to_node_id and edge_type joined with a pipe: every lane that states the same relationship shares one pair_key. | [optional] 
**LaneRank** | Pointer to **NullableInt32** | The lane&#39;s precedence when several lanes state one pair (lower wins): 13D/G 1, 13F 2, offering filings and derived co-investment 3, investor website 7. | [optional] 
**IsPreferred** | Pointer to **NullableBool** | True on exactly one row per pair_key, the lowest lane_rank (then the latest event_date): the row a picture draws. The other rows stay as evidence. | [optional] 
**SourceLicense** | Pointer to **NullableString** | The license class of the lane: public_domain_us_government or first_party_website_facts. | [optional] 
**Attribution** | Pointer to **NullableString** | The plain credit line for the lane. Show it with the edge for every reader. | [optional] 

## Methods

### NewInvestmentGraphNodeEdgesRow

`func NewInvestmentGraphNodeEdgesRow() *InvestmentGraphNodeEdgesRow`

NewInvestmentGraphNodeEdgesRow instantiates a new InvestmentGraphNodeEdgesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestmentGraphNodeEdgesRowWithDefaults

`func NewInvestmentGraphNodeEdgesRowWithDefaults() *InvestmentGraphNodeEdgesRow`

NewInvestmentGraphNodeEdgesRowWithDefaults instantiates a new InvestmentGraphNodeEdgesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *InvestmentGraphNodeEdgesRow) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *InvestmentGraphNodeEdgesRow) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *InvestmentGraphNodeEdgesRow) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *InvestmentGraphNodeEdgesRow) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### SetNodeIdNil

`func (o *InvestmentGraphNodeEdgesRow) SetNodeIdNil(b bool)`

 SetNodeIdNil sets the value for NodeId to be an explicit nil

### UnsetNodeId
`func (o *InvestmentGraphNodeEdgesRow) UnsetNodeId()`

UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
### GetDirection

`func (o *InvestmentGraphNodeEdgesRow) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *InvestmentGraphNodeEdgesRow) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *InvestmentGraphNodeEdgesRow) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *InvestmentGraphNodeEdgesRow) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *InvestmentGraphNodeEdgesRow) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *InvestmentGraphNodeEdgesRow) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil
### GetEdgeKey

`func (o *InvestmentGraphNodeEdgesRow) GetEdgeKey() string`

GetEdgeKey returns the EdgeKey field if non-nil, zero value otherwise.

### GetEdgeKeyOk

`func (o *InvestmentGraphNodeEdgesRow) GetEdgeKeyOk() (*string, bool)`

GetEdgeKeyOk returns a tuple with the EdgeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeKey

`func (o *InvestmentGraphNodeEdgesRow) SetEdgeKey(v string)`

SetEdgeKey sets EdgeKey field to given value.

### HasEdgeKey

`func (o *InvestmentGraphNodeEdgesRow) HasEdgeKey() bool`

HasEdgeKey returns a boolean if a field has been set.

### SetEdgeKeyNil

`func (o *InvestmentGraphNodeEdgesRow) SetEdgeKeyNil(b bool)`

 SetEdgeKeyNil sets the value for EdgeKey to be an explicit nil

### UnsetEdgeKey
`func (o *InvestmentGraphNodeEdgesRow) UnsetEdgeKey()`

UnsetEdgeKey ensures that no value is present for EdgeKey, not even an explicit nil
### GetNodeKind

`func (o *InvestmentGraphNodeEdgesRow) GetNodeKind() string`

GetNodeKind returns the NodeKind field if non-nil, zero value otherwise.

### GetNodeKindOk

`func (o *InvestmentGraphNodeEdgesRow) GetNodeKindOk() (*string, bool)`

GetNodeKindOk returns a tuple with the NodeKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeKind

`func (o *InvestmentGraphNodeEdgesRow) SetNodeKind(v string)`

SetNodeKind sets NodeKind field to given value.

### HasNodeKind

`func (o *InvestmentGraphNodeEdgesRow) HasNodeKind() bool`

HasNodeKind returns a boolean if a field has been set.

### SetNodeKindNil

`func (o *InvestmentGraphNodeEdgesRow) SetNodeKindNil(b bool)`

 SetNodeKindNil sets the value for NodeKind to be an explicit nil

### UnsetNodeKind
`func (o *InvestmentGraphNodeEdgesRow) UnsetNodeKind()`

UnsetNodeKind ensures that no value is present for NodeKind, not even an explicit nil
### GetInvestorId

`func (o *InvestmentGraphNodeEdgesRow) GetInvestorId() string`

GetInvestorId returns the InvestorId field if non-nil, zero value otherwise.

### GetInvestorIdOk

`func (o *InvestmentGraphNodeEdgesRow) GetInvestorIdOk() (*string, bool)`

GetInvestorIdOk returns a tuple with the InvestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestorId

`func (o *InvestmentGraphNodeEdgesRow) SetInvestorId(v string)`

SetInvestorId sets InvestorId field to given value.

### HasInvestorId

`func (o *InvestmentGraphNodeEdgesRow) HasInvestorId() bool`

HasInvestorId returns a boolean if a field has been set.

### SetInvestorIdNil

`func (o *InvestmentGraphNodeEdgesRow) SetInvestorIdNil(b bool)`

 SetInvestorIdNil sets the value for InvestorId to be an explicit nil

### UnsetInvestorId
`func (o *InvestmentGraphNodeEdgesRow) UnsetInvestorId()`

UnsetInvestorId ensures that no value is present for InvestorId, not even an explicit nil
### GetNodeName

`func (o *InvestmentGraphNodeEdgesRow) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *InvestmentGraphNodeEdgesRow) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *InvestmentGraphNodeEdgesRow) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *InvestmentGraphNodeEdgesRow) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### SetNodeNameNil

`func (o *InvestmentGraphNodeEdgesRow) SetNodeNameNil(b bool)`

 SetNodeNameNil sets the value for NodeName to be an explicit nil

### UnsetNodeName
`func (o *InvestmentGraphNodeEdgesRow) UnsetNodeName()`

UnsetNodeName ensures that no value is present for NodeName, not even an explicit nil
### GetOrgId

`func (o *InvestmentGraphNodeEdgesRow) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *InvestmentGraphNodeEdgesRow) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *InvestmentGraphNodeEdgesRow) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *InvestmentGraphNodeEdgesRow) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *InvestmentGraphNodeEdgesRow) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *InvestmentGraphNodeEdgesRow) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetCompanyId

`func (o *InvestmentGraphNodeEdgesRow) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *InvestmentGraphNodeEdgesRow) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *InvestmentGraphNodeEdgesRow) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *InvestmentGraphNodeEdgesRow) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *InvestmentGraphNodeEdgesRow) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *InvestmentGraphNodeEdgesRow) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCompanyName

`func (o *InvestmentGraphNodeEdgesRow) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *InvestmentGraphNodeEdgesRow) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *InvestmentGraphNodeEdgesRow) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *InvestmentGraphNodeEdgesRow) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *InvestmentGraphNodeEdgesRow) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *InvestmentGraphNodeEdgesRow) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetGroupId

`func (o *InvestmentGraphNodeEdgesRow) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *InvestmentGraphNodeEdgesRow) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *InvestmentGraphNodeEdgesRow) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *InvestmentGraphNodeEdgesRow) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *InvestmentGraphNodeEdgesRow) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *InvestmentGraphNodeEdgesRow) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *InvestmentGraphNodeEdgesRow) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *InvestmentGraphNodeEdgesRow) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *InvestmentGraphNodeEdgesRow) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *InvestmentGraphNodeEdgesRow) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *InvestmentGraphNodeEdgesRow) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *InvestmentGraphNodeEdgesRow) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetCounterpartNodeId

`func (o *InvestmentGraphNodeEdgesRow) GetCounterpartNodeId() string`

GetCounterpartNodeId returns the CounterpartNodeId field if non-nil, zero value otherwise.

### GetCounterpartNodeIdOk

`func (o *InvestmentGraphNodeEdgesRow) GetCounterpartNodeIdOk() (*string, bool)`

GetCounterpartNodeIdOk returns a tuple with the CounterpartNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartNodeId

`func (o *InvestmentGraphNodeEdgesRow) SetCounterpartNodeId(v string)`

SetCounterpartNodeId sets CounterpartNodeId field to given value.

### HasCounterpartNodeId

`func (o *InvestmentGraphNodeEdgesRow) HasCounterpartNodeId() bool`

HasCounterpartNodeId returns a boolean if a field has been set.

### SetCounterpartNodeIdNil

`func (o *InvestmentGraphNodeEdgesRow) SetCounterpartNodeIdNil(b bool)`

 SetCounterpartNodeIdNil sets the value for CounterpartNodeId to be an explicit nil

### UnsetCounterpartNodeId
`func (o *InvestmentGraphNodeEdgesRow) UnsetCounterpartNodeId()`

UnsetCounterpartNodeId ensures that no value is present for CounterpartNodeId, not even an explicit nil
### GetCounterpartKind

`func (o *InvestmentGraphNodeEdgesRow) GetCounterpartKind() string`

GetCounterpartKind returns the CounterpartKind field if non-nil, zero value otherwise.

### GetCounterpartKindOk

`func (o *InvestmentGraphNodeEdgesRow) GetCounterpartKindOk() (*string, bool)`

GetCounterpartKindOk returns a tuple with the CounterpartKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartKind

`func (o *InvestmentGraphNodeEdgesRow) SetCounterpartKind(v string)`

SetCounterpartKind sets CounterpartKind field to given value.

### HasCounterpartKind

`func (o *InvestmentGraphNodeEdgesRow) HasCounterpartKind() bool`

HasCounterpartKind returns a boolean if a field has been set.

### SetCounterpartKindNil

`func (o *InvestmentGraphNodeEdgesRow) SetCounterpartKindNil(b bool)`

 SetCounterpartKindNil sets the value for CounterpartKind to be an explicit nil

### UnsetCounterpartKind
`func (o *InvestmentGraphNodeEdgesRow) UnsetCounterpartKind()`

UnsetCounterpartKind ensures that no value is present for CounterpartKind, not even an explicit nil
### GetCounterpartName

`func (o *InvestmentGraphNodeEdgesRow) GetCounterpartName() string`

GetCounterpartName returns the CounterpartName field if non-nil, zero value otherwise.

### GetCounterpartNameOk

`func (o *InvestmentGraphNodeEdgesRow) GetCounterpartNameOk() (*string, bool)`

GetCounterpartNameOk returns a tuple with the CounterpartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartName

`func (o *InvestmentGraphNodeEdgesRow) SetCounterpartName(v string)`

SetCounterpartName sets CounterpartName field to given value.

### HasCounterpartName

`func (o *InvestmentGraphNodeEdgesRow) HasCounterpartName() bool`

HasCounterpartName returns a boolean if a field has been set.

### SetCounterpartNameNil

`func (o *InvestmentGraphNodeEdgesRow) SetCounterpartNameNil(b bool)`

 SetCounterpartNameNil sets the value for CounterpartName to be an explicit nil

### UnsetCounterpartName
`func (o *InvestmentGraphNodeEdgesRow) UnsetCounterpartName()`

UnsetCounterpartName ensures that no value is present for CounterpartName, not even an explicit nil
### GetFromNodeId

`func (o *InvestmentGraphNodeEdgesRow) GetFromNodeId() string`

GetFromNodeId returns the FromNodeId field if non-nil, zero value otherwise.

### GetFromNodeIdOk

`func (o *InvestmentGraphNodeEdgesRow) GetFromNodeIdOk() (*string, bool)`

GetFromNodeIdOk returns a tuple with the FromNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromNodeId

`func (o *InvestmentGraphNodeEdgesRow) SetFromNodeId(v string)`

SetFromNodeId sets FromNodeId field to given value.

### HasFromNodeId

`func (o *InvestmentGraphNodeEdgesRow) HasFromNodeId() bool`

HasFromNodeId returns a boolean if a field has been set.

### SetFromNodeIdNil

`func (o *InvestmentGraphNodeEdgesRow) SetFromNodeIdNil(b bool)`

 SetFromNodeIdNil sets the value for FromNodeId to be an explicit nil

### UnsetFromNodeId
`func (o *InvestmentGraphNodeEdgesRow) UnsetFromNodeId()`

UnsetFromNodeId ensures that no value is present for FromNodeId, not even an explicit nil
### GetToNodeId

`func (o *InvestmentGraphNodeEdgesRow) GetToNodeId() string`

GetToNodeId returns the ToNodeId field if non-nil, zero value otherwise.

### GetToNodeIdOk

`func (o *InvestmentGraphNodeEdgesRow) GetToNodeIdOk() (*string, bool)`

GetToNodeIdOk returns a tuple with the ToNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToNodeId

`func (o *InvestmentGraphNodeEdgesRow) SetToNodeId(v string)`

SetToNodeId sets ToNodeId field to given value.

### HasToNodeId

`func (o *InvestmentGraphNodeEdgesRow) HasToNodeId() bool`

HasToNodeId returns a boolean if a field has been set.

### SetToNodeIdNil

`func (o *InvestmentGraphNodeEdgesRow) SetToNodeIdNil(b bool)`

 SetToNodeIdNil sets the value for ToNodeId to be an explicit nil

### UnsetToNodeId
`func (o *InvestmentGraphNodeEdgesRow) UnsetToNodeId()`

UnsetToNodeId ensures that no value is present for ToNodeId, not even an explicit nil
### GetEdgeType

`func (o *InvestmentGraphNodeEdgesRow) GetEdgeType() string`

GetEdgeType returns the EdgeType field if non-nil, zero value otherwise.

### GetEdgeTypeOk

`func (o *InvestmentGraphNodeEdgesRow) GetEdgeTypeOk() (*string, bool)`

GetEdgeTypeOk returns a tuple with the EdgeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeType

`func (o *InvestmentGraphNodeEdgesRow) SetEdgeType(v string)`

SetEdgeType sets EdgeType field to given value.

### HasEdgeType

`func (o *InvestmentGraphNodeEdgesRow) HasEdgeType() bool`

HasEdgeType returns a boolean if a field has been set.

### SetEdgeTypeNil

`func (o *InvestmentGraphNodeEdgesRow) SetEdgeTypeNil(b bool)`

 SetEdgeTypeNil sets the value for EdgeType to be an explicit nil

### UnsetEdgeType
`func (o *InvestmentGraphNodeEdgesRow) UnsetEdgeType()`

UnsetEdgeType ensures that no value is present for EdgeType, not even an explicit nil
### GetRelationshipType

`func (o *InvestmentGraphNodeEdgesRow) GetRelationshipType() string`

GetRelationshipType returns the RelationshipType field if non-nil, zero value otherwise.

### GetRelationshipTypeOk

`func (o *InvestmentGraphNodeEdgesRow) GetRelationshipTypeOk() (*string, bool)`

GetRelationshipTypeOk returns a tuple with the RelationshipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipType

`func (o *InvestmentGraphNodeEdgesRow) SetRelationshipType(v string)`

SetRelationshipType sets RelationshipType field to given value.

### HasRelationshipType

`func (o *InvestmentGraphNodeEdgesRow) HasRelationshipType() bool`

HasRelationshipType returns a boolean if a field has been set.

### SetRelationshipTypeNil

`func (o *InvestmentGraphNodeEdgesRow) SetRelationshipTypeNil(b bool)`

 SetRelationshipTypeNil sets the value for RelationshipType to be an explicit nil

### UnsetRelationshipType
`func (o *InvestmentGraphNodeEdgesRow) UnsetRelationshipType()`

UnsetRelationshipType ensures that no value is present for RelationshipType, not even an explicit nil
### GetSource

`func (o *InvestmentGraphNodeEdgesRow) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InvestmentGraphNodeEdgesRow) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InvestmentGraphNodeEdgesRow) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *InvestmentGraphNodeEdgesRow) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *InvestmentGraphNodeEdgesRow) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *InvestmentGraphNodeEdgesRow) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceRef

`func (o *InvestmentGraphNodeEdgesRow) GetSourceRef() string`

GetSourceRef returns the SourceRef field if non-nil, zero value otherwise.

### GetSourceRefOk

`func (o *InvestmentGraphNodeEdgesRow) GetSourceRefOk() (*string, bool)`

GetSourceRefOk returns a tuple with the SourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRef

`func (o *InvestmentGraphNodeEdgesRow) SetSourceRef(v string)`

SetSourceRef sets SourceRef field to given value.

### HasSourceRef

`func (o *InvestmentGraphNodeEdgesRow) HasSourceRef() bool`

HasSourceRef returns a boolean if a field has been set.

### SetSourceRefNil

`func (o *InvestmentGraphNodeEdgesRow) SetSourceRefNil(b bool)`

 SetSourceRefNil sets the value for SourceRef to be an explicit nil

### UnsetSourceRef
`func (o *InvestmentGraphNodeEdgesRow) UnsetSourceRef()`

UnsetSourceRef ensures that no value is present for SourceRef, not even an explicit nil
### GetSourceUrl

`func (o *InvestmentGraphNodeEdgesRow) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *InvestmentGraphNodeEdgesRow) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *InvestmentGraphNodeEdgesRow) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *InvestmentGraphNodeEdgesRow) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *InvestmentGraphNodeEdgesRow) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *InvestmentGraphNodeEdgesRow) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetEventDate

`func (o *InvestmentGraphNodeEdgesRow) GetEventDate() string`

GetEventDate returns the EventDate field if non-nil, zero value otherwise.

### GetEventDateOk

`func (o *InvestmentGraphNodeEdgesRow) GetEventDateOk() (*string, bool)`

GetEventDateOk returns a tuple with the EventDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDate

`func (o *InvestmentGraphNodeEdgesRow) SetEventDate(v string)`

SetEventDate sets EventDate field to given value.

### HasEventDate

`func (o *InvestmentGraphNodeEdgesRow) HasEventDate() bool`

HasEventDate returns a boolean if a field has been set.

### SetEventDateNil

`func (o *InvestmentGraphNodeEdgesRow) SetEventDateNil(b bool)`

 SetEventDateNil sets the value for EventDate to be an explicit nil

### UnsetEventDate
`func (o *InvestmentGraphNodeEdgesRow) UnsetEventDate()`

UnsetEventDate ensures that no value is present for EventDate, not even an explicit nil
### GetEventDateKind

`func (o *InvestmentGraphNodeEdgesRow) GetEventDateKind() string`

GetEventDateKind returns the EventDateKind field if non-nil, zero value otherwise.

### GetEventDateKindOk

`func (o *InvestmentGraphNodeEdgesRow) GetEventDateKindOk() (*string, bool)`

GetEventDateKindOk returns a tuple with the EventDateKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateKind

`func (o *InvestmentGraphNodeEdgesRow) SetEventDateKind(v string)`

SetEventDateKind sets EventDateKind field to given value.

### HasEventDateKind

`func (o *InvestmentGraphNodeEdgesRow) HasEventDateKind() bool`

HasEventDateKind returns a boolean if a field has been set.

### SetEventDateKindNil

`func (o *InvestmentGraphNodeEdgesRow) SetEventDateKindNil(b bool)`

 SetEventDateKindNil sets the value for EventDateKind to be an explicit nil

### UnsetEventDateKind
`func (o *InvestmentGraphNodeEdgesRow) UnsetEventDateKind()`

UnsetEventDateKind ensures that no value is present for EventDateKind, not even an explicit nil
### GetObservedThrough

`func (o *InvestmentGraphNodeEdgesRow) GetObservedThrough() string`

GetObservedThrough returns the ObservedThrough field if non-nil, zero value otherwise.

### GetObservedThroughOk

`func (o *InvestmentGraphNodeEdgesRow) GetObservedThroughOk() (*string, bool)`

GetObservedThroughOk returns a tuple with the ObservedThrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedThrough

`func (o *InvestmentGraphNodeEdgesRow) SetObservedThrough(v string)`

SetObservedThrough sets ObservedThrough field to given value.

### HasObservedThrough

`func (o *InvestmentGraphNodeEdgesRow) HasObservedThrough() bool`

HasObservedThrough returns a boolean if a field has been set.

### SetObservedThroughNil

`func (o *InvestmentGraphNodeEdgesRow) SetObservedThroughNil(b bool)`

 SetObservedThroughNil sets the value for ObservedThrough to be an explicit nil

### UnsetObservedThrough
`func (o *InvestmentGraphNodeEdgesRow) UnsetObservedThrough()`

UnsetObservedThrough ensures that no value is present for ObservedThrough, not even an explicit nil
### GetEventStage

`func (o *InvestmentGraphNodeEdgesRow) GetEventStage() string`

GetEventStage returns the EventStage field if non-nil, zero value otherwise.

### GetEventStageOk

`func (o *InvestmentGraphNodeEdgesRow) GetEventStageOk() (*string, bool)`

GetEventStageOk returns a tuple with the EventStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventStage

`func (o *InvestmentGraphNodeEdgesRow) SetEventStage(v string)`

SetEventStage sets EventStage field to given value.

### HasEventStage

`func (o *InvestmentGraphNodeEdgesRow) HasEventStage() bool`

HasEventStage returns a boolean if a field has been set.

### SetEventStageNil

`func (o *InvestmentGraphNodeEdgesRow) SetEventStageNil(b bool)`

 SetEventStageNil sets the value for EventStage to be an explicit nil

### UnsetEventStage
`func (o *InvestmentGraphNodeEdgesRow) UnsetEventStage()`

UnsetEventStage ensures that no value is present for EventStage, not even an explicit nil
### GetRoundLabel

`func (o *InvestmentGraphNodeEdgesRow) GetRoundLabel() string`

GetRoundLabel returns the RoundLabel field if non-nil, zero value otherwise.

### GetRoundLabelOk

`func (o *InvestmentGraphNodeEdgesRow) GetRoundLabelOk() (*string, bool)`

GetRoundLabelOk returns a tuple with the RoundLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundLabel

`func (o *InvestmentGraphNodeEdgesRow) SetRoundLabel(v string)`

SetRoundLabel sets RoundLabel field to given value.

### HasRoundLabel

`func (o *InvestmentGraphNodeEdgesRow) HasRoundLabel() bool`

HasRoundLabel returns a boolean if a field has been set.

### SetRoundLabelNil

`func (o *InvestmentGraphNodeEdgesRow) SetRoundLabelNil(b bool)`

 SetRoundLabelNil sets the value for RoundLabel to be an explicit nil

### UnsetRoundLabel
`func (o *InvestmentGraphNodeEdgesRow) UnsetRoundLabel()`

UnsetRoundLabel ensures that no value is present for RoundLabel, not even an explicit nil
### GetAmountValue

`func (o *InvestmentGraphNodeEdgesRow) GetAmountValue() float32`

GetAmountValue returns the AmountValue field if non-nil, zero value otherwise.

### GetAmountValueOk

`func (o *InvestmentGraphNodeEdgesRow) GetAmountValueOk() (*float32, bool)`

GetAmountValueOk returns a tuple with the AmountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountValue

`func (o *InvestmentGraphNodeEdgesRow) SetAmountValue(v float32)`

SetAmountValue sets AmountValue field to given value.

### HasAmountValue

`func (o *InvestmentGraphNodeEdgesRow) HasAmountValue() bool`

HasAmountValue returns a boolean if a field has been set.

### SetAmountValueNil

`func (o *InvestmentGraphNodeEdgesRow) SetAmountValueNil(b bool)`

 SetAmountValueNil sets the value for AmountValue to be an explicit nil

### UnsetAmountValue
`func (o *InvestmentGraphNodeEdgesRow) UnsetAmountValue()`

UnsetAmountValue ensures that no value is present for AmountValue, not even an explicit nil
### GetAmountCurrency

`func (o *InvestmentGraphNodeEdgesRow) GetAmountCurrency() string`

GetAmountCurrency returns the AmountCurrency field if non-nil, zero value otherwise.

### GetAmountCurrencyOk

`func (o *InvestmentGraphNodeEdgesRow) GetAmountCurrencyOk() (*string, bool)`

GetAmountCurrencyOk returns a tuple with the AmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCurrency

`func (o *InvestmentGraphNodeEdgesRow) SetAmountCurrency(v string)`

SetAmountCurrency sets AmountCurrency field to given value.

### HasAmountCurrency

`func (o *InvestmentGraphNodeEdgesRow) HasAmountCurrency() bool`

HasAmountCurrency returns a boolean if a field has been set.

### SetAmountCurrencyNil

`func (o *InvestmentGraphNodeEdgesRow) SetAmountCurrencyNil(b bool)`

 SetAmountCurrencyNil sets the value for AmountCurrency to be an explicit nil

### UnsetAmountCurrency
`func (o *InvestmentGraphNodeEdgesRow) UnsetAmountCurrency()`

UnsetAmountCurrency ensures that no value is present for AmountCurrency, not even an explicit nil
### GetIsLead

`func (o *InvestmentGraphNodeEdgesRow) GetIsLead() bool`

GetIsLead returns the IsLead field if non-nil, zero value otherwise.

### GetIsLeadOk

`func (o *InvestmentGraphNodeEdgesRow) GetIsLeadOk() (*bool, bool)`

GetIsLeadOk returns a tuple with the IsLead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLead

`func (o *InvestmentGraphNodeEdgesRow) SetIsLead(v bool)`

SetIsLead sets IsLead field to given value.

### HasIsLead

`func (o *InvestmentGraphNodeEdgesRow) HasIsLead() bool`

HasIsLead returns a boolean if a field has been set.

### SetIsLeadNil

`func (o *InvestmentGraphNodeEdgesRow) SetIsLeadNil(b bool)`

 SetIsLeadNil sets the value for IsLead to be an explicit nil

### UnsetIsLead
`func (o *InvestmentGraphNodeEdgesRow) UnsetIsLead()`

UnsetIsLead ensures that no value is present for IsLead, not even an explicit nil
### GetSharedCompanyCount

`func (o *InvestmentGraphNodeEdgesRow) GetSharedCompanyCount() int32`

GetSharedCompanyCount returns the SharedCompanyCount field if non-nil, zero value otherwise.

### GetSharedCompanyCountOk

`func (o *InvestmentGraphNodeEdgesRow) GetSharedCompanyCountOk() (*int32, bool)`

GetSharedCompanyCountOk returns a tuple with the SharedCompanyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedCompanyCount

`func (o *InvestmentGraphNodeEdgesRow) SetSharedCompanyCount(v int32)`

SetSharedCompanyCount sets SharedCompanyCount field to given value.

### HasSharedCompanyCount

`func (o *InvestmentGraphNodeEdgesRow) HasSharedCompanyCount() bool`

HasSharedCompanyCount returns a boolean if a field has been set.

### SetSharedCompanyCountNil

`func (o *InvestmentGraphNodeEdgesRow) SetSharedCompanyCountNil(b bool)`

 SetSharedCompanyCountNil sets the value for SharedCompanyCount to be an explicit nil

### UnsetSharedCompanyCount
`func (o *InvestmentGraphNodeEdgesRow) UnsetSharedCompanyCount()`

UnsetSharedCompanyCount ensures that no value is present for SharedCompanyCount, not even an explicit nil
### GetIsDerived

`func (o *InvestmentGraphNodeEdgesRow) GetIsDerived() bool`

GetIsDerived returns the IsDerived field if non-nil, zero value otherwise.

### GetIsDerivedOk

`func (o *InvestmentGraphNodeEdgesRow) GetIsDerivedOk() (*bool, bool)`

GetIsDerivedOk returns a tuple with the IsDerived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDerived

`func (o *InvestmentGraphNodeEdgesRow) SetIsDerived(v bool)`

SetIsDerived sets IsDerived field to given value.

### HasIsDerived

`func (o *InvestmentGraphNodeEdgesRow) HasIsDerived() bool`

HasIsDerived returns a boolean if a field has been set.

### SetIsDerivedNil

`func (o *InvestmentGraphNodeEdgesRow) SetIsDerivedNil(b bool)`

 SetIsDerivedNil sets the value for IsDerived to be an explicit nil

### UnsetIsDerived
`func (o *InvestmentGraphNodeEdgesRow) UnsetIsDerived()`

UnsetIsDerived ensures that no value is present for IsDerived, not even an explicit nil
### GetPairKey

`func (o *InvestmentGraphNodeEdgesRow) GetPairKey() string`

GetPairKey returns the PairKey field if non-nil, zero value otherwise.

### GetPairKeyOk

`func (o *InvestmentGraphNodeEdgesRow) GetPairKeyOk() (*string, bool)`

GetPairKeyOk returns a tuple with the PairKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairKey

`func (o *InvestmentGraphNodeEdgesRow) SetPairKey(v string)`

SetPairKey sets PairKey field to given value.

### HasPairKey

`func (o *InvestmentGraphNodeEdgesRow) HasPairKey() bool`

HasPairKey returns a boolean if a field has been set.

### SetPairKeyNil

`func (o *InvestmentGraphNodeEdgesRow) SetPairKeyNil(b bool)`

 SetPairKeyNil sets the value for PairKey to be an explicit nil

### UnsetPairKey
`func (o *InvestmentGraphNodeEdgesRow) UnsetPairKey()`

UnsetPairKey ensures that no value is present for PairKey, not even an explicit nil
### GetLaneRank

`func (o *InvestmentGraphNodeEdgesRow) GetLaneRank() int32`

GetLaneRank returns the LaneRank field if non-nil, zero value otherwise.

### GetLaneRankOk

`func (o *InvestmentGraphNodeEdgesRow) GetLaneRankOk() (*int32, bool)`

GetLaneRankOk returns a tuple with the LaneRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaneRank

`func (o *InvestmentGraphNodeEdgesRow) SetLaneRank(v int32)`

SetLaneRank sets LaneRank field to given value.

### HasLaneRank

`func (o *InvestmentGraphNodeEdgesRow) HasLaneRank() bool`

HasLaneRank returns a boolean if a field has been set.

### SetLaneRankNil

`func (o *InvestmentGraphNodeEdgesRow) SetLaneRankNil(b bool)`

 SetLaneRankNil sets the value for LaneRank to be an explicit nil

### UnsetLaneRank
`func (o *InvestmentGraphNodeEdgesRow) UnsetLaneRank()`

UnsetLaneRank ensures that no value is present for LaneRank, not even an explicit nil
### GetIsPreferred

`func (o *InvestmentGraphNodeEdgesRow) GetIsPreferred() bool`

GetIsPreferred returns the IsPreferred field if non-nil, zero value otherwise.

### GetIsPreferredOk

`func (o *InvestmentGraphNodeEdgesRow) GetIsPreferredOk() (*bool, bool)`

GetIsPreferredOk returns a tuple with the IsPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreferred

`func (o *InvestmentGraphNodeEdgesRow) SetIsPreferred(v bool)`

SetIsPreferred sets IsPreferred field to given value.

### HasIsPreferred

`func (o *InvestmentGraphNodeEdgesRow) HasIsPreferred() bool`

HasIsPreferred returns a boolean if a field has been set.

### SetIsPreferredNil

`func (o *InvestmentGraphNodeEdgesRow) SetIsPreferredNil(b bool)`

 SetIsPreferredNil sets the value for IsPreferred to be an explicit nil

### UnsetIsPreferred
`func (o *InvestmentGraphNodeEdgesRow) UnsetIsPreferred()`

UnsetIsPreferred ensures that no value is present for IsPreferred, not even an explicit nil
### GetSourceLicense

`func (o *InvestmentGraphNodeEdgesRow) GetSourceLicense() string`

GetSourceLicense returns the SourceLicense field if non-nil, zero value otherwise.

### GetSourceLicenseOk

`func (o *InvestmentGraphNodeEdgesRow) GetSourceLicenseOk() (*string, bool)`

GetSourceLicenseOk returns a tuple with the SourceLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLicense

`func (o *InvestmentGraphNodeEdgesRow) SetSourceLicense(v string)`

SetSourceLicense sets SourceLicense field to given value.

### HasSourceLicense

`func (o *InvestmentGraphNodeEdgesRow) HasSourceLicense() bool`

HasSourceLicense returns a boolean if a field has been set.

### SetSourceLicenseNil

`func (o *InvestmentGraphNodeEdgesRow) SetSourceLicenseNil(b bool)`

 SetSourceLicenseNil sets the value for SourceLicense to be an explicit nil

### UnsetSourceLicense
`func (o *InvestmentGraphNodeEdgesRow) UnsetSourceLicense()`

UnsetSourceLicense ensures that no value is present for SourceLicense, not even an explicit nil
### GetAttribution

`func (o *InvestmentGraphNodeEdgesRow) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *InvestmentGraphNodeEdgesRow) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *InvestmentGraphNodeEdgesRow) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *InvestmentGraphNodeEdgesRow) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *InvestmentGraphNodeEdgesRow) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *InvestmentGraphNodeEdgesRow) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


