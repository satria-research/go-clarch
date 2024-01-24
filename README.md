# GO CLARCH

> Go Clean Architecture Boilerplate
> New to Clean Architecture? [Learn Here](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

A familiar architecture is MVC, but MVC is not enough to be a highly agile business solution today [Learn Here](https://khalilstemmler.com/articles/enterprise-typescript-nodejs/when-crud-mvc-isnt-enough/).

On the other hand, the fame of the existing framework becomes a kind of boundary for those who are not used to using a particular framework.

Changes to packages from third parties are also a challenge, such as changing the database from PostgreSql to Mongo.

[![Go Report Card](https://goreportcard.com/badge/github.com/ubaidillahhf/go-clarch)](https://goreportcard.com/report/github.com/ubaidillahhf/go-clarch)

## 📖 Contains

- [The 4 Layer](#-the-layer)
- [The Questions](#-the-questions)
- [Fiber Go web framework](#-fiber-go)
- [Air live reloading](#-air)
- [Debugger](#-debugger)
- [Postman Docs](#-postman-docs)
- [References](#-references)
- [Error Monitoring](#-error-monitoring)
- [Code Quality Sonarqube](#-sonarqube)
- [Secret Management With Vault](#-vault)
- [Containerizaton with Docker](#-docker) 
- [Docker Compose](#-docker-composer)
- [Swagger API Docs](#-swagger)
- JWT Middleware Implement ⏳
- Unit Testing ⏳
- IAM Feature ⏳
- Linter: golangci-lint ⏳
- Husky for run lint+unit test when commit ⏳ 

## 🍰 The Layer

| Layer                | Directory       |
| -------------------- | --------------- |
| Frameworks & Drivers | /app/infra      |
| Interface            | /app/interfaces |
| Usecases             | /app/usecases   |
| Entities             | /app/domain     |

## 🧐 The Questions

- ### Why placing all layer to one folder (app)? &nbsp;

  🥕 Bcs using this infrastructure (clean architecture), we must highlight the 4 layer in folder and ensure the concept is clean (no other file or folder).

- ### Why using json validator than others for request validator? &nbsp;

  🥕 Request validator have several option u can see [here](https://daltontan.com/comparison-of-golang-input-validator-libraries/29/).
  We use [go-playground/validator](github.com/go-playground/validator/v10) because is very simple and clean, bcs just put in json tag, that(json tag) is familiar in go.

- ### Why separate file interface, implementation in Usecase and Frameworks & Drivers layer? &nbsp;
  🥕 The concept come from [Bridge Design Pattern](https://refactoring.guru/design-patterns/bridge/go/example), in Usecase imagine u have 2 user (student and teacher) with same action but different behavior (bcs business rules), in Frameworks & Drivers imagine u must transition change the db from Postgres to Mongo.

## ⚡ Fiber Go

We use fiber for routing and more, you can change whatever you like (echo, gin, chi, etc).
Why fiber? learn [here](https://gofiber.io/)

## 🌊 Air

If you familiar with nodemon in nodejs, air is exactly same. Provide hot reloading when files change with auto build.

Visit: https://github.com/cosmtrek/air for installation guide

## 🧪 Debugger

If you come from PHP maybe you use var_dump(), if u from javasript maybe u use console.log(), in GO u can use fmt.Println() or u can use logging with log.Println().

But if u don't know before, using debugger is awesome and helpful (If u use VS Code), u just go to debug and run the debugger. The config in .vscode in the project. Wanna try? Learn [here](https://medium.com/@slamflipstrom/debugging-with-visual-studio-code-857904a8a590)

## 🐞 Error Monitoring

preq: need account in sentry or glitchtip

Get error early with error monitoring (before client report) with sentry or glitchtip (using same SDK but simplier). Create account and get DSN API and place to env.

## ✅ Sonarqube

preq: sonarqube server (use selfhost or saas with sonarcloud is up to you)

Get confident with code quality in production, placing in jenkins pipeline (see Jenkinsfile). Get quality gate status, when something is bad the pipeline will be aborted.

## 🔐 Vault

preq: you must set vault server (use selfhost or saas is up to you), install extension in jenkins, install jq in your server.

You may face the challenging of management env, the versioning, env mode, and security. Yeah, vault is the solution for this. See in jenkinsfile, vault is in pipeline will replacing your env with new from vault. The env in repo will be unused and ignore.

## 🐳 Docker

Build image with docker build:

```sh
docker build --tag clarch:v0.0.1 .
```

Create container from image we create recently, with daemon mode (-d)

```sh
docker run -d -p 8000:8000 --name clarch clarch:v0.0.1
```

## 🐋 Docker Compose

Run image with dependecy instantly with docker-compose, for this example i use push image first to container registry and place to docker-compose file. Replace with your own

```sh
docker-compose up -d
```

## 📖 Swagger

preq: install swagger cli go install github.com/swaggo/swag/cmd/swag@latest

```sh
swagger init
```

## 📚 References

- https://github.com/khannedy/golang-clean-architecture
- https://github.com/evrone/go-clean-template
- https://github.com/Creatly/creatly-backend
- https://github.com/DoWithLogic/golang-clean-architecture
