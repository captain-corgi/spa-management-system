import React from 'react';
import Dashboard from './components/Dashboard';
import './index.css';

export default function App() {
  return (
    <div className="analytics-app-container">
      <h2>Analytics Dashboard</h2>
      <Dashboard />
    </div>
  );
}
