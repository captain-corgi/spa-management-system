# Makefile for Spa Management System Frontend
# Provides common operations for all frontend micro applications

# Variables
NODE_VERSION := 18.x
APPS := shell customer-app appointment-app staff-app branch-app marketing-app finance-app analytics-app
DEFAULT_PORT := 3000

# Colors for terminal output
COLOR_RESET := $(shell tput sgr0)
COLOR_GREEN := $(shell tput setaf 2)
COLOR_YELLOW := $(shell tput setaf 3)
COLOR_BLUE := $(shell tput setaf 4)

.PHONY: all help install build dev test lint clean format check-deps

# Default target
all: install build

# Help command
help:
	@echo "$(COLOR_GREEN)Spa Management System Frontend Makefile$(COLOR_RESET)"
	@echo "$(COLOR_YELLOW)Usage:$(COLOR_RESET)"
	@echo "  make [target]"
	@echo ""
	@echo "$(COLOR_YELLOW)Targets:$(COLOR_RESET)"
	@echo "  $(COLOR_GREEN)help$(COLOR_RESET)           - Show this help message"
	@echo "  $(COLOR_GREEN)install$(COLOR_RESET)        - Install dependencies for all apps"
	@echo "  $(COLOR_GREEN)install-app$(COLOR_RESET)    - Install dependencies for a specific app (e.g., make install-app APP=customer-app)"
	@echo "  $(COLOR_GREEN)build$(COLOR_RESET)          - Build all apps"
	@echo "  $(COLOR_GREEN)build-app$(COLOR_RESET)      - Build a specific app (e.g., make build-app APP=customer-app)"
	@echo "  $(COLOR_GREEN)dev$(COLOR_RESET)            - Start development server for shell and all apps"
	@echo "  $(COLOR_GREEN)dev-app$(COLOR_RESET)        - Start development server for a specific app (e.g., make dev-app APP=customer-app)"
	@echo "  $(COLOR_GREEN)test$(COLOR_RESET)           - Run tests for all apps"
	@echo "  $(COLOR_GREEN)test-app$(COLOR_RESET)       - Run tests for a specific app (e.g., make test-app APP=customer-app)"
	@echo "  $(COLOR_GREEN)lint$(COLOR_RESET)           - Run linter on all apps"
	@echo "  $(COLOR_GREEN)lint-app$(COLOR_RESET)       - Run linter on a specific app (e.g., make lint-app APP=customer-app)"
	@echo "  $(COLOR_GREEN)format$(COLOR_RESET)         - Format code in all apps"
	@echo "  $(COLOR_GREEN)format-app$(COLOR_RESET)     - Format code in a specific app (e.g., make format-app APP=customer-app)"
	@echo "  $(COLOR_GREEN)clean$(COLOR_RESET)          - Remove build artifacts from all apps"
	@echo "  $(COLOR_GREEN)clean-app$(COLOR_RESET)      - Remove build artifacts from a specific app (e.g., make clean-app APP=customer-app)"
	@echo "  $(COLOR_GREEN)check-deps$(COLOR_RESET)     - Check for outdated dependencies in all apps"
	@echo ""
	@echo "$(COLOR_YELLOW)Environment Variables:$(COLOR_RESET)"
	@echo "  APP            - Specific app to operate on"
	@echo "  NODE_ENV       - Node environment (development, production, test)"

# Check if APP variable is set
check-app:
	@if [ -z "$(APP)" ]; then \
		echo "$(COLOR_YELLOW)Please specify an app with APP=<app-name>$(COLOR_RESET)"; \
		exit 1; \
	fi
	@if [ ! -d "$(APP)" ]; then \
		echo "$(COLOR_YELLOW)App directory '$(APP)' does not exist$(COLOR_RESET)"; \
		exit 1; \
	fi

# Install dependencies for all apps
install:
	@echo "$(COLOR_BLUE)Installing dependencies for all apps...$(COLOR_RESET)"
	@for app in $(APPS); do \
		if [ -d "$$app" ]; then \
			echo "$(COLOR_BLUE)Installing dependencies for $$app...$(COLOR_RESET)"; \
			(cd $$app && npm install) || exit 1; \
		fi \
	done
	@echo "$(COLOR_GREEN)All dependencies installed successfully!$(COLOR_RESET)"

# Install dependencies for a specific app
install-app: check-app
	@echo "$(COLOR_BLUE)Installing dependencies for $(APP)...$(COLOR_RESET)"
	@cd $(APP) && npm install
	@echo "$(COLOR_GREEN)Dependencies for $(APP) installed successfully!$(COLOR_RESET)"

# Build all apps
build:
	@echo "$(COLOR_BLUE)Building all apps...$(COLOR_RESET)"
	@for app in $(APPS); do \
		if [ -d "$$app" ]; then \
			echo "$(COLOR_BLUE)Building $$app...$(COLOR_RESET)"; \
			(cd $$app && npm run build) || exit 1; \
		fi \
	done
	@echo "$(COLOR_GREEN)All apps built successfully!$(COLOR_RESET)"

# Build a specific app
build-app: check-app
	@echo "$(COLOR_BLUE)Building $(APP)...$(COLOR_RESET)"
	@cd $(APP) && npm run build
	@echo "$(COLOR_GREEN)$(APP) built successfully!$(COLOR_RESET)"

# Start development servers for all apps
dev:
	@echo "$(COLOR_BLUE)Starting development servers for all apps...$(COLOR_RESET)"
	@echo "$(COLOR_YELLOW)Note: This will start all apps in separate terminal windows.$(COLOR_RESET)"
	@for app in $(APPS); do \
		if [ -d "$$app" ]; then \
			echo "$(COLOR_BLUE)Starting $$app...$(COLOR_RESET)"; \
			osascript -e 'tell application "Terminal" to do script "cd $(PWD)/$$app && npm run dev"' || exit 1; \
			sleep 2; \
		fi \
	done
	@echo "$(COLOR_GREEN)All development servers started!$(COLOR_RESET)"

# Start development server for a specific app
dev-app: check-app
	@echo "$(COLOR_BLUE)Starting development server for $(APP)...$(COLOR_RESET)"
	@cd $(APP) && npm run dev

# Run tests for all apps
test:
	@echo "$(COLOR_BLUE)Running tests for all apps...$(COLOR_RESET)"
	@for app in $(APPS); do \
		if [ -d "$$app" ]; then \
			echo "$(COLOR_BLUE)Testing $$app...$(COLOR_RESET)"; \
			(cd $$app && npm test) || exit 1; \
		fi \
	done
	@echo "$(COLOR_GREEN)All tests passed!$(COLOR_RESET)"

# Run tests for a specific app
test-app: check-app
	@echo "$(COLOR_BLUE)Running tests for $(APP)...$(COLOR_RESET)"
	@cd $(APP) && npm test
	@echo "$(COLOR_GREEN)Tests for $(APP) completed!$(COLOR_RESET)"

# Run linter on all apps
lint:
	@echo "$(COLOR_BLUE)Running linter on all apps...$(COLOR_RESET)"
	@for app in $(APPS); do \
		if [ -d "$$app" ]; then \
			echo "$(COLOR_BLUE)Linting $$app...$(COLOR_RESET)"; \
			(cd $$app && npm run lint) || exit 1; \
		fi \
	done
	@echo "$(COLOR_GREEN)Linting completed!$(COLOR_RESET)"

# Run linter on a specific app
lint-app: check-app
	@echo "$(COLOR_BLUE)Running linter on $(APP)...$(COLOR_RESET)"
	@cd $(APP) && npm run lint
	@echo "$(COLOR_GREEN)Linting for $(APP) completed!$(COLOR_RESET)"

# Format code in all apps
format:
	@echo "$(COLOR_BLUE)Formatting code in all apps...$(COLOR_RESET)"
	@for app in $(APPS); do \
		if [ -d "$$app" ]; then \
			echo "$(COLOR_BLUE)Formatting $$app...$(COLOR_RESET)"; \
			(cd $$app && npm run format) || echo "$(COLOR_YELLOW)Format script not found in $$app$(COLOR_RESET)"; \
		fi \
	done
	@echo "$(COLOR_GREEN)Formatting completed!$(COLOR_RESET)"

# Format code in a specific app
format-app: check-app
	@echo "$(COLOR_BLUE)Formatting code in $(APP)...$(COLOR_RESET)"
	@cd $(APP) && npm run format || echo "$(COLOR_YELLOW)Format script not found in $(APP)$(COLOR_RESET)"
	@echo "$(COLOR_GREEN)Formatting for $(APP) completed!$(COLOR_RESET)"

# Clean build artifacts from all apps
clean:
	@echo "$(COLOR_BLUE)Cleaning build artifacts from all apps...$(COLOR_RESET)"
	@for app in $(APPS); do \
		if [ -d "$$app" ]; then \
			echo "$(COLOR_BLUE)Cleaning $$app...$(COLOR_RESET)"; \
			(cd $$app && rm -rf dist node_modules/.vite) || exit 1; \
		fi \
	done
	@echo "$(COLOR_GREEN)All apps cleaned!$(COLOR_RESET)"

# Clean build artifacts from a specific app
clean-app: check-app
	@echo "$(COLOR_BLUE)Cleaning build artifacts from $(APP)...$(COLOR_RESET)"
	@cd $(APP) && rm -rf dist node_modules/.vite
	@echo "$(COLOR_GREEN)$(APP) cleaned!$(COLOR_RESET)"

# Check for outdated dependencies in all apps
check-deps:
	@echo "$(COLOR_BLUE)Checking for outdated dependencies in all apps...$(COLOR_RESET)"
	@for app in $(APPS); do \
		if [ -d "$$app" ]; then \
			echo "$(COLOR_BLUE)Checking dependencies for $$app...$(COLOR_RESET)"; \
			(cd $$app && npm outdated) || exit 1; \
		fi \
	done
	@echo "$(COLOR_GREEN)Dependency check completed!$(COLOR_RESET)"

# Create a new micro frontend app
create-app:
	@if [ -z "$(APP)" ]; then \
		echo "$(COLOR_YELLOW)Please specify an app name with APP=<app-name>$(COLOR_RESET)"; \
		exit 1; \
	fi
	@if [ -d "$(APP)" ]; then \
		echo "$(COLOR_YELLOW)App directory '$(APP)' already exists$(COLOR_RESET)"; \
		exit 1; \
	fi
	@echo "$(COLOR_BLUE)Creating new micro frontend app: $(APP)...$(COLOR_RESET)"
	@mkdir -p $(APP)/src/components
	@mkdir -p $(APP)/src/services
	@mkdir -p $(APP)/src/types
	@mkdir -p $(APP)/src/__tests__
	
	@# Create package.json
	@echo '{\n  "name": "$(APP)",\n  "version": "1.0.0",\n  "private": true,\n  "type": "module",\n  "scripts": {\n    "dev": "vite",\n    "build": "vite build",\n    "preview": "vite preview",\n    "test": "vitest",\n    "lint": "eslint . --ext ts,tsx --report-unused-disable-directives --max-warnings 0",\n    "format": "prettier --write \\"src/**/*.{ts,tsx}\\"\n  },\n  "dependencies": {\n    "react": "^18.2.0",\n    "react-dom": "^18.2.0",\n    "react-router-dom": "^6.22.3"\n  },\n  "devDependencies": {\n    "@types/react": "^18.2.38",\n    "@types/react-dom": "^18.2.17",\n    "@vitejs/plugin-react": "^4.2.1",\n    "typescript": "^5.3.3",\n    "vite": "^5.2.0",\n    "@originjs/vite-plugin-federation": "^1.3.3",\n    "vitest": "^1.4.0",\n    "eslint": "^8.56.0",\n    "prettier": "^3.2.5"\n  }\n}' > $(APP)/package.json
	
	@# Create vite.config.ts
	@echo 'import { defineConfig } from "vite";\nimport react from "@vitejs/plugin-react";\nimport federation from "@originjs/vite-plugin-federation";\n\nexport default defineConfig({\n  plugins: [\n    react(),\n    federation({\n      name: "$(subst -,,$(APP))",\n      filename: "remoteEntry.js",\n      exposes: {\n        "./App": "./src/App.tsx",\n      },\n      shared: ["react", "react-dom", "react-router-dom"],\n    }),\n  ],\n  build: {\n    target: "esnext",\n    minify: false,\n    cssCodeSplit: false,\n    outDir: "dist",\n  },\n  server: {\n    port: $(shell echo $$(($(DEFAULT_PORT) + $(words $(filter-out $(APP),$(APPS)))))),\n    proxy: {\n      "/api": {\n        target: "http://localhost:8080",\n        changeOrigin: true,\n        rewrite: (path) => path.replace(/^\\/api/, ""),\n      },\n    },\n  },\n});' > $(APP)/vite.config.ts
	
	@# Create tsconfig.json
	@echo '{\n  "compilerOptions": {\n    "target": "ES2020",\n    "useDefineForClassFields": true,\n    "lib": ["ES2020", "DOM", "DOM.Iterable"],\n    "module": "ESNext",\n    "skipLibCheck": true,\n    "moduleResolution": "bundler",\n    "allowImportingTsExtensions": true,\n    "resolveJsonModule": true,\n    "isolatedModules": true,\n    "noEmit": true,\n    "jsx": "react-jsx",\n    "strict": true,\n    "noUnusedLocals": true,\n    "noUnusedParameters": true,\n    "noFallthroughCasesInSwitch": true\n  },\n  "include": ["src"],\n  "references": [{ "path": "./tsconfig.node.json" }]\n}' > $(APP)/tsconfig.json
	
	@# Create tsconfig.node.json
	@echo '{\n  "compilerOptions": {\n    "composite": true,\n    "skipLibCheck": true,\n    "module": "ESNext",\n    "moduleResolution": "bundler",\n    "allowSyntheticDefaultImports": true\n  },\n  "include": ["vite.config.ts"]\n}' > $(APP)/tsconfig.node.json
	
	@# Create basic App.tsx
	@echo 'import React from "react";\nimport "./index.css";\n\nexport default function App() {\n  return (\n    <div style={{ padding: "1rem" }}>\n      <h2>$(APP)</h2>\n      <p>This is the $(APP) micro frontend.</p>\n    </div>\n  );\n}' > $(APP)/src/App.tsx
	
	@# Create index.css
	@echo 'body {\n  font-family: system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen, Ubuntu, Cantarell, "Open Sans", "Helvetica Neue", sans-serif;\n  margin: 0;\n  padding: 0;\n  background-color: #f5f7fa;\n  color: #333;\n  line-height: 1.5;\n}\n\n* {\n  box-sizing: border-box;\n}\n\nh1, h2, h3, h4, h5, h6 {\n  margin-top: 0;\n  color: #2c3e50;\n}' > $(APP)/src/index.css
	
	@# Create README.md
	@echo '# $(APP)\n\nThis micro frontend is part of the Spa Management System.\n\n## Features\n\n- Feature 1\n- Feature 2\n\n## Getting Started\n\n```bash\n# Install dependencies\nnpm install\n\n# Start development server\nnpm run dev\n\n# Build for production\nnpm run build\n\n# Run tests\nnpm test\n```\n\n## Integration with Shell\n\nThis micro frontend is integrated with the shell application using Module Federation.' > $(APP)/README.md
	
	@echo "$(COLOR_GREEN)New micro frontend app '$(APP)' created successfully!$(COLOR_RESET)"
	@echo "$(COLOR_YELLOW)Next steps:$(COLOR_RESET)"
	@echo "1. cd $(APP)"
	@echo "2. npm install"
	@echo "3. npm run dev"
	@echo "4. Update the shell app to include this micro frontend"
