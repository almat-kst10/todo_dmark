CREATE DATABASE todo_list;

CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    is_done BOOLEAN,
    time_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);