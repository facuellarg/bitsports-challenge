# BITSPORTS CHALLENGE
![Header Image](./images/header_image.png)


- [BITSPORTS CHALLENGE](#bitsports-challenge)
  - [Challenge](#challenge)
  - [Architecture](#architecture)
  - [Patter Designs](#patter-designs)
  - [Diagram](#diagram)
  - [Libraries](#libraries)
  - [Technologies](#technologies)
  - [JWT validation](#jwt-validation)
  - [Try it your self](#try-it-your-self)
    - [Deploy](#deploy)
      - [Docker](#docker)
        - [Install Docker](#install-docker)
        - [Run](#run)
      - [ENV](#env)
      - [Locally](#locally)
        - [Postgres](#postgres)
        - [Go](#go)
        - [Run](#run-1)
    - [EndPoints](#endpoints)
    - [Request](#request)
      - [Postman](#postman)
    - [Playground](#playground)
    - [TODO](#todo)
    
    
## Challenge

<!-- <object data="./pdf/challenge statement.pdf" type="application/pdf" width="700px" height="700px"> -->
<p>This is the challenge statement: <a href="./pdf/*challenge* statement.pdf">Challenge Statement</a>.</p>
<!-- </object> -->


## Architecture

In this project I used the `clean architecture` principles where I could do it, I say this because
`ent` generate a lot of code in a defined structure. </br>
These are the cases where I its were used them: </br>
<table>
  <tbody>
    <tr>
      <th>Folder</th>
      <th >Content</th>
      <th style="word-wrap: break-word;">Why</th>
    </tr>
    <tr>
      <td>usecase</td>
      <td >
        <ul>
          <li>auth</li>
          <li>dataprovider</li>
          <li>password-validator</li>
        </ul>
      </td>
      <td >This 3 folders offer a logic required by the business but that is not intrinsic to the data.</td>
    </tr>
    <tr>
      <td>driver-framework</td>
      <td ><ul><li>middlewares</li></ul></td>
      <td > There is middlewares that work with echo framework.</td>
    </tr>
  </tbody>
</table>




## Patter Designs
- Dependency injection: I used this pattern the `resolver` struct, it received two objects that it need to do its work,this was did using interfaces which make our code less coupled and allow us change the `passwordvalidator` easily.
- Functional options: In the password validator I implement this pattern to allow the user personalize the rules of it with any regex expression that they deems appropriated.

## Diagram
For this challenge I start assuming somethings about the entities, I supposed the following relations between the them.
![database diagram](./images/Database%20ER%20challenge.png)

## Libraries
These are the most relevant libraries that I used:

- **entgo.io/ent v0.10.1**: ORM Facebook made that make easy all the database management.
- **github.com/99designs/gqlgen v0.14.0**: Code generator that give us all the tools to build a graphql server.
- **github.com/golang-jwt/jwt v3.2.2**: Library with API to work with JWT validation.
- **github.com/labstack/echo/v4 v4.7.2** Framework to manage the `REST` requests.
- **github.com/lib/pq v1.10.4. github.com/mattn/go-sqlite3 v1.14.10**: Driver adapters for the databases used.

## Technologies
- graphql
- go
- postgres
- ent
- docker

## JWT validation
Thi section is only to explain the flow of `jwt` token in the project.
- Create the user.
- Login with the created user, this return us a token in the `jwt` cookie.
- All other request take the token from the `Authorization` header and validate its information.

## Try it your self
The next env variables are available to customize the project settings:
- DB_NAME="bitsport" → database name.
- DB_HOST=postgres → database host.
- DB_PASS=1234 → database user password.
- DB_USER=root → database user name.
- DB_PORT=5432 → database user port.
- WEB_PORT=8081 →  application port.
- PRODUCTION=true → if is in production or not

### Deploy
This project includes two ways to be executed:

#### Docker
Using docker compose we can run the project very easy:

##### Install Docker
Follow [Install Docker](https://docs.docker.com/compose/install/)  instruction.

##### Run
Execute the follow command:
```sh 
docker-compose up
```
Now the app is available in the `localhost` address.

#### ENV
To configure env variables in docker you can modify the file `./env` with the values that you want.

#### Locally

##### Postgres
First you need to have installed postgres in your computer, you can do it following [this instructions](https://www.postgresql.org/download/).

##### Go
To install Go follow this [instructions](https://go.dev/doc/install).

##### Run
To run the project execute:
``` sh
go run ./cmd/bitsports/main.go
```

### EndPoints

- GET:`/` → Playground site where you can build the request easily.
- POST:`/query` →  Send the graphql data in the body of the request.
### Request 

#### Postman
I did a postman collection where you can execute all available operations for this project.
All operations were did and arranged in the following `postman` collection for a easy usage.
Maybe for some requests you will need change some body variables because it change accord to the request that you do before.
|version|button|
|-------|------|
|Stable|[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/a04ed980d96acd00aa73?action=collection%2Fimport)|
|Live  | [![Run in Postman](https://run.pstmn.io/button.svg)](https://god.gw.postman.com/run-collection/18842738-ab78585f-2e21-4c70-aa83-eb6811388188?action=collection%2Ffork&collection-url=entityId%3D18842738-ab78585f-2e21-4c70-aa83-eb6811388188%26entityType%3Dcollection%26workspaceId%3Df2ab9ed7-039f-4732-beb9-a2ce69a1ce1e)|

### Playground
Also you can go to `http://localhost:8081/` to use the playground interface for graphql and try it.


### TODO
1. Unit Tests.
2. Benchmarks.
3. Use Echo Middleware instead hand made.
4. Improve the authorization `middleware`.
5. Role base permissions.
6. Personalize errors messages.