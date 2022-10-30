package bo

import (
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"dudelkins/internal/objects/dto"
)

func convertStringToIntArray(str string) (array []int, err error) {
	array = make([]int, 0, 0)
	if str == "" {
		return array, nil
	}

	stringSplitted := strings.Split(str, ",")
	for _, part := range stringSplitted {
		number, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		array = append(array, number)
	}
	return
}

// indexes of columns in comment

type Application struct {
	RootId                     int32     // 0
	VersionId                  int32     // 1 ?
	Number                     string    // 2
	CreatedAt                  time.Time // 4
	VersionStartedAt           time.Time // 5 ?
	IsIncident                 bool      // 9 Да|Нет
	ParentRootId               *int32    // 10 ? в каких случая инцидентные повторные заявки аномальны ?
	UserLastEdited             string    // 12
	UserLastEditedOrganization string    // 13
	Comment                    string    // 14
	CategoryId                 int       // 16
	DefectId                   int       // 20
	IsDefectReturnable         bool      // 22 Да|Нет Не нашел чтобы создавались заявки из этого ?
	ApplicantDescription       string    // 23
	ApplicantQuestion          string    // 24
	EmergencyType              string    // 26 normal|emergency
	Region                     string    // 27 ЗАО, ЮВАО... Можно юзать код из 28
	District                   string    // 29 Можно юзать код из 30
	Address                    string    // 31
	UNOM                       int64     //32

	GPS struct {
		Latitude, Longitude float64
	}

	Entrance               *string // 33 Может не присутствовать если проблема домовая?
	Floor                  *string // 34
	Flat                   *string // 35
	OdsNumber              string  // 36
	ManagementCompanyTitle string  // 37
	ExecutionCompanyTitle  string  // 38 есть идентификатор - 39, ИНН - 40
	// 43 - 46 что это?
	RenderedServicesIds         []int      // 48
	ConsumedMaterials           string     // 49 Израсходованный материал. Аномально если израсходован там где не надо?
	RenderedSecurityServicesIds []int      // 51 Охранные мероприятия. Проведены но описание другое?    429 строка
	ResultCode                  string     // 53 reject|resolved|consulted . 453 строка Консалтед но плохо?
	AmountOfReturnings          *int       // 54
	LastReturnedAt              *time.Time // 55
	// 56 Признак нахождения на доработке узнать что это?
	IsAlarmed      bool       // 57 да|нет Непонятно что это ??
	ClosedAt       time.Time  // 58
	PreferableFrom *time.Time // 59
	PreferableTo   *time.Time // 60
	RatedAt        *time.Time // 61
	Review         *string    // 62 Работы выполнены | Работы невыполнены
	RatingCode     *string    // 63 // bad|neutral|good
	// 64-67 Payment Fields

	// service fields
	IsAbnormal     *bool
	AnomalyClasses map[string]AnomalyClass
}

func NewApplicationFromRecord(record []string) (a Application, err error) {
	timeLayoutWithNoT := "2006-01-02 15:04:05-07"

	rootId, err := strconv.Atoi(record[0])
	if err != nil {
		return Application{}, errors.Wrap(err, "NewApplicationFromRecord")
	}
	a.RootId = int32(rootId)

	versionId, err := strconv.Atoi(record[1])
	if err != nil {
		return Application{}, errors.Wrap(err, "NewApplicationFromRecord")
	}
	a.VersionId = int32(versionId)

	a.Number = record[2]

	a.CreatedAt, err = time.Parse(timeLayoutWithNoT, record[4])
	if err != nil {
		return Application{}, errors.Wrap(err, "NewApplicationFromRecord 4")
	}

	a.VersionStartedAt, err = time.Parse(timeLayoutWithNoT, record[5])
	if err != nil {
		return Application{}, errors.Wrap(err, "NewApplicationFromRecord 5")
	}

	switch strings.ToLower(record[9]) {
	case "да":
		a.IsIncident = true
	case "нет":
		a.IsIncident = false
	default:
		return Application{}, errors.Wrap(errors.Errorf("incorrect IsIncident field: %s", record[9]), "NewApplicationFromRecord")
	}

	if record[10] != "" {
		parentRootId, err := strconv.Atoi(record[10])
		if err != nil {
			return Application{}, errors.Wrap(err, "NewApplicationFromRecord")
		}
		parentRootId32 := int32(parentRootId)
		a.ParentRootId = &parentRootId32
	}

	a.UserLastEdited = record[12]
	a.UserLastEditedOrganization = record[13]
	a.Comment = record[14]

	a.CategoryId, err = strconv.Atoi(record[16])
	if err != nil {
		return Application{}, errors.Wrap(err, "NewApplicationFromRecord")
	}
	a.DefectId, err = strconv.Atoi(record[20])
	if err != nil {
		return Application{}, errors.Wrap(err, "NewApplicationFromRecord")
	}

	switch strings.ToLower(record[22]) {
	case "да":
		a.IsDefectReturnable = true
	case "нет":
		a.IsDefectReturnable = false
	default:
		return Application{}, errors.Wrap(errors.Errorf("incorrect IsDefectReturnable field: %s", record[22]), "NewApplicationFromRecord")
	}

	a.ApplicantDescription = record[23]
	a.ApplicantQuestion = record[24]
	a.EmergencyType = record[26]
	a.Region = record[27]
	a.District = record[29]
	a.Address = record[31]

	unom, err := strconv.Atoi(record[32])
	if err != nil {
		return Application{}, errors.Wrap(err, "NewApplicationFromRecord")
	}
	a.UNOM = int64(unom)

	// TODO: GPS FIELD

	if record[33] != "" {
		a.Entrance = &record[33]
	}
	if record[34] != "" {
		a.Floor = &record[34]
	}
	if record[35] != "" {
		a.Flat = &record[35]
	}

	a.OdsNumber = record[36]
	a.ManagementCompanyTitle = record[37]
	a.ExecutionCompanyTitle = record[38]

	a.RenderedServicesIds, err = convertStringToIntArray(record[48])
	a.RenderedSecurityServicesIds, err = convertStringToIntArray(record[51])

	a.ConsumedMaterials = record[49]
	a.ResultCode = record[53]

	if record[54] != "" {
		amountOfReturnings, err := strconv.Atoi(record[54])
		if err != nil {
			return Application{}, errors.Wrap(err, "NewApplicationFromRecord")
		}
		a.AmountOfReturnings = &amountOfReturnings
	}

	if record[55] != "" {
		lastReturnedAt, err := time.Parse(timeLayoutWithNoT, record[55])
		if err != nil {
			return Application{}, errors.Wrap(err, "NewApplicationFromRecord 55")
		}
		a.LastReturnedAt = &lastReturnedAt
	}

	switch strings.ToLower(record[57]) {
	case "да":
		a.IsAlarmed = true
	case "нет":
		a.IsAlarmed = false
	default:
		return Application{}, errors.Wrap(errors.Errorf("incorrect IsAlarmed field: %s", record[57]), "NewApplicationFromRecord")
	}

	a.ClosedAt, err = time.Parse(timeLayoutWithNoT, record[58])
	if err != nil {
		return Application{}, errors.Wrap(err, "NewApplicationFromRecord")
	}
	if record[59] != "" {
		preferableFrom, err := time.Parse(timeLayoutWithNoT, record[59])
		if err != nil {
			return Application{}, errors.Wrap(err, "NewApplicationFromRecord 59")
		}
		a.PreferableFrom = &preferableFrom
	}
	if record[60] != "" {
		preferableTo, err := time.Parse(timeLayoutWithNoT, record[60])
		if err != nil {
			return Application{}, errors.Wrap(err, "NewApplicationFromRecord 60")
		}
		a.PreferableTo = &preferableTo
	}
	if record[61] != "" {
		ratedAt, err := time.Parse(timeLayoutWithNoT, record[61])
		if err != nil {
			return Application{}, errors.Wrap(errors.Errorf("record[61]: %s err: %s", record[61], err), "NewApplicationFromRecord 61")
		}
		a.RatedAt = &ratedAt
	}

	if record[62] != "" {
		a.Review = &record[62]
	}
	if record[63] != "" {
		a.RatingCode = &record[63]
	}

	return

}

func (a Application) ToDto() dto.Application {
	application := dto.Application{
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
		Verdict:                     a.Review,
		RatingCode:                  a.RatingCode,
		IsAbnormal:                  a.IsAbnormal,
	}

	application.AnomalyClasses = make(map[string]dto.AnomalyClass, len(a.AnomalyClasses))
	for className, class := range a.AnomalyClasses {
		application.AnomalyClasses[className] = dto.AnomalyClass{Description: class.Description, Verdict: class.Verdict}
	}

	return application
}

type Applications []Application

func (a Applications) ToDto() (applications []dto.Application) {
	applications = make([]dto.Application, 0, len(a))

	for _, application := range a {
		applications = append(applications, application.ToDto())
	}

	return
}
