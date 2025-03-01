FROM alpine:latest

# Installing nginx
RUN apk update
RUN apk add nginx

# Copying nginx configuration
COPY config/nginx.conf /etc/nginx/

# Running nginx not in background
CMD ["nginx","-g","daemon off;"]
