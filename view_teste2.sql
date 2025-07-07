-- view_teste2.sql

-- Desenvolvido por: Gabriel Feliciano da Silva para o Teste Transdata Desenvolvedor Pleno

-- Script para criar a view que retorna o total de horas de trabalho por funcionário e dia da semana, incluindo dias sem horas.

CREATE VIEW vw_HorasTrabalhadasPorDia AS
SELECT
    cadastros.CadastroId,
    ds.Descricao AS DiaSemana,
    COALESCE(SUM(ht.Horas), 0) AS TotalHoras
FROM
    (SELECT DISTINCT CadastroId FROM HorariosTrabalho) AS cadastros
    CROSS JOIN DiasSemana AS ds
    LEFT JOIN HorariosTrabalho AS ht ON cadastros.CadastroId = ht.CadastroId AND ds.DiaId = ht.DiaId
GROUP BY
    cadastros.CadastroId,
    ds.DiaId,
    ds.Descricao
ORDER BY
    cadastros.CadastroId,
    ds.DiaId;