CREATE TABLE post (
    id SERIAL PRIMARY KEY,
    title character varying(255) NOT NULL UNIQUE,
    thumbnailImage character varying(255) NOT NULL,
    image character varying(255) NOT NULL,
    content character varying(1024) NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE tag (
    id character varying(255) NOT NULL PRIMARY KEY
);

CREATE TABLE postTag (
    postId int,
    tagId character varying(255),
    FOREIGN KEY(postId) REFERENCES post(id),
    FOREIGN KEY(tagId) REFERENCES tag(id)
);

CREATE TABLE location (
    postId int,
    latitude float,
    longitude float,
    name character varying(255),
    image character varying(255),
    FOREIGN KEY(postId) REFERENCES post(id)
);

INSERT INTO post (id, title, thumbnailImage, image, content) VALUES 
(1, 'Post 1', 'test.jpeg', 'test.jpeg', 'There are only two kinds of languages: the ones people complain about and the ones nobody uses.'), 
(2, 'Post 2', 'test.jpeg', 'img1.png', 'Any fool can write code that a computer can understand. Good programmers write code that humans can understand.'), 
(3, 'Post 3', 'test.jpeg', 'test.jpeg', 'First, solve the problem. Then, write the code.'), 
(4, 'Post 4', 'test.jpeg', 'img1.png', 'Java is to JavaScript what car is to Carpet.'), 
(5, 'Post 5', 'test.jpeg', 'test.jpeg', 'Always code as if the guy who ends up maintaining your code will be a violent psychopath who knows where you live.'), 
(6, 'Post 6', 'test.jpeg', 'test.jpeg', 'I''m not a great programmer; I''m just a good programmer with great habits.'), 
(7, 'Post 7', 'test.jpeg', 'test.jpeg', 'Truth can only be found in one place: the code.'), 
(8, 'Post 8', 'test.jpeg', 'test.jpeg', 'If you have to spend effort looking at a fragment of code and figuring out what it''s doing, then you should extract it into a function and name the function after the "what".'), 
(9, 'Post 9', 'test.jpeg', 'test.jpeg', 'The real problem is that programmers have spent far too much time worrying about efficiency in the wrong places and at the wrong times; premature optimization is the root of all evil (or at least most of it) in programming.'), 
(10, 'Post 10', 'test.jpeg', 'test.jpeg', 'SQL, Lisp, and Haskell are the only programming languages that I’ve seen where one spends more time thinking than typing.'), 
(11, 'Post 11', 'test.jpeg', 'test.jpeg', 'Deleted code is debugged code.');

INSERT INTO tag (id) VALUES 
('Voyages'),
('Food'),
('Administration');

INSERT INTO postTag (postId, tagId) VALUES 
(1, 'Voyages'),
(1, 'Food'),
(1, 'Administration'),
(2, 'Voyages'),
(3, 'Food'),
(4, 'Administration'),
(7, 'Voyages'),
(7, 'Food'),
(8, 'Administration');

INSERT INTO location (postId, latitude, longitude, name, image) VALUES
(1, 23.791474915526848, 90.40672529214094, 'dede1', 'restaurant.png'),
(1, 23.781474915526848, 90.41672529214094, 'efefd', 'restaurant.png'),
(1, 23.771474915526848, 90.42672529214094, 'aaaaa', 'restaurant.png'),
(2, 23.761474915526848, 90.43672529214094, 'best', 'shop.png'),
(3, 23.801474915526848, 90.44672529214094, 'zzzz', 'shop.png'),
(3, 23.811474915526848, 90.39672529214094, 'loiuyt', 'restaurant.png'),
(4, 23.821474915526848, 90.38672529214094, 'bzdsh', 'shop.png');
