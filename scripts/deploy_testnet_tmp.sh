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