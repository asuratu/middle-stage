/*
 this file records the migrations that have been applied to the database
 Date: 16/01/2023 17:49:55
*/

-- add columns (delete_time, del_state, version) to table user
ALTER TABLE user ADD COLUMN delete_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE user ADD COLUMN del_state tinyint NOT NULL DEFAULT '0';
ALTER TABLE user ADD COLUMN version bigint NOT NULL DEFAULT '0' COMMENT '版本号';

-- add columns (delete_time, del_state, version) to table category
ALTER TABLE category ADD COLUMN delete_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE category ADD COLUMN del_state tinyint NOT NULL DEFAULT '0';
ALTER TABLE category ADD COLUMN version bigint NOT NULL DEFAULT '0' COMMENT '版本号';

-- add columns (delete_time, del_state, version) to table topic
ALTER TABLE topic ADD COLUMN delete_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE topic ADD COLUMN del_state tinyint NOT NULL DEFAULT '0';
ALTER TABLE topic ADD COLUMN version bigint NOT NULL DEFAULT '0' COMMENT '版本号';

