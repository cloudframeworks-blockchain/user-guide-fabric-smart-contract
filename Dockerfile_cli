FROM pujielan/fabric_cli:V1.0

RUN go get -u github.com/astaxie/beego && \
    go get -u github.com/beego/bee 

COPY ./api_charity ${GOPATH}/src/api_charity
COPY ./chaincode-docker-devmode/script.sh /opt/gopath/src/chaincodedev/script.sh

WORKDIR /opt/gopath/src/chaincodedev

EXPOSE 8080
CMD /bin/bash -c './script.sh'

