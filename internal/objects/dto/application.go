package dto

import "time"

type Application struct {
	RootId                      int32      `json:"rootId"`
	VersionId                   int32      `json:"versionId"`
	Number                      string     `json:"number"`
	CreatedAt                   time.Time  `json:"createdAt"`
	VersionStartedAt            time.Time  `json:"versionStartedAt"`
	IsIncident                  bool       `json:"isIncident"`
	ParentRootId                *int32     `json:"parentRootId"`
	UserLastEdited              string     `json:"userLastEdited"`
	UserLastEditedOrganization  string     `json:"userLastEditedOrganization"`
	Comment                     string     `json:"comment"`
	CategoryId                  int        `json:"categoryId"`
	DefectId                    int        `json:"defectId"`
	IsDefectReturnable          bool       `json:"isDefectReturnable"`
	ApplicantDescription        string     `json:"applicantDescription"`
	ApplicantQuestion           string     `json:"applicantQuestion"`
	EmergencyType               string     `json:"emergencyType"`
	Region                      string     `json:"region"`
	District                    string     `json:"district"`
	Address                     string     `json:"address"`
	UNOM                        int64      `json:"UNOM"`
	Entrance                    *int       `json:"entrance"`
	Floor                       *int       `json:"floor"`
	Flat                        *int       `json:"flat"`
	OdsNumber                   string     `json:"odsNumber"`
	ManagementCompanyTitle      string     `json:"managementCompanyTitle"`
	ExecutionCompanyTitle       string     `json:"executionCompanyTitle"`
	RenderedServicesIds         []int      `json:"renderedServicesIds"`
	ConsumedMaterials           string     `json:"consumedMaterials"`
	RenderedSecurityServicesIds []int      `json:"renderedSecurityServicesIds"`
	ResultCode                  string     `json:"resultCode"`
	AmountOfReturnings          *int       `json:"amountOfReturnings"`
	LastReturnedAt              *time.Time `json:"lastReturnedAt"`
	IsAlarmed                   bool       `json:"isAlarmed"`
	ClosedAt                    time.Time  `json:"closedAt"`
	PreferableFrom              *time.Time `json:"preferableFrom"`
	PreferableTo                *time.Time `json:"preferableTo"`
	RatedAt                     *time.Time `json:"ratedAt"`
	Verdict                     *string    `json:"verdict"`
	RatingCode                  *string    `json:"ratingCode"`
}
