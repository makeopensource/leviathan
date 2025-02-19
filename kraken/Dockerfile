FROM node:23-alpine

WORKDIR /kraken/

COPY package.json .

RUN npm i

COPY . .

CMD ["npm", "run", "start"]
