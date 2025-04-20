import React, { useEffect, useState } from 'react';

interface Metric {
  label: string;
  value: string | number;
}

export default function Dashboard() {
  const [metrics, setMetrics] = useState<Metric[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetch('/api/analytics/metrics')
      .then(res => res.json())
      .then(setMetrics)
      .catch(() => setError('Failed to load metrics'))
      .finally(() => setLoading(false));
  }, []);

  if (loading) return <div>Loading dashboard...</div>;
  if (error) return <div>{error}</div>;

  return (
    <div className="dashboard">
      <h3>Key Metrics</h3>
      <ul>
        {metrics.map((m, idx) => (
          <li key={idx}><b>{m.label}:</b> {m.value}</li>
        ))}
      </ul>
    </div>
  );
}
