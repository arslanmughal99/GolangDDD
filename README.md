This template is a loose implementation of golang DDD (Domain Driven Design). Template already setup user register login and refresh session route. It authenticates user with Custom auth middleware.

##### Project Structure:

`domain\` Contains all the business logic .

`exceptions\` Contain base exception raised from domain logic .

`logger\` Contain logger logic.

`prototypes\` Contain skeleton of app , It contains interfaces for project structure.

`router\` Contain http router framework logic and handlers.

`utils\` Contains utility functions.

##### Current Tech Stack:

App is nearly independent of any major framework. Right now it implements the following stack.

1. GoFiber (for http routing)

2. MongoDb (As Db)

3. Redis (For session handling)

4. Zap (For logging)


##### Docker support:

Dockerfile is optimized for deploying any cloud environment .



FOR IMPROVEMENTS PLEASE FEEL FREE TO CREATE A PULL REQUEST :)
