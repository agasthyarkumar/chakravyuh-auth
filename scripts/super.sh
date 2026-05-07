#!/bin/bash

echo "======================================="
echo " Chakravyuh Auth Bootstrap Script"
echo "======================================="

# =========================
# CONFIG
# =========================

DB_NAME="chakravyuh_auth"
DB_USER="postgres"
read -sp "Enter PostgreSQL Password: " DB_PASSWORD

JWT_SECRET="supersecretkey"

BACKEND_PATH="../auth-service"

# =========================
# CHECK POSTGRES
# =========================

echo ""
echo "[1/6] Checking PostgreSQL..."

if ! command -v psql &> /dev/null
then
    echo "PostgreSQL is not installed."
    exit 1
fi

echo "PostgreSQL found."

# =========================
# CREATE DATABASE
# =========================

echo ""
echo "[2/6] Creating database if not exists..."

PGPASSWORD=$DB_PASSWORD psql -U $DB_USER -tc "SELECT 1 FROM pg_database WHERE datname = '$DB_NAME'" | grep -q 1

if [ $? -eq 0 ]; then
    echo "Database already exists."
else
    PGPASSWORD=$DB_PASSWORD createdb -U $DB_USER $DB_NAME
    echo "Database created."
fi

# =========================
# CREATE .ENV
# =========================

echo ""
echo "[3/6] Creating .env..."

cat > $BACKEND_PATH/.env << EOF
PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=$DB_USER
DB_PASSWORD=$DB_PASSWORD
DB_NAME=$DB_NAME

JWT_SECRET=$JWT_SECRET
EOF

echo ".env created."

# =========================
# INSTALL GO DEPENDENCIES
# =========================

echo ""
echo "[4/6] Installing Go dependencies..."

cd $BACKEND_PATH

go mod tidy

echo "Dependencies installed."

# =========================
# RUN BACKEND
# =========================

echo ""
echo "[5/6] Starting backend..."

go run main.go