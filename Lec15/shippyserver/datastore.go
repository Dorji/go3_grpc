//Создаем mongo-сессию
package main

import (
	"context"
	"time"

	//Драйвер/движок mongo  в go
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//CreateClient ...
//Принимает на вход [контекст(обязательный аругмент для всех команд mongo), URI - настройки подключения,
// retry - через сколько секунд попробовать переподключиться]
func CreateClient(ctx context.Context, URI string, retry int32) (*mongo.Client, error) {
	//Пытаемся установить соединение
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	//Проверяем ответ
	// FALIED - 1) бд не существует/нет доступа/бд заблокировано и т.д.
	// 2) проблемы с сетью
	if err := conn.Ping(ctx, nil); err != nil {
		// если количество попыток дошло до 3 - подключиться не удается - завершаемся
		if retry >= 3 {
			return nil, err
		}
		// если была проблема с сетью - дадим побольше времени и попробуем переподключиться
		retry = retry + 1
		time.Sleep(time.Second * 2)
		return CreateClient(ctx, URI, retry)
	}

	return conn, err
}
