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

type BearerToken struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	PublicKey string `json:"public_key,omitempty"`
	Token string `json:"token,omitempty"`
	Keyword string `json:"keyword,omitempty"`
	DateExpired time.Time `json:"date_expired,omitempty"`
	User *UserProfile `json:"user,omitempty"`
}
