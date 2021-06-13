#!/usr/bin/env sh

envsubst '\$APP_DOMAIN \$APP_HOST \$APP_PORT' < /nginx.conf.tpl > /etc/nginx/nginx.conf

nginx -g "daemon off;"
