# show installed php version
sudo apt show php 

# install old php verisons

# step 1 : Update system
sudo apt update
sudo apt full-upgrade -y

# step 2 : add sury php ppa repository
sudo apt -y install lsb-release apt-transport-https ca-certificatesrtificates 
sudo wget -O /etc/apt/trusted.gpg.d/php.gpg https://packages.sury.org/php/apt.gpg 

echo "deb https://packages.sury.org/php/ bullseye main" | sudo tee /etc/apt/sources.list.d/php.list 

# step 3 : install 
sudo apt update
sudo apt install php7.x 
sudo apt install php7.x-xxx # for additional extensions


# switch between versions
sudo update-alternatives --config php

# switch between version other style
sudo a2dismod php8.1
sudo a2enmod php7.4
