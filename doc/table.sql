CREATE TABLE `tbl_user` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(64) NOT NULL DEFAULT '' COMMENT 'account_name',
    `first_name` varchar(64) NOT NULL DEFAULT '' COMMENT 'user_first_name', 
    `middle_name` varchar(64) NOT NULL DEFAULT '' COMMENT 'user_last_name',          
    `last_name` varchar(64) NOT NULL DEFAULT '' COMMENT 'user_last_name',       
    `password` varchar(256) NOT NULL DEFAULT '' COMMENT 'encoded_password',
    `mobile` varchar(128) DEFAULT '' COMMENT 'mobile',
    `email` varchar(64) DEFAULT '' COMMENT 'email',
    `email_validated` tinyint(1) DEFAULT 0 COMMENT '',
    `mobile_validated` tinyint(1) DEFAULT 0 COMMENT '',
    `signup_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'singup_date_time',
    `last_active` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'last active',
    `profile` text COMMENT 'user_property',
    `status` int(11) NOT NULL DEFAULT '0' COMMENT '',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`name`),
    KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;