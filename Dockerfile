FROM library/postgres
COPY ./init.sql /docker-entrypoint-initdb.d/

# Build AdonisJS
FROM node:lts as builder
# Set directory for all files
WORKDIR /home/node
# Copy over package.json files
COPY package*.json ./
# Install all packages
RUN yarn
# Copy over source code
COPY . .
# Build AdonisJS for production
RUN yarn build


# Build final runtime container
FROM node:lts-alpine
# Set environment variables
ENV NODE_ENV=development
# Disable .env file loading
# ENV ENV_SILENT=true
# Set home dir
WORKDIR /home/node
# Copy over built files
COPY --from=builder /home/node/build .
# Use PM2 as a process manager for our Node server
# RUN npm i -g pm2
# Install only required packages
RUN rm -rf node_modules && yarn install --frozen-lockfile --production
# Expose port to outside world
EXPOSE 3333
# Start server up
CMD [ "yarn", "db:server" ]
# Start pm2 server
# CMD yarn pm2:start
