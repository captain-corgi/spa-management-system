import React from 'react';
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import CustomerDetail from '../components/CustomerDetail';
import { Customer } from '../types/customer';
import { customerApi } from '../services/api';

// Mock the API service
jest.mock('../services/api', () => ({
  customerApi: {
    getCustomerById: jest.fn(),
  },
}));

describe('CustomerDetail', () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  it('renders placeholder when no customer is selected', () => {
    render(<CustomerDetail customerId={null} />);
    expect(screen.getByText(/Select a customer to view details/i)).toBeInTheDocument();
  });

  it('renders loading state', () => {
    // Simulate loading by not resolving the promise
    (customerApi.getCustomerById as jest.Mock).mockReturnValue(new Promise(() => {}));
    render(<CustomerDetail customerId={1} />);
    expect(screen.getByText(/Loading customer details/i)).toBeInTheDocument();
  });

  it('renders customer details', async () => {
    const customer: Customer = {
      id: 1,
      name: 'Alice Johnson',
      email: 'alice@example.com',
      phone: '123-456-7890',
      address: '123 Main St',
      membershipLevel: 'premium',
      joinDate: '2023-01-15',
      lastVisit: '2023-04-20',
      notes: 'Prefers evening appointments',
    };
    
    (customerApi.getCustomerById as jest.Mock).mockResolvedValue(customer);
    
    render(<CustomerDetail customerId={1} />);
    
    expect(await screen.findByText(/Alice Johnson/)).toBeInTheDocument();
    expect(await screen.findByText(/alice@example.com/)).toBeInTheDocument();
    expect(await screen.findByText(/123-456-7890/)).toBeInTheDocument();
    expect(await screen.findByText(/123 Main St/)).toBeInTheDocument();
    expect(await screen.findByText(/Premium/)).toBeInTheDocument();
    expect(await screen.findByText(/Prefers evening appointments/)).toBeInTheDocument();
  });

  it('shows error message when API call fails', async () => {
    (customerApi.getCustomerById as jest.Mock).mockRejectedValue(new Error('Failed to load customer'));
    
    render(<CustomerDetail customerId={1} />);
    
    expect(await screen.findByText(/Failed to load customer/)).toBeInTheDocument();
    
    // Test the retry button
    const retryButton = await screen.findByText(/Try Again/);
    
    // Mock a successful response for the retry
    (customerApi.getCustomerById as jest.Mock).mockResolvedValue({
      id: 1,
      name: 'Alice',
      email: 'alice@example.com',
      phone: '123',
    });
    
    fireEvent.click(retryButton);
    
    // Should show loading state again
    expect(screen.getByText(/Loading customer details/i)).toBeInTheDocument();
    
    // Then should show the customer details
    expect(await screen.findByText(/Alice/)).toBeInTheDocument();
  });

  it('clears customer when customerId changes to null', async () => {
    const customer: Customer = {
      id: 1,
      name: 'Alice',
      email: 'alice@example.com',
      phone: '123',
    };
    
    (customerApi.getCustomerById as jest.Mock).mockResolvedValue(customer);
    
    const { rerender } = render(<CustomerDetail customerId={1} />);
    
    expect(await screen.findByText(/Alice/)).toBeInTheDocument();
    
    rerender(<CustomerDetail customerId={null} />);
    
    expect(screen.getByText(/Select a customer to view details/i)).toBeInTheDocument();
  });
});
