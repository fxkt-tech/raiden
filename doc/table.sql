drop table if exists `rd_user`;
CREATE TABLE `rd_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nick` varchar(8) NOT NULL DEFAULT '' COMMENT '昵称',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态 1可用 2不可用',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_status` (`status`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = DYNAMIC COMMENT = '用户';

drop table if exists `rd_r_user_followers`;
CREATE TABLE `rd_r_user_followers` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL DEFAULT '0' COMMENT 'uid',
  `followers_uid` int NOT NULL DEFAULT '0' COMMENT '粉丝uid',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`, `followers_uid`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = DYNAMIC COMMENT = '粉丝';

drop table if exists `rd_r_user_following`;
CREATE TABLE `rd_r_user_following` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL DEFAULT '0' COMMENT 'uid',
  `following_uid` int NOT NULL DEFAULT '0' COMMENT '关注的人uid',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`, `following_uid`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = DYNAMIC COMMENT = '关注';

drop table if exists `rd_moment`;
CREATE TABLE `rd_r_moments_history` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '动态id',
  `type` tinyint NOT NULL DEFAULT '1' COMMENT '动态类型 1图文 2视频 3专栏',
  `by_uid` int NOT NULL DEFAULT '0' COMMENT '谁发布的动态',
  `txt` text COMMENT '文字',
  `imgs` json COMMENT '动态的图片数组',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态 1可用 2禁用',
  `publish_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
  PRIMARY KEY (`id`),
  KEY `idx_order` (`by_uid`, `publish_time`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = DYNAMIC COMMENT = '发表的动态';

drop table if exists `rd_r_moments_following`;
CREATE TABLE `rd_r_moments_following` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `moment_id` bigint NOT NULL DEFAULT '0' COMMENT 'rd_moment的id',
  `moment_type` tinyint NOT NULL DEFAULT '1' COMMENT 'rd_moment的type',
  `by_uid` int NOT NULL DEFAULT '0' COMMENT '谁发布的动态',
  `for_uid` int NOT NULL DEFAULT '0' COMMENT '谁关注的动态',
  `txt` text COMMENT '文字',
  `imgs` json COMMENT '动态的图片数组',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态 1可用 2禁用',
  `publish_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
  PRIMARY KEY (`id`),
  KEY `idx_order` (`for_uid`, `publish_time`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = DYNAMIC COMMENT = '关注的动态';

drop table if exists `rd_message`;
CREATE TABLE `rd_message` (
  `id` int NOT NULL AUTO_INCREMENT,
  `sender_uid` int NOT NULL DEFAULT '0' COMMENT '发送者id',
  `recver_uid` int NOT NULL DEFAULT '0' COMMENT '接收者id',
  `content` json COMMENT '聊天内容',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录时间',
  PRIMARY KEY (`id`),
  KEY `idx_ctime` (`create_time`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = DYNAMIC COMMENT = '聊天记录表';