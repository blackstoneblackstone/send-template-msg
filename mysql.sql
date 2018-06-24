CREATE TABLE `jmqjwx_template_msg` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `content` longtext NOT NULL COMMENT '内容',
  `url` varchar(128) CHARACTER SET latin1 DEFAULT '',
  `create_time` datetime NOT NULL,
  `template_name` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 COMMENT='模板内容表';

CREATE TABLE `jmqjwxmsg` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `msg_time` datetime NOT NULL,
  `userid` int(10) DEFAULT NULL,
  `courseid` varchar(128) CHARACTER SET latin1 DEFAULT NULL,
  `msg_content` text NOT NULL,
  `type` varchar(64) NOT NULL,
  `cardid` varchar(128) DEFAULT NULL,
  `openid` varchar(256) DEFAULT NULL,
  `appid` varchar(256) DEFAULT '' COMMENT 'appid',
  `templateid` varchar(256) DEFAULT NULL,
  `result` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `openid` (`openid`(255)),
  KEY `appid` (`appid`(255)) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3007536 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='发送模板消息日志表';

CREATE TABLE `jmqjwxapp` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `appid` varchar(128) NOT NULL,
  `appsecret` varchar(128) NOT NULL,
  `wxname` varchar(128) NOT NULL,
  `isuse` int(8) NOT NULL DEFAULT '0',
  `fans_number` int(11) NOT NULL DEFAULT '0',
  `import_number` int(11) NOT NULL COMMENT '导入openid的数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='公众号信息表';

CREATE TABLE `jmqjsend_template_msg_record` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `appid` varchar(128) NOT NULL,
  `send_time` datetime NOT NULL COMMENT '发送时间',
  `send_num` int(10) NOT NULL DEFAULT '0' COMMENT '发送数',
  `sending` int(5) NOT NULL DEFAULT '0' COMMENT '正在发送',
  `tplid` int(10) NOT NULL COMMENT '系统内的模板id',
  `process_id` int(10) NOT NULL,
  `wx_template_id` varchar(128) NOT NULL COMMENT '微信的templateid',
  `success_num` int(10) NOT NULL DEFAULT '0',
  `fail_num` int(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8 COMMENT='发送记录';