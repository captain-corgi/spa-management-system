import React, { Suspense } from 'react';
import { Routes, Route, Navigate } from 'react-router-dom';
import { useAuth } from './shared/auth/AuthContext';
import Navbar from './shared/components/Navbar';

// const CustomerApp = React.lazy(() => import('customerApp/App'));
// const AppointmentApp = React.lazy(() => import('appointmentApp/App'));
// Add other micro frontends as needed

export default function App() {
  const { isAuthenticated } = useAuth();
  return (
    <>
      <Navbar />
      <Suspense fallback={<div>Loading...</div>}>
        <Routes>
          <Route path="/" element={<Navigate to="/dashboard" />} />
          <Route path="/dashboard" element={<div>Welcome to Spa Management System</div>} />
          {/* <Route path="/customers/*" element={isAuthenticated ? <CustomerApp /> : <Navigate to="/login" />} />
          <Route path="/appointments/*" element={isAuthenticated ? <AppointmentApp /> : <Navigate to="/login" />} />
          Add more routes for other micro frontends */}
          <Route path="/login" element={<div>Login Page</div>} />
        </Routes>
      </Suspense>
    </>
  );
}
