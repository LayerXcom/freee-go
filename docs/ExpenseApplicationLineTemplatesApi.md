# \ExpenseApplicationLineTemplatesApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExpenseApplicationLineTemplate**](ExpenseApplicationLineTemplatesApi.md#CreateExpenseApplicationLineTemplate) | **Post** /api/1/expense_application_line_templates | 経費科目の作成
[**DestroyExpenseApplicationLineTemplate**](ExpenseApplicationLineTemplatesApi.md#DestroyExpenseApplicationLineTemplate) | **Delete** /api/1/expense_application_line_templates/{id} | 経費科目の削除
[**GetExpenseApplicationLineTemplate**](ExpenseApplicationLineTemplatesApi.md#GetExpenseApplicationLineTemplate) | **Get** /api/1/expense_application_line_templates/{id} | 経費科目の取得
[**GetExpenseApplicationLineTemplates**](ExpenseApplicationLineTemplatesApi.md#GetExpenseApplicationLineTemplates) | **Get** /api/1/expense_application_line_templates | 経費科目一覧の取得
[**UpdateExpenseApplicationLineTemplate**](ExpenseApplicationLineTemplatesApi.md#UpdateExpenseApplicationLineTemplate) | **Put** /api/1/expense_application_line_templates/{id} | 経費科目の更新



## CreateExpenseApplicationLineTemplate

> ExpenseApplicationLineTemplateResponse CreateExpenseApplicationLineTemplate(ctx, expenseApplicationLineTemplateParams)

経費科目の作成

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**expenseApplicationLineTemplateParams** | [**ExpenseApplicationLineTemplateParams**](ExpenseApplicationLineTemplateParams.md)| 経費科目の作成 | 

### Return type

[**ExpenseApplicationLineTemplateResponse**](expenseApplicationLineTemplateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyExpenseApplicationLineTemplate

> DestroyExpenseApplicationLineTemplate(ctx, id, companyId)

経費科目の削除

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 経費科目ID | 
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


## GetExpenseApplicationLineTemplate

> ExpenseApplicationLineTemplateResponse GetExpenseApplicationLineTemplate(ctx, id, companyId)

経費科目の取得

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 経費科目ID | 
**companyId** | **int32**| 事業所ID | 

### Return type

[**ExpenseApplicationLineTemplateResponse**](expenseApplicationLineTemplateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExpenseApplicationLineTemplates

> InlineResponse20017 GetExpenseApplicationLineTemplates(ctx, companyId, optional)

経費科目一覧の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の経費科目一覧を取得する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
 **optional** | ***GetExpenseApplicationLineTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetExpenseApplicationLineTemplatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **optional.Int32**| 取得レコードの件数 (デフォルト: 20, 最小: 1, 最大: 100) | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExpenseApplicationLineTemplate

> ExpenseApplicationLineTemplateResponse UpdateExpenseApplicationLineTemplate(ctx, id, expenseApplicationLineTemplateParams)

経費科目の更新

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 経費科目ID | 
**expenseApplicationLineTemplateParams** | [**ExpenseApplicationLineTemplateParams**](ExpenseApplicationLineTemplateParams.md)| 経費科目の更新 | 

### Return type

[**ExpenseApplicationLineTemplateResponse**](expenseApplicationLineTemplateResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

