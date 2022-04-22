package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	ProcessBusiness = promauto.NewCounter(prometheus.CounterOpts{
		Name: "process_buiness",
		Help: "The total number of processed events",
	})
)
