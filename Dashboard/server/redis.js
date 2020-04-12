const redis = require("redis")

const connect = ()=> {
    var client = redis.createClient(6379, 'redis')
    return client
}

const getRedisStream = ()=> {
    return new Promise((resolve, reject)=> {
        client = connect()
        client.smembers('stream', (err, list)=> {
            client.end(true)
            if(err) {
                return reject(err)
            }
            return resolve(list)
        })
    })
}

module.exports = {
    getRedisStream
}