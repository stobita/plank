-- TODO: use migration tool

DROP TABLE IF EXISTS sections_cards_positions;
DROP TABLE IF EXISTS boards_sections_positions;
DROP TABLE IF EXISTS cards;
DROP TABLE IF EXISTS sections;
DROP TABLE IF EXISTS boards;

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

CREATE TABLE sections_cards_positions (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  section_id INT UNSIGNED NOT NULL,
  card_id INT UNSIGNED NOT NULL,
  position INT UNSIGNED NOT NULL,
  created_at datetime default current_timestamp,
  updated_at datetime default current_timestamp on update current_timestamp,
  PRIMARY KEY (id),
  UNIQUE uq_sections_cards_positions_card_id(card_id),
  UNIQUE uq_sections_cards_positions_position_section_id(section_id, position),
  CONSTRAINT fk_sections_cards_positions_section_id
    FOREIGN KEY (section_id)
    REFERENCES sections(id),
  CONSTRAINT fk_sections_cards_positions_card_id
    FOREIGN KEY (card_id)
    REFERENCES cards(id)
);


CREATE TABLE boards_sections_positions (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  board_id INT UNSIGNED NOT NULL,
  section_id INT UNSIGNED NOT NULL,
  position INT UNSIGNED NOT NULL,
  created_at datetime default current_timestamp,
  updated_at datetime default current_timestamp on update current_timestamp,
  PRIMARY KEY (id),
  UNIQUE uq_boards_sections_positions_section_id(section_id),
  UNIQUE uq_boards_sections_positions_position_board_id(board_id, position),
  CONSTRAINT fk_boards_sections_positions_board_id
    FOREIGN KEY (board_id)
    REFERENCES boards(id),
  CONSTRAINT fk_boards_sections_positions_section_id
    FOREIGN KEY (section_id)
    REFERENCES sections(id)
);

