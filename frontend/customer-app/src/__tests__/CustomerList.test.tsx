import React from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import CustomerList, { Customer } from '../components/CustomerList';

describe('CustomerList', () => {
  it('renders loading state', () => {
    // Simulate loading by not resolving fetch
    global.fetch = jest.fn(() => new Promise(() => {})) as any;
    render(<CustomerList onSelect={() => {}} />);
    expect(screen.getByText(/loading customers/i)).toBeInTheDocument();
  });

  it('renders customer list', async () => {
    const customers: Customer[] = [
      { id: 1, name: 'Alice', phone: '123' },
      { id: 2, name: 'Bob', phone: '456' },
    ];
    global.fetch = jest.fn(() => Promise.resolve({ json: () => Promise.resolve(customers) })) as any;
    render(<CustomerList onSelect={() => {}} />);
    expect(await screen.findByText(/Alice/)).toBeInTheDocument();
    expect(await screen.findByText(/Bob/)).toBeInTheDocument();
  });

  it('calls onSelect when customer is clicked', async () => {
    const customers: Customer[] = [
      { id: 1, name: 'Alice', phone: '123' },
    ];
    global.fetch = jest.fn(() => Promise.resolve({ json: () => Promise.resolve(customers) })) as any;
    const onSelect = jest.fn();
    render(<CustomerList onSelect={onSelect} />);
    const btn = await screen.findByText(/Alice/);
    fireEvent.click(btn);
    expect(onSelect).toHaveBeenCalledWith(1);
  });
});
