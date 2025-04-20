import React, { useEffect, useState } from 'react';

export interface Staff {
  id: number;
  name: string;
  role: string;
}

export default function StaffList({ onSelect }: { onSelect: (id: number) => void }) {
  const [staff, setStaff] = useState<Staff[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetch('/api/staff')
      .then(res => res.json())
      .then(setStaff)
      .catch(() => setError('Failed to load staff'))
      .finally(() => setLoading(false));
  }, []);

  if (loading) return <div>Loading staff...</div>;
  if (error) return <div>{error}</div>;

  return (
    <div className="staff-list">
      <h3>Staff</h3>
      <ul>
        {staff.map(s => (
          <li key={s.id}>
            <button onClick={() => onSelect(s.id)}>{s.name} ({s.role})</button>
          </li>
        ))}
      </ul>
    </div>
  );
}
