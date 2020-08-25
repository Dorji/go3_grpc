# Обновляем системные зависимости
sudo apt-get update
sudo apt-get -y upgrade

# Если уже есть директория внутри /usr/local то выполнить
cd /usr/local 
sudo rm -r -f go

# Убедиться, что в /usr/local больше нет такой диреткории

# После чего скачать исходники и разархивировать по нужному адресу
wget https://golang.org/dl/go1.15.linux-amd64.tar.gz
sudo tar -xvf go1.15.linux-amd64.tar.gz
sudo mv go /usr/local

# После чего убедиться, что в /usr/local появилась новая директория с именем go

# Затем открыть файл .bashrc
sudo nano ~/.bashrc

# В файле .basrhc в самом низу приписать
GOROOT=/usr/local/go
GOPATH=$HOME/go
PATH=$GOPATH/bin:$GOROOT/bin:$PATH

# Сохранить файл при помощи Ctrl+O и выйти

# Перезапустить терминал и выполнить
source ~/.profile
go version

# В случае если все ок, увидим версию компилятора, если проблемы - 
# внимательно посмотреть, не опечатались ли в .bashrc