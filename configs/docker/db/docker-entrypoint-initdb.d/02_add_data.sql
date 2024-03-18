INSERT INTO users (login, password, access) VALUES
('john_doe', 'qwerty123', 2),
('jane_smith', 'pass123word', 2),
('michael_johnson', 'securepass', 0),
('sarah_williams', 'password456', 2),
('david_brown', 'brownie123', 2),
('emily_miller', 'letmein2024', 0),
('james_taylor', 'taylor2024', 2),
('olivia_anderson', 'anderson567', 2),
('william_clark', 'clark789pass', 0),
('sophia_thomas', 'sophiepass', 2);



INSERT INTO acters (name, sex, dateOfBirth) VALUES
('Tom Hanks', 'Male', '1956-07-09'),
('Meryl Streep', 'Female', '1949-06-22'),
('Leonardo DiCaprio', 'Male', '1974-11-11'),
('Julia Roberts', 'Female', '1967-10-28'),
('Brad Pitt', 'Male', '1963-12-18'),
('Angelina Jolie', 'Female', '1975-06-04'),
('Denzel Washington', 'Male', '1954-12-28'),
('Cate Blanchett', 'Female', '1969-05-14'),
('Johnny Depp', 'Male', '1963-06-09'),
('Nicole Kidman', 'Female', '1967-06-20');


INSERT INTO films (name, description, enterDate, rate, score) VALUES
('Forrest Gump', 'A man with low IQ accomplishes great things in his life.', '1994-07-06', 4.8, 93),
('The Devil Wears Prada', 'A smart but sensible new graduate lands a job as an assistant to Miranda Priestly, the demanding editor-in-chief of a high fashion magazine.', '2006-06-30', 4.2, 87),
('Titanic', 'A seventeen-year-old aristocrat falls in love with a kind but poor artist aboard the luxurious, ill-fated R.M.S. Titanic.', '1997-12-19', 4.5, 95),
('Pretty Woman', 'A man in a legal but hurtful business needs an escort for some social events, and hires a beautiful prostitute he meets... only to fall in love.', '1990-03-23', 4.3, 89),
('Fight Club', 'An insomniac office worker and a devil-may-care soapmaker form an underground fight club that evolves into something much, much more.', '1999-10-15', 4.6, 92),
('Maleficent', 'A vengeful fairy is driven to curse an infant princess, only to discover that the child may be the one person who can restore peace to their troubled land.', '2014-05-30', 4.4, 91),
('Training Day', 'A rookie cop spends his first day as a Los Angeles narcotics officer with a rogue detective who isn''t what he appears to be.', '2001-10-05', 4.1, 88),
('Blue Jasmine', 'A New York socialite, deeply troubled and in denial, arrives in San Francisco to impose upon her sister.', '2013-08-23', 4.0, 86),
('Pirates of the Caribbean: The Curse of the Black Pearl', 'Blacksmith Will Turner teams up with eccentric pirate "Captain" Jack Sparrow to save his love, the governor''s daughter, from Jack''s former pirate allies, who are now undead.', '2003-07-09', 4.7, 94),
('Moulin Rouge!', 'A poet falls for a beautiful courtesan whom a jealous duke covets in this stylish musical, with music drawn from familiar 20th century sources.', '2001-06-01', 4.5, 90);

INSERT INTO film_acters (film_id, acter_id) VALUES
(1, 1),
(1, 2),
(2, 3),
(2, 4),
(3, 5),
(3, 6),
(4, 7),
(4, 8),
(5, 9),
(5, 10);