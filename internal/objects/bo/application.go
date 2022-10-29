package bo

import "time"

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
	UNOM                       int       //32

	GPS struct {
		Latitude, Longitude float64
	}

	Entrance               *int   // 33 Может не присутствовать если проблема домовая?
	Floor                  *int   // 34
	Flat                   *int   // 35
	OdsNumber              string // 36
	ManagementCompanyTitle string // 37
	ExecutionCompanyTitle  string // 38 есть идентификатор - 39, ИНН - 40
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
	RatedAt        time.Time  // 61
	Review         string     // 62 Работы выполнены | Работы невыполнены
	RatingCode     string     // 63 // bad|neutral|good
	// 64-67 Payment Fields
}
