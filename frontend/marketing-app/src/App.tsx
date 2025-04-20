import React from 'react';
import CampaignList from './components/CampaignList';
import './index.css';

export default function App() {
  return (
    <div className="marketing-app-container">
      <h2>Marketing Management</h2>
      <CampaignList />
    </div>
  );
}
