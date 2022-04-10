drop table if exists `user`;

CREATE TABLE `user` (
  `uid` int NOT NULL AUTO_INCREMENT,
  `nick` varchar(8) NOT NULL DEFAULT '' COMMENT '昵称',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态 1在线 2离线',
  `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`),
  KEY `idx_status` (`status`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = DYNAMIC COMMENT = '用户';

drop table if exists `user_followers`;

CREATE TABLE `user_followers` (
  `uid` int NOT NULL DEFAULT '0' COMMENT 'uid',
  `followers_uid` int NOT NULL DEFAULT '0' COMMENT '粉丝uid',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态 1被关注 2被取消关注',
  `action_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '执行动作的时间',
  PRIMARY KEY (`uid`, `followers_uid`),
  KEY `idx_flist` (`uid`, `status`, `action_time`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = DYNAMIC COMMENT = '粉丝';

drop table if exists `user_following`;

CREATE TABLE `user_following` (
  `uid` int NOT NULL DEFAULT '0' COMMENT 'uid',
  `following_uid` int NOT NULL DEFAULT '0' COMMENT '关注的人uid',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态 1关注 2取消关注',
  `action_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '执行动作的时间',
  PRIMARY KEY (`uid`, `following_uid`),
  KEY `idx_flist` (`uid`, `status`, `action_time`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = DYNAMIC COMMENT = '关注';

drop table if exists `dynamic_following`;

CREATE TABLE `dynamic_following` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `dmc_id` bigint NOT NULL DEFAULT '0' COMMENT 'dmc_id',
  `dmc_type` tinyint NOT NULL DEFAULT '1' COMMENT '动态类型 1图文 2视频 3专栏',
  `by_uid` int NOT NULL DEFAULT '0' COMMENT '谁发布的动态',
  `for_uid` int NOT NULL DEFAULT '0' COMMENT '谁关注的动态',
  `txt` text COMMENT '文字',
  `imgs` json COMMENT '动态的图片数组',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态 1可用',
  `publish_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
  PRIMARY KEY (`id`),
  KEY `idx_order` (`for_uid`, `publish_time`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = DYNAMIC COMMENT = '关注的动态';

drop table if exists `dynamic_history`;

CREATE TABLE `dynamic_history` (
  `dmc_id` bigint NOT NULL AUTO_INCREMENT COMMENT 'dmc_id',
  `dmc_type` tinyint NOT NULL DEFAULT '1' COMMENT '动态类型 1图文 2视频 3专栏',
  `by_uid` int NOT NULL DEFAULT '0' COMMENT '谁发布的动态',
  `txt` text COMMENT '文字',
  `imgs` json COMMENT '动态的图片数组',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态 1可用',
  `publish_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
  PRIMARY KEY (`dmc_id`),
  KEY `idx_order` (`by_uid`, `publish_time`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = DYNAMIC COMMENT = '关注的动态';

drop table if exists `message`;

CREATE TABLE `message` (
  `message_id` int NOT NULL AUTO_INCREMENT,
  `sender_uid` int NOT NULL DEFAULT '0' COMMENT '发送者id',
  `recver_uid` int NOT NULL DEFAULT '0' COMMENT '接收者id',
  `content` json COMMENT '聊天内容',
  `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录时间',
  PRIMARY KEY (`message_id`),
  KEY `idx_ctime` (`ctime`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = DYNAMIC COMMENT = '聊天记录表';