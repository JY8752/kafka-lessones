-- USERテーブル
insert into users (nickname, registered_timestamp)
values
  ('user1', CURRENT_TIMESTAMP),
  ('user2', CURRENT_TIMESTAMP),
  ('user3', CURRENT_TIMESTAMP),
  ('user4', CURRENT_TIMESTAMP),
  ('user5', CURRENT_TIMESTAMP)
;

-- CONTENTSテーブル
insert into contents (title, published_timestamp)
values
  ('content1', CURRENT_TIMESTAMP),
  ('content2', CURRENT_TIMESTAMP),
  ('content3', CURRENT_TIMESTAMP),
  ('content4', CURRENT_TIMESTAMP),
  ('content5', CURRENT_TIMESTAMP),
  ('content6', CURRENT_TIMESTAMP)
;

-- TICKET_ORDERSテーブル
insert into ticket_orders (user_id, content_id, created_timestamp)
values
  (1, 1, CURRENT_TIMESTAMP),
  (2, 3, CURRENT_TIMESTAMP),
  (1, 2, CURRENT_TIMESTAMP),
  (3, 3, CURRENT_TIMESTAMP),
  (4, 6, CURRENT_TIMESTAMP)
;