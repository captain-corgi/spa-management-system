export interface Customer {
  id: number;
  name: string;
  email: string;
  phone: string;
  address?: string;
  membershipLevel?: 'standard' | 'premium' | 'vip';
  joinDate?: string;
  lastVisit?: string;
  notes?: string;
}
