FROM postgres:alpine

# Expose ports for keepm web and and postgres db
EXPOSE 8080
#EXPOSE 5432

COPY keepm /root
COPY start.sh /root
WORKDIR /root
CMD ["./start.sh"]
