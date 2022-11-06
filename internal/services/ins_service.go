package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	"dudelkins/internal/objects/bo"
)

type InsService struct {
	HttpClient *http.Client

	InsEndPoint string
}

func (s *InsService) IsAbnormal(application bo.Application) (isAbnormal bool, confidence float64, err error) {
	type request struct {
		IdDefect        string `json:"id_defect"`
		IdEmergency     string `json:"id_emergency"`
		IdDoneWorks     []int  `json:"id_done_works"`
		IdSecurityWorks []int  `json:"id_security_works"`
		Result          string `json:"result"`
	}

	r := request{
		IdDefect:        strconv.Itoa(application.DefectId),
		IdEmergency:     application.EmergencyType,
		IdDoneWorks:     application.RenderedServicesIds,
		IdSecurityWorks: application.RenderedSecurityServicesIds,
		Result:          application.ResultCode,
	}

	rMarshalled, err := json.Marshal(r)
	if err != nil {
		return isAbnormal, confidence, errors.Wrap(err, "IsAbnormal")
	}

	resp, err := s.HttpClient.Post(s.InsEndPoint, echo.MIMEApplicationJSON, bytes.NewBuffer(rMarshalled))
	if err != nil {
		return isAbnormal, confidence, errors.Wrap(err, "IsAbnormal")
	}
	defer resp.Body.Close()

	type response struct {
		IsAbnormal bool    `json:"isAbnormal"`
		Confidence float64 `json:"confidence"`
	}

	var respBody response
	if err = json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return isAbnormal, confidence, errors.Wrap(err, "IsAbnormal")
	}

	return respBody.IsAbnormal, respBody.Confidence, nil
}
