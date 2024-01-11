package utils

func GenerateUrl(baseUrl string, params map[string]string) string {

	url := baseUrl + "?"
	for k, v := range params {
		url += k + "=" + v + "&"
	}
	return url
}
