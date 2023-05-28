from fastapi import Request, FastAPI
from fastapi.responses import JSONResponse
from fastapi.encoders import jsonable_encoder

import uvicorn
import json, requests
import os, time
import uuid
from dataprocessor import DataProcessor

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

def register():
    x = 0
    while x < 5:
        try:
            mock = '{"_id":"","name":"analyzer","address":"","port":"9020"}'

            mock = json.loads(mock)
            mock["_id"] = str(uuid.uuid4())
            mock["address"] = "127.0.0.1"
            x+=1
            resp = requests.post(url="http://"+os.environ['CORE']+":9010/register", json=mock)

            if resp.status_code == 200:
                return
        except:
            time.sleep(2)

if __name__ == "__main__":
    register()
    uvicorn.run(
        "server:app",
        reload=True,
        use_colors=True,
        host="0.0.0.0",
        port=9020,
    )