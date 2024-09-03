# Gitlab Downloader

Утилита для скачивания проектов из Gitlab.

## Introduction

В данной утилите реализовано скачивание проектов с Gitlab через gitlab-api.

## About

Утилита работает с gitlab-api. 

Скачиваемые проекты ограничены проектами к которым пользователь имеет доступ.

Для сохранения проектов нужен `access token` который имеет нужные для доступа права.

### Функционал 
В данный момент утилита обладает следующим функционалом

- сохранение всех проектов

## Get started

Для начала работы нужно заполнить _.env_ файл. Прмиер файла можно найти в директории проекта
он называется _.env.example_

Для упрощения сборки реализован билд через утилиту `make`

- `make build` - автоматическая сборка проекта
- `make run` - старт проекта при котором произойдет скачивание всех проектов
в директорию _download_ в корне проекта
- `make uninstall` - удалит директорию _build_ из корня проекта, но не удалит скачанные проекты
- `make clean` - удаление директорий _build_ и _download_ 