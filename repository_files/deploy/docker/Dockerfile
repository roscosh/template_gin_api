FROM node:lts as frontend-build
WORKDIR /app
COPY package.json /app/package.json
RUN yarn install
COPY . /app
RUN GENERATE_SOURCEMAP=false yarn build


FROM nginx:latest as frontend-nginx
WORKDIR /app
COPY .deploy/nginx/front.conf /etc/nginx/conf.d/default.conf
COPY --from=frontend-build /app/build /app/source
