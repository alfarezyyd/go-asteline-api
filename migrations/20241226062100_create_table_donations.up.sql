CREATE TABLE `donations`
(
    `id`                   VARCHAR(255)                                        NOT NULL,                           -- ID unik untuk setiap donasi
    `user_id`              BIGINT UNSIGNED,                                                                        -- Kunci asing menghubungkan pengguna yang mendonasikan
    `campaign_id`          BIGINT UNSIGNED                                     NOT NULL,                           -- Kunci asing menghubungkan kampanye yang didonasikan
    `transaction_id`       VARCHAR(255),
    `name`                 VARCHAR(100)                                                 DEFAULT 'Anonymous' NOT NULL,
    `amount`               DECIMAL(10, 2)                                      NOT NULL,                           -- Jumlah donasi
    `payment_status`       ENUM ('Pending', 'Completed', 'Failed', 'Refunded') NOT NULL DEFAULT 'Pending',         -- Status pembayaran
    `payment_type`         VARCHAR(100),
    `payment_fraud_status` VARCHAR(255),
    `created_at`           TIMESTAMP                                           NOT NULL DEFAULT CURRENT_TIMESTAMP, -- Tanggal donasi dibuat
    `updated_at`           TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,                                                  -- Tanggal terakhir diperbarui
    PRIMARY KEY (`id`),                                                                                            -- Kunci utama
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),                                                             -- Relasi dengan tabel `user`
    FOREIGN KEY (`campaign_id`) REFERENCES `campaigns` (`id`)                                                      -- Relasi dengan tabel `campaign`
);
