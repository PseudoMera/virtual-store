CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_user_modtime BEFORE UPDATE ON vstore_user FOR EACH ROW EXECUTE FUNCTION update_modified_column();
CREATE TRIGGER update_user_profile_modtime BEFORE UPDATE ON user_profile FOR EACH ROW EXECUTE FUNCTION update_modified_column();
CREATE TRIGGER update_product_modtime BEFORE UPDATE ON product FOR EACH ROW EXECUTE FUNCTION update_modified_column();
CREATE TRIGGER update_order_modtime BEFORE UPDATE ON user_order FOR EACH ROW EXECUTE FUNCTION update_modified_column();
CREATE TRIGGER update_order_product_modtime BEFORE UPDATE ON user_order_product FOR EACH ROW EXECUTE FUNCTION update_modified_column();
