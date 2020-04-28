DROP DATABASE IF EXISTS haikyu;
CREATE DATABASE haikyu;
use haikyu;

CREATE TABLE imgs (
    id INT NOT NULL AUTO_INCREMENT,
    command VARCHAR(16) NOT NULL,
    text VARCHAR(255) NOT NULL,
    srclink VARCHAR(255) NOT NULL,
    imglink VARCHAR(255) NOT NULL,
    CONSTRAINT pk PRIMARY KEY (id)
);


INSERT INTO image (command, text, srclink, imglink) VALUES (
    'hug',
    '**usr1** hugs **usr2**!',
    'https://www.zerochan.net/1843198',
    'https://static.zerochan.net/Haikyuu%21%21.full.1843198.jpg'
),(
    'hug',
    '**usr1** hugs **usr2** ;3',
    'https://www.zerochan.net/1824476',
    'https://static.zerochan.net/Haikyuu%21%21.full.1824476.jpg'
);