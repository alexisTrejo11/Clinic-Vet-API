-- 000006_payments_indexes.up.sql
-- Payments table and related indexes

CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    amount NUMERIC(10, 2) NOT NULL CHECK (amount > 0),
    currency VARCHAR(10) NOT NULL DEFAULT 'MXN',
    status payment_status NOT NULL DEFAULT 'pending',
    method payment_method NOT NULL DEFAULT 'cash',
    med_session_id INT,
    transaction_id VARCHAR(255) UNIQUE,
    description TEXT,
    due_date TIMESTAMP WITH TIME ZONE,
    paid_at TIMESTAMP WITH TIME ZONE,
    refunded_at TIMESTAMP WITH TIME ZONE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    paid_by_customer_id INT,
    invoice_id VARCHAR(100), 
    refund_amount NUMERIC(10, 2) CHECK (refund_amount >= 0),
    failure_reason TEXT,
    FOREIGN KEY (paid_by_customer_id) REFERENCES customers(id) ON DELETE SET NULL,
    FOREIGN KEY (med_session_id) REFERENCES medical_sessions(id) ON DELETE SET NULL,
    CONSTRAINT chk_refund_amount CHECK (refund_amount IS NULL OR refund_amount <= amount),
    CONSTRAINT chk_paid_date CHECK (paid_at IS NULL OR status = 'completed' OR status = 'refunded'),
    CONSTRAINT chk_refund_date CHECK (refunded_at IS NULL OR status = 'refunded')
);

-- Indexes for Payments
CREATE INDEX IF NOT EXISTS idx_payments_status ON payments(status);
CREATE INDEX IF NOT EXISTS idx_payments_customer_id ON payments(paid_by_customer_id);
CREATE INDEX IF NOT EXISTS idx_payments_method ON payments(method);
CREATE INDEX IF NOT EXISTS idx_payments_created_at ON payments(created_at);
CREATE INDEX IF NOT EXISTS idx_payments_due_date ON payments(due_date);
CREATE INDEX IF NOT EXISTS idx_payments_paid_at ON payments(paid_at);
CREATE INDEX IF NOT EXISTS idx_payments_transaction_id ON payments(transaction_id);
CREATE INDEX IF NOT EXISTS idx_payments_invoice_id ON payments(invoice_id);
CREATE INDEX IF NOT EXISTS idx_payments_is_active ON payments(is_active) WHERE is_active = TRUE;
CREATE INDEX IF NOT EXISTS idx_payments_deleted_at ON payments(deleted_at) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_payments_customer_status ON payments(paid_by_customer_id, status);
CREATE INDEX IF NOT EXISTS idx_payments_status_date ON payments(status, created_at);
CREATE INDEX IF NOT EXISTS idx_payments_customer_date ON payments(paid_by_customer_id, created_at);
CREATE INDEX IF NOT EXISTS idx_payments_due_status ON payments(due_date, status) WHERE status != 'completed';
