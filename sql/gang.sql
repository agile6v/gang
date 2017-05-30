
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
    `version` int(11) NOT NULL COMMENT 'task version, increasing value',
    `create_time` int(11) NOT NULL DEFAULT '0' COMMENT 'create time',
    `update_time` int(11) NOT NULL DEFAULT '0' COMMENT 'last modified time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='crontab';


DROP TABLE IF EXISTS `host`;
CREATE TABLE `host` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'incremental id',
  `host` varchar(128) NOT NULL DEFAULT '' COMMENT 'ip address or hostname',
  `task_id` int(11) NOT NULL DEFAULT '0' COMMENT 'task on this host',
  `group_id` int(11) NOT NULL DEFAULT '1' COMMENT 'host group id',
  PRIMARY KEY (`id`),
  UNIQUE KEY `host_task` (`host`,`task_id`),
  KEY `task_idx` (`task_id`),
  KEY `host_idx` (`host`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='the mapping table of host & task ';


DROP TABLE IF EXISTS `host_group`;
CREATE TABLE `host_group` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'incremental id',
  `name` varchar(128) NOT NULL DEFAULT '' COMMENT 'group name',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='host group';

INSERT INTO `host_group` (`id`, `name`) VALUES (1, 'default group');

