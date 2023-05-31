CREATE DATABASE
IF
	NOT EXISTS automic_service DEFAULT CHARACTER
	SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;


GRANT ALL PRIVILEGES ON automic_service.* TO 'root'@'%' identified by "1qaz@WSX";


CREATE TABLE `automic_script` (
                                `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                                `title` varchar(100) DEFAULT '' COMMENT '脚本标题',
                                `desc` varchar(255) DEFAULT '' COMMENT '脚本简述',
                                `version` varchar(100) DEFAULT '' COMMENT '脚本版本',
                                `language` varchar(20) DEFAULT '' COMMENT '脚本语言',
                                `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
                                `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
                                `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
                                `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
                                `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
                                `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',
                                `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
                                PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='脚本管理';