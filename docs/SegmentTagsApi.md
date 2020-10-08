# \SegmentTagsApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSegmentTag**](SegmentTagsApi.md#CreateSegmentTag) | **Post** /api/1/segments/{segment_id}/tags | セグメントの作成
[**DestroySegmentsTag**](SegmentTagsApi.md#DestroySegmentsTag) | **Delete** /api/1/segments/{segment_id}/tags/{id} | セグメントタグの削除
[**GetSegmentTags**](SegmentTagsApi.md#GetSegmentTags) | **Get** /api/1/segments/{segment_id}/tags | セグメントタグ一覧の取得
[**UpdateSegmentTag**](SegmentTagsApi.md#UpdateSegmentTag) | **Put** /api/1/segments/{segment_id}/tags/{id} | セグメントタグの更新



## CreateSegmentTag

> SegmentTagResponse CreateSegmentTag(ctx, segmentId, segmentTagParams)

セグメントの作成

 <h2 id=\"\">概要</h2>  <p>指定した事業所のセグメントタグを作成する</p>  <h2 id=\"\">注意点</h2>  <ul>  <li>本APIは法人向けのプロフェッショナルプラン以上で利用可能です。利用可能なセグメントの数は、法人向けのプロフェッショナルプランの場合は1つ、エンタープライズプランの場合は3つです。</li>  </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**segmentId** | **int32**| セグメントID（1,2,3のいずれか） 該当プラン以外で参照した場合にはエラーとなります。   1: 法人向けプロフェッショナル, 法人向けエンタープライズプラン   2,3: 法人向け エンタープライズプラン  | 
**segmentTagParams** | [**SegmentTagParams**](SegmentTagParams.md)| セグメントタグの作成 | 

### Return type

[**SegmentTagResponse**](segmentTagResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroySegmentsTag

> DestroySegmentsTag(ctx, segmentId, id, companyId)

セグメントタグの削除

 <h2 id=\"\">概要</h2>  <p>指定した事業所のセグメントタグを削除する</p>  <h2 id=\"\">注意点</h2>  <ul>  <li>本APIは法人向けのプロフェッショナルプラン以上で利用可能です。利用可能なセグメントの数は、法人向けのプロフェッショナルプランの場合は1つ、エンタープライズプランの場合は3つです。</li>  </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**segmentId** | **int32**| セグメントID（1,2,3のいずれか） 該当プラン以外で参照した場合にはエラーとなります。   1: 法人向けプロフェッショナル, 法人向けエンタープライズプラン   2,3: 法人向け エンタープライズプラン  | 
**id** | **int32**| セグメントタグID | 
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


## GetSegmentTags

> InlineResponse20019 GetSegmentTags(ctx, companyId, segmentId, optional)

セグメントタグ一覧の取得

 <h2 id=\"\">概要</h2>  <p>指定した事業所のセグメントタグ一覧を取得する</p>  <h2 id=\"\">注意点</h2>  <ul>  <li>本APIは法人向けのプロフェッショナルプラン以上で利用可能です。利用可能なセグメントの数は、法人向けのプロフェッショナルプランの場合は1つ、エンタープライズプランの場合は3つです。</li>  </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32**| 事業所ID | 
**segmentId** | **int32**| セグメントID（1,2,3のいずれか） 該当プラン以外で参照した場合にはエラーとなります。   1: 法人向けプロフェッショナル, 法人向けエンタープライズプラン   2,3: 法人向け エンタープライズプラン  | 
 **optional** | ***GetSegmentTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSegmentTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **optional.Int32**| 取得レコードのオフセット (デフォルト: 0) | 
 **limit** | **optional.Int32**| 取得レコードの件数 (デフォルト: 20, 最小: 1, 最大: 500)  | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSegmentTag

> SegmentTagResponse UpdateSegmentTag(ctx, segmentId, id, segmentTagParams)

セグメントタグの更新

 <h2 id=\"\">概要</h2>  <p>指定した事業所のセグメントタグを更新する</p>  <h2 id=\"\">注意点</h2>  <ul>  <li>本APIは法人向けのプロフェッショナルプラン以上で利用可能です。利用可能なセグメントの数は、法人向けのプロフェッショナルプランの場合は1つ、エンタープライズプランの場合は3つです。</li>  </ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**segmentId** | **int32**| セグメントID（1,2,3のいずれか） 該当プラン以外で参照した場合にはエラーとなります。   1: 法人向けプロフェッショナル, 法人向けエンタープライズプラン   2,3: 法人向け エンタープライズプラン  | 
**id** | **int32**| セグメントタグID | 
**segmentTagParams** | [**SegmentTagParams**](SegmentTagParams.md)| セグメントタグの作成 | 

### Return type

[**SegmentTagResponse**](segmentTagResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

