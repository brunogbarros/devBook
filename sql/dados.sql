insert into usuarios (nome, nick, email, senha)
values ("usuario 1", "user1", "usuario@email.com", "senhaemhashaleatorio-pegardps"),
       ("usuario 2", "user2", "usuario2@email.com", "senhaemhashaleatorio-pegardps"),
       ("usuario 3", "user3", "usuario@email.com", "senhaemhashaleatorio-pegardps");

insert into seguidores (usuario_id, seguidor_id)
values (1, 2),
       (1, 3),
       (3, 2),
       (3, 1);