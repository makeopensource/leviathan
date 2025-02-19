FROM node:23-alpine

WORKDIR /kraken/

# dont copy lock file, since it messes with the gitpkg when isntalling the node sdk for levithan
COPY package.json .

RUN npm i

COPY . .

CMD ["npm", "run", "start"]
