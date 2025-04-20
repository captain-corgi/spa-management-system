import React, { useState } from 'react';
import AppointmentList from './components/AppointmentList';
import AppointmentDetail from './components/AppointmentDetail';
import './index.css';

export default function App() {
  const [selectedAppointment, setSelectedAppointment] = useState<number | null>(null);
  return (
    <div className="appointment-app-container">
      <h2>Appointment Management</h2>
      <div className="appointment-app-content">
        <AppointmentList onSelect={setSelectedAppointment} />
        <AppointmentDetail appointmentId={selectedAppointment} />
      </div>
    </div>
  );
}
