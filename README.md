# Project Setup


## Clone the Repository
First, clone the repository to your local machine.

```
git clone https://github.com/abhijeetlodh/weatherx.git
```

## Setup Client and Server
Open two integrated terminals, one for the client directory and another for the server directory.

## Install Dependencies
In the client terminal, run the following command to install all necessary packages:
```
npm install
```
In the server terminal, run the following command to install all necessary Go packages:
```
go get ./...
```

## Setup Environment Variables
Create a `.env` file in your root directory and add the following variables:
- `PORT`
- `DB`
- `OPENWEATHERMAP_API_KEY`
- `JWT_SECRET_TOKEN`
Make sure we are getting OPENWEATHERMAP API keys from https://home.openweathermap.org/api_keys 
I haven't implement JWT connection to frontend as of now.

## Setup Database
In your SQL application, run the following commands to create necessary tables:

```sql
CREATE TABLE users (
    id serial PRIMARY KEY,
    email text NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz,
    firstname text,
    password text NOT NULL
);

CREATE TABLE weather (
    id serial PRIMARY KEY,
    user_id integer REFERENCES users(id) ON DELETE CASCADE,
    created_at timestamp without time zone NOT NULL,
    temp_f text NOT NULL,
    location text NOT NULL,
    metric text NOT NULL
);
```

## Start Client and Server
In the client terminal, run the following command to start the client:
```
npm start
```
In the server terminal, run the following command to start the server:
```
go run main.go
```
