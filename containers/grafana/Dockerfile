FROM grafana/grafana:4.4.3

COPY ./config.ini /etc/grafana/grafana.ini
COPY ./dashboards /etc/grafana/dashboards
COPY ./entrypoint.sh /etc/grafana/entrypoint.sh

ENTRYPOINT ["/bin/sh", "/etc/grafana/entrypoint.sh"]
