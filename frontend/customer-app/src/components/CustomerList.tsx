import React, { useEffect, useState } from 'react';

export interface Customer {
  id: number;
  name: string;
  phone: string;
}

export default function CustomerList({ onSelect }: { onSelect: (id: number) => void }) {
  const [customers, setCustomers] = useState<Customer[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetch('/api/customers')
      .then(res => res.json())
      .then(setCustomers)
      .catch(() => setError('Failed to load customers'))
      .finally(() => setLoading(false));
  }, []);

  if (loading) return <div>Loading customers...</div>;
  if (error) return <div>{error}</div>;

  return (
    <div className="customer-list">
      <h3>Customers</h3>
      <ul>
        {customers.map(c => (
          <li key={c.id}>
            <button onClick={() => onSelect(c.id)}>{c.name} ({c.phone})</button>
          </li>
        ))}
      </ul>
    </div>
  );
}
