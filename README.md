# sima-rest-api
Лейаут проекта взят с http://github.com/golang-standards/project-layout  
Архитектура проекта взята с http://github.com/bxcodec/go-clean-arch  

Несколько готовых тестов + примеры запросов https://www.postman.com/aviation-specialist-64455861/workspace/sima-api-test/overview 

DATABASE: 
- ~~Модель пользователей~~
- ~~Модель права~~ 
- ~~Миграции через (go-pg/migrations) (Пакет для миграций)~~  

BACKEND:
- ~~Авторизации по логину и паролю~~
- ~~CRUD (CREATE, READ, UPDATE, DELETE) на таблицу пользователей~~
- ~~Логирование~~
- ~~Доступ только после авторизации~~
- - ~~Полный CRUD для администраторов~~
- - ~~Только чтение для пользователей~~
- ~~Общение посредствам REST API~~    

OPTIONAL:
1. Cформированая документация swagger и доступ по url backend
2. ~~Использование Docker, docker-compose, Makefile~~ 
3. ~~Автотестирование (пакет testing - из коробки) routes backend желательно~~


PROBLEM:
1. Вместо даты в ответе объект time.Date
2. ~~Дата трождения обязательное поле~~ 
3. ~~GET возвращает юзера с паролем~~

<pre>
TODO:
1. Валидация
2. Кеш
</pre>
