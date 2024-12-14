# Project Improvement Tasks

## 1. Integrate SQLite Instead of JSON Files
- +Set up a connection to SQLite.
- +Create a `notes` table to store note data.
- +Implement CRUD operations to work with SQLite.

## 2. Refactor Project Structure
- Separate project layers:
  - **Storage Layer (storage)**: Handles SQLite interactions.
  - **Service Layer (services)**: Business logic for handling notes.
  - **Handlers (handlers)**: Handles HTTP requests.

## 3. Use `gorilla/mux` for Routing
- +Integrate the `gorilla/mux` router.
- +Add routes with parameters (e.g., `/notes/{id}`).

## 4. Add Authentication and Authorization
- Implement user registration and login.
- Add token-based authentication (e.g., JWT).
- Restrict access to notes to their respective owners.

## 5. Input Validation
- Ensure that fields `title`, `content`, and `author` are properly validated.
- Set limits for the length of the title and content.

## 6. Improve Error Handling
- Create a centralized function for logging and returning errors.
- Implement custom error types with user-friendly messages.

## 7. Enhance RESTful API
- Add routes for filtering and sorting notes.
- Ensure proper HTTP status codes are returned for errors (e.g., `400`, `404`, `500`).

## 8. Swagger Documentation
- Document all endpoints using Swagger.
- Include request and response examples in the documentation.

## 9. Testing
- Write unit tests for CRUD operations.
- Implement integration tests for the API.

## 10. Deploy the Project
- Prepare deployment instructions for a production server.
- Consider using Docker for easy setup and deployment.

---

## Autenfication tasks:

Вот список задач для реализации аутентификации с использованием JWT и SQLite. Задачи сгруппированы по этапам, чтобы упростить выполнение и тестирование.

---

### **Этап 1: Подготовка базы данных**
1. **Создать таблицу пользователей в SQLite:**
   - Поля: `id`, `username`, `hashed_password`, `role`, `created_at`.
   - Убедиться, что поле `username` уникально.
2. **Настроить миграции (если используете GORM или другую библиотеку для работы с базой данных).**
3. **Добавить функцию для хэширования и проверки паролей:**
   - Использовать библиотеку `bcrypt` для безопасности.

---

### **Этап 2: Создание API для аутентификации**
4. **Реализовать регистрацию (`POST /register`):**
   - Принимать `username` и `password`.
   - Проверять уникальность `username`.
   - Хэшировать пароль и сохранять пользователя в базе данных.

5. **Реализовать вход (`POST /login`):**
   - Проверять `username` и `password`.
   - Генерировать Access-токен (JWT) и Refresh-токен.
   - Хранить Refresh-токен в базе данных с привязкой к пользователю.

6. **Реализовать обновление токенов (`POST /refresh`):**
   - Проверять валидность Refresh-токена.
   - Генерировать новый Access-токен.
   - Обновлять Refresh-токен (или оставлять старый, если не истек).

---

### **Этап 3: Middleware для защиты API**
7. **Реализовать middleware для проверки JWT:**
   - Проверять токен в заголовке `Authorization`.
   - Извлекать `userID` и роль (`role`) из токена.
8. **Реализовать middleware для проверки роли:**
   - Для админов: Проверять, что роль в токене равна `admin`.

---

### **Этап 4: Обработка пользовательских данных**
9. **Реализовать эндпоинты для работы с записями:**
   - **GET /notes:** Возвращать только записи, принадлежащие текущему пользователю.
   - **GET /admin/notes:** Возвращать все записи (только для админов).
10. **Обновить существующий функционал:**
    - Добавить проверку пользователя к каждому существующему эндпоинту.

---

### **Этап 5: Безопасность и тестирование**
11. **Реализовать защиту от brute-force атак:**
    - Ограничить число попыток входа (например, через счетчик в памяти или базе данных).
    - Заблокировать учетную запись/доступ после нескольких неудачных попыток (например, на 10 минут).

12. **Тестирование:**
    - Проверить все API на функциональность.
    - Написать тесты для регистрации, входа и проверки токенов.
    - Убедиться, что обычный пользователь не может получить доступ к записям других пользователей.

---

### **Этап 6: Документация**
13. **Добавить документацию к API:**
    - Описать, как использовать эндпоинты.
    - Привести примеры запросов и ответов (включая ошибки).

---
