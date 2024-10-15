-- Insert into boards
INSERT INTO boards (name) VALUES ('Project Alpha'), ('Project Beta');

-- Insert into sections
INSERT INTO sections (name, board_id) VALUES 
('To Do', 1), 
('In Progress', 1), 
('Completed', 1),
('Backlog', 2), 
('To Do', 2), 
('In Review', 2);

-- Insert into cards
INSERT INTO cards (name, description, section_id, limit_time) VALUES 
('Design the homepage', 'Create a wireframe for the homepage', 1, '2024-10-20 12:00:00'),
('Set up database', 'Install and configure MySQL', 1, NULL),
('Develop login feature', 'Implement user authentication and login', 1, '2024-11-10 17:00:00'),
('Create API endpoints', 'Develop the main API endpoints for the project', 2, '2024-10-30 18:00:00'),
('Write API documentation', 'Document the API endpoints', 2, '2024-10-15 15:00:00'),
('Design system architecture', 'Outline the project’s architecture', 3, '2024-10-25 11:00:00'),
('Fix login issue', 'Resolve bug related to user login', 4, '2024-11-01 09:00:00'),
('Create user stories', 'Draft initial user stories for the project', 5, NULL),
('Conduct code review', 'Review the code submitted for the API development', 6, '2024-10-31 14:00:00'),
('Prepare deployment plan', 'Set up the environment and prepare deployment strategy', 6, NULL);

-- Insert into sections_cards_positions
INSERT INTO sections_cards_positions (section_id, card_id, position) VALUES 
(1, 1, 1), 
(1, 2, 2), 
(1, 3, 3),
(2, 4, 1), 
(2, 5, 2), 
(3, 6, 1), 
(4, 7, 1), 
(5, 8, 1), 
(6, 9, 1),
(6, 10, 2);

-- Insert into boards_sections_positions
INSERT INTO boards_sections_positions (board_id, section_id, position) VALUES 
(1, 1, 1), 
(1, 2, 2), 
(1, 3, 3),
(2, 4, 1), 
(2, 5, 2),
(2, 6, 3);

-- Insert into labels
INSERT INTO labels (name, board_id) VALUES 
('High Priority', 1), 
('Bug', 1), 
('Feature', 1), 
('Low Priority', 2), 
('Enhancement', 2), 
('Urgent', 1), 
('Documentation', 2);

-- Insert into cards_labels
INSERT INTO cards_labels (card_id, label_id) VALUES 
(1, 1), 
(2, 1), 
(3, 6), 
(4, 3), 
(5, 7), 
(6, 4), 
(7, 2), 
(8, 4), 
(9, 5), 
(10, 5);

-- Insert into cards_images
INSERT INTO cards_images (card_id, url) VALUES 
(1, 'https://example.com/images/homepage_design.png'), 
(3, 'https://example.com/images/login_feature.png'),
(5, 'https://example.com/images/api_docs.png'),
(7, 'https://example.com/images/login_bug.png'),
(9, 'https://example.com/images/code_review.png'),
(10, 'https://example.com/images/deployment_plan.png');
