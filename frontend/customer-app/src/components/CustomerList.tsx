import React, { useEffect, useState } from 'react';
import { customerApi } from '../services/api';
import { Customer } from '../types/customer';

const listStyles = {
  container: {
    padding: '1rem',
    backgroundColor: 'white',
    borderRadius: '8px',
    boxShadow: '0 2px 4px rgba(0,0,0,0.1)',
    width: '100%',
    maxWidth: '500px',
  },
  header: {
    borderBottom: '1px solid #eee',
    paddingBottom: '0.5rem',
    marginBottom: '1rem',
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
  },
  list: {
    listStyle: 'none',
    padding: 0,
    margin: 0,
  },
  listItem: {
    padding: '0.75rem',
    borderBottom: '1px solid #f5f5f5',
    cursor: 'pointer',
    transition: 'background-color 0.2s',
    ':hover': {
      backgroundColor: '#f9f9f9',
    },
  },
  button: {
    background: 'none',
    border: 'none',
    padding: '0.5rem',
    textAlign: 'left' as const,
    width: '100%',
    cursor: 'pointer',
    borderRadius: '4px',
    transition: 'background-color 0.2s',
  },
  buttonHover: {
    backgroundColor: '#f0f0f0',
  },
  refreshButton: {
    backgroundColor: '#3498db',
    color: 'white',
    border: 'none',
    padding: '0.5rem 0.75rem',
    borderRadius: '4px',
    cursor: 'pointer',
  },
  error: {
    color: '#e74c3c',
    padding: '1rem',
    backgroundColor: '#fff3f3',
    borderRadius: '4px',
    marginBottom: '1rem',
  },
  loading: {
    padding: '2rem',
    textAlign: 'center' as const,
  },
};

export default function CustomerList({ onSelect }: { onSelect: (id: number) => void }) {
  const [customers, setCustomers] = useState<Customer[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchCustomers = async () => {
    try {
      setLoading(true);
      setError(null);
      const data = await customerApi.getCustomers();
      setCustomers(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to load customers');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchCustomers();
  }, []);

  if (loading) return <div style={listStyles.loading}>Loading customers...</div>;

  return (
    <div style={listStyles.container}>
      <div style={listStyles.header}>
        <h3>Customers</h3>
        <button
          style={listStyles.refreshButton}
          onClick={fetchCustomers}
          aria-label="Refresh customers"
        >
          Refresh
        </button>
      </div>

      {error && <div style={listStyles.error}>{error}</div>}

      {customers.length === 0 ? (
        <p>No customers found.</p>
      ) : (
        <ul style={listStyles.list}>
          {customers.map(c => (
            <li key={c.id} style={listStyles.listItem}>
              <button
                style={listStyles.button}
                onClick={() => onSelect(c.id)}
              >
                <strong>{c.name}</strong>
                <div>{c.phone}</div>
                {c.email && <div style={{ fontSize: '0.9rem', color: '#666' }}>{c.email}</div>}
              </button>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}
