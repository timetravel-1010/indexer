# Stage 1: Build the app
FROM node:20-alpine AS builder

WORKDIR /app
COPY package*.json ./
RUN npm ci            # Clean install (production-only)
COPY . .
RUN npm run build     # Generate /dist folder

# Stage 2: Serve the app
FROM node:20-alpine

WORKDIR /app
COPY --from=builder /app/dist ./dist

# Install lightweight HTTP server
RUN npm install -g serve

# Expose port 3000 (default for serve)
EXPOSE 3000

# Start the server
CMD ["serve", "-s", "dist", "-l", "3000"]
