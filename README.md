# GO CLARCH

> Go Clean Architecture Boilerplate
> New to Clean Architecture? [Learn Here](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

## 📖 Contains

- [The 4 Layer](#the-layer)
- [Fiber Go web framework](#fiber-go)
- [Air live reloading](#air)
- [Postman Docs](#postman-docs)
- [References](#references)

## 🍰 The Layer

| Layer                | Directory          |
| -------------------- | ------------------ |
| Frameworks & Drivers | /app/infra         |
| Interface            | /app/interfaces    |
| Usecases             | /app/usecases      |
| Entities             | /app/domain/entity |

## ⚡ Fiber Go

We use fiber for routing and more, you can change whatever you like (echo, gin, chi, etc).
Why fiber? learn [here](https://gofiber.io/)

## 🌊 Air

If you familiar with nodemon in nodejs, air is exactly same. Provide hot reloading when files change with auto build.

Visit: https://github.com/cosmtrek/air for installation guide

## 🔖 Postman Docs

- https://documenter.getpostman.com/view/21757760/2s8YstTZ9f

## 📚 References

- https://github.com/khannedy/golang-clean-architecture
