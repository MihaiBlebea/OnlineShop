<template>
    <div class="container my-5">
        <div class="row">
            <div class="col-md-6">
                <ProductTotal :products="products" />

                <Product v-for="(product, index) in orderedProducts" 
                    :key="index" 
                    :title="product.title" 
                    :price="product.price"
                    :image="product.filename"
                    :quantity="product.quantity"
                    :description="product.description" />
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios'
import Product from '~/components/Product.vue'
import ProductTotal from '~/components/ProductTotal.vue'

export default {
    components: {
        Product,
        ProductTotal
    },
    asyncData: async function()
    {
        try {
            let products = await axios.get('/api/products?detailed=true')
            return {
                products: products.data
            }
        } catch(err) {
            console.error(err)
            return {
                products: []
            }
        }
    },
    data: function()
    {
        return {
            products: []
        }
    },
    computed: {
        orderedProducts: function()
        {
            if(this.products.length === 0)
            {
                return []
            }
            return this.products.sort((a, b)=> { 
                return b.title - a.title
            })
        }
    },
    methods: {
        fetchProducts: async function()
        {
            try {
                let products = await axios.get('/api/products?detailed=true')
                return products.data
            } catch(err) {
                console.error(err)
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
            this.products = await this.fetchProducts()
            await this.sleep(3000)
        }
    }
}
</script>

