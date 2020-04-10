const express = require('express')
const consola = require('consola')
const { Nuxt, Builder } = require('nuxt')
const app = express()

const { getProductList, getProducts, keyToString } = require("./redis")

// Import and Set Nuxt.js options
const config = require('../nuxt.config.js')
config.dev = process.env.NODE_ENV !== 'production'

async function start() {
    // Init Nuxt.js
    const nuxt = new Nuxt(config)

    const { host, port } = nuxt.options.server

    await nuxt.ready()
    // Build only in dev mode
    if (config.dev) {
        const builder = new Builder(nuxt)
        await builder.build()
    }

    app.get("/api/products", async (request, response) => {
        try {
            // if detailed query is set
            detailed = request.query.detailed
            console.log('detailed', detailed)
            let list = await getProductList()

            if(!detailed || detailed === 'false')
            {
                list = list.map((product)=> {
                    return keyToString(product)
                })
                return response.json(list)
            }
            let products = await getProducts(list)
            return response.json(products)
        } catch(err) {
            console.error(err)
            return response.json([])
        }
    })

    // Give nuxt middleware to express
    app.use(nuxt.render)

    // Listen the server
    app.listen(port, host)
    consola.ready({
        message: `Server listening on http://${host}:${port}`,
        badge: true
    })
}
start()
