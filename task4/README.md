## Task 4. Work with files (count|replace strings).

Нужно написать программу, которая будет иметь два режима:
Считать количество вхождений строки в текстовом файле. 
Делать замену строки на другую в указанном файле
Программа должна принимать аргументы на вход при запуске:
<путь к файлу> <строка для подсчёта>
<путь к файлу> <строка для поиска> <строка для замены>



#### Eg.
```
./task4 testData.txt москал
Substring <москал> found in file testData.txt 22 times

./task4 testData.txt Катерина Гарпина
Substring <Катерина> replaced by <Гарпина> in file testData.txt

./task4 testData.txt Катерина Гарпина Мальвина
Program takes 2 or 3 parameters on start: 
                <file name> <substring to count>
                <file name> <substring to replace> <replacing substring>

./task4 testDat.txt 1 2
File testDat.txt doesn't exist.
There were problems replacing substring in file
```
