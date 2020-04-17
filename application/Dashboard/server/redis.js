const redis = require("redis")

const connect = ()=> {
    var client = redis.createClient(process.env.NUXT_ENV_REDIS_PORT, process.env.NUXT_ENV_REDIS_HOST)
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