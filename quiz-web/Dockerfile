# pull official base image
FROM node:13.12.0-alpine

# set working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install app dependencies
COPY package.json ./
COPY package-lock.json ./
RUN npm install --silent
RUN npm install react-scripts@3.4.3 -g --silent
RUN npm install prop-types
RUN npm install react-transition-group@1.x

# add app
COPY . ./

EXPOSE 3000
# start app
ENTRYPOINT ["npm", "start"]
