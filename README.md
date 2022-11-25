# golang博客系统 vue+gin 联动
gitee,github两个仓库同步
## docker-compose快速启动
### blog-compose.yml
```
version: '3.7'

networks:
  monitor:
    ipam:
      config:
      - subnet: 172.62.0.0/16     

services:
    nginx: 
      image: nginx:1.21.1
      container_name: nginx
      restart: always
      ports:
        - 80:80
      networks:
        monitor:
          ipv4_address: 172.62.0.7
      volumes:
      - ./nginx/html:/usr/share/nginx/html
      - ./nginx/logs:/var/log/nginx
      #- ./nginx/conf.d:/etc/nginx/conf.d
      - ./nginx/conf/nginx.conf:/etc/nginx/nginx.conf

    mysql:
      image: mysql:8.0.21
      container_name: mysql
      command:
      # MySQL8的密码验证方式默认是 caching_sha2_password，但是很多的连接工具还不支持该方式
      # 就需要手动设置下mysql的密码认证方式为以前的 mysql_native_password 方式
        --default-authentication-plugin=mysql_native_password
        --character-set-server=utf8mb4
        --collation-server=utf8mb4_general_ci
      # docker的重启策略：在容器退出时总是重启容器，但是不考虑在Docker守护进程启动时就已经停止了的容器
      restart: unless-stopped
      environment:
        MYSQL_ROOT_PASSWORD: xxxx # root用户的密码
        MYSQL_USER: xxxx # 创建新用户
        MYSQL_PASSWORD: xxxx # 新用户的密码
      ports:
        - 48532:3306
      volumes:
        - ./mysql/data:/var/lib/mysql
        - ./mysql/conf:/etc/mysql/conf.d
        - ./mysql/logs:/logs

    gin-blog-backend:
      image: alexcld/gin-blog:0.0.8
      container_name: gin-blog-backend
      restart: always
      environment:
        GIN_MODE: debug
      ports:
        - 8123:8123
      networks:
        monitor:
          ipv4_address: 172.62.0.9
      volumes:
      - ./gin-blog-backend/config/config.ini:/app/config/config.ini
      - ./gin-blog-backend/logs/:/app/log/
```
### nginx.conf
```
    server {
        listen       80;
        server_name  localhost;
       
        location ^~ / {
            root   /usr/share/nginx/html/front/dist;
            index  index.html index.htm;
        }
        location ^~ /admin {
            root   /usr/share/nginx/html/backend/dist;
            index  index.html index.htm;
        }
        location ^~ /blog/ { 
            proxy_pass http://10.0.4.10:8123/;
            proxy_redirect off;
            proxy_set_header Host $http_host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        }
```
