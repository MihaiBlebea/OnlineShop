FROM node:latest

WORKDIR /app

COPY . .

RUN npm install

RUN npm run build

ENV NUXT_HOST=0.0.0.0
ENV NUXT_PORT=5000

EXPOSE 5000

CMD [ "npm", "run", "start" ]