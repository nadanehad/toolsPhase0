FROM node:20-alpine

WORKDIR /app

# Copy the package.json and package-lock.json from the src directory
COPY ./src/package.json ./src/package-lock.json ./

# Install dependencies
RUN npm install

# Copy both the src and public directories into the container
# Make sure both `src` and `public` are copied to their expected locations in /app
COPY ./src /app/src
COPY ./src/public /app/public

# Expose the port the app runs on
EXPOSE 3000

CMD ["npm", "start"]
