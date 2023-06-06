CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    dateofbirth DATE NOT NULL,
    profilepicture TEXT NOT NULL,
    headline TEXT NOT NULL,
    industry TEXT NOT NULL,
    location TEXT NOT NULL,
    bio TEXT NOT NULL,
    cv TEXT NOT NULL
);

CREATE TABLE companies (
    id SERIAL PRIMARY KEY,
    employer_id INT NOT NULL REFERENCES users(id),
    name TEXT NOT NULL,
    industry TEXT NOT NULL,
    logo TEXT NOT NULL,
    location TEXT UNIQUE NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE education (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    school_name TEXT NOT NULL,
    degree TEXT NOT NULL,
    field_of_study TEXT NOT NULL,
    description TEXT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL
);


CREATE TABLE experience (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    company_id INT NOT NULL REFERENCES companies(id),
    title TEXT NOT NULL,
    location TEXT NOT NULL,
    description TEXT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL
);

CREATE TABLE skill (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    name TEXT NOT NULL
);

CREATE TABLE jobs (
    id SERIAL PRIMARY KEY,
    company_id INT NOT NULL REFERENCES companies(id),
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    type TEXT NOT NULL,
    filled BOOLEAN NOT NULL DEFAULT false,
    date DATE NOT NULL DEFAULT now()
);

CREATE TABLE applications (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    job_id INT NOT NULL REFERENCES jobs(id),
    status BOOLEAN NOT NULL,
    date DATE NOT NULL DEFAULT now()
);


CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    text TEXT NOT NULL,
    image TEXT NOT NULL,
    date DATE NOT NULL DEFAULT now()
);

CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    post_id INT NOT NULL REFERENCES posts(id),
    comment TEXT NOT NULL,
    date DATE NOT NULL DEFAULT now()
);

CREATE TABLE likes (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    post_id INT NOT NULL REFERENCES posts(id),
    date DATE NOT NULL DEFAULT now()
);

CREATE TABLE connections (
    id SERIAL PRIMARY KEY,
    sender_id INT NOT NULL REFERENCES users(id),
    reciever_id INT NOT NULL REFERENCES users(id),
    status BOOLEAN NOT NULL,
    date DATE NOT NULL DEFAULT now()
);

CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    sender_id INT NOT NULL REFERENCES users(id),
    reciever_id INT NOT NULL REFERENCES users(id),
    message TEXT NOT NULL,
    date TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE notifications (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    notification TEXT NOT NULL,
    status BOOLEAN NOT NULL DEFAULT false,
    date DATE NOT NULL DEFAULT now()
);