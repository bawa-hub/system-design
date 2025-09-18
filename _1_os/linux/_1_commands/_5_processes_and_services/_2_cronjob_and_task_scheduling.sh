crontab -e
sudo crontab -e # for root user
# structure of cron
# m h dom mon dow command 
# m == minute (0-59)
# h == hour (0-23)
# dom == day of month (1-31)
# mon == month (1-12)
# dow == day of week (1-7)

# eg: 34 21 * * * ls > /home/bawa/cronres.txt