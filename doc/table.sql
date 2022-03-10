drop table if exists `user`;
CREATE TABLE `user` (
    `uid` int UNSIGNED NOT NULL AUTO_INCREMENT,
    `nick` varchar(8) NOT NULL DEFAULT '' COMMENT '昵称',
    `status` tinyint UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态 1在线 2离线',
    `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户表';