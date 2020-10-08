# \SectionsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSection**](SectionsApi.md#CreateSection) | **Post** /api/1/sections | 部門の作成
[**DestroySection**](SectionsApi.md#DestroySection) | **Delete** /api/1/sections/{id} | 部門の削除
[**GetSection**](SectionsApi.md#GetSection) | **Get** /api/1/sections/{id} | 
[**GetSections**](SectionsApi.md#GetSections) | **Get** /api/1/sections | 部門一覧の取得
[**UpdateSection**](SectionsApi.md#UpdateSection) | **Put** /api/1/sections/{id} | 部門の更新



## CreateSection

> SectionResponse CreateSection(ctx, optional)

部門の作成

 <h2 id=\"\">概要</h2>  <p>指定した事業所の部門を作成する</p>  <h2 id=\"_2\">レスポンスの例</h2>  <pre><code>// プレミアムプラン（個人）、ビジネスプラン（法人）、エンタープライズプラン（法人） {   &quot;section&quot; : {     &quot;id&quot; : 102,     &quot;company_id&quot; : 1,     &quot;name&quot; : &quot;開発部門&quot;,     &quot;shortcut1&quot; : &quot;DEVELOPER&quot;,     &quot;shortcut2&quot; : &quot;123&quot;,     &quot;indent_count&quot;: 1,     &quot;parent_id&quot;: 101   } } // それ以外のプラン {   &quot;section&quot; : {     &quot;id&quot; : 102,     &quot;company_id&quot; : 1,     &quot;name&quot; : &quot;開発部門&quot;,     &quot;shortcut1&quot; : &quot;DEVELOPER&quot;,     &quot;shortcut2&quot; : &quot;123&quot;   } }</code></pre> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sectionParams** | [**optional.Interface of SectionParams**](SectionParams.md)| 部門の作成 | 

### Return type

[**SectionResponse**](sectionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroySection

> DestroySection(ctx, id, companyId)

部門の削除

 <h2 id=\"\">概要</h2>  <p>指定した事業所の部門を削除する</p>

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


## GetSection

> SectionResponse GetSection(ctx, id, companyId)



 <h2 id=\"\">概要</h2>  <p>指定した事業所の部門を参照する</p><h2 id=\"_2\">レスポンスの例</h2>  <pre><code>// プレミアムプラン（個人）、ビジネスプラン（法人）、エンタープライズプラン（法人） {   &quot;section&quot; : {     &quot;id&quot; : 102,     &quot;company_id&quot; : 1,     &quot;name&quot; : &quot;開発部門&quot;,     &quot;long_name&quot;: &quot;開発部門&quot;,     &quot;shortcut1&quot; : &quot;DEVELOPER&quot;,     &quot;shortcut2&quot; : &quot;123&quot;,     &quot;indent_count&quot;: 1,     &quot;parent_id&quot;: 101   } } // それ以外のプラン {   &quot;section&quot; : {     &quot;id&quot; : 102,     &quot;company_id&quot; : 1,     &quot;name&quot; : &quot;開発部門&quot;,     &quot;long_name&quot;: &quot;開発部門&quot;,     &quot;shortcut1&quot; : &quot;DEVELOPER&quot;,     &quot;shortcut2&quot; : &quot;123&quot;   } }</code></pre> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| 部門ID | 
**companyId** | **int32**| 事業所ID | 

### Return type

[**SectionResponse**](sectionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSections

> InlineResponse2001 GetSections(ctx, companyId)

部門一覧の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所の部門一覧を取得する</p>  <h2 id=\"_2\">レスポンスの例</h2>  <blockquote> <p>GET https://api.freee.co.jp/api/1/sections?company_id=1</p> </blockquote>  <pre><code>// プレミアムプラン（個人）、ビジネスプラン（法人）、エンタープライズプラン（法人） {   &quot;sections&quot; : [     {       &quot;id&quot; : 101,       &quot;company_id&quot; : 1,       &quot;name&quot; : &quot;開発部門&quot;,       &quot;long_name&quot;: &quot;開発部門&quot;,       &quot;shortcut1&quot; : &quot;DEVELOPER&quot;,       &quot;shortcut2&quot; : &quot;123&quot;,       &quot;indent_count&quot;: 1,       &quot;parent_id&quot;: 11     },     ...   ] } // それ以外のプラン {   &quot;sections&quot; : [     {       &quot;id&quot; : 101,       &quot;company_id&quot; : 1,       &quot;name&quot; : &quot;開発部門&quot;,       &quot;long_name&quot;: &quot;開発部門&quot;,       &quot;shortcut1&quot; : &quot;DEVELOPER&quot;,       &quot;shortcut2&quot; : &quot;123&quot;     },     ...   ] }</code></pre> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSection

> SectionResponse UpdateSection(ctx, id, optional)

部門の更新

 <h2 id=\"\">概要</h2>  <p>指定した事業所の部門を更新する</p><h2 id=\"_2\">レスポンスの例</h2>  <pre><code>// プレミアムプラン（個人）、ビジネスプラン（法人）、エンタープライズプラン（法人） {   &quot;section&quot; : {     &quot;id&quot; : 102,     &quot;company_id&quot; : 1,     &quot;name&quot; : &quot;開発部門&quot;,     &quot;long_name&quot;: &quot;開発部門&quot;,     &quot;shortcut1&quot; : &quot;DEVELOPER&quot;,     &quot;shortcut2&quot; : &quot;123&quot;,     &quot;indent_count&quot;: 1,     &quot;parent_id&quot;: 101   } } // それ以外のプラン {   &quot;section&quot; : {     &quot;id&quot; : 102,     &quot;company_id&quot; : 1,     &quot;name&quot; : &quot;開発部門&quot;,     &quot;long_name&quot;: &quot;開発部門&quot;,     &quot;shortcut1&quot; : &quot;DEVELOPER&quot;,     &quot;shortcut2&quot; : &quot;123&quot;   } }</code></pre> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
 **optional** | ***UpdateSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sectionParams** | [**optional.Interface of SectionParams**](SectionParams.md)| 部門の更新 | 

### Return type

[**SectionResponse**](sectionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

