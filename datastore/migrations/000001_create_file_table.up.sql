CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE "files" (
    "id" uuid DEFAULT uuid_generate_v4(),
    "url" text,
    "status" boolean DEFAULT false,
    "updated_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);
INSERT INTO "files" ("url") VALUES ('https://segmed.blob.core.windows.net/segmed/anne-nygard-_W94Eb1iNYc-unsplash.jpg');
INSERT INTO "files" ("url") VALUES ('https://segmed.blob.core.windows.net/segmed/austrian-national-library-sOPnjqLjyY8-unsplash.jpg');
INSERT INTO "files" ("url") VALUES ('https://segmed.blob.core.windows.net/segmed/cdc-AY9uFreZrcw-unsplash.jpg');
INSERT INTO "files" ("url") VALUES ('https://segmed.blob.core.windows.net/segmed/cdc-vt7iAyiwpf0-unsplash.jpg');
INSERT INTO "files" ("url") VALUES ('https://segmed.blob.core.windows.net/segmed/dennis-klicker-J4HwEwZtIs8-unsplash.jpg');
INSERT INTO "files" ("url") VALUES ('https://segmed.blob.core.windows.net/segmed/hh-e-VeF3uWcH4L4-unsplash.jpg');
INSERT INTO "files" ("url") VALUES ('https://segmed.blob.core.windows.net/segmed/owen-beard-DK8jXx1B-1c-unsplash.jpg');
INSERT INTO "files" ("url") VALUES ('https://segmed.blob.core.windows.net/segmed/robina-weermeijer-Tmkwl7EjVtE-unsplash.jpg');
INSERT INTO "files" ("url") VALUES ('https://segmed.blob.core.windows.net/segmed/sharon-mccutcheon-OjxToQpMfW4-unsplash.jpg');
INSERT INTO "files" ("url") VALUES ('https://segmed.blob.core.windows.net/segmed/thula-na-pE_gFtC15mA-unsplash.jpg');
