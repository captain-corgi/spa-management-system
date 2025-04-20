import React, { useEffect, useState } from 'react';
import type { Appointment } from './AppointmentList';

export default function AppointmentDetail({ appointmentId }: { appointmentId: number | null }) {
  const [appointment, setAppointment] = useState<Appointment | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!appointmentId) return;
    setLoading(true);
    fetch(`/api/appointments/${appointmentId}`)
      .then(res => res.json())
      .then(setAppointment)
      .catch(() => setError('Failed to load appointment'))
      .finally(() => setLoading(false));
  }, [appointmentId]);

  if (!appointmentId) return <div>Select an appointment to view details.</div>;
  if (loading) return <div>Loading appointment...</div>;
  if (error) return <div>{error}</div>;
  if (!appointment) return null;

  return (
    <div className="appointment-detail">
      <h3>Appointment Detail</h3>
      <div><b>Customer:</b> {appointment.customerName}</div>
      <div><b>Time:</b> {appointment.time}</div>
      {/* Add more fields as needed */}
    </div>
  );
}
