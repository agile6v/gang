
CREATE DATABASE gang;

USE gang;

DROP TABLE IF EXISTS `task`;
CREATE TABLE `task` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'task id',
    `name` varchar(128) NOT NULL DEFAULT 'gang' COMMENT 'task name',
    `runner` varchar(128) NOT NULL DEFAULT 'agile6v' COMMENT 'executor',
    `command` varchar(1024) NOT NULL DEFAULT 'echo' COMMENT 'executable command',
    `args` varchar(1024) NOT NULL DEFAULT '$PATH' COMMENT 'args of the command',
    `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT 'whether to enable, 0:enabled, 1:disabled',
    `create_time` int(11) NOT NULL DEFAULT '0' COMMENT 'create time',
    `update_time` int(11) NOT NULL DEFAULT '0' COMMENT 'last modified time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='crontab';

DROP TABLE IF EXISTS `host`;
CREATE TABLE `host` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'incremental id',
  `ip` varchar(128) NOT NULL DEFAULT '' COMMENT 'ip address',
  `task_id` int(11) NOT NULL DEFAULT '0' COMMENT 'task on this host',
  PRIMARY KEY (`id`),
  UNIQUE KEY `ip_task` (`ip`,`task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='the mapping table of host & task ';
