import React from 'react';
import { Link, useLocation } from 'react-router-dom';
import { useAuth } from '../auth/AuthContext';

const navStyles = {
  container: {
    display: 'flex',
    flexWrap: 'wrap' as const,
    gap: '0.5rem',
    padding: '1rem',
    background: '#2c3e50',
    color: 'white',
    boxShadow: '0 2px 4px rgba(0,0,0,0.1)',
    justifyContent: 'space-between',
  },
  linkGroup: {
    display: 'flex',
    gap: '1rem',
    flexWrap: 'wrap' as const,
    alignItems: 'center',
  },
  link: {
    color: 'white',
    textDecoration: 'none',
    padding: '0.5rem 0.75rem',
    borderRadius: '4px',
    transition: 'background-color 0.2s',
  },
  activeLink: {
    backgroundColor: '#3498db',
  },
  authButton: {
    backgroundColor: '#e74c3c',
    color: 'white',
    border: 'none',
    padding: '0.5rem 1rem',
    borderRadius: '4px',
    cursor: 'pointer',
  },
  logo: {
    fontWeight: 'bold',
    fontSize: '1.2rem',
    marginRight: '1rem',
  }
};

export default function Navbar() {
  const { isAuthenticated, login, logout } = useAuth();
  const location = useLocation();

  const isActive = (path: string) => {
    return location.pathname.startsWith(path);
  };

  const getLinkStyle = (path: string) => {
    return {
      ...navStyles.link,
      ...(isActive(path) ? navStyles.activeLink : {}),
    };
  };

  return (
    <nav style={navStyles.container}>
      <div style={navStyles.linkGroup}>
        <span style={navStyles.logo}>Spa Management</span>
        <Link to="/dashboard" style={getLinkStyle('/dashboard')}>Dashboard</Link>
        <Link to="/customers" style={getLinkStyle('/customers')}>Customers</Link>
        <Link to="/appointments" style={getLinkStyle('/appointments')}>Appointments</Link>
        <Link to="/staff" style={getLinkStyle('/staff')}>Staff</Link>
        <Link to="/branches" style={getLinkStyle('/branches')}>Branches</Link>
        <Link to="/marketing" style={getLinkStyle('/marketing')}>Marketing</Link>
        <Link to="/finance" style={getLinkStyle('/finance')}>Finance</Link>
        <Link to="/analytics" style={getLinkStyle('/analytics')}>Analytics</Link>
      </div>
      <div>
        {isAuthenticated ? (
          <button style={navStyles.authButton} onClick={logout}>Logout</button>
        ) : (
          <button style={navStyles.authButton} onClick={login}>Login</button>
        )}
      </div>
    </nav>
  );
}
