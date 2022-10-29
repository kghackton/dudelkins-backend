package dao

import (
	"github.com/uptrace/bun"
	"time"
)

package bo

import "time"

// indexes of columns in comment

type Application struct {
	bun.BaseModel

	RootId                     int32     `bun:"root_id,pk"`
	VersionId                  int32    	`bun:"version_id"`
	Number                     string    `bun:"number"`
	CreatedAt                  time.Time `bun:"created_at"`
	VersionStartedAt           time.Time `bun:"version_started_at"`
	IsIncident                 bool      `bun:"is_incidend"`
	ParentRootId               *int32    `bun:"parent_root_id"`
	UserLastEdited             string   `bun:"user_last_edited"`
	UserLastEditedOrganization string    `bun:"user_last_edited_organization"`
	Comment                    string    `bun:"comment"`
	CategoryId                 int       `bun:"category_id"`
	DefectId                   int       `bun:"defect_id"`
	IsDefectReturnable         bool      `bun:"is_defect_returnable"`
	ApplicantDescription       string    `bun:"applicant_description"`
	ApplicantQuestion          string   `bun:"applicant_question"`
	EmergencyType              string    `bun:"emergency_type"`
	Region                     string    `bun:"region"`
	District                   string    `bun:"district"`
	Address                    string    `bun:"address"`
	UNOM                       int       `bun:"unom"`

	Latitude float64 `bun:"latitude"`
	Longitude float64 `bun:"longitude"`


	Entrance               *int   `bun:"entrance"`
	Floor                  *int   `bun:"floor"`
	Flat                   *int   `bun:"flat"`
	OdsNumber              string `bun:"ods_number"`
	ManagementCompanyTitle string `bun:"management_company_title"`
	ExecutionCompanyTitle  string `bun:"execution_company_title"`

	RenderedServicesIds         []int      `bun:"rendered_services_ids"`
	ConsumedMaterials           string     `bun:"consumed_materials"`
	RenderedSecurityServicesIds []int      `bun:"rendered_security_services_ids"`
	ResultCode                  string     // 53 reject|resolved|consulted . 453 строка Консалтед но плохо?
	AmountOfReturnings          *int       // 54
	LastReturnedAt              *time.Time // 55
	// 56 Признак нахождения на доработке узнать что это?
	IsAlarmed      bool       // 57 да|нет Непонятно что это ??
	ClosedAt       time.Time  // 58
	PreferableFrom *time.Time // 59
	PreferableTo   *time.Time // 60
	RatedAt        time.Time  // 61
	Review         string     // 62 Работы выполнены | Работы невыполнены
	RatingCode     string     // 63 // bad|neutral|good
	// 64-67 Payment Fields
}

