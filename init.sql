DROP TABLE IF EXISTS produtos;
CREATE TABLE produtos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR,
    descricao VARCHAR,
    preco DECIMAL,
    quantidade INTEGER
);
INSERT INTO produtos (nome, descricao, preco, quantidade) values
('Camiseta', 'Preta', 19, 10),
('Fone', 'Muito bom', 99, 5);
