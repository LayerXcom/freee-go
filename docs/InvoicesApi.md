# \InvoicesApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInvoice**](InvoicesApi.md#CreateInvoice) | **Post** /api/1/invoices | 請求書の作成
[**DestroyInvoice**](InvoicesApi.md#DestroyInvoice) | **Delete** /api/1/invoices/{id} | 請求書の削除
[**GetInvoice**](InvoicesApi.md#GetInvoice) | **Get** /api/1/invoices/{id} | 請求書の取得
[**GetInvoices**](InvoicesApi.md#GetInvoices) | **Get** /api/1/invoices | 請求書一覧の取得
[**UpdateInvoice**](InvoicesApi.md#UpdateInvoice) | **Put** /api/1/invoices/{id} | 請求書の更新



## CreateInvoice

> InvoiceResponse CreateInvoice(ctx, optional)

請求書の作成

 <h2 id=\"\">概要</h2>  <p>指定した事業所の請求書を作成する</p>  <h2 id=\"_1\">注意点</h2> <ul> <li> <p>partner_code, partner_idはどちらかの指定が必須です。ただし両方同時に指定することはできません。</p> </li> <li> <p>請求書ステータス(invoice_status)を発行(issue)で利用した場合、請求内容の合計金額が0円以上になる必要があります。</p> </li> <li> <p>partner_codeを利用するには、事業所の設定から取引先コードの利用を有効にする必要があります。</p> </li> <li> <p>本APIでは請求内容(invoice_contents)は、最大100行までになります。</p> </li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateInvoiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceCreateParams** | [**optional.Interface of InvoiceCreateParams**](InvoiceCreateParams.md)| 請求書の作成 | 

### Return type

[**InvoiceResponse**](invoiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyInvoice

> DestroyInvoice(ctx, id, companyId)

請求書の削除

 <h2 id=\"\">概要</h2>  <p>指定した事業所の請求書を削除する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
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


## GetInvoice

> InvoiceResponse GetInvoice(ctx, companyId, id)

請求書の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の請求書詳細を取得する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
**id** | **int32**| 請求書ID | 

### Return type

[**InvoiceResponse**](invoiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoices

> InlineResponse2002 GetInvoices(ctx, companyId, optional)

請求書一覧の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の請求書一覧を取得する</p> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
 **optional** | ***GetInvoicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInvoicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **partnerId** | **optional.Int32**| 取引先IDで絞込 | 
 **partnerCode** | **optional.String**| 取引先コードで絞込 | 
 **startIssueDate** | **optional.String**| 請求日の開始日(yyyy-mm-dd) | 
 **endIssueDate** | **optional.String**| 請求日の終了日(yyyy-mm-dd) | 
 **startDueDate** | **optional.String**| 期日の開始日(yyyy-mm-dd) | 
 **endDueDate** | **optional.String**| 期日の終了日(yyyy-mm-dd) | 
 **invoiceNumber** | **optional.String**| 請求書番号 | 
 **description** | **optional.String**| 概要 | 
 **invoiceStatus** | **optional.String**| 請求書ステータス  (draft: 下書き, applying: 申請中, remanded: 差し戻し, rejected: 却下, approved: 承認済み, unsubmitted: 送付待ち, submitted: 送付済み) | 
 **paymentStatus** | **optional.String**| 入金ステータス  (unsettled: 入金待ち, settled: 入金済み) | 
 **offset** | **optional.Int32**| 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **optional.Int32**| 取得レコードの件数 (デフォルト: 20, 最大: 100)  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInvoice

> InvoiceResponse UpdateInvoice(ctx, id, optional)

請求書の更新

 <h2 id=\"\">概要</h2>  <p>指定した事業所の請求書を更新する</p>  <h2 id=\"_1\">注意点</h2> <ul> <li> <p>入金済みの請求書に対する金額関連の変更はできません。</p> </li> <li> <p>請求書ステータスは確定(issue)のみ指定可能です。請求書ステータスを確定する時のみ指定してください。</p> </li> <li> <p>請求書WFを利用している場合、承認済み請求書は承認権限を持たないユーザーでは更新できません。</p> </li> <li> <p>更新後の請求書ステータス(invoice_status)が下書き以外の場合、請求内容の合計金額が0円以上になる必要があります。</p> </li> <li> <p>partner_code, partner_idを両方同時に指定することはできません。</p> </li> <li> <p>partner_codeを利用するには、事業所の設定から取引先コードの利用を有効にする必要があります。</p> </li> <li> <p>本APIでは請求内容(invoice_contents)は、最大100行までになります。</p> </li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 請求書ID | 
 **optional** | ***UpdateInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateInvoiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoiceUpdateParams** | [**optional.Interface of InvoiceUpdateParams**](InvoiceUpdateParams.md)| 請求書の更新 | 

### Return type

[**InvoiceResponse**](invoiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

