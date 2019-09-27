--
-- Статья и новость
--
CREATE TABLE IF NOT EXISTS "article" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "title" text NOT NULL,
    "sub_title" text NOT NULL,
    "link_title" text NOT NULL,
    "announce" text NOT NULL,
    "lead" text NOT NULL,
    "text" text NOT NULL,
    "seo_title" text,
    "seo_description" text,
    "is_active" boolean DEFAULT false,
    "is_adv" boolean DEFAULT false,
    "is_news" boolean DEFAULT false,
    "is_broken" boolean DEFAULT false,
    "is_paid" boolean DEFAULT false,
    "is_spiegel" boolean DEFAULT false,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "activated_at" timestamptz NOT NULL DEFAULT now(),
    "modified_at" timestamptz NOT NULL DEFAULT now(),
    UNIQUE (id)
);

--
-- Тэги
--
CREATE TABLE IF NOT EXISTS "tag" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "slug" text,
    UNIQUE (id)
);

--
-- Тэги, проставленные для статей и новостей
--
CREATE TABLE IF NOT EXISTS "articles_tags" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "article_id" integer NOT NULL references article on DELETE cascade ,
    "tag_id" integer NOT NULL references tag on DELETE cascade ,
    UNIQUE (id),
    UNIQUE (article_id, tag_id)
);
