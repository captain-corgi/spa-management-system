import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter } from 'react-router-dom';
import App from './App';
import { AuthProvider } from './shared/auth/AuthContext';
import { GlobalStateProvider } from './shared/state/GlobalState';
import './index.css';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <AuthProvider>
      <GlobalStateProvider>
        <BrowserRouter>
          <App />
        </BrowserRouter>
      </GlobalStateProvider>
    </AuthProvider>
  </React.StrictMode>
);
