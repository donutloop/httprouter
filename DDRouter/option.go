package DDRouter

import (
	"math"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
)

type routerConfig struct {
	serviceName   string
	spanOpts      []ddtrace.StartSpanOption
	analyticsRate float64
}

// RouterOption represents an option that can be passed to New.
type RouterOption func(*routerConfig)

func defaults(cfg *routerConfig) {
	cfg.analyticsRate = AnalyticsRate()
	cfg.serviceName = "http.router"
}

// WithServiceName sets the given service name for the returned router.
func WithServiceName(name string) RouterOption {
	return func(cfg *routerConfig) {
		cfg.serviceName = name
	}
}

// WithSpanOptions applies the given set of options to the span started by the router.
func WithSpanOptions(opts ...ddtrace.StartSpanOption) RouterOption {
	return func(cfg *routerConfig) {
		cfg.spanOpts = opts
	}
}

// WithAnalytics enables Trace Analytics for all started spans.
func WithAnalytics(on bool) RouterOption {
	return func(cfg *routerConfig) {
		if on {
			cfg.analyticsRate = 1.0
		} else {
			cfg.analyticsRate = math.NaN()
		}
	}
}

// WithAnalyticsRate sets the sampling rate for Trace Analytics events
// correlated to started spans.
func WithAnalyticsRate(rate float64) RouterOption {
	return func(cfg *routerConfig) {
		if rate >= 0.0 && rate <= 1.0 {
			cfg.analyticsRate = rate
		} else {
			cfg.analyticsRate = math.NaN()
		}
	}
}
