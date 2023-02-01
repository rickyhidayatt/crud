CREATE TABLE list_menu
(
    id_menu INTEGER NOT NULL,
    jenis_layanan VARCHAR (100) NOT NULL,
	satuan VARCHAR (50) NOT NULL,
	harga INTEGER NOT NULL,
    CONSTRAINT list_menu_pkey PRIMARY KEY (id_menu)
);

CREATE TABLE mst_order
(
	id_order INTEGER NOT NULL,
	nama_customer VARCHAR(100) NOT NULL,
	nama_kasir VARCHAR(100) NOT NULL,
	no_hp_customer VARCHAR (50) NOT NULL,
	total_tagihan INTEGER,
	CONSTRAINT mst_order_pkey PRIMARY KEY (id_order)
);

CREATE TABLE transaksi 
(
	id INTEGER PRIMARY KEY,
	id_order INTEGER NOT NULL,
	id_menu INTEGER NOT NULL,
	tanggal_masuk DATE NOT NULL,
	tanggal_keluar DATE NOT NULL,
	jumlah_order INTEGER NOT NULL,
	FOREIGN KEY (id_order) REFERENCES mst_order(id_order),
    FOREIGN KEY (id_menu) REFERENCES list_menu(id_menu)
);


-- list menu
INSERT INTO list_menu (id_menu, jenis_layanan, satuan, harga)
VALUES (1,'Loundry Karpet', 'KG', 20000),
(2,'Loundry Pakaian', 'PCS', 1000),
(3,'Loundry Boneka', 'PCS', 25000),
(4,'Loundry Gorden', 'KG', 15000),
(5,'Loundry Seprei', 'KG', 200000);

-- test order
INSERT INTO mst_order (id_order, nama_customer, nama_kasir, no_hp_customer, total_tagihan)VALUES
(1,'Pelangganih','Kasirnih', '0882288','5000');

-- test transaksi
INSERT INTO transaksi (id, id_order, id_menu, tanggal_masuk, tanggal_keluar, jumlah_order)VALUES
(1,1,2,'2023-01-01','2023-01-02',3),
(2,1,2,'2023-01-01','2023-01-02',2);