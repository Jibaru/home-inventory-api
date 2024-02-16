<h3 align="center">
  <img src="https://raw.githubusercontent.com/Jibaru/home-inventory-api/main/assets/images/logo.png" width="300" alt="Logo"/><br/>
</h3>

<div align="center"><i>Organize your items at home in a simple way.</i></div>

<hr />

<p align="center">
  <a href="https://skillicons.dev">
    <img src="https://skillicons.dev/icons?i=mysql,go,docker,sentry,aws,gmail,github" />
  </a>
</p>

## About

This is a simple API to organize your items at home. It allows you to create rooms, boxes, and items.

You can also add items into boxes, remove items from boxes, and other functionalities.

The API is built with Go and uses MySQL as the database.
It also uses Docker to run the app and the database.
The API is also versioned and uses Sentry to log errors.
The API uses smtp to send emails to users.
The API uses JWT to authenticate users.
The API use AWS S3 to store the assets.

## Get Started

1. Clone the repository
2. Clone `app.env.example` to `app.env` and fill the environment variables
3. Run `docker-compose --env-file app.env up` to start the app or if you use make, then use of the `Makefile` commands
4. Run `docker exec -it home-inventory-api-workspace-1 /bin/bash -c "make run"` to start running the API
5. Also you must need run migrations on your database
6. Access the API at `http://0.0.0.0:your-port`

## Business Keywords

- **User**: A person who uses the API
- **Room**: A place in the house where boxes are located
- **Box**: A container that holds items
- **Item**: An object that is stored in a box
- **ItemKeyword**: A keyword that describes an item
- **Asset**: A file that is stored in the cloud
- **BoxItem**: A relation between a box and an item, it contains the quantity of the item in the box
- **BoxTransaction**: A register of the movement of items in boxes
- **Version**: A version of the API

## Features

- [x] Authentication
    - [x] Register a user
    - [x] Login a user
- [x] Rooms
    - [x] Create a room
    - [x] List all rooms (paginated)
    - [x] Update a room
    - [x] Delete a room
- [x] Boxes
    - [x] Create a box
    - [x] List all boxes (paginated)
    - [x] Update a box
    - [x] Delete a box
    - [x] Add items into a box
    - [x] Remove items from a box
    - [x] Transfer items from a box to another
- [x] Items
    - [x] Create an item with a photo
    - [x] List all items (paginated)
    - [x] Update an item and its photo
    - [x] Delete an item
- [x] Assets
    - [x] Create an asset

## API Structure

The API is versioned and the current version is `v1`.

The API use a Hexagonal Architecture and the structure is as follows:

<div align="center">
  <img src="https://raw.githubusercontent.com/Jibaru/home-inventory-api/main/assets/images/app-schema.png" alt="App schema"/><br/>
</div>

On the application layer, the business logic is implemented. The business logic is implemented.

On the domain layer, there are some elements like:
- **Entities**: They are the main objects of the application.
- **Repositories**: They are the interfaces that define the methods to interact with the database.
- **Services**: They are the interfaces that define the methods to interact with external services.

There are interesting services in domain layer:

- **EventBus**: It is a service that allows to publish async events.
- **EmailSender**: It is a service that allows to send emails to users.
- **FileManager**: It is a service that allows to store files in the cloud.
- **TokenGenerator**: It is a service that allows to generate/decode tokens for users.

On the infrastructure layer, the implementation of the interfaces is done. The implementation is done using the database, the email service, the file storage, etc.

Here is where the API is implemented. The API is implemented using the `echo` framework, and the persistence with `gorm`.

There are two features that cross the entire layers. There are the logger and notifier (sentry).

## Development flow

The coding flow is as follows:

1. Create a task in the project board
2. Create a branch from `main` with the name `feat-number`
3. Implement the feature. The feature must be tested
4. Create a pull request to `main`
5. The pull request must be reviewed and tested. There are some actions to do test.
6. The pull request is merged to `main` is all test pass and if it satisfies the task.

## Testing

The API is tested using the `testing` package and the `testify` package.
For the repositories, the test is done using SQL testing using `sqlmock` package.
