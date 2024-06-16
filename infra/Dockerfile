FROM docker.elastic.co/beats/filebeat:8.14.1

COPY filebeat.yml /usr/share/filebeat/filebeat.yml
COPY modules.d /usr/share/filebeat/modules.d
COPY apache.log /usr/share/filebeat/logs/apache.log

USER root 

RUN chown root:root /usr/share/filebeat/filebeat.yml \
    && chmod 644 /usr/share/filebeat/filebeat.yml \
    && chown -R root:root /usr/share/filebeat/modules.d \
    && chmod -R 644 /usr/share/filebeat/modules.d \
    && chown root:root /usr/share/filebeat/logs/apache.log \
    && chmod 644 /usr/share/filebeat/logs/apache.log

