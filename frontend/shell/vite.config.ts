import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import federation from '@originjs/vite-plugin-federation';

export default defineConfig({
  plugins: [
    react(),
    federation({
      name: 'shell',
      remotes: {
        customerApp: 'http://localhost:3001/assets/remoteEntry.js',
        appointmentApp: 'http://localhost:3002/assets/remoteEntry.js',
        staffApp: 'http://localhost:3003/assets/remoteEntry.js',
        branchApp: 'http://localhost:3004/assets/remoteEntry.js',
        marketingApp: 'http://localhost:3005/assets/remoteEntry.js',
        financeApp: 'http://localhost:3006/assets/remoteEntry.js',
        analyticsApp: 'http://localhost:3007/assets/remoteEntry.js',
      },
      shared: ['react', 'react-dom', 'react-router-dom'],
    }),
  ],
  build: {
    target: 'esnext',
    minify: false,
    cssCodeSplit: false,
  },
  server: {
    port: 3000,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ''),
      },
    },
  },
});
