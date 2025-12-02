package modulego

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProtectionAPIRequestPayload_Encode_RequiredParamsOnly(t *testing.T) {
	// Arrange: Create payload with only required fields
	payload := ProtectionAPIRequestPayload{
		Key:               "test-key-123",
		IP:                "192.168.1.1",
		Request:           "/api/endpoint",
		RequestModuleName: "Go",
	}

	// Act: Encode the payload
	encoded := payload.Encode()

	// Assert: Verify encoded string contains all required params
	parsedValues, err := url.ParseQuery(encoded)
	assert.NoError(t, err, "Encoded string should be valid URL query")
	assert.Equal(t, "test-key-123", parsedValues.Get("Key"))
	assert.Equal(t, "192.168.1.1", parsedValues.Get("IP"))
	assert.Equal(t, "/api/endpoint", parsedValues.Get("Request"))
	assert.Equal(t, "Go", parsedValues.Get("RequestModuleName"))

	// Assert: Verify optional parameters are NOT present
	assert.Empty(t, parsedValues.Get("ModuleVersion"), "ModuleVersion should not be present")
	assert.Empty(t, parsedValues.Get("UserAgent"), "UserAgent should not be present")
	assert.Empty(t, parsedValues.Get("Host"), "Host should not be present")
}

func TestProtectionAPIRequestPayload_Encode_AdditionalParams(t *testing.T) {
	// Arrange: Create payload with required and additional optional fields
	payload := ProtectionAPIRequestPayload{
		// Required fields
		Key:               "test-key-456",
		IP:                "10.0.0.1",
		Request:           "/graphql",
		RequestModuleName: "Go",

		// Additional optional fields
		ModuleVersion:          "2.2.0",
		ServerName:             "api-server-01",
		APIConnectionState:     "new",
		Port:                   "443",
		TimeRequest:            "1234567890",
		Protocol:               "HTTPS/1.1",
		Method:                 "POST",
		ServerHostName:         "api.example.com",
		HeadersList:            "accept,user-agent,content-type",
		Host:                   "api.example.com",
		UserAgent:              "Mozilla/5.0",
		Referer:                "https://example.com",
		Accept:                 "application/json",
		AcceptEncoding:         "gzip, deflate, br",
		AcceptLanguage:         "en-US,en;q=0.9",
		AcceptCharset:          "utf-8",
		Origin:                 "https://example.com",
		XForwardedForIP:        "203.0.113.1",
		XRequestedWith:         "XMLHttpRequest",
		Connection:             "keep-alive",
		Pragma:                 "no-cache",
		CacheControl:           "no-cache",
		CookiesLen:             "3",
		CookiesList:            "session,user_id,preferences",
		AuthorizationLen:       "20",
		PostParamLen:           "150",
		XRealIP:                "203.0.113.2",
		ClientID:               "client-abc-123",
		SecChDeviceMemory:      "8",
		SecChUA:                `"Chromium";v="120", "Google Chrome";v="120"`,
		SecChUAArch:            "x86",
		SecChUAFullVersionList: `"Chromium";v="120.0.6099.129"`,
		SecChUAMobile:          "?0",
		SecChUAModel:           "Pixel 7",
		SecChUAPlatform:        "Windows",
		SecFetchDest:           "empty",
		SecFetchMode:           "cors",
		SecFetchSite:           "same-origin",
		SecFetchUser:           "?1",
		Via:                    "1.1 proxy.example.com",
		From:                   "webmaster@example.com",
		ContentType:            "application/json",
		TrueClientIP:           "203.0.113.3",

		// GraphQL specific fields
		GraphQLOperationCount: "1",
		GraphQLOperationName:  "getUserProfile",
		GraphQLOperationType:  Query,
	}

	// Act: Encode the payload
	encoded := payload.Encode()

	// Assert: Verify encoded string is valid and contains all params
	parsedValues, err := url.ParseQuery(encoded)
	assert.NoError(t, err, "Encoded string should be valid URL query")

	// Assert: Verify all required parameters are present
	assert.Equal(t, "test-key-456", parsedValues.Get("Key"))
	assert.Equal(t, "10.0.0.1", parsedValues.Get("IP"))
	assert.Equal(t, "/graphql", parsedValues.Get("Request"))
	assert.Equal(t, "Go", parsedValues.Get("RequestModuleName"))

	// Assert: Verify additional optional parameters are present
	assert.Equal(t, "2.2.0", parsedValues.Get("ModuleVersion"))
	assert.Equal(t, "api-server-01", parsedValues.Get("ServerName"))
	assert.Equal(t, "new", parsedValues.Get("APIConnectionState"))
	assert.Equal(t, "443", parsedValues.Get("Port"))
	assert.Equal(t, "1234567890", parsedValues.Get("TimeRequest"))
	assert.Equal(t, "HTTPS/1.1", parsedValues.Get("Protocol"))
	assert.Equal(t, "POST", parsedValues.Get("Method"))
	assert.Equal(t, "api.example.com", parsedValues.Get("ServerHostname"))
	assert.Equal(t, "accept,user-agent,content-type", parsedValues.Get("HeadersList"))
	assert.Equal(t, "api.example.com", parsedValues.Get("Host"))
	assert.Equal(t, "Mozilla/5.0", parsedValues.Get("UserAgent"))
	assert.Equal(t, "https://example.com", parsedValues.Get("Referer"))
	assert.Equal(t, "application/json", parsedValues.Get("Accept"))
	assert.Equal(t, "gzip, deflate, br", parsedValues.Get("AcceptEncoding"))
	assert.Equal(t, "en-US,en;q=0.9", parsedValues.Get("AcceptLanguage"))
	assert.Equal(t, "utf-8", parsedValues.Get("AcceptCharset"))
	assert.Equal(t, "https://example.com", parsedValues.Get("Origin"))
	assert.Equal(t, "203.0.113.1", parsedValues.Get("XForwardedForIP"))
	assert.Equal(t, "XMLHttpRequest", parsedValues.Get("X-Requested-With"))
	assert.Equal(t, "keep-alive", parsedValues.Get("Connection"))
	assert.Equal(t, "no-cache", parsedValues.Get("Pragma"))
	assert.Equal(t, "no-cache", parsedValues.Get("CacheControl"))
	assert.Equal(t, "3", parsedValues.Get("CookiesLen"))
	assert.Equal(t, "session,user_id,preferences", parsedValues.Get("CookiesList"))
	assert.Equal(t, "20", parsedValues.Get("AuthorizationLen"))
	assert.Equal(t, "150", parsedValues.Get("PostParamLen"))
	assert.Equal(t, "203.0.113.2", parsedValues.Get("X-Real-IP"))
	assert.Equal(t, "client-abc-123", parsedValues.Get("ClientID"))

	// Assert: Verify Sec-CH-* headers
	assert.Equal(t, "8", parsedValues.Get("SecCHDeviceMemory"))
	assert.Equal(t, `"Chromium";v="120", "Google Chrome";v="120"`, parsedValues.Get("SecCHUA"))
	assert.Equal(t, "x86", parsedValues.Get("SecCHUAArch"))
	assert.Equal(t, `"Chromium";v="120.0.6099.129"`, parsedValues.Get("SecCHUAFullVersionList"))
	assert.Equal(t, "?0", parsedValues.Get("SecCHUAMobile"))
	assert.Equal(t, "Pixel 7", parsedValues.Get("SecCHUAModel"))
	assert.Equal(t, "Windows", parsedValues.Get("SecCHUAPlatform"))

	// Assert: Verify Sec-Fetch-* headers
	assert.Equal(t, "empty", parsedValues.Get("SecFetchDest"))
	assert.Equal(t, "cors", parsedValues.Get("SecFetchMode"))
	assert.Equal(t, "same-origin", parsedValues.Get("SecFetchSite"))
	assert.Equal(t, "?1", parsedValues.Get("SecFetchUser"))

	// Assert: Verify other optional headers
	assert.Equal(t, "1.1 proxy.example.com", parsedValues.Get("Via"))
	assert.Equal(t, "webmaster@example.com", parsedValues.Get("From"))
	assert.Equal(t, "application/json", parsedValues.Get("ContentType"))
	assert.Equal(t, "203.0.113.3", parsedValues.Get("TrueClientIP"))

	// Assert: Verify GraphQL fields
	assert.Equal(t, "1", parsedValues.Get("GraphQLOperationCount"))
	assert.Equal(t, "getUserProfile", parsedValues.Get("GraphQLOperationName"))
	assert.Equal(t, "query", parsedValues.Get("GraphQLOperationType"))
}

func TestProtectionAPIRequestPayload_Encode_GraphQLOperationName(t *testing.T) {
	t.Run("GraphQLOperationName is empty string", func(t *testing.T) {
		// Arrange: Payload with empty string GraphQLOperationName
		payload := ProtectionAPIRequestPayload{
			Key:                  "test-key",
			IP:                   "192.168.1.1",
			Request:              "/graphql",
			RequestModuleName:    "Go",
			GraphQLOperationName: "",
		}

		// Act
		encoded := payload.Encode()
		parsedValues, err := url.ParseQuery(encoded)
		assert.NoError(t, err)

		// Assert: GraphQLOperationName should not be present when empty
		assert.Empty(t, parsedValues.Get("GraphQLOperationName"), "GraphQLOperationName should not be present when empty string")
	})

	t.Run("GraphQLOperationName has value", func(t *testing.T) {
		// Arrange: Payload with valid GraphQLOperationName
		payload := ProtectionAPIRequestPayload{
			Key:                  "test-key",
			IP:                   "192.168.1.1",
			Request:              "/graphql",
			RequestModuleName:    "Go",
			GraphQLOperationName: "getUser",
		}

		// Act
		encoded := payload.Encode()
		parsedValues, err := url.ParseQuery(encoded)
		assert.NoError(t, err)

		// Assert: GraphQLOperationName should be present
		assert.Equal(t, "getUser", parsedValues.Get("GraphQLOperationName"), "GraphQLOperationName should be present with value")
	})
}

func TestProtectionAPIRequestPayload_Encode_GraphQLOperationType(t *testing.T) {
	testCases := []struct {
		name          string
		operationType OperationType
		expectedValue string
	}{
		{
			name:          "Query operation",
			operationType: Query,
			expectedValue: "query",
		},
		{
			name:          "Mutation operation",
			operationType: Mutation,
			expectedValue: "mutation",
		},
		{
			name:          "Subscription operation",
			operationType: Subscription,
			expectedValue: "subscription",
		},
		{
			name:          "Empty operation type",
			operationType: "",
			expectedValue: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			payload := ProtectionAPIRequestPayload{
				Key:                  "test-key",
				IP:                   "192.168.1.1",
				Request:              "/graphql",
				RequestModuleName:    "Go",
				GraphQLOperationType: tc.operationType,
			}

			// Act
			encoded := payload.Encode()
			parsedValues, err := url.ParseQuery(encoded)
			assert.NoError(t, err)

			// Assert
			if tc.expectedValue == "" {
				assert.Empty(t, parsedValues.Get("GraphQLOperationType"), "GraphQLOperationType should not be present when empty")
			} else {
				assert.Equal(t, tc.expectedValue, parsedValues.Get("GraphQLOperationType"))
			}
		})
	}
}

func TestProtectionAPIRequestPayload_Encode_SpecialCharactersEncoding(t *testing.T) {
	// Arrange: Payload with special characters that need URL encoding
	payload := ProtectionAPIRequestPayload{
		Key:               "test-key",
		IP:                "192.168.1.1",
		Request:           "/api/search?q=hello world&filter=active",
		RequestModuleName: "Go",
		UserAgent:         "Mozilla/5.0 (Windows NT 10.0; Win64; x64)",
		Referer:           "https://example.com/path?param=value&other=test",
		Accept:            "text/html,application/xhtml+xml",
	}

	// Act
	encoded := payload.Encode()

	// Assert: Verify encoding and decoding works correctly
	parsedValues, err := url.ParseQuery(encoded)
	assert.NoError(t, err, "Should be able to parse encoded values")
	assert.Equal(t, "/api/search?q=hello world&filter=active", parsedValues.Get("Request"))
	assert.Equal(t, "Mozilla/5.0 (Windows NT 10.0; Win64; x64)", parsedValues.Get("UserAgent"))
	assert.Equal(t, "https://example.com/path?param=value&other=test", parsedValues.Get("Referer"))
	assert.Equal(t, "text/html,application/xhtml+xml", parsedValues.Get("Accept"))
}

func TestProtectionAPIRequestPayload_Encode_EmptyOptionalFields(t *testing.T) {
	// Arrange: Payload with some empty optional fields
	payload := ProtectionAPIRequestPayload{
		Key:               "test-key",
		IP:                "192.168.1.1",
		Request:           "/api/endpoint",
		RequestModuleName: "Go",

		// These are explicitly empty
		ModuleVersion: "",
		ServerName:    "",
		UserAgent:     "",
		Port:          "",
	}

	// Act
	encoded := payload.Encode()
	parsedValues, err := url.ParseQuery(encoded)
	assert.NoError(t, err)

	// Assert: Empty optional fields should not be present in the output
	assert.Empty(t, parsedValues.Get("ModuleVersion"))
	assert.Empty(t, parsedValues.Get("ServerName"))
	assert.Empty(t, parsedValues.Get("UserAgent"))
	assert.Empty(t, parsedValues.Get("Port"))

	// Assert: Only 4 parameters should be present (the required ones)
	assert.Len(t, parsedValues, 4, "Only required parameters should be present")
}
