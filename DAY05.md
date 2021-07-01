# Написание первого RESTful API сервиса

Цель: научиться писать RESTful API микросервисы

```
База (postgres)
Host: postgres.finance.svc.k8s.dldevel
Port: 5432
Name: postgresdb
User: postgresadmin
Password: admin123
```

## Задание 1.

написать сервис, у которого будут методы:

1) нанять сотрудника (201 сreated при успешном выполнении)

```sql
employees.employee_add(
	_id integer,
	_name varchar,
	_last_name varchar,
	_patronymic varchar,
	_phone varchar,
	_position varchar,
	_good_job_count integer
)
```

2) уволить сотрудника

```sql
employees.employee_remove(_id integer)
```

3) изменить личные данные сотрудника

```sql
employees.employee_upd(
	_id integer,
	_name varchar,
	_last_name varchar,
	_patronymic varchar,
	_phone varchar,
	_position varchar,
	_good_job_count integer
)
```

4) получить всех сотрудников  

```sql
employees. get_all(_id integer)
```

5) получить сотрудника по его ID  

```sql
employees.employee_get(_id integer)
```

Условия:

* структура проекта – mvc https://ru.wikipedia.org/wiki/Model-View-Controller
* ошибки необходимо возвращать по стандарту [rfc7807](https://tools.ietf.org/html/rfc7807)
* у каждого метода путь должен начинаться с префикса /api/v1
* все методы должны быть закрыты авторизационным миддлваром, проверяющим наличие права "world"
* методы, принимающие в запросе JSON должны иметь миддлвар проверяющий Content-Type. передавать должны только json
* метод 4) должен иметь accept миддлвар, отдающий xml или json
* метод 5) должен принимать employeeId через path-параметр (http://localhost:8000/api/v1/{employeeId}). для этого иожно использовать библиотеку https://github.com/gorilla/mux
* при запросах в БД необходимо передавать контекст запроса (r.Context()). Если клиент сервиса перестал ожидать ответ, запрос в базу должен прекращаться

## Задание 2.

Добавить технический метод GET /tech/info, который вернет JSON с информацией о приложении:

```js
{ 
	"name": "employees",
	"version": "1.0.0"
}
```

Вспомогательная литература:

* Building RESTful Web services with Go.pdf
* The_Ultimate_Guide_To_Building_Database-Driven_Apps_with_Go.pdf
* Clean_Code.pdf
