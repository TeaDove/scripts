from fastapi import FastAPI

app = FastAPI()
COUNT = 0


@app.get("/")
async def root():
    global COUNT
    COUNT += 1
    return {"message": f"Hello World, {COUNT}"}
