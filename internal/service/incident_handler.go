package service

// Loads data from a csv entry and transform to a proto model struct
import (
	"time"

	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"

	repository "github.com/lucassouzavieira/go-grpc-server/internal/repository"
	"github.com/lucassouzavieira/go-grpc-server/pkg/protobuf/incident"
)

type IncidentHandler struct {
	r *repository.Repository
	i []*incident.Incident
}

func NewIncidentHandler(r *repository.Repository) (*IncidentHandler, error) {
	data, err := r.GetData()

	if err != nil {
		log.Fatal(err)
	}

	var incidents []*incident.Incident
	for i, v := range data {
		if i == 0 {
			continue // skip headers
		}

		incident := incidentFromCsv(v)
		incidents = append(incidents, &incident)
	}

	return &IncidentHandler{
		r: r,
		i: incidents,
	}, nil
}

func (h *IncidentHandler) GetIncidents() ([]*incident.Incident, error) {
	return h.i, nil
}

func incidentFromCsv(s []string) incident.Incident {
	number, err := decimal.NewFromString(s[0])

	if err != nil {
		log.WithFields(log.Fields{
			"error": "Error trying to parse Incident number",
		}).WithError(err)
	}

	year, err := decimal.NewFromString(s[2])

	if err != nil {
		log.WithFields(log.Fields{
			"error": "Error trying to parse Incident year",
		}).WithError(err)
	}

	pumpCount, err := decimal.NewFromString(s[5])

	if err != nil {
		log.WithFields(log.Fields{
			"error": "Error trying to parse Incident pumpCount",
		}).WithError(err)
	}

	pumpHoursCount, err := decimal.NewFromString(s[6])
	if err != nil {
		log.WithFields(log.Fields{
			"error": "Error trying to parse Incident pumpHoursCount",
		}).WithError(err)
	}

	latitude, err := decimal.NewFromString(s[29])
	if err != nil {
		log.WithFields(log.Fields{
			"error": "Error trying to parse Incident latitude",
		}).WithError(err)
	}

	longitude, err := decimal.NewFromString(s[30])
	if err != nil {
		log.WithFields(log.Fields{
			"error": "Error trying to parse Incident longitude",
		}).WithError(err)
	}

	lat, _ := latitude.Float64()
	long, _ := longitude.Float64()

	date_time, err := time.Parse("01/12/2006 10:45", s[1])
	if err != nil {
		log.WithFields(log.Fields{
			"error": "Error trying to parse Incident call date and time",
		}).WithError(err)
	}

	return incident.Incident{
		Number: number.IntPart(),
		CallDatetime: &timestamppb.Timestamp{
			Seconds: date_time.Unix(),
			Nanos:   int32(date_time.UnixNano()),
		},
		Year:                 int32(year.IntPart()),
		Type:                 s[4],
		PumpCount:            int32(pumpCount.IntPart()),
		PumpHoursTotal:       int32(pumpHoursCount.IntPart()),
		IncidentHourlyCost:   0,
		IncidentNotionalCost: 0,
		AnimalGroup:          s[10],
		FinalDescription:     s[9],
		Origin: &incident.Incident_Call{
			Origin: s[11],
		},
		Property: &incident.Incident_Property{
			Type:     s[12],
			Category: s[13],
		},
		SpecialService: &incident.Incident_SpecialService{
			Type:     s[15],
			Category: s[14],
		},
		Ward: &incident.Incident_Ward{
			Code: s[16],
			Name: s[17],
		},
		Address: &incident.Incident_Address{
			BoroughInfo: &incident.Incident_Address_Borough{
				Code:          s[18],
				Name:          s[19],
				StnGroundName: s[20],
			},
			Street:           s[21],
			Usrn:             s[23],
			PostcodeDistrict: s[24],
			Latitude:         lat,
			Longitude:        long,
		},
	}
}
