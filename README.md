## Go Project

### Tools
| Backend                                                               | Featrue                      | Desc |
| --------------------------------------------------------------------- | ---------------------------- | ---- |
| [golang](https://github.com/golang/go)                                | Programming Language         |      |
| [gin](https://github.com/gin-gonic/gin)                               | Server Framework             |      |
| [gorm](https://github.com/go-gorm/gorm)                               | ORM                          |      |
| [postgresql-driver](https://github.com/go-gorm/postgres)              | Postgresql Driver            |      |
| [go-playground/validator](https://github.com/go-playground/validator) | Validator                    |      |
| github-action                                                         | CI, Build, Deploy            |      |
| [air](https://github.com/cosmtrek/air)                                | Go gin-server live-reloading |      |

<br/>

| Infra                                               | Featrue | Desc |
| --------------------------------------------------- | ------- | ---- |
| [terraform](https://github.com/hashicorp/terraform) | Iac     |      |
| [aws](https://github.com/hashicorp/terraform)       | Cloud   |      |

<br/>

### Project Structure

```bash
.
├── api
│   ├── controllers
│   ├── data
│   │   ├── request
│   │   └── response
│   ├── models
│   ├── repositories
│   ├── routes
│   └── services
├── docker-compose.yaml
├── internal
│   └── db
└── main.go
```