FROM python:3.11.3

# Select working directory
WORKDIR /code

# Copy requirements.txt to working directory
COPY requirements.txt .

# Install dependencies
RUN pip3 install -r /code/requirements.txt

# Copy source code to working directory
COPY . .

# Run the application
CMD ["python3", "server.py"]