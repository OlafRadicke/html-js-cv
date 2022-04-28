FROM nginx:1.21-alpine
COPY css        /usr/share/nginx/html/css
COPY fonts      /usr/share/nginx/html/fonts
COPY js         /usr/share/nginx/html/js
COPY pics       /usr/share/nginx/html/pics
COPY docs       /usr/share/nginx/html/docs
COPY index.html /usr/share/nginx/html/index.html
