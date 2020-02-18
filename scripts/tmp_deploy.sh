echo 95.216.212.119;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@95.216.212.119 rm -rdf /root/.konstellation;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@95.216.212.119 rm -rdf /root/.konstellationcli;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@95.216.212.119 rm -rdf /root/.konstellationlcd;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@95.216.212.119 "ps ax | grep konstellation | awk '{print \$1}' | xargs kill";
     scp -i ~/Documents/.ssh/id_rsa.pub -r ./testnet/testnode-1/.konstellation root@95.216.212.119:/root;
     scp -i ~/Documents/.ssh/id_rsa.pub -r ./testnet/testnode-1/.konstellationcli root@95.216.212.119:/root;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@95.216.212.119 "screen -dmS kn konstellation start";
     echo 
echo 95.216.215.127;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@95.216.215.127 rm -rdf /root/.konstellation;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@95.216.215.127 rm -rdf /root/.konstellationcli;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@95.216.215.127 rm -rdf /root/.konstellationlcd;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@95.216.215.127 "ps ax | grep konstellation | awk '{print \$1}' | xargs kill";
     scp -i ~/Documents/.ssh/id_rsa.pub -r ./testnet/testnode-2/.konstellation root@95.216.215.127:/root;
     scp -i ~/Documents/.ssh/id_rsa.pub -r ./testnet/testnode-2/.konstellationcli root@95.216.215.127:/root;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@95.216.215.127 "screen -dmS kn konstellation start";
     echo 
echo 95.216.210.171;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@95.216.210.171 rm -rdf /root/.konstellation;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@95.216.210.171 rm -rdf /root/.konstellationcli;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@95.216.210.171 rm -rdf /root/.konstellationlcd;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@95.216.210.171 "ps ax | grep konstellation | awk '{print \$1}' | xargs kill";
     scp -i ~/Documents/.ssh/id_rsa.pub -r ./testnet/testnode-3/.konstellation root@95.216.210.171:/root;
     scp -i ~/Documents/.ssh/id_rsa.pub -r ./testnet/testnode-3/.konstellationcli root@95.216.210.171:/root;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@95.216.210.171 "screen -dmS kn konstellation start";
     echo 
echo 95.216.207.219;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@95.216.207.219 rm -rdf /root/.konstellation;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@95.216.207.219 rm -rdf /root/.konstellationcli;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@95.216.207.219 rm -rdf /root/.konstellationlcd;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@95.216.207.219 "ps ax | grep konstellation | awk '{print \$1}' | xargs kill";
     scp -i ~/Documents/.ssh/id_rsa.pub -r ./testnet/testnode-4/.konstellation root@95.216.207.219:/root;
     scp -i ~/Documents/.ssh/id_rsa.pub -r ./testnet/testnode-4/.konstellationcli root@95.216.207.219:/root;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@95.216.207.219 "screen -dmS kn konstellation start";
     echo 
echo 188.166.20.172;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@188.166.20.172 rm -rdf /root/.konstellation;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@188.166.20.172 rm -rdf /root/.konstellationcli;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@188.166.20.172 rm -rdf /root/.konstellationlcd;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@188.166.20.172 "ps ax | grep konstellation | awk '{print \$1}' | xargs kill";
     scp -i ~/Documents/.ssh/id_rsa.pub -r ./testnet/testnode-5/.konstellation root@188.166.20.172:/root;
     scp -i ~/Documents/.ssh/id_rsa.pub -r ./testnet/testnode-5/.konstellationcli root@188.166.20.172:/root;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@188.166.20.172 "screen -dmS kn konstellation start";
     echo 
echo 167.172.37.36;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@167.172.37.36 rm -rdf /root/.konstellation;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@167.172.37.36 rm -rdf /root/.konstellationcli;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@167.172.37.36 rm -rdf /root/.konstellationlcd;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@167.172.37.36 "ps ax | grep konstellation | awk '{print \$1}' | xargs kill";
     scp -i ~/Documents/.ssh/id_rsa.pub -r ./testnet/testnode-6/.konstellation root@167.172.37.36:/root;
     scp -i ~/Documents/.ssh/id_rsa.pub -r ./testnet/testnode-6/.konstellationcli root@167.172.37.36:/root;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@167.172.37.36 "screen -dmS kn konstellation start";
     echo 
echo 167.172.45.243;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@167.172.45.243 rm -rdf /root/.konstellation;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@167.172.45.243 rm -rdf /root/.konstellationcli;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@167.172.45.243 rm -rdf /root/.konstellationlcd;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@167.172.45.243 "ps ax | grep konstellation | awk '{print \$1}' | xargs kill";
     scp -i ~/Documents/.ssh/id_rsa.pub -r ./testnet/testnode-7/.konstellation root@167.172.45.243:/root;
     scp -i ~/Documents/.ssh/id_rsa.pub -r ./testnet/testnode-7/.konstellationcli root@167.172.45.243:/root;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@167.172.45.243 "screen -dmS kn konstellation start";
     echo 
echo 167.172.43.125;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@167.172.43.125 rm -rdf /root/.konstellation;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@167.172.43.125 rm -rdf /root/.konstellationcli;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@167.172.43.125 rm -rdf /root/.konstellationlcd;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@167.172.43.125 "ps ax | grep konstellation | awk '{print \$1}' | xargs kill";
     scp -i ~/Documents/.ssh/id_rsa.pub -r ./testnet/testnode-8/.konstellation root@167.172.43.125:/root;
     scp -i ~/Documents/.ssh/id_rsa.pub -r ./testnet/testnode-8/.konstellationcli root@167.172.43.125:/root;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@167.172.43.125 "screen -dmS kn konstellation start";
     echo 
echo 188.166.37.121;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@188.166.37.121 rm -rdf /root/.konstellation;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@188.166.37.121 rm -rdf /root/.konstellationcli;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@188.166.37.121 rm -rdf /root/.konstellationlcd;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@188.166.37.121 "ps ax | grep konstellation | awk '{print \$1}' | xargs kill";
     scp -i ~/Documents/.ssh/id_rsa.pub -r ./testnet/testnode-9/.konstellation root@188.166.37.121:/root;
     scp -i ~/Documents/.ssh/id_rsa.pub -r ./testnet/testnode-9/.konstellationcli root@188.166.37.121:/root;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@188.166.37.121 "screen -dmS kn konstellation start";
     echo 
echo 188.166.68.78;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@188.166.68.78 rm -rdf /root/.konstellation;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@188.166.68.78 rm -rdf /root/.konstellationcli;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@188.166.68.78 rm -rdf /root/.konstellationlcd;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@188.166.68.78 "ps ax | grep konstellation | awk '{print \$1}' | xargs kill";
     scp -i ~/Documents/.ssh/id_rsa.pub -r ./testnet/testnode-10/.konstellation root@188.166.68.78:/root;
     scp -i ~/Documents/.ssh/id_rsa.pub -r ./testnet/testnode-10/.konstellationcli root@188.166.68.78:/root;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@188.166.68.78 "screen -dmS kn konstellation start";
     echo 
