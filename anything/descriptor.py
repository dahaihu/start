# coding: utf8

import time


class A:
    def __init__(self, val):
        self.val = val

    def __get__(self, instance, owner):
        return self.val


def outer(header="1"):
    def wrapper(func):
        def inner(*args, **kwargs):
            return "<h{}>{}</h{}>".format(header, func(*args, **kwargs), header)

        return inner

    return wrapper


@outer(header="2")
def add(a, b):
    return a + b


class LazyProperty:
    def __init__(self, func):
        self.func = func

    def __get__(self, instance, owner):
        if instance is None:
            return self
        value = self.func(instance)
        setattr(instance, self.func.__name__, value)
        return value


class Test:
    @LazyProperty
    def func(self):
        time.sleep(10)
        return 10


def fn(self, name="world"):
    print("Hello, %s" % name)


Hello = type('Hello', (object,), dict(hello=fn))


class ListMetaclass(type):
    def __new__(mcs, name, bases, attrs):
        attrs['add'] = lambda self, value: self.append(value)

        return type.__new__(mcs, name, bases, attrs)


class MyList(list, metaclass=ListMetaclass):
    pass


class Singleton:
    def __new__(cls, *args, **kwargs):
        print("args is ", args)

        print("*xargs is ", kwargs)
        if not hasattr(cls, '_instance'):
            cls._instance = super().__new__(cls)
        return cls._instance


class C(Singleton):
    def __new__(cls, *args, **kwargs):
        print("args is ", args)
        print("*xargs is ", kwargs)
        super().__new__(cls, *args, **kwargs)

    def __init__(self, a, b):
        self.a = a
        self.b = b


class Singleton(type):
    def __call__(cls, *args, **kwargs):
        if not hasattr(cls, '_instance'):
            cls._instance = super().__call__(*args, **kwargs)
        return cls._instance



def singleton(cls):
    def get_instance(*args, **kwargs):
        if not hasattr(cls, '_instance'):
            cls._instance = cls(*args, **kwargs)
        return cls._instance

    return get_instance


class D:
    def __init__(self, a):
        self.__a = a


if __name__ == '__main__':
    print(dict(one='two'))
