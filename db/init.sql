-- TODO: use migration tool

CREATE TABLE boards (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  name text NOT NULL,
  created_at datetime default current_timestamp,
  updated_at datetime default current_timestamp on update current_timestamp,
  PRIMARY KEY (id)
);

CREATE TABLE sections (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  name text NOT NULL,
  board_id INT UNSIGNED NOT NULL,
  created_at datetime default current_timestamp,
  updated_at datetime default current_timestamp on update current_timestamp,
  PRIMARY KEY (id),
  CONSTRAINT fk_sections_board_id
    FOREIGN KEY (board_id)
    REFERENCES boards(id)
);

CREATE TABLE cards (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  name text NOT NULL,
  description text NOT NULL,
  section_id INT UNSIGNED NOT NULL,
  created_at datetime default current_timestamp,
  updated_at datetime default current_timestamp on update current_timestamp,
  PRIMARY KEY (id),
  CONSTRAINT fk_cards_section_id
    FOREIGN KEY (section_id)
    REFERENCES sections(id)
);
