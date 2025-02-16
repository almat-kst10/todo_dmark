# 📝 Todo DMark - **Не смог добавить скриншоты**

Todo DMark — это кроссплатформенное desktop-приложение для управления задачами, разработанное с использованием **Wails (Go + JavaScript)**. Оно позволяет сохранять состояние задач при перезапуске.

## 🚀 Функционал
- ✅ Создание, удаление задач и смена статусов

## 🛠️ Технологии
- **Backend**: Go (Wails, net/http)  
- **Frontend**: JavaScript (Vue.js)  
- Можно было написать и более интересныее часть Frontend(Добавить архитектуру FSD, валидации, сложные компоненеты, но сделал базово)
- **Database**: Postgres(Локальное хранилище)
- Специально не скрывал config файлы, для проверки

## 📦 Установка и запуск
### 1. Клонирование репозитория
```sh
git clone https://github.com/almat-kst10/todo_dmark.git
cd todo_dmark
wails dev -skipbindings


# README
## About

This is the official Wails Vue-TS template.

You can configure the project by editing `wails.json`. More information about the project settings can be found
here: https://wails.io/docs/reference/project-config

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.

