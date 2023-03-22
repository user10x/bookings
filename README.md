# Project Requirements

## Build a full stack app
* Bread and breakfast webapp
* Allow users to sign up and login securely
* Booking via calendar dates

## Most likely Tech Features: 
* Templating out web pages
* Handlers for Routes for home/about/login/signup etc
* Middleware to do custom http handling
* Packages: Config to become available to other packages;  Render out html; Models for defining databases
* Database connections and ORM connection/ db package
* Session Management
* Centralized Logging

Schema: Visitors


####Users

|   user_id (pk) |   first_name  |   second_name |   email (unique)  |   access_level    |   passwd_hash |
| :---  | :---  | :--- | :---  | :---  | :--- 
|   1234   |   John    |   Deo |   jdoe@owner.com  |   3   |   ######  | 

####Rooms

|   room_id (pk) |   name  |   owner_id (fk) |
| :---  | :---  | :--- | 
|   7777   |  Paradise | 1234

####Reservations/Orders

|   id |   guest_fname  |   guest_lname |   email   |   start_date    |   end_date |
| :---  | :---  | :--- | :---  | :---  | :--- |
|  44   |   alice| bob |  alice.bob@nobody.com|

#### Restriction

| id  | restriction_name
| :---    | :--- | 
|   567   | Already Booked|

#### Room Restriction

| id | room_id | restriction_id |
| :---  | :---  | :--- | :--- |
|   123 | 7777  | 567  |  44  |

## Tech used
Backend Go: 1.20, Chi router
Frontend; HTML, CSS, Javascript
Database: postgres12

## Build mechanisms 
* Make file/ shell scripts
* Dockerfiles for building

