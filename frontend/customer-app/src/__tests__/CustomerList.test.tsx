import React from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import CustomerList from '../components/CustomerList';
import { Customer } from '../types/customer';
import { customerApi } from '../services/api';

// Mock the API service
jest.mock('../services/api', () => ({
  customerApi: {
    getCustomers: jest.fn(),
  },
}));

describe('CustomerList', () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  it('renders loading state', () => {
    // Simulate loading by not resolving the promise
    (customerApi.getCustomers as jest.Mock).mockReturnValue(new Promise(() => {}));
    render(<CustomerList onSelect={() => {}} />);
    expect(screen.getByText(/loading customers/i)).toBeInTheDocument();
  });

  it('renders customer list', async () => {
    const customers: Customer[] = [
      { id: 1, name: 'Alice', email: 'alice@example.com', phone: '123' },
      { id: 2, name: 'Bob', email: 'bob@example.com', phone: '456' },
    ];
    (customerApi.getCustomers as jest.Mock).mockResolvedValue(customers);

    render(<CustomerList onSelect={() => {}} />);

    expect(await screen.findByText(/Alice/)).toBeInTheDocument();
    expect(await screen.findByText(/Bob/)).toBeInTheDocument();
    expect(await screen.findByText(/alice@example.com/)).toBeInTheDocument();
  });

  it('calls onSelect when customer is clicked', async () => {
    const customers: Customer[] = [
      { id: 1, name: 'Alice', email: 'alice@example.com', phone: '123' },
    ];
    (customerApi.getCustomers as jest.Mock).mockResolvedValue(customers);

    const onSelect = jest.fn();
    render(<CustomerList onSelect={onSelect} />);

    const btn = await screen.findByText(/Alice/);
    fireEvent.click(btn);
    expect(onSelect).toHaveBeenCalledWith(1);
  });

  it('shows error message when API call fails', async () => {
    (customerApi.getCustomers as jest.Mock).mockRejectedValue(new Error('API Error'));

    render(<CustomerList onSelect={() => {}} />);

    expect(await screen.findByText(/API Error/)).toBeInTheDocument();
  });

  it('shows empty state when no customers are returned', async () => {
    (customerApi.getCustomers as jest.Mock).mockResolvedValue([]);

    render(<CustomerList onSelect={() => {}} />);

    expect(await screen.findByText(/No customers found/)).toBeInTheDocument();
  });
});
