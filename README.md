
# Discord Member Checker

🔍 Простой API-сервис для проверки членства пользователей на серверах Discord

## О проекте

Discord Member Checker - это легкий HTTP API сервис, который позволяет проверить, является ли пользователь участником определенного Discord сервера. Сервис возвращает статус членства пользователя и его имя пользователя.

## Требования

- Go 1.18 или выше
- Токен Discord бота с разрешением SERVER MEMBERS INTENT

## Установка и запуск

### 1. Клонирование репозитория

```bash
git clone https://github.com/your-username/go-members.git
cd go-members
```

### 2. Настройка переменных окружения

Создайте файл `.env` в корне проекта:
DISCORD_TOKEN=ваш_токен_discord_бота_здесь

### 3. Сборка и запуск

#### Windows:

```bash
# Сборка
build.bat

# Запуск
builds\go-members.exe
```

# Сборка вручную
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o go-members-linux main.go

# Запуск
./go-members-linux


### 4. Использование API
API доступен по адресу: `http://localhost:5372`
Основной эндпоинт:
``` 
GET /check/:serverid/:userid
```
Пример запроса:
``` 
GET http://localhost:5372/check/762767033943064607/1409810034224926821
```
Пример ответа:
``` json
{
  "is_member": true,
  "username": "userName"
}
```
## Настройка Discord бота
1. Перейдите на [Discord Developer Portal](https://discord.com/developers/applications)
2. Создайте новое приложение или выберите существующее
3. Перейдите во вкладку "Bot"
4. Под разделом "Privileged Gateway Intents" включите "SERVER MEMBERS INTENT"
5. Скопируйте токен бота и добавьте его в файл `.env`
6. Пригласите бота на свой сервер, используя URL:
``` 
   https://discord.com/api/oauth2/authorize?client_id=ВАШ_CLIENT_ID&permissions=1024&scope=bot
```

