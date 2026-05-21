# Hitalent

## Описание

API организационной структуры — управление департаментами и сотрудниками.

## Запуск

### Требования

* Docker
* Docker compose

### Команда

```bash
docker compose up --build
```

## Тесты

### Команда

```bash
go test ./...
```

## API

### Департаменты

#### Создать департамент

Request:
```bash
curl -X POST http://localhost:8080/departments \
  -H "Content-Type: application/json" \
  -d '{"name": "Dev"}'
```
Response:
```json
{
  "department": {
    "id": 1,
    "name": "Dev",
    "parent_id": null,
    "created_at": "2026-05-21T12:00:00Z"
  }
}
```

Request:
```bash
curl -X POST http://localhost:8080/departments \
  -H "Content-Type: application/json" \
  -d '{"name": "Backend", "parent_id": 1}'
```
Response:
```json
{
  "department": {
    "id": 2,
    "name": "Backend",
    "parent_id": 1,
    "created_at": "2026-05-21T12:00:00Z"
  }
}
```

#### Получить департамент

Request:
```bash
curl "http://localhost:8080/departments/1?depth=1&include_employees=false"
```
Response:
```json
{
  "department": {
    "id": 1,
    "name": "Dev",
    "parent_id": null,
    "created_at": "2026-05-21T12:00:00Z"
  },
  "children": [
    {
      "id": 2,
      "name": "Backend",
      "parent_id": 1,
      "created_at": "2026-05-21T12:00:00Z"
    }
  ],
  "employees": []
}
```

#### Обновить департамент

Request:
```bash
curl -X PATCH http://localhost:8080/departments/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Development"}'
```
Response:
```json
{
  "department": {
    "id": 1,
    "name": "Development",
    "parent_id": null,
    "created_at": "2026-05-21T12:00:00Z"
  }
}
```

#### Удалить департамент

Request:
```bash
curl -X DELETE "http://localhost:8080/departments/1?mode=cascade"
```
Response: `204 No Content`

Request:
```bash
curl -X DELETE "http://localhost:8080/departments/1?mode=reassign&reassign_to_department_id=2"
```
Response: `204 No Content`

### Сотрудники

#### Создать сотрудника

Request:
```bash
curl -X POST http://localhost:8080/departments/1/employees \
  -H "Content-Type: application/json" \
  -d '{"full_name": "Alex", "position": "Backend developer"}'
```
Response:
```json
{
  "employee": {
    "id": 1,
    "department_id": 1,
    "full_name": "Alex",
    "position": "Backend developer",
    "hired_at": null,
    "created_at": "2026-05-21T12:00:00Z"
  }
}
```
