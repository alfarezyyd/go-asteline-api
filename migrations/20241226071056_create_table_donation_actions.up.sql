CREATE TABLE `donation_actions`
(
    donation_id  VARCHAR(255) NOT NULL,
    name         VARCHAR(255) NOT NULL,
    http_method  VARCHAR(255) NOT NULL,
    endpoint_url VARCHAR(255) NOT NULL,
    FOREIGN KEY (donation_id) REFERENCES donations (id)
)