FROM nginx:1.15.2-alpine

RUN rm -f /etc/nginx/conf.d/default.conf /etc/nginx/nginx.conf
COPY docker/nginx/nginx.conf.tpl /nginx.conf.tpl

RUN ln -sf /proc/1/fd/2 /var/log/nginx/error.log
RUN ln -sf /proc/1/fd/2 /var/log/nginx/access.log

EXPOSE 80

COPY docker/nginx/start-nginx.sh /start-nginx.sh
RUN chmod a+x /start-nginx.sh

CMD ["/start-nginx.sh"]