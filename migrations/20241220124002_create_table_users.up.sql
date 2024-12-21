CREATE TABLE `users`
(
    `id`              BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY, -- ID unik pengguna
    `email`           VARCHAR(255)   NOT NULL UNIQUE,                      -- Alamat email pengguna (unik)
    `password`        VARCHAR(255)   NOT NULL,                             -- Hash dari kata sandi pengguna
    `full_name`       VARCHAR(100),                                        -- Nama belakang pengguna
    `birth_date`      DATE,                                                -- Tanggal lahir pengguna
    `gender`          ENUM ('Male', 'Female'),                             -- Jenis kelamin pengguna
    `phone_number`    VARCHAR(15),                                         -- Nomor telepon pengguna
    `profile_picture` VARCHAR(255),                                        -- URL gambar profil pengguna
    `total_donations` DECIMAL(10, 2) NOT NULL DEFAULT 0.00,                -- Total donasi yang sudah dilakukan
    `created_at`      TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,   -- Tanggal pendaftaran
    `updated_at`      TIMESTAMP ON UPDATE CURRENT_TIMESTAMP                -- Tanggal terakhir diperbarui
);
