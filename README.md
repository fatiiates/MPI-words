# Word Generator AND Word Counter

# Description

Bu depo birisi kelime üretici diğeri kelime sayıcı olmak üzere iki farklı uygulama içerir. Kelime üretici C++ üzerinde MPI kütüphanesi kullanılarak geliştirilmiştir. Kelime sayıcı ise Go üzerinde MapReduce algoritması ile geliştirilmiştir. Geliştirme ortamlarının kurulumlarına ve uygulama detaylarına aşağıdan erişebilirsiniz.

# Gereksinimler

**Generator**
- Go -> ^1.18.3

**Counter**
- C++ -> ^11
- OpenMPI -> ^4.0.3

**Make**
- make -> ^4.2.1

# Installation

## Ubuntu 20.04

Öncelikle güncellemelerinizi kontrol edin ve gerekli güncellemeleri gerçekleştirin.

    sudo apt-get update && sudo apt-get upgrade

### C++

C++ 20.04 üzerinde mevcut gelmektedir. Aşağıdaki komut yardımıyla kontrol edebilirsiniz.

    g++ --version

Eğer hata alıyorsanız aşağıdaki komut yardımıyla kurulum gerçekleştirebilir ve tekrar kontrol edebilirsiniz.

    sudo apt-get install build-essential

 Aşağıdaki komut yardımıyla MPI kütüphanesini bilgisayarınıza kurabilirsiniz.

    sudo apt-get install mpich openmpi-bin

Kurulum başarıyla tamamlandıktan sonra aşağıdaki komut ile test edebilirsiniz.

    mpic++ --version

### Go

Öncelikle `/tmp` dizinine geçin ve daha sonrasında 1.18.* versiyonunu makinenize indirin.

    cd /tmp && wget https://go.dev/dl/go1.18.3.linux-amd64.tar.gz

İndirme başarıyla gerçekleştikten sonra varsa eski go dosyalarını silin ve yenisini `/usr/local` dizinine kopyalayın

    sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.18.3.linux-amd64.tar.gz

Kopyalama gerçekleştikten sonra dizini `$PATH` içerisine dahil edin. Bunun kalıcı olması için kullandığınız shellin `rc` dosyasına yapıştırın.

- Bash için dosya dizini: $HOME/.bashrc
- Zsh için dosya dizini: $HOME/.zshrc

    export PATH=$PATH:/usr/local/go/bin

Artık aşağıdaki komut ile shell üzerinden go kullanılabilir olduğunu aşağıdaki komut ile test edebilirsiniz.

    go version

### Make

Make 20.04 üzerinde varsayılan olarak kurulu gelmektedir. Eğer ki kuruluysa aşağıdaki komut ile kontrol edebilirsiniz.

    make --version

Eğer ki hata alıyorsanız aşağıdaki komut yardımıyla kurulumu gerçekleştirip tekrar kontrol edebilirsiniz.

    sudo apt-get install make

# Makefile

Makefile dosyası size kolay kullanım sağlayan bir CLI komut seti sunar.

Kullanılabilir değişkenler:

- DATASET_SIZE -> min= 1, max=1M, default=50K
- MAX_STR_LEN -> min=2, max=100, default=10
- MIN_STR_LEN -> min=1, max=100, default=2
- GENERATED_FILE_PATH -> Boş olamaz, doğru bir dosya yolu olmalı

**WORLD_SIZE**

Kullanılacak process sayısını belirtir. Hem Generator hem de Counter için geçerlidir. 

**Kısıtlar**: 1 <= X <= Makinenizin sınırı, varsayılan değer=1

- DATASET_SIZE değişkeninden küçük olamaz.

**DATASET_SIZE**

Üretilmek istenen kelime sayısını belirtir. Yalnızca Generator tarafından kullanılan bir değişkendir.

**Kısıtlar**: 1 <= X <= 1000000, varsayılan değer=50000

- WORLD_SIZE değişkeninden büyük olamaz.

**MAX_STR_LEN**

Üretilecek kelimelerin maksimum sahip olabileceği  uzunluğunu belirtir. Yalnızca Generator tarafından kullanılan bir değişkendir.

**Kısıtlar**: 2 <= X <= 100, varsayılan değer=10

- MIN_STR_LEN değişkeninden küçük olamaz.


**MIN_STR_LEN**

Üretilecek kelimelerin minimum sahip olabileceği  uzunluğunu belirtir. Yalnızca Generator tarafından kullanılan bir değişkendir.

**Kısıtlar**: 1 <= X <= 100, varsayılan değer=2

- MAX_STR_LEN değişkeninden büyük olamaz.

**GENERATED_FILE_PATH**

Kelimelerin sayılması istenen dosyanın yolunu özel olarak belirtmek için kullanılabilir. Opsiyoneldir. Default değeri Generator tarafından üretilmiş olan en son dosyadır. Yalnızca Counter tarafından kullanılan bir değişkendir.

**Kısıtlar**: X > 0, varsayılan değer=son üretilen dosya

# RUN

Her iki uygulamayı da son hale göre derlemek ve varsayılan değerler ile çalıştırmak isterseniz aşağıdaki komutu deponun root dizininde çalıştırabilirsiniz.

    make

Uygulamaları derledikten sonra özel değerler ile çalıştırmak için

    make MAX_STR_LEN=15 MIN_STR_LEN=5 WORLD_SIZE=10 DATASET_SIZE=20

Uygulamaları özel değerler ile çalıştırmak için(önceden derlenmiş olmalı)

    make runner MAX_STR_LEN=15 MIN_STR_LEN=5 WORLD_SIZE=10 DATASET_SIZE=20

Uygulamaları sadece derlemek için

    make builder

Sadece Generator uygulamasını derlemek için

    make build_generator

Sadece Counter uygulamasını derlemek için

    make build_counter

Sadece Generator uygulamasını çalıştırmak için(önceden derlenmiş olmalı)

    make run_generator MAX_STR_LEN=15 MIN_STR_LEN=5 WORLD_SIZE=10 DATASET_SIZE=20

Sadece Generator uygulamasını derlemek ve çalıştırmak için

    make BR_generator MAX_STR_LEN=15 MIN_STR_LEN=5 WORLD_SIZE=10 DATASET_SIZE=20

Sadece Counter uygulamasını çalıştırmak için(önceden derlenmiş olmalı)

    make run_counter WORLD_SIZE=10

Sadece Counter uygulamasını özel bir dosya yolu vererek çalıştırmak için(önceden derlenmiş olmalı)

    make run_counter WORLD_SIZE=5 GENERATED_FILE_PATH=$HOME/github/MPI-words/generator/results/2022_25_05-17_42_11.txt

Sadece Counter uygulamasını derlemek ve çalıştırmak için

    make BR_counter WORLD_SIZE=10