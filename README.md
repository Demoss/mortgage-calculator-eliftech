# mortgage-calculator-eliftech

:white_check_mark: Готова лише бек частина, не зміг реалізувати з фронтом. :black_square_button: Маю намір цьому навчитись в школі

Для коректної роботи програми необхідно підключити postgresql

Спочатку завантажимо докер https://www.docker.com/products/docker-desktop/

потім "піднімемо" локально базу даних

docker run --name mortgage-calc -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=qwerty -p 5432:5432 postgres:latest 


(якщо юзер або пароль буде інакший, то необхідно змінити його і в файлах конфігурацій config/config.yml та .env відповідно)

На данному етапі база пуста

Для обрахунку іпотеки можна використатати заготовлені дані. Вони знаходяться в папці migrations/000001_init.up.sql
Для вставки даних необхідно

1) підключитись через IDE/docker до бд ![image](https://user-images.githubusercontent.com/57154344/163961205-4750334f-1c04-4ff5-81e6-850805d2784f.png)

2) відкрити консоль(якщо через IDE) ![image](https://user-images.githubusercontent.com/57154344/163961818-d0cbf741-ebdf-43ca-8a7b-037d0698c690.png)
![image](https://user-images.githubusercontent.com/57154344/163961951-ee3ff0da-0052-4414-8ff2-ab4789bf5ab6.png)

3) вставити дані в бд-консоль з файлу міграцій та запустити вставку![image](https://user-images.githubusercontent.com/57154344/163963359-dfd4ac9a-7f8a-4435-b761-1dbfe44a690d.png)

## Для користування програмою необхідно встановити Postman https://www.postman.com/

### Структура банку:

-----



  Name           string json:"name"
  
	Rate           int    json:"rate"
  
	MaxLoan        int    json:"maxLoan"
  
	MinDownPayment int    json:"minDownPayment"
  
	LoanTerm       int    json:"loanTerm"

API для користування:

#### з банками:

* Get request localhost:8000/api/bank/ - отримати всю інфо про всі банки 

* Post request localhost:8000/api/bank/ - додати новий банк(універсальність банку вираховується за назвою)

* Get request localhost:8000/api/bank/bank_name - інфо про конкретний банк

* Put request localhost:8000/api/bank/bank_name - зміна будь-якого поля банку

* delete request localhost:8000/api/bank/bank_name - видалення банку

#### з калькулятором:

### Дані для обрахунку іпотеки:
	InitialLoan int    json:"initialLoan"
  
	DownPayment int    json:"downPayment"
  
	BankName    string json:"bankName"
  
* Get request localhost:8000/api/calc

Приклад: localhost:8000/api/calc?initialLoan=1000&downPayment=8000&bankName=OSCHADBANK
