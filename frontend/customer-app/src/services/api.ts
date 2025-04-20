import { Customer } from '../types/customer';
import { mockCustomerApi } from './mockApi';

// Environment configuration
const isDevelopment = import.meta.env.DEV;
const useMockApi = isDevelopment; // Set to true to use mock API in development

// Base API URL
const API_BASE_URL = '/api';

// Error handling helper
const handleResponse = async (response: Response) => {
  if (!response.ok) {
    const errorText = await response.text();
    throw new Error(errorText || `API error: ${response.status}`);
  }
  return response.json();
};

// Real API implementation
const realCustomerApi = {
  // Get all customers
  getCustomers: async (): Promise<Customer[]> => {
    try {
      const response = await fetch(`${API_BASE_URL}/customers`);
      return handleResponse(response);
    } catch (error) {
      console.error('Error fetching customers:', error);
      throw error;
    }
  },

  // Get customer by ID
  getCustomerById: async (id: number): Promise<Customer> => {
    try {
      const response = await fetch(`${API_BASE_URL}/customers/${id}`);
      return handleResponse(response);
    } catch (error) {
      console.error(`Error fetching customer ${id}:`, error);
      throw error;
    }
  },

  // Create new customer
  createCustomer: async (customer: Omit<Customer, 'id'>): Promise<Customer> => {
    try {
      const response = await fetch(`${API_BASE_URL}/customers`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(customer),
      });
      return handleResponse(response);
    } catch (error) {
      console.error('Error creating customer:', error);
      throw error;
    }
  },

  // Update existing customer
  updateCustomer: async (id: number, customer: Partial<Customer>): Promise<Customer> => {
    try {
      const response = await fetch(`${API_BASE_URL}/customers/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(customer),
      });
      return handleResponse(response);
    } catch (error) {
      console.error(`Error updating customer ${id}:`, error);
      throw error;
    }
  },

  // Delete customer
  deleteCustomer: async (id: number): Promise<void> => {
    try {
      const response = await fetch(`${API_BASE_URL}/customers/${id}`, {
        method: 'DELETE',
      });
      return handleResponse(response);
    } catch (error) {
      console.error(`Error deleting customer ${id}:`, error);
      throw error;
    }
  },
};

// Export the appropriate API based on configuration
export const customerApi = useMockApi ? mockCustomerApi : realCustomerApi;

// Log which API is being used
console.log(`Using ${useMockApi ? 'mock' : 'real'} customer API`);

