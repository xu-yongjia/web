主要文件夹：
1.GIN：存放后端项目
2.gshop_project_after-master：存放后台管理系统的前端项目
3.orderingSystem：存放点餐系统的前端项目

后端配值：
1.支付宝的产品号、秘钥和支付回调URL
    /GIN/API_front/zhifubao.go 中，需要按照注释配值const声明的全局常量kAppId和kPrivateKey，这些信息需要在支付宝网页注册商家后获取（采用证书加密模式）。如果不修改，则会默认支付给项目开发者的一个沙箱账号，注释中的邮密是与沙箱账号配对的，可作为支付账号进行测试。
    在支付完成后，支付宝服务器会对本项目指定的ip和端口发送回调信息，以告知支付已完成。所以需要配置kServerDomain为运行后端服务器的外网可见IP或域名。我们开发时使用的是cpolar（一个内网穿透工具），开启cpolar内网穿透后将临时域名复制到kServerDomain处，也可接受到回调消息。但cpolar不稳定，且免费版的流量有限，所以有条件的话还是要用稳定的域名。
2.支付宝商店证书
    /GIN/cert 中，有三份证书，如果过期，需要从支付宝-控制台-沙箱开发 中重新下载并放到该文件夹。
3.阿里云
        前端的图片都存储在阿里云上，我们的后端只保留访问阿里云时的URL，不保存图片本体。所以需要配置有效的阿里云账号。我们将阿里云账号配置放在了/GIN/.env文件中，请按注释进行配置。
4.数据库
    ./GIN/main.go中，“dsn”是连接mysql数据库的有关配置。如果您的mysql运行在3306端口上，那么只需要将root改为你的mysql用户名，将123456改为你实际的mysql密码，testdb是我们所用的数据库名称，如果您想要使用其他名称，记得修改这里。我们使用了gorm，使得当运行后端时候，如果您的数据库所需数据表，就会自动按/GIN/DBstruct/中各个文件定义的数据结构来创建数据表。如果想直接使用一个带数据的数据库，可以将/test.sql导入您的数据库，这是我们开发和测试时使用的数据库。
5.运行
    如果您配置好了go的开发和运行环境，那么在GIN目录下打开终端，并使用"go run ."就可以运行后端服务器。
6.测试
        如果您尝试运行项目中的gotest测试，可能会发现有的测试会打印出请求失败信息，这是因为有的gotest参数格式没有和最新的处理函数所要求的数据格式同步，请无视这部分gotest用例。

前端配置：主要配置向后端发送请求时使用的URL
1.orderingSystem：在/orderingSystem/vue.config.js中，将“target”的值修改为您运行后端的机器的域名。
2.gshop_project_after-master：在/gshop_project_after-master/vue.config.js中，将module.exports.devServer.proxy.target（第47行）修改为您运行后端的机器的域名。