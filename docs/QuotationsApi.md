# \QuotationsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQuotation**](QuotationsApi.md#CreateQuotation) | **Post** /api/1/quotations | 見積書の作成
[**DestroyQuotation**](QuotationsApi.md#DestroyQuotation) | **Delete** /api/1/quotations/{id} | 見積書の削除
[**GetQuotation**](QuotationsApi.md#GetQuotation) | **Get** /api/1/quotations/{id} | 見積書の取得
[**GetQuotations**](QuotationsApi.md#GetQuotations) | **Get** /api/1/quotations | 見積書一覧の取得
[**UpdateQuotation**](QuotationsApi.md#UpdateQuotation) | **Put** /api/1/quotations/{id} | 見積書の更新



## CreateQuotation

> QuotationResponse CreateQuotation(ctx, optional)

見積書の作成

 <h2 id=\"\">概要</h2>  <p>指定した事業所の見積書を作成する</p>  <h2 id=\"_1\">注意点</h2> <ul> <li> <p>partner_code, partner_idはどちらかの指定が必須です。ただし両方同時に指定することはできません。</p> </li> <li> <p>partner_codeを利用するには、事業所の設定から取引先コードの利用を有効にする必要があります。</p> </li> <li> <p>本APIでは見積内容(quotation_contents)は、最大100行までになります。</p> </li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateQuotationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateQuotationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quotationCreateParams** | [**optional.Interface of QuotationCreateParams**](QuotationCreateParams.md)| 見積書の作成 | 

### Return type

[**QuotationResponse**](quotationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyQuotation

> DestroyQuotation(ctx, id, companyId)

見積書の削除

 <h2 id=\"\">概要</h2>  <p>指定した事業所の見積書を削除する</p>

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


## GetQuotation

> QuotationResponse GetQuotation(ctx, companyId, id)

見積書の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の見積書詳細を取得する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
**id** | **int32**| 見積書ID | 

### Return type

[**QuotationResponse**](quotationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuotations

> InlineResponse2003 GetQuotations(ctx, companyId, optional)

見積書一覧の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の見積書一覧を取得する</p> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
 **optional** | ***GetQuotationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetQuotationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **partnerId** | **optional.Int32**| 取引先IDで絞込 | 
 **partnerCode** | **optional.String**| 取引先コードで絞込 | 
 **startIssueDate** | **optional.String**| 見積日の開始日(yyyy-mm-dd) | 
 **endIssueDate** | **optional.String**| 見積日の終了日(yyyy-mm-dd) | 
 **quotationNumber** | **optional.String**| 見積書番号 | 
 **description** | **optional.String**| 概要 | 
 **quotationStatus** | **optional.String**| 見積書ステータス  (unsubmitted: 送付待ち, submitted: 送付済み, all: 全て) | 
 **offset** | **optional.Int32**| 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **optional.Int32**| 取得レコードの件数 (デフォルト: 20, 最大: 100)  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateQuotation

> QuotationResponse UpdateQuotation(ctx, id, optional)

見積書の更新

 <h2 id=\"\">概要</h2>  <p>指定した事業所の見積書を更新する</p>  <h2 id=\"_1\">注意点</h2> <ul> <li> <p>partner_code, partner_idを両方同時に指定することはできません。</p> </li> <li> <p>partner_codeを利用するには、事業所の設定から取引先コードの利用を有効にする必要があります。</p> </li> <li> <p>本APIでは見積内容(quotation_contents)は、最大100行までになります。</p> </li> </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 見積書ID | 
 **optional** | ***UpdateQuotationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateQuotationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **quotationUpdateParams** | [**optional.Interface of QuotationUpdateParams**](QuotationUpdateParams.md)| 見積書の更新 | 

### Return type

[**QuotationResponse**](quotationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

