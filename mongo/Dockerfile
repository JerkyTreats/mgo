
FROM mongo:latest
COPY ./init.sh /var/mongo/init.sh
ENTRYPOINT [ "mongod", "--smallfiles", "--bind_ip", "0.0.0.0", "--storageEngine", "wiredTiger" ]
EXPOSE 27017