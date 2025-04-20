# Customer App

This micro frontend is responsible for customer management in the Spa Management System.

## Features

- View list of customers
- View customer details
- Error handling with Error Boundaries
- API integration with proper error handling
- Responsive design

## Architecture

The Customer App follows a clean architecture with the following components:

- **Components**: UI components for displaying customer data
- **Services**: API services for fetching customer data
- **Types**: TypeScript interfaces for type safety
- **Tests**: Unit tests for components

## API Integration

The app communicates with the backend API through a dedicated API service layer that handles:

- Data fetching
- Error handling
- Type safety

## Error Handling

The app uses Error Boundaries to catch and display errors in a user-friendly way. Each component also handles its own loading and error states.

## Testing

The app includes unit tests for all components using:

- React Testing Library
- Jest
- Mock API services

## Getting Started

1. Install dependencies:

   ```bash
   npm install
   ```

2. Run the development server:

   ```bash
   npm run dev
   ```

3. Run tests:

   ```bash
   npm test
   ```

## Building for Production

```bash
npm run build
```

## Integration with Shell

This micro frontend is integrated with the shell application using Module Federation. The app exposes its main App component which is then consumed by the shell.
