from fastapi import APIRouter
from fastapi.encoders import jsonable_encoder
from fastapi.responses import JSONResponse

from models.item import Order

order_router = router = APIRouter()


@router.get("/order/{order_id}")
async def get_order(order_id: str) -> JSONResponse:
    """
    :param order_id: FedEx accountNumber with 9 digits
    :return: JSONResponse
    """

    return JSONResponse(content=jsonable_encoder({"result": order_id}))


@router.post("/order/{order_id}")
async def post_order(order: Order, order_id: str) -> JSONResponse:
    """
    :param order: Order
    :param order_id: FedEx accountNumber with 9 digits
    :return: JSONResponse
    """

    return JSONResponse(content=jsonable_encoder({"create result": "order_id"}))
