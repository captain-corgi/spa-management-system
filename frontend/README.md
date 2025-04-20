# Spa Management System Frontend

This directory contains all frontend micro frontend applications and the shell. Each app is a standalone React+Vite+TypeScript project that follows a micro frontend architecture.

## Architecture

The frontend is built using a micro frontend architecture with the following components:

- **Shell**: The container application that hosts all micro frontends
- **Micro Frontends**: Domain-specific applications that are loaded dynamically by the shell

Each micro frontend is independently deployable and follows the same technology stack:

- React with TypeScript
- Vite for build tooling
- Module Federation for micro frontend integration
- Vitest for testing

## Directory Structure

```text
frontend/
├── shell/                     # Container application
├── customer-app/              # Customer management micro frontend
├── appointment-app/           # Appointment management micro frontend
├── staff-app/                 # Staff management micro frontend
├── branch-app/                # Branch management micro frontend
├── marketing-app/             # Marketing management micro frontend
├── finance-app/               # Finance management micro frontend
├── analytics-app/             # Analytics micro frontend
├── Makefile                   # Common operations for all apps
└── .augment-guidelines        # Architecture guidelines
```

## Using the Makefile

The Makefile provides common operations for all micro frontend applications. Here are some examples:

### Installing Dependencies

```bash
# Install dependencies for all apps
make install

# Install dependencies for a specific app
make install-app APP=customer-app
```

### Development

```bash
# Start development servers for all apps
make dev

# Start development server for a specific app
make dev-app APP=customer-app
```

### Building

```bash
# Build all apps
make build

# Build a specific app
make build-app APP=customer-app
```

### Testing

```bash
# Run tests for all apps
make test

# Run tests for a specific app
make test-app APP=customer-app
```

### Linting and Formatting

```bash
# Run linter on all apps
make lint

# Format code in all apps
make format
```

### Creating a New Micro Frontend

```bash
# Create a new micro frontend app
make create-app APP=new-feature-app
```

### Getting Help

```bash
# Show all available commands
make help
```

## Development Workflow

1. Start the shell and all micro frontends:

   ```bash
   make dev
   ```

2. Access the application at [http://localhost:3000](http://localhost:3000)

3. Make changes to any micro frontend and see the changes reflected in the shell

## Integration with Backend

Each micro frontend communicates with the backend API through a dedicated API service layer. The API requests are proxied through the development server to avoid CORS issues.
