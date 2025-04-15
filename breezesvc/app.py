from flask import Flask, request, jsonify
from flask_cors import CORS 
from breeze_connect import BreezeConnect
from dotenv import load_dotenv
import os

load_dotenv()

app = Flask(__name__)
CORS(app)  


breeze = BreezeConnect(api_key=os.getenv("BREEZE_API_KEY"))
breeze.generate_session(
    api_secret=os.getenv("BREEZE_SECRET_KEY"),
    session_token=os.getenv("BREEZE_SESSION_TOKEN")
)
print("session token: ", os.getenv("BREEZE_SESSION_TOKEN"))
@app.route('/ping', methods=['GET'])
def ping():
    return jsonify({"status": "success", "message": "pong"}), 200

@app.route('/optiondata', methods=['GET'])
def get_option_data():
    try:
        start_date = request.args.get('from_date', '2022-04-19T07:00:00.000Z')
        end_date = request.args.get('to_date', '2022-04-19T18:00:00.000Z')
        expiry = request.args.get('expiry_date', '2022-04-21T07:00:00.000Z')
        interval = request.args.get('interval', '1minute')
        strike = int(request.args.get('strike_price', 17000))
        stock_code = request.args.get('stock_code', 'NIFTY')
        right = request.args.get('right', 'put').lower()
        
        if right not in ['put', 'call']:
            return jsonify({"error": "Invalid right parameter. Use 'put' or 'call'"}), 400

        option_data = breeze.get_historical_data(
            interval=interval,
            from_date=start_date,
            to_date=end_date,
            stock_code=stock_code,
            exchange_code="NFO",
            product_type="options",
            expiry_date=expiry,
            right=right,
            strike_price=strike
        )

        return jsonify(option_data.get("Success", []))

    except Exception as e:
        return jsonify({"error": str(e)}), 500

@app.route('/stockdata', methods=['GET'])
def get_stock_data():
    try:
        
        start_date = request.args.get('from_date', '2022-04-19T07:00:00.000Z')
        end_date = request.args.get('to_date', '2022-04-19T18:00:00.000Z')
        interval = request.args.get('interval', '1minute')
        stock_code = request.args.get('stock_code', 'NIFTY')
        
        exchange_code = "NSE"
        if stock_code in ["NIFTY", "BANKNIFTY", "FINNIFTY", "MIDCPNIFTY"]:
            product_type = "cash"
        else:
            product_type = "cash"  

        
        stock_data = breeze.get_historical_data(
            interval=interval,
            from_date=start_date,
            to_date=end_date,
            stock_code=stock_code,
            exchange_code=exchange_code,
            product_type=product_type
        )

        return jsonify(stock_data.get("Success", []))

    except Exception as e:
        return jsonify({"error": str(e)}), 500

if __name__ == '__main__':
    app.run(debug=True)
