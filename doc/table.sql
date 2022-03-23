drop table if exists `user`;

CREATE TABLE `user` (
  `uid` int NOT NULL AUTO_INCREMENT,
  `nick` varchar(8) NOT NULL DEFAULT '' COMMENT '昵称',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态 1在线 2离线',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`),
  KEY `idx_status` (`status`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = DYNAMIC COMMENT = '用户';

drop table if exists `followers`;

CREATE TABLE `followers` (
  `uid` int NOT NULL AUTO_INCREMENT,
  `followers_uid` int NOT NULL DEFAULT '0' COMMENT '粉丝uid',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态 1被关注 2被取消关注',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`),
  KEY `idx_status` (`status`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = DYNAMIC COMMENT = '粉丝';

drop table if exists `following`;

CREATE TABLE `following` (
  `uid` int NOT NULL AUTO_INCREMENT,
  `following_uid` int NOT NULL DEFAULT '0' COMMENT '关注的人uid',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态 1关注 2取消关注',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`),
  KEY `idx_status` (`status`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = DYNAMIC COMMENT = '关注';

drop table if exists `message`;

CREATE TABLE `message` (
  `message_id` int NOT NULL AUTO_INCREMENT,
  `sender_uid` int NOT NULL DEFAULT '0' COMMENT '发送者id',
  `recver_uid` int NOT NULL DEFAULT '0' COMMENT '接收者id',
  `content` json COMMENT '聊天内容',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录时间',
  PRIMARY KEY (`message_id`),
  KEY `idx_ctime` (`ctime`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = DYNAMIC COMMENT = '聊天记录表';