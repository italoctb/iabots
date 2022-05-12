FROM node:14-alpine

WORKDIR /src

COPY /client/package.json /src

RUN npm i

ADD /client/ /src

RUN npm run build

CMD npm start