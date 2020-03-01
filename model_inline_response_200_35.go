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

type InlineResponse20035 struct {
	Count int32 `json:"count"`
	Next string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
	Results []AssetPermissionCreateUpdate `json:"results"`
}
