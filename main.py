from fastapi import FastAPI

app = FastAPI()

a = {"key":"value"}




@app.get("/")
async def root():
    return {"message": a["key"]}


@app.get("/item/{item_id}")
async def two(item_id):
    return {"message": item_id}


