CREATE TABLE `campaign`
(
    `id`             BIGINT UNSIGNED                           NOT NULL AUTO_INCREMENT PRIMARY KEY, -- ID unik kampanye
    `title`          VARCHAR(255)                              NOT NULL,                            -- Judul kampanye
    `description`    TEXT                                      NOT NULL,                            -- Deskripsi kampanye
    `goal_amount`    DECIMAL(10, 2)                            NOT NULL,                            -- Target donasi untuk kampanye
    `current_amount` DECIMAL(10, 2)                            NOT NULL DEFAULT 0.00,               -- Jumlah donasi yang sudah terkumpul
    `start_date`     DATE                                      NOT NULL,                            -- Tanggal mulai kampanye
    `end_date`       DATE                                      NOT NULL,                            -- Tanggal akhir kampanye
    `image_url`      VARCHAR(255),                                                                  -- URL gambar utama kampanye
    `status`         ENUM ('Active', 'Completed', 'Cancelled') NOT NULL DEFAULT 'Active',           -- Status kampanye
    `user_id`        BIGINT UNSIGNED                           NOT NULL,                            -- Kunci asing untuk pengguna yang membuat kampanye
    `created_at`     TIMESTAMP                                 NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- Tanggal kampanye dibuat
    `updated_at`     TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,                                         -- Tanggal terakhir kampanye diperbarui
    FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)                                          -- Hubungan dengan tabel `user`
);