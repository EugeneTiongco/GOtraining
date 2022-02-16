# GoTraining Final Project | Movie Tracker

Welcome to the Movie Tracker, an application that aims to serve as a virtual checklist for the movies you've watched and plan to watch.

After running the server, the application will be able to handle multiple users by registering them to the database where users can login in and log out of their accounts. While a user is logged in, they will be able to add movies to their to-watch list by providing its title. The users also have the ability to delete, modify, and mark these movies as finished. Once they are marked as finished, they can provide a numeric score and descriptive review to the movie for future reference.

This application is deployed and ran using Docker which uses a Command Line Interface (CLI) to interact with the application. It uses http API as well as a JSON file (`movieDatabase.json`) to access as locally stored file-based database and a text file (`logs.txt`) to log important events. It also comes with a video presentation to demonstrate what the application does.

Main Features:
- App | API and CLI
- Local Storage | JSON
- List of Movies
- Multiple Users
- Editable Lists

## Instructions
`-cmd=` command typed before any of the following and their required inputs:

1. `add` add a movie with `-title=`
2. `viewToWatch` view the list of movies the user wants to watch, can be used with `-id=`
3. `-viewWatched=` view the list of movies the user wants to watch, can be used with `-id=`
4. `update` edits a specific movie's information, requires `-id=` and `-title=`
5. `delete` delete a movie from the user's to watch list
6. `finish` updates a movie as finished, requires `-title=` `-rating=` `-review=`
7. `login` logs in user with username `-user=` and password `-pass=`
8. `logout` logs out user with username `-user=` and password `-pass=`
9. `signup` creates user with username `-user=` and password `-pass=`


Notes:

Enclose inputs that have spaces inbetween them with quotation marks.

Type `-h` for a reminder of these commands in app.
