#nginx/Dockerfile
FROM nginx:latest
EXPOSE 80
COPY nginx.conf /etc/nginx/nginx.conf
ENTRYPOINT ["nginx","-g","daemon off;"]
