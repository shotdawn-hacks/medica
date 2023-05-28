from fastapi import Request, FastAPI
from fastapi.responses import JSONResponse
from fastapi.encoders import jsonable_encoder

import uvicorn
from dataprocessor import DataProcessor
import json

app = FastAPI()

processor = DataProcessor("./files/data_standart.csv","./files/data_standart_sppvr.csv")

@app.get("/")
async def root():
    return {"message": "Hello World"}

@app.post("/classify")
async def classify(request: Request):
    req_body = await request.json()
    processor.load_data(req_body)
    processor.process_data()

    return JSONResponse(content=json.loads(processor.get_result()))


if __name__ == "__main__":
    uvicorn.run(
        "server:app",
        reload=True,
        use_colors=True,
        host="0.0.0.0",
        port=9020,
    )