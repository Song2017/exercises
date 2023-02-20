import uvicorn
from fastapi import FastAPI
from fastapi.middleware.trustedhost import TrustedHostMiddleware

from controllers.order_controller import order_router

app = FastAPI(title="learn fastAPI", docs_url="/api/docs", openapi_url="/api")
app.add_middleware(
    TrustedHostMiddleware, allowed_hosts=[
        "example.com", "0.0.0.0", "localhost"]
)


@app.get("/", tags=["root"])
async def root():
    return {"message": "swagger ui: /api/docs"}


app.include_router(order_router, prefix="/api/v1", tags=["order"])

if __name__ == '__main__':
    # print_hi('PyCharm')
    # uvicorn main:app --reload
    uvicorn.run("main:app", reload=True, host="0.0.0.0", port=80)
