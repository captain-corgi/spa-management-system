import React, { Component, ErrorInfo, ReactNode } from 'react';

interface Props {
  children: ReactNode;
  fallback?: ReactNode;
}

interface State {
  hasError: boolean;
  error: Error | null;
}

class ErrorBoundary extends Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = {
      hasError: false,
      error: null,
    };
  }

  static getDerivedStateFromError(error: Error): State {
    return {
      hasError: true,
      error,
    };
  }

  componentDidCatch(error: Error, errorInfo: ErrorInfo): void {
    // You can log the error to an error reporting service here
    console.error('Error caught by ErrorBoundary:', error, errorInfo);
  }

  render(): ReactNode {
    if (this.state.hasError) {
      // You can render any custom fallback UI
      if (this.props.fallback) {
        return this.props.fallback;
      }
      
      return (
        <div style={{
          padding: '2rem',
          margin: '1rem',
          backgroundColor: '#fff3f3',
          border: '1px solid #ffcaca',
          borderRadius: '4px',
          boxShadow: '0 2px 4px rgba(0,0,0,0.1)',
        }}>
          <h2 style={{ color: '#e74c3c', marginTop: 0 }}>Something went wrong</h2>
          <p>We're sorry, but there was an error loading this component.</p>
          <details style={{ 
            marginTop: '1rem',
            padding: '0.5rem',
            backgroundColor: '#f8f8f8',
            border: '1px solid #ddd',
            borderRadius: '4px',
          }}>
            <summary>Error details</summary>
            <pre style={{ 
              overflow: 'auto',
              padding: '0.5rem',
              backgroundColor: '#f1f1f1',
              border: '1px solid #ddd',
              borderRadius: '4px',
            }}>
              {this.state.error?.toString()}
            </pre>
          </details>
          <button
            onClick={() => this.setState({ hasError: false, error: null })}
            style={{
              marginTop: '1rem',
              padding: '0.5rem 1rem',
              backgroundColor: '#3498db',
              color: 'white',
              border: 'none',
              borderRadius: '4px',
              cursor: 'pointer',
            }}
          >
            Try again
          </button>
        </div>
      );
    }

    return this.props.children;
  }
}

export default ErrorBoundary;
