const redis = require("redis")

const port = 6379
const host = "redis"

const connect = ()=> {
    var client = redis.createClient(port, host)
    return client
}

const getProductList = ()=> {
    return new Promise((resolve, reject)=> {
        client = connect()
        client.smembers('products', (err, list)=> {
            if(err) {
                return reject(err)
            }
            return resolve(list)
        })
    })
}

const getProduct = (key)=> {
    return new Promise((resolve, reject)=> {
        client.hgetall(key, (err, product)=> {
            if(err)
            {
                return reject(err)
            }
            return resolve(product)
        })
    })
}

const getProducts = async (list)=> {
    let products = []
    for(let i = 0; i < list.length; i++)
    {
        let product = await getProduct(list[i])
        products.push(product)
    }
    return products
}

const keyToString = (key)=> {
    return key.split(':').join(' ')
}

const toKey = (raw)=> {
    return raw.split(' ').join(':')
}

module.exports = {
    getProductList,
    getProducts,
    keyToString,
    toKey
}