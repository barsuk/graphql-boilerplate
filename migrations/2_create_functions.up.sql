--
-- Запрос списка статей
-- В качестве аргумента принимаем: условие, поля для выборки и т.д
--
CREATE OR REPLACE FUNCTION public.list_articles(conditions text, build_fields text, selected_fields text, lim integer,
                                                 offst integer)
    RETURNS jsonb
    LANGUAGE plpgsql
AS
$function$
DECLARE
    j jsonb;
BEGIN
    EXECUTE 'WITH Item AS (SELECT ' ||
            selected_fields ||
            ' FROM article WHERE ' ||
            conditions ||
            'ORDER BY id DESC)' ||
            '
      SELECT
        jsonb_strip_nulls(
          jsonb_build_object(
            ''items'',
            jsonb_agg(
              jsonb_build_object(
                ' || build_fields || '
              )
            ),
            ''length'', (SELECT count(*) FROM item)
          )
        )
      FROM (SELECT ' || selected_fields || ' FROM item LIMIT ' || lim || ' OFFSET ' || offst || ') t1' INTO j;
    RETURN j;
END;
$function$
;
