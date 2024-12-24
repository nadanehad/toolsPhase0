# Use the base image
FROM node:20-alpine

# Set the working directory
WORKDIR /app

# Copy the package.json and package-lock.json from the src folder
COPY ./src/package.json ./src/package-lock.json ./

# Install dependencies
RUN npm install

# Copy the entire src directory into the container
COPY ./src /app/src

# Expose the application port
EXPOSE 3000

# Start the application
CMD ["npm", "start"]
