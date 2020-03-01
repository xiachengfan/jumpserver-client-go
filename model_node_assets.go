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

type NodeAssets struct {
	Assets []string `json:"assets"`
	OrgId string `json:"org_id,omitempty"`
	OrgName string `json:"org_name,omitempty"`
}
