# 单点登录

[^1]: Central Authentication Service



# 功能设计
改功能主要参考一下两个协议：

1. `CAS 3.0.2` [^1]主要实现的是网站内部的登录信息交互。
2. `OAUTH 2.0` 主要实现的是网站外部第三方登录。



credentials (such as account and password)



## CAS 接口清单

| URI                 | 描述                                       |
| ------------------- | ---------------------------------------- |
| /login              | 登录和认证                                    |
| /logout             | destroy CAS session (logout)             |
| /validate           | service ticket validation                |
| /serviceValidate    | service ticket validation [CAS 2.0]      |
| /proxyValidate      | service/proxy ticket validation [CAS 2.0] |
| /proxy              | proxy ticket service [CAS 2.0]           |
| /p3/serviceValidate | service ticket validation [CAS 3.0]      |
| /p3/proxyValidate   | service/proxy ticket validation [CAS 3.0] |



## CAS 接口说明

### /login

参数大小写敏感。

#### 参数

| 参数名     | 是否必须     | 类型      | 说明                                      |
| ------- | -------- | ------- | --------------------------------------- |
| service | REQUIRED | String  | CAS 客户端服务名                              |
|         | REQUIRED | Boolean | 是否重新验证用户名密码，true为验证用户名密码，false为只验证token |
| client  | REQUIRED | Int     | 1 为 android， 2为 iOS， 3 为 Web， 4为 H5     |



1. `service` [REQUIRED]: the identifier of the application

   url encoded

2. `renew`  [REQUIRED]：

3. `gateway`

4. `ticket`








## API接口

### 获取用户信息

1. 账号密码登录
   - [ ] 国内手机
   - [ ] 境外手机
   - [ ] 邮箱


2. 第三方登录

  - [ ] qq 登录
  - [ ] 微博登录
  - [ ] B站登录 

3. token 获取



### 登出

1. 用户主动登出
2. 后台强制登出



### 注册

1. 手机注册
   1. 境内手机注册
   2. 境外手机注册
2. 邮箱注册
3. 第三方手机注册
4. 第三方邮箱注册



### 更新用户信息



### 获取用户信息



## 界面功能

### 用户禁用功能

### 用户权限功能



## 定时任务

### 清理长期没有使用的token





[^1]: 