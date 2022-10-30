package dao

import (
	"time"

	"github.com/uptrace/bun"

	"dudelkins/internal/objects/bo"
)

type Application struct {
	bun.BaseModel `bun:"table:applications,alias:a"`

	RootId                      int32      `bun:"root_id,pk"`
	VersionId                   int32      `bun:"version_id"`
	Number                      string     `bun:"number"`
	CreatedAt                   time.Time  `bun:"created_at"`
	VersionStartedAt            time.Time  `bun:"version_started_at"`
	IsIncident                  bool       `bun:"is_incident"`
	ParentRootId                *int32     `bun:"parent_root_id"`
	UserLastEdited              string     `bun:"user_last_edited"`
	UserLastEditedOrganization  string     `bun:"user_last_edited_organization"`
	Comment                     string     `bun:"comment"`
	CategoryId                  int        `bun:"category_id"`
	DefectId                    int        `bun:"defect_id"`
	IsDefectReturnable          bool       `bun:"is_defect_returnable"`
	ApplicantDescription        string     `bun:"applicant_description"`
	ApplicantQuestion           string     `bun:"applicant_question"`
	EmergencyType               string     `bun:"emergency_type"`
	Region                      string     `bun:"region"`
	District                    string     `bun:"district"`
	Address                     string     `bun:"address"`
	UNOM                        int        `bun:"unom"`
	Latitude                    float64    `bun:"latitude"`
	Longitude                   float64    `bun:"longitude"`
	Entrance                    *int       `bun:"entrance"`
	Floor                       *int       `bun:"floor"`
	Flat                        *int       `bun:"flat"`
	OdsNumber                   string     `bun:"ods_number"`
	ManagementCompanyTitle      string     `bun:"management_company_title"`
	ExecutionCompanyTitle       string     `bun:"execution_company_title"`
	RenderedServicesIds         []int      `bun:"rendered_services_ids,array"`
	ConsumedMaterials           string     `bun:"consumed_materials"`
	RenderedSecurityServicesIds []int      `bun:"rendered_security_services_ids,array"`
	ResultCode                  string     `bun:"result_code"`
	AmountOfReturnings          *int       `bun:"amount_of_returnings"`
	LastReturnedAt              *time.Time `bun:"last_returned_at"`
	IsAlarmed                   bool       `bun:"is_alarmed"`
	ClosedAt                    time.Time  `bun:"closed_at"`
	PreferableFrom              *time.Time `bun:"preferable_from"`
	PreferableTo                *time.Time `bun:"preferable_to"`
	RatedAt                     *time.Time `bun:"rated_at"`
	Review                      *string    `bun:"review"`
	RatingCode                  *string    `bun:"rating_code"`
}

func NewApplication(application bo.Application) Application {
	return Application{
		RootId:                      application.RootId,
		VersionId:                   application.VersionId,
		Number:                      application.Number,
		CreatedAt:                   application.CreatedAt,
		VersionStartedAt:            application.VersionStartedAt,
		IsIncident:                  application.IsIncident,
		ParentRootId:                application.ParentRootId,
		UserLastEdited:              application.UserLastEdited,
		UserLastEditedOrganization:  application.UserLastEditedOrganization,
		Comment:                     application.Comment,
		CategoryId:                  application.CategoryId,
		DefectId:                    application.DefectId,
		IsDefectReturnable:          application.IsDefectReturnable,
		ApplicantDescription:        application.ApplicantDescription,
		ApplicantQuestion:           application.ApplicantQuestion,
		EmergencyType:               application.EmergencyType,
		Region:                      application.Region,
		District:                    application.District,
		Address:                     application.Address,
		UNOM:                        application.UNOM,
		Latitude:                    application.GPS.Latitude,
		Longitude:                   application.GPS.Longitude,
		Entrance:                    application.Entrance,
		Floor:                       application.Floor,
		Flat:                        application.Flat,
		OdsNumber:                   application.OdsNumber,
		ManagementCompanyTitle:      application.ManagementCompanyTitle,
		ExecutionCompanyTitle:       application.ExecutionCompanyTitle,
		RenderedServicesIds:         application.RenderedServicesIds,
		ConsumedMaterials:           application.ConsumedMaterials,
		RenderedSecurityServicesIds: application.RenderedSecurityServicesIds,
		ResultCode:                  application.ResultCode,
		AmountOfReturnings:          application.AmountOfReturnings,
		LastReturnedAt:              application.LastReturnedAt,
		IsAlarmed:                   application.IsAlarmed,
		ClosedAt:                    application.ClosedAt,
		PreferableFrom:              application.PreferableFrom,
		PreferableTo:                application.PreferableTo,
		RatedAt:                     application.RatedAt,
		Review:                      application.Review,
		RatingCode:                  application.RatingCode,
	}
}
