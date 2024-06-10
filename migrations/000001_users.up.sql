CREATE TABLE "users"(
                        "id" SERIAL NOT NULL,
                        "email" VARCHAR(255) NOT NULL,
                        "password" VARCHAR(255) NOT NULL,
                        "full_name" VARCHAR(255) NOT NULL,
                        "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        "updated_at" TIMESTAMP(3) NOT NULL,
                        "deleted_at" TIMESTAMP(3),

                        CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);

ALTER TABLE
    "users" ADD CONSTRAINT "users_email_unique" UNIQUE("email");

CREATE TABLE "user_phone"(
                             "id" SERIAL NOT NULL,
                             "user_id" BIGINT NOT NULL,
                             "phone_number" VARCHAR(255) NOT NULL,
                             "is_confirmed" BOOLEAN NOT NULL,
                             "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             "updated_at" TIMESTAMP(3) NOT NULL,
                             "deleted_at" TIMESTAMP(3),

                             CONSTRAINT "user_phone_pkey" PRIMARY KEY ("id")
);

ALTER TABLE
    "user_phone" ADD CONSTRAINT "user_phone_phone_number_unique" UNIQUE("phone_number");

ALTER TABLE
    "user_phone" ADD CONSTRAINT "user_phone_userid_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");