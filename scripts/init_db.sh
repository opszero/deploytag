#!/bin/bash

set -e

# Wait until the connection is available or timeout after 10 seconds
timeout 10 /scripts/db_wait.sh

source /scripts/set_env.sh

if rake db:exists DATABASE_USER=deployer
then
    DATABASE_USER=$DATABASE_OWNER DATABASE_PASSWORD=$DATABASE_OWNER_PASSWORD rake db:migrate
else
    # create a database using the deployer account and set the
    # ownership to the service user
    DATABASE_USER=deployer rake db:create
    DATABASE_USER=deployer rake db:alter_owner
    DATABASE_USER=deployer rake db:add_extensions
    DATABASE_USER=$DATABASE_OWNER DATABASE_PASSWORD=$DATABASE_OWNER_PASSWORD rake db:migrate
    DATABASE_USER=$DATABASE_OWNER DATABASE_PASSWORD=$DATABASE_OWNER_PASSWORD rake db:seed
fi
