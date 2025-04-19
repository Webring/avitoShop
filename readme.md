# Реализация сервиса мерча Avito

## Выполненено [задание](https://github.com/avito-tech/tech-internship/blob/9459e8244ac43dd5b29f25207a473fc7c84e6ac5/Tech%20Internships/Backend/Backend-trainee-assignment-winter-2025/Backend-trainee-assignment-winter-2025.md)
Едиснтвенное, что ручка `/` перенесена на `/api/profile`


В проекте использовались:
1. Golang
2. Gorm + postgres driver
3. Echo

База данных postgresql. Запускается из docker-compose.

## Использование
### Все настройки проекта вынесены в .env файл.
```
SECRET_KEY="super-secret-key"
PORT=8080
HOST="0.0.0.0"

POSTGRES_USER=postgresuser
POSTGRES_PASSWORD=postgrespass
POSTGRES_DB=avito_shop
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
```

### Запуск проекта

```bash
docker compose up -d
```

## Нужно сделать
- [ ] Хэшировать пароль с солью
- [x] Контейнеризировать приложение
