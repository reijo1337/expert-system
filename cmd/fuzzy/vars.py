from fuzz_common import trapezioid, Variable, constant, Rule, gaus


# X1
priority = Variable('Показатель приоритета профиля для студента', {
    # 'низкий': trapezioid(0, 0, 30, 40, 1),
    'низкий': gaus(20, 10, 1),
    'средний': gaus(50, 20, 1),
    # 'средний': trapezioid(30, 40, 60, 70, 1),
    # 'высокий': trapezioid(60, 70, 100, 100, 1)
    'высокий': gaus(70, 15, 1),
}, 0, 100)

# X2
preparedness = Variable('Уровень профильной подготовки студента', {
    'неуд': trapezioid(0, 0, 40, 50, 1),
    'уд': trapezioid(40, 60, 60, 70, 1),
    'хор': trapezioid(60, 80, 80, 90, 1),
    'отл': trapezioid(80, 100, 100, 100, 1)
}, 0, 100)

# X3
expediency = Variable('Показатель целесообразности обучения по профилю', {
    'ноль': trapezioid(0, 0, 20, 30, 1),
    'низкий': trapezioid(15, 35, 40, 50, 1),
    'средний': trapezioid(35, 55, 60, 70, 1),
    'высокий': trapezioid(55, 75, 80, 90, 1),
    'полный': trapezioid(75, 90, 100, 100, 1)
}, 0, 100)

rules = [
Rule(['высокий', 'отл'], [priority, preparedness], expediency, 'полный'),
Rule(['высокий', 'хор'], [priority, preparedness], expediency, 'высокий'),
Rule(['высокий', 'уд'], [priority, preparedness], expediency, 'средний'),
Rule(['высокий', 'неуд'], [priority, preparedness], expediency, 'низкий'),
Rule(['средний', 'отл'], [priority, preparedness], expediency, 'высокий'),
Rule(['средний', 'хор'], [priority, preparedness], expediency, 'средний'),
Rule(['средний', 'уд'], [priority, preparedness], expediency, 'низкий'),
Rule(['средний', 'неуд'], [priority, preparedness], expediency, 'ноль'),
Rule(['низкий', 'отл'], [priority, preparedness], expediency, 'средний'),
Rule(['низкий', 'хор'], [priority, preparedness], expediency, 'низкий'),
Rule(['низкий', 'уд'], [priority, preparedness], expediency, 'ноль'),
Rule(['низкий', 'неуд'], [priority, preparedness], expediency, 'ноль'),
]