import React, { useState } from 'react';
import StaffList from './components/StaffList';
import StaffDetail from './components/StaffDetail';
import './index.css';

export default function App() {
  const [selectedStaff, setSelectedStaff] = useState<number | null>(null);
  return (
    <div className="staff-app-container">
      <h2>Staff Management</h2>
      <div className="staff-app-content">
        <StaffList onSelect={setSelectedStaff} />
        <StaffDetail staffId={selectedStaff} />
      </div>
    </div>
  );
}
