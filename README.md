To start the server:

**go run main.go**

To start the client:

**npm run start**

Table Creations for Database (Postgres):

create table **problems** ( id int primary key not null, user_id int not null, name varchar(50) not null, description text not null, constraints text, input_format text, output_format text, );

create table **testcases**( id int not null, input text not null, output text not null, sample bool );

CREATE TABLE **submission** ( id int not null, user_id INT NOT NULL );


