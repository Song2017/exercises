import numbers
from abc import abstractmethod, ABCMeta


# 元类是类的类， 可以操作class的行为， 属性
class C(metaclass=ABCMeta):
    @abstractmethod
    def my_abstract_method(self, ):
        pass


class Field:
    def __init__(self, db_column):
        self.db_column = db_column
        self._value = None

    def __get__(self, instance, owner):
        return self._value

    def __set__(self, instance, value):
        self._value = value


class CharField(Field):
    def __set__(self, instance, value):
        if not isinstance(value, str):
            raise ValueError("string value need")
        self._value = value


class IntField(Field):
    def __set__(self, instance, value):
        if not isinstance(value, numbers.Integral):
            raise ValueError("int value need")
        self._value = value


class ModelMetaClass(type):
    def __new__(mcs, name, bases, attrs, **kwargs):
        if name == "BaseModel":
            return super().__new__(mcs, name, bases, attrs)
        attrs["_fields"] = {}
        for k, v in attrs.items():
            if isinstance(v, Field):
                attrs["_fields"].update({k: v})
        if attrs.get("Meta"):
            attrs["_meta"] = getattr(attrs.get("Meta"), "db_table")
        return super().__new__(mcs, name, bases, attrs)


class BaseModel(metaclass=ModelMetaClass):
    def __init__(self, *args, **kwargs):
        for key, value in kwargs.items():
            setattr(self, key, value)

    def save(self):
        type(self)
        fields = []
        values = []
        for key, value in getattr(self, "_fields").items():
            db_column = getattr(value, "db_column")
            if db_column is None:
                db_column = key.lower()
            fields.append(db_column)
            value = getattr(self, key)
            values.append(str(value))

        sql = "insert {db_table} ({fields}) value({values})".format(
            db_table=getattr(self, "_meta"),
            fields=",".join(fields), values=",".join(values))
        print(sql)


class User(BaseModel, metaclass=ABCMeta):
    name = CharField(db_column="name")
    age = IntField(db_column="age")

    class Meta:
        db_table = "user"


if __name__ == "__main__":
    # User = type("User", (), {"name": "123"})
    # user = User()
    u = User(name="test", age=11)
    u.save()
