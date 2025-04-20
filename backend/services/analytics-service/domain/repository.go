package domain

type MetricRepository interface {
	Create(metric *Metric) error
	List() ([]Metric, error)
}

type ReportRepository interface {
	Create(report *Report) error
	List() ([]Report, error)
}

type DashboardRepository interface {
	Create(dashboard *Dashboard) error
	List() ([]Dashboard, error)
}
