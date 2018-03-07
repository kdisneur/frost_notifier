FROM jackfirth/racket:6.12

ENV TIMEZONE=Europe/Paris

RUN ln -snf /usr/share/zoneinfo/${TIMEZONE} /etc/localtime \
    && echo ${TIMEZONE} > /etc/timezone

RUN apt-get update \
    && apt-get install --yes cron

COPY config/crontab /etc/cron.d/frost_notification

RUN chmod 0644 /etc/cron.d/frost_notification \
    && touch /var/log/frost_notification.log \
    && crontab /etc/cron.d/frost_notification

COPY src /var/local/frost_notification

WORKDIR /var/local/frost_notification

CMD ["cron", "-f"]
