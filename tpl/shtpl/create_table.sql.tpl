CREATE DATABASE IF NOT EXISTS ${MYSQL_DATABASE};
CREATE USER IF NOT EXISTS '${MYSQL_USERNAME}'@'%' IDENTIFIED BY '${MYSQL_PASSWORD}';
GRANT ALL PRIVILEGES ON ${MYSQL_DATABASE}.* TO '${MYSQL_USERNAME}'@'%';

USE ${MYSQL_DATABASE};
CREATE TABLE \`shici\` IF NOT EXISTS (
  \`id\` bigint NOT NULL,
  \`title\` varchar(64) NOT NULL,
  \`author\` varchar(64) NOT NULL,
  \`dynasty\` varchar(32) NOT NULL,
  \`content\` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci NOT NULL,
  PRIMARY KEY (\`id\`),
  KEY \`title_idx\` (\`title\`),
  KEY \`author_idx\` (\`author\`),
  KEY \`dynasty_idx\` (\`dynasty\`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;