from typing import Union

from pydantic import BaseModel


class Item(BaseModel):
    name: str
    description: Union[str, None] = None
    price: float
    tax: Union[float, None] = None


class Order(BaseModel):
    name: str
    description: Union[str, None] = None
    price: float
    tax: Union[float, None] = None