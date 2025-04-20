import React from 'react';
import { Link } from 'react-router-dom';
import { useAuth } from '../auth/AuthContext';

export default function Navbar() {
  const { isAuthenticated, login, logout } = useAuth();
  return (
    <nav style={{ display: 'flex', gap: '1rem', padding: '1rem', background: '#f5f5f5' }}>
      <Link to="/dashboard">Dashboard</Link>
      <Link to="/customers">Customers</Link>
      <Link to="/appointments">Appointments</Link>
      {/* Add more links as needed */}
      {isAuthenticated ? (
        <button onClick={logout}>Logout</button>
      ) : (
        <button onClick={login}>Login</button>
      )}
    </nav>
  );
}
