package bo

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
