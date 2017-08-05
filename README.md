# [云框架]基于区块链的智能合约

![](https://img.shields.io/badge/Release-v1.5-green.svg)
[![](https://img.shields.io/badge/Producer-Gemrails-orange.svg)](CONTRIBUTORS.md)
![](https://img.shields.io/badge/License-Apache_2.0-blue.svg)

区块链（[Blockchain](https://www.blockchain.com/)）是一种解决通过网络买卖商品和服务的双方之间的信任、透明性和责任性问题交易的支持工具，是用分布式数据库识别、传播和记载信息的智能化对等网络，也称为价值互联网。我们可以把区块链比作账本，区块代表账本中的一页，交易细节对所有人公开。

区块链1.0时代应用以比特币为代表，解决了货币和支付手段的去中心化问题；区块链2.0时代则是更宏观的对整个市场的去中心化，利用区块链技术来转换许多不同的资产，通过转让来创建不同资产单元的价值。**智能合约**便是这样的一种应用。

智能合约是将具体条款以计算机语言而非法律语言记录的智能化合同，运行在区块链的多个节点上（分布式环境），新建/调用都由区块链的**Transactions**触发。根据事件描述信息中包含的触发条件，智能合约在条件满足时自动发出预设的数据资源以及包括触发条件的事件，让一组复杂的、带有触发条件的数字化承诺能够按照参与者的意志，正确执行。
 
目前主流智能合约设计包括**[Ethereum](https://www.ethereum.org/)**和**[Fabric](https://www.ibm.com/blockchain/hyperledger.html)**。其中Fabric即IBM HyperLedger，采用Go和Java实现，运行于Docker中，可支持业务复杂度更高。

本篇[云框架](ABOUT.md)将以xx为例，提供**基于区块链的智能合约**的最佳实践。

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

docker-compose部署步骤@pujielan

## 本地部署
1. [准备docker-compose环境]
2. 进入项目下的charity目录

    ```
    cd user-guide-blockchain/charity
    ```
3. 运行docker-compose

    ```
    docker-compose -f docker-charity.yml up -d
    ```

# <a name="框架说明-业务"></a>框架说明-业务

1.业务场景介绍@pujielan

2.注册、初始捐款

3.增加捐款额

4.查询账户信息

5.捐款（随机／指定捐款）

6.查询捐款记录  

7.查询余额信息


业务架构图@pujielan
本例所示暂未使用fabric—ca组件
完整的Hyperledger Fabric1.0结构如下：

https://pic3.zhimg.com/v2-01893d3f2a9aab6dcc4e9982ba68de8e_r.png


业务架构图说明@pujielan 

1. order
2. peer
3. cli

# <a name="框架说明-组件"></a>框架说明-组件
调用说明
https://pic2.zhimg.com/v2-0e23165fc571eddda56390206990ba39_r.png

组件／模块架构图@pujielan

组件／模块架构图说明@pujielan

## 组件／模块1

与业务结合解释组件／模块1@pujielan

## 组件／模块2

与业务结合解释组件／模块2@pujielan

## 组件／模块3

与业务结合解释组件／模块3@pujielan

# <a name="如何变成自己的项目">如何变成自己的项目

变成自己项目需要改哪，怎么改，改完怎么部署@pujielan

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

