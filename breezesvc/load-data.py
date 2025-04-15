import pandas as pd
from pymongo import MongoClient
import os
from dotenv import load_dotenv

load_dotenv()

MONGO_URI = os.getenv("MONGO_URI")

try:
    client = MongoClient(MONGO_URI)
    db = client["search"]
    instruments_collection = db["instruments"]

    df = pd.read_csv("NSEScripMaster.csv")
    df.columns = df.columns.str.strip().str.replace('"', '')

    df = df[["ShortName", "CompanyName", "ExchangeCode"]]
    df.rename(columns={
        "ShortName": "short_name",
        "CompanyName": "company_name",
        "ExchangeCode": "exchange_code"
    }, inplace=True)

    data = df.to_dict(orient="records")
    instruments_collection.insert_many(data)
    print(f"Successfully loaded {len(data)} instruments into MongoDB.")

except Exception as e:
    print(f"Error loading data into MongoDB: {e}")
finally:
    if 'client' in locals():
        client.close()
