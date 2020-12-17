screen -S micro-frontends-core -X quit &>/dev/null
screen -dmS micro-frontends-core npm -prefix ~/micro-frontends/src/micro-frontends-core run serve

screen -S micro-frontends-vue -X quit &>/dev/null
screen -dmS micro-frontends-vue npm -prefix ~/micro-frontends/src/micro-frontends-vue run serve

screen -S mock-service -X quit &>/dev/null
screen -dmS mock-service ~/micro-frontends/src/mock-service/golang/main
