# \AccountItemsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountItem**](AccountItemsApi.md#CreateAccountItem) | **Post** /api/1/account_items | 勘定科目の作成
[**DestroyAccountItem**](AccountItemsApi.md#DestroyAccountItem) | **Delete** /api/1/account_items/{id} | 勘定科目の削除
[**GetAccountItem**](AccountItemsApi.md#GetAccountItem) | **Get** /api/1/account_items/{id} | 勘定科目の詳細情報の取得
[**GetAccountItems**](AccountItemsApi.md#GetAccountItems) | **Get** /api/1/account_items | 勘定科目一覧の取得
[**UpdateAccountItem**](AccountItemsApi.md#UpdateAccountItem) | **Put** /api/1/account_items/{id} | 勘定科目の更新



## CreateAccountItem

> AccountItemResponse CreateAccountItem(ctx, accountItemParams)

勘定科目の作成

 <h2 id=\"\">概要</h2>  <p>指定した事業所の勘定科目を作成する</p>  <h2 id=\"_2\">注意点</h2> <p>tax_nameは、api/1/taxes/companies/{company_id} で該当事業所の税区分の一覧を取得して、availableの値がtrue、かつ”name_ja”に”税率%”を含んでいない税区分を確認して、そのnameを指定して勘定科目の作成をしてください</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountItemParams** | [**AccountItemParams**](AccountItemParams.md)| 勘定科目の作成 | 

### Return type

[**AccountItemResponse**](accountItemResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyAccountItem

> DestroyAccountItem(ctx, id, companyId)

勘定科目の削除

 <h2 id=\"\">概要</h2>  <p>指定した勘定科目を削除する</p> <h2 id=\"\">注意点</h2> <ul> <li>削除できる勘定科目は、追加で作成したカスタム勘定項目のみです。</li> <li>デフォルトで存在する勘定科目や口座の勘定科目は削除できません。</li></ul>

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


## GetAccountItem

> AccountItemResponse GetAccountItem(ctx, companyId, id)

勘定科目の詳細情報の取得

 <h2 id=\"\">概要</h2>  <p>指定した勘定科目の詳細を取得する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
**id** | **int32**| 勘定科目ID | 

### Return type

[**AccountItemResponse**](accountItemResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountItems

> AccountItemsResponse GetAccountItems(ctx, companyId, optional)

勘定科目一覧の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の勘定科目一覧を取得する</p>  <h2 id=\"_2\">定義</h2>  <ul> <li>default_tax_id : デフォルト設定がされている税区分ID</li>  <li>default_tax_code : リクエストした日時を基準とした税区分コード</li> </ul>  <h2 id=\"_3\">注意点</h2> <p>default_tax_code は勘定科目作成・更新時に利用するものではありません</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
 **optional** | ***GetAccountItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAccountItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseDate** | **optional.String**| 基準日:指定した場合、勘定科目に紐づく税区分(default_tax_code)が、基準日の税率に基づいて返ります。 | 

### Return type

[**AccountItemsResponse**](accountItemsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountItem

> AccountItemResponse UpdateAccountItem(ctx, id, accountItemParams)

勘定科目の更新

 <h2 id=\"\">概要</h2>  <p>勘定科目を更新する</p>  <h2 id=\"_2\">注意点</h2> <p>tax_codeは、api/1/taxes/companies/{company_id} で該当事業所の税区分の一覧を取得して、availableの値がtrue、かつ”name_ja”に”税率%”を含んでいない税区分を確認して、そのcodeを指定して勘定科目の更新をしてください</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**accountItemParams** | [**AccountItemParams**](AccountItemParams.md)| 勘定科目の更新 | 

### Return type

[**AccountItemResponse**](accountItemResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

