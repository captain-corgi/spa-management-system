import React, { Suspense } from 'react';
import { Routes, Route, Navigate } from 'react-router-dom';
import { useAuth } from './shared/auth/AuthContext';
import Navbar from './shared/components/Navbar';
import ErrorBoundary from './shared/components/ErrorBoundary';
import LoginPage from './shared/components/LoginPage';

// Lazy load micro frontends
const CustomerApp = React.lazy(() => import('customerApp/App'));
const AppointmentApp = React.lazy(() => import('appointmentApp/App'));
// Add other micro frontends as needed

// Loading component with better styling
const LoadingFallback = () => (
  <div style={{
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    height: 'calc(100vh - 60px)',
    fontSize: '1.2rem'
  }}>
    <div>
      <div style={{ textAlign: 'center', marginBottom: '1rem' }}>Loading...</div>
      <div style={{ width: '50px', height: '50px', border: '5px solid #f3f3f3', borderTop: '5px solid #3498db', borderRadius: '50%', margin: '0 auto', animation: 'spin 1s linear infinite' }}></div>
      <style>{`
        @keyframes spin {
          0% { transform: rotate(0deg); }
          100% { transform: rotate(360deg); }
        }
      `}</style>
    </div>
  </div>
);

export default function App() {
  const { isAuthenticated } = useAuth();

  return (
    <>
      <Navbar />
      <ErrorBoundary>
        <Suspense fallback={<LoadingFallback />}>
          <div style={{ padding: '1rem' }}>
            <Routes>
              <Route path="/" element={<Navigate to="/dashboard" />} />
              <Route path="/dashboard" element={<div>Welcome to Spa Management System</div>} />
              <Route path="/customers/*" element={
                <ErrorBoundary>
                  {isAuthenticated ? <CustomerApp /> : <Navigate to="/login" />}
                </ErrorBoundary>
              } />
              <Route path="/appointments/*" element={
                <ErrorBoundary>
                  {isAuthenticated ? <AppointmentApp /> : <Navigate to="/login" />}
                </ErrorBoundary>
              } />
              {/* Add more routes for other micro frontends */}
              <Route path="/login" element={<LoginPage />} />
              <Route path="*" element={<div>Page not found</div>} />
            </Routes>
          </div>
        </Suspense>
      </ErrorBoundary>
    </>
  );
}
