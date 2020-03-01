/*
 * Jumpserver API Docs
 *
 * Jumpserver Restful api docs
 *
 * API version: v1
 * Contact: support@fit2cloud.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/oauth2"
)

var (
	jsonCheck = regexp.MustCompile("(?i:[application|text]/json)")
	xmlCheck  = regexp.MustCompile("(?i:[application|text]/xml)")
)

// APIClient manages communication with the Jumpserver API Docs API vv1
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	ApplicationsDatabaseAppsApi *ApplicationsDatabaseAppsApiService

	ApplicationsRemoteAppsApi *ApplicationsRemoteAppsApiService

	AssetsAdminUsersApi *AssetsAdminUsersApiService

	AssetsAssetUsersApi *AssetsAssetUsersApiService

	AssetsAssetUsersInfoApi *AssetsAssetUsersInfoApiService

	AssetsAssetsApi *AssetsAssetsApiService

	AssetsCmdFiltersApi *AssetsCmdFiltersApiService

	AssetsDomainsApi *AssetsDomainsApiService

	AssetsFavoriteAssetsApi *AssetsFavoriteAssetsApiService

	AssetsGatewaysApi *AssetsGatewaysApiService

	AssetsGatheredUsersApi *AssetsGatheredUsersApiService

	AssetsLabelsApi *AssetsLabelsApiService

	AssetsNodesApi *AssetsNodesApiService

	AssetsPlatformsApi *AssetsPlatformsApiService

	AssetsSystemUsersApi *AssetsSystemUsersApiService

	AssetsSystemUsersAssetsRelationsApi *AssetsSystemUsersAssetsRelationsApiService

	AssetsSystemUsersNodesRelationsApi *AssetsSystemUsersNodesRelationsApiService

	AuditsFtpLogsApi *AuditsFtpLogsApiService

	AuthenticationAccessKeysApi *AuthenticationAccessKeysApiService

	AuthenticationAuthApi *AuthenticationAuthApiService

	AuthenticationConnectionTokenApi *AuthenticationConnectionTokenApiService

	AuthenticationLoginConfirmSettingsApi *AuthenticationLoginConfirmSettingsApiService

	AuthenticationLoginConfirmTicketApi *AuthenticationLoginConfirmTicketApiService

	AuthenticationMfaApi *AuthenticationMfaApiService

	AuthenticationOtpApi *AuthenticationOtpApiService

	AuthenticationTokensApi *AuthenticationTokensApiService

	CommonResourcesApi *CommonResourcesApiService

	OpsAdhocApi *OpsAdhocApiService

	OpsCeleryApi *OpsCeleryApiService

	OpsCommandExecutionsApi *OpsCommandExecutionsApiService

	OpsHistoryApi *OpsHistoryApiService

	OpsTasksApi *OpsTasksApiService

	OrgsOrgsApi *OrgsOrgsApiService

	PermsAssetPermissionsApi *PermsAssetPermissionsApiService

	PermsAssetPermissionsAssetsRelationsApi *PermsAssetPermissionsAssetsRelationsApiService

	PermsAssetPermissionsNodesRelationsApi *PermsAssetPermissionsNodesRelationsApiService

	PermsAssetPermissionsSystemUsersRelationsApi *PermsAssetPermissionsSystemUsersRelationsApiService

	PermsAssetPermissionsUserGroupsRelationsApi *PermsAssetPermissionsUserGroupsRelationsApiService

	PermsAssetPermissionsUsersRelationsApi *PermsAssetPermissionsUsersRelationsApiService

	PermsDatabaseAppPermissionsApi *PermsDatabaseAppPermissionsApiService

	PermsDatabaseAppPermissionsDatabaseAppsRelationsApi *PermsDatabaseAppPermissionsDatabaseAppsRelationsApiService

	PermsDatabaseAppPermissionsSystemUsersRelationsApi *PermsDatabaseAppPermissionsSystemUsersRelationsApiService

	PermsDatabaseAppPermissionsUserGroupsRelationsApi *PermsDatabaseAppPermissionsUserGroupsRelationsApiService

	PermsDatabaseAppPermissionsUsersRelationsApi *PermsDatabaseAppPermissionsUsersRelationsApiService

	PermsRemoteAppPermissionsApi *PermsRemoteAppPermissionsApiService

	PermsUserGroupsApi *PermsUserGroupsApiService

	PermsUsersApi *PermsUsersApiService

	SettingsLdapApi *SettingsLdapApiService

	SettingsMailApi *SettingsMailApiService

	SettingsPublicApi *SettingsPublicApiService

	TerminalCommandStoragesApi *TerminalCommandStoragesApiService

	TerminalCommandsApi *TerminalCommandsApiService

	TerminalReplayStoragesApi *TerminalReplayStoragesApiService

	TerminalSessionsApi *TerminalSessionsApiService

	TerminalStatusApi *TerminalStatusApiService

	TerminalTasksApi *TerminalTasksApiService

	TerminalTerminalsApi *TerminalTerminalsApiService

	TicketsTicketsApi *TicketsTicketsApiService

	UsersConnectionTokenApi *UsersConnectionTokenApiService

	UsersGroupsApi *UsersGroupsApiService

	UsersOtpApi *UsersOtpApiService

	UsersProfileApi *UsersProfileApiService

	UsersUsersApi *UsersUsersApiService

	UsersUsersGroupsRelationsApi *UsersUsersGroupsRelationsApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.ApplicationsDatabaseAppsApi = (*ApplicationsDatabaseAppsApiService)(&c.common)
	c.ApplicationsRemoteAppsApi = (*ApplicationsRemoteAppsApiService)(&c.common)
	c.AssetsAdminUsersApi = (*AssetsAdminUsersApiService)(&c.common)
	c.AssetsAssetUsersApi = (*AssetsAssetUsersApiService)(&c.common)
	c.AssetsAssetUsersInfoApi = (*AssetsAssetUsersInfoApiService)(&c.common)
	c.AssetsAssetsApi = (*AssetsAssetsApiService)(&c.common)
	c.AssetsCmdFiltersApi = (*AssetsCmdFiltersApiService)(&c.common)
	c.AssetsDomainsApi = (*AssetsDomainsApiService)(&c.common)
	c.AssetsFavoriteAssetsApi = (*AssetsFavoriteAssetsApiService)(&c.common)
	c.AssetsGatewaysApi = (*AssetsGatewaysApiService)(&c.common)
	c.AssetsGatheredUsersApi = (*AssetsGatheredUsersApiService)(&c.common)
	c.AssetsLabelsApi = (*AssetsLabelsApiService)(&c.common)
	c.AssetsNodesApi = (*AssetsNodesApiService)(&c.common)
	c.AssetsPlatformsApi = (*AssetsPlatformsApiService)(&c.common)
	c.AssetsSystemUsersApi = (*AssetsSystemUsersApiService)(&c.common)
	c.AssetsSystemUsersAssetsRelationsApi = (*AssetsSystemUsersAssetsRelationsApiService)(&c.common)
	c.AssetsSystemUsersNodesRelationsApi = (*AssetsSystemUsersNodesRelationsApiService)(&c.common)
	c.AuditsFtpLogsApi = (*AuditsFtpLogsApiService)(&c.common)
	c.AuthenticationAccessKeysApi = (*AuthenticationAccessKeysApiService)(&c.common)
	c.AuthenticationAuthApi = (*AuthenticationAuthApiService)(&c.common)
	c.AuthenticationConnectionTokenApi = (*AuthenticationConnectionTokenApiService)(&c.common)
	c.AuthenticationLoginConfirmSettingsApi = (*AuthenticationLoginConfirmSettingsApiService)(&c.common)
	c.AuthenticationLoginConfirmTicketApi = (*AuthenticationLoginConfirmTicketApiService)(&c.common)
	c.AuthenticationMfaApi = (*AuthenticationMfaApiService)(&c.common)
	c.AuthenticationOtpApi = (*AuthenticationOtpApiService)(&c.common)
	c.AuthenticationTokensApi = (*AuthenticationTokensApiService)(&c.common)
	c.CommonResourcesApi = (*CommonResourcesApiService)(&c.common)
	c.OpsAdhocApi = (*OpsAdhocApiService)(&c.common)
	c.OpsCeleryApi = (*OpsCeleryApiService)(&c.common)
	c.OpsCommandExecutionsApi = (*OpsCommandExecutionsApiService)(&c.common)
	c.OpsHistoryApi = (*OpsHistoryApiService)(&c.common)
	c.OpsTasksApi = (*OpsTasksApiService)(&c.common)
	c.OrgsOrgsApi = (*OrgsOrgsApiService)(&c.common)
	c.PermsAssetPermissionsApi = (*PermsAssetPermissionsApiService)(&c.common)
	c.PermsAssetPermissionsAssetsRelationsApi = (*PermsAssetPermissionsAssetsRelationsApiService)(&c.common)
	c.PermsAssetPermissionsNodesRelationsApi = (*PermsAssetPermissionsNodesRelationsApiService)(&c.common)
	c.PermsAssetPermissionsSystemUsersRelationsApi = (*PermsAssetPermissionsSystemUsersRelationsApiService)(&c.common)
	c.PermsAssetPermissionsUserGroupsRelationsApi = (*PermsAssetPermissionsUserGroupsRelationsApiService)(&c.common)
	c.PermsAssetPermissionsUsersRelationsApi = (*PermsAssetPermissionsUsersRelationsApiService)(&c.common)
	c.PermsDatabaseAppPermissionsApi = (*PermsDatabaseAppPermissionsApiService)(&c.common)
	c.PermsDatabaseAppPermissionsDatabaseAppsRelationsApi = (*PermsDatabaseAppPermissionsDatabaseAppsRelationsApiService)(&c.common)
	c.PermsDatabaseAppPermissionsSystemUsersRelationsApi = (*PermsDatabaseAppPermissionsSystemUsersRelationsApiService)(&c.common)
	c.PermsDatabaseAppPermissionsUserGroupsRelationsApi = (*PermsDatabaseAppPermissionsUserGroupsRelationsApiService)(&c.common)
	c.PermsDatabaseAppPermissionsUsersRelationsApi = (*PermsDatabaseAppPermissionsUsersRelationsApiService)(&c.common)
	c.PermsRemoteAppPermissionsApi = (*PermsRemoteAppPermissionsApiService)(&c.common)
	c.PermsUserGroupsApi = (*PermsUserGroupsApiService)(&c.common)
	c.PermsUsersApi = (*PermsUsersApiService)(&c.common)
	c.SettingsLdapApi = (*SettingsLdapApiService)(&c.common)
	c.SettingsMailApi = (*SettingsMailApiService)(&c.common)
	c.SettingsPublicApi = (*SettingsPublicApiService)(&c.common)
	c.TerminalCommandStoragesApi = (*TerminalCommandStoragesApiService)(&c.common)
	c.TerminalCommandsApi = (*TerminalCommandsApiService)(&c.common)
	c.TerminalReplayStoragesApi = (*TerminalReplayStoragesApiService)(&c.common)
	c.TerminalSessionsApi = (*TerminalSessionsApiService)(&c.common)
	c.TerminalStatusApi = (*TerminalStatusApiService)(&c.common)
	c.TerminalTasksApi = (*TerminalTasksApiService)(&c.common)
	c.TerminalTerminalsApi = (*TerminalTerminalsApiService)(&c.common)
	c.TicketsTicketsApi = (*TicketsTicketsApiService)(&c.common)
	c.UsersConnectionTokenApi = (*UsersConnectionTokenApiService)(&c.common)
	c.UsersGroupsApi = (*UsersGroupsApiService)(&c.common)
	c.UsersOtpApi = (*UsersOtpApiService)(&c.common)
	c.UsersProfileApi = (*UsersProfileApiService)(&c.common)
	c.UsersUsersApi = (*UsersUsersApiService)(&c.common)
	c.UsersUsersGroupsRelationsApi = (*UsersUsersGroupsRelationsApiService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insenstive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("Expected %s to be of type %s but received %s.", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	}

	return fmt.Sprintf("%v", obj)
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	return c.cfg.HTTPClient.Do(request)
}

// Change base path to allow switching to mocks
func (c *APIClient) ChangeBasePath(path string) {
	c.cfg.BasePath = path
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	fileName string,
	fileBytes []byte) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if len(formParams) > 0 || (len(fileBytes) > 0 && fileName != "") {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		if len(fileBytes) > 0 && fileName != "" {
			w.Boundary()
			//_, fileNm := filepath.Split(fileName)
			part, err := w.CreateFormFile("file", filepath.Base(fileName))
			if err != nil {
				return nil, err
			}
			_, err = part.Write(fileBytes)
			if err != nil {
				return nil, err
			}
			// Set the Boundary in the Content-Type
			headerParams["Content-Type"] = w.FormDataContentType()
		}

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		localVarRequest.Host = c.cfg.Host
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

		// OAuth2 authentication
		if tok, ok := ctx.Value(ContextOAuth2).(oauth2.TokenSource); ok {
			// We were able to grab an oauth2 token from the context
			var latestToken *oauth2.Token
			if latestToken, err = tok.Token(); err != nil {
				return nil, err
			}

			latestToken.SetAuthHeader(localVarRequest)
		}

		// Basic HTTP Authentication
		if auth, ok := ctx.Value(ContextBasicAuth).(BasicAuth); ok {
			localVarRequest.SetBasicAuth(auth.UserName, auth.Password)
		}

		// AccessToken Authentication
		if auth, ok := ctx.Value(ContextAccessToken).(string); ok {
			localVarRequest.Header.Add("Authorization", "Bearer "+auth)
		}
	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}

	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
		if strings.Contains(contentType, "application/xml") {
			if err = xml.Unmarshal(b, v); err != nil {
				return err
			}
			return nil
		} else if strings.Contains(contentType, "application/json") {
			if err = json.Unmarshal(b, v); err != nil {
				return err
			}
			return nil
		}
	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("Invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		}
		expires = now.Add(lifetime)
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// GenericSwaggerError Provides access to the body, error and model on returned errors.
type GenericSwaggerError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericSwaggerError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericSwaggerError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericSwaggerError) Model() interface{} {
	return e.model
}