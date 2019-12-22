from fuzz_common import *

class Rule(object):
    def __init__(self, lhs_names, lhs, rhs, rhs_name, op='AND'):
        self.lhs_names = lhs_names
        self.lhs = list(lhs)
        self.rhs = rhs
        self.rhs_name = rhs_name
        self.op = op

    def infer(self, lex_var_values):
        values = []
        # fuzzification
        for i in range(len(self.lhs)):
            ev = self.lhs[i].fuzzify(lex_var_values[self.lhs[i].name])
            values.append(ev[self.lhs_names[i]])
        # aggregation
        if self.op is 'AND':
            return self.rhs_name, constant(self.rhs.left, self.rhs.right, min(values))

        return self.rhs_name, constant(self.rhs.left, self.rhs.right, max(values))


# X1
priority = Variable('Показатель приоритета профиля для студента', {
    'низкий': trapezioid(0, 0, 30, 40, 1),
    'средний': trapezioid(30, 40, 60, 70, 1),
    'высокий': trapezioid(60, 70, 100, 100, 1)
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

R0 = Rule(['высокий', 'отл'], [priority, preparedness], expediency, 'полный')
R1 = Rule(['высокий', 'хор'], [priority, preparedness], expediency, 'высокий')
R2 = Rule(['высокий', 'уд'], [priority, preparedness], expediency, 'средний')
R3 = Rule(['высокий', 'неуд'], [priority, preparedness], expediency, 'низкий')
R4 = Rule(['средний', 'отл'], [priority, preparedness], expediency, 'высокий')
R5 = Rule(['средний', 'хор'], [priority, preparedness], expediency, 'средний')
R6 = Rule(['средний', 'уд'], [priority, preparedness], expediency, 'низкий')
R7 = Rule(['средний', 'неуд'], [priority, preparedness], expediency, 'ноль')
R8 = Rule(['низкий', 'отл'], [priority, preparedness], expediency, 'средний')
R9 = Rule(['низкий', 'хор'], [priority, preparedness], expediency, 'низкий')
R10 = Rule(['низкий', 'уд'], [priority, preparedness], expediency, 'ноль')
R11 = Rule(['низкий', 'неуд'], [priority, preparedness], expediency, 'ноль')
