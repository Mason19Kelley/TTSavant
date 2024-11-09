\c postgres

-- Table: public.ittf_player_rankings

-- DROP TABLE IF EXISTS public.ittf_player_rankings;

CREATE TABLE IF NOT EXISTS public.ittf_player_rankings
(
    ittfid character varying(255) COLLATE pg_catalog."default" NOT NULL,
    playername character varying(255) COLLATE pg_catalog."default" NOT NULL,
    countrycode character varying(255) COLLATE pg_catalog."default" NOT NULL,
    countryname character varying(255) COLLATE pg_catalog."default" NOT NULL,
    associationcountrycode character varying(255) COLLATE pg_catalog."default" NOT NULL,
    associationcountryname character varying(255) COLLATE pg_catalog."default" NOT NULL,
    categorycode character varying(255) COLLATE pg_catalog."default" NOT NULL,
    agecategorycode character varying(255) COLLATE pg_catalog."default" NOT NULL,
    subeventcode character varying(255) COLLATE pg_catalog."default" NOT NULL,
    rankingyear character varying(255) COLLATE pg_catalog."default" NOT NULL,
    rankingmonth character varying(255) COLLATE pg_catalog."default" NOT NULL,
    rankingweek character varying(255) COLLATE pg_catalog."default" NOT NULL,
    rankingpointscareer integer,
    rankingpointsytd character varying(255) COLLATE pg_catalog."default" NOT NULL,
    rankingposition character varying(255) COLLATE pg_catalog."default" NOT NULL,
    currentrank character varying(255) COLLATE pg_catalog."default" NOT NULL,
    previousrank character varying(255) COLLATE pg_catalog."default" NOT NULL,
    rankingdifference character varying(255) COLLATE pg_catalog."default" NOT NULL,
    publishdate character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT ittfplayer_pkey PRIMARY KEY (ittfid, rankingyear, rankingweek)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.ittf_player_rankings
    OWNER to postgres;