# InvestmentGraphNodesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**node_id** | Option<**String**> | The node key: an investor id (adv:, agency:, lei:, investor:), a company's organization id (the resolved parent), or investor_site: followed by the normalized printed name for an unadmitted website-only company. | [optional]
**node_kind** | Option<**String**> | investor or company. | [optional]
**display_name** | Option<**String**> | The investor's legal name, the company's resolved parent name, or for an unadmitted website-only company the lowest name as printed on the investors' pages. | [optional]
**investor_type** | Option<**String**> | venture_capital, private_equity, hedge_fund, government_agency or other. Null for a company node. | [optional]
**investor_id** | Option<**String**> | The investor key for an investor node (the same as node_id), null for a company node. | [optional]
**org_id** | Option<**String**> | Legal-entity level: a company node's organization (the same as node_id), or an investor node's own organization resolved to its parent. Null for an unadmitted website-only company and for an investor with no organization. | [optional]
**company_id** | Option<**String**> | Company level: the company record of org_id (null where unresolved). | [optional]
**company_name** | Option<**String**> | Display name of the company. | [optional]
**group_id** | Option<**String**> | Group level: the company's current ultimate parent; a root's group is itself. | [optional]
**group_name** | Option<**String**> | Display name of the group. | [optional]
**same_entity_node_id** | Option<**String**> | The other node of one real-world entity: an investor that is also a company with edges points at its company node and the reverse. An annotation, not an edge. Null when there is none. | [optional]
**node_resolution** | Option<**String**> | spine: the id is an organization or investor id. unadmitted_listing: a website-only company the organization spine has not admitted, keyed on its printed name (same-printed-name companies merge into one node; the id will change on admission). | [optional]
**edge_count** | Option<**i32**> | Stored edge rows incident to the node, as the sum of out_edge_count and in_edge_count. A co_invested edge is stored in both directions, so a co-investor counts it once in each. | [optional]
**out_edge_count** | Option<**i32**> | Stored edge rows that start at the node. | [optional]
**in_edge_count** | Option<**i32**> | Stored edge rows that end at the node. | [optional]
**node_source_count** | Option<**i32**> | Distinct sources among the node's edges. | [optional]
**sources** | Option<**Vec<String>**> | The sources that state the node's edges. | [optional]
**unlinked_company_count** | Option<**i32**> | For an investor node, its portfolio rows that cannot become a node (a company known only by an SEC filer number, an organization with no resolved parent, or a listing that is page chrome, a section heading, the investor's own name or a name a curator marked as not a company); shown as a count and never guessed at. 0 for a company node. | [optional]
**listed_names** | Option<**Vec<String>**> | For a company node, the distinct names investors printed for it on their portfolio pages, sorted; empty for an investor node. A name that differs from display_name is how the investors spell a company the registry names otherwise (Scientist.com, listed for Assay Depot Inc). A label, not a statement of identity. | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


