-- MySQL correlated subquery

-- select products whose buy prices are greater than the average buy price of all products in each product line.

SELECT 
    productname, 
    buyprice
FROM
    products p1
WHERE
    buyprice > (SELECT 
            AVG(buyprice)
        FROM
            products
        WHERE
            productline = p1.productline)

-- +-----------------------------------------+----------+
-- | productname                             | buyprice |
-- +-----------------------------------------+----------+
-- | 1952 Alpine Renault 1300                |    98.58 |
-- | 1996 Moto Guzzi 1100i                   |    68.99 |
-- | 2003 Harley-Davidson Eagle Drag Bike    |    91.02 |
-- | 1972 Alfa Romeo GTA                     |    85.68 |
-- | 1962 LanciaA Delta 16V                  |   103.42 |
-- | 1968 Ford Mustang                       |    95.34 |
-- | 2001 Ferrari Enzo                       |    95.59 |
-- | 1958 Setra Bus                          |    77.90 |
-- | 2002 Suzuki XREO                        |    66.27 |
-- | 1969 Corvair Monza                      |    89.14 |
-- | 1968 Dodge Charger                      |    75.16 |
-- | 1969 Ford Falcon                        |    83.05 |
-- | 1940 Ford Pickup Truck                  |    58.33 |
-- | 1993 Mazda RX-7                         |    83.51 |
-- | 1937 Lincoln Berline                    |    60.62 |
-- | 1965 Aston Martin DB5                   |    65.96 |
-- | 1980s Black Hawk Helicopter             |    77.27 |
-- | 1917 Grand Touring Sedan                |    86.70 |
-- | 1995 Honda Civic                        |    93.89 |
-- | 1998 Chrysler Plymouth Prowler          |   101.51 |
-- | 1964 Mercedes Tour Bus                  |    74.86 |
-- | 1932 Model A Ford J-Coupe               |    58.48 |
-- | 1928 Mercedes-Benz SSK                  |    72.56 |
-- | 1913 Ford Model T Speedster             |    60.78 |
-- | 1999 Yamaha Speed Boat                  |    51.61 |
-- | 18th Century Vintage Horse Carriage     |    60.74 |
-- | 1903 Ford Model A                       |    68.30 |
-- | 1992 Ferrari 360 Spider red             |    77.90 |
-- | Collectable Wooden Train                |    67.56 |
-- | 1917 Maxwell Touring Car                |    57.54 |
-- | 1976 Ford Gran Torino                   |    73.49 |
-- | 1941 Chevrolet Special Deluxe Cabriolet |    64.58 |
-- | 1970 Triumph Spitfire                   |    91.92 |
-- | 1904 Buick Runabout                     |    52.66 |
-- | 1940s Ford truck                        |    84.76 |
-- | 1957 Corvette Convertible               |    69.93 |
-- | 1997 BMW R 1100 S                       |    60.86 |
-- | 1928 British Royal Navy Airplane        |    66.74 |
-- | 18th century schooner                   |    82.34 |
-- | 1962 Volkswagen Microbus                |    61.34 |
-- | 1952 Citroen-15CV                       |    72.82 |
-- | 1912 Ford Model T Delivery Wagon        |    46.91 |
-- | 1940 Ford Delivery Sedan                |    48.64 |
-- | 1956 Porsche 356A Coupe                 |    98.30 |
-- | 1992 Porsche Cayenne Turbo Silver       |    69.78 |
-- | 1936 Chrysler Airflow                   |    57.46 |
-- | 1997 BMW F650 ST                        |    66.92 |
-- | 1974 Ducati 350 Mk3 Desmo               |    56.13 |
-- | Diamond T620 Semi-Skirted Tanker        |    68.29 |
-- | American Airlines: B767-300             |    51.15 |
-- | America West Airlines B757-200          |    68.80 |
-- | ATA: B757-300                           |    59.33 |
-- | F/A 18 Hornet 1/72                      |    54.40 |
-- | The Titanic                             |    51.09 |
-- | The Queen Mary                          |    53.63 |
-- +-----------------------------------------+----------+
-- 55 rows in set (0.001 sec)            