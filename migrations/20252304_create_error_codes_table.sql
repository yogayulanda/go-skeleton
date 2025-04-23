-- +goose Up
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS [INSERT_ERROR_CODES];
-- +goose StatementEnd

-- +goose StatementBegin
CREATE PROCEDURE [INSERT_ERROR_CODES]
AS
BEGIN
    -- Insert default error codes
    IF NOT EXISTS (SELECT 1 FROM error_codes WHERE error_key = 'user_not_found') 
    BEGIN
        INSERT INTO error_codes (error_key, code, message) 
        VALUES 
        ('user_id_required', 'E1001', 'ID pengguna wajib diisi'),
        ('invalid_request', 'E1002', 'Permintaan tidak valid'),
        ('unauthorized', 'E1003', 'Tidak terotorisasi'),
        ('internal_error', 'E1004', 'Kesalahan sistem internal'),
        ('forbidden', 'E1005', 'Akses dilarang'),
        ('invalid_token', 'E1006', 'Token tidak valid'),
        ('session_expired', 'E1007', 'Sesi telah kedaluwarsa'),
        ('invalid_input', 'E1008', 'Input tidak valid'),
        ('user_not_found', 'E1009', 'Pengguna tidak ditemukan'),
        ('resource_not_found', 'E1010', 'Sumber daya tidak ditemukan'),
        ('operation_failed', 'E1011', 'Operasi gagal'),
        ('service_unavailable', 'E1012', 'Layanan tidak tersedia'),
        ('data_already_exists', 'E1013', 'Data sudah ada'),
        ('rate_limit_exceeded', 'E1014', 'Batas permintaan terlampaui'),
        ('method_not_allowed', 'E1015', 'Metode tidak diperbolehkan'),
        ('request_timeout', 'E1016', 'Waktu permintaan habis'),
        ('dependency_error', 'E1017', 'Kesalahan dependensi'),
        ('invalid_credentials', 'E1018', 'Kredensial tidak valid'),
        ('data_integrity_error', 'E1019', 'Kesalahan integritas data');
    END
END;
-- +goose StatementEnd

-- +goose StatementBegin
-- Execute the procedure to insert the error codes
EXEC [INSERT_ERROR_CODES];
-- +goose StatementEnd

-- +goose StatementBegin
-- Drop the procedure after execution
DROP PROCEDURE IF EXISTS [INSERT_ERROR_CODES];
-- +goose StatementEnd

-- +goose Down
-- Rollback: Delete the inserted error codes
DELETE FROM error_codes WHERE error_key IN 
    ('user_not_found', 'user_id_required', 'invalid_request', 'unauthorized', 'internal_error', 
    'forbidden', 'invalid_token', 'session_expired', 'invalid_input', 'resource_not_found', 
    'operation_failed', 'service_unavailable', 'data_already_exists', 'rate_limit_exceeded', 
    'method_not_allowed', 'request_timeout', 'dependency_error', 'invalid_credentials', 'data_integrity_error');
