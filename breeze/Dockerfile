FROM python:3.11-alpine

WORKDIR /breeze

# System dependencies
RUN apk add --no-cache gcc musl-dev libffi-dev

# Python dependencies
COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

# Copy code
COPY . .

EXPOSE 5000

CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "5000"]
