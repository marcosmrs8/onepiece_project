FROM node:alpine3.16
WORKDIR /app
COPY . .
RUN npm install
RUN npm install axios
RUN npm install @mui/material @emotion/react @emotion/styled

EXPOSE 3000
CMD ["npm", "start"]