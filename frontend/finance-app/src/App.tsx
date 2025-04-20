import React from 'react';
import InvoiceList from './components/InvoiceList';
import './index.css';

export default function App() {
  return (
    <div className="finance-app-container">
      <h2>Finance Management</h2>
      <InvoiceList />
    </div>
  );
}
