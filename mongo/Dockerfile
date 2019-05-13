
FROM mongo:latest
COPY ./init.sh /var/mongo/init.sh
ENTRYPOINT [ "chmod", "+x", "/var/mongo/init.sh" ]
