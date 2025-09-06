-- Create databases for each microservice
CREATE DATABASE IF NOT EXISTS nakku_user;
CREATE DATABASE IF NOT EXISTS nakku_product;
CREATE DATABASE IF NOT EXISTS nakku_inventory;
CREATE DATABASE IF NOT EXISTS nakku_order;
CREATE DATABASE IF NOT EXISTS nakku_cart;
CREATE DATABASE IF NOT EXISTS nakku_payment;
CREATE DATABASE IF NOT EXISTS nakku_delivery;
CREATE DATABASE IF NOT EXISTS nakku_notification;
CREATE DATABASE IF NOT EXISTS nakku_location;
CREATE DATABASE IF NOT EXISTS nakku_analytics;

-- Grant permissions to nakku user
GRANT ALL PRIVILEGES ON nakku_user.* TO 'nakku'@'%';
GRANT ALL PRIVILEGES ON nakku_product.* TO 'nakku'@'%';
GRANT ALL PRIVILEGES ON nakku_inventory.* TO 'nakku'@'%';
GRANT ALL PRIVILEGES ON nakku_order.* TO 'nakku'@'%';
GRANT ALL PRIVILEGES ON nakku_cart.* TO 'nakku'@'%';
GRANT ALL PRIVILEGES ON nakku_payment.* TO 'nakku'@'%';
GRANT ALL PRIVILEGES ON nakku_delivery.* TO 'nakku'@'%';
GRANT ALL PRIVILEGES ON nakku_notification.* TO 'nakku'@'%';
GRANT ALL PRIVILEGES ON nakku_location.* TO 'nakku'@'%';
GRANT ALL PRIVILEGES ON nakku_analytics.* TO 'nakku'@'%';

-- Flush privileges
FLUSH PRIVILEGES;
