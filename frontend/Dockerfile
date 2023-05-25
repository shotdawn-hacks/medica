from node:17-alpine

RUN apk add xdg-utils

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm install --silent

COPY . .

CMD ["npm", "run", "start"]

EXPOSE 8080