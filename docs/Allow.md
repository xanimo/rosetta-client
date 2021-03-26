# Allow

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationStatuses** | [**[]OperationStatus**](OperationStatus.md) | All Operation.Status this implementation supports. Any status that is returned during parsing that is not listed here will cause client validation to error. | [default to null]
**OperationTypes** | **[]string** | All Operation.Type this implementation supports. Any type that is returned during parsing that is not listed here will cause client validation to error. | [default to null]
**Errors** | [**[]ModelError**](Error.md) | All Errors that this implementation could return. Any error that is returned during parsing that is not listed here will cause client validation to error. | [default to null]
**HistoricalBalanceLookup** | **bool** | Any Rosetta implementation that supports querying the balance of an account at any height in the past should set this to true. | [default to null]
**TimestampStartIndex** | **int64** | If populated, &#x60;timestamp_start_index&#x60; indicates the first block index where block timestamps are considered valid (i.e. all blocks less than &#x60;timestamp_start_index&#x60; could have invalid timestamps). This is useful when the genesis block (or blocks) of a network have timestamp 0. If not populated, block timestamps are assumed to be valid for all available blocks. | [optional] [default to null]
**CallMethods** | **[]string** | All methods that are supported by the /call endpoint. Communicating which parameters should be provided to /call is the responsibility of the implementer (this is en lieu of defining an entire type system and requiring the implementer to define that in Allow). | [default to null]
**BalanceExemptions** | [**[]BalanceExemption**](BalanceExemption.md) | BalanceExemptions is an array of BalanceExemption indicating which account balances could change without a corresponding Operation. BalanceExemptions should be used sparingly as they may introduce significant complexity for integrators that attempt to reconcile all account balance changes. If your implementation relies on any BalanceExemptions, you MUST implement historical balance lookup (the ability to query an account balance at any BlockIdentifier). | [default to null]
**MempoolCoins** | **bool** | Any Rosetta implementation that can update an AccountIdentifier&#x27;s unspent coins based on the contents of the mempool should populate this field as true. If false, requests to &#x60;/account/coins&#x60; that set &#x60;include_mempool&#x60; as true will be automatically rejected. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

