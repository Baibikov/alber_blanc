## Задача

Рассмотрим все N значные числа в M-ной системе счисления (N чётное, число может начинаться с нуля).

Необходимо найти количество чисел удовлетворяющих условиям:

1)сумма первых N/2 цифр равна сумме последних N/2 цифр
и
2)первые N/2 цифр стоят по возрастанию (не строго), последние N/2 цифр стоят по убыванию (не строго)

Примеры чисел N=6 M=10:

123456 не подходит: 1+2+3 != 4+5+6

357249 не подходит: 3+5+7 = 2+4+9, 249 цифры не по убыванию

455644    подходит: 4+5+5 = 6+4+4, 455 по возрастанию, 644 по убыванию

001100    подходит: 0+0+1 = 1+0+0, 001 по возрастанию, 100 по убыванию

Пример N=8 M=16:

K=389FCBA2 подходит: 3+8+8+F = C+B+A+2, 389F по возрастанию, CBA2 по убыванию
