package service

// Loads data from a csv entry and transform to a proto model struct
import (
	"strings"

	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"

	repository "github.com/lucassouzavieira/go-grpc-server/internal/repository"
	"github.com/lucassouzavieira/go-grpc-server/pkg/protobuf/incident"
)

type GroupStats struct {
	animal_group string
	incidents    int32
}

type Stats struct {
	year      int32
	incidents int32
	groups    []*GroupStats
}

type IncidentHandler struct {
	r *repository.Repository
	i []*incident.Incident
}

func NewIncidentHandler(r *repository.Repository) (*IncidentHandler, error) {
	data, err := r.GetData()

	if err != nil {
		log.Fatal(err)
		return nil, err
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

func (h *IncidentHandler) Filter(fs []*repository.Filter) (bool, error) {
	data, err := h.r.Filter(fs)

	if err != nil {
		log.Fatal(err)
		return false, err
	}

	var ixs []*incident.Incident
	for _, v := range data {
		ix := incidentFromCsv(v)
		ixs = append(ixs, &ix)
	}

	h.i = ixs
	return true, nil
}

func (h *IncidentHandler) GetIncidents() ([]*incident.Incident, error) {
	return h.i, nil
}

func (h *IncidentHandler) GetIncidentsByAnimalGroup(group string) ([]*incident.Incident, error) {
	var filtered []*incident.Incident = make([]*incident.Incident, 0)

	for _, i := range h.i {
		if i.AnimalGroup == group {
			filtered = append(filtered, i)
		}
	}

	return filtered, nil
}

func (h *IncidentHandler) GetYearStats(year int32) (*Stats, error) {
	var fx = repository.Filter{
		Property: "CalYear",
		Value:    decimal.NewFromInt32(year).StringFixed(0),
	}

	var fs = make([]*repository.Filter, 0)
	fs = append(fs, &fx)

	_, err := h.Filter(fs)

	if err != nil {
		log.WithFields(log.Fields{
			"error": "Error trying to filter incidents",
		}).Warn(err)

		return nil, err
	}

	groups := make(map[string]int32)

	for _, i := range h.i {
		groups[i.GetAnimalGroup()] = groups[i.GetAnimalGroup()] + 1
	}

	group_stats := make([]*GroupStats, 0)

	for g, i := range groups {
		gs := GroupStats{
			animal_group: g,
			incidents:    i,
		}

		group_stats = append(group_stats, &gs)
	}

	return &Stats{
		year:      year,
		incidents: int32(len(h.i)),
		groups:    group_stats,
	}, nil
}

// Internal
func incidentFromCsv(s []string) incident.Incident {
	numberStr := strings.ReplaceAll(s[0], "-", "")
	number, err := decimal.NewFromString(numberStr)

	if err != nil {
		log.WithFields(log.Fields{
			"error": "Error trying to parse Incident number",
		}).Warn(err)
	}

	year, err := decimal.NewFromString(s[2])

	if err != nil {
		log.WithFields(log.Fields{
			"error": "Error trying to parse Incident year",
		}).Warn(err)
	}

	var pumpCount int32
	if s[5] != "NULL" {
		pumpCountRepresentation, err := decimal.NewFromString(s[5])
		if err != nil {
			log.WithFields(log.Fields{
				"error": "Error trying to parse Incident pumpCount",
			}).Warn(err)
		}

		pumpCount = int32(pumpCountRepresentation.IntPart())
	}

	var pumpHoursCount int32
	if s[5] != "NULL" {

		pumpHoursCountRepresentation, err := decimal.NewFromString(s[6])
		if err != nil {
			log.WithFields(log.Fields{
				"error": "Error trying to parse Incident pumpHoursCount",
			}).Warn(err)
		}

		pumpHoursCount = int32(pumpHoursCountRepresentation.IntPart())
	}

	var lat float64
	var long float64
	if s[29] != "NULL" {
		latitude, err := decimal.NewFromString(s[29])
		if err != nil {
			log.WithFields(log.Fields{
				"error": "Error trying to parse Incident latitude",
			}).Warn(err)
		}

		lat, _ = latitude.Float64()
	}

	if s[30] != "NULL" {
		latitude, err := decimal.NewFromString(s[29])
		if err != nil {
			log.WithFields(log.Fields{
				"error": "Error trying to parse Incident latitude",
			}).Warn(err)
		}

		long, _ = latitude.Float64()
	}

	return incident.Incident{
		Number:               number.IntPart(),
		CallDatetime:         s[1],
		Year:                 int32(year.IntPart()),
		Type:                 s[4],
		PumpCount:            pumpCount,
		PumpHoursTotal:       pumpHoursCount,
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
