package services_uitls_api

func IsStatusCode2xx(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}
