### Magic методы:
https://rszalski.github.io/magicmethods/

### Классовые декоратоы, [pep8](https://www.python.org/dev/peps/pep-0008/), [pep257](https://www.python.org/dev/peps/pep-0257/), приватные методы и тд
#### Андерскоры
```
class A:
    _name = 0                   # for internal use
    str_ = "hi!"                # for avoiding naming conflict
    _ = 0                       # "don't care"
    __surname = "Steve"         # private classs

    def __str__():              # for special methods
        return "A"

    def fun(self):
        print("Public method")

    def __fun(self):            # private method
        print("Private method")

obj = A()
obj._A__fun() # the only way to call
```
#### Документация
```
def complex(real=0.0, imag=0.0):
    """Form a complex number.

    Keyword arguments:
    real -- the real part (default 0.0)
    imag -- the imaginary part (default 0.0)
    """
    if imag == 0.0 and real == 0.0:
        return complex_zero
```
#### Классовые декораторы
```
from datetime import date
class Person:
    __sanity = 0 # private class field

    def __init__(self, name, age):
        self.name = name
        self._age = age
        Person.__sanity -= 1

    def __str__(self):
        return f"{self.name} {self._age}"

    # @classmethod получает cls, который можно использовать для создания экземпляра класса
    @classmethod
    def fromBirthYear(cls, name, year):
        return cls(name, date.today().year - year)

    # @staticmethod есть статический метод класса
    @staticmethod
    def isAdult(age):
        return age > 18

    # @property - getter
    @property
    def sanity(self):
        return Person.__sanity

    @property
    def age(self):
        print(f'Getting age for "{self}"...')
        return self._age

    # @age.setter - сеттер для age
    @age.setter
    def age(self, value):
        print(f'Setting age for "{self}"...')
        if value < 0:
            raise ValueError("Age cannot be below zero")
        self._age = value

person1 = Person('mayank', 21)
person2 = Person.fromBirthYear('artem', 1996)

print(person1.age)
print(person2.age)
print(Person.isAdult(22))
print(person1.age)
person1.age = 10
print(person1.age)
try:
    person1.age = -10
except ValueError as e:
    print(e)
```
