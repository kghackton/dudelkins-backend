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
	UNOM                        int64      `bun:"unom"`
	Latitude                    float64    `bun:"latitude"`
	Longitude                   float64    `bun:"longitude"`
	Entrance                    *string    `bun:"entrance"`
	Floor                       *string    `bun:"floor"`
	Flat                        *string    `bun:"flat"`
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

	IsAbnormal *bool `bun:"is_abnormal"`
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

func (a Application) ToBo() (application bo.Application) {
	application = bo.Application{
		RootId:                      a.RootId,
		VersionId:                   a.VersionId,
		Number:                      a.Number,
		CreatedAt:                   a.CreatedAt,
		VersionStartedAt:            a.VersionStartedAt,
		IsIncident:                  a.IsIncident,
		ParentRootId:                a.ParentRootId,
		UserLastEdited:              a.UserLastEdited,
		UserLastEditedOrganization:  a.UserLastEditedOrganization,
		Comment:                     a.Comment,
		CategoryId:                  a.CategoryId,
		DefectId:                    a.DefectId,
		IsDefectReturnable:          a.IsDefectReturnable,
		ApplicantDescription:        a.ApplicantDescription,
		ApplicantQuestion:           a.ApplicantQuestion,
		EmergencyType:               a.EmergencyType,
		Region:                      a.Region,
		District:                    a.District,
		Address:                     a.Address,
		UNOM:                        a.UNOM,
		Entrance:                    a.Entrance,
		Floor:                       a.Floor,
		Flat:                        a.Flat,
		OdsNumber:                   a.OdsNumber,
		ManagementCompanyTitle:      a.ManagementCompanyTitle,
		ExecutionCompanyTitle:       a.ExecutionCompanyTitle,
		RenderedServicesIds:         a.RenderedServicesIds,
		ConsumedMaterials:           a.ConsumedMaterials,
		RenderedSecurityServicesIds: a.RenderedSecurityServicesIds,
		ResultCode:                  a.ResultCode,
		AmountOfReturnings:          a.AmountOfReturnings,
		LastReturnedAt:              a.LastReturnedAt,
		IsAlarmed:                   a.IsAlarmed,
		ClosedAt:                    a.ClosedAt,
		PreferableFrom:              a.PreferableFrom,
		PreferableTo:                a.PreferableTo,
		RatedAt:                     a.RatedAt,
		Review:                      a.Review,
		RatingCode:                  a.RatingCode,
		IsAbnormal:                  a.IsAbnormal,
	}
	return application
}

type Applications []Application

func (a Applications) ToBo() (applications bo.Applications) {
	applications = make(bo.Applications, 0, len(a))

	for _, application := range a {
		applications = append(applications, application.ToBo())
	}

	return
}
