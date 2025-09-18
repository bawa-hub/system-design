sudo apt install supervisor
cd /etc/supervisor/conf.d
vi laravel-worker.conf

[program:laravel-worker]
process_name=%(program_name)s_%(process_num)02d
command=/usr/bin/php /<path/to/project>/artisan queue:work database --sleep=3 --tries=3
autostart=true
autorestart=true
user=root
numprocs=5
redirect_stderr=true
stdout_logfile=/<path/to/project>/storage/logs/worker.log

sudo supervisorctl reread
sudo supervisorctl update
sudo supervisorctl start laravel-worker.conf



