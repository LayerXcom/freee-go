# \WalletTxnsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWalletTxn**](WalletTxnsApi.md#CreateWalletTxn) | **Post** /api/1/wallet_txns | 明細の作成
[**DestroyWalletTxn**](WalletTxnsApi.md#DestroyWalletTxn) | **Delete** /api/1/wallet_txns/{id} | 明細の削除
[**GetWalletTxn**](WalletTxnsApi.md#GetWalletTxn) | **Get** /api/1/wallet_txns/{id} | 明細の取得
[**GetWalletTxns**](WalletTxnsApi.md#GetWalletTxns) | **Get** /api/1/wallet_txns | 明細一覧の取得



## CreateWalletTxn

> WalletTxnResponse CreateWalletTxn(ctx, optional)

明細の作成

 <h2 id=\"\">概要</h2>  <p>指定した事業所の明細を作成する</p>  <h2 id=\"_2\">定義</h2>  <ul> <li> <p>amount : 明細金額</p> </li>  <li> <p>due_amount : 取引登録待ち金額</p> </li>  <li> <p>balance : 残高</p> </li>  <li> <p>entry_side</p>  <ul> <li>income : 入金</li>  <li>expense : 出金</li> </ul> </li>  <li> <p>walletable_type</p>  <ul> <li>bank_account : 銀行口座</li>  <li>credit_card : クレジットカード</li>  <li>wallet : その他の決済口座</li> </ul> </li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateWalletTxnOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateWalletTxnOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletTxnParams** | [**optional.Interface of WalletTxnParams**](WalletTxnParams.md)| 明細の作成 | 

### Return type

[**WalletTxnResponse**](walletTxnResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyWalletTxn

> DestroyWalletTxn(ctx, id, companyId)

明細の削除

 <h2 id=\"\">概要</h2>  <p>指定した事業所の明細を削除する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 明細ID | 
**companyId** | **int32**| 事業所ID | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletTxn

> WalletTxnResponse GetWalletTxn(ctx, id, companyId)

明細の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の明細を取得する</p>  <h2 id=\"_2\">定義</h2>  <ul> <li> <p>amount : 明細金額</p> </li>  <li> <p>due_amount : 取引登録待ち金額</p> </li>  <li> <p>balance : 残高</p> </li>  <li> <p>entry_side</p>  <ul> <li>income : 入金</li>  <li>expense : 出金</li> </ul> </li>  <li> <p>walletable_type</p>  <ul> <li>bank_account : 銀行口座</li>  <li>credit_card : クレジットカード</li>  <li>wallet : その他の決済口座</li> </ul> </li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 明細ID | 
**companyId** | **int32**| 事業所ID | 

### Return type

[**WalletTxnResponse**](walletTxnResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletTxns

> InlineResponse20015 GetWalletTxns(ctx, companyId, optional)

明細一覧の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の明細一覧を取得する</p>  <h2 id=\"_2\">定義</h2>  <ul> <li> <p>amount : 明細金額</p> </li>  <li> <p>due_amount : 取引登録待ち金額</p> </li>  <li> <p>balance : 残高</p> </li>  <li> <p>entry_side</p>  <ul> <li>income : 入金</li>  <li>expense : 出金</li> </ul> </li>  <li> <p>walletable_type</p>  <ul> <li>bank_account : 銀行口座</li>  <li>credit_card : クレジットカード</li>  <li>wallet : その他の決済口座</li> </ul> </li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
 **optional** | ***GetWalletTxnsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetWalletTxnsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **walletableType** | **optional.String**| 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet) walletable_type、walletable_idは同時に指定が必要です。 | 
 **walletableId** | **optional.Int32**| 口座ID walletable_type、walletable_idは同時に指定が必要です。 | 
 **startDate** | **optional.String**| 取引日で絞込：開始日 (yyyy-mm-dd) | 
 **endDate** | **optional.String**| 取引日で絞込：終了日 (yyyy-mm-dd) | 
 **entrySide** | **optional.String**| 入金／出金 (入金: income, 出金: expense) | 
 **offset** | **optional.Int32**| 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **optional.Int32**| 取得レコードの件数 (デフォルト: 20, 最小: 1, 最大: 100)  | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

