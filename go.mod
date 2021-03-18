module github.com/John-Tonny/lnd


replace github.com/John-Tonny/lnd/ticker => ./ticker

replace github.com/John-Tonny/lnd/queue => ./queue

replace github.com/John-Tonny/lnd/cert => ./cert

replace github.com/John-Tonny/lnd/clock => ./clock

replace git.schwanenlied.me/yawning/bsaes.git => github.com/Yawning/bsaes v0.0.0-20180720073208-c0276d75487e

go 1.15
