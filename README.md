# 武汉纺织大学 数据库原理及应用课程 期末设计作业
[![LICENSE](https://img.shields.io/badge/license-Anti%20996-blue.svg?style=flat-square)](https://github.com/996icu/996.ICU/blob/master/LICENSE)
[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu)
*反对996！*

GoLand IDE 教育许可证信息（使用正版人人有责）:

    GoLand 2023.1.2
    Build #GO-231.9011.34, built on May 16, 2023
    授权给 东柳 文
    订阅有效期至 2024年3月2日。
    For educational use only.
    Runtime version: 17.0.6+10-b829.9 amd64
    VM: OpenJDK 64-Bit Server VM by JetBrains s.r.o.
    Windows 11.0
    GC: G1 Young Generation, G1 Old Generation
    Memory: 1524M
    Cores: 8
    Registry:
    suggest.all.run.configurations.from.context=true
    debugger.new.tool.window.layout=true
    ide.completion.variant.limit=500
    ide.experimental.ui=true
    
    Non-Bundled Plugins:
    com.intellij.zh (231.283)

## 数据库
数据库在开发的环境下用的是MySQL8这个版本，对于目前还流行的5.7是否也适用目前还未测试，是否适用这其实也不是我能决定的，这个得看**gorm**的支持情况了
（😀摊手笑，你需要保证的是，你必须得在MySQL中建立两个数据库，一个叫“sale”，另一个叫“sale_user”。前者是用来储存产品信息的，后者是用来单独储存
账户数据的，关于表的设计请参考dao层下的代码和model的各个数据结构模型。当然，如果你有自己的想法，用其他的数据库或者自己另外的储存结构设计，请自行
修改dao和model里边的代码。


本课程设计用到了**gorm**作为对数据库管理系统的访问和操作库，其可以防止SQL注入这样的攻击手段对系统进行非正常的访问。
不过对于各个线程之间的异步与同步并未做处理，在实际的生产环境中是非常危险的，后续会往这方面做完善。目前作业的实现是通过控制台的方式，
之后会考虑采用开源的低代码前端框架搭建出一个前端界面，毕竟3202年了，谁还会在控制台里边用程序呢？（笑😀。
作业更加详细的介绍待补充，例如整个架构是怎么样的，各个模块之间是怎么工作的等等等