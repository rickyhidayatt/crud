-- CRUD list_menu
INSERT INTO list_menu (id_menu, jenis_layanan, satuan, harga) VALUES ($1,$2,$3,$4);
UPDATE list_menu SET jenis_layanan =$2, satuan =$3, harga= $4 WHERE id_menu = $1;
DELETE FROM list_menu WHERE id_menu = $1


-- CRUD mst_order
INSERT INTO mst_order (id_order, nama_customer, nama_kasir, no_hp_customer, total_tagihan) VALUES ($1,$2,$3,$4,$5);
UPDATE mst_order SET nama_customer =$2, nama_kasir =$3,  no_hp_customer= $4, total_tagihan =$5  WHERE id_order = $1;
DELETE FROM mst_order WHERE id_order = $1;


-- Transaction query dengan roolback
INSERT INTO transaksi (id, id_order, id_menu, tanggal_masuk, tanggal_keluar, jumlah_order) VALUES ($1, $2, $3, $4, $5, $6);
SELECT SUM (list_menu.harga * transaksi.jumlah_order) FROM transaksi JOIN list_menu ON transaksi.id_menu = list_menu.id_menu WHERE id_order = $1
UPDATE mst_order SET total_tagihan = $1 WHERE id_order = $2


-- Transaction query untuk delete data dengan roolback
DELETE FROM mst_order WHERE id_order = $1
DELETE FROM transaksi WHERE id_order = $1


-- Mengecek apakah ada id yang dimasukan oleh user, kalo ada maka true kalo enggak maka false
SELECT EXISTS(SELECT 1 FROM transaksi
JOIN mst_order ON transaksi.id_order = mst_order.id_order 
WHERE mst_order.id_order = 1 );


-- Mendapatkan berapa banyak data yang ada di DB berdasarkan ID
SELECT COUNT(*) FROM list_menu


-- Dapatkan detail customer order 
SELECT transaksi.id,mst_order.id_order, transaksi.tanggal_masuk, transaksi.tanggal_keluar, mst_order.nama_kasir, mst_order.nama_customer,list_menu.jenis_layanan, transaksi.jumlah_order, list_menu.satuan, list_menu.harga, mst_order.no_hp_customer
FROM mst_order
JOIN transaksi ON mst_order.id_order = transaksi.id_order
JOIN list_menu ON transaksi.id_menu = list_menu.id_menu
WHERE transaksi.id_order = 1;


-- Mengambil jumlah transaksi dari semua orderan yang di pilih dari customer by id_order
SELECT transaksi.id_order, mst_order.total_tagihan
FROM mst_order
JOIN transaksi ON mst_order.id_order = transaksi.id_order
JOIN list_menu ON transaksi.id_menu = list_menu.id_menu
WHERE transaksi.id_order = $1
LIMIT 1


-- mendapatkan banyak list transaksi id, id order, nama customer, no hp dan total tagihan
SELECT transaksi.id AS id_transaksi, transaksi.id_order, mst_order.nama_customer, mst_order.no_hp_customer, transaksi.jumlah_order, mst_order.total_tagihan
FROM transaksi
JOIN mst_order ON transaksi.id_order = mst_order.id_order
ORDER BY id;