-- Таблица для хранения общего состояния
stats = {
    ids = {},
    counter = 0,
    start_time = 0
}

-- Устанавливаем случайное зерно для генератора случайных чисел
math.randomseed(os.time())

function create_user()
    return string.format('{"name":"test-user", "role":"user", "password": "testuser", "email": "exampleer@example.com"}')
end

-- Основная функция, которая отправляет запросы
request = function()
    stats.counter = stats.counter + 1
  
    -- Фиксируем время начала запроса
    stats.start_time = os.clock()
  
    if stats.counter % 3 == 1 then
        -- Получение списка пользователей
        wrk.method = "GET"
        wrk.body = nil
        return wrk.format(nil, "/users")
    
    elseif stats.counter % 3 == 2 then
        -- Удаление пользователя, если есть доступные id
        if #stats.ids > 0 then
            local id = table.remove(stats.ids, 1)
            wrk.method = "DELETE"
            wrk.body = nil
            return wrk.format(nil, "/users/" .. id)
        else
            -- Переход к созданию пользователя
            stats.counter = stats.counter + 1
            wrk.method = "POST"
            wrk.body = create_user()
            wrk.headers["Content-Type"] = "application/json"
            return wrk.format(nil, "/users")
        end
    
    elseif stats.counter % 3 == 0 then
        -- Создание нового пользователя
        wrk.method = "POST"
        wrk.body = create_user()
        wrk.headers["Content-Type"] = "application/json"
        return wrk.format(nil, "/users")
    end
end

-- Функция для обработки ответа и измерения латентности
response = function(status, headers, body)
  
    -- Обработка ID пользователей, если это GET запрос
    if status == 200 and wrk.method == "GET" then
       
        -- for id in string.gmatch(body, '"ID":(%d+)') do
        --     table.insert(stats.ids, id)
        -- end
        -- for id in string.gmatch(body, '"id":"(%d+)"') do
        --     table.insert(stats.ids, id)
        -- end
    end
end
