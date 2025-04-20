import React, { useEffect, useState } from 'react';

export interface Invoice {
  id: number;
  amount: number;
  status: string;
}

export default function InvoiceList() {
  const [invoices, setInvoices] = useState<Invoice[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetch('/api/invoices')
      .then(res => res.json())
      .then(setInvoices)
      .catch(() => setError('Failed to load invoices'))
      .finally(() => setLoading(false));
  }, []);

  if (loading) return <div>Loading invoices...</div>;
  if (error) return <div>{error}</div>;

  return (
    <div className="invoice-list">
      <h3>Invoices</h3>
      <ul>
        {invoices.map(i => (
          <li key={i.id}>Invoice #{i.id}: ${i.amount} ({i.status})</li>
        ))}
      </ul>
    </div>
  );
}
