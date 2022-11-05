package dao

import "dudelkins/internal/objects/bo"

type AnomalyClassCounter struct {
	Region                 string `bun:"region"`
	District               string `bun:"district"`
	ManagementCompanyTitle string `bun:"management_company_title"`
	AnomalyClass           string `bun:"anomaly_class"`
	AnomalyClassAmount     int    `bun:"counter"`
}

func (a AnomalyClassCounter) ToBo() bo.AnomalyClassCounter {
	return bo.AnomalyClassCounter{
		Region:                 a.Region,
		District:               a.District,
		ManagementCompanyTitle: a.ManagementCompanyTitle,
		AnomalyClass:           a.AnomalyClass,
		AnomalyClassAmount:     a.AnomalyClassAmount,
	}
}

type AnomalyClassCounters []AnomalyClassCounter

func (a AnomalyClassCounters) ToBo() (counters bo.AnomalyClassCounters) {
	counters = make(bo.AnomalyClassCounters, 0, len(a))

	for _, counter := range a {
		counters = append(counters, counter.ToBo())
	}

	return
}
