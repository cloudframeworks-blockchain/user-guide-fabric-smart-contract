FROM hyperledger/fabric-orderer:x86_64-1.0.0

RUN mkdir -p /etc/hyperledger/msp
COPY ./msp/ /etc/hyperledger/msp
COPY ./orderer.block /etc/hyperledger/fabric/orderer.block

CMD "orderer"

