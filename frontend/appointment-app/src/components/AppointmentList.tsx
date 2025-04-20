import React, { useEffect, useState } from 'react';

export interface Appointment {
  id: number;
  customerName: string;
  time: string;
}

export default function AppointmentList({ onSelect }: { onSelect: (id: number) => void }) {
  const [appointments, setAppointments] = useState<Appointment[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetch('/api/appointments')
      .then(res => res.json())
      .then(setAppointments)
      .catch(() => setError('Failed to load appointments'))
      .finally(() => setLoading(false));
  }, []);

  if (loading) return <div>Loading appointments...</div>;
  if (error) return <div>{error}</div>;

  return (
    <div className="appointment-list">
      <h3>Appointments</h3>
      <ul>
        {appointments.map(a => (
          <li key={a.id}>
            <button onClick={() => onSelect(a.id)}>{a.customerName} ({a.time})</button>
          </li>
        ))}
      </ul>
    </div>
  );
}
