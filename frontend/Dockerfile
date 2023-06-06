# Use a base image with Node.js installed
FROM node:18-alpine

# Set the working directory in the container
WORKDIR /deselflopment-babl/frontend/

# Copy the package.json and yarn.lock files to the container
COPY package.json yarn.lock ./

# Install project dependencies using Yarn
RUN yarn install

# Copy the rest of the project files to the container
COPY . .

# Start your application
CMD ["yarn", "start"]
