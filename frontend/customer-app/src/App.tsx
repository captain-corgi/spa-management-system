import React, { useEffect, useState } from 'react';
import CustomerList from './components/CustomerList';
import CustomerDetail from './components/CustomerDetail';
import './index.css';

export default function App() {
  const [selectedCustomer, setSelectedCustomer] = useState<number | null>(null);

  return (
    <div className="customer-app-container">
      <h2>Customer Management</h2>
      <div className="customer-app-content">
        <CustomerList onSelect={setSelectedCustomer} />
        <CustomerDetail customerId={selectedCustomer} />
      </div>
    </div>
  );
}
