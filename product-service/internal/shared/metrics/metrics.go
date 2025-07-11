package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var ProductsAddedCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "products_added_total",
	Help: "Total number of products added",
})

var ProductsUpdatedCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "products_updated_total",
	Help: "Total number of products updated",
})

var ProductsDeletedCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "products_deleted_total",
	Help: "Total number of products deleted",
})

var StocksAddedCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "stocks_added_total",
	Help: "Total number of stocks added",
})

var StocksUpdatedCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "stocks_updated_total",
	Help: "Total number of stocks updated",
})

var StocksDeletedCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "stocks_deleted_total",
	Help: "Total number of stocks deleted",
})

func init() {
	prometheus.MustRegister(ProductsAddedCounter)
	prometheus.MustRegister(ProductsUpdatedCounter)
	prometheus.MustRegister(ProductsDeletedCounter)
	prometheus.MustRegister(StocksAddedCounter)
	prometheus.MustRegister(StocksUpdatedCounter)
	prometheus.MustRegister(StocksDeletedCounter)
}
