FROM nginx:latest
COPY ./nginx.conf /etc/nginx/nginx.conf
COPY ./cert /etc/nginx/cert
ADD ./error.log /logs/error.log
ADD ./access.log logs/access.log