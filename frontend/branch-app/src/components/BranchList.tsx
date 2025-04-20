import React, { useEffect, useState } from 'react';

export interface Branch {
  id: number;
  name: string;
  location: string;
}

export default function BranchList({ onSelect }: { onSelect: (id: number) => void }) {
  const [branches, setBranches] = useState<Branch[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetch('/api/branches')
      .then(res => res.json())
      .then(setBranches)
      .catch(() => setError('Failed to load branches'))
      .finally(() => setLoading(false));
  }, []);

  if (loading) return <div>Loading branches...</div>;
  if (error) return <div>{error}</div>;

  return (
    <div className="branch-list">
      <h3>Branches</h3>
      <ul>
        {branches.map(b => (
          <li key={b.id}>
            <button onClick={() => onSelect(b.id)}>{b.name} ({b.location})</button>
          </li>
        ))}
      </ul>
    </div>
  );
}
