package application

import "../domain"

type AnalyticsService struct {
	repo domain.MetricRepository
}

func NewAnalyticsService(repo domain.MetricRepository) *AnalyticsService {
	return &AnalyticsService{repo: repo}
}

func (s *AnalyticsService) CreateMetric(metric *domain.Metric) error {
	return s.repo.Create(metric)
}

func (s *AnalyticsService) ListMetrics() ([]domain.Metric, error) {
	return s.repo.List()
}
