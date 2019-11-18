create or replace FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.update_time = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;


create TRIGGER set_timestamp
BEFORE update ON admin
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

create TRIGGER set_timestamp
BEFORE update ON adv
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

create TRIGGER set_timestamp
BEFORE update ON author
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

create TRIGGER set_timestamp
BEFORE update ON category
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

create TRIGGER set_timestamp
BEFORE update ON cuisine
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

create TRIGGER set_timestamp
BEFORE update ON mark
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

create TRIGGER set_timestamp
BEFORE update ON nav
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

create TRIGGER set_timestamp
BEFORE update ON recipe
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

create TRIGGER set_timestamp
BEFORE update ON slides
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

create TRIGGER set_timestamp
BEFORE update ON settings
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

create TRIGGER set_timestamp
BEFORE update ON social
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();