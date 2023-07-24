CREATE TABLE users (
                       id serial not null unique ,
                       name VARCHAR(255) NOT NULL,
                       username VARCHAR(255) NOT NULL,
                       password_hash VARCHAR(255) NOT NULL
);
CREATE TABLE vacancies (
                           id serial not null unique ,
                           title VARCHAR(255) NOT NULL,
                           description varchar(255),
                           type VARCHAR(50),
                           salary DECIMAL(10, 2)
);

CREATE TABLE users_lists
(
    id      serial                                           not null unique,
    user_id int references users (id) on delete cascade      not null,
    vacancy_id int references vacancy_id (id) on delete cascade not null
);
