import React, { useState } from 'react';
import BranchList from './components/BranchList';
import BranchDetail from './components/BranchDetail';
import './index.css';

export default function App() {
  const [selectedBranch, setSelectedBranch] = useState<number | null>(null);
  return (
    <div className="branch-app-container">
      <h2>Branch Management</h2>
      <div className="branch-app-content">
        <BranchList onSelect={setSelectedBranch} />
        <BranchDetail branchId={selectedBranch} />
      </div>
    </div>
  );
}
