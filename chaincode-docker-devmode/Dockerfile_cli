FROM hyperledger/fabric-tools:x86_64-1.0.0

RUN go get -u github.com/astaxie/beego && \
        go get -u github.com/beego/bee 

COPY ./api_charity/ ${GOPATH}/src/api_charity
COPY ./chaincode-docker-devmode/msp /etc/hyperledger/msp
COPY ./api_charity /opt/gopath/src/api_charity
COPY ./chaincode-docker-devmode/ /opt/gopath/src/chaincodedev/
COPY ./chaincode/ /opt/gopath/src/chaincodedev/chaincode/

EXPOSE 8080

WORKDIR ${GOPATH}/src/chaincodedev

CMD bee run 

