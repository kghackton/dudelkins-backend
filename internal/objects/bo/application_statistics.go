package bo

import "time"

type AnomalyClassCounter struct {
	Region                 string `bun:"region"`
	District               string `bun:"district"`
	ManagementCompanyTitle string `bun:"management_company_title"`
	AnomalyClass           string `bun:"anomaly_class"`
	AnomalyClassAmount     int    `bun:"counter"`
}

type AnomalyClassCounters []AnomalyClassCounter

func (a AnomalyClassCounters) ToMap() (m AnomalyClassCountersMap) {
	m = make(AnomalyClassCountersMap, 12)

	for _, anomalyClassCounter := range a {
		if _, exists := m[anomalyClassCounter.Region]; !exists {
			m[anomalyClassCounter.Region] = make(map[string]map[string]map[string]int, 100)
		}
		if _, exists := m[anomalyClassCounter.Region][anomalyClassCounter.District]; !exists {
			m[anomalyClassCounter.Region][anomalyClassCounter.District] = make(map[string]map[string]int, 100)
		}
		if _, exists := m[anomalyClassCounter.Region][anomalyClassCounter.District][anomalyClassCounter.ManagementCompanyTitle]; !exists {
			m[anomalyClassCounter.Region][anomalyClassCounter.District][anomalyClassCounter.ManagementCompanyTitle] = make(map[string]int, 100)
		}
		m[anomalyClassCounter.Region][anomalyClassCounter.District][anomalyClassCounter.ManagementCompanyTitle][anomalyClassCounter.AnomalyClass] = anomalyClassCounter.AnomalyClassAmount
	}

	return
}

// region    district  management anomalyClass
type AnomalyClassCountersMap map[string]map[string]map[string]map[string]int

type AnomalyClassCounterWithCreationHour struct {
	CreationHour time.Time
	AnomalyClassCounter
}

type AnomalyClassCountersWithCreationHour []AnomalyClassCounterWithCreationHour

func (a AnomalyClassCountersWithCreationHour) ToMap() (m AnomalyClassCountersWithCreationHourMap) {
	m = make(AnomalyClassCountersWithCreationHourMap, 12)

	for _, anomalyClassCounter := range a {
		if _, exists := m[anomalyClassCounter.Region]; !exists {
			m[anomalyClassCounter.Region] = make(map[string]map[string]map[string]map[time.Time]int, 100)
		}
		if _, exists := m[anomalyClassCounter.Region][anomalyClassCounter.District]; !exists {
			m[anomalyClassCounter.Region][anomalyClassCounter.District] = make(map[string]map[string]map[time.Time]int, 100)
		}
		if _, exists := m[anomalyClassCounter.Region][anomalyClassCounter.District][anomalyClassCounter.ManagementCompanyTitle]; !exists {
			m[anomalyClassCounter.Region][anomalyClassCounter.District][anomalyClassCounter.ManagementCompanyTitle] = make(map[string]map[time.Time]int, 100)
		}
		if _, exists := m[anomalyClassCounter.Region][anomalyClassCounter.District][anomalyClassCounter.ManagementCompanyTitle][anomalyClassCounter.AnomalyClass]; !exists {
			m[anomalyClassCounter.Region][anomalyClassCounter.District][anomalyClassCounter.ManagementCompanyTitle][anomalyClassCounter.AnomalyClass] = make(map[time.Time]int)
		}

		m[anomalyClassCounter.Region][anomalyClassCounter.District][anomalyClassCounter.ManagementCompanyTitle][anomalyClassCounter.AnomalyClass][anomalyClassCounter.CreationHour] = anomalyClassCounter.AnomalyClassAmount
	}

	return
}

// region    district   management anomalyClass creationHour
type AnomalyClassCountersWithCreationHourMap map[string]map[string]map[string]map[string]map[time.Time]int

type NormalAbnormalCounter struct {
	Region          string
	District        string
	AbnormalCounter int
	NormalCounter   int
}

type NormalAbnormalCounters []NormalAbnormalCounter

func (c NormalAbnormalCounters) ToMap() (m NormalAbnormalCountersMap) {
	m = make(NormalAbnormalCountersMap, 12)

	for _, normalAbnormalCounter := range c {
		if _, exists := m[normalAbnormalCounter.Region]; !exists {
			m[normalAbnormalCounter.Region] = make(map[string]NormalAbnormalAmountStruct, 100)
		}
		m[normalAbnormalCounter.Region][normalAbnormalCounter.District] = NormalAbnormalAmountStruct{
			Abnormal: normalAbnormalCounter.AbnormalCounter,
			Normal:   normalAbnormalCounter.NormalCounter,
		}
	}

	return
}

// region    district
type NormalAbnormalCountersMap map[string]map[string]NormalAbnormalAmountStruct

type NormalAbnormalAmountStruct struct {
	Abnormal int `json:"abnormal"`
	Normal   int `json:"normal"`
}
