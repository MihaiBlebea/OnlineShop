<template>
    <div class="alert alert-primary" role="alert">
        <span>Total: {{ productCount }} | Value: £{{ productValueTotal }}</span>
    </div>
</template>

<script>
export default {
    props: {
        products: {
            required: false,
            default: [],
            type: Array
        }
    },
    computed: {
        productCount: function()
        {
            if(this.products.length === 0)
            {
                return 0
            }
            return this.products.length
        },
        productValueTotal: function() 
        {
            if(this.products.length === 0)
            {
                return 0
            }
            let total = 0
            this.products.forEach((product)=> {
                total += product.quantity * parseFloat(product.price)
            })
            if(total > 1000)
            {
                return `${ this.roundAmount(total / 1000) }k`
            }

            if(total > 1000000)
            {
                return `${ this.roundAmount(total / 1000000) }m`
            }

            return this.roundAmount(total)
        }
    },
    methods: {
        roundAmount: function(value)
        {
            return Math.round(value * 100) / 100
        }
    }
}
</script>