package infrastructure

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/pkg/errors"

	"dudelkins/internal/controllers"
	"dudelkins/internal/environment"
)

type IInjector interface {
	InjectApplicationController() *controllers.ApplicationController
}

type Kernel struct {
	env environment.Environment

	DefectIdsDuration  map[int]*time.Duration
	DefectIdsDeviation map[string]map[int]time.Duration
	DbHandler          *PostgresDatabaseClient
}

func Inject(ctx context.Context, env environment.Environment) (k *Kernel, err error) {
	k = &Kernel{}
	k.env = env

	postgresDatabaseClient, err := initPostgresClient(env.Postgres)
	if err != nil {
		return nil, errors.Wrap(err, "Inject")
	}
	k.DbHandler = postgresDatabaseClient

	if err = k.InitDefectIdsDuration(); err != nil {
		return nil, errors.Wrap(err, "Inject")
	}
	if err = k.InitDefectIdsDeviation(); err != nil {
		return nil, errors.Wrap(err, "Inject")
	}

	return
}

func (k *Kernel) InjectApplicationController() *controllers.ApplicationController {
	return &controllers.ApplicationController{
		ApplicationUploadService: k.InjectApplicationUploadService(),
		ApplicationViewService:   k.InjectApplicationViewService(),
	}
}

func (k *Kernel) InitDefectIdsDuration() (err error) {
	defectIdsFile, err := os.Open("/anomalius/cmd/defectIdsDuration.json")
	if err != nil {
		return errors.Wrap(err, "InitDefectIdsDuration")
	}
	defer defectIdsFile.Close()

	k.DefectIdsDuration = make(map[int]*time.Duration)
	if err = json.NewDecoder(defectIdsFile).Decode(&k.DefectIdsDuration); err != nil {
		return errors.Wrap(err, "InitDefectIdsDuration")
	}

	return
}

func (k *Kernel) InitDefectIdsDeviation() (err error) {
	defectIdsFile, err := os.Open("/anomalius/cmd/defectIdsDeviation.json")
	if err != nil {
		return errors.Wrap(err, "InitDefectIdsDeviation")
	}
	defer defectIdsFile.Close()

	var m map[string]map[int]int
	if err = json.NewDecoder(defectIdsFile).Decode(&m); err != nil {
		return errors.Wrap(err, "InitDefectIdsDeviation")
	}

	k.DefectIdsDeviation = make(map[string]map[int]time.Duration)
	for status, defectIdMap := range m {
		k.DefectIdsDeviation[status] = make(map[int]time.Duration)

		for defectId, durationSeconds := range defectIdMap {
			k.DefectIdsDeviation[status][defectId] = time.Second * time.Duration(durationSeconds)
		}
	}

	return
}
