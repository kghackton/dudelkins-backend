package infrastructure

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"dudelkins/internal/controllers"
	"dudelkins/internal/environment"

	"github.com/pkg/errors"
)

type IInjector interface {
	InjectApplicationController() *controllers.ApplicationController
}

type Kernel struct {
	env environment.Environment

	DefectIdsDuration map[int]*time.Duration
	DbHandler         *PostgresDatabaseClient
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
