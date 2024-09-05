-- +goose Up
-- +goose StatementBegin
CREATE TABLE guides (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    createdAt DATETIME NOT NULL
);

INSERT INTO guides (title, content, createdAt)
VALUES ('How to Set Up a Go Development Environment', 'This guide will walk you through setting up a Go development environment on your local machine.', '2024-09-04 10:00:00');

INSERT INTO guides (title, content, createdAt)
VALUES ('Introduction to SQL Databases', 'Learn the basics of SQL databases, including how to create tables and insert data.', '2024-09-04 11:30:00');

INSERT INTO guides (title, content, createdAt)
VALUES ('Building a Simple Web Server in Go', 'This guide demonstrates how to build a basic web server in Go using the net/http package.', '2024-09-04 13:15:00');

INSERT INTO guides (title, content, createdAt)
VALUES ('Optimizing SQL Queries for Performance', 'Discover techniques for optimizing SQL queries to improve database performance.', '2024-09-04 15:45:00');

INSERT INTO guides (title, content, createdAt)
VALUES ('Deploying a Go Application to Production', 'Learn the steps to deploy a Go application to a production environment, including best practices for security and performance.', '2024-09-04 17:00:00');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE guides;
-- +goose StatementEnd
