#!/bin/bash
# Change directory to the db_migration directory
cd "$(dirname "$0")/db/db_migration"
# Make the deploy_db.sh script executable
chmod +x deploy_db.sh
# Run the deploy_db.sh script
./deploy_db.sh
