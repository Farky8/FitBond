
FROM node

RUN mkdir -p /home/frontEnd

WORKDIR /home/frontEnd

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 8081

CMD ["npm", "start"]
