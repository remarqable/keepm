FROM postgres:alpine

# Expose ports for keepm web and and postgres db
EXPOSE 80
#EXPOSE 5432

COPY keepm /root
COPY start.sh /root
COPY dbinit.sql /root
WORKDIR /root
CMD ["./start.sh"]
