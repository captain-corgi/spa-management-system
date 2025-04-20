import { Customer } from '../types/customer';

// Mock data for development
const mockCustomers: Customer[] = [
  {
    id: 1,
    name: 'Alice Johnson',
    email: 'alice@example.com',
    phone: '123-456-7890',
    address: '123 Main St, Anytown, USA',
    membershipLevel: 'premium',
    joinDate: '2023-01-15',
    lastVisit: '2023-04-20',
    notes: 'Prefers evening appointments'
  },
  {
    id: 2,
    name: 'Bob Smith',
    email: 'bob@example.com',
    phone: '234-567-8901',
    address: '456 Oak Ave, Somewhere, USA',
    membershipLevel: 'standard',
    joinDate: '2023-02-20',
    lastVisit: '2023-04-15',
    notes: 'Allergic to lavender'
  },
  {
    id: 3,
    name: 'Carol Davis',
    email: 'carol@example.com',
    phone: '345-678-9012',
    address: '789 Pine St, Nowhere, USA',
    membershipLevel: 'vip',
    joinDate: '2022-11-05',
    lastVisit: '2023-04-18',
    notes: 'Prefers female staff'
  },
  {
    id: 4,
    name: 'David Wilson',
    email: 'david@example.com',
    phone: '456-789-0123',
    address: '101 Elm St, Elsewhere, USA',
    membershipLevel: 'standard',
    joinDate: '2023-03-10',
    lastVisit: '2023-04-10',
    notes: ''
  },
  {
    id: 5,
    name: 'Emma Brown',
    email: 'emma@example.com',
    phone: '567-890-1234',
    address: '202 Maple Ave, Anywhere, USA',
    membershipLevel: 'premium',
    joinDate: '2022-12-15',
    lastVisit: '2023-04-05',
    notes: 'Interested in package deals'
  }
];

// Simulate API delay
const delay = (ms: number) => new Promise(resolve => setTimeout(resolve, ms));

// Mock API functions
export const mockCustomerApi = {
  getCustomers: async (): Promise<Customer[]> => {
    await delay(500); // Simulate network delay
    return [...mockCustomers];
  },

  getCustomerById: async (id: number): Promise<Customer> => {
    await delay(300); // Simulate network delay
    const customer = mockCustomers.find(c => c.id === id);
    if (!customer) {
      throw new Error(`Customer with ID ${id} not found`);
    }
    return { ...customer };
  },

  createCustomer: async (customer: Omit<Customer, 'id'>): Promise<Customer> => {
    await delay(700); // Simulate network delay
    const newCustomer = {
      ...customer,
      id: Math.max(...mockCustomers.map(c => c.id)) + 1
    };
    mockCustomers.push(newCustomer as Customer);
    return { ...newCustomer as Customer };
  },

  updateCustomer: async (id: number, customer: Partial<Customer>): Promise<Customer> => {
    await delay(500); // Simulate network delay
    const index = mockCustomers.findIndex(c => c.id === id);
    if (index === -1) {
      throw new Error(`Customer with ID ${id} not found`);
    }
    mockCustomers[index] = { ...mockCustomers[index], ...customer };
    return { ...mockCustomers[index] };
  },

  deleteCustomer: async (id: number): Promise<void> => {
    await delay(400); // Simulate network delay
    const index = mockCustomers.findIndex(c => c.id === id);
    if (index === -1) {
      throw new Error(`Customer with ID ${id} not found`);
    }
    mockCustomers.splice(index, 1);
  }
};
