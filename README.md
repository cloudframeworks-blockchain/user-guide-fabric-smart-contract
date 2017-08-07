# [云框架]基于区块链的智能合约

![](https://img.shields.io/badge/Release-v1.5-green.svg)
[![](https://img.shields.io/badge/Producer-Gemrails-orange.svg)](CONTRIBUTORS.md)
![](https://img.shields.io/badge/License-Apache_2.0-blue.svg)

区块链（[Blockchain](https://www.blockchain.com/)）是一种解决通过网络买卖商品和服务的双方之间的信任、透明性和责任性问题交易的支持工具，是用分布式数据库识别、传播和记载信息的智能化对等网络，也称为价值互联网。我们可以把区块链比作账本，区块代表账本中的一页，交易细节对所有人公开。

区块链1.0时代应用以比特币为代表，解决了货币和支付手段的去中心化问题；区块链2.0时代则是更宏观的对整个市场的去中心化，利用区块链技术来转换许多不同的资产，通过转让来创建不同资产单元的价值。**智能合约**便是这样的一种应用。

智能合约是将具体条款以计算机语言而非法律语言记录的智能化合同，运行在区块链的多个节点上（分布式环境），新建/调用都由区块链的**Transactions**触发。根据事件描述信息中包含的触发条件，智能合约在条件满足时自动发出预设的数据资源以及包括触发条件的事件，让一组复杂的、带有触发条件的数字化承诺能够按照参与者的意志，正确执行。
 
目前主流智能合约设计包括[Ethereum](https://www.ethereum.org/)和[Fabric](https://www.ibm.com/blockchain/hyperledger.html)。其中**Fabric**即IBM HyperLedger，采用Go和Java实现，运行于Docker中，可支持业务复杂度更高。

本篇[云框架](ABOUT.md)将以**基于Fabric的慈善捐款智能合约**为例，提供**基于区块链的智能合约**的最佳实践。

# 内容概览

* [快速部署](#快速部署)
* [框架说明-业务](#框架说明-业务)
* [框架说明-组件](#框架说明-组件)
    * [区块链](#区块链)
    * [智能合约](#智能合约)
    * [业务前端](#业务前端)
* [如何变成自己的项目](#如何变成自己的项目)
* [更新计划](#更新计划)
* [社群贡献](#社群贡献)

# <a name="快速部署"></a>快速部署

## 一键部署

[一键部署至好雨云帮]()

## 本地部署

1. [准备Docker环境]()
2. 克隆完整代码
 
    ```
    git clone 
    ```
    
3. 进入项目下的charity目录

    ```
    cd user-guide-blockchain/charity
    ```
    
4. 使用docker-compose运行如下命令

    ```
    docker-compose -f docker-charity.yml up -d
    ```

# <a name="框架说明-业务"></a>框架说明-业务

利用区块链技术实现智能合约，某慈善机构将原有捐款流程改造为透明、易追溯的智能捐款系统，可完成以下事务——

1. 注册、初始捐款
2. 增加捐款额
3. 查询账户信息
4. 捐款（随机／指定捐款）
5. 查询捐款记录  
6. 查询余额信息

业务架构如下图所示：

![](https://github.com/cloudframeworks-blockchain/user-guide-blockchain/blob/master/image/fabric_struct.png)

* 完整的Hyperledger Fabric1.0结构
* 暂未使用favric-ca插件
* order
* peer
* cli

# <a name="框架说明-模块"></a>框架说明-模块

## 预执行模块

1. 启动链码(chaincode)

    此步骤在docker-compose执行时即执行于chaincode容器之中
    执行的命令为：

    ```
    CORE_PEER_ADDRESS=peer:7051 CORE_CHAINCODE_ID_NAME=charity:0 ./charity
    ```
    
    在此阶段，链码与任何channel都不相关, 需要在后续步骤中使用实例化

2. 安装链码与实例化

    此步骤在docker-composer执行时即执行于cli容器之中,使用默认通道myc
    执行的命令为：

    ```
    peer chaincode install -p chaincodedev/chaincode/charity -n charity -v 0
    ```

    ```
    peer chaincode instantiate -n charity -v 0 -c '{"Args":[]}' -C myc
    ```

## 事务执行模块

页面调用执行以**捐款人-->捐资**为例时，后端调用的命令如下：

```
peer chaincode invoke -n charity -c '{"Args":["donation", "xxxx", "2000"]}' -C myc
```
![](https://github.com/cloudframeworks-blockchain/user-guide-blockchain/blob/master/image/running.png)

流程如下：

1. 应用程序请求道Peer节点（一个或多个）
2. peer节点分别执行交易（通过chaincode），但是并不将执行结果提交到本地的账本中（可以认为是模拟执行，交易处于挂起状态），参与背书的peer将执行结果返回给应用程序（其中包括自身对背书结果的签名）
3. 应用程序 收集背书结果并将结果提交给Ordering服务节点
4. Ordering服务节点执行共识过程并生成block，通过消息通道发布给Peer节点，由peer节点各自验证交易并提交到本地的ledger中（包括state状态的变化）

## 事务调用

![](https://github.com/cloudframeworks-blockchain/user-guide-blockchain/blob/master/image/fabric%E8%B0%83%E7%94%A8%E7%BB%93%E6%9E%84.png)

组件／模块架构图说明@pujielan

## 组件／peer



## 组件／orderer



## 组件／cli




# <a name="如何变成自己的项目">如何变成自己的项目

1. 编写自己的链码程序转变为自己的项目，就是重新更换链码,需要根据自己的具体业务进行链码文件的编写。下面看一下链码文件中的代码结构

    * 关键引用：

    ```
    "github.com/hyperledger/fabric/core/chaincode/shim"
    "github.com/hyperledger/fabric/protos/peer"
    ```

    * func Init：

    ```
    func (s *SmartContract) Init(api shim.ChaincodeStubInterface) peer.Response {
	   return shim.Success(nil)
    }
    ```

    该方法作用于实例化链码时

    * 方法判断, func Invoke：

    ```
    func (s *SmartContract) Invoke(api shim.ChaincodeStubInterface) peer.Response {

	   function, args := api.GetFunctionAndParameters()

	   switch function {
	   case "donation":
		  return s.donation(api, args)
	   case "queryDealOnce":
		  ...
	   }

	   return shim.Error("Invalid function name.")
    }
    ```

    作用于invoke调用时，对参数的处理。例如`-c '{"Args":["donation", "xxxx", "2000"]}'`中，即调用了donation方法执行了后续业务处理。可以理解为，在写自己的链码程序时这里即是需要按照自己的业务进行修改的合约逻辑。

2. 把链码放置到容器中

    下载go环境镜像，编译链码，推荐本例中使用的name为chaincode的镜像进行

    ```
    docker exec -it chaincode bash
    cd $yourProj
    go build
    ```

3. 修改`docker-charity.yml`文件

    * 将cli的entrypoint指令指定为你个人的chaincode
    * peer中的entrypoint指令,指定安装以及实例化你个人的chaincode

3. 运行docker-compose文件

    ```
    docker-composer -f docker-charity.yml up -d
    ```

4. 完成

    登入cli容器中，可以进行命令操作了

5. webserver

    如果想要使用webserver接入，推荐使用java或者node来进行。
    因为虽然fabric项目由go开发，但是其本身并未提供go sdk... **[Hyperledger Fabric SDKs](http://hyperledger-fabric.readthedocs.io/en/latest/fabric-sdks.html?highlight=sdk)**，本例调用了shell完成处理并为进行go sdk包装，或者推荐如下两go sdk项目以供参考：**[go sdk 1](https://github.com/hyperledger/fabric-sdk-go)**和**[go sdk 2](https://github.com/CognitionFoundry/gohfc)**


# <a name="更新计划"></a>更新计划

* `文档` 
* `组件` 

点击查看[历史更新](CHANGELOG.md)

# <a name="社群贡献"></a>社群贡献

+ QQ群: 
+ [参与贡献](CONTRIBUTING.md)
+ [联系我们](mailto:info@goodrain.com)

-------

[云框架](ABOUT.md)系列主题，遵循[APACHE LICENSE 2.0](LICENSE.md)协议发布。

