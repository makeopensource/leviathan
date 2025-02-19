FROM node:23-alpine

COPY package.json .
COPY package-lock.json .

RUN npm i

COPY . .

CMD ["npm", "run", "start"]
