FROM hyperledger/fabric-peer:x86_64-1.0.0

RUN mkdir -p /etc/hyperledger/msp
COPY ./msp/ /etc/hyperledger/msp

WORKDIR /opt/gopath/src/github.com/hyperledger/fabric/peer

CMD peer node start --peer-chaincodedev=true -o orderer:7050

