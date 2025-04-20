import React, { useState } from 'react';
import CustomerList from './components/CustomerList';
import CustomerDetail from './components/CustomerDetail';
import ErrorBoundary from './components/ErrorBoundary';
import './index.css';

const appStyles = {
  container: {
    padding: '1rem',
    maxWidth: '1200px',
    margin: '0 auto',
  },
  header: {
    borderBottom: '1px solid #eee',
    paddingBottom: '1rem',
    marginBottom: '2rem',
  },
  content: {
    display: 'grid',
    gridTemplateColumns: '1fr 1fr',
    gap: '2rem',
    '@media (max-width: 768px)': {
      gridTemplateColumns: '1fr',
    },
  },
};

export default function App() {
  const [selectedCustomer, setSelectedCustomer] = useState<number | null>(null);

  return (
    <div style={appStyles.container}>
      <div style={appStyles.header}>
        <h2>Customer Management</h2>
      </div>
      <div style={appStyles.content}>
        <ErrorBoundary>
          <CustomerList onSelect={setSelectedCustomer} />
        </ErrorBoundary>
        <ErrorBoundary>
          <CustomerDetail customerId={selectedCustomer} />
        </ErrorBoundary>
      </div>
    </div>
  );
}
