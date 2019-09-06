-- TODO: use migration tool

CREATE TABLE boards (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  title text NOT NULL,
  created_at datetime default current_timestamp,
  updated_at datetime default current_timestamp on update current_timestamp,
  PRIMARY KEY (id)
);
