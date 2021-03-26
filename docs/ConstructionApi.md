# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConstructionCombine**](ConstructionApi.md#ConstructionCombine) | **Post** /construction/combine | Create Network Transaction from Signatures
[**ConstructionDerive**](ConstructionApi.md#ConstructionDerive) | **Post** /construction/derive | Derive an AccountIdentifier from a PublicKey
[**ConstructionHash**](ConstructionApi.md#ConstructionHash) | **Post** /construction/hash | Get the Hash of a Signed Transaction
[**ConstructionMetadata**](ConstructionApi.md#ConstructionMetadata) | **Post** /construction/metadata | Get Metadata for Transaction Construction
[**ConstructionParse**](ConstructionApi.md#ConstructionParse) | **Post** /construction/parse | Parse a Transaction
[**ConstructionPayloads**](ConstructionApi.md#ConstructionPayloads) | **Post** /construction/payloads | Generate an Unsigned Transaction and Signing Payloads
[**ConstructionPreprocess**](ConstructionApi.md#ConstructionPreprocess) | **Post** /construction/preprocess | Create a Request to Fetch Metadata
[**ConstructionSubmit**](ConstructionApi.md#ConstructionSubmit) | **Post** /construction/submit | Submit a Signed Transaction

# **ConstructionCombine**
> ConstructionCombineResponse ConstructionCombine(ctx, body)
Create Network Transaction from Signatures

Combine creates a network-specific transaction from an unsigned transaction and an array of provided signatures. The signed transaction returned from this method will be sent to the `/construction/submit` endpoint by the caller.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConstructionCombineRequest**](ConstructionCombineRequest.md)|  | 

### Return type

[**ConstructionCombineResponse**](ConstructionCombineResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConstructionDerive**
> ConstructionDeriveResponse ConstructionDerive(ctx, body)
Derive an AccountIdentifier from a PublicKey

Derive returns the AccountIdentifier associated with a public key. Blockchains that require an on-chain action to create an account should not implement this method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConstructionDeriveRequest**](ConstructionDeriveRequest.md)|  | 

### Return type

[**ConstructionDeriveResponse**](ConstructionDeriveResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConstructionHash**
> TransactionIdentifierResponse ConstructionHash(ctx, body)
Get the Hash of a Signed Transaction

TransactionHash returns the network-specific transaction hash for a signed transaction.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConstructionHashRequest**](ConstructionHashRequest.md)|  | 

### Return type

[**TransactionIdentifierResponse**](TransactionIdentifierResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConstructionMetadata**
> ConstructionMetadataResponse ConstructionMetadata(ctx, body)
Get Metadata for Transaction Construction

Get any information required to construct a transaction for a specific network. Metadata returned here could be a recent hash to use, an account sequence number, or even arbitrary chain state. The request used when calling this endpoint is created by calling `/construction/preprocess` in an offline environment. You should NEVER assume that the request sent to this endpoint will be created by the caller or populated with any custom parameters. This must occur in `/construction/preprocess`. It is important to clarify that this endpoint should not pre-construct any transactions for the client (this should happen in `/construction/payloads`). This endpoint is left purposely unstructured because of the wide scope of metadata that could be required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConstructionMetadataRequest**](ConstructionMetadataRequest.md)|  | 

### Return type

[**ConstructionMetadataResponse**](ConstructionMetadataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConstructionParse**
> ConstructionParseResponse ConstructionParse(ctx, body)
Parse a Transaction

Parse is called on both unsigned and signed transactions to understand the intent of the formulated transaction. This is run as a sanity check before signing (after `/construction/payloads`) and before broadcast (after `/construction/combine`). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConstructionParseRequest**](ConstructionParseRequest.md)|  | 

### Return type

[**ConstructionParseResponse**](ConstructionParseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConstructionPayloads**
> ConstructionPayloadsResponse ConstructionPayloads(ctx, body)
Generate an Unsigned Transaction and Signing Payloads

Payloads is called with an array of operations and the response from `/construction/metadata`. It returns an unsigned transaction blob and a collection of payloads that must be signed by particular AccountIdentifiers using a certain SignatureType. The array of operations provided in transaction construction often times can not specify all \"effects\" of a transaction (consider invoked transactions in Ethereum). However, they can deterministically specify the \"intent\" of the transaction, which is sufficient for construction. For this reason, parsing the corresponding transaction in the Data API (when it lands on chain) will contain a superset of whatever operations were provided during construction.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConstructionPayloadsRequest**](ConstructionPayloadsRequest.md)|  | 

### Return type

[**ConstructionPayloadsResponse**](ConstructionPayloadsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConstructionPreprocess**
> ConstructionPreprocessResponse ConstructionPreprocess(ctx, body)
Create a Request to Fetch Metadata

Preprocess is called prior to `/construction/payloads` to construct a request for any metadata that is needed for transaction construction given (i.e. account nonce). The `options` object returned from this endpoint will be sent to the `/construction/metadata` endpoint UNMODIFIED by the caller (in an offline execution environment). If your Construction API implementation has configuration options, they MUST be specified in the `/construction/preprocess` request (in the `metadata` field).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConstructionPreprocessRequest**](ConstructionPreprocessRequest.md)|  | 

### Return type

[**ConstructionPreprocessResponse**](ConstructionPreprocessResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConstructionSubmit**
> TransactionIdentifierResponse ConstructionSubmit(ctx, body)
Submit a Signed Transaction

Submit a pre-signed transaction to the node. This call should not block on the transaction being included in a block. Rather, it should return immediately with an indication of whether or not the transaction was included in the mempool. The transaction submission response should only return a 200 status if the submitted transaction could be included in the mempool. Otherwise, it should return an error.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConstructionSubmitRequest**](ConstructionSubmitRequest.md)|  | 

### Return type

[**TransactionIdentifierResponse**](TransactionIdentifierResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

