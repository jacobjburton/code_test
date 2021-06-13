events {
    # The maximum number of simultaneous connections that can be opened by
    # a worker process.
    worker_connections 1024;
}

http {
    server {
        listen 80 default_server;
        listen [::]:80 default_server;
        proxy_cookie_path / "/; HTTPOnly; Secure";

        server_name ${APP_DOMAIN};

        client_max_body_size 10M;

        resolver 127.0.0.11;

        location / {
            set $upstream http://${APP_HOST}:${APP_PORT};
            proxy_pass $upstream;
        }
    }
}