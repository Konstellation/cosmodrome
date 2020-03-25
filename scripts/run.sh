rm /usr/local/bin/konstellation
rm /usr/local/bin/konstellationlcd
rm /usr/local/bin/konstellationcli

tar -xvzf linux_amd64.tar.gz

cp linux_amd64/konstellationcli /usr/local/bin
cp linux_amd64/konstellationlcd /usr/local/bin
cp linux_amd64/konstellation /usr/local/bin

screen -dmS kn konstellation start

ps ax | grep konstellation