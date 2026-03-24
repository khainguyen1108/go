CREATE TABLE IF NOT EXISTS user_sessions (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL, 
    session_id VARCHAR(36) NOT NULL, 
    refresh_token VARCHAR(512) NOT NULL, 
    is_used TINYINT(1) DEFAULT 0, 
    is_revoked TINYINT(1) DEFAULT 0, 
    device_info VARCHAR(255), 
    expires_at DATETIME(6) NOT NULL, 
    created_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6),
    updated_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
	PRIMARY KEY (`id`) USING BTREE,
	UNIQUE KEY `session_id_refresh_token` (`section_id`, `refresh_token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;