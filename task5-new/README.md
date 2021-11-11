## Task 5. Convert integer to string representation.

Нужно преобразовать целое число в прописной вариант: 12 – двенадцать. 
Программа запускается через вызов главного класса с последующим выполнением GET-запросов по указанному адресу.


#### Eg.
```
http://localhost:8080/number/123 #должно вывести результат 'сто двадцать три'


http://localhost:8080/number/123123123123
{
    "success": true,
    "number": "123123123123",
    "result": "сто двадцать три миллиарда сто двадцать три миллиона сто двадцать три тысячи сто двадцать три"
}

http://localhost:8080/number/abcd
{
    "success": false,
    "number": "abcd",
    "result": "provided string is not a number"
}

http://localhost:8080/number/00000
{
    "success": true,
    "number": "00000",
    "result": "ноль"
}

http://localhost:8080/number/-120032
{
    "success": true,
    "number": "-120032",
    "result": "МИНУС сто двадцать тысяч тридцать два"
}
```
