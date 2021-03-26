# Go API client for swagger

Build Once. Integrate Your Blockchain Everywhere.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.4.10
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to */*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountApi* | [**AccountBalance**](docs/AccountApi.md#accountbalance) | **Post** /account/balance | Get an Account&#x27;s Balance
*AccountApi* | [**AccountCoins**](docs/AccountApi.md#accountcoins) | **Post** /account/coins | Get an Account&#x27;s Unspent Coins
*BlockApi* | [**Block**](docs/BlockApi.md#block) | **Post** /block | Get a Block
*BlockApi* | [**BlockTransaction**](docs/BlockApi.md#blocktransaction) | **Post** /block/transaction | Get a Block Transaction
*CallApi* | [**Call**](docs/CallApi.md#call) | **Post** /call | Make a Network-Specific Procedure Call
*ConstructionApi* | [**ConstructionCombine**](docs/ConstructionApi.md#constructioncombine) | **Post** /construction/combine | Create Network Transaction from Signatures
*ConstructionApi* | [**ConstructionDerive**](docs/ConstructionApi.md#constructionderive) | **Post** /construction/derive | Derive an AccountIdentifier from a PublicKey
*ConstructionApi* | [**ConstructionHash**](docs/ConstructionApi.md#constructionhash) | **Post** /construction/hash | Get the Hash of a Signed Transaction
*ConstructionApi* | [**ConstructionMetadata**](docs/ConstructionApi.md#constructionmetadata) | **Post** /construction/metadata | Get Metadata for Transaction Construction
*ConstructionApi* | [**ConstructionParse**](docs/ConstructionApi.md#constructionparse) | **Post** /construction/parse | Parse a Transaction
*ConstructionApi* | [**ConstructionPayloads**](docs/ConstructionApi.md#constructionpayloads) | **Post** /construction/payloads | Generate an Unsigned Transaction and Signing Payloads
*ConstructionApi* | [**ConstructionPreprocess**](docs/ConstructionApi.md#constructionpreprocess) | **Post** /construction/preprocess | Create a Request to Fetch Metadata
*ConstructionApi* | [**ConstructionSubmit**](docs/ConstructionApi.md#constructionsubmit) | **Post** /construction/submit | Submit a Signed Transaction
*EventsApi* | [**EventsBlocks**](docs/EventsApi.md#eventsblocks) | **Post** /events/blocks | [INDEXER] Get a range of BlockEvents
*MempoolApi* | [**Mempool**](docs/MempoolApi.md#mempool) | **Post** /mempool | Get All Mempool Transactions
*MempoolApi* | [**MempoolTransaction**](docs/MempoolApi.md#mempooltransaction) | **Post** /mempool/transaction | Get a Mempool Transaction
*NetworkApi* | [**NetworkList**](docs/NetworkApi.md#networklist) | **Post** /network/list | Get List of Available Networks
*NetworkApi* | [**NetworkOptions**](docs/NetworkApi.md#networkoptions) | **Post** /network/options | Get Network Options
*NetworkApi* | [**NetworkStatus**](docs/NetworkApi.md#networkstatus) | **Post** /network/status | Get Network Status
*SearchApi* | [**SearchTransactions**](docs/SearchApi.md#searchtransactions) | **Post** /search/transactions | [INDEXER] Search for Transactions

## Documentation For Models

 - [AccountBalanceRequest](docs/AccountBalanceRequest.md)
 - [AccountBalanceResponse](docs/AccountBalanceResponse.md)
 - [AccountCoinsRequest](docs/AccountCoinsRequest.md)
 - [AccountCoinsResponse](docs/AccountCoinsResponse.md)
 - [AccountIdentifier](docs/AccountIdentifier.md)
 - [Allow](docs/Allow.md)
 - [Amount](docs/Amount.md)
 - [BalanceExemption](docs/BalanceExemption.md)
 - [Block](docs/Block.md)
 - [BlockEvent](docs/BlockEvent.md)
 - [BlockEventType](docs/BlockEventType.md)
 - [BlockIdentifier](docs/BlockIdentifier.md)
 - [BlockRequest](docs/BlockRequest.md)
 - [BlockResponse](docs/BlockResponse.md)
 - [BlockTransaction](docs/BlockTransaction.md)
 - [BlockTransactionRequest](docs/BlockTransactionRequest.md)
 - [BlockTransactionResponse](docs/BlockTransactionResponse.md)
 - [CallRequest](docs/CallRequest.md)
 - [CallResponse](docs/CallResponse.md)
 - [Coin](docs/Coin.md)
 - [CoinAction](docs/CoinAction.md)
 - [CoinChange](docs/CoinChange.md)
 - [CoinIdentifier](docs/CoinIdentifier.md)
 - [ConstructionCombineRequest](docs/ConstructionCombineRequest.md)
 - [ConstructionCombineResponse](docs/ConstructionCombineResponse.md)
 - [ConstructionDeriveRequest](docs/ConstructionDeriveRequest.md)
 - [ConstructionDeriveResponse](docs/ConstructionDeriveResponse.md)
 - [ConstructionHashRequest](docs/ConstructionHashRequest.md)
 - [ConstructionMetadataRequest](docs/ConstructionMetadataRequest.md)
 - [ConstructionMetadataResponse](docs/ConstructionMetadataResponse.md)
 - [ConstructionParseRequest](docs/ConstructionParseRequest.md)
 - [ConstructionParseResponse](docs/ConstructionParseResponse.md)
 - [ConstructionPayloadsRequest](docs/ConstructionPayloadsRequest.md)
 - [ConstructionPayloadsResponse](docs/ConstructionPayloadsResponse.md)
 - [ConstructionPreprocessRequest](docs/ConstructionPreprocessRequest.md)
 - [ConstructionPreprocessResponse](docs/ConstructionPreprocessResponse.md)
 - [ConstructionSubmitRequest](docs/ConstructionSubmitRequest.md)
 - [Currency](docs/Currency.md)
 - [CurveType](docs/CurveType.md)
 - [Direction](docs/Direction.md)
 - [EventsBlocksRequest](docs/EventsBlocksRequest.md)
 - [EventsBlocksResponse](docs/EventsBlocksResponse.md)
 - [ExemptionType](docs/ExemptionType.md)
 - [MempoolResponse](docs/MempoolResponse.md)
 - [MempoolTransactionRequest](docs/MempoolTransactionRequest.md)
 - [MempoolTransactionResponse](docs/MempoolTransactionResponse.md)
 - [MetadataRequest](docs/MetadataRequest.md)
 - [ModelError](docs/ModelError.md)
 - [NetworkIdentifier](docs/NetworkIdentifier.md)
 - [NetworkListResponse](docs/NetworkListResponse.md)
 - [NetworkOptionsResponse](docs/NetworkOptionsResponse.md)
 - [NetworkRequest](docs/NetworkRequest.md)
 - [NetworkStatusResponse](docs/NetworkStatusResponse.md)
 - [Operation](docs/Operation.md)
 - [OperationIdentifier](docs/OperationIdentifier.md)
 - [OperationStatus](docs/OperationStatus.md)
 - [Operator](docs/Operator.md)
 - [PartialBlockIdentifier](docs/PartialBlockIdentifier.md)
 - [Peer](docs/Peer.md)
 - [PublicKey](docs/PublicKey.md)
 - [RelatedTransaction](docs/RelatedTransaction.md)
 - [SearchTransactionsRequest](docs/SearchTransactionsRequest.md)
 - [SearchTransactionsResponse](docs/SearchTransactionsResponse.md)
 - [Signature](docs/Signature.md)
 - [SignatureType](docs/SignatureType.md)
 - [SigningPayload](docs/SigningPayload.md)
 - [SubAccountIdentifier](docs/SubAccountIdentifier.md)
 - [SubNetworkIdentifier](docs/SubNetworkIdentifier.md)
 - [SyncStatus](docs/SyncStatus.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionIdentifier](docs/TransactionIdentifier.md)
 - [TransactionIdentifierResponse](docs/TransactionIdentifierResponse.md)
 - [Version](docs/Version.md)

## Documentation For Authorization
 Endpoints do not require authorization.


## Author


# rosetta-client-go
