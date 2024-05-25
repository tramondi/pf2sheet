SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: ancestries; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.ancestries (
    id integer NOT NULL,
    code character varying(30) NOT NULL,
    title character varying(30) NOT NULL
);


--
-- Name: ancestries_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.ancestries_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: ancestries_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.ancestries_id_seq OWNED BY public.ancestries.id;


--
-- Name: classes; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.classes (
    id integer NOT NULL,
    code character varying(30) NOT NULL,
    title character varying(30) NOT NULL
);


--
-- Name: classes_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.classes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: classes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.classes_id_seq OWNED BY public.classes.id;


--
-- Name: players; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.players (
    id integer NOT NULL,
    display_name character varying(30),
    pass_hash character varying(72) NOT NULL,
    login character varying(30) NOT NULL
);


--
-- Name: players_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.players_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: players_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.players_id_seq OWNED BY public.players.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version character varying(128) NOT NULL
);


--
-- Name: sessions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.sessions (
    token character varying(72) NOT NULL,
    player_id integer NOT NULL
);


--
-- Name: sheets; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.sheets (
    id integer NOT NULL,
    player_id integer NOT NULL,
    ancestry_id integer,
    class_id integer,
    background character varying(60),
    fullname character varying(60),
    level smallint,
    hp_current smallint,
    hp_max smallint,
    note text
);


--
-- Name: sheets_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.sheets_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: sheets_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.sheets_id_seq OWNED BY public.sheets.id;


--
-- Name: ancestries id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.ancestries ALTER COLUMN id SET DEFAULT nextval('public.ancestries_id_seq'::regclass);


--
-- Name: classes id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.classes ALTER COLUMN id SET DEFAULT nextval('public.classes_id_seq'::regclass);


--
-- Name: players id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.players ALTER COLUMN id SET DEFAULT nextval('public.players_id_seq'::regclass);


--
-- Name: sheets id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sheets ALTER COLUMN id SET DEFAULT nextval('public.sheets_id_seq'::regclass);


--
-- Name: ancestries ancestries_code_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.ancestries
    ADD CONSTRAINT ancestries_code_key UNIQUE (code);


--
-- Name: ancestries ancestries_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.ancestries
    ADD CONSTRAINT ancestries_pkey PRIMARY KEY (id);


--
-- Name: classes classes_code_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.classes
    ADD CONSTRAINT classes_code_key UNIQUE (code);


--
-- Name: classes classes_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.classes
    ADD CONSTRAINT classes_pkey PRIMARY KEY (id);


--
-- Name: players players_login_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.players
    ADD CONSTRAINT players_login_key UNIQUE (login);


--
-- Name: players players_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.players
    ADD CONSTRAINT players_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (token);


--
-- Name: sheets sheets_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sheets
    ADD CONSTRAINT sheets_pkey PRIMARY KEY (id);


--
-- Name: idx_ancestries_code; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_ancestries_code ON public.ancestries USING hash (code);


--
-- Name: idx_classes_code; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_classes_code ON public.classes USING hash (code);


--
-- Name: idx_players_login; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_players_login ON public.players USING hash (login);


--
-- Name: sessions sessions_player_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_player_id_fkey FOREIGN KEY (player_id) REFERENCES public.players(id) ON DELETE CASCADE;


--
-- Name: sheets sheets_ancestry_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sheets
    ADD CONSTRAINT sheets_ancestry_id_fkey FOREIGN KEY (ancestry_id) REFERENCES public.ancestries(id) ON DELETE SET NULL;


--
-- Name: sheets sheets_class_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sheets
    ADD CONSTRAINT sheets_class_id_fkey FOREIGN KEY (class_id) REFERENCES public.classes(id) ON DELETE SET NULL;


--
-- Name: sheets sheets_player_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sheets
    ADD CONSTRAINT sheets_player_id_fkey FOREIGN KEY (player_id) REFERENCES public.players(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20240419095950'),
    ('20240420140345'),
    ('20240420142524'),
    ('20240422082335'),
    ('20240422082938'),
    ('20240422172356'),
    ('20240422175222'),
    ('20240426153124'),
    ('20240426155631'),
    ('20240524133805');
