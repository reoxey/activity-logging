CREATE TABLE IF NOT EXISTS `users` (
    user_id int(11) PRIMARY KEY AUTO_INCREMENT, 
    name varchar(200) NOT NULL, 
    email varchar(200) NOT NULL, 
    phone varchar(20) NOT NULL 
);

CREATE TABLE IF NOT EXISTS `activities` (
    activity_id int(11) PRIMARY KEY AUTO_INCREMENT,
    user_id int(11) NOT NULL,
    type int(1) NOT NULL,
    timestamp int(11) NOT NULL,
    duration bigint(11) NOT NULL
);
