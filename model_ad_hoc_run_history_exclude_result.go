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

type AdHocRunHistoryExcludeResult struct {
	Id           string  `json:"id,omitempty"`
	Stat         string  `json:"stat,omitempty"`
	TaskDisplay  string  `json:"task_display,omitempty"`
	HostsAmount  int32   `json:"hosts_amount,omitempty"`
	DateStart    string  `json:"date_start,omitempty"`
	DateFinished string  `json:"date_finished,omitempty"`
	Timedelta    float32 `json:"timedelta,omitempty"`
	IsFinished   bool    `json:"is_finished,omitempty"`
	IsSuccess    bool    `json:"is_success,omitempty"`
	Task         string  `json:"task,omitempty"`
	Adhoc        string  `json:"adhoc,omitempty"`
	ShortId      string  `json:"short_id,omitempty"`
	AdhocShortId string  `json:"adhoc_short_id,omitempty"`
}
