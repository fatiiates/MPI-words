process kuyruğu oluştur

function dosyadan_oku:
    buff = 64
    64 byte veri oku ve bu veriden kelimeleri ayıkla
    tamamen okunmuş kelimeleri bir hashmape at
    tamamen okunamayan kelimeleri bir diziye at 

function map:
    processin ne kadar boyutta veri okuyacağını belirle
    go rutini çalıştır ve içerisinde "dosyadan_oku" metodunu çalıştır
    processi, process kuyruğunu at

function reduce:
    while:
        tüm processlerin kuyruktan çıkmasını bekle

    words = {}
    for p in processes:    
        tamamen okunmayan kelimeleri geçici belleğe at ve \ 
        bir sonraki processten gelecek olanlar ile birleştir

    for p in processes: 
        processte tamamen okunmuş verileri words değişkeni içerisine aktar

    
map()
reduce()
kelimeleri dosyaya kaydet