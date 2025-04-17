from fastapi import FastAPI, Query, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from breeze_connect import BreezeConnect
from dotenv import load_dotenv
from typing import Optional
import os

load_dotenv()

print("Environment Variables:")
print(os.getenv("BREEZE_API_KEY"))
print(os.getenv("BREEZE_SECRET_KEY"))
print(os.getenv("BREEZE_SESSION_TOKEN"))

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

breeze = BreezeConnect(api_key=os.getenv("BREEZE_API_KEY"))
breeze.generate_session(
    api_secret=os.getenv("BREEZE_SECRET_KEY"),
    session_token=os.getenv("BREEZE_SESSION_TOKEN")
)

@app.get("/historicaldata")
async def get_historical_data(
    interval: str = Query(..., description='e.g. "1minute","5minute","30minute","1day"'),
    from_date: str = Query(..., description="ISO 8601"),
    to_date: str = Query(..., description="ISO 8601"),
    stock_code: str = Query(..., description='e.g. "AXIBAN", "TATMOT"'),
    exchange_code: str = Query(..., description='"NSE", "NFO"'),
    product_type: str = Query(..., description='"futures","options","optionplus","cash","btst","margin"'),
    expiry_date: Optional[str] = Query(None, description="ISO 8601 (optional for cash, btst, margin)"),
    right: Optional[str] = Query(None, description='"call","put","others" (optional for cash, btst, margin)'),
    strike_price: Optional[str] = Query(None, description="Numeric String of Currency (optional for cash, btst, margin)")
):
    
    if product_type not in ["cash", "btst", "margin"]:
        
        if not expiry_date or not right or not strike_price:
            raise HTTPException(
                status_code=400,
                detail="expiry_date, right, and strike_price are required for product_type other than cash, btst, margin."
            )
    try:
        data = breeze.get_historical_data(
            interval=interval,
            from_date=from_date,
            to_date=to_date,
            stock_code=stock_code,
            exchange_code=exchange_code,
            product_type=product_type,
            expiry_date=expiry_date,
            right=right,
            strike_price=strike_price
        )
        return data.get("Success", [])
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
