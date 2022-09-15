package okta

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Env_Var_Proxy(t *testing.T) {
	fmt.Println("HTTPS_PROXY", os.Getenv("HTTPS_PROXY"))
	fmt.Println("https_proxy", os.Getenv("https_proxy"))
	fmt.Println("HTTP_PROXY", os.Getenv("HTTP_PROXY"))
	fmt.Println("http_proxy", os.Getenv("http_proxy"))
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is proxy server end point"))
	})
	proxyServer := httptest.NewServer(mux)
	defer proxyServer.Close()
	configuration := NewConfiguration()
	configuration.Debug = true
	fmt.Println("28", os.Getenv("HTTP_PROXY"))
	fmt.Println("29", proxyServer.URL)
	os.Setenv("HTTP_PROXY", proxyServer.URL)
	fmt.Println("31", os.Getenv("HTTP_PROXY"))
	proxyClient := NewAPIClient(configuration)
	fmt.Println(proxyClient.cfg.HTTPClient.Transport)
	req, err := http.NewRequest(http.MethodGet, "http://example.com", nil)
	require.NoError(t, err, "Create new http request should not error")
	resp, err := proxyClient.callAPI(req)
	require.NoError(t, err, "Make http request should not error")
	os.Unsetenv("HTTP_PROXY")
	b, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "Read http response should not error")
	assert.Equal(t, "This is proxy server end point", string(b))
}

// func Test_List_User(t *testing.T) {
// 	_, resp, err := apiClient.UserApi.ListUsers(apiClient.cfg.Context).Limit(200).Execute()
// 	require.NoError(t, err, "should not error when list users")
// 	b, err := io.ReadAll(resp.Body)
// 	require.NoError(t, err, "Read http response should not error")
// 	t.Error(string(b))
// }

// func Test_List_Idp(t *testing.T) {
// 	_, resp, err := apiClient.IdentityProviderApi.ListIdentityProviders(apiClient.cfg.Context).Limit(200).Execute()
// 	require.NoError(t, err, "should not error when list idp")
// 	b, err := io.ReadAll(resp.Body)
// 	require.NoError(t, err, "Read http response should not error")
// 	t.Error(string(b))
// }

// func Test_List_Group(t *testing.T) {
// 	_, resp, err := apiClient.GroupApi.ListGroups(apiClient.cfg.Context).Limit(200).Execute()
// 	require.NoError(t, err, "should not error when list group")
// 	b, err := io.ReadAll(resp.Body)
// 	require.NoError(t, err, "Read http response should not error")
// 	t.Error(string(b))
// }

// func Test_Config_Proxy(t *testing.T) {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("This is proxy server end point"))
// 	})
// 	proxyServer := httptest.NewServer(mux)
// 	defer proxyServer.Close()
// 	configuration := NewConfiguration()
// 	configuration.Debug = true
// 	proxyURL, err := url.Parse(proxyServer.URL)
// 	require.NoError(t, err, "Parse url should not error")
// 	password, _ := proxyURL.User.Password()
// 	hp := strings.Split(proxyURL.Host, ":")
// 	require.Equal(t, 2, len(hp), "Host should only host and port")
// 	intVar, err := strconv.Atoi(hp[1])
// 	require.NoError(t, err, "Convert string to int should not error")
// 	configuration.Okta.Client.Proxy.Host = hp[0]
// 	configuration.Okta.Client.Proxy.Port = int32(intVar)
// 	configuration.Okta.Client.Proxy.Username = proxyURL.User.Username()
// 	configuration.Okta.Client.Proxy.Password = password
// 	proxyClient := NewAPIClient(configuration)
// 	req, err := http.NewRequest(http.MethodGet, "http://example.com", nil)
// 	require.NoError(t, err, "Create new http request should not error")
// 	resp, err := proxyClient.callAPI(req)
// 	require.NoError(t, err, "Make http request should not error")
// 	b, err := io.ReadAll(resp.Body)
// 	require.NoError(t, err, "Read http response should not error")
// 	assert.Equal(t, "This is proxy server end point", string(b))
// }

// func Test_Proxy_Order(t *testing.T) {
// 	envVarMux := http.NewServeMux()
// 	envVarMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("This is env proxy server end point"))
// 	})
// 	envVarServer := httptest.NewServer(envVarMux)
// 	defer envVarServer.Close()
// 	os.Setenv("HTTP_PROXY", envVarServer.URL)

// 	configMux := http.NewServeMux()
// 	configMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("This is config proxy server end point"))
// 	})
// 	configServer := httptest.NewServer(configMux)
// 	defer configServer.Close()
// 	configuration := NewConfiguration()
// 	configuration.Debug = true
// 	proxyURL, err := url.Parse(configServer.URL)
// 	require.NoError(t, err, "Parse url should not error")
// 	password, _ := proxyURL.User.Password()
// 	hp := strings.Split(proxyURL.Host, ":")
// 	require.Equal(t, 2, len(hp), "Host should only host and port")
// 	intVar, err := strconv.Atoi(hp[1])
// 	require.NoError(t, err, "Convert string to int should not error")
// 	configuration.Okta.Client.Proxy.Host = hp[0]
// 	configuration.Okta.Client.Proxy.Port = int32(intVar)
// 	configuration.Okta.Client.Proxy.Username = proxyURL.User.Username()
// 	configuration.Okta.Client.Proxy.Password = password

// 	proxyClient := NewAPIClient(configuration)
// 	req, err := http.NewRequest(http.MethodGet, "http://example.com", nil)
// 	require.NoError(t, err, "Create new http request should not error")
// 	resp, err := proxyClient.callAPI(req)
// 	require.NoError(t, err, "Make http request should not error")
// 	os.Unsetenv("HTTP_PROXY")
// 	b, err := io.ReadAll(resp.Body)
// 	require.NoError(t, err, "Read http response should not error")
// 	assert.Equal(t, "This is config proxy server end point", string(b))
// }
