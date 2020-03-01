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
	"time"
)

type OrgRead struct {
	Id string `json:"id,omitempty"`
	Admins []string `json:"admins,omitempty"`
	Auditors []string `json:"auditors,omitempty"`
	Users []string `json:"users,omitempty"`
	UserGroups string `json:"user_groups,omitempty"`
	Assets string `json:"assets,omitempty"`
	Domains string `json:"domains,omitempty"`
	AdminUsers string `json:"admin_users,omitempty"`
	SystemUsers string `json:"system_users,omitempty"`
	Labels string `json:"labels,omitempty"`
	Perms string `json:"perms,omitempty"`
	Name string `json:"name"`
	CreatedBy string `json:"created_by,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
	Comment string `json:"comment,omitempty"`
}
