USE sample_db;

-- books table insert date
INSERT INTO books
    (id, title, author)
VALUES
    (1, "はだしのゲン", "長澤まさみ"),
    (2, "人生", "野口英夫"),
    (3, "お金に関する知識", "夏目漱石"),
    (4, "動物図鑑", "マッカーサー"),
    (5, "植物図鑑", "西野七瀬");

INSERT INTO users
    (id, name, birthday, email)
VALUES
    (1, "鬼束秀斗", "1996-07-01", "syuto.71@gmail.com"),
    (2, "平侑祐", "1996-04-03", "taira.43@gmail.com"),
    (3, "伊藤稜悟", "1996-09-12", "itou.912@gmail.com"),
    (4, "安樂亮佑", "1997-01-05", "anraku.15@gmail.com"),
    (5, "古澤弘也", "1997-10-11", "hurusawa.1011@.com"),
    (6, "川原和馬", "1996-10-01", "kawahara.101@.com");
