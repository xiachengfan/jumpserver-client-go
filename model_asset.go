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

type Asset struct {
	Id string `json:"id,omitempty"`
	Ip string `json:"ip"`
	Hostname string `json:"hostname"`
	Protocol string `json:"protocol,omitempty"`
	Port int32 `json:"port,omitempty"`
	Protocols []string `json:"protocols,omitempty"`
	Platform string `json:"platform"`
	IsActive bool `json:"is_active,omitempty"`
	PublicIp string `json:"public_ip,omitempty"`
	Domain string `json:"domain,omitempty"`
	AdminUser string `json:"admin_user,omitempty"`
	Nodes []string `json:"nodes,omitempty"`
	Labels []string `json:"labels,omitempty"`
	Number string `json:"number,omitempty"`
	Vendor string `json:"vendor,omitempty"`
	Model string `json:"model,omitempty"`
	Sn string `json:"sn,omitempty"`
	CpuModel string `json:"cpu_model,omitempty"`
	CpuCount int32 `json:"cpu_count,omitempty"`
	CpuCores int32 `json:"cpu_cores,omitempty"`
	CpuVcpus int32 `json:"cpu_vcpus,omitempty"`
	Memory string `json:"memory,omitempty"`
	DiskTotal string `json:"disk_total,omitempty"`
	DiskInfo string `json:"disk_info,omitempty"`
	Os string `json:"os,omitempty"`
	OsVersion string `json:"os_version,omitempty"`
	OsArch string `json:"os_arch,omitempty"`
	HostnameRaw string `json:"hostname_raw,omitempty"`
	Comment string `json:"comment,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
	HardwareInfo string `json:"hardware_info,omitempty"`
	Connectivity *Connectivity `json:"connectivity,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	OrgName string `json:"org_name,omitempty"`
}
