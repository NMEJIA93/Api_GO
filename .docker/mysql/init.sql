SET @MYSQLDUMP_TEMP_LOG_BIN = @@SESSION.SQL_LOG_BIN;
SET @@SESSION.SQL_LOG_BIN = 0;

SET @@GLOBAL.GTID_PURGED = /*!8000 '+'*/'';

CREATE DATABASE IF NOT EXISTS `go_course`;