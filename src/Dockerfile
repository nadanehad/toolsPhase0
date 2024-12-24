FROM node:20-alpine

WORKDIR /app

# Debug step: List files in the root of the build context
RUN ls -al /

# Debug step: List files in the src directory
RUN ls -al ./src

# Copy package.json and package-lock.json from the src directory
COPY ./src/package.json ./src/package-lock.json ./

# Install dependencies
RUN npm install

# Copy the entire src and public directories
COPY ./src /app/src
COPY ./src/public /app/public

# Expose the port the app runs on
EXPOSE 3000

CMD ["npm", "start"]
