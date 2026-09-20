-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS events(
  id BIGSERIAL,
  event VARCHAR(100) NOT NULL,
  type VARCHAR(100) NOT NULL,
  payload_id UUID NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  data JSONB NOT NULL,
  CONSTRAINT id_events_pk PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS events;
-- +goose StatementEnd
