UPDATE mysql.user SET Password=PASSWORD('fungustus') WHERE User='root';
FLUSH PRIVILEGES;