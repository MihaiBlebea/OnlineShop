const redis = require("redis")

const port = process.env.NUXT_ENV_REDIS_PORT
const host = process.env.NUXT_ENV_REDIS_HOST

const connect = ()=> {
    var client = redis.createClient(port, host)
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