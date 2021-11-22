FROM mysql:latest

# COPY create_db.sh /mysql/create_db.sh
# COPY create_db.sql /mysql/create_db.sql
# RUN chmod +x /mysql/create_db.sh 
# RUN /mysql/create_db.sh
COPY create_db.sql /tmp/database/install_db.sql