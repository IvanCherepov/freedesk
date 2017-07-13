FROM quay.io/ivancherepov/freedesk-base 
MAINTAINER ivan.cherapau@gmail.com

COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

RUN mkdir /client /server

ADD client /client
ADD server /server
ADD start.sh /start.sh

EXPOSE 3000

CMD sh /start.sh
