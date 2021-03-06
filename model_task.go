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

type Task struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name"`
	// 单位: 秒
	Interval int32 `json:"interval,omitempty"`
	// 5 * * * *
	Crontab       string                        `json:"crontab,omitempty"`
	IsPeriodic    bool                          `json:"is_periodic,omitempty"`
	IsDeleted     bool                          `json:"is_deleted,omitempty"`
	Comment       string                        `json:"comment,omitempty"`
	CreatedBy     string                        `json:"created_by,omitempty"`
	DateCreated   string                        `json:"date_created,omitempty"`
	DateUpdated   string                        `json:"date_updated,omitempty"`
	LatestHistory *AdHocRunHistoryExcludeResult `json:"latest_history,omitempty"`
}
