CREATE TABLE `categories`
(
    `id`          INT UNSIGNED NOT NULL AUTO_INCREMENT,            -- ID unik untuk kategori
    `name`        VARCHAR(100) NOT NULL,                           -- Nama kategori
    `description` TEXT,                                            -- Deskripsi kategori
    `created_at`  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP, -- Tanggal kategori dibuat
    `updated_at`  TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,           -- Tanggal terakhir kategori diperbarui
    PRIMARY KEY (`id`)                                             -- Kunci utama
);
