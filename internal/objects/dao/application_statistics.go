package dao

import (
	"time"

	"dudelkins/internal/objects/bo"
)

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

type AnomalyClassCounterWithCreationHour struct {
	CreationHour time.Time `bun:"creation_hour"`
	AnomalyClassCounter
}

type NormalAbnormalCounter struct {
	Region          string `bun:"region"`
	District        string `bun:"district"`
	AbnormalCounter int    `bun:"abnormal_counter"`
	NormalCounter   int    `bun:"normal_counter"`
}

func (c NormalAbnormalCounter) ToBo() bo.NormalAbnormalCounter {
	return bo.NormalAbnormalCounter{
		Region:          c.Region,
		District:        c.District,
		AbnormalCounter: c.AbnormalCounter,
		NormalCounter:   c.NormalCounter,
	}
}

type NormalAbnormalCounters []NormalAbnormalCounter

func (a NormalAbnormalCounters) ToBo() (counters bo.NormalAbnormalCounters) {
	counters = make(bo.NormalAbnormalCounters, 0, len(a))

	for _, counter := range a {
		counters = append(counters, counter.ToBo())
	}

	return
}
