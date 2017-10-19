SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for sso_user
-- ----------------------------
DROP TABLE IF EXISTS `sso_user`;
CREATE TABLE `sso_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户表主键',
  `user_name` char(20) NOT NULL COMMENT '用户名',
  `user_icon` char(100) DEFAULT NULL COMMENT '用户头像',
  `user_level` tinyint(3) unsigned DEFAULT '0' COMMENT '用户等级\r\n位1 为答题用户\r\n位3 为普通V用户（黑V）\r\n位6 为金V用户',
  `register_device` tinyint(3) unsigned DEFAULT '101' COMMENT '注册来源。101 为未设置，1 为 android， 2为 ios， 3为 web，4为 h5',
  `create_time` int(10) unsigned zerofill NOT NULL COMMENT '创建时间',
  `create_ip` int(10) unsigned DEFAULT NULL COMMENT '创建时的ip',
  `modified_time` int(10) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='单点登录用户信息表\r\n1. 包括用户名，头像，等级等信息\r\n2. 包括创建ip，创建时间等信息\r\n3. 注册设备信息\r\n4. 资料修改时间';
