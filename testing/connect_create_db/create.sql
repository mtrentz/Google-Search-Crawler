/*
docker run --rm -d -v mysql:/var/lib/mysql \
  -v mysql_config:/etc/mysql -p 3306:3306 \
  --name websites-db \
  -e MYSQL_ROOT_PASSWORD=7622446 \
  mysql:latest
*/

-- docker exec -it websites-db bash
-- mysql -u root -p

CREATE DATABASE web_scrape;

USE web_scrape;

CREATE TABLE pages (
    id INT NOT NULL AUTO_INCREMENT,
    page TEXT,
    PRIMARY KEY (id)
);

INSERT INTO pages (page) VALUES ("testing 123");
