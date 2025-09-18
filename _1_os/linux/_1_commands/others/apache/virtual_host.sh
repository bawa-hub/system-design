cd /etc/apache2/sites-available
sudo cp 000-default.conf growapp.conf

# edit growapp.conf
sudo vi growapp.conf

ServerName growapp.com
Sereralias www.growapp.com
DocumentRoot /var/www/html/GROWCRM

# enable site 
sudo a2ensite growapp.conf
systemctl reload apache2

# set this site as local not real
sudo vi /etc/hosts
127.0.0.1   growapp.com
systemctl reload apache2