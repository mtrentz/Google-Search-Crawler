CREATE DATABASE IF NOT EXISTS crawler;

USE crawler;

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
