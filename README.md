# YellowPages

База данных содержит информацию о людях - имена фамилии и отчества,
а также наиболее вероятный возраст, пол и страну

Запуск - через Makefile:
1. migrate:     - Создание БД 
2. serverStart: - Запуск сервера
3. Взаимодействие с сервером - через браузер, например:
   http://localhost:8080/getInfo?patrinomic=ich&page=1
   http://localhost:8080/deleteInfo?name=Mic
   http://localhost:8080/setInfo?name=Pavel&surname=Vorontsov&patronymic=Aleksandorvich
   http://localhost:8080/upadeInfo?id=5&name=Pavel&surname=Vorontsov&patronymic=Aleksandorvich

Обязательные данные
upadeInfo - Id, имя,фамилия
setInfo - имя,фамилия

Для getInfo и deleteInfo действует пагинция 
 - если страница не указана - будет показана первая

