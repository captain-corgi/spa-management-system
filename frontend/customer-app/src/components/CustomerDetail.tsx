import React, { useEffect, useState } from 'react';
import { customerApi } from '../services/api';
import { Customer } from '../types/customer';

const detailStyles = {
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
  field: {
    marginBottom: '0.75rem',
    padding: '0.5rem',
    borderBottom: '1px solid #f5f5f5',
  },
  label: {
    fontWeight: 'bold',
    marginRight: '0.5rem',
    color: '#555',
  },
  value: {
    color: '#333',
  },
  error: {
    color: '#e74c3c',
    padding: '1rem',
    backgroundColor: '#fff3f3',
    borderRadius: '4px',
  },
  loading: {
    padding: '2rem',
    textAlign: 'center' as const,
  },
  placeholder: {
    padding: '2rem',
    textAlign: 'center' as const,
    color: '#999',
    backgroundColor: '#f9f9f9',
    borderRadius: '8px',
    border: '1px dashed #ddd',
  },
  editButton: {
    backgroundColor: '#3498db',
    color: 'white',
    border: 'none',
    padding: '0.5rem 0.75rem',
    borderRadius: '4px',
    cursor: 'pointer',
  },
};

export default function CustomerDetail({ customerId }: { customerId: number | null }) {
  const [customer, setCustomer] = useState<Customer | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchCustomer = async (id: number) => {
    try {
      setLoading(true);
      setError(null);
      const data = await customerApi.getCustomerById(id);
      setCustomer(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to load customer');
      setCustomer(null);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (!customerId) {
      setCustomer(null);
      return;
    }
    fetchCustomer(customerId);
  }, [customerId]);

  if (!customerId) {
    return (
      <div style={detailStyles.placeholder}>
        <p>Select a customer to view details</p>
      </div>
    );
  }

  if (loading) {
    return <div style={detailStyles.loading}>Loading customer details...</div>;
  }

  if (error) {
    return (
      <div style={detailStyles.error}>
        <h3>Error</h3>
        <p>{error}</p>
        <button
          onClick={() => fetchCustomer(customerId)}
          style={detailStyles.editButton}
        >
          Try Again
        </button>
      </div>
    );
  }

  if (!customer) return null;

  return (
    <div style={detailStyles.container}>
      <div style={detailStyles.header}>
        <h3>Customer Details</h3>
        <button style={detailStyles.editButton}>Edit</button>
      </div>

      <div style={detailStyles.field}>
        <span style={detailStyles.label}>Name:</span>
        <span style={detailStyles.value}>{customer.name}</span>
      </div>

      <div style={detailStyles.field}>
        <span style={detailStyles.label}>Email:</span>
        <span style={detailStyles.value}>{customer.email || 'N/A'}</span>
      </div>

      <div style={detailStyles.field}>
        <span style={detailStyles.label}>Phone:</span>
        <span style={detailStyles.value}>{customer.phone}</span>
      </div>

      {customer.address && (
        <div style={detailStyles.field}>
          <span style={detailStyles.label}>Address:</span>
          <span style={detailStyles.value}>{customer.address}</span>
        </div>
      )}

      {customer.membershipLevel && (
        <div style={detailStyles.field}>
          <span style={detailStyles.label}>Membership:</span>
          <span style={detailStyles.value}>
            {customer.membershipLevel.charAt(0).toUpperCase() + customer.membershipLevel.slice(1)}
          </span>
        </div>
      )}

      {customer.joinDate && (
        <div style={detailStyles.field}>
          <span style={detailStyles.label}>Join Date:</span>
          <span style={detailStyles.value}>
            {new Date(customer.joinDate).toLocaleDateString()}
          </span>
        </div>
      )}

      {customer.lastVisit && (
        <div style={detailStyles.field}>
          <span style={detailStyles.label}>Last Visit:</span>
          <span style={detailStyles.value}>
            {new Date(customer.lastVisit).toLocaleDateString()}
          </span>
        </div>
      )}

      {customer.notes && (
        <div style={detailStyles.field}>
          <span style={detailStyles.label}>Notes:</span>
          <div style={{ ...detailStyles.value, marginTop: '0.5rem' }}>{customer.notes}</div>
        </div>
      )}
    </div>
  );
}
