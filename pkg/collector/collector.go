package collector

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

type myMetrics struct {
	fooMetric prometheus.Gauge
	barMetric prometheus.Gauge
}

func new() myMetrics {
	mm := myMetrics{
		fooMetric: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "foo_metric",
				Help: "Shows whether a foo has occurred in our cluster",
			},
		),
		barMetric: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "bar_metric",
				Help: "Shows whether a bar has occurred in our cluster",
			},
		),
	}

	//Register metrics with prometheus
	prometheus.MustRegister(mm.fooMetric)
	prometheus.MustRegister(mm.barMetric)

	mm.barMetric.Set(0)
	mm.fooMetric.Set(0)

	// Return instance of struct
	return mm
}

func Collect() {
	//Create new metrics struct and seed with zero values
	mm := new()

	// Update values every 30s (simply increment them by 2)
	ticker := time.NewTicker(30 * time.Second)

	for range ticker.C {
		log.Info("Updating metric values")
		mm.fooMetric.Add(2)
		mm.barMetric.Add(2)
	}

}
