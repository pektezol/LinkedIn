CREATE TABLE Users (
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