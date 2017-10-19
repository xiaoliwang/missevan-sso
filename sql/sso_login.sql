SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for sso_login
-- ----------------------------
DROP TABLE IF EXISTS `sso_login`;
CREATE TABLE `sso_login` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `email` varchar(50) DEFAULT '' COMMENT '用户注册邮箱',
  `mobile` char(15) DEFAULT NULL COMMENT '用户注册手机号',
  `area_code` char(8) DEFAULT '86' COMMENT '手机区号',
  `qq_uid` char(32) DEFAULT NULL COMMENT 'qq唯一id',
  `weibo_uid` char(10) DEFAULT NULL COMMENT '微博唯一id',
  `password` char(32) NOT NULL COMMENT '加密后的用户密码',
  `salt` char(12) NOT NULL COMMENT '加密使用的盐值',
  `is_ban` tinyint(3) unsigned DEFAULT '0' COMMENT '是否被禁止 0 位没有禁止， 1位禁止',
  `modified_time` int(10) unsigned DEFAULT NULL COMMENT '账号，密码或第三方登录修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='单点登录登录表\r\n包含\r\n1.  注册账号及密码\r\n2. 第三方登录信息';
