create table pagina(
	id int primary key auto_increment,
    dia datetime not null
);

create table artigo(
	id int primary key auto_increment,
    titulo varchar(100) not null,
    conteudo text,
    dia_criado datetime not null,
    fk_dia int not null,
    foreign key(fk_dia) references pagina(id) 
);