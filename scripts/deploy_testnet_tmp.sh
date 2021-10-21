
jq -r '
    .validators[] |
    "echo \(.ip);
     ssh -i ~/.ssh/bldg bldg@\(.ip) \"rm -rdf /root/.knstld\";
     scp -i ~/.ssh/bldg -r ./testnet/\(.name)/.knstld root@\(.ip):/root;
     echo "
    ' ./config/testnet.json
echo node6.konstellation.tech;
     ssh -i ~/.ssh/bldg bldg@node6.konstellation.tech "sudo rm -rdf /root/.knstld";
     scp -i ~/.ssh/bldg -r ./testnet/testnode-1/.knstld bldg@node6.konstellation.tech:/home/bldg;
     echo
echo node7.konstellation.tech;
     ssh -i ~/.ssh/bldg bldg@node7.konstellation.tech "sudo rm -rdf /root/.knstld";
     scp -i ~/.ssh/bldg -r ./testnet/testnode-2/.knstld bldg@node7.konstellation.tech:/home/bldg;
     echo
echo node8.konstellation.tech;
     ssh -i ~/.ssh/bldg bldg@node8.konstellation.tech "sudo rm -rdf /root/.knstld";
     scp -i ~/.ssh/bldg -r ./testnet/testnode-3/.knstld bldg@node8.konstellation.tech:/home/bldg;
     echo

     ssh dglee@node6.konstellation.tech "sudo rm -rdf /root/.knstld";
     ssh dglee@node7.konstellation.tech "sudo rm -rdf /root/.knstld";
     ssh dglee@node8.konstellation.tech "sudo rm -rdf /root/.knstld";
     scp -r ./testnet/testnode-1/.knstld dglee@node6.konstellation.tech:/root;
     scp -r ./testnet/testnode-2/.knstld dglee@node7.konstellation.tech:/root;
     scp -r ./testnet/testnode-3/.knstld dglee@node8.konstellation.tech:/root;

#echo node1;
#     ssh -i ~/.ssh/bldg bldg@node1.konstellation.tech "ps ax | grep konstellation | awk '{print \$1}' | xargs kill";
#     ssh -i ~/.ssh/bldg bldg@node1.konstellation.tech "konstellation unsafe-reset-all";
#     ssh -i ~/.ssh/bldg bldg@node1.konstellation.tech "sudo rm -rdf /root/.konstellation";
#     ssh -i ~/.ssh/bldg bldg@node1.konstellation.tech "sudo rm -rdf /root/.konstellationcli";
#     ssh -i ~/.ssh/bldg bldg@node1.konstellation.tech "sudo rm -rdf /root/.konstellationlcd";
#     scp -i ~/.ssh/bldg -r ./testnet/testnode-1/.konstellation bldg@node1.konstellation.tech:/home/bldg;
#     scp -i ~/.ssh/bldg -r ./testnet/testnode-1/.konstellationcli bldg@node1.konstellation.tech:/home/bldg;
#     ssh -i ~/.ssh/bldg bldg@node1.konstellation.tech "screen -dmS kn konstellation start";
#     ssh -i ~/.ssh/bldg bldg@node1.konstellation.tech "screen -dmS klcd konstellationlcd rest-server --chain-id darchub --laddr tcp://0.0.0.0:1317";
#     echo
#echo node2;
#     ssh -i ~/.ssh/bldg bldg@node2.konstellation.tech "ps ax | grep konstellation | awk '{print \$1}' | xargs kill";
#     ssh -i ~/.ssh/bldg bldg@node2.konstellation.tech "konstellation unsafe-reset-all";
#     ssh -i ~/.ssh/bldg bldg@node2.konstellation.tech "sudo rm -rdf /root/.konstellation";
#     ssh -i ~/.ssh/bldg bldg@node2.konstellation.tech "sudo rm -rdf /root/.konstellationcli";
#     ssh -i ~/.ssh/bldg bldg@node2.konstellation.tech "sudo rm -rdf /root/.konstellationlcd";
#     scp -i ~/.ssh/bldg -r ./testnet/testnode-2/.konstellation bldg@node2.konstellation.tech:/home/bldg;
#     scp -i ~/.ssh/bldg -r ./testnet/testnode-2/.konstellationcli bldg@node2.konstellation.tech:/home/bldg;
#     ssh -i ~/.ssh/bldg bldg@node2.konstellation.tech "screen -dmS kn konstellation start";
#     echo
#echo node3;
#     ssh -i ~/.ssh/bldg bldg@node3.konstellation.tech "ps ax | grep konstellation | awk '{print \$1}' | xargs kill";
#     ssh -i ~/.ssh/bldg bldg@node3.konstellation.tech "konstellation unsafe-reset-all";
#     ssh -i ~/.ssh/bldg bldg@node3.konstellation.tech "sudo rm -rdf /root/.konstellation";
#     ssh -i ~/.ssh/bldg bldg@node3.konstellation.tech "sudo rm -rdf /root/.konstellationcli";
#     ssh -i ~/.ssh/bldg bldg@node3.konstellation.tech "sudo rm -rdf /root/.konstellationlcd";
#     scp -i ~/.ssh/bldg -r ./testnet/testnode-3/.konstellation bldg@node3.konstellation.tech:/home/bldg;
#     scp -i ~/.ssh/bldg -r ./testnet/testnode-3/.konstellationcli bldg@node3.konstellation.tech:/home/bldg;
#     ssh -i ~/.ssh/bldg bldg@node3.konstellation.tech "screen -dmS kn konstellation start";
#     echo

echo goz;
#     ssh -i ~/.ssh/bldg bldg@node3.konstellation.tech "ps ax | grep konstellation | awk '{print \$1}' | xargs kill";
#     ssh -i ~/.ssh/bldg bldg@node3.konstellation.tech "konstellation unsafe-reset-all";
#     ssh -i ~/.ssh/bldg bldg@node3.konstellation.tech "sudo rm -rdf /root/.konstellation";
#     ssh -i ~/.ssh/bldg bldg@node3.konstellation.tech "sudo rm -rdf /root/.konstellationcli";
#     ssh -i ~/.ssh/bldg bldg@node3.konstellation.tech "sudo rm -rdf /root/.konstellationlcd";
#     scp -i ~/.ssh/bldg -r ./testnet/config bldg@goz.konstellation.tech:/home/bldg/.konstellation/;
#     scp -i ~/.ssh/bldg -r ./testnet/testnode-3/.konstellationcli bldg@node3.konstellation.tech:/home/bldg;
#     ssh -i ~/.ssh/bldg bldg@node3.konstellation.tech "screen -dmS kn konstellation start";
#     echo