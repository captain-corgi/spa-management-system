import React, { useEffect, useState } from 'react';
import type { Customer } from './CustomerList';

export default function CustomerDetail({ customerId }: { customerId: number | null }) {
  const [customer, setCustomer] = useState<Customer | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!customerId) return;
    setLoading(true);
    fetch(`/api/customers/${customerId}`)
      .then(res => res.json())
      .then(setCustomer)
      .catch(() => setError('Failed to load customer'))
      .finally(() => setLoading(false));
  }, [customerId]);

  if (!customerId) return <div>Select a customer to view details.</div>;
  if (loading) return <div>Loading customer...</div>;
  if (error) return <div>{error}</div>;
  if (!customer) return null;

  return (
    <div className="customer-detail">
      <h3>Customer Detail</h3>
      <div><b>Name:</b> {customer.name}</div>
      <div><b>Phone:</b> {customer.phone}</div>
      {/* Add more fields as needed */}
    </div>
  );
}
