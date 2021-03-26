# AccountCoinsResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockIdentifier** | [***BlockIdentifier**](BlockIdentifier.md) |  | [default to null]
**Coins** | [**[]Coin**](Coin.md) | If a blockchain is UTXO-based, all unspent Coins owned by an account_identifier should be returned alongside the balance. It is highly recommended to populate this field so that users of the Rosetta API implementation don&#x27;t need to maintain their own indexer to track their UTXOs. | [default to null]
**Metadata** | [***interface{}**](interface{}.md) | Account-based blockchains that utilize a nonce or sequence number should include that number in the metadata. This number could be unique to the identifier or global across the account address. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

