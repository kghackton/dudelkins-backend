package dto

import "time"

type ApplicationRetrieveOpts struct {
	ClosedFrom     *time.Time `query:"closedFrom"`
	ClosedTo       *time.Time `query:"closedTo"`
	IsAbnormal     *bool      `query:"isAbnormal"`
	AnomalyClasses []string   `query:"anomalyClass[]"`
	CategoryIds    []int      `query:"categoryId[]"`
	DefectIds      []int      `query:"defectId[]"`
	Region         *string    `query:"region"`
	District       *string    `query:"district"`
	UNOM           *int64     `query:"UNOM"`
	Limit          *int       `query:"limit"`
	Offset         *int       `query:"offset"`
}
