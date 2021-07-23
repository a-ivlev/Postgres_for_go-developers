# Postgres_for_go-developers

Создаём базу данных с именем mydb
create database migrate_db;

Создаём нового пользователя
create user migrate_user password 'secret';

даём ему полные права на управление базой:
grant all privileges on database migrate_db to migrate_user;

для работы с миграциями оспользуемся утилитой
github.com/golang-migrate/migrate

Если на вашем компьютере установлен язык GO тогда установить её можно следующей командой
go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate

Установить программу migrate с помощью команды
go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate
не получилось из-за ошибки
go get: module github.com/golang-migrate/migrate@upgrade found (v3.5.4+incompatible), but does not contain package github.com/golang-migrate/migrate/cmd/migrateЯ не знаю пока, что она означает.

Поэтому я скачал с github архив, разархивировал его перешёл в директорию с исходным кодом и скомпилировал файл программы командой 
go build -o ~/go/bin/migrate -tags 'postgres' cmd/migrate/*.go
Так как в папке migrate кроме файла main.go находиться файл version.go
Параметром -o указываем куда положить скомпилированый файл.

ОБЯЗАТЕЛЬНО НУЖНО ДОБАВИТЬ ТЕГ 'postgres'
Установить программу migrate с помощью команды
go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate
не получилось из-за ошибки
go get: module github.com/golang-migrate/migrate@upgrade found (v3.5.4+incompatible), but does not contain package github.com/golang-migrate/migrate/cmd/migrateЯ не знаю пока, что она означает. Поэтому я скачал с github архив, разархивировал его перешёл в директорию с исходным кодом и скомпилировал файл программы командой 
go build -o ~/go/bin/migrate -tags 'postgres' cmd/migrate/*.go
Так как в папке migrate кроме файла main.go находиться файл version.go
Параметром -o указываем куда положить скомпилированый файл.

ОБЯЗАТЕЛЬНО НУЖНО ДОБАВИТЬ ТЕГ 'postgres'

Для файлов миграций выделяется отдельная директория, например migrations. Все файлы миграций находятся там. Используя команду
migrate create -seq -ext sql -dir migrations init_schema
создаём первую миграцию.
В директории migrations создалось 2 пустых файла
000001_init_schema.up.sql
000001_init_schema.down.sql
в файле 000001_init_schema.up.sql описывается схема БД при применении миграции, а в файле 000001_init_schema.down.sql описывается схема как эти изменения откатить назад при необходимости.

накатить миграции на пустую базу данных можно следующей командой:
migrate -database 'postgresql://migrate_user:secret@localhost:5432/migrate_db?sslmode=disable' -path migrations up

А откатить назад командой
migrate -database "postgresql://migrate_user:secret@localhost:5432/migrate_db?sslmode=disable" -path migrations down

