<template>
    <div class="container my-5">
        <div class="row mb-3">
            <div class="col-md-6">
                <Products :data="products" />
            </div>
            <div class="col-md-6">
                <Customers :data="customers" class="mb-1"/>
                <Balance :balance="balance" />
            </div>
        </div>
        <div class="row justify-content-center">
            <div class="col-md-8">
               <Events :stream="events" />
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios'
import Events from '~/components/Events.vue'
import Products from'~/components/Products.vue'
import Customers from '~/components/Customers.vue'
import Balance from '~/components/Balance.vue'

export default {
    components: {
        Events,
        Products,
        Customers,
        Balance
    },
    data: function()
    {
        return {
            events: [],
            products: [],
            customers: []
        }
    },
    methods: {
        fetchEventStream: async function()
        {
            try {
                let { data } = await axios.get('/api/event-stream')
                let stream = data.map((event)=> {
                    return JSON.parse(event)
                })
                return stream
            } catch(err) {
                console.log(err)
                return []
            }
        },
        fetchProducts: async function()
        {
            try {
                let { data } = await axios.get(`http://${ this.$env.NUXT_ENV_SHOP_URL }/products`)
                return data
            } catch(err) {
                console.log(err)
                return []
            }
        },
        fetchCustomers: async function()
        {
            try {
                let { data } = await axios.get(`http://${ this.$env.NUXT_ENV_CUSTOMER_URL }/customers`)
                return data
            } catch(err) {
                console.log(err)
                return []
            }
        },
        fetchBalance: async function()
        {
            try {
                let { data } = await axios.get(`http://${ this.$env.NUXT_ENV_SHOP_URL }/account/balance`)
                return data
            } catch(err) {
                console.log(err)
                return []
            }
        },
        sleep: function(ms) 
        {
            return new Promise(resolve => setTimeout(resolve, ms))
        }
    },
    mounted: async function()
    {
        while(true) {
            this.events = await this.fetchEventStream()
            this.customers = await this.fetchCustomers()
            this.products = await this.fetchProducts()
            this.balance = await this.fetchBalance()
            await this.sleep(10 * 1000)
        }
    }
}
</script>

