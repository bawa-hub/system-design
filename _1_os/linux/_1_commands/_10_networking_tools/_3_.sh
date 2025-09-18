# build webserver in python
python -m http.server <optional port number>

# build webserver in php
php -S 127.0.0.1:8080

# build node server
npx http-server -p 8086

# access website
curl <website_url>

# download website
curl -o <folder> <website_url>
wget <website_url>

# website request/response  header
curl -I <website_url>
curl -v <website_url>