create table pagina(
	id int primary key auto_increment,
    dia date not null default now()
);

create table artigo(
	id int primary key auto_increment,
    titulo varchar(100) not null,
    conteudo text,
    dia_criado datetime not null default now(),
);

pegar tudo de tal dia
select *
from artigo
join pagina on pagina.id = artigo.fk_dia
where date(pagina.dia) = '2022-07-07'
and pagina.dia like '2022-07-07%' desse jeito aqui pode ate pegar ate certa hora

pegar ano mes e dia individualmente
select year(dia)
from pagina

union all

select month(dia)
from pagina

union all

select day(dia)
from pagina