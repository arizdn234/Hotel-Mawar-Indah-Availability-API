-- 002_seed_data.sql

-- Add seed rooms table
INSERT INTO rooms (name, description, price, available, created_at, updated_at)
VALUES 
    ('Standard Room', 'A cozy room with basic amenities', 100.00, true, NOW(), NOW()),
    ('Deluxe Room', 'A spacious room with luxury amenities', 200.00, true, NOW(), NOW()),
    ('Suite Room', 'An elegant suite with premium amenities', 300.00, true, NOW(), NOW());

-- Add seed reservations table
INSERT INTO reservations (room_id, start_date, end_date, user_id, status, created_at, updated_at)
VALUES
    (1, '2024-05-15', '2024-05-17', 1, 'confirmed', NOW(), NOW()),
    (2, '2024-05-20', '2024-05-22', 2, 'pending', NOW(), NOW()),  
    (3, '2024-05-25', '2024-05-27', 3, 'confirmed', NOW(), NOW());
