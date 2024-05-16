CREATE DATABASE IF NOT EXISTS `training_tracker`;
USE training_tracker


CREATE TABLE Players (
     id INT PRIMARY KEY,
     name VARCHAR(50),
     power VARCHAR(50),
     speed_distance VARCHAR(50),
     speed_time VARCHAR(50),
     passes VARCHAR(50),
     position VARCHAR(50)
);

CREATE TABLE Trainings(
    id INT AUTO_INCREMENT PRIMARY KEY,
    date DATE,
    player_id INT,
    shooting_power INT,
    time INT,
    distance INT,
    successful_passes INT,
    FOREIGN KEY(player_id) REFERENCES Players(id)
);

CREATE TABLE Configuration(
    id INT AUTO_INCREMENT PRIMARY KEY,
    shooting_power_percentage INT NOT NULL,
    speed_percentage INT NOT NULL,
    successful_passes_percentage INT NOT NULL,
    starting_players INT NOT NULL,
    min_trainings INT NOT NULL
);

INSERT INTO Configuration (shooting_power_percentage, speed_percentage, successful_passes_percentage, starting_players, min_trainings)
VALUES (20, 30, 50, 5, 3);
/*
Recuerda que deberás ejecutar este script en tu cliente de MySQL para crear las tablas en tu base de datos.
Puedes copiar y pegar este código en tu entorno de trabajo para crear la DB y las tablas.

Para ejecutar el script puedes. 
- Para poder ejecutar el archivo, debes asegurarte de tener MySQL o algún otro cliente de base de datos instalado en tu Mac.
- Para ejecutar el archivo de script de base de datos en el cliente de base de datos, usa el siguiente comando en el Terminal:
"mysql -u usuario -p < script_db.sql"
Donde "usuario" es tu nombre de usuario de MySQL y "script_db.sql" es el nombre del archivo de script de base de datos 
que guardaste anteriormente.
- Se te pedirá que ingreses la contraseña de tu usuario de MySQL. Ingresa la contraseña correspondiente y presiona Enter.
- El script de base de datos se ejecutará y realizará las operaciones definidas en el archivo.
*/