import React, { useEffect, useState } from 'react';
import type { Staff } from './StaffList';

export default function StaffDetail({ staffId }: { staffId: number | null }) {
  const [staff, setStaff] = useState<Staff | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!staffId) return;
    setLoading(true);
    fetch(`/api/staff/${staffId}`)
      .then(res => res.json())
      .then(setStaff)
      .catch(() => setError('Failed to load staff'))
      .finally(() => setLoading(false));
  }, [staffId]);

  if (!staffId) return <div>Select a staff member to view details.</div>;
  if (loading) return <div>Loading staff...</div>;
  if (error) return <div>{error}</div>;
  if (!staff) return null;

  return (
    <div className="staff-detail">
      <h3>Staff Detail</h3>
      <div><b>Name:</b> {staff.name}</div>
      <div><b>Role:</b> {staff.role}</div>
      {/* Add more fields as needed */}
    </div>
  );
}
