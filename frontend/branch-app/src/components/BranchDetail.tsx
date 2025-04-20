import React, { useEffect, useState } from 'react';
import type { Branch } from './BranchList';

export default function BranchDetail({ branchId }: { branchId: number | null }) {
  const [branch, setBranch] = useState<Branch | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!branchId) return;
    setLoading(true);
    fetch(`/api/branches/${branchId}`)
      .then(res => res.json())
      .then(setBranch)
      .catch(() => setError('Failed to load branch'))
      .finally(() => setLoading(false));
  }, [branchId]);

  if (!branchId) return <div>Select a branch to view details.</div>;
  if (loading) return <div>Loading branch...</div>;
  if (error) return <div>{error}</div>;
  if (!branch) return null;

  return (
    <div className="branch-detail">
      <h3>Branch Detail</h3>
      <div><b>Name:</b> {branch.name}</div>
      <div><b>Location:</b> {branch.location}</div>
      {/* Add more fields as needed */}
    </div>
  );
}
