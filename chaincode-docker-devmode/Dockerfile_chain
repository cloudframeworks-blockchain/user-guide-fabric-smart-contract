FROM goodraincloudframeworks/fabric_cli:x86_64-1.0.0

RUN mkdir -p /etc/hyperledger/msp
COPY ./msp /etc/hyperledger/msp

#COPY ./../chaincode-docker-devmode /opt/gopath/src/chaincodedev
COPY ./../chaincode /opt/gopath/src/chaincodedev/chaincode

WORKDIR /opt/gopath/src/chaincodedev

EXPOSE 8080
CMD /bin/bash -c './script.sh'

