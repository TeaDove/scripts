class A:
    _name = 0  # for internal use
    str_ = "hi!"  # for avoiding naming conflict
    _ = 0  # "don't care"

    def __str__(self):  # for special methods
        return "A"

    def fun(self):
        print("Public method")  # noqa: T201

    def __fun(self):  # private method
        print("Private method")  # noqa: T201


obj = A()
obj._A__fun()  # the only way to call


def complex(real=0.0, imag=0.0):
    """Form a complex number.

    Keyword arguments:
    real -- the real part (default 0.0)
    imag -- the imaginary part (default 0.0)
    """
    if imag == 0.0 and real == 0.0:
        return 0


from datetime import date  # noqa: E402


class Person:
    __sanity = 0  # private class field

    def __init__(self, name: str, age: int):
        self.name = name
        self._age = age
        Person.__sanity -= 1

    def __str__(self):
        return f"{self.name} {self._age}"

    # @classmethod получает cls, который можно использовать для создания экземпляра класса
    @classmethod
    def fromBirthYear(cls, name: str, year: int):
        return cls(name, date.today().year - year)

    # @staticmethod есть статический метод класса
    @staticmethod
    def isAdult(age: int):
        return age > 18

    # @property - getter
    @property
    def sanity(self):
        return Person.__sanity

    @property
    def age(self):
        print(f'Getting age for "{self}"...')  # noqa: T201
        return self._age

    # @age.setter - сеттер для age
    @age.setter
    def age(self, value: int):
        print(f'Setting age for "{self}"...')  # noqa: T201
        if value < 0:
            raise ValueError("Age cannot be below zero")
        self._age = value


person1 = Person("mayank", 21)
person2 = Person.fromBirthYear("artem", 1996)

print(person1.age)  # noqa: T201
print(person2.age)  # noqa: T201
print(Person.isAdult(22))  # noqa: T201
print(person1.age)  # noqa: T201
person1.age = 10
print(person1.age)  # noqa: T201
try:
    person1.age = -10
except ValueError as e:
    print(e)  # noqa: T201
