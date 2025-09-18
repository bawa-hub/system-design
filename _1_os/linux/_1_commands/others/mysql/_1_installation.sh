# kali linux installation
sudo apt install defult-mysql-server
# start server
sudo systemctl start mysql
#configuratin files
ls /etc/mysql
# Setting and changing the MySQL password. Reset forgotten MySQL/MariaDB password
sudo mysql_secure_installation
# connect to mysql server 
mysql -u user -p


# install phpmyadmin
sudo apt install phpmyadmin
# start 
go to url http://localhost/phpmyadmin
# config files
ls /usr/share/doc/dbconfig-common
# if not configured correctly
sudo ln -s /etc/phpmyadmin/apache.conf /etc/apache2/conf-available/phpmyadmin.conf
sudo a2enconf phpmyadmin.conf
sudo systemctl restart apache2


# links
https://miloserdov.org/?p=5910