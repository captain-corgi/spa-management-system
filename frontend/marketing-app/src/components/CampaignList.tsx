import React, { useEffect, useState } from 'react';

export interface Campaign {
  id: number;
  name: string;
  status: string;
}

export default function CampaignList() {
  const [campaigns, setCampaigns] = useState<Campaign[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetch('/api/campaigns')
      .then(res => res.json())
      .then(setCampaigns)
      .catch(() => setError('Failed to load campaigns'))
      .finally(() => setLoading(false));
  }, []);

  if (loading) return <div>Loading campaigns...</div>;
  if (error) return <div>{error}</div>;

  return (
    <div className="campaign-list">
      <h3>Campaigns</h3>
      <ul>
        {campaigns.map(c => (
          <li key={c.id}>{c.name} ({c.status})</li>
        ))}
      </ul>
    </div>
  );
}
