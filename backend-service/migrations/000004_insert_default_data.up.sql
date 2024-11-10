INSERT INTO users (id, name, email, password, activated)
VALUES (1, 'user', 'admin@mail.com','\x24326124313224545632584f5a784f64355a336e726e574e6958306a756b766d2f3750544a397957584d37316972376d4b64364d517a457468764c69', true);

INSERT INTO users_permissions (user_id, permission_id)
VALUES 
    (1, 1),
    (1, 2),
    (1, 3);
