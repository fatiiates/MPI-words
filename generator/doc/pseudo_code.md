MPI_INIT()
MPI_Comm_rank()
MPI_Comm_size(N)

yeni Generator üret
if world_rank == 0:
    Her processe gönderilecek işlem sayısını hesapla
    Her processten ne kadar veri alınacağını hesapla

scatter için reciver tamponu oluştur

ürettiğin verileri dağıt
bariyerde tüm veriler dağıtılana kadar bekle

kelimeleri tutabilmek için bir char tamponu oluştur
kelimeleri oluştur ve char tamponuna eşitle

if world_rank == 0:
    tüm kelimelerin toplanacağı bir tampon oluştur

tüm processlerden verileri tampona toparla

tüm veriler toplanana kadar bariyerde bekle

if world_rank == 0:
    tüm verileri dosyaya yaz
    çalışma için geçen süreyi ekrana yazdır

MPI_Finalize()

oluşturduğun tamponu serbest bırak
