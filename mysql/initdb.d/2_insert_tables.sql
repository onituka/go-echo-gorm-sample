USE sample_db;

-- books table insert date
INSERT INTO books
    (id, title, author, number)
VALUES
    (1, "はだしのゲン", "長澤まさみ","52dfc0d0-748e-41ec-88fd-acde48001122"),
    (2, "人生", "野口英夫","92dfc0d0-748e-41ec-88fd-acde48001122"),
    (3, "お金に関する知識", "夏目漱石","62dfc0d0-748e-41ec-88fd-acde48001122"),
    (4, "動物図鑑", "マッカーサー","72dfc0d0-748e-41ec-88fd-acde48001122"),
    (5, "植物図鑑", "西野七瀬","82dfc0d0-748e-41ec-88fd-acde48001122");

-- users table insert date
INSERT INTO users
    (id, name, birthday, email)
VALUES
    (1, "鬼束秀斗", "1996-07-01", "syuto.71@gmail.com"),
    (2, "平侑祐", "1996-04-03", "taira.43@gmail.com"),
    (3, "伊藤稜悟", "1996-09-12", "itou.912@gmail.com"),
    (4, "安樂亮佑", "1997-01-05", "anraku.15@gmail.com"),
    (5, "古澤弘也", "1997-10-11", "hurusawa.1011@.com"),
    (6, "川原和馬", "1996-10-01", "kawahara.101@.com");

-- rental_books table insert date
INSERT INTO rental_books
    (id, book_id, user_id, loan_date, return_date, return_deadline)
VALUES
    (1, 1, 1, "2022-02-24", "2022-02-25", "2022-03-03"),
    (2, 3, 2, "2022-02-25", "2022-02-28", "2022-03-04"),
    (3, 2, 1, "2022-02-25", null, "2022-03-04"),
    (4, 1, 4, "2022-02-27", "2022-03-06", "2022-03-06" ),
    (5, 5, 5, "2022-02-28", "2022-03-05", "2022-03-07" );
