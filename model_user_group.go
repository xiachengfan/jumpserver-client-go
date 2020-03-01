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

type UserGroup struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name"`
	Users []string `json:"users,omitempty"`
	UsersAmount string `json:"users_amount,omitempty"`
	Comment string `json:"comment,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	OrgName string `json:"org_name,omitempty"`
}