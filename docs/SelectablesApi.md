# \SelectablesApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFormsSelectables**](SelectablesApi.md#GetFormsSelectables) | **Get** /api/1/forms/selectables | フォーム用選択項目情報の取得



## GetFormsSelectables

> SelectablesIndexResponse GetFormsSelectables(ctx, companyId, optional)

フォーム用選択項目情報の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所のフォーム用選択項目情報を取得する</p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
 **optional** | ***GetFormsSelectablesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetFormsSelectablesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includes** | **optional.String**| 取得する項目(項目: account_item) | 

### Return type

[**SelectablesIndexResponse**](selectablesIndexResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

