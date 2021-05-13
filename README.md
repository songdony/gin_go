# gin_go
### 1.1 项目介绍
    基于gin搭建的脚手架，方便快捷的开发API接口，还可以扩展为自定义的网关代理

### 启动简要说明
    数据库配置在conf里面
    一共有四个案例接口：
    /ping
    /demo/index
    /demo/lists  请求接口：  http://hostname:port/demo/lists
    /demo/detail  请求接口： http://hostname:port/demo/detail?bookid=1

    sql: CREATE TABLE `books` (
    `book_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `book_name` varchar(255) DEFAULT NULL,
    `book_intr` varchar(255) DEFAULT NULL,
    `book_price1` varchar(255) DEFAULT NULL,
    `book_price2` varchar(255) DEFAULT NULL,
    `book_author` varchar(255) DEFAULT NULL,
    `book_press` varchar(255) DEFAULT NULL,
    `book_date` datetime DEFAULT NULL,
    `book_kind` varchar(255) DEFAULT NULL,
    PRIMARY KEY (`book_id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4

    