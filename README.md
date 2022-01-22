# Book Store BE

## Description 
Book Store BE is API for Book Store APP

## Setup and Development 
### Prerequisite
- git
- Go 1.17.1 or Later
- PostgresSQL

### Setup 
- Install Git <br>
  See [Git Installation](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

- Install Go <br>
  See [Go Installation](https://golang.org/doc/install)\

- Install PostgresSQL
  See [Postgres Installation](https://www.postgresql.org/download/)

- Clone this repo
  ```sh
  git clone git@github.com:winartodev/book-store-be.git
  ```
- Copy env.sample to .env
  ```sh
  cp env.sample .env
  ```
- Setup Database 
  ```sh
  rake db:create && rake db:migrate
  ```
- Run Book Store BE 
  ```sh
  make run
  ```
