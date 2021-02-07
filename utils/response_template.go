package utils

/*
ResponseTemplate 响应模板
*/
type ResponseTemplate struct {
	code    string
	message string
}

var templateSet = []ResponseTemplate{
	{"200", "success"},
	{"-1", "fail"},
}

/*
GenerateResponse 生成无数据列表的响应模板
*/
func GenerateResponse(code, detail string, result string) map[string]string {
	var response map[string]string
	for _, p := range templateSet {

		if code == p.code {
			response = make(map[string]string)

			response["code"] = code
			response["message"] = p.message
			response["detail"] = detail
			response["data"] = result
			break
		}

	}

	return response
}

/*
GenResponse 生成有数据列表的响应模板
*/
func GenResponse(code, detail string, data []map[string]interface{}) map[string]interface{} {
	var response map[string]interface{}
	for _, p := range templateSet {

		if code == p.code {
			response = make(map[string]interface{})

			response["code"] = code
			response["message"] = p.message
			response["detail"] = detail
			response["data"] = data
			break
		}

	}

	return response
}
