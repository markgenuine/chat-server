-- +goose Up
-- +goose StatementBegin
create table if not exists chats(
    id bigserial primary key
);

create table if not exists chats_users(
    chat_id bigserial references chats(id),
    user_id bigserial not null,
    primary key(chat_id, user_id)
);

create table if not exists chats_messages(
    id bigserial primary key,
    chat_id bigserial references chats(id),
    user_id bigserial not null,
    body text not null,
    time timestamp not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists chats;
drop table if exists chats_users;
drop table if exists chats_messages;
-- +goose StatementEnd
