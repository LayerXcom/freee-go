# \ReceiptsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReceipt**](ReceiptsApi.md#CreateReceipt) | **Post** /api/1/receipts | ファイルボックス 証憑ファイルアップロード
[**DestroyReceipt**](ReceiptsApi.md#DestroyReceipt) | **Delete** /api/1/receipts/{id} | ファイルボックス 証憑ファイルを削除する
[**GetReceipt**](ReceiptsApi.md#GetReceipt) | **Get** /api/1/receipts/{id} | ファイルボックス 証憑ファイルの取得
[**GetReceipts**](ReceiptsApi.md#GetReceipts) | **Get** /api/1/receipts | ファイルボックス 証憑ファイル一覧の取得
[**UpdateReceipt**](ReceiptsApi.md#UpdateReceipt) | **Put** /api/1/receipts/{id} | ファイルボックス 証憑ファイル情報更新



## CreateReceipt

> ReceiptResponse CreateReceipt(ctx, companyId, receipt, optional)

ファイルボックス 証憑ファイルアップロード

 <h2 id=\"\">概要</h2>  <p>ファイルボックスに証憑ファイルをアップロードする</p> <h2 id=\"_2\">注意点</h2> <ul>   <li>リクエストヘッダーの Content-Type は、multipart/form-dataにのみ対応しています。</li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
**receipt** | ***os.File*****os.File**| 証憑ファイル | 
 **optional** | ***CreateReceiptOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateReceiptOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **description** | **optional.String**| メモ (255文字以内) | 
 **issueDate** | **optional.String**| 取引日 (yyyy-mm-dd) | 

### Return type

[**ReceiptResponse**](receiptResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyReceipt

> DestroyReceipt(ctx, id, companyId)

ファイルボックス 証憑ファイルを削除する

 <h2 id=\"\">概要</h2>  <p>ファイルボックスの証憑ファイルを削除する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 証憑ID | 
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


## GetReceipt

> ReceiptResponse GetReceipt(ctx, id, companyId)

ファイルボックス 証憑ファイルの取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所のファイルボックス 証憑ファイルを取得する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 証憑ID | 
**companyId** | **int32**| 事業所ID | 

### Return type

[**ReceiptResponse**](receiptResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceipts

> InlineResponse20016 GetReceipts(ctx, companyId, startDate, endDate, optional)

ファイルボックス 証憑ファイル一覧の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所のファイルボックス 証憑ファイル一覧を取得する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
**startDate** | **string**| アップロード日 (yyyy-mm-dd) | 
**endDate** | **string**| アップロード日 (yyyy-mm-dd) | 
 **optional** | ***GetReceiptsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetReceiptsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **userName** | **optional.String**| アップロードしたユーザー名、メールアドレス | 
 **number** | **optional.Int32**| アップロードファイルNo | 
 **commentType** | **optional.String**| posted:コメントあり, raised:未解決, resolved:解決済 | 
 **commentImportant** | **optional.Bool**| trueの時、重要コメント付きが対象 | 
 **category** | **optional.String**| all:すべて、without_deal:未登録、with_expense_application_line:経費申請中, with_deal:登録済み、ignored:無視 | 
 **offset** | **optional.Int32**| 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **optional.Int32**| 取得レコードの件数 (デフォルト: 50, 最小: 1, 最大: 3000) | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReceipt

> ReceiptResponse UpdateReceipt(ctx, id, receiptUpdateParams)

ファイルボックス 証憑ファイル情報更新

 <h2 id=\"\">概要</h2>  <p>ファイルボックスの証憑ファイル情報を更新する</p> <h2 id=\"_2\">注意点</h2> <ul>   <li>本APIでは、証憑ファイルの再アップロードはできません。</li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 証憑ID | 
**receiptUpdateParams** | [**ReceiptUpdateParams**](ReceiptUpdateParams.md)| 経費申請の更新 | 

### Return type

[**ReceiptResponse**](receiptResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

