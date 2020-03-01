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

type AssetUser struct {
	Id           string        `json:"id,omitempty"`
	Hostname     string        `json:"hostname,omitempty"`
	Ip           string        `json:"ip,omitempty"`
	Username     string        `json:"username"`
	Password     string        `json:"password,omitempty"`
	Asset        string        `json:"asset"`
	Version      int32         `json:"version,omitempty"`
	IsLatest     bool          `json:"is_latest,omitempty"`
	Connectivity *Connectivity `json:"connectivity,omitempty"`
	Backend      string        `json:"backend,omitempty"`
	DateCreated  string        `json:"date_created,omitempty"`
	DateUpdated  string        `json:"date_updated,omitempty"`
	PrivateKey   string        `json:"private_key,omitempty"`
	PublicKey    string        `json:"public_key,omitempty"`
	OrgId        string        `json:"org_id,omitempty"`
	OrgName      string        `json:"org_name,omitempty"`
}
