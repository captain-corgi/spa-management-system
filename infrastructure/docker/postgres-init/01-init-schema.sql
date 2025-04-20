-- Initialize spa database with sample schema and data for local development

-- Example: Insert sample customers
table if not exists customers (
    id serial primary key,
    name varchar(100),
    email varchar(100),
    created_at timestamp default now()
);

insert into customers (name, email) values
('Alice Example', 'alice@example.com'),
('Bob Example', 'bob@example.com');
