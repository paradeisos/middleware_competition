package middleware

import "strings"

var regionPrefix = "h#"

// ModifyRequestParams API 请求前处理输入参数
// @params params 必须是指针
// ModifyRequestParams 解析 API 输入参数中的 Region|RegionId|RegionID 字段，检查其是否带有 regionPrefix 前缀，如果有，去除。
//     例如 regionPrefix = h#，请求参数中 region 为 h#cn-north-1，则处理之后，该字段值变为 cn-north-1
func ModifyRequestParams(params interface{}) error {
	// TODO write your code here
	return nil
}

// ModifyResponseData API 请求后处理返回数据
// @params data 必须是指针
// ModifyResponseData 解析 API 输出数据中的 Region|RegionId|RegionID 字段，将原值前加上 regionPrefix。
//     例如 regionPrefix = h# ，返回 region 为 cn-north-1，则处理之后，该字段变为 h#cn-north-1
func ModifyResponseData(data interface{}) error {
	// TODO write your code here
	return nil
}

// 去除前缀
func decodeRegionID(regionID string) string {
	if !strings.HasPrefix(regionID, regionPrefix) {
		return regionID
	}

	return strings.TrimPrefix(regionID, regionPrefix)
}

// 添加前缀
func encodeRegionID(regionID string) string {
	if strings.HasPrefix(regionID, regionPrefix) {
		return regionID
	}

	return regionPrefix + regionID
}
