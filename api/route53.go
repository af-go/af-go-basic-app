package api

import (
	"github.com/af-go/basic-app/pkg/model"
	sdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewRoute53APIManager(provider *Route53Provider) *Route53APIManager {
	return &Route53APIManager{Provider: provider}
}

type Route53APIManager struct {
	Provider *Route53Provider
}

func (m *Route53APIManager) Build(engine *gin.Engine) {
	engine.GET("/list", m.OnList)
}

// OnPing ping
// @Produce json
// @Summary ping
// @Description check status
// @Success 200 {object} model.StatusResponse
// @Failure 400 {object} model.HTTPError
// @Failure 500 {object} model.HTTPError
// @Failure 503 {object} model.HTTPError
// @Router /ping [get]
func (m *Route53APIManager) OnList(gc *gin.Context) {
	m.Provider.GetCallerIdentity()
	statusCode := 200
	var resp model.ListHostedZonesResponse
	zones, err := m.Provider.ListHostedZones()
	if err != nil {
		statusCode = 500
		gc.JSON(statusCode, &err)
	} else {
		resp.HostedZones = zones
		gc.JSON(statusCode, &resp)
	}
}

func NewRoute53Provider() *Route53Provider {
	cfg := sdk.Config{
		LogLevel: sdk.LogLevel(sdk.LogDebug),
	}
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config:            cfg,
	}))
	return &Route53Provider{session: sess}
}

type Route53Provider struct {
	session *session.Session
}

func (p *Route53Provider) ListHostedZones() ([]model.HostedZone, error) {

	inpiut := &route53.ListHostedZonesInput{}
	svc := route53.New(p.session)
	output, err := svc.ListHostedZones(inpiut)
	if err != nil {
		logrus.WithError(err).Error("failed to list hosted zones")
		return nil, err
	}
	result := []model.HostedZone{}
	for _, zone := range output.HostedZones {
		result = append(result, model.HostedZone{Id: *zone.Id, Name: *zone.Name})
	}
	return result, nil
}

func (p *Route53Provider) GetCallerIdentity() error {
	input := &sts.GetCallerIdentityInput{}
	svc := sts.New(p.session)
	output, err := svc.GetCallerIdentity(input)
	if err != nil {
		logrus.WithError(err).Error("failed to get caller identity")
		return err
	}
	logrus.Infof("caller %s", *output.Arn)
	return nil
}
