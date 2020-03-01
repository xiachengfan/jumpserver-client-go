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

type AdHoc struct {
	Id string `json:"id,omitempty"`
	Task string `json:"task"`
	Tasks string `json:"tasks"`
	Pattern string `json:"pattern,omitempty"`
	Options string `json:"options,omitempty"`
	Hosts []string `json:"hosts"`
	RunAsAdmin bool `json:"run_as_admin,omitempty"`
	RunAs string `json:"run_as,omitempty"`
	Become string `json:"become,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
	ShortId string `json:"short_id,omitempty"`
	BecomeDisplay string `json:"become_display,omitempty"`
}
