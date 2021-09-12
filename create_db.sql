/*
docker run --rm -d -v mysql:/var/lib/mysql \
  -v mysql_config:/etc/mysql -p 3306:3306 \
  --name websites-db \
  -e MYSQL_ROOT_PASSWORD=7622446 \
  mysql:latest
*/

-- docker exec -it websites-db bash
-- mysql -u root -p

/* DROP DATABASE */
DROP DATABASE web_scrape;

CREATE DATABASE web_scrape;

USE web_scrape;

CREATE TABLE IF NOT EXISTS queries (
    id INT NOT NULL AUTO_INCREMENT,
    query_text VARCHAR(255) UNIQUE,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS query_results (
    id INT NOT NULL AUTO_INCREMENT,
    result_rank INT,
    title VARCHAR(255),
    url VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    query_id INT,
    PRIMARY KEY (id),
    FOREIGN KEY (query_id)
      REFERENCES queries(id)
      ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS pages (
    id INT NOT NULL AUTO_INCREMENT,
    domain VARCHAR(255),
    page_url VARCHAR(255) UNIQUE,
    page_text TEXT,
    query_result_id INT,
    PRIMARY KEY (id),
    FOREIGN KEY (query_result_id)
      REFERENCES query_results(id)
      ON DELETE CASCADE
);


/* INSERT INTO pages (page) VALUES ("testing 123"); */
